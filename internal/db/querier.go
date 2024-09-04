// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"context"
)

type Querier interface {
	CountAnime(ctx context.Context) (int64, error)
	CountSeason(ctx context.Context, animeID int64) (int64, error)
	CreateAnime(ctx context.Context, arg CreateAnimeParams) (Anime, error)
	CreateEpisode(ctx context.Context, arg CreateEpisodeParams) (Episode, error)
	CreateSeason(ctx context.Context, arg CreateSeasonParams) (Season, error)
	CreateTheater(ctx context.Context, arg CreateTheaterParams) (Theater, error)
	DeleteAnime(ctx context.Context, id int64) error
	DeleteEpisode(ctx context.Context, id int64) error
	DeleteSeason(ctx context.Context, id int64) error
	DeleteTheater(ctx context.Context, id int64) error
	GetAnime(ctx context.Context, id int64) (Anime, error)
	GetEpisode(ctx context.Context, id int64) (Episode, error)
	GetSeason(ctx context.Context, id int64) (Season, error)
	GetTheater(ctx context.Context, id int64) (Theater, error)
	ListAnime(ctx context.Context, arg ListAnimeParams) ([]Anime, error)
	ListEpisode(ctx context.Context, arg ListEpisodeParams) ([]Episode, error)
	ListSeason(ctx context.Context, arg ListSeasonParams) ([]Season, error)
	ListTheater(ctx context.Context, arg ListTheaterParams) ([]Theater, error)
	UpdateAnime(ctx context.Context, arg UpdateAnimeParams) (Anime, error)
	UpdateEpisode(ctx context.Context, arg UpdateEpisodeParams) (Episode, error)
	UpdateSeason(ctx context.Context, arg UpdateSeasonParams) (Season, error)
	UpdateTheater(ctx context.Context, arg UpdateTheaterParams) (Theater, error)
}

var _ Querier = (*Queries)(nil)
