package sqlmarshaler

import "errors"

// QueryBuilder is the wrapper for all the configuration used by wrapper methods of sql.
// It contains fields that represent SQL clauses. During each operation, each clause will be checked
// in the following order for validity: InsertInto > DeleteFrom > Select. In this checking process,
// if a clause is found to be valid, others will be ignored.
type QueryBuilder struct {
	// Name of the table that data will be inserted to.
	InsertInto string

	// The name of the table from which rows will be deleted. If this clause is non-empty,
	// Where clause must be non-empty.
	DeleteFrom string

	// Columns to be selected. If this clause is non-empty, From clause must be non-empty.
	Select []string

	// Phrases after the FROM clause
	From []string

	// Phrases after the WHERE clause as a string
	Where string

	// Destination of unmarshaling or source of marshaling. If clause is InsertInto, Data
	// can be a slice. If clause is Select and query result contains multiple rows, Data
	// will be a slice; otherwise it will be of the type of the requested object.
	Data interface{}
}

// NewInsertionQuery creates a QueryBuilder for inserting rows.
func NewInsertionQuery(insertInto string, data interface{}) (*QueryBuilder, error) {
	return nil, errors.New("not implemented")
}

// NewDeletionQuery creates a QueryBuilder for deleting rows.
func NewDeletionQuery(deleteFrom string, where string) (*QueryBuilder, error) {
	return nil, errors.New("not implemented")
}

// NewSelectionQuery creates a QueryBuilder for selecting rows.
func NewSelectionQuery(Select []string, data interface{}) (*QueryBuilder, error) {
	return nil, errors.New("not implemented")
}

// IsValid reports whether QueryBuilder is valid or not. If it's valid, returns true and the type
// of the query; otherwise, it returns false and the reason why.
func (qb *QueryBuilder) IsValid() (bool, string) {
	return false, "not implemented"
}

// Build gathers up the clauses in the QueryBuilder and returns the query string that
// this QueryBuilder represents and the arguments to be passed on to sql functions.
func (qb *QueryBuilder) Build() (string, []interface{}) {
	return qb.buildQuery(), qb.buildArgs()
}

// buildQuery returns the sql query string that this QueryBuilder represents
func (qb *QueryBuilder) buildQuery() string {
	return "not implemented"
}

// buildQuery returns the arguments to be passed on to sql functions.
func (qb *QueryBuilder) buildArgs() []interface{} {
	return nil
}
