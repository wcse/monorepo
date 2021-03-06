// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/google/uuid"
	"github.com/wcse/monorepo/backend/feed/ent/feed"
	"github.com/wcse/monorepo/backend/feed/ent/predicate"
)

// FeedUpdate is the builder for updating Feed entities.
type FeedUpdate struct {
	config
	hooks      []Hook
	mutation   *FeedMutation
	predicates []predicate.Feed
}

// Where adds a new predicate for the builder.
func (fu *FeedUpdate) Where(ps ...predicate.Feed) *FeedUpdate {
	fu.predicates = append(fu.predicates, ps...)
	return fu
}

// SetFeedID sets the feed_id field.
func (fu *FeedUpdate) SetFeedID(u uuid.UUID) *FeedUpdate {
	fu.mutation.SetFeedID(u)
	return fu
}

// SetAudioURL sets the audio_url field.
func (fu *FeedUpdate) SetAudioURL(s string) *FeedUpdate {
	fu.mutation.SetAudioURL(s)
	return fu
}

// SetCaption sets the caption field.
func (fu *FeedUpdate) SetCaption(s string) *FeedUpdate {
	fu.mutation.SetCaption(s)
	return fu
}

// SetNillableCaption sets the caption field if the given value is not nil.
func (fu *FeedUpdate) SetNillableCaption(s *string) *FeedUpdate {
	if s != nil {
		fu.SetCaption(*s)
	}
	return fu
}

// ClearCaption clears the value of caption.
func (fu *FeedUpdate) ClearCaption() *FeedUpdate {
	fu.mutation.ClearCaption()
	return fu
}

// SetTranscript sets the transcript field.
func (fu *FeedUpdate) SetTranscript(s string) *FeedUpdate {
	fu.mutation.SetTranscript(s)
	return fu
}

// SetNillableTranscript sets the transcript field if the given value is not nil.
func (fu *FeedUpdate) SetNillableTranscript(s *string) *FeedUpdate {
	if s != nil {
		fu.SetTranscript(*s)
	}
	return fu
}

// ClearTranscript clears the value of transcript.
func (fu *FeedUpdate) ClearTranscript() *FeedUpdate {
	fu.mutation.ClearTranscript()
	return fu
}

// SetPrivacy sets the privacy field.
func (fu *FeedUpdate) SetPrivacy(f feed.Privacy) *FeedUpdate {
	fu.mutation.SetPrivacy(f)
	return fu
}

// SetNillablePrivacy sets the privacy field if the given value is not nil.
func (fu *FeedUpdate) SetNillablePrivacy(f *feed.Privacy) *FeedUpdate {
	if f != nil {
		fu.SetPrivacy(*f)
	}
	return fu
}

// SetOwnerID sets the owner_id field.
func (fu *FeedUpdate) SetOwnerID(s string) *FeedUpdate {
	fu.mutation.SetOwnerID(s)
	return fu
}

// SetCreatedAt sets the created_at field.
func (fu *FeedUpdate) SetCreatedAt(t time.Time) *FeedUpdate {
	fu.mutation.SetCreatedAt(t)
	return fu
}

// SetNillableCreatedAt sets the created_at field if the given value is not nil.
func (fu *FeedUpdate) SetNillableCreatedAt(t *time.Time) *FeedUpdate {
	if t != nil {
		fu.SetCreatedAt(*t)
	}
	return fu
}

// SetUpdatedAt sets the updated_at field.
func (fu *FeedUpdate) SetUpdatedAt(t time.Time) *FeedUpdate {
	fu.mutation.SetUpdatedAt(t)
	return fu
}

// SetNillableUpdatedAt sets the updated_at field if the given value is not nil.
func (fu *FeedUpdate) SetNillableUpdatedAt(t *time.Time) *FeedUpdate {
	if t != nil {
		fu.SetUpdatedAt(*t)
	}
	return fu
}

// SetDeletedAt sets the deleted_at field.
func (fu *FeedUpdate) SetDeletedAt(t time.Time) *FeedUpdate {
	fu.mutation.SetDeletedAt(t)
	return fu
}

// SetNillableDeletedAt sets the deleted_at field if the given value is not nil.
func (fu *FeedUpdate) SetNillableDeletedAt(t *time.Time) *FeedUpdate {
	if t != nil {
		fu.SetDeletedAt(*t)
	}
	return fu
}

// ClearDeletedAt clears the value of deleted_at.
func (fu *FeedUpdate) ClearDeletedAt() *FeedUpdate {
	fu.mutation.ClearDeletedAt()
	return fu
}

