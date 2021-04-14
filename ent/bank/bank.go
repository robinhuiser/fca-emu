// Code generated by entc, DO NOT EDIT.

package bank

const (
	// Label holds the string label denoting the bank type in the database.
	Label = "bank"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldBankCode holds the string denoting the bankcode field in the database.
	FieldBankCode = "bank_code"
	// FieldBankName holds the string denoting the bankname field in the database.
	FieldBankName = "bank_name"
	// FieldURL holds the string denoting the url field in the database.
	FieldURL = "url"
	// FieldSwift holds the string denoting the swift field in the database.
	FieldSwift = "swift"
	// EdgeBranches holds the string denoting the branches edge name in mutations.
	EdgeBranches = "branches"
	// Table holds the table name of the bank in the database.
	Table = "banks"
	// BranchesTable is the table the holds the branches relation/edge.
	BranchesTable = "branches"
	// BranchesInverseTable is the table name for the Branch entity.
	// It exists in this package in order to avoid circular dependency with the "branch" package.
	BranchesInverseTable = "branches"
	// BranchesColumn is the table column denoting the branches relation/edge.
	BranchesColumn = "bank_branches"
)

// Columns holds all SQL columns for bank fields.
var Columns = []string{
	FieldID,
	FieldBankCode,
	FieldBankName,
	FieldURL,
	FieldSwift,
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