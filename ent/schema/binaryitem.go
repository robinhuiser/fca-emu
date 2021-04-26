package schema

import (
	"errors"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// BinaryItem holds the schema definition for the BinaryItem entity.
type BinaryItem struct {
	ent.Schema
}

// Fields of the BinaryItem.
func (BinaryItem) Fields() []ent.Field {
	return []ent.Field{
		field.String("format").
			Validate(func(s string) error {
				values := []string{
					"application/pdf",
					"image/jpeg",
					"image/png",
					"image/svg",
					"image/gif",
					"text/html",
					"text/plain",
				}
				valid := false
				for _, value := range values {
					if s == value {
						valid = true
					}
				}
				if !valid {
					return errors.New("not a valid binary type")
				}
				return nil
			}),
		field.Int("length").
			Optional(),
		field.Bytes("content"),
		field.String("url").
			Optional(),
	}
}

// Edges of the BinaryItem.
func (BinaryItem) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("transaction", Transaction.Type).
			Ref("images").
			Unique(),
	}
}
