// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/dongwlin/anime-list/internal/db (interfaces: Store)
//
// Generated by this command:
//
//	mockgen -package mock -destination internal/mock/store.go github.com/dongwlin/anime-list/internal/db Store
//

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	db "github.com/dongwlin/anime-list/internal/db"
	gomock "go.uber.org/mock/gomock"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// CountAnime mocks base method.
func (m *MockStore) CountAnime(arg0 context.Context) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountAnime", arg0)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountAnime indicates an expected call of CountAnime.
func (mr *MockStoreMockRecorder) CountAnime(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountAnime", reflect.TypeOf((*MockStore)(nil).CountAnime), arg0)
}

// CreateAnime mocks base method.
func (m *MockStore) CreateAnime(arg0 context.Context, arg1 db.CreateAnimeParams) (db.Anime, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAnime", arg0, arg1)
	ret0, _ := ret[0].(db.Anime)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAnime indicates an expected call of CreateAnime.
func (mr *MockStoreMockRecorder) CreateAnime(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAnime", reflect.TypeOf((*MockStore)(nil).CreateAnime), arg0, arg1)
}

// CreateEpisode mocks base method.
func (m *MockStore) CreateEpisode(arg0 context.Context, arg1 db.CreateEpisodeParams) (db.Episode, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateEpisode", arg0, arg1)
	ret0, _ := ret[0].(db.Episode)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateEpisode indicates an expected call of CreateEpisode.
func (mr *MockStoreMockRecorder) CreateEpisode(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEpisode", reflect.TypeOf((*MockStore)(nil).CreateEpisode), arg0, arg1)
}

// CreateSeason mocks base method.
func (m *MockStore) CreateSeason(arg0 context.Context, arg1 db.CreateSeasonParams) (db.Season, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSeason", arg0, arg1)
	ret0, _ := ret[0].(db.Season)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSeason indicates an expected call of CreateSeason.
func (mr *MockStoreMockRecorder) CreateSeason(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSeason", reflect.TypeOf((*MockStore)(nil).CreateSeason), arg0, arg1)
}

// CreateTheater mocks base method.
func (m *MockStore) CreateTheater(arg0 context.Context, arg1 db.CreateTheaterParams) (db.Theater, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTheater", arg0, arg1)
	ret0, _ := ret[0].(db.Theater)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTheater indicates an expected call of CreateTheater.
func (mr *MockStoreMockRecorder) CreateTheater(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTheater", reflect.TypeOf((*MockStore)(nil).CreateTheater), arg0, arg1)
}

// DeleteAnime mocks base method.
func (m *MockStore) DeleteAnime(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAnime", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAnime indicates an expected call of DeleteAnime.
func (mr *MockStoreMockRecorder) DeleteAnime(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAnime", reflect.TypeOf((*MockStore)(nil).DeleteAnime), arg0, arg1)
}

// DeleteEpisode mocks base method.
func (m *MockStore) DeleteEpisode(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteEpisode", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteEpisode indicates an expected call of DeleteEpisode.
func (mr *MockStoreMockRecorder) DeleteEpisode(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteEpisode", reflect.TypeOf((*MockStore)(nil).DeleteEpisode), arg0, arg1)
}

// DeleteSeason mocks base method.
func (m *MockStore) DeleteSeason(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSeason", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSeason indicates an expected call of DeleteSeason.
func (mr *MockStoreMockRecorder) DeleteSeason(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSeason", reflect.TypeOf((*MockStore)(nil).DeleteSeason), arg0, arg1)
}

// DeleteTheater mocks base method.
func (m *MockStore) DeleteTheater(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTheater", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTheater indicates an expected call of DeleteTheater.
func (mr *MockStoreMockRecorder) DeleteTheater(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTheater", reflect.TypeOf((*MockStore)(nil).DeleteTheater), arg0, arg1)
}

// GetAnime mocks base method.
func (m *MockStore) GetAnime(arg0 context.Context, arg1 int64) (db.Anime, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAnime", arg0, arg1)
	ret0, _ := ret[0].(db.Anime)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAnime indicates an expected call of GetAnime.
func (mr *MockStoreMockRecorder) GetAnime(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAnime", reflect.TypeOf((*MockStore)(nil).GetAnime), arg0, arg1)
}

// GetEpisode mocks base method.
func (m *MockStore) GetEpisode(arg0 context.Context, arg1 int64) (db.Episode, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEpisode", arg0, arg1)
	ret0, _ := ret[0].(db.Episode)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEpisode indicates an expected call of GetEpisode.
func (mr *MockStoreMockRecorder) GetEpisode(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEpisode", reflect.TypeOf((*MockStore)(nil).GetEpisode), arg0, arg1)
}

// GetSeason mocks base method.
func (m *MockStore) GetSeason(arg0 context.Context, arg1 int64) (db.Season, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSeason", arg0, arg1)
	ret0, _ := ret[0].(db.Season)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSeason indicates an expected call of GetSeason.
func (mr *MockStoreMockRecorder) GetSeason(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSeason", reflect.TypeOf((*MockStore)(nil).GetSeason), arg0, arg1)
}

// GetTheater mocks base method.
func (m *MockStore) GetTheater(arg0 context.Context, arg1 int64) (db.Theater, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTheater", arg0, arg1)
	ret0, _ := ret[0].(db.Theater)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTheater indicates an expected call of GetTheater.
func (mr *MockStoreMockRecorder) GetTheater(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTheater", reflect.TypeOf((*MockStore)(nil).GetTheater), arg0, arg1)
}

// ListAnime mocks base method.
func (m *MockStore) ListAnime(arg0 context.Context, arg1 db.ListAnimeParams) ([]db.Anime, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAnime", arg0, arg1)
	ret0, _ := ret[0].([]db.Anime)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAnime indicates an expected call of ListAnime.
func (mr *MockStoreMockRecorder) ListAnime(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAnime", reflect.TypeOf((*MockStore)(nil).ListAnime), arg0, arg1)
}

// ListEpisode mocks base method.
func (m *MockStore) ListEpisode(arg0 context.Context, arg1 db.ListEpisodeParams) ([]db.Episode, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListEpisode", arg0, arg1)
	ret0, _ := ret[0].([]db.Episode)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListEpisode indicates an expected call of ListEpisode.
func (mr *MockStoreMockRecorder) ListEpisode(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListEpisode", reflect.TypeOf((*MockStore)(nil).ListEpisode), arg0, arg1)
}

// ListSeason mocks base method.
func (m *MockStore) ListSeason(arg0 context.Context, arg1 db.ListSeasonParams) ([]db.Season, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSeason", arg0, arg1)
	ret0, _ := ret[0].([]db.Season)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSeason indicates an expected call of ListSeason.
func (mr *MockStoreMockRecorder) ListSeason(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSeason", reflect.TypeOf((*MockStore)(nil).ListSeason), arg0, arg1)
}

// ListTheater mocks base method.
func (m *MockStore) ListTheater(arg0 context.Context, arg1 db.ListTheaterParams) ([]db.Theater, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTheater", arg0, arg1)
	ret0, _ := ret[0].([]db.Theater)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTheater indicates an expected call of ListTheater.
func (mr *MockStoreMockRecorder) ListTheater(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTheater", reflect.TypeOf((*MockStore)(nil).ListTheater), arg0, arg1)
}

// UpdateAnime mocks base method.
func (m *MockStore) UpdateAnime(arg0 context.Context, arg1 db.UpdateAnimeParams) (db.Anime, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAnime", arg0, arg1)
	ret0, _ := ret[0].(db.Anime)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAnime indicates an expected call of UpdateAnime.
func (mr *MockStoreMockRecorder) UpdateAnime(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAnime", reflect.TypeOf((*MockStore)(nil).UpdateAnime), arg0, arg1)
}

// UpdateEpisode mocks base method.
func (m *MockStore) UpdateEpisode(arg0 context.Context, arg1 db.UpdateEpisodeParams) (db.Episode, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateEpisode", arg0, arg1)
	ret0, _ := ret[0].(db.Episode)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateEpisode indicates an expected call of UpdateEpisode.
func (mr *MockStoreMockRecorder) UpdateEpisode(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateEpisode", reflect.TypeOf((*MockStore)(nil).UpdateEpisode), arg0, arg1)
}

// UpdateSeason mocks base method.
func (m *MockStore) UpdateSeason(arg0 context.Context, arg1 db.UpdateSeasonParams) (db.Season, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSeason", arg0, arg1)
	ret0, _ := ret[0].(db.Season)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateSeason indicates an expected call of UpdateSeason.
func (mr *MockStoreMockRecorder) UpdateSeason(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSeason", reflect.TypeOf((*MockStore)(nil).UpdateSeason), arg0, arg1)
}

// UpdateTheater mocks base method.
func (m *MockStore) UpdateTheater(arg0 context.Context, arg1 db.UpdateTheaterParams) (db.Theater, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTheater", arg0, arg1)
	ret0, _ := ret[0].(db.Theater)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTheater indicates an expected call of UpdateTheater.
func (mr *MockStoreMockRecorder) UpdateTheater(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTheater", reflect.TypeOf((*MockStore)(nil).UpdateTheater), arg0, arg1)
}
