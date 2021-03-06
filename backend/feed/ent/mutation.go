// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/wcse/monorepo/backend/feed/ent/feed"

	"github.com/facebook/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeFeed = "Feed"
)

// FeedMutation represents an operation that mutate the Feeds
// nodes in the graph.
type FeedMutation struct {
	config
	op            Op
	typ           string
	id            *int64
	feed_id       *uuid.UUID
	audio_url     *string
	caption       *string
	transcript    *string
	privacy       *feed.Privacy
	owner_id      *string
	created_at    *time.Time
	updated_at    *time.Time
	deleted_at    *time.Time
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Feed, error)
}

var _ ent.Mutation = (*FeedMutation)(nil)

// feedOption allows to manage the mutation configuration using functional options.
type feedOption func(*FeedMutation)

// newFeedMutation creates new mutation for $n.Name.
func newFeedMutation(c config, op Op, opts ...feedOption) *FeedMutation {
	m := &FeedMutation{
		config:        c,
		op:            op,
		typ:           TypeFeed,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withFeedID sets the id field of the mutation.
func withFeedID(id int64) feedOption {
	return func(m *FeedMutation) {
		var (
			err   error
			once  sync.Once
			value *Feed
		)
		m.oldValue = func(ctx context.Context) (*Feed, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Feed.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withFeed sets the old Feed of the mutation.
func withFeed(node *Feed) feedOption {
	return func(m *FeedMutation) {
		m.oldValue = func(context.Context) (*Feed, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m FeedMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m FeedMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that, this
// operation is accepted only on Feed creation.
func (m *FeedMutation) SetID(id int64) {
	m.id = &id
}

// ID returns the id value in the mutation. Note that, the id
// is available only if it was provided to the builder.
func (m *FeedMutation) ID() (id int64, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetFeedID sets the feed_id field.
func (m *FeedMutation) SetFeedID(u uuid.UUID) {
	m.feed_id = &u
}

// FeedID returns the feed_id value in the mutation.
func (m *FeedMutation) FeedID() (r uuid.UUID, exists bool) {
	v := m.feed_id
	if v == nil {
		return
	}
	return *v, true
}

// OldFeedID returns the old feed_id value of the Feed.
// If the Feed object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *FeedMutation) OldFeedID(ctx context.Context) (v uuid.UUID, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldFeedID is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldFeedID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldFeedID: %w", err)
	}
	return oldValue.FeedID, nil
}

// ResetFeedID reset all changes of the "feed_id" field.
func (m *FeedMutation) ResetFeedID() {
	m.feed_id = nil
}

// SetAudioURL sets the audio_url field.
func (m *FeedMutation) SetAudioURL(s string) {
	m.audio_url = &s
}

// AudioURL returns the audio_url value in the mutation.
func (m *FeedMutation) AudioURL() (r string, exists bool) {
	v := m.audio_url
	if v == nil {
		return
	}
	return *v, true
}

// OldAudioURL returns the old audio_url value of the Feed.
// If the Feed object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *FeedMutation) OldAudioURL(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldAudioURL is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldAudioURL requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldAudioURL: %w", err)
	}
	return oldValue.AudioURL, nil
}

// ResetAudioURL reset all changes of the "audio_url" field.
func (m *FeedMutation) ResetAudioURL() {
	m.audio_url = nil
}

// SetCaption sets the caption field.
func (m *FeedMutation) SetCaption(s string) {
	m.caption = &s
}

// Caption returns the caption value in the mutation.
func (m *FeedMutation) Caption() (r string, exists bool) {
	v := m.caption
	if v == nil {
		return
	}
	return *v, true
}

// OldCaption returns the old caption value of the Feed.
// If the Feed object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *FeedMutation) OldCaption(ctx context.Context) (v *string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldCaption is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldCaption requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCaption: %w", err)
	}
	return oldValue.Caption, nil
}

// ClearCaption clears the value of caption.
func (m *FeedMutation) ClearCaption() {
	m.caption = nil
	m.clearedFields[feed.FieldCaption] = struct{}{}
}

// CaptionCleared returns if the field caption was cleared in this mutation.
func (m *FeedMutation) CaptionCleared() bool {
	_, ok := m.clearedFields[feed.FieldCaption]
	return ok
}

// ResetCaption reset all changes of the "caption" field.
func (m *FeedMutation) ResetCaption() {
	m.caption = nil
	delete(m.clearedFields, feed.FieldCaption)
}

// SetTranscript sets the transcript field.
func (m *FeedMutation) SetTranscript(s string) {
	m.transcript = &s
}

// Transcript returns the transcript value in the mutation.
func (m *FeedMutation) Transcript() (r string, exists bool) {
	v := m.transcript
	if v == nil {
		return
	}
	return *v, true
}

// OldTranscript returns the old transcript value of the Feed.
// If the Feed object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *FeedMutation) OldTranscript(ctx context.Context) (v *string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldTranscript is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldTranscript requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldTranscript: %w", err)
	}
	return oldValue.Transcript, nil
}

// ClearTranscript clears the value of transcript.
func (m *FeedMutation) ClearTranscript() {
	m.transcript = nil
	m.clearedFields[feed.FieldTranscript] = struct{}{}
}

// TranscriptCleared returns if the field transcript was cleared in this mutation.
func (m *FeedMutation) TranscriptCleared() bool {
	_, ok := m.clearedFields[feed.FieldTranscript]
	return ok
}

// ResetTranscript reset all changes of the "transcript" field.
func (m *FeedMutation) ResetTranscript() {
	m.transcript = nil
	delete(m.clearedFields, feed.FieldTranscript)
}

// SetPrivacy sets the privacy field.
func (m *FeedMutation) SetPrivacy(f feed.Privacy) {
	m.privacy = &f
}

// Privacy returns the privacy value in the mutation.
func (m *FeedMutation) Privacy() (r feed.Privacy, exists bool) {
	v := m.privacy
	if v == nil {
		return
	}
	return *v, true
}

// OldPrivacy returns the old privacy value of the Feed.
// If the Feed object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *FeedMutation) OldPrivacy(ctx context.Context) (v feed.Privacy, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldPrivacy is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldPrivacy requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPrivacy: %w", err)
	}
	return oldValue.Privacy, nil
}

