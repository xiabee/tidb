// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/pingcap/tidb/pkg/disttask/framework/dispatcher (interfaces: Dispatcher,CleanUpRoutine,TaskManager)
//
// Generated by this command:
//
//	mockgen -package mock github.com/pingcap/tidb/pkg/disttask/framework/dispatcher Dispatcher,CleanUpRoutine,TaskManager
//
// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	proto "github.com/pingcap/tidb/pkg/disttask/framework/proto"
	sessionctx "github.com/pingcap/tidb/pkg/sessionctx"
	gomock "go.uber.org/mock/gomock"
)

// MockDispatcher is a mock of Dispatcher interface.
type MockDispatcher struct {
	ctrl     *gomock.Controller
	recorder *MockDispatcherMockRecorder
}

// MockDispatcherMockRecorder is the mock recorder for MockDispatcher.
type MockDispatcherMockRecorder struct {
	mock *MockDispatcher
}

// NewMockDispatcher creates a new mock instance.
func NewMockDispatcher(ctrl *gomock.Controller) *MockDispatcher {
	mock := &MockDispatcher{ctrl: ctrl}
	mock.recorder = &MockDispatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDispatcher) EXPECT() *MockDispatcherMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockDispatcher) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockDispatcherMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockDispatcher)(nil).Close))
}

// ExecuteTask mocks base method.
func (m *MockDispatcher) ExecuteTask() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ExecuteTask")
}

// ExecuteTask indicates an expected call of ExecuteTask.
func (mr *MockDispatcherMockRecorder) ExecuteTask() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecuteTask", reflect.TypeOf((*MockDispatcher)(nil).ExecuteTask))
}

// Init mocks base method.
func (m *MockDispatcher) Init() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Init")
	ret0, _ := ret[0].(error)
	return ret0
}

// Init indicates an expected call of Init.
func (mr *MockDispatcherMockRecorder) Init() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockDispatcher)(nil).Init))
}

// MockCleanUpRoutine is a mock of CleanUpRoutine interface.
type MockCleanUpRoutine struct {
	ctrl     *gomock.Controller
	recorder *MockCleanUpRoutineMockRecorder
}

// MockCleanUpRoutineMockRecorder is the mock recorder for MockCleanUpRoutine.
type MockCleanUpRoutineMockRecorder struct {
	mock *MockCleanUpRoutine
}

// NewMockCleanUpRoutine creates a new mock instance.
func NewMockCleanUpRoutine(ctrl *gomock.Controller) *MockCleanUpRoutine {
	mock := &MockCleanUpRoutine{ctrl: ctrl}
	mock.recorder = &MockCleanUpRoutineMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCleanUpRoutine) EXPECT() *MockCleanUpRoutineMockRecorder {
	return m.recorder
}

// CleanUp mocks base method.
func (m *MockCleanUpRoutine) CleanUp(arg0 context.Context, arg1 *proto.Task) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CleanUp", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CleanUp indicates an expected call of CleanUp.
func (mr *MockCleanUpRoutineMockRecorder) CleanUp(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CleanUp", reflect.TypeOf((*MockCleanUpRoutine)(nil).CleanUp), arg0, arg1)
}

// MockTaskManager is a mock of TaskManager interface.
type MockTaskManager struct {
	ctrl     *gomock.Controller
	recorder *MockTaskManagerMockRecorder
}

// MockTaskManagerMockRecorder is the mock recorder for MockTaskManager.
type MockTaskManagerMockRecorder struct {
	mock *MockTaskManager
}

// NewMockTaskManager creates a new mock instance.
func NewMockTaskManager(ctrl *gomock.Controller) *MockTaskManager {
	mock := &MockTaskManager{ctrl: ctrl}
	mock.recorder = &MockTaskManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTaskManager) EXPECT() *MockTaskManagerMockRecorder {
	return m.recorder
}

