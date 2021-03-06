// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/robinhuiser/fca-emu/ent/bank"
	"github.com/robinhuiser/fca-emu/ent/branch"
)

// Branch is the model entity for the Branch schema.
type Branch struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// BranchCode holds the value of the "branchCode" field.
	BranchCode string `json:"branchCode,omitempty"`
	// StreetNumber holds the value of the "streetNumber" field.
	StreetNumber string `json:"streetNumber,omitempty"`
	// StreetName holds the value of the "streetName" field.
	StreetName string `json:"streetName,omitempty"`
	// City holds the value of the "city" field.
	City string `json:"city,omitempty"`
	// State holds the value of the "state" field.
	State string `json:"state,omitempty"`
	// Zip holds the value of the "zip" field.
	Zip string `json:"zip,omitempty"`
	// Latitude holds the value of the "latitude" field.
	Latitude float64 `json:"latitude,omitempty"`
	// Longitude holds the value of the "longitude" field.
	Longitude float64 `json:"longitude,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the BranchQuery when eager-loading is set.
	Edges         BranchEdges `json:"edges"`
	bank_branches *int
}

// BranchEdges holds the relations/edges for other nodes in the graph.
type BranchEdges struct {
	// Owner holds the value of the owner edge.
	Owner *Bank `json:"owner,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BranchEdges) OwnerOrErr() (*Bank, error) {
	if e.loadedTypes[0] {
		if e.Owner == nil {
			// The edge owner was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: bank.Label}
		}
		return e.Owner, nil
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Branch) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case branch.FieldLatitude, branch.FieldLongitude:
			values[i] = &sql.NullFloat64{}
		case branch.FieldID:
			values[i] = &sql.NullInt64{}
		case branch.FieldBranchCode, branch.FieldStreetNumber, branch.FieldStreetName, branch.FieldCity, branch.FieldState, branch.FieldZip:
			values[i] = &sql.NullString{}
		case branch.ForeignKeys[0]: // bank_branches
			values[i] = &sql.NullInt64{}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Branch", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Branch fields.
func (b *Branch) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case branch.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			b.ID = int(value.Int64)
		case branch.FieldBranchCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field branchCode", values[i])
			} else if value.Valid {
				b.BranchCode = value.String
			}
		case branch.FieldStreetNumber:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field streetNumber", values[i])
			} else if value.Valid {
				b.StreetNumber = value.String
			}
		case branch.FieldStreetName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field streetName", values[i])
			} else if value.Valid {
				b.StreetName = value.String
			}
		case branch.FieldCity:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field city", values[i])
			} else if value.Valid {
				b.City = value.String
			}
		case branch.FieldState:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field state", values[i])
			} else if value.Valid {
				b.State = value.String
			}
		case branch.FieldZip:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field zip", values[i])
			} else if value.Valid {
				b.Zip = value.String
			}
		case branch.FieldLatitude:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field latitude", values[i])
			} else if value.Valid {
				b.Latitude = value.Float64
			}
		case branch.FieldLongitude:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field longitude", values[i])
			} else if value.Valid {
				b.Longitude = value.Float64
			}
		case branch.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field bank_branches", value)
			} else if value.Valid {
				b.bank_branches = new(int)
				*b.bank_branches = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryOwner queries the "owner" edge of the Branch entity.
func (b *Branch) QueryOwner() *BankQuery {
	return (&BranchClient{config: b.config}).QueryOwner(b)
}

// Update returns a builder for updating this Branch.
// Note that you need to call Branch.Unwrap() before calling this method if this Branch
// was returned from a transaction, and the transaction was committed or rolled back.
func (b *Branch) Update() *BranchUpdateOne {
	return (&BranchClient{config: b.config}).UpdateOne(b)
}

// Unwrap unwraps the Branch entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (b *Branch) Unwrap() *Branch {
	tx, ok := b.config.driver.(*txDriver)
	if !ok {
		panic("ent: Branch is not a transactional entity")
	}
	b.config.driver = tx.drv
	return b
}

// String implements the fmt.Stringer.
func (b *Branch) String() string {
	var builder strings.Builder
	builder.WriteString("Branch(")
	builder.WriteString(fmt.Sprintf("id=%v", b.ID))
	builder.WriteString(", branchCode=")
	builder.WriteString(b.BranchCode)
	builder.WriteString(", streetNumber=")
	builder.WriteString(b.StreetNumber)
	builder.WriteString(", streetName=")
	builder.WriteString(b.StreetName)
	builder.WriteString(", city=")
	builder.WriteString(b.City)
	builder.WriteString(", state=")
	builder.WriteString(b.State)
	builder.WriteString(", zip=")
	builder.WriteString(b.Zip)
	builder.WriteString(", latitude=")
	builder.WriteString(fmt.Sprintf("%v", b.Latitude))
	builder.WriteString(", longitude=")
	builder.WriteString(fmt.Sprintf("%v", b.Longitude))
	builder.WriteByte(')')
	return builder.String()
}

// Branches is a parsable slice of Branch.
type Branches []*Branch

func (b Branches) config(cfg config) {
	for _i := range b {
		b[_i].config = cfg
	}
}
