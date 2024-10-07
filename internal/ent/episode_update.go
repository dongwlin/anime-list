// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/dongwlin/anime-list/internal/ent/episode"
	"github.com/dongwlin/anime-list/internal/ent/predicate"
	"github.com/dongwlin/anime-list/internal/ent/season"
)

// EpisodeUpdate is the builder for updating Episode entities.
type EpisodeUpdate struct {
	config
	hooks    []Hook
	mutation *EpisodeMutation
}

// Where appends a list predicates to the EpisodeUpdate builder.
func (eu *EpisodeUpdate) Where(ps ...predicate.Episode) *EpisodeUpdate {
	eu.mutation.Where(ps...)
	return eu
}

// SetName sets the "name" field.
func (eu *EpisodeUpdate) SetName(s string) *EpisodeUpdate {
	eu.mutation.SetName(s)
	return eu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (eu *EpisodeUpdate) SetNillableName(s *string) *EpisodeUpdate {
	if s != nil {
		eu.SetName(*s)
	}
	return eu
}

// SetValue sets the "value" field.
func (eu *EpisodeUpdate) SetValue(i int64) *EpisodeUpdate {
	eu.mutation.ResetValue()
	eu.mutation.SetValue(i)
	return eu
}

// SetNillableValue sets the "value" field if the given value is not nil.
func (eu *EpisodeUpdate) SetNillableValue(i *int64) *EpisodeUpdate {
	if i != nil {
		eu.SetValue(*i)
	}
	return eu
}

// AddValue adds i to the "value" field.
func (eu *EpisodeUpdate) AddValue(i int64) *EpisodeUpdate {
	eu.mutation.AddValue(i)
	return eu
}

// SetDescription sets the "description" field.
func (eu *EpisodeUpdate) SetDescription(s string) *EpisodeUpdate {
	eu.mutation.SetDescription(s)
	return eu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (eu *EpisodeUpdate) SetNillableDescription(s *string) *EpisodeUpdate {
	if s != nil {
		eu.SetDescription(*s)
	}
	return eu
}

// SetStatus sets the "status" field.
func (eu *EpisodeUpdate) SetStatus(i int) *EpisodeUpdate {
	eu.mutation.ResetStatus()
	eu.mutation.SetStatus(i)
	return eu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (eu *EpisodeUpdate) SetNillableStatus(i *int) *EpisodeUpdate {
	if i != nil {
		eu.SetStatus(*i)
	}
	return eu
}

// AddStatus adds i to the "status" field.
func (eu *EpisodeUpdate) AddStatus(i int) *EpisodeUpdate {
	eu.mutation.AddStatus(i)
	return eu
}

// SetCreatedAt sets the "created_at" field.
func (eu *EpisodeUpdate) SetCreatedAt(t time.Time) *EpisodeUpdate {
	eu.mutation.SetCreatedAt(t)
	return eu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (eu *EpisodeUpdate) SetNillableCreatedAt(t *time.Time) *EpisodeUpdate {
	if t != nil {
		eu.SetCreatedAt(*t)
	}
	return eu
}

// SetSeasonID sets the "season" edge to the Season entity by ID.
func (eu *EpisodeUpdate) SetSeasonID(id int) *EpisodeUpdate {
	eu.mutation.SetSeasonID(id)
	return eu
}

// SetNillableSeasonID sets the "season" edge to the Season entity by ID if the given value is not nil.
func (eu *EpisodeUpdate) SetNillableSeasonID(id *int) *EpisodeUpdate {
	if id != nil {
		eu = eu.SetSeasonID(*id)
	}
	return eu
}

// SetSeason sets the "season" edge to the Season entity.
func (eu *EpisodeUpdate) SetSeason(s *Season) *EpisodeUpdate {
	return eu.SetSeasonID(s.ID)
}

// Mutation returns the EpisodeMutation object of the builder.
func (eu *EpisodeUpdate) Mutation() *EpisodeMutation {
	return eu.mutation
}

// ClearSeason clears the "season" edge to the Season entity.
func (eu *EpisodeUpdate) ClearSeason() *EpisodeUpdate {
	eu.mutation.ClearSeason()
	return eu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (eu *EpisodeUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, eu.sqlSave, eu.mutation, eu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (eu *EpisodeUpdate) SaveX(ctx context.Context) int {
	affected, err := eu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (eu *EpisodeUpdate) Exec(ctx context.Context) error {
	_, err := eu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eu *EpisodeUpdate) ExecX(ctx context.Context) {
	if err := eu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (eu *EpisodeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(episode.Table, episode.Columns, sqlgraph.NewFieldSpec(episode.FieldID, field.TypeInt))
	if ps := eu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := eu.mutation.Name(); ok {
		_spec.SetField(episode.FieldName, field.TypeString, value)
	}
	if value, ok := eu.mutation.Value(); ok {
		_spec.SetField(episode.FieldValue, field.TypeInt64, value)
	}
	if value, ok := eu.mutation.AddedValue(); ok {
		_spec.AddField(episode.FieldValue, field.TypeInt64, value)
	}
	if value, ok := eu.mutation.Description(); ok {
		_spec.SetField(episode.FieldDescription, field.TypeString, value)
	}
	if value, ok := eu.mutation.Status(); ok {
		_spec.SetField(episode.FieldStatus, field.TypeInt, value)
	}
	if value, ok := eu.mutation.AddedStatus(); ok {
		_spec.AddField(episode.FieldStatus, field.TypeInt, value)
	}
	if value, ok := eu.mutation.CreatedAt(); ok {
		_spec.SetField(episode.FieldCreatedAt, field.TypeTime, value)
	}
	if eu.mutation.SeasonCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   episode.SeasonTable,
			Columns: []string{episode.SeasonColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(season.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.SeasonIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   episode.SeasonTable,
			Columns: []string{episode.SeasonColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(season.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, eu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{episode.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	eu.mutation.done = true
	return n, nil
}

// EpisodeUpdateOne is the builder for updating a single Episode entity.
type EpisodeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *EpisodeMutation
}

// SetName sets the "name" field.
func (euo *EpisodeUpdateOne) SetName(s string) *EpisodeUpdateOne {
	euo.mutation.SetName(s)
	return euo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (euo *EpisodeUpdateOne) SetNillableName(s *string) *EpisodeUpdateOne {
	if s != nil {
		euo.SetName(*s)
	}
	return euo
}

// SetValue sets the "value" field.
func (euo *EpisodeUpdateOne) SetValue(i int64) *EpisodeUpdateOne {
	euo.mutation.ResetValue()
	euo.mutation.SetValue(i)
	return euo
}

// SetNillableValue sets the "value" field if the given value is not nil.
func (euo *EpisodeUpdateOne) SetNillableValue(i *int64) *EpisodeUpdateOne {
	if i != nil {
		euo.SetValue(*i)
	}
	return euo
}

// AddValue adds i to the "value" field.
func (euo *EpisodeUpdateOne) AddValue(i int64) *EpisodeUpdateOne {
	euo.mutation.AddValue(i)
	return euo
}

// SetDescription sets the "description" field.
func (euo *EpisodeUpdateOne) SetDescription(s string) *EpisodeUpdateOne {
	euo.mutation.SetDescription(s)
	return euo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (euo *EpisodeUpdateOne) SetNillableDescription(s *string) *EpisodeUpdateOne {
	if s != nil {
		euo.SetDescription(*s)
	}
	return euo
}

// SetStatus sets the "status" field.
func (euo *EpisodeUpdateOne) SetStatus(i int) *EpisodeUpdateOne {
	euo.mutation.ResetStatus()
	euo.mutation.SetStatus(i)
	return euo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (euo *EpisodeUpdateOne) SetNillableStatus(i *int) *EpisodeUpdateOne {
	if i != nil {
		euo.SetStatus(*i)
	}
	return euo
}

// AddStatus adds i to the "status" field.
func (euo *EpisodeUpdateOne) AddStatus(i int) *EpisodeUpdateOne {
	euo.mutation.AddStatus(i)
	return euo
}

// SetCreatedAt sets the "created_at" field.
func (euo *EpisodeUpdateOne) SetCreatedAt(t time.Time) *EpisodeUpdateOne {
	euo.mutation.SetCreatedAt(t)
	return euo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (euo *EpisodeUpdateOne) SetNillableCreatedAt(t *time.Time) *EpisodeUpdateOne {
	if t != nil {
		euo.SetCreatedAt(*t)
	}
	return euo
}

// SetSeasonID sets the "season" edge to the Season entity by ID.
func (euo *EpisodeUpdateOne) SetSeasonID(id int) *EpisodeUpdateOne {
	euo.mutation.SetSeasonID(id)
	return euo
}

// SetNillableSeasonID sets the "season" edge to the Season entity by ID if the given value is not nil.
func (euo *EpisodeUpdateOne) SetNillableSeasonID(id *int) *EpisodeUpdateOne {
	if id != nil {
		euo = euo.SetSeasonID(*id)
	}
	return euo
}

// SetSeason sets the "season" edge to the Season entity.
func (euo *EpisodeUpdateOne) SetSeason(s *Season) *EpisodeUpdateOne {
	return euo.SetSeasonID(s.ID)
}

// Mutation returns the EpisodeMutation object of the builder.
func (euo *EpisodeUpdateOne) Mutation() *EpisodeMutation {
	return euo.mutation
}

// ClearSeason clears the "season" edge to the Season entity.
func (euo *EpisodeUpdateOne) ClearSeason() *EpisodeUpdateOne {
	euo.mutation.ClearSeason()
	return euo
}

// Where appends a list predicates to the EpisodeUpdate builder.
func (euo *EpisodeUpdateOne) Where(ps ...predicate.Episode) *EpisodeUpdateOne {
	euo.mutation.Where(ps...)
	return euo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (euo *EpisodeUpdateOne) Select(field string, fields ...string) *EpisodeUpdateOne {
	euo.fields = append([]string{field}, fields...)
	return euo
}

// Save executes the query and returns the updated Episode entity.
func (euo *EpisodeUpdateOne) Save(ctx context.Context) (*Episode, error) {
	return withHooks(ctx, euo.sqlSave, euo.mutation, euo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (euo *EpisodeUpdateOne) SaveX(ctx context.Context) *Episode {
	node, err := euo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (euo *EpisodeUpdateOne) Exec(ctx context.Context) error {
	_, err := euo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (euo *EpisodeUpdateOne) ExecX(ctx context.Context) {
	if err := euo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (euo *EpisodeUpdateOne) sqlSave(ctx context.Context) (_node *Episode, err error) {
	_spec := sqlgraph.NewUpdateSpec(episode.Table, episode.Columns, sqlgraph.NewFieldSpec(episode.FieldID, field.TypeInt))
	id, ok := euo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Episode.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := euo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, episode.FieldID)
		for _, f := range fields {
			if !episode.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != episode.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := euo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := euo.mutation.Name(); ok {
		_spec.SetField(episode.FieldName, field.TypeString, value)
	}
	if value, ok := euo.mutation.Value(); ok {
		_spec.SetField(episode.FieldValue, field.TypeInt64, value)
	}
	if value, ok := euo.mutation.AddedValue(); ok {
		_spec.AddField(episode.FieldValue, field.TypeInt64, value)
	}
	if value, ok := euo.mutation.Description(); ok {
		_spec.SetField(episode.FieldDescription, field.TypeString, value)
	}
	if value, ok := euo.mutation.Status(); ok {
		_spec.SetField(episode.FieldStatus, field.TypeInt, value)
	}
	if value, ok := euo.mutation.AddedStatus(); ok {
		_spec.AddField(episode.FieldStatus, field.TypeInt, value)
	}
	if value, ok := euo.mutation.CreatedAt(); ok {
		_spec.SetField(episode.FieldCreatedAt, field.TypeTime, value)
	}
	if euo.mutation.SeasonCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   episode.SeasonTable,
			Columns: []string{episode.SeasonColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(season.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.SeasonIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   episode.SeasonTable,
			Columns: []string{episode.SeasonColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(season.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Episode{config: euo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, euo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{episode.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	euo.mutation.done = true
	return _node, nil
}