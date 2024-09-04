package db

import (
	"context"
	"database/sql"
	"github.com/dongwlin/anime-list/internal/pkg/util"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func createRandomTheater(t *testing.T, anime Anime) Theater {
	arg := CreateTheaterParams{
		AnimeID:     anime.ID,
		Name:        util.RandomString(6),
		Cover:       util.RandomString(255),
		ReleasedAt:  time.Now(),
		Description: util.RandomString(16),
		Status:      0,
	}

	theater, err := testQueries.CreateTheater(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, theater)

	require.Equal(t, arg.AnimeID, theater.AnimeID)
	require.Equal(t, arg.Name, theater.Name)
	require.Equal(t, arg.Cover, theater.Cover)
	require.WithinDuration(t, arg.ReleasedAt, theater.ReleasedAt, time.Second)
	require.Equal(t, arg.Description, theater.Description)
	require.Equal(t, arg.Status, theater.Status)
	require.NotZero(t, theater.CreatedAt)
	require.NotZero(t, theater.UpdatedAt)

	return theater
}

func TestQueries_CreateTheater(t *testing.T) {
	anime := createRandomAnime(t)
	createRandomTheater(t, anime)
}

func TestQueries_GetTheater(t *testing.T) {
	anime := createRandomAnime(t)
	theater1 := createRandomTheater(t, anime)

	theater2, err := testQueries.GetTheater(context.Background(), theater1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, theater2)

	require.Equal(t, theater1.ID, theater2.ID)
	require.Equal(t, theater1.AnimeID, theater2.AnimeID)
	require.Equal(t, theater1.Name, theater2.Name)
	require.Equal(t, theater1.Cover, theater2.Cover)
	require.WithinDuration(t, theater1.ReleasedAt, theater2.ReleasedAt, time.Second)
	require.Equal(t, theater1.Description, theater2.Description)
	require.Equal(t, theater1.Status, theater2.Status)
	require.WithinDuration(t, theater1.CreatedAt, theater2.CreatedAt, time.Second)
	require.WithinDuration(t, theater1.UpdatedAt, theater2.UpdatedAt, time.Second)
}

func TestQueries_ListTheater(t *testing.T) {
	anime := createRandomAnime(t)
	for i := 0; i < 10; i++ {
		createRandomTheater(t, anime)
	}

	arg := ListTheaterParams{
		Limit:  5,
		Offset: 5,
	}

	theaters, err := testQueries.ListTheater(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, theaters, 5)

	for _, theater := range theaters {
		require.NotEmpty(t, theater)
	}
}

func TestQueries_UpdateTheater(t *testing.T) {
	anime := createRandomAnime(t)
	theater1 := createRandomTheater(t, anime)

	arg := UpdateTheaterParams{
		ID:         theater1.ID,
		Name:       util.RandomString(6),
		Cover:      util.RandomString(255),
		ReleasedAt: time.Now(),
	}

	theater2, err := testQueries.UpdateTheater(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, theater2)

	require.Equal(t, theater1.ID, theater2.ID)
	require.Equal(t, theater1.AnimeID, theater2.AnimeID)
	require.Equal(t, arg.Name, theater2.Name)
	require.Equal(t, arg.Cover, theater2.Cover)
	require.WithinDuration(t, arg.ReleasedAt, theater2.ReleasedAt, time.Second)
	require.Equal(t, arg.Description, theater2.Description)
	require.Equal(t, arg.Status, theater2.Status)
	require.WithinDuration(t, theater1.CreatedAt, theater2.CreatedAt, time.Second)
	//require.NotEqual(t, theater1.UpdatedAt, theater2.UpdatedAt)
}

func TestQueries_DeleteTheater(t *testing.T) {
	anime := createRandomAnime(t)
	theater1 := createRandomTheater(t, anime)

	err := testQueries.DeleteTheater(context.Background(), theater1.ID)
	require.NoError(t, err)

	theater2, err := testQueries.GetTheater(context.Background(), theater1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, theater2)
}
