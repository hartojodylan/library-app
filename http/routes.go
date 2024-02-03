package http

import (
	"github.com/dylanh/library-app/http/controller"
	"github.com/gookit/rux"
)

// AddRoutes add http routes
func AddRoutes(r *rux.Router) {
	// static assets
	r.StaticDir("/static", "static")

	r.GET("/", controller.Home)
	r.GET("/apidoc", controller.SwagDoc)

	// status
	r.GET("/health", controller.AppHealth)
	r.GET("/status", controller.AppStatus)

	r.GET("/ping", func(c *rux.Context) {
		c.Text(200, "pong")
	})

	r.Group("/v1", func() {
		r.GET("/health", controller.AppHealth)

		internal := new(controller.InternalApi)
		r.GET("/config", internal.Config)

		user := new(controller.BooksApi)
		r.GET("/books/{subject}", user.GetBooksListBySubject)
		r.POST("/books", user.SaveBookBooking)
	})

	// not found routes
	r.NotFound(func(c *rux.Context) {
		c.JSONBytes(404, []byte(`{"code": 0, "message": "page not found", data: {}}`))
	})
}
