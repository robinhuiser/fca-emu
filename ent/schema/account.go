package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Account holds the schema definition for the Account entity.
type Account struct {
	ent.Schema
}

// Fields of the Account.
func (Account) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.String("type"),
		field.String("number"),
		field.UUID("parentId", uuid.UUID{}),
		field.String("name"),
		field.String("title"),
		field.String("iban"),
		field.Time("dateCreated"),
		field.Time("dateOpened"),
		field.Time("dateLastUpdated"),
		field.Time("dateClosed").Optional(),
		field.String("currencyCode"),
		field.String("status"),
		field.String("source"),
		field.Bool("interestReporting"),
	}
}

// Edges of the Account.
func (Account) Edges() []ent.Edge {
	return nil
}
