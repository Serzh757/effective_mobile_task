package swagger

import (
	"net/http"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gin-gonic/gin"
	"github.com/swaggest/swgui/v4cdn"
)

func Register(r gin.IRouter, swFn func() (*openapi3.T, error), basePath string) error {
	swDoc, err := swFn()
	if err != nil {
		return err
	}
	swaggerHandler := v4cdn.NewHandler("App API "+basePath, basePath+"/docs.json", "/")

	r.GET("/documentation/*any", gin.WrapH(swaggerHandler))
	r.GET("/docs.json", func(c *gin.Context) {
		c.JSON(http.StatusOK, &swDoc)
	})
	return nil
}
