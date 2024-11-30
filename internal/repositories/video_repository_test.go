package repositories

import (
	"context"
	"golang_template/internal/mocks"
	"golang_template/internal/repositories/dao"
	"golang_template/internal/repositories/dto"
	"testing"

	"github.com/arangodb/go-driver/v2/arangodb"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestVideoRepository(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockdb := mocks.NewMockArangoDB(ctrl)
	ctx := context.Background()

	var videoCollection arangodb.Collection
	mockdb.EXPECT().VideoCollection(ctx).Return(videoCollection, nil)

	repo, err := NewVideoRepository(mockdb, ctx)
	require.NoError(t, err)

	t.Run("Get video by key successfully", func(t *testing.T) {
		var meta arangodb.DocumentMeta
		mockdb.EXPECT().ReadDocumentWithOptions(ctx, videoCollection, "123", gomock.Any(), gomock.Any()).Return(meta, nil)
		video, err := repo.Get("123")
		require.NoError(t, err)

		assert.IsType(t, &dao.Video{}, video)
	})

	t.Run("Create video successfully", func(t *testing.T) {
		var response arangodb.CollectionDocumentCreateResponse
		mockdb.EXPECT().CreateDocumentWithOptions(ctx, videoCollection, gomock.Any(), gomock.Any()).Return(response, nil)
		video := dto.Video{
			Name:        "name",
			Publishable: true,
			Categories:  []string{"cat1", "cat2"},
			Description: "desc",
			Type:        "type",
		}
		err := repo.Create(video)
		require.NoError(t, err)
		assert.Equal(t, err, nil)
	})

	t.Run("Update video successfully", func(t *testing.T) {
		var response arangodb.CollectionDocumentUpdateResponse
		mockdb.EXPECT().UpdateDocumentWithOptions(ctx, videoCollection, "123", gomock.Any(), gomock.Any()).Return(response, nil)
		videoUpdate := dto.VideoUpdate{
			Key:         "123",
			Categories:  []string{"cat1", "cat2"},
			Description: "desc",
			Name:        "name",
			Views:       1,
		}
		video, err := repo.Update(videoUpdate)
		require.NoError(t, err)

		assert.IsType(t, &dao.Video{}, video)
	})

	t.Run("Delete video successfully", func(t *testing.T) {
		mockdb.EXPECT().DeleteDocumentWithOptions(ctx, videoCollection, "123", gomock.Any()).Return(nil)
		err := repo.Delete("123")
		require.NoError(t, err)

		assert.Equal(t, err, nil)
	})
}
