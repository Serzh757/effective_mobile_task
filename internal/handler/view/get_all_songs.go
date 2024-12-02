package view

import (
	"github.com/effective_mobile_task/internal/models"
	"github.com/effective_mobile_task/internal/view"
)

func NewSongs(in []models.Song) view.SongsListResponse {
	if len(in) == 0 {
		return view.SongsListResponse{}
	}

	result := make([]view.Song, 0, len(in))

	for _, song := range in {
		result = append(result, view.Song{
			Group:       song.Group,
			Id:          int(song.ID),
			Link:        song.Link,
			Song:        song.Song,
			Text:        song.Text,
			ReleaseDate: song.ReleaseDate,
		})
	}

	return view.SongsListResponse{
		Data: result,
	}
}
