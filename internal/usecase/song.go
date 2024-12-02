package usecase

import (
	"github.com/effective_mobile_task/internal/models"
	"github.com/effective_mobile_task/internal/repository"
)

type SongUseCase struct {
	repository *repository.SongRepository
}

func NewSongUseCase(repository *repository.SongRepository) *SongUseCase {
	return &SongUseCase{
		repository: repository,
	}
}

// AllSongs - Получение всех песен с пагинацией
func (uc *SongUseCase) AllSongs(page, limit int) ([]models.Song, error) {
	return uc.repository.GetAllSongs(page, limit)
}

// SongByID - Получение песни по идентификатору
func (uc *SongUseCase) SongByID(songID int) (*models.Song, error) {
	return uc.repository.GetSongByID(uint(songID))
}

// SaveSong - Сохранение новой песни
func (uc *SongUseCase) SaveSong(req models.Song) (*models.Song, error) {
	return uc.repository.CreateSong(req)
}

// RemoveSong - Удаление песни
func (uc *SongUseCase) RemoveSong(songID int) error {
	return uc.repository.DeleteSong(uint(songID))
}

// UpdatedSong - Обновление данных песни
func (uc *SongUseCase) UpdatedSong(req models.Song) (*models.Song, error) {
	return uc.repository.UpdateSong(req)
}
