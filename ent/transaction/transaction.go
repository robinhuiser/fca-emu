// Code generated by entc, DO NOT EDIT.

package transaction

import (
	"fmt"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the transaction type in the database.
	Label = "transaction"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldSequenceInDay holds the string denoting the sequenceinday field in the database.
	FieldSequenceInDay = "sequence_in_day"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldExecutedAmount holds the string denoting the executedamount field in the database.
	FieldExecutedAmount = "executed_amount"
	// FieldExecutedCurrencyCode holds the string denoting the executedcurrencycode field in the database.
	FieldExecutedCurrencyCode = "executed_currency_code"
	// FieldExchangeRate holds the string denoting the exchangerate field in the database.
	FieldExchangeRate = "exchange_rate"
	// FieldOriginatingAmount holds the string denoting the originatingamount field in the database.
	FieldOriginatingAmount = "originating_amount"
	// FieldOriginatingCurrencyCode holds the string denoting the originatingcurrencycode field in the database.
	FieldOriginatingCurrencyCode = "originating_currency_code"
	// FieldDirection holds the string denoting the direction field in the database.
	FieldDirection = "direction"
	// FieldRunningBalance holds the string denoting the runningbalance field in the database.
	FieldRunningBalance = "running_balance"
	// FieldCreatedDate holds the string denoting the createddate field in the database.
	FieldCreatedDate = "created_date"
	// FieldPostedDate holds the string denoting the posteddate field in the database.
	FieldPostedDate = "posted_date"
	// FieldExecutedDate holds the string denoting the executeddate field in the database.
	FieldExecutedDate = "executed_date"
	// FieldUpdatedDate holds the string denoting the updateddate field in the database.
	FieldUpdatedDate = "updated_date"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldMemo holds the string denoting the memo field in the database.
	FieldMemo = "memo"
	// FieldGroup holds the string denoting the group field in the database.
	FieldGroup = "group"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldMainCategory holds the string denoting the maincategory field in the database.
	FieldMainCategory = "main_category"
	// FieldSubCategory holds the string denoting the subcategory field in the database.
	FieldSubCategory = "sub_category"
	// FieldCheckNumber holds the string denoting the checknumber field in the database.
	FieldCheckNumber = "check_number"
	// FieldLatitude holds the string denoting the latitude field in the database.
	FieldLatitude = "latitude"
	// FieldLongitude holds the string denoting the longitude field in the database.
	FieldLongitude = "longitude"
	// FieldMerchantCode holds the string denoting the merchantcode field in the database.
	FieldMerchantCode = "merchant_code"
	// FieldReversal holds the string denoting the reversal field in the database.
	FieldReversal = "reversal"
	// FieldReversalFor holds the string denoting the reversalfor field in the database.
	FieldReversalFor = "reversal_for"
	// FieldReversed holds the string denoting the reversed field in the database.
	FieldReversed = "reversed"
	// FieldReversedBy holds the string denoting the reversedby field in the database.
	FieldReversedBy = "reversed_by"
	// FieldURL holds the string denoting the url field in the database.
	FieldURL = "url"
	// EdgeImages holds the string denoting the images edge name in mutations.
	EdgeImages = "images"
	// EdgeAccount holds the string denoting the account edge name in mutations.
	EdgeAccount = "account"
	// Table holds the table name of the transaction in the database.
	Table = "transactions"
	// ImagesTable is the table the holds the images relation/edge.
	ImagesTable = "binary_items"
	// ImagesInverseTable is the table name for the BinaryItem entity.
	// It exists in this package in order to avoid circular dependency with the "binaryitem" package.
	ImagesInverseTable = "binary_items"
	// ImagesColumn is the table column denoting the images relation/edge.
	ImagesColumn = "transaction_images"
	// AccountTable is the table the holds the account relation/edge.
	AccountTable = "transactions"
	// AccountInverseTable is the table name for the Account entity.
	// It exists in this package in order to avoid circular dependency with the "account" package.
	AccountInverseTable = "accounts"
	// AccountColumn is the table column denoting the account relation/edge.
	AccountColumn = "account_transactions"
)

// Columns holds all SQL columns for transaction fields.
var Columns = []string{
	FieldID,
	FieldSequenceInDay,
	FieldStatus,
	FieldExecutedAmount,
	FieldExecutedCurrencyCode,
	FieldExchangeRate,
	FieldOriginatingAmount,
	FieldOriginatingCurrencyCode,
	FieldDirection,
	FieldRunningBalance,
	FieldCreatedDate,
	FieldPostedDate,
	FieldExecutedDate,
	FieldUpdatedDate,
	FieldDescription,
	FieldMemo,
	FieldGroup,
	FieldType,
	FieldMainCategory,
	FieldSubCategory,
	FieldCheckNumber,
	FieldLatitude,
	FieldLongitude,
	FieldMerchantCode,
	FieldReversal,
	FieldReversalFor,
	FieldReversed,
	FieldReversedBy,
	FieldURL,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "transactions"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"account_transactions",
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

var (
	// ExecutedCurrencyCodeValidator is a validator for the "executedCurrencyCode" field. It is called by the builders before save.
	ExecutedCurrencyCodeValidator func(string) error
	// OriginatingCurrencyCodeValidator is a validator for the "originatingCurrencyCode" field. It is called by the builders before save.
	OriginatingCurrencyCodeValidator func(string) error
	// DefaultReversal holds the default value on creation for the "reversal" field.
	DefaultReversal bool
	// DefaultReversed holds the default value on creation for the "reversed" field.
	DefaultReversed bool
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// Status defines the type for the "status" enum field.
type Status string

// Status values.
const (
	StatusPENDING Status = "PENDING"
	StatusPOSTED  Status = "POSTED"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusPENDING, StatusPOSTED:
		return nil
	default:
		return fmt.Errorf("transaction: invalid enum value for status field: %q", s)
	}
}

// Direction defines the type for the "direction" enum field.
type Direction string

// Direction values.
const (
	DirectionDEBIT  Direction = "DEBIT"
	DirectionCREDIT Direction = "CREDIT"
)

func (d Direction) String() string {
	return string(d)
}

// DirectionValidator is a validator for the "direction" field enum values. It is called by the builders before save.
func DirectionValidator(d Direction) error {
	switch d {
	case DirectionDEBIT, DirectionCREDIT:
		return nil
	default:
		return fmt.Errorf("transaction: invalid enum value for direction field: %q", d)
	}
}
