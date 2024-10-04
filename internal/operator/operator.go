package operator

import (
	"context"

	"github.com/dongwlin/anime-list/internal/ent"
)

type EpisodeOperator interface {
	Create(ctx context.Context)
	List(ctx context.Context)
	Get(ctx context.Context)
	Update(ctx context.Context)
	Delete(ctx context.Context)
}

type TheaterOperator interface {
	Create(ctx context.Context)
	List(ctx context.Context)
	Get(ctx context.Context)
	Update(ctx context.Context)
	Delete(ctx context.Context)
}

type Operator struct {
	Anime   AnimeOperator
	Season  SeasonOperator
	Episode EpisodeOperator
	Theater TheaterOperator
}

func New(db *ent.Client) *Operator {
	return &Operator{
		Anime:  NewAnimeOperator(db),
		Season: NewSeasonOperator(db),
	}
}
