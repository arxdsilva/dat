package dat

// Result serves the same purpose as sql.Result. Defining
// it for the package avoids tight coupling with database/sql.
type Result struct {
	LastInsertID int64
	RowsAffected int64
}

// Execer is any object that executes and queries SQL.
type Execer interface {
	Exec() (*Result, error)
	QueryScalar(destinations ...interface{}) error
	QuerySlice(dest interface{}) error
	QueryStruct(dest interface{}) error
	QueryStructs(dest interface{}) error
}

const panicExecerMsg = "dat builders are disconnected, use a runner package"

// panicExecer is the execer for instances of dat builders from
// data package.
type panicExecer struct{}

// Exec panics when Exec is called.
func (nop *panicExecer) Exec() (*Result, error) {
	panic(panicExecerMsg)
}

// QueryScalar panics when QueryScalar is called.
func (nop *panicExecer) QueryScalar(destinations ...interface{}) error {
	panic(panicExecerMsg)
}

// QuerySlice panics when QuerySlice is called.
func (nop *panicExecer) QuerySlice(dest interface{}) error {
	panic(panicExecerMsg)
}

// QueryStruct panics when QueryStruct is called.
func (nop *panicExecer) QueryStruct(dest interface{}) error {
	panic(panicExecerMsg)
}

// QueryStructs panics when QueryStructs is called.
func (nop *panicExecer) QueryStructs(dest interface{}) error {
	panic(panicExecerMsg)
}