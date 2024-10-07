package operator

import (
	"context"
	"github.com/dongwlin/anime-list/internal/ent/episode"
	"github.com/dongwlin/anime-list/internal/ent/season"
	"github.com/dongwlin/anime-list/internal/ent/theater"

	"github.com/dongwlin/anime-list/internal/ent"
	"github.com/dongwlin/anime-list/internal/ent/anime"
)

type CreateAnimeParams struct {
	Name        string
	Description string
	Status      int
}

type ListAnimeParams struct {
	Offset int
	Limit  int
}

type SearchAnimeParams struct {
	Query  string
	Offset int
	Limit  int
}

type UpdateAnimeParams struct {
	ID          int
	Name        string
	Description string
	Status      int
}

type AnimeOperator interface {
	Create(ctx context.Context, params CreateAnimeParams) (*ent.Anime, error)
	List(ctx context.Context, params ListAnimeParams) (int, []*ent.Anime, error)
	Search(ctx context.Context, params SearchAnimeParams) (int, []*ent.Anime, error)
	Get(ctx context.Context, id int) (*ent.Anime, error)
	Update(ctx context.Context, params UpdateAnimeParams) (*ent.Anime, error)
	Delete(ctx context.Context, id int) error
}

type animeOperator struct {
	db *ent.Client
}

// Create implements AnimeOperator.
func (ao *animeOperator) Create(ctx context.Context, params CreateAnimeParams) (*ent.Anime, error) {
	a, err := ao.db.Anime.Create().
		SetName(params.Name).
		SetDescription(params.Description).
		SetStatus(params.Status).
		Save(ctx)
	return a, err
}

// List implements AnimeOperator.
func (ao *animeOperator) List(ctx context.Context, params ListAnimeParams) (int, []*ent.Anime, error) {
	query := ao.db.Anime.Query()

	count, err := query.Count(ctx)
	if err != nil {
		return 0, nil, err
	}

	animes, err := query.
		Offset(params.Offset).
		Limit(params.Limit).
		All(ctx)
	if err != nil {
		return 0, nil, err
	}

	return count, animes, nil
}

func (ao *animeOperator) Search(ctx context.Context, params SearchAnimeParams) (int, []*ent.Anime, error) {
	query := ao.db.Anime.Query().
		Where(anime.NameContains(params.Query))

	count, err := query.Count(ctx)
	if err != nil {
		return 0, nil, err
	}

	animes, err := query.
		Offset(params.Offset).
		Limit(params.Limit).
		All(ctx)
	if err != nil {
		return 0, nil, err
	}

	return count, animes, nil
}

// Get implements AnimeOperator.
func (ao *animeOperator) Get(ctx context.Context, id int) (*ent.Anime, error) {
	a, err := ao.db.Anime.
		Get(ctx, id)
	return a, err
}

// Update implements AnimeOperator.
func (ao *animeOperator) Update(ctx context.Context, params UpdateAnimeParams) (*ent.Anime, error) {
	a, err := ao.db.Anime.
		UpdateOneID(params.ID).
		SetName(params.Name).
		SetDescription(params.Description).
		SetStatus(params.Status).
		Save(ctx)
	return a, err
}

// Delete implements AnimeOperator.
func (ao *animeOperator) Delete(ctx context.Context, id int) error {
	tx, err := ao.db.Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	_, err = tx.Theater.
		Delete().
		Where(theater.HasAnimeWith(anime.ID(id))).
		Exec(ctx)
	if err != nil {
		return err
	}

	seasons, err := tx.Season.
		Query().
		Where(season.HasAnimeWith(anime.ID(id))).
		All(ctx)
	if err != nil {
		return err
	}

	for _, s := range seasons {
		_, err = tx.Episode.
			Delete().
			Where(episode.HasSeasonWith(season.ID(s.ID))).
			Exec(ctx)
		if err != nil {
			return err
		}

		err = tx.Season.
			DeleteOneID(s.ID).
			Exec(ctx)
		if err != nil {
			return err
		}
	}

	err = tx.Anime.
		DeleteOneID(id).
		Exec(ctx)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func NewAnimeOperator(db *ent.Client) AnimeOperator {
	return &animeOperator{
		db: db,
	}
}