package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

type Feed struct {
	ent.Schema
}

// Fields of the Player.
func (Feed) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.UUID("feed_id", uuid.UUID{}).
			Default(uuid.New),
		field.String("audio_url"),
		field.String("caption").
			Optional().
			Nillable(),
		field.String("transcript").
			Optional().
			Nillable(),
		field.Enum("privacy").
			Values("PUBLIC", "PRIVATE").
			Default("PUBLIC"),
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
