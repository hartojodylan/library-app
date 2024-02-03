package books

import (
	"github.com/dylanh/library-app/core/books"
	"github.com/dylanh/library-app/model/form"
)

type BookClient interface {
	GetBookDetailsFromDB(bookIDs []string) (res []form.BooksList, err error)
	InsertBookBooking(bookIDs []string, pickUpSchedule string, userID int64) (successBookIDs []string, partialSuccess bool, err error)
	GetBooksListAPIBySubject(subject string, limit int, offset int) (res form.GetBooksListApiResponse, err error)
}

var (
	Client BookClient
)

func init() {
	Client = &books.BookCore{}
}

// mock section

type ClientMock struct {
	GetBookDetailsFromDBFunc     func(bookIDs []string) (res []form.BooksList, err error)
	InsertBookBookingFunc        func(bookIDs []string, pickUpSchedule string, userID int64) (successBookIDs []string, partialSuccess bool, err error)
	GetBooksListAPIBySubjectFunc func(subject string, limit int, offset int) (res form.GetBooksListApiResponse, err error)
}

var (
	mockGetBookDetailsFromDBFunc func(bookIDs []string) (res []form.BooksList, err error)
	mockInsertBookBookingFunc    func(bookIDs []string, pickUpSchedule string, userID int64) (successBookIDs []string, partialSuccess bool, err error)
	mockGetBooksListAPIBySubject func(subject string, limit int, offset int) (res form.GetBooksListApiResponse, err error)
)

func (m *ClientMock) GetBookDetailsFromDB(bookIDs []string) (res []form.BooksList, err error) {
	return mockGetBookDetailsFromDBFunc(bookIDs)
}

func (m *ClientMock) InsertBookBooking(bookIDs []string, pickUpSchedule string, userID int64) (successBookIDs []string, partialSuccess bool, err error) {
	return mockInsertBookBookingFunc(bookIDs, pickUpSchedule, userID)
}

func (m *ClientMock) GetBooksListAPIBySubject(subject string, limit int, offset int) (res form.GetBooksListApiResponse, err error) {
	return mockGetBooksListAPIBySubject(subject, limit, offset)
}
