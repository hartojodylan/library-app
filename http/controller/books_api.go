package controller

import (
	"fmt"
	"github.com/dylanh/library-app/app/errcode"
	"github.com/dylanh/library-app/model/form"
	"github.com/gookit/rux"
	"strconv"
	"time"
)

type BooksApi struct {
	BaseApi
}

// AddRoutes for the API controller
func (u *BooksApi) AddRoutes(g *rux.Router) {
	g.GET("/books/{subject}", u.GetBooksListBySubject)
	g.POST("/books", u.SaveBookBooking)
}

// @Tags BooksApi
// @Summary Get multiple book details per page
// @Description get book details
// @Param   subject     path    string     true        "book subject"
// @Failure 200 {object} model.JsonMapData "Need book subject"
// @Failure 404 {object} model.JsonMapData "Can't find book subject"
// @Router /books/{subject} [get]
func (u *BooksApi) GetBooksListBySubject(c *rux.Context) {
	// path param
	subject := c.Param("subject")

	// query params
	limit, ok := c.QueryParam("limit")
	if !ok {
		c.JSON(400, u.MakeRes(errcode.ErrParam, fmt.Errorf("limit is empty: %v", limit), "fail", nil))
		return
	}
	page, ok := c.QueryParam("page")
	if !ok {
		c.JSON(400, u.MakeRes(errcode.ErrParam, fmt.Errorf("page is empty: %v", page), "fail", nil))
		return
	}

	// convert limit and page to int type
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		c.JSON(400, u.MakeRes(errcode.ErrParam, err, "fail to convert limit to int", nil))
		return
	}

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		c.JSON(400, u.MakeRes(errcode.ErrParam, err, "fail to convert page to int", nil))
		return
	}

	// get book details
	res, err := Client.GetBooksListBySubject(subject, limitInt, pageInt)
	if err != nil {
		c.JSON(500, u.MakeRes(errcode.ErrNotFound, err, "fail", nil))
		return
	}

	c.JSON(200, u.MakeRes(200, nil, "success", res))
	return
}

// @Tags BooksApi
// @Summary Create a new book booking
// @Description insert book booking data
// @Param   bodyData     body    form.SaveBookBookingRequest     true  "pickUpSchedule format: 2006-01-02 15:04:05"
// @Failure 200 {object} model.JsonMapData "Need booking data!!"
// @Failure 404 {object} model.JsonMapData "Cannot insert booking data"
// @Router /books [post]
func (u *BooksApi) SaveBookBooking(c *rux.Context) {
	var f form.SaveBookBookingRequest

	if err := c.Bind(&f); err != nil {
		c.AbortThen().JSON(400, u.MakeRes(400, err, "error param", []string{}))
		return
	}

	// validate pick up schedule is valid datetime format
	_, err := time.Parse("2006-01-02 15:04:05", f.PickUpSchedule)
	if err != nil {
		c.JSON(400, u.MakeRes(400, err, "invalid pick up schedule format", nil))
		return
	}

	res, err := Client.SaveBookBooking(&f)
	if err != nil {
		c.JSON(res.Code, u.MakeRes(res.Code, err, res.Message, res.SuccessfullyBookedBookIDs))
		return
	}

	c.JSON(res.Code, u.MakeRes(res.Code, nil, res.Message, res.SuccessfullyBookedBookIDs))
}
