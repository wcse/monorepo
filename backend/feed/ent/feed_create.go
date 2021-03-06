// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/google/uuid"
	"github.com/wcse/monorepo/backend/feed/ent/feed"
)

// FeedCreate is the builder for creating a Feed entity.
type FeedCreate struct {
	config
	mutation *FeedMutation
	hooks    []Hook
}

// SetFeedID sets the feed_id field.
func (fc *FeedCreate) SetFeedID(u uuid.UUID) *FeedCreate {
	fc.mutation.SetFeedID(u)
	return fc
}

// SetAudioURL sets the audio_url field.
func (fc *FeedCreate) SetAudioURL(s string) *FeedCreate {
	fc.mutation.SetAudioURL(s)
	return fc
}

// SetCaption sets the caption field.
func (fc *FeedCreate) SetCaption(s string) *FeedCreate {
	fc.mutation.SetCaption(s)
	return fc
}

// SetNillableCaption sets the caption field if the given value is not nil.
func (fc *FeedCreate) SetNillableCaption(s *string) *FeedCreate {
	if s != nil {
		fc.SetCaption(*s)
	}
	return fc
}

// SetTranscript sets the transcript field.
func (fc *FeedCreate) SetTranscript(s string) *FeedCreate {
	fc.mutation.SetTranscript(s)
	return fc
}

// SetNillableTranscript sets the transcript field if the given value is not nil.
func (fc *FeedCreate) SetNillableTranscript(s *string) *FeedCreate {
	if s != nil {
		fc.SetTranscript(*s)
	}
	return fc
}

// SetPrivacy sets the privacy field.
func (fc *FeedCreate) SetPrivacy(f feed.Privacy) *FeedCreate {
	fc.mutation.SetPrivacy(f)
	return fc
}

// SetNillablePrivacy sets the privacy field if the given value is not nil.
func (fc *FeedCreate) SetNillablePrivacy(f *feed.Privacy) *FeedCreate {
	if f != nil {
		fc.SetPrivacy(*f)
	}
	return fc
}

// SetOwnerID sets the owner_id field.
func (fc *FeedCreate) SetOwnerID(s string) *FeedCreate {
	fc.mutation.SetOwnerID(s)
	return fc
}

// SetCreatedAt sets the created_at field.
func (fc *FeedCreate) SetCreatedAt(t time.Time) *FeedCreate {
	fc.mutation.SetCreatedAt(t)
	return fc
}

// SetNillableCreatedAt sets the created_at field if the given value is not nil.
func (fc *FeedCreate) SetNillableCreatedAt(t *time.Time) *FeedCreate {
	if t != nil {
		fc.SetCreatedAt(*t)
	}
	return fc
}

// SetUpdatedAt sets the updated_at field.
func (fc *FeedCreate) SetUpdatedAt(t time.Time) *FeedCreate {
	fc.mutation.SetUpdatedAt(t)
	return fc
}

// SetNillableUpdatedAt sets the updated_at field if the given value is not nil.
func (fc *FeedCreate) SetNillableUpdatedAt(t *time.Time) *FeedCreate {
	if t != nil {
		fc.SetUpdatedAt(*t)
	}
	return fc
}

// SetDeletedAt sets the deleted_at field.
func (fc *FeedCreate) SetDeletedAt(t time.Time) *FeedCreate {
	fc.mutation.SetDeletedAt(t)
	return fc
}

// SetNillableDeletedAt sets the deleted_at field if the given value is not nil.
func (fc *FeedCreate) SetNillableDeletedAt(t *time.Time) *FeedCreate {
	if t != nil {
		fc.SetDeletedAt(*t)
	}
	return fc
}

// SetID sets the id field.
func (fc *FeedCreate) SetID(i int64) *FeedCreate {
	fc.mutation.SetID(i)
	return fc
}

// Mutation returns the FeedMutation object of the builder.
func (fc *FeedCreate) Mutation() *FeedMutation {
	return fc.mutation
}

