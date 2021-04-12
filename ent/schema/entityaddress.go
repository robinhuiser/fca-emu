package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// EntityAddress holds the schema definition for the EntityAddress entity.
type EntityAddress struct {
	ent.Schema
}

// Fields of the EntityAddress.
func (EntityAddress) Fields() []ent.Field {
	return []ent.Field{
		field.String("country"),
		field.String("city"),
		field.String("postalCode"),
		field.String("state"),
		field.Enum("type").Values("BUSINESS", "PRIVATE", "MAILBOX"),
		field.String("line1"),
		field.String("line2").Optional(),
		field.String("line3").Optional(),
		field.Bool("primary"),
	}
}

// Edges of the EntityAddress.
func (EntityAddress) Edges() []ent.Edge {
	return nil
}
