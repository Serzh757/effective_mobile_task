package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	dview "github.com/effective_mobile_task/internal/handler/view"
	"github.com/effective_mobile_task/internal/models"
	"github.com/effective_mobile_task/internal/view"
)

// AddSong /api/v1/song [POST]
// Добавление новой песни
func (h *MusicHandler) AddSong(ctx *gin.Context) {
	var req view.PostSongRequest
	if err := ctx.ShouldBind(&req); err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response, err := h.uc.SaveSong(makeAddSongRequest(req))
	if err != nil {
		_ = ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, dview.NewSong(response))
}

func makeAddSongRequest(req view.PostSongRequest) *models.Song {
	return &models.Song{
		Group: req.Group,
		Song:  req.Song,
	}
}

// UpdateSong /api/v1/song [PUT]
// Обновление существующей песни
func (h *MusicHandler) UpdateSong(ctx *gin.Context) {
	var req view.PutSongRequest
	if err := ctx.ShouldBind(&req); err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	oldModel, err := h.uc.SongByID(req.SongID)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if oldModel == nil {
		_ = ctx.AbortWithError(http.StatusUnprocessableEntity, errors.New("песня не найдена"))
		return
	}

	response, err := h.uc.UpdatedSong(makeUpdateSongRequest(req, oldModel))
	if err != nil {
		_ = ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, dview.NewSong(response))
}

func makeUpdateSongRequest(req view.PutSongRequest, oldModel *models.Song) *models.Song {
	if v := req.Song; v != nil {
		oldModel.Song = *v
	}
	if v := req.Link; v != nil {
		oldModel.Link = *v
	}
	if v := req.Text; v != nil {
		oldModel.Text = *v
	}
	if v := req.ReleaseDate; v != nil {
		oldModel.ReleaseDate = *v
	}
	if v := req.Group; v != nil {
		oldModel.Group = *v
	}

	return oldModel
}

// SongByID /api/v1/song/:id [GET]
// Получение песни по ИД
func (h *MusicHandler) SongByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		_ = ctx.AbortWithError(http.StatusUnprocessableEntity, errors.New("id песни не найден"))
		return
	}

	response, err := h.uc.SongByID(id)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, dview.NewSong(response))
}

// DeleteSongByID /api/v1/song/:id [DELETE]
// Удаление песни по ИД
func (h *MusicHandler) DeleteSongByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		_ = ctx.AbortWithError(http.StatusUnprocessableEntity, errors.New("id песни не найден"))
		return
	}

	err = h.uc.RemoveSong(id)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.Status(http.StatusNoContent)
}

// AllSongs /api/v1/songs [GET]
// Получение всех песен
func (h *MusicHandler) AllSongs(ctx *gin.Context) {
	var req view.GetAllSongsParams
	if err := ctx.ShouldBind(&req); err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response, err := h.uc.AllSongs(makeParamsForAllSongs(req))
	if err != nil {
		_ = ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, dview.NewSongs(response))
}

func makeParamsForAllSongs(req view.GetAllSongsParams) (int, int) {
	var (
		page = 1
		size = 5
	)

	if v := req.Page; v != nil {
		page = *v
	}

	if v := req.PerPage; v != nil {
		size = *v
	}

	return page, size
}
