package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// EntityContactPoints holds the schema definition for the EntityContactPoints entity.
type EntityContactPoint struct {
	ent.Schema
}

// Fields of the EntityContactPoints.
func (EntityContactPoint) Fields() []ent.Field {
	return []ent.Field{
		field.Int("prefix").Optional(),
		field.String("name"),
		field.Enum("type").Values("SMS", "EMAIL", "PHONE", "WHATSAPP", "SKYPE"),
		field.Int("suffix").Optional(),
		field.String("value"),
	}
}

// Edges of the EntityContactPoints.
func (EntityContactPoint) Edges() []ent.Edge {
	return nil
}
