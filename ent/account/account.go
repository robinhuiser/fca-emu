// Code generated by entc, DO NOT EDIT.

package account

import (
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the account type in the database.
	Label = "account"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldNumber holds the string denoting the number field in the database.
	FieldNumber = "number"
	// FieldParentId holds the string denoting the parentid field in the database.
	FieldParentId = "parent_id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldIban holds the string denoting the iban field in the database.
	FieldIban = "iban"
	// FieldDateCreated holds the string denoting the datecreated field in the database.
	FieldDateCreated = "date_created"
	// FieldDateOpened holds the string denoting the dateopened field in the database.
	FieldDateOpened = "date_opened"
	// FieldDateLastUpdated holds the string denoting the datelastupdated field in the database.
	FieldDateLastUpdated = "date_last_updated"
	// FieldDateClosed holds the string denoting the dateclosed field in the database.
	FieldDateClosed = "date_closed"
	// FieldCurrencyCode holds the string denoting the currencycode field in the database.
	FieldCurrencyCode = "currency_code"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldSource holds the string denoting the source field in the database.
	FieldSource = "source"
	// FieldInterestReporting holds the string denoting the interestreporting field in the database.
	FieldInterestReporting = "interest_reporting"
	// Table holds the table name of the account in the database.
	Table = "accounts"
)

// Columns holds all SQL columns for account fields.
var Columns = []string{
	FieldID,
	FieldType,
	FieldNumber,
	FieldParentId,
	FieldName,
	FieldTitle,
	FieldIban,
	FieldDateCreated,
	FieldDateOpened,
	FieldDateLastUpdated,
	FieldDateClosed,
	FieldCurrencyCode,
	FieldStatus,
	FieldSource,
	FieldInterestReporting,
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

var (
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
