package sqlmarshaler

import (
	"context"
	"database/sql"
	"errors"
)

type SqlMarshaler struct {
	db *sql.DB
}

// Creates and returns a new SqlMarshaler with the given parameters.
func NewSqlMarshaler(dbURL, driver string) (*SqlMarshaler, error) {
	return nil, errors.New("not implemented")
}

// Exec is ExecContext with empty context.
func (sm *SqlMarshaler) Exec(ctx context.Context, query *QueryBuilder) error {
	return sm.ExecContext(context.Background(), query)
}

func (sm *SqlMarshaler) ExecContext(ctx context.Context, query *QueryBuilder) error {
	return errors.New("not implemented")
}

// Query is QueryContext with empty context.
func (sm *SqlMarshaler) Query(query *QueryBuilder) error {
	return sm.QueryContext(context.Background(), query)
}

func (sm *SqlMarshaler) QueryContext(ctx context.Context, query *QueryBuilder) error {
	return errors.New("not implemented")
}

// QueryRow is QueryRowContext with empty context.
func (sm *SqlMarshaler) QueryRow(query *QueryBuilder) error {
	return sm.QueryRowContext(context.Background(), query)
}

func (sm *SqlMarshaler) QueryRowContext(ctx context.Context, query *QueryBuilder) error {
	return errors.New("not implemented")
}
