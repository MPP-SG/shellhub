// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	context "context"

	models "github.com/shellhub-io/shellhub/pkg/models"
	mock "github.com/stretchr/testify/mock"

	paginator "github.com/shellhub-io/shellhub/pkg/api/paginator"
)

// Store is an autogenerated mock type for the Store type
type Store struct {
	mock.Mock
}

// AddDevice provides a mock function with given fields: ctx, d
func (_m *Store) AddDevice(ctx context.Context, d models.Device) error {
	ret := _m.Called(ctx, d)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.Device) error); ok {
		r0 = rf(ctx, d)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateFirewallRule provides a mock function with given fields: ctx, rule
func (_m *Store) CreateFirewallRule(ctx context.Context, rule *models.FirewallRule) error {
	ret := _m.Called(ctx, rule)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.FirewallRule) error); ok {
		r0 = rf(ctx, rule)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateSession provides a mock function with given fields: ctx, session
func (_m *Store) CreateSession(ctx context.Context, session models.Session) (*models.Session, error) {
	ret := _m.Called(ctx, session)

	var r0 *models.Session
	if rf, ok := ret.Get(0).(func(context.Context, models.Session) *models.Session); ok {
		r0 = rf(ctx, session)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Session)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.Session) error); ok {
		r1 = rf(ctx, session)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeactivateSession provides a mock function with given fields: ctx, uid
func (_m *Store) DeactivateSession(ctx context.Context, uid models.UID) error {
	ret := _m.Called(ctx, uid)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.UID) error); ok {
		r0 = rf(ctx, uid)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteDevice provides a mock function with given fields: ctx, uid
func (_m *Store) DeleteDevice(ctx context.Context, uid models.UID) error {
	ret := _m.Called(ctx, uid)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.UID) error); ok {
		r0 = rf(ctx, uid)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteFirewallRule provides a mock function with given fields: ctx, id
func (_m *Store) DeleteFirewallRule(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetDevice provides a mock function with given fields: ctx, uid
func (_m *Store) GetDevice(ctx context.Context, uid models.UID) (*models.Device, error) {
	ret := _m.Called(ctx, uid)

	var r0 *models.Device
	if rf, ok := ret.Get(0).(func(context.Context, models.UID) *models.Device); ok {
		r0 = rf(ctx, uid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Device)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.UID) error); ok {
		r1 = rf(ctx, uid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDeviceByMac provides a mock function with given fields: ctx, mac, tenant
func (_m *Store) GetDeviceByMac(ctx context.Context, mac string, tenant string) (*models.Device, error) {
	ret := _m.Called(ctx, mac, tenant)

	var r0 *models.Device
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *models.Device); ok {
		r0 = rf(ctx, mac, tenant)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Device)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, mac, tenant)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDeviceByName provides a mock function with given fields: ctx, name, tenant
func (_m *Store) GetDeviceByName(ctx context.Context, name string, tenant string) (*models.Device, error) {
	ret := _m.Called(ctx, name, tenant)

	var r0 *models.Device
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *models.Device); ok {
		r0 = rf(ctx, name, tenant)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Device)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, name, tenant)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDeviceByUID provides a mock function with given fields: ctx, uid, tenant
func (_m *Store) GetDeviceByUID(ctx context.Context, uid models.UID, tenant string) (*models.Device, error) {
	ret := _m.Called(ctx, uid, tenant)

	var r0 *models.Device
	if rf, ok := ret.Get(0).(func(context.Context, models.UID, string) *models.Device); ok {
		r0 = rf(ctx, uid, tenant)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Device)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.UID, string) error); ok {
		r1 = rf(ctx, uid, tenant)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetFirewallRule provides a mock function with given fields: ctx, id
func (_m *Store) GetFirewallRule(ctx context.Context, id string) (*models.FirewallRule, error) {
	ret := _m.Called(ctx, id)

	var r0 *models.FirewallRule
	if rf, ok := ret.Get(0).(func(context.Context, string) *models.FirewallRule); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.FirewallRule)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSession provides a mock function with given fields: ctx, uid
func (_m *Store) GetSession(ctx context.Context, uid models.UID) (*models.Session, error) {
	ret := _m.Called(ctx, uid)

	var r0 *models.Session
	if rf, ok := ret.Get(0).(func(context.Context, models.UID) *models.Session); ok {
		r0 = rf(ctx, uid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Session)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.UID) error); ok {
		r1 = rf(ctx, uid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStats provides a mock function with given fields: ctx
func (_m *Store) GetStats(ctx context.Context) (*models.Stats, error) {
	ret := _m.Called(ctx)

	var r0 *models.Stats
	if rf, ok := ret.Get(0).(func(context.Context) *models.Stats); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Stats)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserByTenant provides a mock function with given fields: ctx, tenant
func (_m *Store) GetUserByTenant(ctx context.Context, tenant string) (*models.User, error) {
	ret := _m.Called(ctx, tenant)

	var r0 *models.User
	if rf, ok := ret.Get(0).(func(context.Context, string) *models.User); ok {
		r0 = rf(ctx, tenant)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, tenant)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserByUsername provides a mock function with given fields: ctx, username
func (_m *Store) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	ret := _m.Called(ctx, username)

	var r0 *models.User
	if rf, ok := ret.Get(0).(func(context.Context, string) *models.User); ok {
		r0 = rf(ctx, username)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// KeepAliveSession provides a mock function with given fields: ctx, uid
func (_m *Store) KeepAliveSession(ctx context.Context, uid models.UID) error {
	ret := _m.Called(ctx, uid)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.UID) error); ok {
		r0 = rf(ctx, uid)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ListDevices provides a mock function with given fields: ctx, pagination, filters, status
func (_m *Store) ListDevices(ctx context.Context, pagination paginator.Query, filters []models.Filter, status string) ([]models.Device, int, error) {
	ret := _m.Called(ctx, pagination, filters, status)

	var r0 []models.Device
	if rf, ok := ret.Get(0).(func(context.Context, paginator.Query, []models.Filter, string) []models.Device); ok {
		r0 = rf(ctx, pagination, filters, status)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Device)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(context.Context, paginator.Query, []models.Filter, string) int); ok {
		r1 = rf(ctx, pagination, filters, status)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, paginator.Query, []models.Filter, string) error); ok {
		r2 = rf(ctx, pagination, filters, status)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ListFirewallRules provides a mock function with given fields: ctx, pagination
func (_m *Store) ListFirewallRules(ctx context.Context, pagination paginator.Query) ([]models.FirewallRule, int, error) {
	ret := _m.Called(ctx, pagination)

	var r0 []models.FirewallRule
	if rf, ok := ret.Get(0).(func(context.Context, paginator.Query) []models.FirewallRule); ok {
		r0 = rf(ctx, pagination)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.FirewallRule)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(context.Context, paginator.Query) int); ok {
		r1 = rf(ctx, pagination)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, paginator.Query) error); ok {
		r2 = rf(ctx, pagination)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ListSessions provides a mock function with given fields: ctx, pagination
func (_m *Store) ListSessions(ctx context.Context, pagination paginator.Query) ([]models.Session, int, error) {
	ret := _m.Called(ctx, pagination)

	var r0 []models.Session
	if rf, ok := ret.Get(0).(func(context.Context, paginator.Query) []models.Session); ok {
		r0 = rf(ctx, pagination)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Session)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(context.Context, paginator.Query) int); ok {
		r1 = rf(ctx, pagination)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, paginator.Query) error); ok {
		r2 = rf(ctx, pagination)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// LookupDevice provides a mock function with given fields: ctx, namespace, name
func (_m *Store) LookupDevice(ctx context.Context, namespace string, name string) (*models.Device, error) {
	ret := _m.Called(ctx, namespace, name)

	var r0 *models.Device
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *models.Device); ok {
		r0 = rf(ctx, namespace, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Device)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, namespace, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RenameDevice provides a mock function with given fields: ctx, uid, name
func (_m *Store) RenameDevice(ctx context.Context, uid models.UID, name string) error {
	ret := _m.Called(ctx, uid, name)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.UID, string) error); ok {
		r0 = rf(ctx, uid, name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetSessionAuthenticated provides a mock function with given fields: ctx, uid, authenticated
func (_m *Store) SetSessionAuthenticated(ctx context.Context, uid models.UID, authenticated bool) error {
	ret := _m.Called(ctx, uid, authenticated)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.UID, bool) error); ok {
		r0 = rf(ctx, uid, authenticated)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateDeviceStatus provides a mock function with given fields: ctx, uid, online
func (_m *Store) UpdateDeviceStatus(ctx context.Context, uid models.UID, online bool) error {
	ret := _m.Called(ctx, uid, online)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.UID, bool) error); ok {
		r0 = rf(ctx, uid, online)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateFirewallRule provides a mock function with given fields: ctx, id, rule
func (_m *Store) UpdateFirewallRule(ctx context.Context, id string, rule models.FirewallRuleUpdate) (*models.FirewallRule, error) {
	ret := _m.Called(ctx, id, rule)

	var r0 *models.FirewallRule
	if rf, ok := ret.Get(0).(func(context.Context, string, models.FirewallRuleUpdate) *models.FirewallRule); ok {
		r0 = rf(ctx, id, rule)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.FirewallRule)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, models.FirewallRuleUpdate) error); ok {
		r1 = rf(ctx, id, rule)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdatePendingStatus provides a mock function with given fields: ctx, uid, status
func (_m *Store) UpdatePendingStatus(ctx context.Context, uid models.UID, status string) error {
	ret := _m.Called(ctx, uid, status)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.UID, string) error); ok {
		r0 = rf(ctx, uid, status)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}