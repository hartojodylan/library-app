package controller

import (
	"github.com/dylanh/library-app/app"
	"github.com/dylanh/library-app/model"
	"strconv"

	"github.com/gookit/goutil/jsonutil"
	"github.com/gookit/rux"
	"github.com/sirupsen/logrus"
)

// BaseApi controller
type BaseApi struct {
	lang string
}

// getPageAndSize get and format page, size params
func (a *BaseApi) getPageAndSize(c *rux.Context) (int, int) {
	pageStr := c.Query("page", "1")
	sizeStr := c.Query("size", app.PageSizeStr)

	page, _ := strconv.Atoi(pageStr)
	size, _ := strconv.Atoi(sizeStr)

	return app.FormatPageAndSize(page, size)
}

func (a *BaseApi) JSON(c *rux.Context, status int, data interface{}) {
	bs, err := jsonutil.Encode(data)
	if err != nil {
		c.AddError(err)
		return
	}

	c.JSONBytes(status, bs)
}

// DataRes response json data
func (a *BaseApi) DataRes(c *rux.Context, data interface{}) *model.JsonData {
	return a.MakeRes(0, nil, "", data)
}

// MakeRes
// code custom error code
// empty map:
//
//	c.DataRes(map[string]string{})
//
// empty list:
//
//	c.DataRes([]int{})
//
// err  real error message, the message will not output, only write to log file.
func (a *BaseApi) MakeRes(code int, err error, message string, data interface{}) *model.JsonData {
	if data == nil {
		// data = map[string]string{}
		data = []string{}
	}

	// log and print error message
	if err != nil {
		logrus.Warn("detected response error", "code", code, "message", err.Error())

		// if open debug
		if app.Debug {
			data = map[string]string{"debug_msg": err.Error()}
		}
	}

	return &model.JsonData{Code: code, Message: message, Data: data}
}
