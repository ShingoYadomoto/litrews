package db

import (
	"context"

	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

// DB wrap sqlx.DB
type DB struct {
	*sqlx.DB
}

// Transaction is transaction manager.
func (db *DB) Transaction(f func(sc SchemaContext) error) error {
	tx, err := db.Beginx()
	if err != nil {
		return NewTransactionErrorCause(err)
	}

	defer func() {
		err := recover()
		if err != nil {
			if txerr := tx.Rollback(); txerr != nil {
				panic(fmt.Sprintf("%s; %s", err, txerr))
			} else {
				panic(err)
			}
		}
	}()

	err = f(schemaContextFromTx(tx))
	if err != nil {
		if txerr := tx.Rollback(); txerr != nil {
			return NewTransactionError(err, txerr)
		} else {
			return NewTransactionErrorCause(err)
		}
	}

	tx.Commit()

	return err
}

// Tx wrap sqlx.Tx
type Tx struct {
	*sqlx.Tx
}

// Transaction is transaction manager.
func (tx *Tx) Transaction(f func(sc SchemaContext) error) error {
	panic("Tx does not support Transaction()")
}

type TransactionError struct {
	Cause    error
	Rollback error
}

func (err *TransactionError) Error() string {
	return fmt.Sprintf("TransactionError{ cause: %+v, rollback: %+v }", err.Cause, err.Rollback)
}

func NewTransactionError(cause error, rollback error) *TransactionError {
	return &TransactionError{cause, rollback}
}

func NewTransactionErrorCause(cause error) *TransactionError {
	return &TransactionError{cause, nil}
}

// AbstractDB is abstract db type.
type AbstractDB interface {
	BindNamed(query string, arg interface{}) (string, []interface{}, error)                       // DB/Tx
	DriverName() string                                                                           // DB/Tx
	Exec(query string, args ...interface{}) (sql.Result, error)                                   // DB/Tx
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)       // DB/Tx
	Get(dest interface{}, query string, args ...interface{}) error                                // DB/Tx
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error    // DB/Tx
	MustExec(query string, args ...interface{}) sql.Result                                        // DB/Tx
	MustExecContext(ctx context.Context, query string, args ...interface{}) sql.Result            // DB/Tx
	NamedExec(query string, arg interface{}) (sql.Result, error)                                  // DB/Tx
	NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error)      // DB/Tx
	NamedQuery(query string, arg interface{}) (*sqlx.Rows, error)                                 // DB/Tx
	Prepare(query string) (*sql.Stmt, error)                                                      // DB/Tx
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)                          // DB/Tx
	PrepareNamed(query string) (*sqlx.NamedStmt, error)                                           // DB/Tx
	Preparex(query string) (*sqlx.Stmt, error)                                                    // DB/Tx
	Query(query string, args ...interface{}) (*sql.Rows, error)                                   // DB/Tx
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)       // DB/Tx
	QueryRow(query string, args ...interface{}) *sql.Row                                          // DB/Tx
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row              // DB/Tx
	QueryRowx(query string, args ...interface{}) *sqlx.Row                                        // DB/Tx
	QueryRowxContext(ctx context.Context, query string, args ...interface{}) *sqlx.Row            // DB/Tx
	Queryx(query string, args ...interface{}) (*sqlx.Rows, error)                                 // DB/Tx
	QueryxContext(ctx context.Context, query string, args ...interface{}) (*sqlx.Rows, error)     // DB/Tx
	Rebind(query string) string                                                                   // DB/Tx
	Select(dest interface{}, query string, args ...interface{}) error                             // DB/Tx
	SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error // DB/Tx
	Transaction(f func(sc SchemaContext) error) error
}

// SchemaContext contains AbstractDB.
type SchemaContext struct {
	db AbstractDB
}

// NewSchemaContext is constructor.
func NewSchemaContext(db *sqlx.DB) SchemaContext {
	return SchemaContext{db: &DB{db}}
}

func schemaContextFromTx(tx *sqlx.Tx) SchemaContext {
	return SchemaContext{db: &Tx{tx}}
}

// DB is accessor to db in SchemaContext.
func (sc SchemaContext) DB() AbstractDB {
	return sc.db
}
