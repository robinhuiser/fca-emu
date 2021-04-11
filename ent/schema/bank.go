package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Bank holds the schema definition for the Bank entity.
type Bank struct {
	ent.Schema
}

// Fields of the Bank.
func (Bank) Fields() []ent.Field {
	return []ent.Field{
		field.String("bankCode"),
		field.String("bankName"),
		field.String("url"),
		field.String("swift"),
	}
}

// Edges of the Bank.
func (Bank) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("branches", Branch.Type),
	}
}
