package repositories

import (
	"context"
	"fmt"
	"golang_template/internal/database/arango"
	"golang_template/internal/repositories/dao"
	"golang_template/internal/repositories/dto"
	"log"

	"github.com/arangodb/go-driver/v2/arangodb"
)

type VideoRepository interface {
	Get(key string) (*dao.Video, error)
	Create(video dto.Video) error
	Update(videoUpdate dto.VideoUpdate) (*dao.Video, error)
	Delete(key string) error
	GetByName(name string) (*dao.VideoByName, error)
}

type videoRepository struct {
	videoCollection arangodb.Collection
	ctx             context.Context
}

func NewVideoRepository(db arango.ArangoDB, ctx context.Context) (VideoRepository, error) {
	videoCollection, err := db.VideoCollection(ctx)
	if err != nil {
		log.Println("Failed to get video collection")
		return nil, err
	}
	return &videoRepository{
		videoCollection: videoCollection,
		ctx:             ctx,
	}, nil
}

func (c videoRepository) Get(key string) (*dao.Video, error) {
	var video dao.Video
	_, err := c.videoCollection.ReadDocument(c.ctx, key, &video)
	if err != nil {
		return nil, err
	}

	return &video, nil
}

func (c videoRepository) Create(video dto.Video) error {
	_, err := c.videoCollection.CreateDocument(c.ctx, video)
	if err != nil {
		return err
	}
	return nil
}

func (c videoRepository) Update(videoUpdate dto.VideoUpdate) (*dao.Video, error) {
	var video dao.Video
	withWaitForSync := true
	keepNull := true
	options := arangodb.CollectionDocumentUpdateOptions{
		WithWaitForSync: &withWaitForSync,
		NewObject:       &video,
		KeepNull:        &keepNull,
	}
	_, err := c.videoCollection.UpdateDocumentWithOptions(c.ctx, videoUpdate.Key, videoUpdate, &options)
	if err != nil {
		return nil, err
	}

	return &video, nil
}

func (c videoRepository) Delete(key string) error {
	_, err := c.videoCollection.DeleteDocument(c.ctx, key)
	return err
}

func (c videoRepository) GetByName(name string) (*dao.VideoByName, error) {
	var video dao.VideoByName
	query := `
	FOR v IN @@collection
		FILTER v.name == @name
		LIMIT 1
		RETURN v
	`
	opts := arangodb.QueryOptions{
		BindVars: map[string]interface{}{
			"@collection": c.videoCollection.Name(),
			"name":        name,
		},
	}
	cursor, err := c.videoCollection.Database().Query(c.ctx, query, &opts)
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
