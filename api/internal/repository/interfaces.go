package repository

import (
	"context"

	"github.com/circuit-shell/playlist-builder-back/internal/model"
)

type SongRepository interface {
	Create(ctx context.Context, song *model.Song) error
	GetAll(ctx context.Context) ([]model.Song, error)
}
