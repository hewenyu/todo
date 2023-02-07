package mixins

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// BaseMixin implements the ent.Mixin for sharing
// 基础公共参数
type BaseMixin struct {
	// We embed the `mixin.Schema` to avoid
	// implementing the rest of the methods.
	mixin.Schema
}

func (BaseMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Comment("id"),
		field.Time("created_at").Optional().
			SchemaType(map[string]string{dialect.MySQL: "datetime"}).
			Default(time.Now).Immutable().
			Annotations(entgql.OrderField("CREATED_AT")),
		field.Time("updated_at").Optional().
			SchemaType(map[string]string{dialect.MySQL: "datetime"}).
			Default(time.Now).UpdateDefault(time.Now).
			Comment("更新时间时间").
			Annotations(entgql.OrderField("UPDATE_AT")),
		field.Enum("status").
			NamedValues(
				"InProgress", "IN_PROGRESS",
				"Completed", "COMPLETED",
			).
			Default("IN_PROGRESS").
			Comment("状态").
			Annotations(
				entgql.OrderField("STATUS"),
			),
		field.Enum("delete_at").
			NamedValues(
				"Enable", "ENABLE",
				"Disable", "DISABLE",
			).
			Default("ENABLE").
			Comment("状态 正常 删除").
			Annotations(entgql.OrderField("DELETE_AT")),
	}
}
