package repository

import (
	"github.com/go-pg/pg/v10"
	"github.com/rs/zerolog/log"

	"github.com/effective_mobile_task/internal/models"
)

type SongRepository struct {
	DB *pg.DB
}

func NewSongRepository(db *pg.DB) *SongRepository {
	return &SongRepository{
		DB: db,
	}
}

// CreateSong - Сохранение песни в БД
func (r *SongRepository) CreateSong(req models.Song) (*models.Song, error) {
	if _, err := r.DB.Model(req).Insert(); err != nil {
		log.Fatal().Err(err).Msg("Failed to add new song")
		return nil, err
	}

	lastInsertSong := &models.Song{}
	err := r.DB.Model(lastInsertSong).
		Where("songs.id = ?", req.ID).
		Select()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to get new song")
		return nil, err
	}

	return lastInsertSong, nil
}

// GetAllSongs - Получение всех песен с пагинацией
func (r *SongRepository) GetAllSongs(page, limit int) ([]models.Song, error) {
	var (
		songs  = make([]models.Song, 0)
		offset = (page - 1) * limit
	)
	err := r.DB.Model(&songs).Offset(offset).Limit(limit).Select()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to get all songs")
		return nil, err
	}
	return songs, nil
}

// GetSongByID - Получение песни по ИД
func (r *SongRepository) GetSongByID(songID uint) (*models.Song, error) {
	var song models.Song
	err := r.DB.Model(song).
		Where("songs.id = ?", songID).
		Select()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to retrieve song with ID")
		return nil, err
	}
	return &song, nil
}

// UpdateSong - Обновление данных существующей песни
func (r *SongRepository) UpdateSong(req models.Song) (*models.Song, error) {
	_, err := r.DB.Model(req).WherePK().Update()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to update song data")
		return nil, err
	}

	lastUpdateSong := &models.Song{}
	err = r.DB.Model(lastUpdateSong).Where("songs.id = ?", req.ID).Select()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to get updated song")
		return nil, err
	}

	return lastUpdateSong, nil
}

// DeleteSong - Удаление песни по ИД
func (r *SongRepository) DeleteSong(songID uint) error {
	song := &models.Song{
		ID: songID,
	}
	if err := r.DB.Model(song).Where("songs.id = ?", song.ID).Select(); err != nil {
		log.Fatal().Err(err).Msg("Failed to get song")
		return err
	}

	_, err := r.DB.Model(song).WherePK().Delete()

	return err
}
