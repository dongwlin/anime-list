package db

import (
	"context"
	"database/sql"
	"github.com/dongwlin/anime-list/internal/pkg/util"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func createRandomAnime(t *testing.T) Anime {

	arg := CreateAnimeParams{
		Name:        util.RandomString(6),
		Description: util.RandomString(16),
		Status:      0,
	}

	anime, err := testQueries.CreateAnime(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, anime)

	require.Equal(t, arg.Name, anime.Name)
	require.Equal(t, arg.Description, anime.Description)
	require.Equal(t, arg.Status, anime.Status)
	require.NotZero(t, anime.CreatedAt)
	require.NotZero(t, anime.UpdatedAt)

	return anime
}

func TestQueries_CreateAnime(t *testing.T) {
	createRandomAnime(t)
}

func TestQueries_GetAnime(t *testing.T) {
	anime1 := createRandomAnime(t)
	anime2, err := testQueries.GetAnime(context.Background(), anime1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, anime2)

	require.Equal(t, anime1.ID, anime2.ID)
	require.Equal(t, anime1.Name, anime2.Name)
	require.Equal(t, anime1.Description, anime2.Description)
	require.Equal(t, anime1.Status, anime2.Status)
	require.WithinDuration(t, anime1.CreatedAt, anime2.CreatedAt, time.Second)
	require.WithinDuration(t, anime1.UpdatedAt, anime2.UpdatedAt, time.Second)
}

func TestQueries_ListAnime(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAnime(t)
	}

	arg := ListAnimeParams{
		Limit:  5,
		Offset: 5,
	}

	animes, err := testQueries.ListAnime(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, animes, 5)

	for _, anime := range animes {
		require.NotEmpty(t, anime)
	}
}

func TestQueries_UpdateAnime(t *testing.T) {
	anime1 := createRandomAnime(t)

	arg := UpdateAnimeParams{
		ID:          anime1.ID,
		Name:        util.RandomString(6),
		Description: util.RandomString(16),
		Status:      util.RandomInt64(1, 7),
	}

	anime2, err := testQueries.UpdateAnime(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, anime2)

	require.Equal(t, anime1.ID, anime2.ID)
	require.Equal(t, arg.Name, anime2.Name)
	require.Equal(t, arg.Description, anime2.Description)
	require.Equal(t, arg.Status, anime2.Status)
	require.WithinDuration(t, anime1.CreatedAt, anime2.CreatedAt, time.Second)
	//require.NotEqual(t, anime1.UpdatedAt, anime2.UpdatedAt)
}

func TestQueries_DeleteAnime(t *testing.T) {
	anime1 := createRandomAnime(t)

	err := testQueries.DeleteAnime(context.Background(), anime1.ID)
	require.NoError(t, err)

	anime2, err := testQueries.GetAnime(context.Background(), anime1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, anime2)
}
