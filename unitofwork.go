package unitofwork

import (
	"github.com/pkg/errors"
)

// UnitOfWork defines a composite set of operations that should be committed or discarded idempotently
type UnitOfWork interface {
	// returns the "active" database connection
	Database() Db

	// starts a new transaction for this unit of work if one is not already present
	Begin() (UnitOfWork, error)
	// commits the active transaction
	Commit() error
	// aborts the active transaction, discarding any changes
	Rollback() error

	// executes the specified action passing in the active database connection, which should be used
	// to perform the relevant database commands.
	Execute(action Op) error
}

// Op defines a function that performs an action against the passed-in database.  If this function returns
// a non-nil error, the active transaction is marked invalid and will be aborted as the call to Execute() unwinds.
type Op func(Db) error

// performOp executes the specified action within a unit of work, committing or rolling back all changes atomically
func performOp(uow UnitOfWork, action Op) (err error) {
	txn, err := uow.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			switch t := p.(type) {
			case error:
				err = t
			default:
				err = errors.Errorf("%s", t)
			}
		}
		if err != nil {
			if e := txn.Rollback(); e != nil {
				err = errors.Wrap(e, "Unable to abort the transaction")
			}
			return
		}
		err = txn.Commit()
	}()
	return action(uow.Database())
}
