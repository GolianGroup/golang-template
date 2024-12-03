package services

import (
	dao "golang_template/handler/daos"
	dto "golang_template/handler/dtos"
	"golang_template/internal/repositories"
	"golang_template/internal/repositories/models"
)

type VideoService interface {
	GetVideo(key string) (*dao.Video, error)
	CreateVideo(video dto.Video) error
	UpdateVideo(videoUpdate dto.VideoUpdate) (*dao.Video, error)
	DeleteVideo(key string) error
	GetVideoByName(name string) (*dao.VideoByName, error)
}

type videoService struct {
	videoRepository repositories.VideoRepository
}

func NewVideoService(videoRepository repositories.VideoRepository) VideoService {
	return &videoService{
		videoRepository: videoRepository,
	}
}

func (s videoService) GetVideo(key string) (*dao.Video, error) {
	foundVideo, err := s.videoRepository.Get(key)
	if err != nil {
		return nil, err
	}

	video := dao.Video{
		Key:         foundVideo.Key,
		Publishable: foundVideo.Publishable,
		Categories:  foundVideo.Categories,
		Description: foundVideo.Description,
		Name:        foundVideo.Name,
		Views:       foundVideo.Views,
		Type:        foundVideo.Type,
	}

	return &video, nil
}

func (s videoService) CreateVideo(video dto.Video) error {
	createVideo := models.Video{
		Publishable: video.Publishable,
		Categories:  video.Categories,
		Description: video.Description,
		Name:        video.Name,
		Type:        video.Type,
	}

	err := s.videoRepository.Create(createVideo)
	if err != nil {
		return err
	}
	return nil
}

func (s videoService) UpdateVideo(videoUpdate dto.VideoUpdate) (*dao.Video, error) {
	updateVideo := models.Video{
		Key:         videoUpdate.Key,
		Categories:  videoUpdate.Categories,
		Description: videoUpdate.Description,
		Name:        videoUpdate.Name,
		Views:       videoUpdate.Views,
	}
	updatedVideo, err := s.videoRepository.Update(updateVideo)
	if err != nil {
		return nil, err
	}

	video := dao.Video{
		Key:         updatedVideo.Key,
		Publishable: updatedVideo.Publishable,
		Categories:  updatedVideo.Categories,
		Description: updatedVideo.Description,
		Name:        updatedVideo.Name,
		Views:       updatedVideo.Views,
		Type:        updatedVideo.Type,
	}

	return &video, nil
}

func (s videoService) DeleteVideo(key string) error {
	err := s.videoRepository.Delete(key)
	if err != nil {
		return err
	}
	return nil
}

func (s videoService) GetVideoByName(name string) (*dao.VideoByName, error) {
	foundVideo, err := s.videoRepository.GetByName(name)
	if err != nil {
		return nil, err
	}

	video := dao.VideoByName{
		Publishable: foundVideo.Publishable,
		Categories:  foundVideo.Categories,
		Description: foundVideo.Description,
		Name:        foundVideo.Name,
		Views:       foundVideo.Views,
		Type:        foundVideo.Type,
	}

	return &video, nil
}
