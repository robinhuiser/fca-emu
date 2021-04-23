package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Transaction holds the schema definition for the Transaction entity.
type Transaction struct {
	ent.Schema
}

// Fields of the Transaction.
func (Transaction) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.Int("sequenceInDay").
			Optional(),
		field.Enum("status").
			Values("PENDING", "POSTED"),
		field.Float("executedAmount"),
		field.String("executedCurrencyCode").
			MinLen(3).
			MaxLen(3),
		field.Float("exchangeRate").
			Optional(),
		field.Float("originatingAmount").
			Optional(),
		field.String("originatingCurrencyCode").
			MinLen(3).
			MaxLen(3).
			Optional(),
		field.Enum("direction").
			Values("DEBIT", "CREDIT"),
		field.Float("runningBalance"),
		field.Time("createdDate"),
		field.Time("postedDate").
			Optional(),
		field.Time("executedDate").
			Optional(),
		field.Time("updatedDate").
			Optional(),
		field.String("description").
			Optional(),
		field.String("memo").
			Optional(),
		field.String("group").
			Optional(),
		field.String("type").
			Optional(),
		field.String("mainCategory").
			Optional(),
		field.String("subCategory").
			Optional(),
		field.String("checkNumber").
			Optional(),
		field.Float("latitude").
			Optional(),
		field.Float("longitude").
			Optional(),
		field.String("merchantCode").
			Optional(),
		field.Bool("reversal"),
		field.String("reversalFor").
			Optional(),
		field.Bool("reversed"),
		field.String("reversedBy").
			Optional(),
		field.String("url").
			Optional(),
	}
}

// Edges of the Transaction.
func (Transaction) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("images", BinaryItem.Type),
	}
}
