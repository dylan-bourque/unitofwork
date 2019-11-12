# Package `unitofwork`
Package `unitofwork` provides scaffolding code around a logical transaction (in this case, backed by
an actual `database/sql.Tx`).

This implementation of the unit of work idiom provides several things:
- A means to decouple logic code from the actual database (via the `unitofwork.Db` interface)
- A mechanism for supporting "nested" transactions (implemented via reference counting, since most native
  database drivers don't support them)
- Automatic transaction rollback if any individual operation fails or panics

## Example Usage
``` go
import "database/sql"

type sqlDb struct {
    *sql.DB
}
// Begin adapts the standard library's 'database/sql' package to be compatible with 'unitofwork.Txn'
func (d sqlDb) Begin() (unitofwork.Txn, error) {
    txn, err := d.DB.Begin()
    // *database/sql.Tx satisfies unitofwork.Txn so you can return it directly
    return txn, err
}

type fooRepo struct {
    uow unitofwork.UnitOfWork
}
func (r *fooRepo) CreateFoo() (int, error) {
    var id int
    err := r.uow.Execute(func(db unitofwork.Db) error {
        // use db to perform relevant database commands bound to the transaction, assign the key of the new record to id
        return nil
    })
    return id, err
}

type barRepo struct {
    uow unitofwork.UnitOfWork
}
func (r *barRepo) CreateBar(fooID int) (int, error) {
    var id int
    err := r.uow.Execute(func(db unitofwork.Db) error {
        // use db to perform relevant database commands bound to the transaction, using fooID as a foreign key, and assign the key of the new record to id
        return nil
    })
    return id, err
}

func doWork(db unitofwork.TransactionalDb) error {
    // the "database connection" should be created and managed at application level and passed in to this function
    uow := unitofwork.NewDatabaseUnitOfWork(db)
    // starts a new transaction
    // . if the provided function returns nil, the transaction will be committed
    // . if any error is returned or the function panics, the transaction is aborted
    return uow.Execute(func(txn unitofwork.Db) error {
        foo := fooRepo{uow: uow}
        bar := barRepo{uow: uow}
        fooID, err := foo.CreateFoo()
        if err != nil {
            return err
        }
        _, err = bar.CreateBar(fooID)
        if err != nil {
            return err
        }
        return nil
    })
}
```
