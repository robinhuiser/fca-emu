// Code generated by entc, DO NOT EDIT.

package product

import (
	"fmt"
)

const (
	// Label holds the string label denoting the product type in the database.
	Label = "product"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldTypeName holds the string denoting the typename field in the database.
	FieldTypeName = "type_name"
	// FieldSubType holds the string denoting the subtype field in the database.
	FieldSubType = "sub_type"
	// FieldSubTypeName holds the string denoting the subtypename field in the database.
	FieldSubTypeName = "sub_type_name"
	// FieldURL holds the string denoting the url field in the database.
	FieldURL = "url"
	// Table holds the table name of the product in the database.
	Table = "products"
)

// Columns holds all SQL columns for product fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldType,
	FieldTypeName,
	FieldSubType,
	FieldSubTypeName,
	FieldURL,
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

// Type defines the type for the "type" enum field.
type Type string

// Type values.
const (
	TypeSAVING     Type = "SAVING"
	TypeLOAN       Type = "LOAN"
	TypeDEPOSIT    Type = "DEPOSIT"
	TypeCHECKING   Type = "CHECKING"
	TypeINVESTMENT Type = "INVESTMENT"
)

func (_type Type) String() string {
	return string(_type)
}

// TypeValidator is a validator for the "type" field enum values. It is called by the builders before save.
func TypeValidator(_type Type) error {
	switch _type {
	case TypeSAVING, TypeLOAN, TypeDEPOSIT, TypeCHECKING, TypeINVESTMENT:
		return nil
	default:
		return fmt.Errorf("product: invalid enum value for type field: %q", _type)
	}
}
