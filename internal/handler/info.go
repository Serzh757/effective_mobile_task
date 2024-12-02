package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/rs/zerolog/log"

	"github.com/effective_mobile_task/internal/models"
)

// GetSongDetailFromAPI
// Выполняем запрос к внешнему API для получения данных о песне
func (h *MusicHandler) GetSongDetailFromAPI(group, song string) (models.SongDetail, error) {
	encodedGroup := url.QueryEscape(group)
	encodedSong := url.QueryEscape(song)
	apiURL := fmt.Sprintf(h.cfg.InfoBackend+"/info?group=%s&song=%s", encodedGroup, encodedSong)

	req, err := http.NewRequest(http.MethodGet, apiURL, nil)
	if err != nil {
		log.Printf("ERROR: Failed make request: %v", err)
		return models.SongDetail{}, err
	}

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("ERROR: Failed to request external API: %v", err)
		return models.SongDetail{}, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		log.Printf("WARNING: External API returned status code %d", response.StatusCode)
		return models.SongDetail{}, err
	}

	var apiData models.SongDetail
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Printf("ERROR: Failed to read API response: %v", err)
		return models.SongDetail{}, err
	}

	if err := json.Unmarshal(body, &apiData); err != nil {
		log.Printf("ERROR: Failed to parse API response: %v", err)
		return models.SongDetail{}, err
	}

	return apiData, nil
}
