package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

// Comment holds the schema definition for the Comment entity.
type Comment struct {
	ent.Schema
}

// Fields of the Comment.
func (Comment) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.UUID("comment_id", uuid.UUID{}).
			Default(uuid.New),
		field.String("feed_id"),
		field.String("content"),
		field.String("owner_id"),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now),
		field.Time("deleted_at").
			Optional().
			Nillable(),
	}
}

// Edges of the Comment.
func (Comment) Edges() []ent.Edge {
	return nil
}
