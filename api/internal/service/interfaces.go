package service

import (
	"context"

	"github.com/circuit-shell/playlist-builder-back/internal/model"
)

type SongServiceInterface interface {
	CreateSong(ctx context.Context, req model.CreateSongRequest) (*model.Song, error)
	GetAllSongs(ctx context.Context) ([]model.Song, error)
}
