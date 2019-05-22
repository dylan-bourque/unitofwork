// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package unitofwork

import (
	"database/sql"
	"sync"
)

var (
	lockTransactionalDbMockBegin    sync.RWMutex
	lockTransactionalDbMockExec     sync.RWMutex
	lockTransactionalDbMockQuery    sync.RWMutex
	lockTransactionalDbMockQueryRow sync.RWMutex
)

// Ensure, that TransactionalDbMock does implement TransactionalDb.
// If this is not the case, regenerate this file with moq.
var _ TransactionalDb = &TransactionalDbMock{}

// TransactionalDbMock is a mock implementation of TransactionalDb.
//
//     func TestSomethingThatUsesTransactionalDb(t *testing.T) {
//
//         // make and configure a mocked TransactionalDb
//         mockedTransactionalDb := &TransactionalDbMock{
//             BeginFunc: func() (Txn, error) {
// 	               panic("mock out the Begin method")
//             },
//             ExecFunc: func(command string, args ...interface{}) (sql.Result, error) {
// 	               panic("mock out the Exec method")
//             },
//             QueryFunc: func(query string, args ...interface{}) (*sql.Rows, error) {
// 	               panic("mock out the Query method")
//             },
//             QueryRowFunc: func(query string, args ...interface{}) *sql.Row {
// 	               panic("mock out the QueryRow method")
//             },
//         }
//
//         // use mockedTransactionalDb in code that requires TransactionalDb
//         // and then make assertions.
//
//     }
type TransactionalDbMock struct {
	// BeginFunc mocks the Begin method.
	BeginFunc func() (Txn, error)

	// ExecFunc mocks the Exec method.
	ExecFunc func(command string, args ...interface{}) (sql.Result, error)

	// QueryFunc mocks the Query method.
	QueryFunc func(query string, args ...interface{}) (*sql.Rows, error)

	// QueryRowFunc mocks the QueryRow method.
	QueryRowFunc func(query string, args ...interface{}) *sql.Row

	// calls tracks calls to the methods.
	calls struct {
		// Begin holds details about calls to the Begin method.
		Begin []struct {
		}
		// Exec holds details about calls to the Exec method.
		Exec []struct {
			// Command is the command argument value.
			Command string
			// Args is the args argument value.
			Args []interface{}
		}
		// Query holds details about calls to the Query method.
		Query []struct {
			// Query is the query argument value.
			Query string
			// Args is the args argument value.
			Args []interface{}
		}
		// QueryRow holds details about calls to the QueryRow method.
		QueryRow []struct {
			// Query is the query argument value.
			Query string
			// Args is the args argument value.
			Args []interface{}
		}
	}
}

// Begin calls BeginFunc.
func (mock *TransactionalDbMock) Begin() (Txn, error) {
	if mock.BeginFunc == nil {
		panic("TransactionalDbMock.BeginFunc: method is nil but TransactionalDb.Begin was just called")
	}
	callInfo := struct {
	}{}
	lockTransactionalDbMockBegin.Lock()
	mock.calls.Begin = append(mock.calls.Begin, callInfo)
	lockTransactionalDbMockBegin.Unlock()
	return mock.BeginFunc()
}

// BeginCalls gets all the calls that were made to Begin.
// Check the length with:
//     len(mockedTransactionalDb.BeginCalls())
func (mock *TransactionalDbMock) BeginCalls() []struct {
} {
	var calls []struct {
	}
	lockTransactionalDbMockBegin.RLock()
	calls = mock.calls.Begin
	lockTransactionalDbMockBegin.RUnlock()
	return calls
}

// Exec calls ExecFunc.
func (mock *TransactionalDbMock) Exec(command string, args ...interface{}) (sql.Result, error) {
	if mock.ExecFunc == nil {
		panic("TransactionalDbMock.ExecFunc: method is nil but TransactionalDb.Exec was just called")
	}
	callInfo := struct {
		Command string
		Args    []interface{}
	}{
		Command: command,
		Args:    args,
	}
	lockTransactionalDbMockExec.Lock()
	mock.calls.Exec = append(mock.calls.Exec, callInfo)
	lockTransactionalDbMockExec.Unlock()
	return mock.ExecFunc(command, args...)
}

// ExecCalls gets all the calls that were made to Exec.
// Check the length with:
//     len(mockedTransactionalDb.ExecCalls())
func (mock *TransactionalDbMock) ExecCalls() []struct {
	Command string
	Args    []interface{}
} {
	var calls []struct {
		Command string
		Args    []interface{}
	}
	lockTransactionalDbMockExec.RLock()
	calls = mock.calls.Exec
	lockTransactionalDbMockExec.RUnlock()
	return calls
}

// Query calls QueryFunc.
func (mock *TransactionalDbMock) Query(query string, args ...interface{}) (*sql.Rows, error) {
	if mock.QueryFunc == nil {
		panic("TransactionalDbMock.QueryFunc: method is nil but TransactionalDb.Query was just called")
	}
	callInfo := struct {
		Query string
		Args  []interface{}
	}{
		Query: query,
		Args:  args,
	}
	lockTransactionalDbMockQuery.Lock()
	mock.calls.Query = append(mock.calls.Query, callInfo)
	lockTransactionalDbMockQuery.Unlock()
	return mock.QueryFunc(query, args...)
}

// QueryCalls gets all the calls that were made to Query.
// Check the length with:
//     len(mockedTransactionalDb.QueryCalls())
func (mock *TransactionalDbMock) QueryCalls() []struct {
	Query string
	Args  []interface{}
} {
	var calls []struct {
		Query string
		Args  []interface{}
	}
	lockTransactionalDbMockQuery.RLock()
	calls = mock.calls.Query
	lockTransactionalDbMockQuery.RUnlock()
	return calls
}

// QueryRow calls QueryRowFunc.
func (mock *TransactionalDbMock) QueryRow(query string, args ...interface{}) *sql.Row {
	if mock.QueryRowFunc == nil {
		panic("TransactionalDbMock.QueryRowFunc: method is nil but TransactionalDb.QueryRow was just called")
	}
	callInfo := struct {
		Query string
		Args  []interface{}
	}{
		Query: query,
		Args:  args,
	}
	lockTransactionalDbMockQueryRow.Lock()
	mock.calls.QueryRow = append(mock.calls.QueryRow, callInfo)
	lockTransactionalDbMockQueryRow.Unlock()
	return mock.QueryRowFunc(query, args...)
}

// QueryRowCalls gets all the calls that were made to QueryRow.
// Check the length with:
//     len(mockedTransactionalDb.QueryRowCalls())
func (mock *TransactionalDbMock) QueryRowCalls() []struct {
	Query string
	Args  []interface{}
} {
	var calls []struct {
		Query string
		Args  []interface{}
	}
	lockTransactionalDbMockQueryRow.RLock()
	calls = mock.calls.QueryRow
	lockTransactionalDbMockQueryRow.RUnlock()
	return calls
}
