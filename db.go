package unitofwork

import (
	"database/sql"
	"fmt"
	"sync"

	"github.com/pkg/errors"
)

// Db defines the contract of a type that can execute commands against a database
type Db interface {
	Exec(command string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}

// TransactionalDb extends the Db interface to add the ability to begin a transaction
type TransactionalDb interface {
	Db
	Begin() (Txn, error)
}

// NewDatabaseUnitOfWork creates and initialized a new unit of work referencing the provided transactional database
func NewDatabaseUnitOfWork(db TransactionalDb) UnitOfWork {
	uow := databaseUnitOfWork{
		dbConn: db,
	}
	return &uow
}

type databaseUnitOfWork struct {
	mtx        sync.Mutex
	dbConn     TransactionalDb
	currentTxn Txn
	txnDepth   int
	aborted    bool
}

func (w *databaseUnitOfWork) Begin() (UnitOfWork, error) {
	w.mtx.Lock()
	defer w.mtx.Unlock()
	if w.currentTxn == nil {
		t, err := w.dbConn.Begin()
		if err != nil {
			return nil, err
		}
		w.currentTxn = t
	}
	w.txnDepth++
	return w, nil
}

func (w *databaseUnitOfWork) Database() Db {
	w.mtx.Lock()
	defer w.mtx.Unlock()
	if w.currentTxn != nil {
		return w.currentTxn
	}
	return w.dbConn
}

func (w *databaseUnitOfWork) Commit() error {
	w.mtx.Lock()
	defer w.mtx.Unlock()
	if w.aborted {
		return errors.New("Cannot commit an aborted unit of work")
	}
	if w.currentTxn == nil {
		return errors.New("A unit of work has not been started")
	}
	if w.txnDepth > 0 {
		w.txnDepth--
	}
	if w.txnDepth < 1 {
		if e := w.currentTxn.Commit(); e != nil {
			return fmt.Errorf("Error committing unit of work (%s)", e)
		}
		w.currentTxn = nil
	}
	return nil
}

func (w *databaseUnitOfWork) Rollback() error {
	w.mtx.Lock()
	defer w.mtx.Unlock()
	if w.currentTxn == nil {
		return errors.New("A unit of work has not been started")
	}
	w.aborted = true
	w.txnDepth = 0
	if e := w.currentTxn.Rollback(); e != nil {
		return fmt.Errorf("Error rolling back unit of work (%s)", e)
	}
	w.currentTxn = nil
	return nil
}

func (w *databaseUnitOfWork) Execute(action Op) error {
	return performOp(w, action)
}
