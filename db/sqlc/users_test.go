package db

import (
	"context"
	"testing"
	"time"

	"github.com/test-go/testify/require"
	"gopkg.in/guregu/null.v4/zero"
)

func createRandomUser(t *testing.T) User {
	arg := CreateUserParams{
		FullName: fake.Person().Name(),
		UserAvatar: int32(fake.RandomDigit()),
		CountPosts: fake.Int32(),
		Email: fake.Internet().Email(),
		Pasword: "secret",
	}

	user, err := testStore.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.UserAvatar, user.UserAvatar)
	require.Equal(t, arg.CountPosts, user.CountPosts)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.Pasword, user.Pasword)
	require.True(t, user.UpdatedAt.Valid)
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestCreateUser(t *testing.T ) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testStore.GetUser(context.Background(), user1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.FullName, user2.FullName)
	require.Equal(t, user1.UserAvatar, user2.UserAvatar)
	require.Equal(t, user1.CountPosts, user2.CountPosts)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.Pasword, user2.Pasword)
	require.WithinDuration(t, user1.UpdatedAt.Time, user2.UpdatedAt.Time, time.Second)
	require.WithinDuration(t, user1.CreatedAt.Time, user2.CreatedAt.Time, time.Second)
}

func TestDeleteUser(t *testing.T) {
	user1 := createRandomUser(t)
	err := testStore.DeleteUser(context.Background(), user1.ID)
	require.NoError(t, err)

	user2, err := testStore.GetUser(context.Background(), user1.ID)
	require.Error(t, err)
	require.EqualError(t, err, ErrRecordNotFound.Error())
	require.Empty(t, user2)
}

func TestListUser(t *testing.T) {
	var lastUser User
	for i := 0; i < 10; i++ {
		lastUser = createRandomUser(t)
	}

	arg := ListUsersParams{
		FullName: lastUser.FullName,
		Limit: 5,
		Offset: 0,
	}

	users, err := testStore.ListUsers(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, users)

	for _, user := range users {
		require.NotEmpty(t, user)
		require.Equal(t, lastUser.FullName, user.FullName)
	}
}

func TestUpdateUser(t *testing.T) {
	user1 := createRandomUser(t)

	arg := UpdateUserParams{
		ID: user1.ID,
		FullName: zero.StringFrom(user1.FullName),
		UserAvatar: zero.IntFrom(int64(user1.UserAvatar)),
		CountPosts: zero.IntFrom(int64(user1.CountPosts)),
		Email: zero.StringFrom(user1.Email),
		Pasword: zero.StringFrom(user1.Pasword),
	}

	user2, err := testStore.UpdateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1, user2)
	require.WithinDuration(t, user1.CreatedAt.Time, user2.CreatedAt.Time, time.Second)
}