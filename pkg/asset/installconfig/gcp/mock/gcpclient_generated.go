// Code generated by MockGen. DO NOT EDIT.
// Source: ./client.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	compute "google.golang.org/api/compute/v1"
	dns "google.golang.org/api/dns/v1"
	reflect "reflect"
)

// MockAPI is a mock of API interface.
type MockAPI struct {
	ctrl     *gomock.Controller
	recorder *MockAPIMockRecorder
}

// MockAPIMockRecorder is the mock recorder for MockAPI.
type MockAPIMockRecorder struct {
	mock *MockAPI
}

// NewMockAPI creates a new mock instance.
func NewMockAPI(ctrl *gomock.Controller) *MockAPI {
	mock := &MockAPI{ctrl: ctrl}
	mock.recorder = &MockAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAPI) EXPECT() *MockAPIMockRecorder {
	return m.recorder
}

// GetNetwork mocks base method.
func (m *MockAPI) GetNetwork(ctx context.Context, network, project string) (*compute.Network, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNetwork", ctx, network, project)
	ret0, _ := ret[0].(*compute.Network)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNetwork indicates an expected call of GetNetwork.
func (mr *MockAPIMockRecorder) GetNetwork(ctx, network, project interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNetwork", reflect.TypeOf((*MockAPI)(nil).GetNetwork), ctx, network, project)
}

// GetPublicDomains mocks base method.
func (m *MockAPI) GetPublicDomains(ctx context.Context, project string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPublicDomains", ctx, project)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPublicDomains indicates an expected call of GetPublicDomains.
func (mr *MockAPIMockRecorder) GetPublicDomains(ctx, project interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPublicDomains", reflect.TypeOf((*MockAPI)(nil).GetPublicDomains), ctx, project)
}

// GetPublicDNSZone mocks base method.
func (m *MockAPI) GetPublicDNSZone(ctx context.Context, baseDomain, project string) (*dns.ManagedZone, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPublicDNSZone", ctx, baseDomain, project)
	ret0, _ := ret[0].(*dns.ManagedZone)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPublicDNSZone indicates an expected call of GetPublicDNSZone.
func (mr *MockAPIMockRecorder) GetPublicDNSZone(ctx, baseDomain, project interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPublicDNSZone", reflect.TypeOf((*MockAPI)(nil).GetPublicDNSZone), ctx, baseDomain, project)
}

// GetSubnetworks mocks base method.
func (m *MockAPI) GetSubnetworks(ctx context.Context, network, project, region string) ([]*compute.Subnetwork, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubnetworks", ctx, network, project, region)
	ret0, _ := ret[0].([]*compute.Subnetwork)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubnetworks indicates an expected call of GetSubnetworks.
func (mr *MockAPIMockRecorder) GetSubnetworks(ctx, network, project, region interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubnetworks", reflect.TypeOf((*MockAPI)(nil).GetSubnetworks), ctx, network, project, region)
}

// GetProjects mocks base method.
func (m *MockAPI) GetProjects(ctx context.Context) (map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProjects", ctx)
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProjects indicates an expected call of GetProjects.
func (mr *MockAPIMockRecorder) GetProjects(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProjects", reflect.TypeOf((*MockAPI)(nil).GetProjects), ctx)
}

// GetRecordSets mocks base method.
func (m *MockAPI) GetRecordSets(ctx context.Context, project, zone string) ([]*dns.ResourceRecordSet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRecordSets", ctx, project, zone)
	ret0, _ := ret[0].([]*dns.ResourceRecordSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRecordSets indicates an expected call of GetRecordSets.
func (mr *MockAPIMockRecorder) GetRecordSets(ctx, project, zone interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRecordSets", reflect.TypeOf((*MockAPI)(nil).GetRecordSets), ctx, project, zone)
}
