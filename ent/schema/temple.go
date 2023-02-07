package schema

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Temple holds the schema definition for the Temple entity.
type Temple struct {
	ent.Schema
}

// Fields of the Temple.
func (Temple) Fields() []ent.Field {
	return []ent.Field{
		field.Text("text").
			NotEmpty().
			Annotations(
				entgql.OrderField("TEXT"),
			),
		field.Time("created_at").
			Default(time.Now).
			Immutable().
			Annotations(
				entgql.OrderField("CREATED_AT"),
			),
		field.Enum("status").
			NamedValues(
				"InProgress", "IN_PROGRESS",
				"Completed", "COMPLETED",
			).
			Default("IN_PROGRESS").
			Annotations(
				entgql.OrderField("STATUS"),
			),
		field.Int("priority").
			Default(0).
			Annotations(
				entgql.OrderField("PRIORITY"),
			),
	}
}

// Edges of the Temple.
func (Temple) Edges() []ent.Edge {
	return nil
}
