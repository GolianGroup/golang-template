package repositories

import (
	"context"
	"fmt"
	"golang_template/internal/database/arango"
	"golang_template/internal/repositories/models"
	"log"

	"github.com/arangodb/go-driver/v2/arangodb"
)

type VideoRepository interface {
	Get(key string) (*models.Video, error)
	Create(video models.Video) error
	Update(videoUpdate models.Video) (*models.Video, error)
	Delete(key string) error
	GetByName(name string) (*models.Video, error)
}

type videoRepository struct {
	db  arangodb.Database
	ctx context.Context
}

func NewVideoRepository(db arango.ArangoDB, ctx context.Context) (VideoRepository, error) {
	database := db.Database(ctx)

	return &videoRepository{
		db:  database,
		ctx: ctx,
	}, nil
}

func (c videoRepository) Get(key string) (*models.Video, error) {
	collection, err := c.db.GetCollection(c.ctx, "videos_collection", nil)
	if err != nil {
		log.Println("Failed to get videos collection")
		return nil, err
	}

	var video models.Video

	opts := arangodb.CollectionDocumentReadOptions{}
	_, err = collection.ReadDocumentWithOptions(c.ctx, key, &video, &opts)
	if err != nil {
		return nil, err
	}

	return &video, nil
}

func (c videoRepository) Create(video models.Video) error {
	collection, err := c.db.GetCollection(c.ctx, "videos_collection", nil)
	if err != nil {
		log.Println("Failed to get videos collection")
		return err
	}

	opts := arangodb.CollectionDocumentCreateOptions{}
	_, err = collection.CreateDocumentWithOptions(c.ctx, video, &opts)
	if err != nil {
		return err
	}
	return nil
}

func (c videoRepository) Update(videoUpdate models.Video) (*models.Video, error) {
	collection, err := c.db.GetCollection(c.ctx, "videos_collection", nil)
	if err != nil {
		log.Println("Failed to get videos collection")
		return nil, err
	}

	var video models.Video
	withWaitForSync := true
	keepNull := true
	options := arangodb.CollectionDocumentUpdateOptions{
		WithWaitForSync: &withWaitForSync,
		NewObject:       &video,
		KeepNull:        &keepNull,
	}
	_, err = collection.UpdateDocumentWithOptions(c.ctx, videoUpdate.Key, videoUpdate, &options)
	if err != nil {
		return nil, err
	}

	return &video, nil
}

func (c videoRepository) Delete(key string) error {
	collection, err := c.db.GetCollection(c.ctx, "videos_collection", nil)
	if err != nil {
		log.Println("Failed to get videos collection")
		return err
	}

	_, err = collection.DeleteDocumentWithOptions(c.ctx, key, nil)
	if err != nil {
		return err
	}
	return nil
}

func (c videoRepository) GetByName(name string) (*models.Video, error) {
	collection, err := c.db.GetCollection(c.ctx, "videos_collection", nil)
	if err != nil {
		log.Println("Failed to get videos collection")
		return nil, err
	}

	var video models.Video
	query := `
	FOR v IN @@collection
		FILTER v.name == @name
		LIMIT 1
		RETURN v
	`
	opts := arangodb.QueryOptions{
		BindVars: map[string]interface{}{
			"@collection": collection.Name(),
			"name":        name,
		},
	}
	cursor, err := c.db.Query(c.ctx, query, &opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close()

	if !cursor.HasMore() {
		return nil, fmt.Errorf("video not found")
	}

	_, err = cursor.ReadDocument(c.ctx, &video)
	if err != nil {
		return nil, err
	}

	return &video, nil
}