// ResetPrivacy reset all changes of the "privacy" field.
func (m *FeedMutation) ResetPrivacy() {
	m.privacy = nil
}

// SetOwnerID sets the owner_id field.
func (m *FeedMutation) SetOwnerID(s string) {
	m.owner_id = &s
}

// OwnerID returns the owner_id value in the mutation.
func (m *FeedMutation) OwnerID() (r string, exists bool) {
	v := m.owner_id
	if v == nil {
		return
	}
	return *v, true
}

// OldOwnerID returns the old owner_id value of the Feed.
// If the Feed object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *FeedMutation) OldOwnerID(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldOwnerID is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldOwnerID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldOwnerID: %w", err)
	}
	return oldValue.OwnerID, nil
}

// ResetOwnerID reset all changes of the "owner_id" field.
func (m *FeedMutation) ResetOwnerID() {
	m.owner_id = nil
}

// SetCreatedAt sets the created_at field.
func (m *FeedMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the created_at value in the mutation.
func (m *FeedMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old created_at value of the Feed.
// If the Feed object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *FeedMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldCreatedAt is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt reset all changes of the "created_at" field.
func (m *FeedMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetUpdatedAt sets the updated_at field.
func (m *FeedMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the updated_at value in the mutation.
func (m *FeedMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old updated_at value of the Feed.
// If the Feed object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *FeedMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldUpdatedAt is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// ResetUpdatedAt reset all changes of the "updated_at" field.
func (m *FeedMutation) ResetUpdatedAt() {
	m.updated_at = nil
}

// SetDeletedAt sets the deleted_at field.
func (m *FeedMutation) SetDeletedAt(t time.Time) {
	m.deleted_at = &t
}

// DeletedAt returns the deleted_at value in the mutation.
func (m *FeedMutation) DeletedAt() (r time.Time, exists bool) {
	v := m.deleted_at
	if v == nil {
		return
	}
	return *v, true
}

// OldDeletedAt returns the old deleted_at value of the Feed.
// If the Feed object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *FeedMutation) OldDeletedAt(ctx context.Context) (v *time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldDeletedAt is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldDeletedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldDeletedAt: %w", err)
	}
	return oldValue.DeletedAt, nil
}

// ClearDeletedAt clears the value of deleted_at.
func (m *FeedMutation) ClearDeletedAt() {
	m.deleted_at = nil
	m.clearedFields[feed.FieldDeletedAt] = struct{}{}
}

// DeletedAtCleared returns if the field deleted_at was cleared in this mutation.
func (m *FeedMutation) DeletedAtCleared() bool {
	_, ok := m.clearedFields[feed.FieldDeletedAt]
	return ok
}

// ResetDeletedAt reset all changes of the "deleted_at" field.
func (m *FeedMutation) ResetDeletedAt() {
	m.deleted_at = nil
	delete(m.clearedFields, feed.FieldDeletedAt)
}

// Op returns the operation name.
func (m *FeedMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Feed).
func (m *FeedMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during
// this mutation. Note that, in order to get all numeric
// fields that were in/decremented, call AddedFields().
func (m *FeedMutation) Fields() []string {
	fields := make([]string, 0, 9)
	if m.feed_id != nil {
		fields = append(fields, feed.FieldFeedID)
	}
	if m.audio_url != nil {
		fields = append(fields, feed.FieldAudioURL)
	}
	if m.caption != nil {
		fields = append(fields, feed.FieldCaption)
	}
	if m.transcript != nil {
		fields = append(fields, feed.FieldTranscript)
	}
	if m.privacy != nil {
		fields = append(fields, feed.FieldPrivacy)
	}
	if m.owner_id != nil {
		fields = append(fields, feed.FieldOwnerID)
	}
	if m.created_at != nil {
		fields = append(fields, feed.FieldCreatedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, feed.FieldUpdatedAt)
	}
	if m.deleted_at != nil {
		fields = append(fields, feed.FieldDeletedAt)
	}
	return fields
}

// Field returns the value of a field with the given name.
// The second boolean value indicates that this field was
// not set, or was not define in the schema.
func (m *FeedMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case feed.FieldFeedID:
		return m.FeedID()
	case feed.FieldAudioURL:
		return m.AudioURL()
	case feed.FieldCaption:
		return m.Caption()
	case feed.FieldTranscript:
		return m.Transcript()
	case feed.FieldPrivacy:
		return m.Privacy()
	case feed.FieldOwnerID:
		return m.OwnerID()
	case feed.FieldCreatedAt:
		return m.CreatedAt()
	case feed.FieldUpdatedAt:
		return m.UpdatedAt()
	case feed.FieldDeletedAt:
		return m.DeletedAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database.
// An error is returned if the mutation operation is not UpdateOne,
// or the query to the database was failed.
func (m *FeedMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case feed.FieldFeedID:
		return m.OldFeedID(ctx)
	case feed.FieldAudioURL:
		return m.OldAudioURL(ctx)
	case feed.FieldCaption:
		return m.OldCaption(ctx)
	case feed.FieldTranscript:
		return m.OldTranscript(ctx)
	case feed.FieldPrivacy:
		return m.OldPrivacy(ctx)
	case feed.FieldOwnerID:
		return m.OldOwnerID(ctx)
	case feed.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case feed.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	case feed.FieldDeletedAt:
		return m.OldDeletedAt(ctx)
	}
	return nil, fmt.Errorf("unknown Feed field %s", name)
}

// SetField sets the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *FeedMutation) SetField(name string, value ent.Value) error {
	switch name {
	case feed.FieldFeedID:
		v, ok := value.(uuid.UUID)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetFeedID(v)
		return nil
	case feed.FieldAudioURL:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetAudioURL(v)
		return nil
	case feed.FieldCaption:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCaption(v)
		return nil
	case feed.FieldTranscript:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetTranscript(v)
		return nil
	case feed.FieldPrivacy:
		v, ok := value.(feed.Privacy)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPrivacy(v)
		return nil
	case feed.FieldOwnerID:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetOwnerID(v)
		return nil
	case feed.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case feed.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	case feed.FieldDeletedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDeletedAt(v)
		return nil
	}
	return fmt.Errorf("unknown Feed field %s", name)
}

// AddedFields returns all numeric fields that were incremented
// or decremented during this mutation.
func (m *FeedMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was in/decremented
// from a field with the given name. The second value indicates
// that this field was not set, or was not define in the schema.
func (m *FeedMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *FeedMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Feed numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared
// during this mutation.
func (m *FeedMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(feed.FieldCaption) {
		fields = append(fields, feed.FieldCaption)
	}
	if m.FieldCleared(feed.FieldTranscript) {
		fields = append(fields, feed.FieldTranscript)
	}
	if m.FieldCleared(feed.FieldDeletedAt) {
		fields = append(fields, feed.FieldDeletedAt)
	}
	return fields
}

// FieldCleared returns a boolean indicates if this field was
// cleared in this mutation.
func (m *FeedMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value for the given name. It returns an
// error if the field is not defined in the schema.
func (m *FeedMutation) ClearField(name string) error {
	switch name {
	case feed.FieldCaption:
		m.ClearCaption()
		return nil
	case feed.FieldTranscript:
		m.ClearTranscript()
		return nil
	case feed.FieldDeletedAt:
		m.ClearDeletedAt()
		return nil
	}
	return fmt.Errorf("unknown Feed nullable field %s", name)
}

// ResetField resets all changes in the mutation regarding the
// given field name. It returns an error if the field is not
// defined in the schema.
func (m *FeedMutation) ResetField(name string) error {
	switch name {
	case feed.FieldFeedID:
		m.ResetFeedID()
		return nil
	case feed.FieldAudioURL:
		m.ResetAudioURL()
		return nil
	case feed.FieldCaption:
		m.ResetCaption()
		return nil
	case feed.FieldTranscript:
		m.ResetTranscript()
		return nil
	case feed.FieldPrivacy:
		m.ResetPrivacy()
		return nil
	case feed.FieldOwnerID:
		m.ResetOwnerID()
		return nil
	case feed.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case feed.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	case feed.FieldDeletedAt:
		m.ResetDeletedAt()
		return nil
	}
	return fmt.Errorf("unknown Feed field %s", name)
}

// AddedEdges returns all edge names that were set/added in this
// mutation.
func (m *FeedMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all ids (to other nodes) that were added for
// the given edge name.
func (m *FeedMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this
// mutation.
func (m *FeedMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all ids (to other nodes) that were removed for
// the given edge name.
func (m *FeedMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this
// mutation.
func (m *FeedMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean indicates if this edge was
// cleared in this mutation.
func (m *FeedMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value for the given name. It returns an
// error if the edge name is not defined in the schema.
func (m *FeedMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Feed unique edge %s", name)
}

// ResetEdge resets all changes in the mutation regarding the
// given edge name. It returns an error if the edge is not
// defined in the schema.
func (m *FeedMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Feed edge %s", name)
}
