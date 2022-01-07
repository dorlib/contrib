// Code generated by entc, DO NOT EDIT.

package nilexample

const (
	// Label holds the string label denoting the nilexample type in the database.
	Label = "nil_example"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldStrNil holds the string denoting the str_nil field in the database.
	FieldStrNil = "str_nil"
	// FieldTimeNil holds the string denoting the time_nil field in the database.
	FieldTimeNil = "time_nil"
	// Table holds the table name of the nilexample in the database.
	Table = "nil_examples"
)

// Columns holds all SQL columns for nilexample fields.
var Columns = []string{
	FieldID,
	FieldStrNil,
	FieldTimeNil,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}