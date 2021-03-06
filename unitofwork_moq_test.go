// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package unitofwork

import (
	"sync"
)

var (
	lockUnitOfWorkMockBegin    sync.RWMutex
	lockUnitOfWorkMockCommit   sync.RWMutex
	lockUnitOfWorkMockDatabase sync.RWMutex
	lockUnitOfWorkMockExecute  sync.RWMutex
	lockUnitOfWorkMockRollback sync.RWMutex
)

// Ensure, that UnitOfWorkMock does implement UnitOfWork.
// If this is not the case, regenerate this file with moq.
var _ UnitOfWork = &UnitOfWorkMock{}

// UnitOfWorkMock is a mock implementation of UnitOfWork.
//
//     func TestSomethingThatUsesUnitOfWork(t *testing.T) {
//
//         // make and configure a mocked UnitOfWork
//         mockedUnitOfWork := &UnitOfWorkMock{
//             BeginFunc: func() (UnitOfWork, error) {
// 	               panic("mock out the Begin method")
//             },
//             CommitFunc: func() error {
// 	               panic("mock out the Commit method")
//             },
//             DatabaseFunc: func() Db {
// 	               panic("mock out the Database method")
//             },
//             ExecuteFunc: func(action Op) error {
// 	               panic("mock out the Execute method")
//             },
//             RollbackFunc: func() error {
// 	               panic("mock out the Rollback method")
//             },
//         }
//
//         // use mockedUnitOfWork in code that requires UnitOfWork
//         // and then make assertions.
//
//     }
type UnitOfWorkMock struct {
	// BeginFunc mocks the Begin method.
	BeginFunc func() (UnitOfWork, error)

	// CommitFunc mocks the Commit method.
	CommitFunc func() error

	// DatabaseFunc mocks the Database method.
	DatabaseFunc func() Db

	// ExecuteFunc mocks the Execute method.
	ExecuteFunc func(action Op) error

	// RollbackFunc mocks the Rollback method.
	RollbackFunc func() error

	// calls tracks calls to the methods.
	calls struct {
		// Begin holds details about calls to the Begin method.
		Begin []struct {
		}
		// Commit holds details about calls to the Commit method.
		Commit []struct {
		}
		// Database holds details about calls to the Database method.
		Database []struct {
		}
		// Execute holds details about calls to the Execute method.
		Execute []struct {
			// Action is the action argument value.
			Action Op
		}
		// Rollback holds details about calls to the Rollback method.
		Rollback []struct {
		}
	}
}

// Begin calls BeginFunc.
func (mock *UnitOfWorkMock) Begin() (UnitOfWork, error) {
	if mock.BeginFunc == nil {
		panic("UnitOfWorkMock.BeginFunc: method is nil but UnitOfWork.Begin was just called")
	}
	callInfo := struct {
	}{}
	lockUnitOfWorkMockBegin.Lock()
	mock.calls.Begin = append(mock.calls.Begin, callInfo)
	lockUnitOfWorkMockBegin.Unlock()
	return mock.BeginFunc()
}

// BeginCalls gets all the calls that were made to Begin.
// Check the length with:
//     len(mockedUnitOfWork.BeginCalls())
func (mock *UnitOfWorkMock) BeginCalls() []struct {
} {
	var calls []struct {
	}
	lockUnitOfWorkMockBegin.RLock()
	calls = mock.calls.Begin
	lockUnitOfWorkMockBegin.RUnlock()
	return calls
}

// Commit calls CommitFunc.
func (mock *UnitOfWorkMock) Commit() error {
	if mock.CommitFunc == nil {
		panic("UnitOfWorkMock.CommitFunc: method is nil but UnitOfWork.Commit was just called")
	}
	callInfo := struct {
	}{}
	lockUnitOfWorkMockCommit.Lock()
	mock.calls.Commit = append(mock.calls.Commit, callInfo)
	lockUnitOfWorkMockCommit.Unlock()
	return mock.CommitFunc()
}

// CommitCalls gets all the calls that were made to Commit.
// Check the length with:
//     len(mockedUnitOfWork.CommitCalls())
func (mock *UnitOfWorkMock) CommitCalls() []struct {
} {
	var calls []struct {
	}
	lockUnitOfWorkMockCommit.RLock()
	calls = mock.calls.Commit
	lockUnitOfWorkMockCommit.RUnlock()
	return calls
}

// Database calls DatabaseFunc.
func (mock *UnitOfWorkMock) Database() Db {
	if mock.DatabaseFunc == nil {
		panic("UnitOfWorkMock.DatabaseFunc: method is nil but UnitOfWork.Database was just called")
	}
	callInfo := struct {
	}{}
	lockUnitOfWorkMockDatabase.Lock()
	mock.calls.Database = append(mock.calls.Database, callInfo)
	lockUnitOfWorkMockDatabase.Unlock()
	return mock.DatabaseFunc()
}

// DatabaseCalls gets all the calls that were made to Database.
// Check the length with:
//     len(mockedUnitOfWork.DatabaseCalls())
func (mock *UnitOfWorkMock) DatabaseCalls() []struct {
} {
	var calls []struct {
	}
	lockUnitOfWorkMockDatabase.RLock()
	calls = mock.calls.Database
	lockUnitOfWorkMockDatabase.RUnlock()
	return calls
}

// Execute calls ExecuteFunc.
func (mock *UnitOfWorkMock) Execute(action Op) error {
	if mock.ExecuteFunc == nil {
		panic("UnitOfWorkMock.ExecuteFunc: method is nil but UnitOfWork.Execute was just called")
	}
	callInfo := struct {
		Action Op
	}{
		Action: action,
	}
	lockUnitOfWorkMockExecute.Lock()
	mock.calls.Execute = append(mock.calls.Execute, callInfo)
	lockUnitOfWorkMockExecute.Unlock()
	return mock.ExecuteFunc(action)
}

// ExecuteCalls gets all the calls that were made to Execute.
// Check the length with:
//     len(mockedUnitOfWork.ExecuteCalls())
func (mock *UnitOfWorkMock) ExecuteCalls() []struct {
	Action Op
} {
	var calls []struct {
		Action Op
	}
	lockUnitOfWorkMockExecute.RLock()
	calls = mock.calls.Execute
	lockUnitOfWorkMockExecute.RUnlock()
	return calls
}

// Rollback calls RollbackFunc.
func (mock *UnitOfWorkMock) Rollback() error {
	if mock.RollbackFunc == nil {
		panic("UnitOfWorkMock.RollbackFunc: method is nil but UnitOfWork.Rollback was just called")
	}
	callInfo := struct {
	}{}
	lockUnitOfWorkMockRollback.Lock()
	mock.calls.Rollback = append(mock.calls.Rollback, callInfo)
	lockUnitOfWorkMockRollback.Unlock()
	return mock.RollbackFunc()
}

// RollbackCalls gets all the calls that were made to Rollback.
// Check the length with:
//     len(mockedUnitOfWork.RollbackCalls())
func (mock *UnitOfWorkMock) RollbackCalls() []struct {
} {
	var calls []struct {
	}
	lockUnitOfWorkMockRollback.RLock()
	calls = mock.calls.Rollback
	lockUnitOfWorkMockRollback.RUnlock()
	return calls
}
