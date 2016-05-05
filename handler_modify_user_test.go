package charon

import (
	"testing"

	"github.com/piotrkowalczuk/ntypes"
)

func TestModifyUserHandler_Firewall(t *testing.T) {
	success := []struct {
		hint   string
		req    ModifyUserRequest
		entity userEntity
		user   userEntity
		perm   Permissions
	}{
		{
			hint:   "superuser should be able to degrade itself",
			req:    ModifyUserRequest{Id: 1, IsSuperuser: &ntypes.Bool{Bool: false, Valid: true}},
			entity: userEntity{ID: 1, IsSuperuser: true},
			user:   userEntity{ID: 1, IsSuperuser: true},
		},
		{
			hint:   "superuser should be able to promote an user",
			req:    ModifyUserRequest{Id: 2, IsSuperuser: &ntypes.Bool{Bool: true, Valid: true}},
			entity: userEntity{ID: 2},
			user:   userEntity{ID: 1, IsSuperuser: true},
		},
		{
			hint:   "superuser should be able to promote a staff user",
			req:    ModifyUserRequest{Id: 2, IsSuperuser: &ntypes.Bool{Bool: true, Valid: true}},
			entity: userEntity{ID: 2},
			user:   userEntity{ID: 1, IsSuperuser: true},
		},
		{
			hint:   "superuser should be able to degrade another superuser",
			req:    ModifyUserRequest{Id: 2, IsSuperuser: &ntypes.Bool{Bool: false, Valid: true}},
			entity: userEntity{ID: 2, IsSuperuser: true},
			user:   userEntity{ID: 1, IsSuperuser: true},
		},
		{
			hint:   "if user has permission to modify an user as a stranger, he should be able to do that",
			req:    ModifyUserRequest{Id: 2},
			entity: userEntity{ID: 2},
			user:   userEntity{ID: 1},
			perm:   Permissions{UserCanModifyAsStranger},
		},
		{
			hint:   "if user has permission to modify an user as an owner, he should be able to do that",
			req:    ModifyUserRequest{Id: 1},
			entity: userEntity{ID: 1},
			user:   userEntity{ID: 1},
			perm:   Permissions{UserCanModifyAsOwner},
		},
	}

	handler := &modifyUserHandler{}
	for _, args := range success {
		msg, ok := handler.firewall(&args.req, &args.entity, &actor{
			user:        &args.user,
			permissions: args.perm,
		})
		if !ok {
			t.Errorf(args.hint+", got: %s", msg)
		} else {
			t.Log(args.hint)
		}
	}
}