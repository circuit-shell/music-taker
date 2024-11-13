package sqlite

import (
	"context"
	"database/sql"

	"github.com/circuit-shell/playlist-builder-back/internal/model"
)

type SongRepository struct {
	db *sql.DB
}

func NewSongRepository(db *sql.DB) *SongRepository {
	return &SongRepository{db: db}
}

func (r *SongRepository) Create(ctx context.Context, song *model.Song) error {
	query := `
        INSERT INTO songs (id, title, artist, album, year, genre, created_at)
        VALUES (?, ?, ?, ?, ?, ?, ?)
    `
	_, err := r.db.ExecContext(ctx, query,
		song.ID, song.Title, song.Artist, song.Album, song.Year, song.Genre, song.CreatedAt)
	return err
}

func (r *SongRepository) GetAll(ctx context.Context) ([]model.Song, error) {
	query := `SELECT id, title, artist, album, year, genre, created_at FROM songs`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var songs []model.Song
	for rows.Next() {
		var song model.Song
		err := rows.Scan(
			&song.ID,
			&song.Title,
			&song.Artist,
			&song.Album,
			&song.Year,
			&song.Genre,
			&song.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		songs = append(songs, song)
	}
	return songs, nil
}
