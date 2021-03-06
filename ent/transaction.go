// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/robinhuiser/fca-emu/ent/account"
	"github.com/robinhuiser/fca-emu/ent/transaction"
)

// Transaction is the model entity for the Transaction schema.
type Transaction struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// SequenceInDay holds the value of the "sequenceInDay" field.
	SequenceInDay int `json:"sequenceInDay,omitempty"`
	// Status holds the value of the "status" field.
	Status transaction.Status `json:"status,omitempty"`
	// ExecutedAmount holds the value of the "executedAmount" field.
	ExecutedAmount float64 `json:"executedAmount,omitempty"`
	// ExecutedCurrencyCode holds the value of the "executedCurrencyCode" field.
	ExecutedCurrencyCode string `json:"executedCurrencyCode,omitempty"`
	// ExchangeRate holds the value of the "exchangeRate" field.
	ExchangeRate float64 `json:"exchangeRate,omitempty"`
	// OriginatingAmount holds the value of the "originatingAmount" field.
	OriginatingAmount float64 `json:"originatingAmount,omitempty"`
	// OriginatingCurrencyCode holds the value of the "originatingCurrencyCode" field.
	OriginatingCurrencyCode string `json:"originatingCurrencyCode,omitempty"`
	// Direction holds the value of the "direction" field.
	Direction transaction.Direction `json:"direction,omitempty"`
	// RunningBalance holds the value of the "runningBalance" field.
	RunningBalance float64 `json:"runningBalance,omitempty"`
	// CreatedDate holds the value of the "createdDate" field.
	CreatedDate time.Time `json:"createdDate,omitempty"`
	// PostedDate holds the value of the "postedDate" field.
	PostedDate time.Time `json:"postedDate,omitempty"`
	// ExecutedDate holds the value of the "executedDate" field.
	ExecutedDate time.Time `json:"executedDate,omitempty"`
	// UpdatedDate holds the value of the "updatedDate" field.
	UpdatedDate time.Time `json:"updatedDate,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Memo holds the value of the "memo" field.
	Memo string `json:"memo,omitempty"`
	// Group holds the value of the "group" field.
	Group string `json:"group,omitempty"`
	// Type holds the value of the "type" field.
	Type string `json:"type,omitempty"`
	// MainCategory holds the value of the "mainCategory" field.
	MainCategory string `json:"mainCategory,omitempty"`
	// SubCategory holds the value of the "subCategory" field.
	SubCategory string `json:"subCategory,omitempty"`
	// CheckNumber holds the value of the "checkNumber" field.
	CheckNumber string `json:"checkNumber,omitempty"`
	// Latitude holds the value of the "latitude" field.
	Latitude float64 `json:"latitude,omitempty"`
	// Longitude holds the value of the "longitude" field.
	Longitude float64 `json:"longitude,omitempty"`
	// MerchantCode holds the value of the "merchantCode" field.
	MerchantCode string `json:"merchantCode,omitempty"`
	// Reversal holds the value of the "reversal" field.
	Reversal bool `json:"reversal,omitempty"`
	// ReversalFor holds the value of the "reversalFor" field.
	ReversalFor string `json:"reversalFor,omitempty"`
	// Reversed holds the value of the "reversed" field.
	Reversed bool `json:"reversed,omitempty"`
	// ReversedBy holds the value of the "reversedBy" field.
	ReversedBy string `json:"reversedBy,omitempty"`
	// URL holds the value of the "url" field.
	URL string `json:"url,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TransactionQuery when eager-loading is set.
	Edges                TransactionEdges `json:"edges"`
	account_transactions *uuid.UUID
}

