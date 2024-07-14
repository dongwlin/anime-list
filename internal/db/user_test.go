package db

import (
	"context"
	"database/sql"
	"github.com/dongwlin/anime-list/internal/pkg/util"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func createRandomUser(t *testing.T) User {
	hashedPassword, err := util.HashPassword(util.RandomString(8))
	require.NoError(t, err)

	arg := CreateUserParams{
		Username:       util.RandomUsername(),
		HashedPassword: hashedPassword,
		IsAdmin:        util.RandomBool(),
		Desc:           util.RandomString(16),
		Status:         0,
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.IsAdmin, user.IsAdmin)
	require.Equal(t, arg.Desc, user.Desc)
	require.Equal(t, arg.Status, user.Status)
	require.NotZero(t, user.CreatedAt)
	require.NotZero(t, user.UpdatedAt)

	return user
}

func TestQueries_CreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestQueries_GetUser(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testQueries.GetUser(context.Background(), user1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.HashedPassword, user2.HashedPassword)
	require.Equal(t, user1.IsAdmin, user2.IsAdmin)
	require.Equal(t, user1.Desc, user2.Desc)
	require.Equal(t, user1.Status, user2.Status)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
	require.WithinDuration(t, user1.UpdatedAt, user2.UpdatedAt, time.Second)
}

func TestQueries_ListUser(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomUser(t)
	}

	arg := ListUserParams{
		Limit:  5,
		Offset: 5,
	}

	users, err := testQueries.ListUser(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, users, 5)

	for _, user := range users {
		require.NotEmpty(t, user)
	}
}

func TestQueries_UpdateUser(t *testing.T) {
	user1 := createRandomUser(t)

	hashedPassword, err := util.HashPassword(util.RandomString(8))
	require.NoError(t, err)

	arg := UpdateUserParams{
		ID:             user1.ID,
		HashedPassword: hashedPassword,
		IsAdmin:        util.RandomBool(),
		Desc:           util.RandomString(16),
		Status:         util.RandomInt64(1, 7),
	}

	user2, err := testQueries.UpdateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, arg.HashedPassword, user2.HashedPassword)
	require.Equal(t, arg.IsAdmin, user2.IsAdmin)
	require.Equal(t, arg.Desc, user2.Desc)
	require.Equal(t, arg.Status, user2.Status)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
	//require.NotEqual(t, user1.UpdatedAt, user2.UpdatedAt)
}

func TestQueries_DeleteUser(t *testing.T) {
	user1 := createRandomUser(t)

	err := testQueries.DeleteUser(context.Background(), user1.ID)
	require.NoError(t, err)

	user2, err := testQueries.GetUser(context.Background(), user1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, user2)
}
