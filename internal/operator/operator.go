package operator

import (
	"github.com/dongwlin/anime-list/internal/ent"
)

type Operator struct {
	Anime   AnimeOperator
	Season  SeasonOperator
	Episode EpisodeOperator
	Theater TheaterOperator
}

func New(db *ent.Client) *Operator {
	return &Operator{
		Anime:   NewAnimeOperator(db),
		Season:  NewSeasonOperator(db),
		Episode: NewEpisodeOperator(db),
		Theater: NewTheaterOperator(db),
	}
}
