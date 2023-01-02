// Code generated by mockery v2.14.0. DO NOT EDIT.

package dashboards

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	models "github.com/grafana/grafana/pkg/models"
)

// FakeDashboardService is an autogenerated mock type for the DashboardService type
type FakeDashboardService struct {
	mock.Mock
}

// BuildSaveDashboardCommand provides a mock function with given fields: ctx, dto, shouldValidateAlerts, validateProvisionedDashboard
func (_m *FakeDashboardService) BuildSaveDashboardCommand(ctx context.Context, dto *SaveDashboardDTO, shouldValidateAlerts bool, validateProvisionedDashboard bool) (*models.SaveDashboardCommand, error) {
	ret := _m.Called(ctx, dto, shouldValidateAlerts, validateProvisionedDashboard)

	var r0 *models.SaveDashboardCommand
	if rf, ok := ret.Get(0).(func(context.Context, *SaveDashboardDTO, bool, bool) *models.SaveDashboardCommand); ok {
		r0 = rf(ctx, dto, shouldValidateAlerts, validateProvisionedDashboard)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.SaveDashboardCommand)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *SaveDashboardDTO, bool, bool) error); ok {
		r1 = rf(ctx, dto, shouldValidateAlerts, validateProvisionedDashboard)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CountDashboardsInFolder provides a mock function with given fields: ctx, query
func (_m *FakeDashboardService) CountDashboardsInFolder(ctx context.Context, query *CountDashboardsInFolderQuery) (int64, error) {
	ret := _m.Called(ctx, query)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, *CountDashboardsInFolderQuery) int64); ok {
		r0 = rf(ctx, query)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *CountDashboardsInFolderQuery) error); ok {
		r1 = rf(ctx, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteACLByUser provides a mock function with given fields: ctx, userID
func (_m *FakeDashboardService) DeleteACLByUser(ctx context.Context, userID int64) error {
	ret := _m.Called(ctx, userID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, userID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteDashboard provides a mock function with given fields: ctx, dashboardId, orgId
func (_m *FakeDashboardService) DeleteDashboard(ctx context.Context, dashboardId int64, orgId int64) error {
	ret := _m.Called(ctx, dashboardId, orgId)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, int64) error); ok {
		r0 = rf(ctx, dashboardId, orgId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindDashboards provides a mock function with given fields: ctx, query
func (_m *FakeDashboardService) FindDashboards(ctx context.Context, query *models.FindPersistedDashboardsQuery) ([]DashboardSearchProjection, error) {
	ret := _m.Called(ctx, query)

	var r0 []DashboardSearchProjection
	if rf, ok := ret.Get(0).(func(context.Context, *models.FindPersistedDashboardsQuery) []DashboardSearchProjection); ok {
		r0 = rf(ctx, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]DashboardSearchProjection)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *models.FindPersistedDashboardsQuery) error); ok {
		r1 = rf(ctx, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDashboard provides a mock function with given fields: ctx, query
func (_m *FakeDashboardService) GetDashboard(ctx context.Context, query *models.GetDashboardQuery) error {
	ret := _m.Called(ctx, query)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.GetDashboardQuery) error); ok {
		r0 = rf(ctx, query)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetDashboardACLInfoList provides a mock function with given fields: ctx, query
func (_m *FakeDashboardService) GetDashboardACLInfoList(ctx context.Context, query *models.GetDashboardACLInfoListQuery) error {
	ret := _m.Called(ctx, query)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.GetDashboardACLInfoListQuery) error); ok {
		r0 = rf(ctx, query)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetDashboardTags provides a mock function with given fields: ctx, query
func (_m *FakeDashboardService) GetDashboardTags(ctx context.Context, query *models.GetDashboardTagsQuery) error {
	ret := _m.Called(ctx, query)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.GetDashboardTagsQuery) error); ok {
		r0 = rf(ctx, query)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetDashboardUIDById provides a mock function with given fields: ctx, query
func (_m *FakeDashboardService) GetDashboardUIDById(ctx context.Context, query *models.GetDashboardRefByIdQuery) error {
	ret := _m.Called(ctx, query)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.GetDashboardRefByIdQuery) error); ok {
		r0 = rf(ctx, query)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetDashboards provides a mock function with given fields: ctx, query
func (_m *FakeDashboardService) GetDashboards(ctx context.Context, query *models.GetDashboardsQuery) error {
	ret := _m.Called(ctx, query)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.GetDashboardsQuery) error); ok {
		r0 = rf(ctx, query)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// HasAdminPermissionInDashboardsOrFolders provides a mock function with given fields: ctx, query
func (_m *FakeDashboardService) HasAdminPermissionInDashboardsOrFolders(ctx context.Context, query *models.HasAdminPermissionInDashboardsOrFoldersQuery) error {
	ret := _m.Called(ctx, query)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.HasAdminPermissionInDashboardsOrFoldersQuery) error); ok {
		r0 = rf(ctx, query)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// HasEditPermissionInFolders provides a mock function with given fields: ctx, query
func (_m *FakeDashboardService) HasEditPermissionInFolders(ctx context.Context, query *models.HasEditPermissionInFoldersQuery) error {
	ret := _m.Called(ctx, query)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.HasEditPermissionInFoldersQuery) error); ok {
		r0 = rf(ctx, query)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ImportDashboard provides a mock function with given fields: ctx, dto
func (_m *FakeDashboardService) ImportDashboard(ctx context.Context, dto *SaveDashboardDTO) (*models.Dashboard, error) {
	ret := _m.Called(ctx, dto)

	var r0 *models.Dashboard
	if rf, ok := ret.Get(0).(func(context.Context, *SaveDashboardDTO) *models.Dashboard); ok {
		r0 = rf(ctx, dto)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Dashboard)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *SaveDashboardDTO) error); ok {
		r1 = rf(ctx, dto)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MakeUserAdmin provides a mock function with given fields: ctx, orgID, userID, dashboardID, setViewAndEditPermissions
func (_m *FakeDashboardService) MakeUserAdmin(ctx context.Context, orgID int64, userID int64, dashboardID int64, setViewAndEditPermissions bool) error {
	ret := _m.Called(ctx, orgID, userID, dashboardID, setViewAndEditPermissions)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, int64, int64, bool) error); ok {
		r0 = rf(ctx, orgID, userID, dashboardID, setViewAndEditPermissions)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SaveDashboard provides a mock function with given fields: ctx, dto, allowUiUpdate
func (_m *FakeDashboardService) SaveDashboard(ctx context.Context, dto *SaveDashboardDTO, allowUiUpdate bool) (*models.Dashboard, error) {
	ret := _m.Called(ctx, dto, allowUiUpdate)

	var r0 *models.Dashboard
	if rf, ok := ret.Get(0).(func(context.Context, *SaveDashboardDTO, bool) *models.Dashboard); ok {
		r0 = rf(ctx, dto, allowUiUpdate)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Dashboard)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *SaveDashboardDTO, bool) error); ok {
		r1 = rf(ctx, dto, allowUiUpdate)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SearchDashboards provides a mock function with given fields: ctx, query
func (_m *FakeDashboardService) SearchDashboards(ctx context.Context, query *models.FindPersistedDashboardsQuery) error {
	ret := _m.Called(ctx, query)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.FindPersistedDashboardsQuery) error); ok {
		r0 = rf(ctx, query)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateDashboardACL provides a mock function with given fields: ctx, uid, items
func (_m *FakeDashboardService) UpdateDashboardACL(ctx context.Context, uid int64, items []*models.DashboardACL) error {
	ret := _m.Called(ctx, uid, items)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, []*models.DashboardACL) error); ok {
		r0 = rf(ctx, uid, items)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewFakeDashboardService interface {
	mock.TestingT
	Cleanup(func())
}

// NewFakeDashboardService creates a new instance of FakeDashboardService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewFakeDashboardService(t mockConstructorTestingTNewFakeDashboardService) *FakeDashboardService {
	mock := &FakeDashboardService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
