package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// RoutingNumber holds the schema definition for the RoutingNumber entity.
type RoutingNumber struct {
	ent.Schema
}

// Fields of the RoutingNumber.
func (RoutingNumber) Fields() []ent.Field {
	return []ent.Field{
		field.String("number"),
		field.Enum("type").Values("WIRE", "ABA"),
	}
}

// Edges of the RoutingNumber.
func (RoutingNumber) Edges() []ent.Edge {
	return nil
}
