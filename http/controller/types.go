package controller

import (
	"github.com/dylanh/library-app/model/form"
	"github.com/dylanh/library-app/service/books"
)

type BookSvcClient interface {
	SaveBookBooking(f *form.SaveBookBookingRequest) (res form.SaveBookBookingResponse, err error)
	GetBooksListBySubject(subject string, limit int, page int) (res []form.BooksList, err error)
}

var (
	Client BookSvcClient
)

func init() {
	Client = &books.BookSvc{}
}

// mock section

type ClientMock struct {
	SaveBookBookingFunc       func(f *form.SaveBookBookingRequest) (res form.SaveBookBookingResponse, err error)
	GetBooksListBySubjectFunc func(subject string, limit int, page int) (res []form.BooksList, err error)
}

var (
	mockSaveBookBookingFunc   func(f *form.SaveBookBookingRequest) (res form.SaveBookBookingResponse, err error)
	mockGetBooksListBySubject func(subject string, limit int, page int) (res []form.BooksList, err error)
)

func (m *ClientMock) SaveBookBooking(f *form.SaveBookBookingRequest) (res form.SaveBookBookingResponse, err error) {
	return mockSaveBookBookingFunc(f)
}

func (m *ClientMock) GetBooksListBySubject(subject string, limit int, page int) (res []form.BooksList, err error) {
	return mockGetBooksListBySubject(subject, limit, page)
}
