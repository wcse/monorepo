// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"github.com/facebook/ent/dialect/sql/schema"
	"github.com/facebook/ent/schema/field"
)

var (
	// FeedsColumns holds the columns for the "feeds" table.
	FeedsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "feed_id", Type: field.TypeUUID},
		{Name: "audio_url", Type: field.TypeString},
		{Name: "caption", Type: field.TypeString, Nullable: true},
		{Name: "transcript", Type: field.TypeString, Nullable: true},
		{Name: "privacy", Type: field.TypeEnum, Enums: []string{"PRIVATE", "PUBLIC"}, Default: "PUBLIC"},
		{Name: "owner_id", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
	}
	// FeedsTable holds the schema information for the "feeds" table.
	FeedsTable = &schema.Table{
		Name:        "feeds",
		Columns:     FeedsColumns,
		PrimaryKey:  []*schema.Column{FeedsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		FeedsTable,
	}
)

func init() {
}
