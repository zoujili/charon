package main

import (
	"github.com/piotrkowalczuk/charon"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type setUserPermissionsHandler struct {
	*handler
}

func (suph *setUserPermissionsHandler) handle(ctx context.Context, req *charon.SetUserPermissionsRequest) (*charon.SetUserPermissionsResponse, error) {
	act, err := suph.retrieveActor(ctx)
	if err != nil {
		return nil, err
	}

	if err = suph.firewall(req, act); err != nil {
		return nil, err
	}

	created, removed, err := suph.repository.user.SetPermissions(req.UserId, charon.NewPermissions(req.Permissions...)...)
	if err != nil {
		return nil, err
	}

	return &charon.SetUserPermissionsResponse{
		Created:   created,
		Removed:   removed,
		Untouched: untouched(int64(len(req.Permissions)), created, removed),
	}, nil
}

func (suph *setUserPermissionsHandler) firewall(req *charon.SetUserPermissionsRequest, act *actor) error {
	if act.user.IsSuperuser {
		return nil
	}
	if act.permissions.Contains(charon.UserPermissionCanCreate, charon.UserPermissionCanDelete) {
		return nil
	}

	return grpc.Errorf(codes.PermissionDenied, "charond: user permissions cannot be set, missing permission")
}
