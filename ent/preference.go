// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/robinhuiser/fca-emu/ent/preference"
)

// Preference is the model entity for the Preference schema.
type Preference struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Value holds the value of the "value" field.
	Value               string `json:"value,omitempty"`
	account_preferences *uuid.UUID
	entity_preferences  *uuid.UUID
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Preference) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case preference.FieldID:
			values[i] = &sql.NullInt64{}
		case preference.FieldName, preference.FieldValue:
			values[i] = &sql.NullString{}
		case preference.ForeignKeys[0]: // account_preferences
			values[i] = &uuid.UUID{}
		case preference.ForeignKeys[1]: // entity_preferences
			values[i] = &uuid.UUID{}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Preference", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Preference fields.
func (pr *Preference) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case preference.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pr.ID = int(value.Int64)
		case preference.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				pr.Name = value.String
			}
		case preference.FieldValue:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field value", values[i])
			} else if value.Valid {
				pr.Value = value.String
			}
		case preference.ForeignKeys[0]:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field account_preferences", values[i])
			} else if value != nil {
				pr.account_preferences = value
			}
		case preference.ForeignKeys[1]:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field entity_preferences", values[i])
			} else if value != nil {
				pr.entity_preferences = value
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Preference.
// Note that you need to call Preference.Unwrap() before calling this method if this Preference
// was returned from a transaction, and the transaction was committed or rolled back.
func (pr *Preference) Update() *PreferenceUpdateOne {
	return (&PreferenceClient{config: pr.config}).UpdateOne(pr)
}

// Unwrap unwraps the Preference entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pr *Preference) Unwrap() *Preference {
	tx, ok := pr.config.driver.(*txDriver)
	if !ok {
		panic("ent: Preference is not a transactional entity")
	}
	pr.config.driver = tx.drv
	return pr
}

// String implements the fmt.Stringer.
func (pr *Preference) String() string {
	var builder strings.Builder
	builder.WriteString("Preference(")
	builder.WriteString(fmt.Sprintf("id=%v", pr.ID))
	builder.WriteString(", name=")
	builder.WriteString(pr.Name)
	builder.WriteString(", value=")
	builder.WriteString(pr.Value)
	builder.WriteByte(')')
	return builder.String()
}

// Preferences is a parsable slice of Preference.
type Preferences []*Preference

func (pr Preferences) config(cfg config) {
	for _i := range pr {
		pr[_i].config = cfg
	}
}
