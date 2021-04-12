package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Product holds the schema definition for the Product entity.
type Product struct {
	ent.Schema
}

// Fields of the Product.
func (Product) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Enum("type").Values("INVESTMENT", "LOAN", "DEPOSIT", "CARD"),
		field.String("typeName"),
		field.String("subType"),
		field.String("subTypeName"),
		field.String("url"),
	}
}

// Edges of the Product.
func (Product) Edges() []ent.Edge {
	return nil
}
