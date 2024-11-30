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
	database        arango.ArangoDB
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
		database:        db,
		videoCollection: videoCollection,
		ctx:             ctx,
	}, nil
}

func (c videoRepository) Get(key string) (*dao.Video, error) {
	var video dao.Video

	opts := arangodb.CollectionDocumentReadOptions{}
	_, err := c.database.ReadDocumentWithOptions(c.ctx, c.videoCollection, key, &video, &opts)
	if err != nil {
		return nil, err
	}

	return &video, nil
}

func (c videoRepository) Create(video dto.Video) error {
	opts := arangodb.CollectionDocumentCreateOptions{}
	_, err := c.database.CreateDocumentWithOptions(c.ctx, c.videoCollection, video, &opts)
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
	_, err := c.database.UpdateDocumentWithOptions(c.ctx, c.videoCollection, videoUpdate.Key, videoUpdate, &options)
	if err != nil {
		return nil, err
	}

	return &video, nil
}

func (c videoRepository) Delete(key string) error {
	err := c.database.DeleteDocumentWithOptions(c.ctx, c.videoCollection, key, nil)
	if err != nil {
		return err
	}
	return nil
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
	cursor, err := c.database.Query(c.ctx, query, opts)
	if err != nil {
		return nil, err
	}
	defer c.database.CloseCursor(cursor)

	if !c.database.CursorHasMore(cursor) {
		return nil, fmt.Errorf("video not found")
	}

	_, err = c.database.CursorReadDocument(c.ctx, cursor, &video)
	if err != nil {
		return nil, err
	}

	return &video, nil
}
