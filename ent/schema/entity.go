package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Entity holds the schema definition for the Entity entity.
type Entity struct {
	ent.Schema
}

// Fields of the Entity.
func (Entity) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.Time("dateCreated"),
		field.String("firstname").Optional(),
		field.String("lastname").Optional(),
		field.String("fullname").Optional(),
		field.Time("dateOfBirth"),
		field.Bool("active").Default(true),
		field.Enum("type").Values("PERSON", "ORGANIZATION", "CORPORATE"),
		field.Time("lastLoginDate"),
		field.String("username"),
		field.String("token"),
		field.String("url"),
	}
}

// Edges of the Entity.
func (Entity) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("taxSpecifications", EntityTaxInformation.Type),
		edge.To("addresses", EntityAddress.Type),
		edge.To("preferences", Preference.Type),
		edge.To("contactPoints", EntityContactPoint.Type),
		edge.From("owns_account", Account.Type).Ref("owners"),
	}
}
