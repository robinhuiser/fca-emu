// Code generated by entc, DO NOT EDIT.

package entitytaxinformation

import (
	"fmt"
)

const (
	// Label holds the string label denoting the entitytaxinformation type in the database.
	Label = "entity_tax_information"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldTaxId holds the string denoting the taxid field in the database.
	FieldTaxId = "tax_id"
	// Table holds the table name of the entitytaxinformation in the database.
	Table = "entity_tax_informations"
)

// Columns holds all SQL columns for entitytaxinformation fields.
var Columns = []string{
	FieldID,
	FieldType,
	FieldTaxId,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "entity_tax_informations"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"entity_tax_specifications",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

// Type defines the type for the "type" enum field.
type Type string

// Type values.
const (
	TypeSSN   Type = "SSN"
	TypeEIN   Type = "EIN"
	TypeITIN  Type = "ITIN"
	TypeATIN  Type = "ATIN"
	TypeOTHER Type = "OTHER"
)

func (_type Type) String() string {
	return string(_type)
}

// TypeValidator is a validator for the "type" field enum values. It is called by the builders before save.
func TypeValidator(_type Type) error {
	switch _type {
	case TypeSSN, TypeEIN, TypeITIN, TypeATIN, TypeOTHER:
		return nil
	default:
		return fmt.Errorf("entitytaxinformation: invalid enum value for type field: %q", _type)
	}
}
