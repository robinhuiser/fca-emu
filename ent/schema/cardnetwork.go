package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// CardNetwork holds the schema definition for the CardNetwork entity.
type CardNetwork struct {
	ent.Schema
}

// Fields of the CardNetwork.
func (CardNetwork) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("code"),
	}
}

// Edges of the CardNetwork.
func (CardNetwork) Edges() []ent.Edge {
	return nil
}