// CancelTask mocks base method.
func (m *MockTaskManager) CancelTask(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CancelTask", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CancelTask indicates an expected call of CancelTask.
func (mr *MockTaskManagerMockRecorder) CancelTask(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelTask", reflect.TypeOf((*MockTaskManager)(nil).CancelTask), arg0, arg1)
}

// CleanUpMeta mocks base method.
func (m *MockTaskManager) CleanUpMeta(arg0 context.Context, arg1 []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CleanUpMeta", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CleanUpMeta indicates an expected call of CleanUpMeta.
func (mr *MockTaskManagerMockRecorder) CleanUpMeta(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CleanUpMeta", reflect.TypeOf((*MockTaskManager)(nil).CleanUpMeta), arg0, arg1)
}

// CollectSubTaskError mocks base method.
func (m *MockTaskManager) CollectSubTaskError(arg0 context.Context, arg1 int64) ([]error, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CollectSubTaskError", arg0, arg1)
	ret0, _ := ret[0].([]error)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CollectSubTaskError indicates an expected call of CollectSubTaskError.
func (mr *MockTaskManagerMockRecorder) CollectSubTaskError(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CollectSubTaskError", reflect.TypeOf((*MockTaskManager)(nil).CollectSubTaskError), arg0, arg1)
}

// FailTask mocks base method.
func (m *MockTaskManager) FailTask(arg0 context.Context, arg1 int64, arg2 proto.TaskState, arg3 error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FailTask", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// FailTask indicates an expected call of FailTask.
func (mr *MockTaskManagerMockRecorder) FailTask(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FailTask", reflect.TypeOf((*MockTaskManager)(nil).FailTask), arg0, arg1, arg2, arg3)
}

// GCSubtasks mocks base method.
func (m *MockTaskManager) GCSubtasks(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GCSubtasks", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// GCSubtasks indicates an expected call of GCSubtasks.
func (mr *MockTaskManagerMockRecorder) GCSubtasks(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GCSubtasks", reflect.TypeOf((*MockTaskManager)(nil).GCSubtasks), arg0)
}

// GetAllNodes mocks base method.
func (m *MockTaskManager) GetAllNodes(arg0 context.Context) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllNodes", arg0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllNodes indicates an expected call of GetAllNodes.
func (mr *MockTaskManagerMockRecorder) GetAllNodes(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllNodes", reflect.TypeOf((*MockTaskManager)(nil).GetAllNodes), arg0)
}

// GetManagedNodes mocks base method.
func (m *MockTaskManager) GetManagedNodes(arg0 context.Context) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetManagedNodes", arg0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetManagedNodes indicates an expected call of GetManagedNodes.
func (mr *MockTaskManagerMockRecorder) GetManagedNodes(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetManagedNodes", reflect.TypeOf((*MockTaskManager)(nil).GetManagedNodes), arg0)
}

// GetSubtaskInStatesCnt mocks base method.
func (m *MockTaskManager) GetSubtaskInStatesCnt(arg0 context.Context, arg1 int64, arg2 ...any) (int64, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetSubtaskInStatesCnt", varargs...)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubtaskInStatesCnt indicates an expected call of GetSubtaskInStatesCnt.
func (mr *MockTaskManagerMockRecorder) GetSubtaskInStatesCnt(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubtaskInStatesCnt", reflect.TypeOf((*MockTaskManager)(nil).GetSubtaskInStatesCnt), varargs...)
}

// GetSubtasksByExecIdsAndStepAndState mocks base method.
func (m *MockTaskManager) GetSubtasksByExecIdsAndStepAndState(arg0 context.Context, arg1 []string, arg2 int64, arg3 proto.Step, arg4 proto.TaskState) ([]*proto.Subtask, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubtasksByExecIdsAndStepAndState", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].([]*proto.Subtask)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubtasksByExecIdsAndStepAndState indicates an expected call of GetSubtasksByExecIdsAndStepAndState.
func (mr *MockTaskManagerMockRecorder) GetSubtasksByExecIdsAndStepAndState(arg0, arg1, arg2, arg3, arg4 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubtasksByExecIdsAndStepAndState", reflect.TypeOf((*MockTaskManager)(nil).GetSubtasksByExecIdsAndStepAndState), arg0, arg1, arg2, arg3, arg4)
}

// GetSubtasksByStepAndState mocks base method.
func (m *MockTaskManager) GetSubtasksByStepAndState(arg0 context.Context, arg1 int64, arg2 proto.Step, arg3 proto.TaskState) ([]*proto.Subtask, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubtasksByStepAndState", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]*proto.Subtask)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubtasksByStepAndState indicates an expected call of GetSubtasksByStepAndState.
func (mr *MockTaskManagerMockRecorder) GetSubtasksByStepAndState(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubtasksByStepAndState", reflect.TypeOf((*MockTaskManager)(nil).GetSubtasksByStepAndState), arg0, arg1, arg2, arg3)
}

// GetTaskByID mocks base method.
func (m *MockTaskManager) GetTaskByID(arg0 context.Context, arg1 int64) (*proto.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTaskByID", arg0, arg1)
	ret0, _ := ret[0].(*proto.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTaskByID indicates an expected call of GetTaskByID.
func (mr *MockTaskManagerMockRecorder) GetTaskByID(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTaskByID", reflect.TypeOf((*MockTaskManager)(nil).GetTaskByID), arg0, arg1)
}

// GetTaskExecutorIDsByTaskID mocks base method.
func (m *MockTaskManager) GetTaskExecutorIDsByTaskID(arg0 context.Context, arg1 int64) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTaskExecutorIDsByTaskID", arg0, arg1)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTaskExecutorIDsByTaskID indicates an expected call of GetTaskExecutorIDsByTaskID.
func (mr *MockTaskManagerMockRecorder) GetTaskExecutorIDsByTaskID(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTaskExecutorIDsByTaskID", reflect.TypeOf((*MockTaskManager)(nil).GetTaskExecutorIDsByTaskID), arg0, arg1)
}

// GetTaskExecutorIDsByTaskIDAndStep mocks base method.
func (m *MockTaskManager) GetTaskExecutorIDsByTaskIDAndStep(arg0 context.Context, arg1 int64, arg2 proto.Step) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTaskExecutorIDsByTaskIDAndStep", arg0, arg1, arg2)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTaskExecutorIDsByTaskIDAndStep indicates an expected call of GetTaskExecutorIDsByTaskIDAndStep.
func (mr *MockTaskManagerMockRecorder) GetTaskExecutorIDsByTaskIDAndStep(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTaskExecutorIDsByTaskIDAndStep", reflect.TypeOf((*MockTaskManager)(nil).GetTaskExecutorIDsByTaskIDAndStep), arg0, arg1, arg2)
}

// GetTasksInStates mocks base method.
func (m *MockTaskManager) GetTasksInStates(arg0 context.Context, arg1 ...any) ([]*proto.Task, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetTasksInStates", varargs...)
	ret0, _ := ret[0].([]*proto.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTasksInStates indicates an expected call of GetTasksInStates.
func (mr *MockTaskManagerMockRecorder) GetTasksInStates(arg0 any, arg1 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTasksInStates", reflect.TypeOf((*MockTaskManager)(nil).GetTasksInStates), varargs...)
}

// GetTopUnfinishedTasks mocks base method.
func (m *MockTaskManager) GetTopUnfinishedTasks(arg0 context.Context) ([]*proto.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTopUnfinishedTasks", arg0)
	ret0, _ := ret[0].([]*proto.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTopUnfinishedTasks indicates an expected call of GetTopUnfinishedTasks.
func (mr *MockTaskManagerMockRecorder) GetTopUnfinishedTasks(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTopUnfinishedTasks", reflect.TypeOf((*MockTaskManager)(nil).GetTopUnfinishedTasks), arg0)
}

// GetUsedSlotsOnNodes mocks base method.
func (m *MockTaskManager) GetUsedSlotsOnNodes(arg0 context.Context) (map[string]int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsedSlotsOnNodes", arg0)
	ret0, _ := ret[0].(map[string]int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUsedSlotsOnNodes indicates an expected call of GetUsedSlotsOnNodes.
func (mr *MockTaskManagerMockRecorder) GetUsedSlotsOnNodes(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsedSlotsOnNodes", reflect.TypeOf((*MockTaskManager)(nil).GetUsedSlotsOnNodes), arg0)
}

// PauseTask mocks base method.
func (m *MockTaskManager) PauseTask(arg0 context.Context, arg1 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PauseTask", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PauseTask indicates an expected call of PauseTask.
func (mr *MockTaskManagerMockRecorder) PauseTask(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PauseTask", reflect.TypeOf((*MockTaskManager)(nil).PauseTask), arg0, arg1)
}

// ResumeSubtasks mocks base method.
func (m *MockTaskManager) ResumeSubtasks(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResumeSubtasks", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ResumeSubtasks indicates an expected call of ResumeSubtasks.
func (mr *MockTaskManagerMockRecorder) ResumeSubtasks(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResumeSubtasks", reflect.TypeOf((*MockTaskManager)(nil).ResumeSubtasks), arg0, arg1)
}

// TransferSubTasks2History mocks base method.
func (m *MockTaskManager) TransferSubTasks2History(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TransferSubTasks2History", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// TransferSubTasks2History indicates an expected call of TransferSubTasks2History.
func (mr *MockTaskManagerMockRecorder) TransferSubTasks2History(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransferSubTasks2History", reflect.TypeOf((*MockTaskManager)(nil).TransferSubTasks2History), arg0, arg1)
}

// TransferTasks2History mocks base method.
func (m *MockTaskManager) TransferTasks2History(arg0 context.Context, arg1 []*proto.Task) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TransferTasks2History", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// TransferTasks2History indicates an expected call of TransferTasks2History.
func (mr *MockTaskManagerMockRecorder) TransferTasks2History(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransferTasks2History", reflect.TypeOf((*MockTaskManager)(nil).TransferTasks2History), arg0, arg1)
}

// UpdateSubtasksExecIDs mocks base method.
func (m *MockTaskManager) UpdateSubtasksExecIDs(arg0 context.Context, arg1 int64, arg2 []*proto.Subtask) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSubtasksExecIDs", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateSubtasksExecIDs indicates an expected call of UpdateSubtasksExecIDs.
func (mr *MockTaskManagerMockRecorder) UpdateSubtasksExecIDs(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSubtasksExecIDs", reflect.TypeOf((*MockTaskManager)(nil).UpdateSubtasksExecIDs), arg0, arg1, arg2)
}

// UpdateTaskAndAddSubTasks mocks base method.
func (m *MockTaskManager) UpdateTaskAndAddSubTasks(arg0 context.Context, arg1 *proto.Task, arg2 []*proto.Subtask, arg3 proto.TaskState) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTaskAndAddSubTasks", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTaskAndAddSubTasks indicates an expected call of UpdateTaskAndAddSubTasks.
func (mr *MockTaskManagerMockRecorder) UpdateTaskAndAddSubTasks(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTaskAndAddSubTasks", reflect.TypeOf((*MockTaskManager)(nil).UpdateTaskAndAddSubTasks), arg0, arg1, arg2, arg3)
}

// WithNewSession mocks base method.
func (m *MockTaskManager) WithNewSession(arg0 func(sessionctx.Context) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithNewSession", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// WithNewSession indicates an expected call of WithNewSession.
func (mr *MockTaskManagerMockRecorder) WithNewSession(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithNewSession", reflect.TypeOf((*MockTaskManager)(nil).WithNewSession), arg0)
}

// WithNewTxn mocks base method.
func (m *MockTaskManager) WithNewTxn(arg0 context.Context, arg1 func(sessionctx.Context) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithNewTxn", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// WithNewTxn indicates an expected call of WithNewTxn.
func (mr *MockTaskManagerMockRecorder) WithNewTxn(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithNewTxn", reflect.TypeOf((*MockTaskManager)(nil).WithNewTxn), arg0, arg1)
}
