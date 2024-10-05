package operator

import (
	"context"

	"github.com/dongwlin/anime-list/internal/ent"
	"github.com/dongwlin/anime-list/internal/ent/episode"
	"github.com/dongwlin/anime-list/internal/ent/season"
)

type CreateEpisodeParams struct {
	SeasonID    int
	Name        string
	Value       int64
	Description string
	Status      int
}

type ListEpisodeParams struct {
	SeasonID int
	Offset   int
	Limit    int
}

type SearchEpisodeParams struct {
	SeasonID int
	Query    string
	Offset   int
	Limit    int
}

type UpdateEpisodeParams struct {
	ID          int
	Name        string
	Value       int64
	Description string
	Status      int
}

type EpisodeOperator interface {
	Create(ctx context.Context, params CreateEpisodeParams) (*ent.Episode, error)
	List(ctx context.Context, params ListEpisodeParams) (int, []*ent.Episode, error)
	Search(ctx context.Context, params SearchEpisodeParams) (int, []*ent.Episode, error)
	Get(ctx context.Context, id int) (*ent.Episode, error)
	Update(ctx context.Context, params UpdateEpisodeParams) (*ent.Episode, error)
	Delete(ctx context.Context, id int) error
}

type episodeOperator struct {
	db *ent.Client
}

// Create implements EpisodeOperator.
func (eo *episodeOperator) Create(ctx context.Context, params CreateEpisodeParams) (*ent.Episode, error) {
	e, err := eo.db.Episode.
		Create().
		SetName(params.Name).
		SetValue(params.Value).
		SetDescription(params.Description).
		SetStatus(params.Status).
		SetSeasonID(params.SeasonID).
		Save(ctx)
	return e, err

}

// List implements EpisodeOperator.
func (eo *episodeOperator) List(ctx context.Context, params ListEpisodeParams) (int, []*ent.Episode, error) {
	count, err := eo.db.Season.
		Query().
		Where(season.ID(params.SeasonID)).
		QueryEpisodes().
		Count(ctx)
	if err != nil {
		return 0, nil, err
	}
	episodes, err := eo.db.Season.
		Query().
		Where(season.ID(params.SeasonID)).
		QueryEpisodes().
		Offset(params.Offset).
		Limit(params.Limit).
		All(ctx)
	if err != nil {
		return 0, nil, err
	}
	return count, episodes, nil
}

// Search implements EpisodeOperator.
func (eo *episodeOperator) Search(ctx context.Context, params SearchEpisodeParams) (int, []*ent.Episode, error) {
	count, err := eo.db.Season.
		Query().
		Where(season.ID(params.SeasonID)).
		QueryEpisodes().
		Where(episode.NameContains(params.Query)).
		Count(ctx)
	if err != nil {
		return 0, nil, err
	}
	episodes, err := eo.db.Season.
		Query().
		Where(season.ID(params.SeasonID)).
		QueryEpisodes().
		Where(episode.NameContains(params.Query)).
		Offset(params.Offset).
		Limit(params.Limit).
		All(ctx)
	if err != nil {
		return 0, nil, err
	}
	return count, episodes, nil
}

// Get implements EpisodeOperator.
func (eo *episodeOperator) Get(ctx context.Context, id int) (*ent.Episode, error) {
	e, err := eo.db.Episode.
		Get(ctx, id)
	return e, err
}

// Update implements EpisodeOperator.
func (eo *episodeOperator) Update(ctx context.Context, params UpdateEpisodeParams) (*ent.Episode, error) {
	e, err := eo.db.Episode.
		UpdateOneID(params.ID).
		SetName(params.Name).
		SetValue(params.Value).
		SetDescription(params.Description).
		SetStatus(params.Status).
		Save(ctx)
	return e, err
}

// Delete implements EpisodeOperator.
func (eo *episodeOperator) Delete(ctx context.Context, id int) error {
	err := eo.db.Episode.
		DeleteOneID(id).
		Exec(ctx)
	return err
}

func NewEpisodeOperator(db *ent.Client) EpisodeOperator {
	return &episodeOperator{
		db: db,
	}
}
