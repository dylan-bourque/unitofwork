package unitofwork

// Txn declares the common behaviors between sql.DB and sql.Tx used by the unit-of-work framework so that they can be treated interchangeably
type Txn interface {
	Db
	Commit() error
	Rollback() error
}
