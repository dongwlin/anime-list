// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"time"
)

type Anime struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Desc      string    `json:"desc"`
	Status    int64     `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Episode struct {
	ID        int64     `json:"id"`
	SeasonID  int64     `json:"season_id"`
	Name      string    `json:"name"`
	Value     int64     `json:"value"`
	Desc      string    `json:"desc"`
	Status    int64     `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Season struct {
	ID         int64     `json:"id"`
	AnimeID    int64     `json:"anime_id"`
	Name       string    `json:"name"`
	Value      int64     `json:"value"`
	Cover      string    `json:"cover"`
	ReleasedAt time.Time `json:"released_at"`
	Desc       string    `json:"desc"`
	Status     int64     `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type Theater struct {
	ID         int64     `json:"id"`
	AnimeID    int64     `json:"anime_id"`
	Name       string    `json:"name"`
	Cover      string    `json:"cover"`
	ReleasedAt time.Time `json:"released_at"`
	Desc       string    `json:"desc"`
	Status     int64     `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