// Save creates the Feed in the database.
func (fc *FeedCreate) Save(ctx context.Context) (*Feed, error) {
	if err := fc.preSave(); err != nil {
		return nil, err
	}
	var (
		err  error
		node *Feed
	)
	if len(fc.hooks) == 0 {
		node, err = fc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FeedMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			fc.mutation = mutation
			node, err = fc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(fc.hooks) - 1; i >= 0; i-- {
			mut = fc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (fc *FeedCreate) SaveX(ctx context.Context) *Feed {
	v, err := fc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (fc *FeedCreate) preSave() error {
	if _, ok := fc.mutation.FeedID(); !ok {
		v := feed.DefaultFeedID()
		fc.mutation.SetFeedID(v)
	}
	if _, ok := fc.mutation.AudioURL(); !ok {
		return &ValidationError{Name: "audio_url", err: errors.New("ent: missing required field \"audio_url\"")}
	}
	if _, ok := fc.mutation.Privacy(); !ok {
		v := feed.DefaultPrivacy
		fc.mutation.SetPrivacy(v)
	}
	if v, ok := fc.mutation.Privacy(); ok {
		if err := feed.PrivacyValidator(v); err != nil {
			return &ValidationError{Name: "privacy", err: fmt.Errorf("ent: validator failed for field \"privacy\": %w", err)}
		}
	}
	if _, ok := fc.mutation.OwnerID(); !ok {
		return &ValidationError{Name: "owner_id", err: errors.New("ent: missing required field \"owner_id\"")}
	}
	if _, ok := fc.mutation.CreatedAt(); !ok {
		v := feed.DefaultCreatedAt()
		fc.mutation.SetCreatedAt(v)
	}
	if _, ok := fc.mutation.UpdatedAt(); !ok {
		v := feed.DefaultUpdatedAt()
		fc.mutation.SetUpdatedAt(v)
	}
	return nil
}

func (fc *FeedCreate) sqlSave(ctx context.Context) (*Feed, error) {
	f, _spec := fc.createSpec()
	if err := sqlgraph.CreateNode(ctx, fc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	if f.ID == 0 {
		id := _spec.ID.Value.(int64)
		f.ID = int64(id)
	}
	return f, nil
}

func (fc *FeedCreate) createSpec() (*Feed, *sqlgraph.CreateSpec) {
	var (
		f     = &Feed{config: fc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: feed.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: feed.FieldID,
			},
		}
	)
	if id, ok := fc.mutation.ID(); ok {
		f.ID = id
		_spec.ID.Value = id
	}
	if value, ok := fc.mutation.FeedID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: feed.FieldFeedID,
		})
		f.FeedID = value
	}
	if value, ok := fc.mutation.AudioURL(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: feed.FieldAudioURL,
		})
		f.AudioURL = value
	}
	if value, ok := fc.mutation.Caption(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: feed.FieldCaption,
		})
		f.Caption = &value
	}
	if value, ok := fc.mutation.Transcript(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: feed.FieldTranscript,
		})
		f.Transcript = &value
	}
	if value, ok := fc.mutation.Privacy(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: feed.FieldPrivacy,
		})
		f.Privacy = value
	}
	if value, ok := fc.mutation.OwnerID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: feed.FieldOwnerID,
		})
		f.OwnerID = value
	}
	if value, ok := fc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: feed.FieldCreatedAt,
		})
		f.CreatedAt = value
	}
	if value, ok := fc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: feed.FieldUpdatedAt,
		})
		f.UpdatedAt = value
	}
	if value, ok := fc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: feed.FieldDeletedAt,
		})
		f.DeletedAt = &value
	}
	return f, _spec
}

// FeedCreateBulk is the builder for creating a bulk of Feed entities.
type FeedCreateBulk struct {
	config
	builders []*FeedCreate
}

// Save creates the Feed entities in the database.
func (fcb *FeedCreateBulk) Save(ctx context.Context) ([]*Feed, error) {
	specs := make([]*sqlgraph.CreateSpec, len(fcb.builders))
	nodes := make([]*Feed, len(fcb.builders))
	mutators := make([]Mutator, len(fcb.builders))
	for i := range fcb.builders {
		func(i int, root context.Context) {
			builder := fcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				if err := builder.preSave(); err != nil {
					return nil, err
				}
				mutation, ok := m.(*FeedMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, fcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, fcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				if nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int64(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, fcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (fcb *FeedCreateBulk) SaveX(ctx context.Context) []*Feed {
	v, err := fcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
