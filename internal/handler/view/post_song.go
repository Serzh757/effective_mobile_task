package view

import (
	"github.com/effective_mobile_task/internal/models"
	"github.com/effective_mobile_task/internal/view"
)

func NewSong(in *models.Song) view.SongsByIdResponse {
	if in == nil {
		return view.SongsByIdResponse{}
	}
	return view.SongsByIdResponse{
		Data: view.Song{
			Group:       in.Group,
			Id:          int(in.ID),
			Link:        in.Link,
			Song:        in.Song,
			Text:        in.Text,
			ReleaseDate: in.ReleaseDate,
		},
	}
}
