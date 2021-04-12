package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// TaxInformation holds the schema definition for the TaxInformation entity.
type EntityTaxInformation struct {
	ent.Schema
}

// Fields of the TaxInformation.
func (EntityTaxInformation) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("type").Values("SSN"),
		field.String("taxId"),
	}
}

// Edges of the TaxInformation.
func (EntityTaxInformation) Edges() []ent.Edge {
	return nil
}
