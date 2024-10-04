package handler

import "github.com/dongwlin/anime-list/internal/operator"

type Handler struct {
	Ping    *PingHandler
	Anime   *AnimeHandler
	Season  *SeasonHandler
	Episode *EpisodeHandler
	Theater *TheaterHandler
}

func New(operator *operator.Operator) *Handler {
	return &Handler{
		Ping:    NewPingHandler(),
		Anime:   NewAnimeHandler(operator.Anime),
		Season:  NewSeasonHandler(operator.Season),
		Episode: NewEpisodeHandler(operator.Episode),
		Theater: NewTheaterHandler(operator.Theater),
	}
}
