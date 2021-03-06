// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"github.com/facebook/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/wcse/monorepo/backend/feed/ent/feed"
)

// Feed is the model entity for the Feed schema.
type Feed struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// FeedID holds the value of the "feed_id" field.
	FeedID uuid.UUID `json:"feed_id,omitempty"`
	// AudioURL holds the value of the "audio_url" field.
	AudioURL string `json:"audio_url,omitempty"`
	// Caption holds the value of the "caption" field.
	Caption *string `json:"caption,omitempty"`
	// Transcript holds the value of the "transcript" field.
	Transcript *string `json:"transcript,omitempty"`
	// Privacy holds the value of the "privacy" field.
	Privacy feed.Privacy `json:"privacy,omitempty"`
	// OwnerID holds the value of the "owner_id" field.
	OwnerID string `json:"owner_id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Feed) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&uuid.UUID{},      // feed_id
		&sql.NullString{}, // audio_url
		&sql.NullString{}, // caption
		&sql.NullString{}, // transcript
		&sql.NullString{}, // privacy
		&sql.NullString{}, // owner_id
		&sql.NullTime{},   // created_at
		&sql.NullTime{},   // updated_at
		&sql.NullTime{},   // deleted_at
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Feed fields.
func (f *Feed) assignValues(values ...interface{}) error {
	if m, n := len(values), len(feed.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	f.ID = int64(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*uuid.UUID); !ok {
		return fmt.Errorf("unexpected type %T for field feed_id", values[0])
	} else if value != nil {
		f.FeedID = *value
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field audio_url", values[1])
	} else if value.Valid {
		f.AudioURL = value.String
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field caption", values[2])
	} else if value.Valid {
		f.Caption = new(string)
		*f.Caption = value.String
	}
	if value, ok := values[3].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field transcript", values[3])
	} else if value.Valid {
		f.Transcript = new(string)
		*f.Transcript = value.String
	}
	if value, ok := values[4].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field privacy", values[4])
	} else if value.Valid {
		f.Privacy = feed.Privacy(value.String)
	}
	if value, ok := values[5].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field owner_id", values[5])
	} else if value.Valid {
		f.OwnerID = value.String
	}
	if value, ok := values[6].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field created_at", values[6])
	} else if value.Valid {
		f.CreatedAt = value.Time
	}
	if value, ok := values[7].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field updated_at", values[7])
	} else if value.Valid {
		f.UpdatedAt = value.Time
	}
	if value, ok := values[8].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field deleted_at", values[8])
	} else if value.Valid {
		f.DeletedAt = new(time.Time)
		*f.DeletedAt = value.Time
	}
	return nil
}

// Update returns a builder for updating this Feed.
// Note that, you need to call Feed.Unwrap() before calling this method, if this Feed
// was returned from a transaction, and the transaction was committed or rolled back.
func (f *Feed) Update() *FeedUpdateOne {
	return (&FeedClient{config: f.config}).UpdateOne(f)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (f *Feed) Unwrap() *Feed {
	tx, ok := f.config.driver.(*txDriver)
	if !ok {
		panic("ent: Feed is not a transactional entity")
	}
	f.config.driver = tx.drv
	return f
}

// String implements the fmt.Stringer.
func (f *Feed) String() string {
	var builder strings.Builder
	builder.WriteString("Feed(")
	builder.WriteString(fmt.Sprintf("id=%v", f.ID))
	builder.WriteString(", feed_id=")
	builder.WriteString(fmt.Sprintf("%v", f.FeedID))
	builder.WriteString(", audio_url=")
	builder.WriteString(f.AudioURL)
	if v := f.Caption; v != nil {
		builder.WriteString(", caption=")
		builder.WriteString(*v)
	}
	if v := f.Transcript; v != nil {
		builder.WriteString(", transcript=")
		builder.WriteString(*v)
	}
	builder.WriteString(", privacy=")
	builder.WriteString(fmt.Sprintf("%v", f.Privacy))
	builder.WriteString(", owner_id=")
	builder.WriteString(f.OwnerID)
	builder.WriteString(", created_at=")
	builder.WriteString(f.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(f.UpdatedAt.Format(time.ANSIC))
	if v := f.DeletedAt; v != nil {
		builder.WriteString(", deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteByte(')')
	return builder.String()
}

// Feeds is a parsable slice of Feed.
type Feeds []*Feed

func (f Feeds) config(cfg config) {
	for _i := range f {
		f[_i].config = cfg
	}
}
