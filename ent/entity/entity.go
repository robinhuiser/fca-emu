// Code generated by entc, DO NOT EDIT.

package entity

import (
	"fmt"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the entity type in the database.
	Label = "entity"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldDateCreated holds the string denoting the datecreated field in the database.
	FieldDateCreated = "date_created"
	// FieldFirstname holds the string denoting the firstname field in the database.
	FieldFirstname = "firstname"
	// FieldLastname holds the string denoting the lastname field in the database.
	FieldLastname = "lastname"
	// FieldFullname holds the string denoting the fullname field in the database.
	FieldFullname = "fullname"
	// FieldDateOfBirth holds the string denoting the dateofbirth field in the database.
	FieldDateOfBirth = "date_of_birth"
	// FieldActive holds the string denoting the active field in the database.
	FieldActive = "active"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldLastLoginDate holds the string denoting the lastlogindate field in the database.
	FieldLastLoginDate = "last_login_date"
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "username"
	// FieldToken holds the string denoting the token field in the database.
	FieldToken = "token"
	// FieldURL holds the string denoting the url field in the database.
	FieldURL = "url"
	// EdgeTaxSpecifications holds the string denoting the taxspecifications edge name in mutations.
	EdgeTaxSpecifications = "taxSpecifications"
	// EdgeAddresses holds the string denoting the addresses edge name in mutations.
	EdgeAddresses = "addresses"
	// EdgePreferences holds the string denoting the preferences edge name in mutations.
	EdgePreferences = "preferences"
	// EdgeContactPoints holds the string denoting the contactpoints edge name in mutations.
	EdgeContactPoints = "contactPoints"
	// EdgeOwnsAccount holds the string denoting the owns_account edge name in mutations.
	EdgeOwnsAccount = "owns_account"
	// Table holds the table name of the entity in the database.
	Table = "entities"
	// TaxSpecificationsTable is the table the holds the taxSpecifications relation/edge.
	TaxSpecificationsTable = "entity_tax_informations"
	// TaxSpecificationsInverseTable is the table name for the EntityTaxInformation entity.
	// It exists in this package in order to avoid circular dependency with the "entitytaxinformation" package.
	TaxSpecificationsInverseTable = "entity_tax_informations"
	// TaxSpecificationsColumn is the table column denoting the taxSpecifications relation/edge.
	TaxSpecificationsColumn = "entity_tax_specifications"
	// AddressesTable is the table the holds the addresses relation/edge.
	AddressesTable = "entity_addresses"
	// AddressesInverseTable is the table name for the EntityAddress entity.
	// It exists in this package in order to avoid circular dependency with the "entityaddress" package.
	AddressesInverseTable = "entity_addresses"
	// AddressesColumn is the table column denoting the addresses relation/edge.
	AddressesColumn = "entity_addresses"
	// PreferencesTable is the table the holds the preferences relation/edge.
	PreferencesTable = "preferences"
	// PreferencesInverseTable is the table name for the Preference entity.
	// It exists in this package in order to avoid circular dependency with the "preference" package.
	PreferencesInverseTable = "preferences"
	// PreferencesColumn is the table column denoting the preferences relation/edge.
	PreferencesColumn = "entity_preferences"
	// ContactPointsTable is the table the holds the contactPoints relation/edge.
	ContactPointsTable = "entity_contact_points"
	// ContactPointsInverseTable is the table name for the EntityContactPoint entity.
	// It exists in this package in order to avoid circular dependency with the "entitycontactpoint" package.
	ContactPointsInverseTable = "entity_contact_points"
	// ContactPointsColumn is the table column denoting the contactPoints relation/edge.
	ContactPointsColumn = "entity_contact_points"
	// OwnsAccountTable is the table the holds the owns_account relation/edge. The primary key declared below.
	OwnsAccountTable = "account_owners"
	// OwnsAccountInverseTable is the table name for the Account entity.
	// It exists in this package in order to avoid circular dependency with the "account" package.
	OwnsAccountInverseTable = "accounts"
)

// Columns holds all SQL columns for entity fields.
var Columns = []string{
	FieldID,
	FieldDateCreated,
	FieldFirstname,
	FieldLastname,
	FieldFullname,
	FieldDateOfBirth,
	FieldActive,
	FieldType,
	FieldLastLoginDate,
	FieldUsername,
	FieldToken,
	FieldURL,
}

var (
	// OwnsAccountPrimaryKey and OwnsAccountColumn2 are the table columns denoting the
	// primary key for the owns_account relation (M2M).
	OwnsAccountPrimaryKey = []string{"account_id", "entity_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultActive holds the default value on creation for the "active" field.
	DefaultActive bool
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// Type defines the type for the "type" enum field.
type Type string

// Type values.
const (
	TypePERSON   Type = "PERSON"
	TypeBUSINESS Type = "BUSINESS"
	TypeSYSTEM   Type = "SYSTEM"
)

func (_type Type) String() string {
	return string(_type)
}

// TypeValidator is a validator for the "type" field enum values. It is called by the builders before save.
func TypeValidator(_type Type) error {
	switch _type {
	case TypePERSON, TypeBUSINESS, TypeSYSTEM:
		return nil
	default:
		return fmt.Errorf("entity: invalid enum value for type field: %q", _type)
	}
}
