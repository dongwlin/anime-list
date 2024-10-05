package operator

import (
	"context"
	"time"

	"github.com/dongwlin/anime-list/internal/ent"
	"github.com/dongwlin/anime-list/internal/ent/anime"
	"github.com/dongwlin/anime-list/internal/ent/theater"
)

type CreateTheaterParams struct {
	AnimeID     int
	Name        string
	Cover       string
	ReleasedAt  time.Time
	Description string
	Status      int
}

type ListTheaterParams struct {
	AnimeID int
	Offset  int
	Limit   int
}

type SearchTheaterParams struct {
	AnimeID int
	Query   string
	Offset  int
	Limit   int
}

type UpdateTheaterParams struct {
	ID          int
	Name        string
	Cover       string
	ReleasedAt  time.Time
	Description string
	Status      int
}

type TheaterOperator interface {
	Create(ctx context.Context, params CreateTheaterParams) (*ent.Theater, error)
	List(ctx context.Context, params ListTheaterParams) (int, []*ent.Theater, error)
	Search(ctx context.Context, params SearchTheaterParams) (int, []*ent.Theater, error)
	Get(ctx context.Context, id int) (*ent.Theater, error)
	Update(ctx context.Context, params UpdateTheaterParams) (*ent.Theater, error)
	Delete(ctx context.Context, id int) error
}

type theaterOperator struct {
	db *ent.Client
}

// Create implements TheaterOperator.
func (to *theaterOperator) Create(ctx context.Context, params CreateTheaterParams) (*ent.Theater, error) {
	t, err := to.db.Theater.
		Create().
		SetName(params.Name).
		SetCover(params.Cover).
		SetReleasedAt(params.ReleasedAt).
		SetDescription(params.Description).
		SetStatus(params.Status).
		SetAnimeID(params.AnimeID).
		Save(ctx)
	return t, err
}

// List implements TheaterOperator.
func (to *theaterOperator) List(ctx context.Context, params ListTheaterParams) (int, []*ent.Theater, error) {
	count, err := to.db.Anime.
		Query().
		Where(anime.ID(params.AnimeID)).
		QueryTheaters().
		Count(ctx)
	if err != nil {
		return 0, nil, err
	}
	theaters, err := to.db.Anime.
		Query().
		Where(anime.ID(params.AnimeID)).
		QueryTheaters().
		Offset(params.Offset).
		Limit(params.Limit).
		All(ctx)
	if err != nil {
		return 0, nil, err
	}
	return count, theaters, nil
}

// Search implements TheaterOperator.
func (to *theaterOperator) Search(ctx context.Context, params SearchTheaterParams) (int, []*ent.Theater, error) {
	count, err := to.db.Anime.
		Query().
		Where(anime.ID(params.AnimeID)).
		QueryTheaters().
		Where(theater.NameContains(params.Query)).
		Count(ctx)
	if err != nil {
		return 0, nil, err
	}
	theaters, err := to.db.Anime.
		Query().
		Where(anime.ID(params.AnimeID)).
		QueryTheaters().
		Where(theater.NameContains(params.Query)).
		Offset(params.Offset).
		Limit(params.Limit).
		All(ctx)
	if err != nil {
		return 0, nil, err
	}
	return count, theaters, nil
}

// Get implements TheaterOperator.
func (to *theaterOperator) Get(ctx context.Context, id int) (*ent.Theater, error) {
	t, err := to.db.Theater.
		Get(ctx, id)
	return t, err
}

// Update implements TheaterOperator.
func (to *theaterOperator) Update(ctx context.Context, params UpdateTheaterParams) (*ent.Theater, error) {
	t, err := to.db.Theater.
		UpdateOneID(params.ID).
		SetName(params.Name).
		SetCover(params.Cover).
		SetReleasedAt(params.ReleasedAt).
		SetDescription(params.Description).
		SetStatus(params.Status).
		Save(ctx)
	return t, err
}

// Delete implements TheaterOperator.
func (to *theaterOperator) Delete(ctx context.Context, id int) error {
	err := to.db.Theater.
		DeleteOneID(id).
		Exec(ctx)
	return err
}

func NewTheaterOperator(db *ent.Client) TheaterOperator {
	return &theaterOperator{
		db: db,
	}
}
