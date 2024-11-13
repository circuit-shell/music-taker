package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/circuit-shell/playlist-builder-back/internal/model"
	"github.com/circuit-shell/playlist-builder-back/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockSongService is a mock implementation of the SongServiceInterface
type MockSongService struct {
	mock.Mock
}

// Ensure MockSongService implements SongServiceInterface
var _ service.SongServiceInterface = (*MockSongService)(nil)

// CreateSong implements the interface method
func (m *MockSongService) CreateSong(ctx context.Context, req model.CreateSongRequest) (*model.Song, error) {
	args := m.Called(ctx, req)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Song), args.Error(1)
}

// GetAllSongs implements the interface method
func (m *MockSongService) GetAllSongs(ctx context.Context) ([]model.Song, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.Song), args.Error(1)
}

func TestSongHandler_CreateSong(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name         string
		request      model.CreateSongRequest
		setupMock    func(*MockSongService)
		expectedCode int
		expectedBody bool
	}{
		{
			name: "valid song",
			request: model.CreateSongRequest{
				Title:  "Test Song",
				Artist: "Test Artist",
				Album:  "Test Album",
				Year:   2024,
				Genre:  "Rock",
			},
			setupMock: func(ms *MockSongService) {
				ms.On("CreateSong", mock.Anything, mock.MatchedBy(func(req model.CreateSongRequest) bool {
					return req.Title == "Test Song" &&
						req.Artist == "Test Artist" &&
						req.Album == "Test Album" &&
						req.Year == 2024 &&
						req.Genre == "Rock"
				})).Return(&model.Song{
					ID:        uuid.New().String(),
					Title:     "Test Song",
					Artist:    "Test Artist",
					Album:     "Test Album",
					Year:      2024,
					Genre:     "Rock",
					CreatedAt: time.Now(),
				}, nil)
			},
			expectedCode: http.StatusCreated,
			expectedBody: true,
		},
		{
			name: "invalid request - missing required fields",
			request: model.CreateSongRequest{
				Title: "Test Song",
				// Missing other required fields
			},
			setupMock:    func(ms *MockSongService) {},
			expectedCode: http.StatusBadRequest,
			expectedBody: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create mock service
			mockService := new(MockSongService)
			tt.setupMock(mockService)

			// Create handler with mock service
			handler := NewSongHandler(mockService)

			// Setup router
			router := gin.New()
			router.POST("/songs", handler.CreateSong)

			// Create request
			jsonData, err := json.Marshal(tt.request)
			assert.NoError(t, err)

			req, err := http.NewRequest(http.MethodPost, "/songs", bytes.NewBuffer(jsonData))
			assert.NoError(t, err)
			req.Header.Set("Content-Type", "application/json")

			// Create response recorder
			resp := httptest.NewRecorder()

			// Perform request
			router.ServeHTTP(resp, req)

			// Assert response
			assert.Equal(t, tt.expectedCode, resp.Code)

			if tt.expectedBody {
				var response model.Song
				err = json.Unmarshal(resp.Body.Bytes(), &response)
				assert.NoError(t, err)
				assert.NotEmpty(t, response.ID)
				assert.Equal(t, tt.request.Title, response.Title)
				assert.Equal(t, tt.request.Artist, response.Artist)
				assert.Equal(t, tt.request.Album, response.Album)
				assert.Equal(t, tt.request.Year, response.Year)
				assert.Equal(t, tt.request.Genre, response.Genre)
			}

			// Verify that all expected mock calls were made
			mockService.AssertExpectations(t)
		})
	}
}

func TestSongHandler_GetAllSongs(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name         string
		setupMock    func(*MockSongService)
		expectedCode int
		expectedLen  int
	}{
		{
			name: "successful fetch - empty list",
			setupMock: func(ms *MockSongService) {
				ms.On("GetAllSongs", mock.Anything).Return([]model.Song{}, nil)
			},
			expectedCode: http.StatusOK,
			expectedLen:  0,
		},
		{
			name: "successful fetch - with songs",
			setupMock: func(ms *MockSongService) {
				songs := []model.Song{
					{
						ID:        uuid.New().String(),
						Title:     "Test Song 1",
						Artist:    "Test Artist 1",
						Album:     "Test Album 1",
						Year:      2024,
						Genre:     "Rock",
						CreatedAt: time.Now(),
					},
					{
						ID:        uuid.New().String(),
						Title:     "Test Song 2",
						Artist:    "Test Artist 2",
						Album:     "Test Album 2",
						Year:      2024,
						Genre:     "Pop",
						CreatedAt: time.Now(),
					},
				}
				ms.On("GetAllSongs", mock.Anything).Return(songs, nil)
			},
			expectedCode: http.StatusOK,
			expectedLen:  2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create mock service
			mockService := new(MockSongService)
			tt.setupMock(mockService)

			// Create handler with mock service
			handler := NewSongHandler(mockService)

			// Setup router
			router := gin.New()
			router.GET("/songs", handler.GetAllSongs)

			// Create request
			req, err := http.NewRequest(http.MethodGet, "/songs", nil)
			assert.NoError(t, err)

			// Create response recorder
			resp := httptest.NewRecorder()

			// Perform request
			router.ServeHTTP(resp, req)

			// Assert response
			assert.Equal(t, tt.expectedCode, resp.Code)

			var response []model.Song
			err = json.Unmarshal(resp.Body.Bytes(), &response)
			assert.NoError(t, err)
			assert.Len(t, response, tt.expectedLen)

			// Verify that all expected mock calls were made
			mockService.AssertExpectations(t)
		})
	}
}