// TransactionEdges holds the relations/edges for other nodes in the graph.
type TransactionEdges struct {
	// Images holds the value of the images edge.
	Images []*BinaryItem `json:"images,omitempty"`
	// Account holds the value of the account edge.
	Account *Account `json:"account,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// ImagesOrErr returns the Images value or an error if the edge
// was not loaded in eager-loading.
func (e TransactionEdges) ImagesOrErr() ([]*BinaryItem, error) {
	if e.loadedTypes[0] {
		return e.Images, nil
	}
	return nil, &NotLoadedError{edge: "images"}
}

// AccountOrErr returns the Account value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TransactionEdges) AccountOrErr() (*Account, error) {
	if e.loadedTypes[1] {
		if e.Account == nil {
			// The edge account was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: account.Label}
		}
		return e.Account, nil
	}
	return nil, &NotLoadedError{edge: "account"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Transaction) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case transaction.FieldReversal, transaction.FieldReversed:
			values[i] = &sql.NullBool{}
		case transaction.FieldExecutedAmount, transaction.FieldExchangeRate, transaction.FieldOriginatingAmount, transaction.FieldRunningBalance, transaction.FieldLatitude, transaction.FieldLongitude:
			values[i] = &sql.NullFloat64{}
		case transaction.FieldSequenceInDay:
			values[i] = &sql.NullInt64{}
		case transaction.FieldStatus, transaction.FieldExecutedCurrencyCode, transaction.FieldOriginatingCurrencyCode, transaction.FieldDirection, transaction.FieldDescription, transaction.FieldMemo, transaction.FieldGroup, transaction.FieldType, transaction.FieldMainCategory, transaction.FieldSubCategory, transaction.FieldCheckNumber, transaction.FieldMerchantCode, transaction.FieldReversalFor, transaction.FieldReversedBy, transaction.FieldURL:
			values[i] = &sql.NullString{}
		case transaction.FieldCreatedDate, transaction.FieldPostedDate, transaction.FieldExecutedDate, transaction.FieldUpdatedDate:
			values[i] = &sql.NullTime{}
		case transaction.FieldID:
			values[i] = &uuid.UUID{}
		case transaction.ForeignKeys[0]: // account_transactions
			values[i] = &uuid.UUID{}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Transaction", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Transaction fields.
func (t *Transaction) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case transaction.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				t.ID = *value
			}
		case transaction.FieldSequenceInDay:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field sequenceInDay", values[i])
			} else if value.Valid {
				t.SequenceInDay = int(value.Int64)
			}
		case transaction.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				t.Status = transaction.Status(value.String)
			}
		case transaction.FieldExecutedAmount:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field executedAmount", values[i])
			} else if value.Valid {
				t.ExecutedAmount = value.Float64
			}
		case transaction.FieldExecutedCurrencyCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field executedCurrencyCode", values[i])
			} else if value.Valid {
				t.ExecutedCurrencyCode = value.String
			}
		case transaction.FieldExchangeRate:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field exchangeRate", values[i])
			} else if value.Valid {
				t.ExchangeRate = value.Float64
			}
		case transaction.FieldOriginatingAmount:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field originatingAmount", values[i])
			} else if value.Valid {
				t.OriginatingAmount = value.Float64
			}
		case transaction.FieldOriginatingCurrencyCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field originatingCurrencyCode", values[i])
			} else if value.Valid {
				t.OriginatingCurrencyCode = value.String
			}
		case transaction.FieldDirection:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field direction", values[i])
			} else if value.Valid {
				t.Direction = transaction.Direction(value.String)
			}
		case transaction.FieldRunningBalance:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field runningBalance", values[i])
			} else if value.Valid {
				t.RunningBalance = value.Float64
			}
		case transaction.FieldCreatedDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field createdDate", values[i])
			} else if value.Valid {
				t.CreatedDate = value.Time
			}
		case transaction.FieldPostedDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field postedDate", values[i])
			} else if value.Valid {
				t.PostedDate = value.Time
			}
		case transaction.FieldExecutedDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field executedDate", values[i])
			} else if value.Valid {
				t.ExecutedDate = value.Time
			}
		case transaction.FieldUpdatedDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updatedDate", values[i])
			} else if value.Valid {
				t.UpdatedDate = value.Time
			}
		case transaction.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				t.Description = value.String
			}
		case transaction.FieldMemo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field memo", values[i])
			} else if value.Valid {
				t.Memo = value.String
			}
		case transaction.FieldGroup:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field group", values[i])
			} else if value.Valid {
				t.Group = value.String
			}
		case transaction.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				t.Type = value.String
			}
		case transaction.FieldMainCategory:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field mainCategory", values[i])
			} else if value.Valid {
				t.MainCategory = value.String
			}
		case transaction.FieldSubCategory:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field subCategory", values[i])
			} else if value.Valid {
				t.SubCategory = value.String
			}
		case transaction.FieldCheckNumber:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field checkNumber", values[i])
			} else if value.Valid {
				t.CheckNumber = value.String
			}
		case transaction.FieldLatitude:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field latitude", values[i])
			} else if value.Valid {
				t.Latitude = value.Float64
			}
		case transaction.FieldLongitude:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field longitude", values[i])
			} else if value.Valid {
				t.Longitude = value.Float64
			}
		case transaction.FieldMerchantCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field merchantCode", values[i])
			} else if value.Valid {
				t.MerchantCode = value.String
			}
		case transaction.FieldReversal:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field reversal", values[i])
			} else if value.Valid {
				t.Reversal = value.Bool
			}
		case transaction.FieldReversalFor:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field reversalFor", values[i])
			} else if value.Valid {
				t.ReversalFor = value.String
			}
		case transaction.FieldReversed:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field reversed", values[i])
			} else if value.Valid {
				t.Reversed = value.Bool
			}
		case transaction.FieldReversedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field reversedBy", values[i])
			} else if value.Valid {
				t.ReversedBy = value.String
			}
		case transaction.FieldURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field url", values[i])
			} else if value.Valid {
				t.URL = value.String
			}
		case transaction.ForeignKeys[0]:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field account_transactions", values[i])
			} else if value != nil {
				t.account_transactions = value
			}
		}
	}
	return nil
}

// QueryImages queries the "images" edge of the Transaction entity.
func (t *Transaction) QueryImages() *BinaryItemQuery {
	return (&TransactionClient{config: t.config}).QueryImages(t)
}

// QueryAccount queries the "account" edge of the Transaction entity.
func (t *Transaction) QueryAccount() *AccountQuery {
	return (&TransactionClient{config: t.config}).QueryAccount(t)
}

// Update returns a builder for updating this Transaction.
// Note that you need to call Transaction.Unwrap() before calling this method if this Transaction
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Transaction) Update() *TransactionUpdateOne {
	return (&TransactionClient{config: t.config}).UpdateOne(t)
}

// Unwrap unwraps the Transaction entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Transaction) Unwrap() *Transaction {
	tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Transaction is not a transactional entity")
	}
	t.config.driver = tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Transaction) String() string {
	var builder strings.Builder
	builder.WriteString("Transaction(")
	builder.WriteString(fmt.Sprintf("id=%v", t.ID))
	builder.WriteString(", sequenceInDay=")
	builder.WriteString(fmt.Sprintf("%v", t.SequenceInDay))
	builder.WriteString(", status=")
	builder.WriteString(fmt.Sprintf("%v", t.Status))
	builder.WriteString(", executedAmount=")
	builder.WriteString(fmt.Sprintf("%v", t.ExecutedAmount))
	builder.WriteString(", executedCurrencyCode=")
	builder.WriteString(t.ExecutedCurrencyCode)
	builder.WriteString(", exchangeRate=")
	builder.WriteString(fmt.Sprintf("%v", t.ExchangeRate))
	builder.WriteString(", originatingAmount=")
	builder.WriteString(fmt.Sprintf("%v", t.OriginatingAmount))
	builder.WriteString(", originatingCurrencyCode=")
	builder.WriteString(t.OriginatingCurrencyCode)
	builder.WriteString(", direction=")
	builder.WriteString(fmt.Sprintf("%v", t.Direction))
	builder.WriteString(", runningBalance=")
	builder.WriteString(fmt.Sprintf("%v", t.RunningBalance))
	builder.WriteString(", createdDate=")
	builder.WriteString(t.CreatedDate.Format(time.ANSIC))
	builder.WriteString(", postedDate=")
	builder.WriteString(t.PostedDate.Format(time.ANSIC))
	builder.WriteString(", executedDate=")
	builder.WriteString(t.ExecutedDate.Format(time.ANSIC))
	builder.WriteString(", updatedDate=")
	builder.WriteString(t.UpdatedDate.Format(time.ANSIC))
	builder.WriteString(", description=")
	builder.WriteString(t.Description)
	builder.WriteString(", memo=")
	builder.WriteString(t.Memo)
	builder.WriteString(", group=")
	builder.WriteString(t.Group)
	builder.WriteString(", type=")
	builder.WriteString(t.Type)
	builder.WriteString(", mainCategory=")
	builder.WriteString(t.MainCategory)
	builder.WriteString(", subCategory=")
	builder.WriteString(t.SubCategory)
	builder.WriteString(", checkNumber=")
	builder.WriteString(t.CheckNumber)
	builder.WriteString(", latitude=")
	builder.WriteString(fmt.Sprintf("%v", t.Latitude))
	builder.WriteString(", longitude=")
	builder.WriteString(fmt.Sprintf("%v", t.Longitude))
	builder.WriteString(", merchantCode=")
	builder.WriteString(t.MerchantCode)
	builder.WriteString(", reversal=")
	builder.WriteString(fmt.Sprintf("%v", t.Reversal))
	builder.WriteString(", reversalFor=")
	builder.WriteString(t.ReversalFor)
	builder.WriteString(", reversed=")
	builder.WriteString(fmt.Sprintf("%v", t.Reversed))
	builder.WriteString(", reversedBy=")
	builder.WriteString(t.ReversedBy)
	builder.WriteString(", url=")
	builder.WriteString(t.URL)
	builder.WriteByte(')')
	return builder.String()
}

// Transactions is a parsable slice of Transaction.
type Transactions []*Transaction

func (t Transactions) config(cfg config) {
	for _i := range t {
		t[_i].config = cfg
	}
}
