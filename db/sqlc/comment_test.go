package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/test-go/testify/require"
	"gopkg.in/guregu/null.v4/zero"
)

func createRandomComment(t *testing.T) Comment {
	user := createRandomUser(t)

	arg := CreateCommentParams{
		Comment: fake.Lorem().Text(30),
		UserID: int32(user.ID),
		LikeCount: int32(fake.RandomDigit()),
		PostID: int32(fake.RandomDigit()),
	}

	comment, err := testStore.CreateComment(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, comment)

	require.Equal(t, arg.Comment, comment.Comment)
	require.Equal(t, arg.UserID, comment.UserID)
	require.Equal(t, arg.LikeCount, comment.LikeCount)
	require.Equal(t, arg.PostID, comment.PostID)

	require.NotZero(t, comment.ID)

	return comment
}

func TestCreateComment(t *testing.T) {
	createRandomComment(t)
}

func TestGetComment(t *testing.T) {
	comment1 := createRandomComment(t)
	comment2, err := testStore.GetComment(context.Background(), comment1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, comment2)

	require.Equal(t, comment1.Comment, comment2.Comment)
	require.Equal(t, comment1.UserID, comment2.UserID)
	require.Equal(t, comment1.LikeCount, comment2.LikeCount)
	require.Equal(t, comment1.PostID, comment2.PostID)
}

func TestListComment(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomComment(t)
	}

	arg := ListCommentsParams{
		Limit: 5,
		Offset: 5,
	}

	comments, err := testStore.ListComments(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, comments, 5)
}

func TestUpdateComment(t *testing.T) {
	comment1 := createRandomComment(t)
	user := createRandomUser(t)

	arg := UpdateCommentParams{
		ID: comment1.ID,
		Comment: zero.StringFrom(fake.Lorem().Text(30)),
		UserID: zero.IntFrom(user.ID),
		LikeCount: zero.IntFrom(int64(fake.RandomDigit())),
		PostID: zero.IntFrom(int64(fake.RandomDigit())),
	}

	comment2, err := testStore.UpdateComment(context.Background(), arg)
	require.NoError(t, err)

	assert.NotEmpty(t, comment2)
	assert.Equal(t, arg.Comment.String, string(comment2.Comment))
}

func TestDeleteComment(t *testing.T) {
	comment1 := createRandomComment(t)
	err := testStore.DeleteComment(context.Background(), comment1.ID)
	assert.NoError(t, err)

	comment2, err := testStore.GetComment(context.Background(), comment1.ID)
	assert.EqualError(t, err, ErrRecordNotFound.Error())
	assert.Empty(t, comment2)
}
