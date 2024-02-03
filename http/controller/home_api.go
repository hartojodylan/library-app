package controller

import (
	"github.com/dylanh/library-app/app"
	"os"

	"github.com/gookit/rux"
	"github.com/gookit/view"
)

func Home(c *rux.Context) {
	c.JSON(200, rux.M{"hello": "welcome"})
}

func SwagDoc(c *rux.Context) {
	fInfo, err := os.Stat("static/swagger.json")
	if err != nil {
		c.AbortWithStatus(404, "swagger doc file not exists")
		return
	}

	data := map[string]string{
		"EnvName":    app.EnvName,
		"AppName":    app.Name,
		"JsonFile":   "/static/swagger.json",
		"SwgUIPath":  "/static/swagger-ui",
		"AssetPath":  "/static",
		"UpdateTime": fInfo.ModTime().Format(app.BaseDate),
	}

	// c.HTML(200, nil)
	view.Partial(c.Resp, "swagger.tpl", data)
}

// @Tags InternalApi
// @Summary Health check
// @Description get app health
// @Success 201 {string} json data
// @Failure 403 body is empty
// @Router /health [get]
func AppHealth(c *rux.Context) {
	data := map[string]interface{}{
		"status": "UP",
		"info":   app.GitInfo,
	}

	c.JSON(200, data)
}

func AppStatus(c *rux.Context) {
	data := map[string]interface{}{
		"status": "UP",
		"info":   app.GitInfo,
	}

	c.JSON(200, data)
}
