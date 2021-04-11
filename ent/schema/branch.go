package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Branch holds the schema definition for the Branch entity.
type Branch struct {
	ent.Schema
}

// Fields of the Branch.
func (Branch) Fields() []ent.Field {
	return []ent.Field{
		field.String("branchCode"),
		field.String("streetNumber"),
		field.String("streetName"),
		field.String("city"),
		field.String("state").
			MaxLen(2),
		field.String("zip"),
		field.Float("latitude"),
		field.Float("longitude"),
	}
}

// Edges of the Branch.
func (Branch) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("branch_owner", Bank.Type).
			Ref("branches").
			Unique(),
	}
}
