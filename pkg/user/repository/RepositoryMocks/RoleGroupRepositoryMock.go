// Code generated by mockery v2.14.0. DO NOT EDIT.

package repomock

import (
	pg "github.com/go-pg/pg"
	mock "github.com/stretchr/testify/mock"

	repository "github.com/devtron-labs/devtron/pkg/user/repository"
)

// RoleGroupRepository is an autogenerated mock type for the RoleGroupRepository type
type RoleGroupRepository struct {
	mock.Mock
}

// CreateRoleGroup provides a mock function with given fields: model, tx
func (_m *RoleGroupRepository) CreateRoleGroup(model *repository.RoleGroup, tx *pg.Tx) (*repository.RoleGroup, error) {
	ret := _m.Called(model, tx)

	var r0 *repository.RoleGroup
	if rf, ok := ret.Get(0).(func(*repository.RoleGroup, *pg.Tx) *repository.RoleGroup); ok {
		r0 = rf(model, tx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*repository.RoleGroup)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*repository.RoleGroup, *pg.Tx) error); ok {
		r1 = rf(model, tx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateRoleGroupRoleMapping provides a mock function with given fields: model, tx
func (_m *RoleGroupRepository) CreateRoleGroupRoleMapping(model *repository.RoleGroupRoleMapping, tx *pg.Tx) (*repository.RoleGroupRoleMapping, error) {
	ret := _m.Called(model, tx)

	var r0 *repository.RoleGroupRoleMapping
	if rf, ok := ret.Get(0).(func(*repository.RoleGroupRoleMapping, *pg.Tx) *repository.RoleGroupRoleMapping); ok {
		r0 = rf(model, tx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*repository.RoleGroupRoleMapping)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*repository.RoleGroupRoleMapping, *pg.Tx) error); ok {
		r1 = rf(model, tx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteRoleGroupRoleMapping provides a mock function with given fields: model, tx
func (_m *RoleGroupRepository) DeleteRoleGroupRoleMapping(model *repository.RoleGroupRoleMapping, tx *pg.Tx) (bool, error) {
	ret := _m.Called(model, tx)

	var r0 bool
	if rf, ok := ret.Get(0).(func(*repository.RoleGroupRoleMapping, *pg.Tx) bool); ok {
		r0 = rf(model, tx)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*repository.RoleGroupRoleMapping, *pg.Tx) error); ok {
		r1 = rf(model, tx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteRoleGroupRoleMappingByRoleId provides a mock function with given fields: roleId, tx
func (_m *RoleGroupRepository) DeleteRoleGroupRoleMappingByRoleId(roleId int, tx *pg.Tx) error {
	ret := _m.Called(roleId, tx)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, *pg.Tx) error); ok {
		r0 = rf(roleId, tx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllRoleGroup provides a mock function with given fields:
func (_m *RoleGroupRepository) GetAllRoleGroup() ([]*repository.RoleGroup, error) {
	ret := _m.Called()

	var r0 []*repository.RoleGroup
	if rf, ok := ret.Get(0).(func() []*repository.RoleGroup); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*repository.RoleGroup)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetConnection provides a mock function with given fields:
func (_m *RoleGroupRepository) GetConnection() *pg.DB {
	ret := _m.Called()

	var r0 *pg.DB
	if rf, ok := ret.Get(0).(func() *pg.DB); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pg.DB)
		}
	}

	return r0
}

// GetRoleGroupById provides a mock function with given fields: id
func (_m *RoleGroupRepository) GetRoleGroupById(id int32) (*repository.RoleGroup, error) {
	ret := _m.Called(id)

	var r0 *repository.RoleGroup
	if rf, ok := ret.Get(0).(func(int32) *repository.RoleGroup); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*repository.RoleGroup)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int32) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRoleGroupByName provides a mock function with given fields: name
func (_m *RoleGroupRepository) GetRoleGroupByName(name string) (*repository.RoleGroup, error) {
	ret := _m.Called(name)

	var r0 *repository.RoleGroup
	if rf, ok := ret.Get(0).(func(string) *repository.RoleGroup); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*repository.RoleGroup)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRoleGroupListByCasbinNames provides a mock function with given fields: name
func (_m *RoleGroupRepository) GetRoleGroupListByCasbinNames(name []string) ([]*repository.RoleGroup, error) {
	ret := _m.Called(name)

	var r0 []*repository.RoleGroup
	if rf, ok := ret.Get(0).(func([]string) []*repository.RoleGroup); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*repository.RoleGroup)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRoleGroupListByName provides a mock function with given fields: name
func (_m *RoleGroupRepository) GetRoleGroupListByName(name string) ([]*repository.RoleGroup, error) {
	ret := _m.Called(name)

	var r0 []*repository.RoleGroup
	if rf, ok := ret.Get(0).(func(string) []*repository.RoleGroup); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*repository.RoleGroup)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRoleGroupListByNames provides a mock function with given fields: groupNames
func (_m *RoleGroupRepository) GetRoleGroupListByNames(groupNames []string) ([]*repository.RoleGroup, error) {
	ret := _m.Called(groupNames)

	var r0 []*repository.RoleGroup
	if rf, ok := ret.Get(0).(func([]string) []*repository.RoleGroup); ok {
		r0 = rf(groupNames)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*repository.RoleGroup)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]string) error); ok {
		r1 = rf(groupNames)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRoleGroupRoleMapping provides a mock function with given fields: model
func (_m *RoleGroupRepository) GetRoleGroupRoleMapping(model int32) (*repository.RoleGroupRoleMapping, error) {
	ret := _m.Called(model)

	var r0 *repository.RoleGroupRoleMapping
	if rf, ok := ret.Get(0).(func(int32) *repository.RoleGroupRoleMapping); ok {
		r0 = rf(model)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*repository.RoleGroupRoleMapping)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int32) error); ok {
		r1 = rf(model)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRoleGroupRoleMappingByRoleGroupId provides a mock function with given fields: roleGroupId
func (_m *RoleGroupRepository) GetRoleGroupRoleMappingByRoleGroupId(roleGroupId int32) ([]*repository.RoleGroupRoleMapping, error) {
	ret := _m.Called(roleGroupId)

	var r0 []*repository.RoleGroupRoleMapping
	if rf, ok := ret.Get(0).(func(int32) []*repository.RoleGroupRoleMapping); ok {
		r0 = rf(roleGroupId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*repository.RoleGroupRoleMapping)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int32) error); ok {
		r1 = rf(roleGroupId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRoleGroupRoleMappingByRoleGroupIds provides a mock function with given fields: roleGroupIds
func (_m *RoleGroupRepository) GetRoleGroupRoleMappingByRoleGroupIds(roleGroupIds []int32) ([]*repository.RoleModel, error) {
	ret := _m.Called(roleGroupIds)

	var r0 []*repository.RoleModel
	if rf, ok := ret.Get(0).(func([]int32) []*repository.RoleModel); ok {
		r0 = rf(roleGroupIds)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*repository.RoleModel)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]int32) error); ok {
		r1 = rf(roleGroupIds)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRolesByGroupCasbinName provides a mock function with given fields: groupName
func (_m *RoleGroupRepository) GetRolesByGroupCasbinName(groupName string) ([]*repository.RoleModel, error) {
	ret := _m.Called(groupName)

	var r0 []*repository.RoleModel
	if rf, ok := ret.Get(0).(func(string) []*repository.RoleModel); ok {
		r0 = rf(groupName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*repository.RoleModel)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(groupName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRolesByGroupNames provides a mock function with given fields: groupNames
func (_m *RoleGroupRepository) GetRolesByGroupNames(groupNames []string) ([]*repository.RoleModel, error) {
	ret := _m.Called(groupNames)

	var r0 []*repository.RoleModel
	if rf, ok := ret.Get(0).(func([]string) []*repository.RoleModel); ok {
		r0 = rf(groupNames)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*repository.RoleModel)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]string) error); ok {
		r1 = rf(groupNames)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateRoleGroup provides a mock function with given fields: model, tx
func (_m *RoleGroupRepository) UpdateRoleGroup(model *repository.RoleGroup, tx *pg.Tx) (*repository.RoleGroup, error) {
	ret := _m.Called(model, tx)

	var r0 *repository.RoleGroup
	if rf, ok := ret.Get(0).(func(*repository.RoleGroup, *pg.Tx) *repository.RoleGroup); ok {
		r0 = rf(model, tx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*repository.RoleGroup)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*repository.RoleGroup, *pg.Tx) error); ok {
		r1 = rf(model, tx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewRoleGroupRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewRoleGroupRepository creates a new instance of RoleGroupRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRoleGroupRepository(t mockConstructorTestingTNewRoleGroupRepository) *RoleGroupRepository {
	mock := &RoleGroupRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
