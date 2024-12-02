package errors

import (
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/ru"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	ru_translations "github.com/go-playground/validator/v10/translations/ru"
	"github.com/rs/zerolog/log"
)

var Wrapper = func() func(c *gin.Context) {
	var trans ut.Translator
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		trans, _ = ut.New(ru.New()).GetTranslator("ru")
		_ = ru_translations.RegisterDefaultTranslations(v, trans)
	}

	m := &middleware{
		trans: trans,
	}
	return m.make
}()

type middleware struct {
	trans ut.Translator
}

func (e *middleware) make(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Next()

	err := c.Errors.Last()
	if err == nil {
		return
	}

	log.Info().Err(err.Err)

	if c.Writer.Status() < http.StatusInternalServerError || errors.Is(err.Err, context.Canceled) {
		c.JSON(-1, newErrorByWrappingError(err.Err))
		return
	}

	c.JSON(-1, newErrorByWrappingError(err.Err))
}
