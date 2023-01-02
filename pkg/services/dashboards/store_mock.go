// Code generated by mockery v2.14.0. DO NOT EDIT.

package dashboards

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	models "github.com/grafana/grafana/pkg/models"
	folder "github.com/grafana/grafana/pkg/services/folder"
	"github.com/grafana/grafana/pkg/services/quota"
)

// FakeDashboardStore is an autogenerated mock type for the Store type
type FakeDashboardStore struct {
	mock.Mock
}

// CountDashboardsInFolder provides a mock function with given fields: ctx, request
func (_m *FakeDashboardStore) CountDashboardsInFolder(ctx context.Context, request *CountDashboardsInFolderRequest) (int64, error) {
	ret := _m.Called(ctx, request)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, *CountDashboardsInFolderRequest) int64); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *CountDashboardsInFolderRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteACLByUser provides a mock function with given fields: _a0, _a1
func (_m *FakeDashboardStore) DeleteACLByUser(_a0 context.Context, _a1 int64) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteDashboard provides a mock function with given fields: ctx, cmd
func (_m *FakeDashboardStore) DeleteDashboard(ctx context.Context, cmd *models.DeleteDashboardCommand) error {
	ret := _m.Called(ctx, cmd)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.DeleteDashboardCommand) error); ok {
		r0 = rf(ctx, cmd)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteOrphanedProvisionedDashboards provides a mock function with given fields: ctx, cmd
func (_m *FakeDashboardStore) DeleteOrphanedProvisionedDashboards(ctx context.Context, cmd *models.DeleteOrphanedProvisionedDashboardsCommand) error {
	ret := _m.Called(ctx, cmd)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.DeleteOrphanedProvisionedDashboardsCommand) error); ok {
		r0 = rf(ctx, cmd)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindDashboards provides a mock function with given fields: ctx, query
func (_m *FakeDashboardStore) FindDashboards(ctx context.Context, query *models.FindPersistedDashboardsQuery) ([]DashboardSearchProjection, error) {
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
func (_m *FakeDashboardStore) GetDashboard(ctx context.Context, query *models.GetDashboardQuery) (*models.Dashboard, error) {
	ret := _m.Called(ctx, query)

	var r0 *models.Dashboard
	if rf, ok := ret.Get(0).(func(context.Context, *models.GetDashboardQuery) *models.Dashboard); ok {
		r0 = rf(ctx, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Dashboard)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *models.GetDashboardQuery) error); ok {
		r1 = rf(ctx, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDashboardACLInfoList provides a mock function with given fields: ctx, query
func (_m *FakeDashboardStore) GetDashboardACLInfoList(ctx context.Context, query *models.GetDashboardACLInfoListQuery) error {
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
func (_m *FakeDashboardStore) GetDashboardTags(ctx context.Context, query *models.GetDashboardTagsQuery) error {
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
func (_m *FakeDashboardStore) GetDashboardUIDById(ctx context.Context, query *models.GetDashboardRefByIdQuery) error {
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
func (_m *FakeDashboardStore) GetDashboards(ctx context.Context, query *models.GetDashboardsQuery) error {
	ret := _m.Called(ctx, query)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.GetDashboardsQuery) error); ok {
		r0 = rf(ctx, query)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetDashboardsByPluginID provides a mock function with given fields: ctx, query
func (_m *FakeDashboardStore) GetDashboardsByPluginID(ctx context.Context, query *models.GetDashboardsByPluginIdQuery) error {
	ret := _m.Called(ctx, query)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.GetDashboardsByPluginIdQuery) error); ok {
		r0 = rf(ctx, query)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetFolderByID provides a mock function with given fields: ctx, orgID, id
func (_m *FakeDashboardStore) GetFolderByID(ctx context.Context, orgID int64, id int64) (*folder.Folder, error) {
	ret := _m.Called(ctx, orgID, id)

	var r0 *folder.Folder
	if rf, ok := ret.Get(0).(func(context.Context, int64, int64) *folder.Folder); ok {
		r0 = rf(ctx, orgID, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*folder.Folder)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64, int64) error); ok {
		r1 = rf(ctx, orgID, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetFolderByTitle provides a mock function with given fields: ctx, orgID, title
func (_m *FakeDashboardStore) GetFolderByTitle(ctx context.Context, orgID int64, title string) (*folder.Folder, error) {
	ret := _m.Called(ctx, orgID, title)

	var r0 *folder.Folder
	if rf, ok := ret.Get(0).(func(context.Context, int64, string) *folder.Folder); ok {
		r0 = rf(ctx, orgID, title)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*folder.Folder)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64, string) error); ok {
		r1 = rf(ctx, orgID, title)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetFolderByUID provides a mock function with given fields: ctx, orgID, uid
func (_m *FakeDashboardStore) GetFolderByUID(ctx context.Context, orgID int64, uid string) (*folder.Folder, error) {
	ret := _m.Called(ctx, orgID, uid)

	var r0 *folder.Folder
	if rf, ok := ret.Get(0).(func(context.Context, int64, string) *folder.Folder); ok {
		r0 = rf(ctx, orgID, uid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*folder.Folder)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64, string) error); ok {
		r1 = rf(ctx, orgID, uid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProvisionedDashboardData provides a mock function with given fields: ctx, name
func (_m *FakeDashboardStore) GetProvisionedDashboardData(ctx context.Context, name string) ([]*models.DashboardProvisioning, error) {
	ret := _m.Called(ctx, name)

	var r0 []*models.DashboardProvisioning
	if rf, ok := ret.Get(0).(func(context.Context, string) []*models.DashboardProvisioning); ok {
		r0 = rf(ctx, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.DashboardProvisioning)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProvisionedDataByDashboardID provides a mock function with given fields: ctx, dashboardID
func (_m *FakeDashboardStore) GetProvisionedDataByDashboardID(ctx context.Context, dashboardID int64) (*models.DashboardProvisioning, error) {
	ret := _m.Called(ctx, dashboardID)

	var r0 *models.DashboardProvisioning
	if rf, ok := ret.Get(0).(func(context.Context, int64) *models.DashboardProvisioning); ok {
		r0 = rf(ctx, dashboardID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.DashboardProvisioning)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, dashboardID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProvisionedDataByDashboardUID provides a mock function with given fields: ctx, orgID, dashboardUID
func (_m *FakeDashboardStore) GetProvisionedDataByDashboardUID(ctx context.Context, orgID int64, dashboardUID string) (*models.DashboardProvisioning, error) {
	ret := _m.Called(ctx, orgID, dashboardUID)

	var r0 *models.DashboardProvisioning
	if rf, ok := ret.Get(0).(func(context.Context, int64, string) *models.DashboardProvisioning); ok {
		r0 = rf(ctx, orgID, dashboardUID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.DashboardProvisioning)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64, string) error); ok {
		r1 = rf(ctx, orgID, dashboardUID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HasAdminPermissionInDashboardsOrFolders provides a mock function with given fields: ctx, query
func (_m *FakeDashboardStore) HasAdminPermissionInDashboardsOrFolders(ctx context.Context, query *models.HasAdminPermissionInDashboardsOrFoldersQuery) error {
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
func (_m *FakeDashboardStore) HasEditPermissionInFolders(ctx context.Context, query *models.HasEditPermissionInFoldersQuery) error {
	ret := _m.Called(ctx, query)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.HasEditPermissionInFoldersQuery) error); ok {
		r0 = rf(ctx, query)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SaveAlerts provides a mock function with given fields: ctx, dashID, alerts
func (_m *FakeDashboardStore) SaveAlerts(ctx context.Context, dashID int64, alerts []*models.Alert) error {
	ret := _m.Called(ctx, dashID, alerts)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, []*models.Alert) error); ok {
		r0 = rf(ctx, dashID, alerts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SaveDashboard provides a mock function with given fields: ctx, cmd
func (_m *FakeDashboardStore) SaveDashboard(ctx context.Context, cmd models.SaveDashboardCommand) (*models.Dashboard, error) {
	ret := _m.Called(ctx, cmd)

	var r0 *models.Dashboard
	if rf, ok := ret.Get(0).(func(context.Context, models.SaveDashboardCommand) *models.Dashboard); ok {
		r0 = rf(ctx, cmd)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Dashboard)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.SaveDashboardCommand) error); ok {
		r1 = rf(ctx, cmd)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SaveProvisionedDashboard provides a mock function with given fields: ctx, cmd, provisioning
func (_m *FakeDashboardStore) SaveProvisionedDashboard(ctx context.Context, cmd models.SaveDashboardCommand, provisioning *models.DashboardProvisioning) (*models.Dashboard, error) {
	ret := _m.Called(ctx, cmd, provisioning)

	var r0 *models.Dashboard
	if rf, ok := ret.Get(0).(func(context.Context, models.SaveDashboardCommand, *models.DashboardProvisioning) *models.Dashboard); ok {
		r0 = rf(ctx, cmd, provisioning)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Dashboard)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.SaveDashboardCommand, *models.DashboardProvisioning) error); ok {
		r1 = rf(ctx, cmd, provisioning)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UnprovisionDashboard provides a mock function with given fields: ctx, id
func (_m *FakeDashboardStore) UnprovisionDashboard(ctx context.Context, id int64) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateDashboardACL provides a mock function with given fields: ctx, uid, items
func (_m *FakeDashboardStore) UpdateDashboardACL(ctx context.Context, uid int64, items []*models.DashboardACL) error {
	ret := _m.Called(ctx, uid, items)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, []*models.DashboardACL) error); ok {
		r0 = rf(ctx, uid, items)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ValidateDashboardBeforeSave provides a mock function with given fields: ctx, dashboard, overwrite
func (_m *FakeDashboardStore) ValidateDashboardBeforeSave(ctx context.Context, dashboard *models.Dashboard, overwrite bool) (bool, error) {
	ret := _m.Called(ctx, dashboard, overwrite)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, *models.Dashboard, bool) bool); ok {
		r0 = rf(ctx, dashboard, overwrite)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *models.Dashboard, bool) error); ok {
		r1 = rf(ctx, dashboard, overwrite)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewFakeDashboardStore interface {
	mock.TestingT
	Cleanup(func())
}

func (_m *FakeDashboardStore) Count(context.Context, *quota.ScopeParameters) (*quota.Map, error) {
	return nil, nil
}

// NewFakeDashboardStore creates a new instance of FakeDashboardStore. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewFakeDashboardStore(t mockConstructorTestingTNewFakeDashboardStore) *FakeDashboardStore {
	mock := &FakeDashboardStore{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
