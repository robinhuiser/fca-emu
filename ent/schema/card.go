package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Card holds the schema definition for the Card entity.
type Card struct {
	ent.Schema
}

// Fields of the Card.
func (Card) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("type").
			Values("CREDIT", "DEBIT", "LOYALTY"),
		field.String("number"),
		field.Time("startDate"),
		field.Time("expiryDate"),
		field.String("holderName"),
		field.Enum("status").
			Values("LOCKED", "OPERATIONAL"),
		field.String("url"),
	}
}

// Edges of the Card.
func (Card) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("network", CardNetwork.Type).Unique(),
	}
}
