package db

import (
	"context"
	"database/sql"
	"github.com/dongwlin/anime-list/internal/pkg/util"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func createRandomEpisode(t *testing.T, season Season) Episode {
	arg := CreateEpisodeParams{
		SeasonID:    season.ID,
		Name:        util.RandomString(6),
		Value:       util.RandomInt64(1, 225),
		Description: util.RandomString(16),
		Status:      0,
	}

	episode, err := testQueries.CreateEpisode(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, episode)

	require.Equal(t, arg.SeasonID, episode.SeasonID)
	require.Equal(t, arg.Name, episode.Name)
	require.Equal(t, arg.Value, episode.Value)
	require.Equal(t, arg.Description, episode.Description)
	require.Equal(t, arg.Status, episode.Status)
	require.NotZero(t, episode.CreatedAt)
	require.NotZero(t, episode.UpdatedAt)

	return episode
}

func TestQueries_CreateEpisode(t *testing.T) {
	anime := createRandomAnime(t)
	season := createRandomSeason(t, anime)
	createRandomEpisode(t, season)
}

func TestQueries_GetEpisode(t *testing.T) {
	anime := createRandomAnime(t)
	season := createRandomSeason(t, anime)
	episode1 := createRandomEpisode(t, season)

	episode2, err := testQueries.GetEpisode(context.Background(), episode1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, episode2)

	require.Equal(t, episode1.ID, episode2.ID)
	require.Equal(t, episode1.SeasonID, episode2.SeasonID)
	require.Equal(t, episode1.Name, episode2.Name)
	require.Equal(t, episode1.Value, episode2.Value)
	require.Equal(t, episode1.Description, episode2.Description)
	require.Equal(t, episode1.Status, episode2.Status)
	require.WithinDuration(t, episode1.CreatedAt, episode2.CreatedAt, time.Second)
	require.WithinDuration(t, episode1.UpdatedAt, episode2.UpdatedAt, time.Second)
}

func TestQueries_ListEpisode(t *testing.T) {
	anime := createRandomAnime(t)
	season := createRandomSeason(t, anime)
	for i := 0; i < 10; i++ {
		createRandomEpisode(t, season)
	}

	arg := ListEpisodeParams{
		SeasonID: season.ID,
		Limit:    5,
		Offset:   5,
	}

	episodes, err := testQueries.ListEpisode(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, episodes, 5)

	for _, episode := range episodes {
		require.NotEmpty(t, episode)
	}
}

func TestQueries_UpdateEpisode(t *testing.T) {
	anime := createRandomAnime(t)
	season := createRandomSeason(t, anime)
	episode1 := createRandomEpisode(t, season)

	arg := UpdateEpisodeParams{
		ID:          episode1.ID,
		Name:        util.RandomString(6),
		Value:       util.RandomInt64(1, 255),
		Description: util.RandomString(16),
		Status:      util.RandomInt64(1, 7),
	}

	episode2, err := testQueries.UpdateEpisode(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, episode2)

	require.Equal(t, episode1.ID, episode2.ID)
	require.Equal(t, episode1.SeasonID, episode2.SeasonID)
	require.Equal(t, arg.Name, episode2.Name)
	require.Equal(t, arg.Value, episode2.Value)
	require.Equal(t, arg.Description, episode2.Description)
	require.Equal(t, arg.Status, episode2.Status)
	require.WithinDuration(t, episode1.CreatedAt, episode2.CreatedAt, time.Second)
	//require.NotEqual(t, episode1.UpdatedAt, episode2.UpdatedAt)
}

func TestQueries_DeleteEpisode(t *testing.T) {
	anime := createRandomAnime(t)
	season := createRandomSeason(t, anime)
	episode1 := createRandomEpisode(t, season)

	err := testQueries.DeleteEpisode(context.Background(), episode1.ID)
	require.NoError(t, err)

	episode2, err := testQueries.GetEpisode(context.Background(), episode1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, episode2)
}
