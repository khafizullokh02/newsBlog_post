package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/guregu/null.v4/zero"
)

func createRandomCategory(t *testing.T) Category {
	arg := CreateCategoryParams{
		Title:        int32(fake.RandomDigit()),
		PostCount:    int32(fake.RandomDigit()),
		ArticleCount: int32(fake.RandomDigit()),
	}

	category, err := testStore.CreateCategory(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, category)

	require.Equal(t, arg.Title, category.Title)
	require.Equal(t, arg.PostCount, category.PostCount)
	require.Equal(t, arg.ArticleCount, category.ArticleCount)

	require.NotZero(t, category.ID)
	return category
}

func TestCreateCategory(t *testing.T) {
	createRandomCategory(t)
}

func TestGetCategory(t *testing.T) {
	category1 := createRandomCategory(t)
	category2, err := testStore.GetCategory(context.Background(), category1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, category2)

	require.Equal(t, category1, category2)
}

func TestUpdateCategory(t *testing.T) {
	category1 := createRandomCategory(t)

	arg := UpdateCategoryParams{
		ID: category1.ID,
		Title: zero.IntFrom(int64(fake.RandomDigit())),
		PostCount: zero.IntFrom(int64(fake.RandomDigit())),
		ArticleCount: zero.IntFrom(int64(fake.RandomDigit())),
	}

	category2, err := testStore.UpdateCategory(context.Background(), arg)
	require.NoError(t, err)

	assert.NotEmpty(t, category2)
	assert.Equal(t, arg.Title.Int64, int64(category2.Title))
}

func TestDeleteCategory(t *testing.T) {
	category1 := createRandomCategory(t)
	err := testStore.DeleteCategory(context.Background(), category1.ID)
	assert.NoError(t, err)

	categor2, err := testStore.GetCategory(context.Background(), category1.ID)
	require.Error(t, err)
	assert.EqualError(t, err, ErrRecordNotFound.Error())
	assert.Empty(t, categor2)
}

func TestListCategory(t *testing.T) {
	var createdCategories []Category
	for i := 0; i < 10; i++ {
		createdCategories = append(createdCategories, createRandomCategory(t))
	}

	arg := ListCategoriesParams{
		Limit: 5,
		Offset: 0,
	}

	categories, err := testStore.ListCategories(context.Background(), arg)
	require.NoError(t, err)

	for idx, category := range categories {
		assert.NotEmpty(t, category)
		assert.Equal(t, createdCategories[len(createdCategories)-idx-1].ID, category.ID)
		assert.Equal(t, createdCategories[len(createdCategories)-idx-1].Title, category.Title)
		assert.Equal(t, createdCategories[len(createdCategories)-idx-1].ArticleCount, category.ArticleCount)
		assert.Equal(t, createdCategories[len(createdCategories)-idx-1].PostCount, category.PostCount)
	}
}