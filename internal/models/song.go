package models

// Song представляет песню в базе данных
type Song struct {
	ID          uint   `pg:",pk" json:"id"`
	Group       string `json:"group"`
	Song        string `json:"song"`
	ReleaseDate string `json:"release_date"`
	Text        string `json:"text"`
	Link        string `json:"link"`
}

// SongDetail представляет детальную информацию о песне
type SongDetail struct {
	ReleaseDate string `json:"release_date"`
	Text        string `json:"text"`
	Link        string `json:"link"`
}

// NewSongRequest используется при добавлении новой песни
type NewSongRequest struct {
	Group string `json:"group"`
	Song  string `json:"song"`
}
