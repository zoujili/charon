// Code generated by mockery v1.0.0. DO NOT EDIT.

package modelmock

import context "context"
import mock "github.com/stretchr/testify/mock"
import model "github.com/piotrkowalczuk/charon/internal/model"

// GroupPermissionsProvider is an autogenerated mock type for the GroupPermissionsProvider type
type GroupPermissionsProvider struct {
	mock.Mock
}

// Insert provides a mock function with given fields: _a0, _a1
func (_m *GroupPermissionsProvider) Insert(_a0 context.Context, _a1 *model.GroupPermissionsEntity) (*model.GroupPermissionsEntity, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *model.GroupPermissionsEntity
	if rf, ok := ret.Get(0).(func(context.Context, *model.GroupPermissionsEntity) *model.GroupPermissionsEntity); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.GroupPermissionsEntity)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *model.GroupPermissionsEntity) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
