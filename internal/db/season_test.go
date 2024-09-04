package db

import (
	"context"
	"database/sql"
	"github.com/dongwlin/anime-list/internal/pkg/util"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func createRandomSeason(t *testing.T, anime Anime) Season {
	arg := CreateSeasonParams{
		AnimeID:     anime.ID,
		Name:        util.RandomString(6),
		Value:       util.RandomInt64(1, 32),
		Cover:       util.RandomString(255),
		ReleasedAt:  time.Now(),
		Description: util.RandomString(16),
		Status:      util.RandomInt64(1, 7),
	}

	season, err := testQueries.CreateSeason(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, season)

	require.Equal(t, arg.AnimeID, season.AnimeID)
	require.Equal(t, arg.Name, season.Name)
	require.Equal(t, arg.Value, season.Value)
	require.Equal(t, arg.Cover, season.Cover)
	require.WithinDuration(t, arg.ReleasedAt, season.ReleasedAt, time.Second)
	require.Equal(t, arg.Description, season.Description)
	require.Equal(t, arg.Status, season.Status)
	require.NotZero(t, season.CreatedAt)
	require.NotZero(t, season.UpdatedAt)

	return season
}

func TestQueries_CreateSeason(t *testing.T) {
	anime := createRandomAnime(t)
	createRandomSeason(t, anime)
}

func TestQueries_GetSeason(t *testing.T) {
	anime := createRandomAnime(t)
	season1 := createRandomSeason(t, anime)

	season2, err := testQueries.GetSeason(context.Background(), season1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, season2)

	require.Equal(t, season1.ID, season2.ID)
	require.Equal(t, season1.AnimeID, season2.AnimeID)
	require.Equal(t, season1.Name, season2.Name)
	require.Equal(t, season1.Value, season2.Value)
	require.Equal(t, season1.Cover, season2.Cover)
	require.WithinDuration(t, season1.ReleasedAt, season2.ReleasedAt, time.Second)
	require.Equal(t, season1.Description, season2.Description)
	require.Equal(t, season1.Status, season2.Status)
	require.WithinDuration(t, season1.CreatedAt, season2.CreatedAt, time.Second)
	require.WithinDuration(t, season1.UpdatedAt, season2.UpdatedAt, time.Second)
}

func TestQueries_ListSeason(t *testing.T) {
	anime := createRandomAnime(t)
	for i := 0; i < 10; i++ {
		createRandomSeason(t, anime)
	}

	arg := ListSeasonParams{
		AnimeID: anime.ID,
		Limit:   5,
		Offset:  5,
	}

	seasons, err := testQueries.ListSeason(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, seasons, 5)

	for _, season := range seasons {
		require.NotEmpty(t, season)
	}
}

func TestQueries_UpdateSeason(t *testing.T) {
	anime := createRandomAnime(t)
	season1 := createRandomSeason(t, anime)

	arg := UpdateSeasonParams{
		ID:          season1.ID,
		Name:        util.RandomString(6),
		Value:       util.RandomInt64(1, 32),
		Cover:       util.RandomString(255),
		ReleasedAt:  time.Now(),
		Description: util.RandomString(16),
		Status:      util.RandomInt64(1, 7),
	}

	season2, err := testQueries.UpdateSeason(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, season2)

	require.Equal(t, season1.ID, season2.ID)
	require.Equal(t, season1.AnimeID, season2.AnimeID)
	require.Equal(t, arg.Name, season2.Name)
	require.Equal(t, arg.Value, season2.Value)
	require.Equal(t, arg.Cover, season2.Cover)
	require.WithinDuration(t, arg.ReleasedAt, season2.ReleasedAt, time.Second)
	require.Equal(t, arg.Description, season2.Description)
	require.Equal(t, arg.Status, season2.Status)
	require.WithinDuration(t, season1.CreatedAt, season2.CreatedAt, time.Second)
	//require.NotEqual(t, season1.UpdatedAt, season2.UpdatedAt)
}

func TestQueries_DeleteSeason(t *testing.T) {
	anime := createRandomAnime(t)
	season1 := createRandomSeason(t, anime)

	err := testQueries.DeleteSeason(context.Background(), season1.ID)
	require.NoError(t, err)

	season2, err := testQueries.GetSeason(context.Background(), season1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, season2)
}
