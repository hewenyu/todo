package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/hewenyu/todo/ent/mixins"
)

// Temple holds the schema definition for the Temple entity.
// 这是一个模板表
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
		field.Int("priority").
			Default(0).
			Annotations(
				entgql.OrderField("PRIORITY"),
			),
	}
}

// Edges of the Temple.
func (Temple) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("parent", Temple.Type).
			Unique().
			From("children"),
		edge.From("owner", User.Type).
			Ref("todos").
			Unique(),
	}
}

// 加载公共参数
func (Temple) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
	}
}

// 定义表相关信息
func (Temple) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "temple"},
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate()),
	}
}
