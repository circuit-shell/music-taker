package model

import "time"

type Song struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Artist    string    `json:"artist"`
	Album     string    `json:"album"`
	Year      int       `json:"year"`
	Genre     string    `json:"genre"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateSongRequest struct {
	Title  string `json:"title" binding:"required"`
	Artist string `json:"artist" binding:"required"`
	Album  string `json:"album" binding:"required"`
	Year   int    `json:"year" binding:"required"`
	Genre  string `json:"genre" binding:"required"`
}
