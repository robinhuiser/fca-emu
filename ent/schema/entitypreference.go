package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// EntityPreferences holds the schema definition for the EntityPreferences entity.
type EntityPreference struct {
	ent.Schema
}

// Fields of the EntityPreferences.
func (EntityPreference) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("value"),
	}
}

// Edges of the EntityPreferences.
func (EntityPreference) Edges() []ent.Edge {
	return nil
}
