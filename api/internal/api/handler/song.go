package handler

import (
	"net/http"

	"github.com/circuit-shell/playlist-builder-back/internal/model"
	"github.com/circuit-shell/playlist-builder-back/internal/service"
	"github.com/gin-gonic/gin"
)

type SongHandler struct {
	songService service.SongServiceInterface
}

func NewSongHandler(songService service.SongServiceInterface) *SongHandler {
	return &SongHandler{songService: songService}
}

func (h *SongHandler) CreateSong(c *gin.Context) {
	var req model.CreateSongRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	song, err := h.songService.CreateSong(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, song)
}

func (h *SongHandler) GetAllSongs(c *gin.Context) {
	songs, err := h.songService.GetAllSongs(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, songs)
}
