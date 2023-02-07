// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// TempleColumns holds the columns for the "temple" table.
	TempleColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "created_at", Type: field.TypeTime, Nullable: true, SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true, SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"IN_PROGRESS", "COMPLETED"}, Default: "IN_PROGRESS"},
		{Name: "delete_at", Type: field.TypeEnum, Enums: []string{"ENABLE", "DISABLE"}, Default: "ENABLE"},
		{Name: "text", Type: field.TypeString, Size: 2147483647},
		{Name: "priority", Type: field.TypeInt, Default: 0},
		{Name: "temple_parent", Type: field.TypeUint64, Nullable: true},
		{Name: "user_todos", Type: field.TypeUint64, Nullable: true},
	}
	// TempleTable holds the schema information for the "temple" table.
	TempleTable = &schema.Table{
		Name:       "temple",
		Columns:    TempleColumns,
		PrimaryKey: []*schema.Column{TempleColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "temple_temple_parent",
				Columns:    []*schema.Column{TempleColumns[7]},
				RefColumns: []*schema.Column{TempleColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "temple_user_todos",
				Columns:    []*schema.Column{TempleColumns[8]},
				RefColumns: []*schema.Column{UserColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UserColumns holds the columns for the "user" table.
	UserColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "created_at", Type: field.TypeTime, Nullable: true, SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true, SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"IN_PROGRESS", "COMPLETED"}, Default: "IN_PROGRESS"},
		{Name: "delete_at", Type: field.TypeEnum, Enums: []string{"ENABLE", "DISABLE"}, Default: "ENABLE"},
	}
	// UserTable holds the schema information for the "user" table.
	UserTable = &schema.Table{
		Name:       "user",
		Columns:    UserColumns,
		PrimaryKey: []*schema.Column{UserColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		TempleTable,
		UserTable,
	}
)

func init() {
	TempleTable.ForeignKeys[0].RefTable = TempleTable
	TempleTable.ForeignKeys[1].RefTable = UserTable
	TempleTable.Annotation = &entsql.Annotation{
		Table: "temple",
	}
	UserTable.Annotation = &entsql.Annotation{
		Table: "user",
	}
}
