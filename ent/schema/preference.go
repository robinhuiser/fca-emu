package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Preferences holds the schema definition for the Preferences entity.
type Preference struct {
	ent.Schema
}

// Fields of the Preferences.
func (Preference) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("value"),
	}
}

// Edges of the EntityPreferences.
func (Preference) Edges() []ent.Edge {
	return nil
}
