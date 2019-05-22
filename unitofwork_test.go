package unitofwork

import (
	"database/sql"
	"testing"
)

//go:generate moq -out unitofwork_moq_test.go . UnitOfWork
//go:generate moq -out db_moq_test.go . Db
//go:generate moq -out txn_moq_test.go . Txn
//go:generate moq -out txndb_moq_test.go . TransactionalDb

func TestScratch(t *testing.T) {
	uow := newMockUnitOfWork()

	err := uow.Execute(func(Db) error { return nil })
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	mk := uow.(*UnitOfWorkMock)
	if n := len(mk.BeginCalls()); n != 1 {
		t.Errorf("Expected Begin() to be called 1 time, actually called %v times", n)
	}
}

func newMockUnitOfWork() UnitOfWork {
	uow := &UnitOfWorkMock{}
	uow.DatabaseFunc = func() Db { return newMockDb() }
	uow.BeginFunc = func() (UnitOfWork, error) { return uow, nil }
	uow.CommitFunc = func() error { return nil }
	uow.RollbackFunc = func() error { return nil }
	uow.ExecuteFunc = func(action Op) error { return performOp(uow, action) }
	return uow
}

func newMockDb() Db {
	return &DbMock{
		ExecFunc: func(_ string, _ ...interface{}) (sql.Result, error) {
			return newMockResult(0, 0, nil), nil
		},
		QueryFunc: func(_ string, _ ...interface{}) (*sql.Rows, error) {
			return &sql.Rows{}, nil
		},
		QueryRowFunc: func(_ string, _ ...interface{}) *sql.Row {
			return &sql.Row{}
		},
	}
}

func newMockTxn() Txn {
	return &TxnMock{
		ExecFunc: func(_ string, _ ...interface{}) (sql.Result, error) {
			return newMockResult(0, 0, nil), nil
		},
		QueryFunc: func(_ string, _ ...interface{}) (*sql.Rows, error) {
			return &sql.Rows{}, nil
		},
		QueryRowFunc: func(_ string, _ ...interface{}) *sql.Row {
			return &sql.Row{}
		},
		CommitFunc: func() error {
			return nil
		},
		RollbackFunc: func() error {
			return nil
		},
	}
}

func newMockResult(id, n int64, e error) sql.Result {
	return mockSqlResult{id: id, n: n, e: e}
}

type mockSqlResult struct {
	id, n int64
	e     error
}

func (r mockSqlResult) LastInsertId() (int64, error) {
	if r.e != nil {
		return 0, r.e
	}
	return r.id, nil
}
func (r mockSqlResult) RowsAffected() (int64, error) {
	if r.e != nil {
		return 0, r.e
	}
	return r.n, nil
}
