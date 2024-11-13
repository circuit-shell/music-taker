package service

import (
	"context"
	"time"

	"github.com/circuit-shell/playlist-builder-back/internal/model"
	"github.com/circuit-shell/playlist-builder-back/internal/repository"
	"github.com/google/uuid"
)

type SongService struct {
	repo repository.SongRepository
}

func NewSongService(repo repository.SongRepository) *SongService {
	return &SongService{repo: repo}
}

func (s *SongService) CreateSong(ctx context.Context, req model.CreateSongRequest) (*model.Song, error) {
	song := &model.Song{
		ID:        uuid.New().String(),
		Title:     req.Title,
		Artist:    req.Artist,
		Album:     req.Album,
		Year:      req.Year,
		Genre:     req.Genre,
		CreatedAt: time.Now(),
	}

	if err := s.repo.Create(ctx, song); err != nil {
		return nil, err
	}

	return song, nil
}

func (s *SongService) GetAllSongs(ctx context.Context) ([]model.Song, error) {
	return s.repo.GetAll(ctx)
}