// Mutation returns the FeedMutation object of the builder.
func (fu *FeedUpdate) Mutation() *FeedMutation {
	return fu.mutation
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (fu *FeedUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := fu.mutation.Privacy(); ok {
		if err := feed.PrivacyValidator(v); err != nil {
			return 0, &ValidationError{Name: "privacy", err: fmt.Errorf("ent: validator failed for field \"privacy\": %w", err)}
		}
	}
	var (
		err      error
		affected int
	)
	if len(fu.hooks) == 0 {
		affected, err = fu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FeedMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			fu.mutation = mutation
			affected, err = fu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(fu.hooks) - 1; i >= 0; i-- {
			mut = fu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (fu *FeedUpdate) SaveX(ctx context.Context) int {
	affected, err := fu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (fu *FeedUpdate) Exec(ctx context.Context) error {
	_, err := fu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fu *FeedUpdate) ExecX(ctx context.Context) {
	if err := fu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (fu *FeedUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   feed.Table,
			Columns: feed.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: feed.FieldID,
			},
		},
	}
	if ps := fu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fu.mutation.FeedID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: feed.FieldFeedID,
		})
	}
	if value, ok := fu.mutation.AudioURL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: feed.FieldAudioURL,
		})
	}
	if value, ok := fu.mutation.Caption(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: feed.FieldCaption,
		})
	}
	if fu.mutation.CaptionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: feed.FieldCaption,
		})
	}
	if value, ok := fu.mutation.Transcript(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: feed.FieldTranscript,
		})
	}
	if fu.mutation.TranscriptCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: feed.FieldTranscript,
		})
	}
	if value, ok := fu.mutation.Privacy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: feed.FieldPrivacy,
		})
	}
	if value, ok := fu.mutation.OwnerID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: feed.FieldOwnerID,
		})
	}
	if value, ok := fu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: feed.FieldCreatedAt,
		})
	}
	if value, ok := fu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: feed.FieldUpdatedAt,
		})
	}
	if value, ok := fu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: feed.FieldDeletedAt,
		})
	}
	if fu.mutation.DeletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: feed.FieldDeletedAt,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, fu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{feed.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// FeedUpdateOne is the builder for updating a single Feed entity.
type FeedUpdateOne struct {
	config
	hooks    []Hook
	mutation *FeedMutation
}

// SetFeedID sets the feed_id field.
func (fuo *FeedUpdateOne) SetFeedID(u uuid.UUID) *FeedUpdateOne {
	fuo.mutation.SetFeedID(u)
	return fuo
}

// SetAudioURL sets the audio_url field.
func (fuo *FeedUpdateOne) SetAudioURL(s string) *FeedUpdateOne {
	fuo.mutation.SetAudioURL(s)
	return fuo
}

// SetCaption sets the caption field.
func (fuo *FeedUpdateOne) SetCaption(s string) *FeedUpdateOne {
	fuo.mutation.SetCaption(s)
	return fuo
}

// SetNillableCaption sets the caption field if the given value is not nil.
func (fuo *FeedUpdateOne) SetNillableCaption(s *string) *FeedUpdateOne {
	if s != nil {
		fuo.SetCaption(*s)
	}
	return fuo
}

// ClearCaption clears the value of caption.
func (fuo *FeedUpdateOne) ClearCaption() *FeedUpdateOne {
	fuo.mutation.ClearCaption()
	return fuo
}

// SetTranscript sets the transcript field.
func (fuo *FeedUpdateOne) SetTranscript(s string) *FeedUpdateOne {
	fuo.mutation.SetTranscript(s)
	return fuo
}

// SetNillableTranscript sets the transcript field if the given value is not nil.
func (fuo *FeedUpdateOne) SetNillableTranscript(s *string) *FeedUpdateOne {
	if s != nil {
		fuo.SetTranscript(*s)
	}
	return fuo
}

// ClearTranscript clears the value of transcript.
func (fuo *FeedUpdateOne) ClearTranscript() *FeedUpdateOne {
	fuo.mutation.ClearTranscript()
	return fuo
}

// SetPrivacy sets the privacy field.
func (fuo *FeedUpdateOne) SetPrivacy(f feed.Privacy) *FeedUpdateOne {
	fuo.mutation.SetPrivacy(f)
	return fuo
}

// SetNillablePrivacy sets the privacy field if the given value is not nil.
func (fuo *FeedUpdateOne) SetNillablePrivacy(f *feed.Privacy) *FeedUpdateOne {
	if f != nil {
		fuo.SetPrivacy(*f)
	}
	return fuo
}

// SetOwnerID sets the owner_id field.
func (fuo *FeedUpdateOne) SetOwnerID(s string) *FeedUpdateOne {
	fuo.mutation.SetOwnerID(s)
	return fuo
}

// SetCreatedAt sets the created_at field.
func (fuo *FeedUpdateOne) SetCreatedAt(t time.Time) *FeedUpdateOne {
	fuo.mutation.SetCreatedAt(t)
	return fuo
}

// SetNillableCreatedAt sets the created_at field if the given value is not nil.
func (fuo *FeedUpdateOne) SetNillableCreatedAt(t *time.Time) *FeedUpdateOne {
	if t != nil {
		fuo.SetCreatedAt(*t)
	}
	return fuo
}

// SetUpdatedAt sets the updated_at field.
func (fuo *FeedUpdateOne) SetUpdatedAt(t time.Time) *FeedUpdateOne {
	fuo.mutation.SetUpdatedAt(t)
	return fuo
}

// SetNillableUpdatedAt sets the updated_at field if the given value is not nil.
func (fuo *FeedUpdateOne) SetNillableUpdatedAt(t *time.Time) *FeedUpdateOne {
	if t != nil {
		fuo.SetUpdatedAt(*t)
	}
	return fuo
}

// SetDeletedAt sets the deleted_at field.
func (fuo *FeedUpdateOne) SetDeletedAt(t time.Time) *FeedUpdateOne {
	fuo.mutation.SetDeletedAt(t)
	return fuo
}

// SetNillableDeletedAt sets the deleted_at field if the given value is not nil.
func (fuo *FeedUpdateOne) SetNillableDeletedAt(t *time.Time) *FeedUpdateOne {
	if t != nil {
		fuo.SetDeletedAt(*t)
	}
	return fuo
}

// ClearDeletedAt clears the value of deleted_at.
func (fuo *FeedUpdateOne) ClearDeletedAt() *FeedUpdateOne {
	fuo.mutation.ClearDeletedAt()
	return fuo
}

// Mutation returns the FeedMutation object of the builder.
func (fuo *FeedUpdateOne) Mutation() *FeedMutation {
	return fuo.mutation
}

// Save executes the query and returns the updated entity.
func (fuo *FeedUpdateOne) Save(ctx context.Context) (*Feed, error) {
	if v, ok := fuo.mutation.Privacy(); ok {
		if err := feed.PrivacyValidator(v); err != nil {
			return nil, &ValidationError{Name: "privacy", err: fmt.Errorf("ent: validator failed for field \"privacy\": %w", err)}
		}
	}
	var (
		err  error
		node *Feed
	)
	if len(fuo.hooks) == 0 {
		node, err = fuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FeedMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			fuo.mutation = mutation
			node, err = fuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(fuo.hooks) - 1; i >= 0; i-- {
			mut = fuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (fuo *FeedUpdateOne) SaveX(ctx context.Context) *Feed {
	f, err := fuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return f
}

// Exec executes the query on the entity.
func (fuo *FeedUpdateOne) Exec(ctx context.Context) error {
	_, err := fuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fuo *FeedUpdateOne) ExecX(ctx context.Context) {
	if err := fuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (fuo *FeedUpdateOne) sqlSave(ctx context.Context) (f *Feed, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   feed.Table,
			Columns: feed.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: feed.FieldID,
			},
		},
	}
	id, ok := fuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Feed.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := fuo.mutation.FeedID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: feed.FieldFeedID,
		})
	}
	if value, ok := fuo.mutation.AudioURL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: feed.FieldAudioURL,
		})
	}
	if value, ok := fuo.mutation.Caption(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: feed.FieldCaption,
		})
	}
	if fuo.mutation.CaptionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: feed.FieldCaption,
		})
	}
	if value, ok := fuo.mutation.Transcript(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: feed.FieldTranscript,
		})
	}
	if fuo.mutation.TranscriptCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: feed.FieldTranscript,
		})
	}
	if value, ok := fuo.mutation.Privacy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: feed.FieldPrivacy,
		})
	}
	if value, ok := fuo.mutation.OwnerID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: feed.FieldOwnerID,
		})
	}
	if value, ok := fuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: feed.FieldCreatedAt,
		})
	}
	if value, ok := fuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: feed.FieldUpdatedAt,
		})
	}
	if value, ok := fuo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: feed.FieldDeletedAt,
		})
	}
	if fuo.mutation.DeletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: feed.FieldDeletedAt,
		})
	}
	f = &Feed{config: fuo.config}
	_spec.Assign = f.assignValues
	_spec.ScanValues = f.scanValues()
	if err = sqlgraph.UpdateNode(ctx, fuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{feed.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return f, nil
}
