package main

import (
	"fmt"
	"github.com/dylanh/library-app/global"
	"github.com/dylanh/library-app/model/form"

	// boot and init some services(log, cache, eureka)
	"github.com/dylanh/library-app/app"
	"github.com/dylanh/library-app/app/bootstrap"
	"github.com/dylanh/library-app/app/clog"
	"github.com/dylanh/library-app/http"
	"os"

	// init redis, mongo, mysql connection

	"github.com/gookit/rux"
	"github.com/gookit/rux/handlers"
)

var router *rux.Router

func init() {
	bootstrap.Web()

	// router and routes
	router = rux.New()
	// global middleware
	router.Use(handlers.RequestLogger())

	http.AddRoutes(router)
}

// @title My Project API
// @version 1.0
// @description My Project API
// @termsOfService https://github.com/inhere

// @contact.name API Support
// @contact.url https://github.com/inhere
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /v1
func main() {
	clog.Printf(
		"======================== Begin Running(PID: %d) ========================\n",
		os.Getpid(),
	)

	global.BooksList = map[string]form.BooksList{
		"1": {
			"1",
			"harry potter",
			"J.K. Rowling",
			1,
		},
		"2": {
			"2",
			"hunger games",
			"Suzanne Collins",
			1,
		},
	}

	global.BooksBooking = map[string]form.BookBooking{
		"2": {
			"2",
			"2020-01-01",
			1,
		},
	}

	// default is listen and serve on 0.0.0.0:8080
	router.Listen(fmt.Sprintf("0.0.0.0:%d", app.HttpPort))
}
