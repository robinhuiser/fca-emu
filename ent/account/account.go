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
	// FieldCurrentBalance holds the string denoting the currentbalance field in the database.
	FieldCurrentBalance = "current_balance"
	// FieldAvailableBalance holds the string denoting the availablebalance field in the database.
	FieldAvailableBalance = "available_balance"
	// FieldURL holds the string denoting the url field in the database.
	FieldURL = "url"
	// EdgeBranch holds the string denoting the branch edge name in mutations.
	EdgeBranch = "branch"
	// EdgeOwner holds the string denoting the owner edge name in mutations.
	EdgeOwner = "owner"
	// EdgePreference holds the string denoting the preference edge name in mutations.
	EdgePreference = "preference"
	// EdgeRoutingnumber holds the string denoting the routingnumber edge name in mutations.
	EdgeRoutingnumber = "routingnumber"
	// Table holds the table name of the account in the database.
	Table = "accounts"
	// BranchTable is the table the holds the branch relation/edge.
	BranchTable = "accounts"
	// BranchInverseTable is the table name for the Branch entity.
	// It exists in this package in order to avoid circular dependency with the "branch" package.
	BranchInverseTable = "branches"
	// BranchColumn is the table column denoting the branch relation/edge.
	BranchColumn = "account_branch"
	// OwnerTable is the table the holds the owner relation/edge. The primary key declared below.
	OwnerTable = "account_owner"
	// OwnerInverseTable is the table name for the Entity entity.
	// It exists in this package in order to avoid circular dependency with the "entity" package.
	OwnerInverseTable = "entities"
	// PreferenceTable is the table the holds the preference relation/edge.
	PreferenceTable = "preferences"
	// PreferenceInverseTable is the table name for the Preference entity.
	// It exists in this package in order to avoid circular dependency with the "preference" package.
	PreferenceInverseTable = "preferences"
	// PreferenceColumn is the table column denoting the preference relation/edge.
	PreferenceColumn = "account_preference"
	// RoutingnumberTable is the table the holds the routingnumber relation/edge.
	RoutingnumberTable = "routing_numbers"
	// RoutingnumberInverseTable is the table name for the RoutingNumber entity.
	// It exists in this package in order to avoid circular dependency with the "routingnumber" package.
	RoutingnumberInverseTable = "routing_numbers"
	// RoutingnumberColumn is the table column denoting the routingnumber relation/edge.
	RoutingnumberColumn = "account_routingnumber"
)

// Columns holds all SQL columns for account fields.
var Columns = []string{
	FieldID,
	FieldType,
	FieldNumber,
	FieldParentId,
	FieldName,
	FieldTitle,
	FieldDateCreated,
	FieldDateOpened,
	FieldDateLastUpdated,
	FieldDateClosed,
	FieldCurrencyCode,
	FieldStatus,
	FieldSource,
	FieldInterestReporting,
	FieldCurrentBalance,
	FieldAvailableBalance,
	FieldURL,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "accounts"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"account_branch",
}

var (
	// OwnerPrimaryKey and OwnerColumn2 are the table columns denoting the
	// primary key for the owner relation (M2M).
	OwnerPrimaryKey = []string{"account_id", "entity_id"}
)

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

var (
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
