package books

import (
	"github.com/dylanh/library-app/global"
	"github.com/dylanh/library-app/model/form"
	"strings"
)

// SaveBookBooking create a new book booking
func (b *BookSvc) SaveBookBooking(f *form.SaveBookBookingRequest) (res form.SaveBookBookingResponse, err error) {
	// check if book exist from db(mock)
	bookDetails, err := Client.GetBookDetailsFromDB(f.BookID)

	// if no books exist
	if len(bookDetails) == 0 || err != nil {
		return form.SaveBookBookingResponse{
			Code:    400,
			Message: "fail to book requested books",
		}, err
	}

	// get book id from bookDetails
	var bookIDs []string
	for _, v := range bookDetails {
		bookIDs = append(bookIDs, v.BookID)
	}

	//hit db to save book booking (mock)
	bookIDs, partialSuccess, err := Client.InsertBookBooking(bookIDs, f.PickUpSchedule, f.UserID)
	if err != nil && !partialSuccess {
		return form.SaveBookBookingResponse{
			Code:                      500,
			Message:                   "bad request",
			SuccessfullyBookedBookIDs: bookIDs,
		}, err
	}

	// some books already booked
	if err != nil && partialSuccess {
		return form.SaveBookBookingResponse{
			Code:                      202,
			Message:                   "partial success with error: " + err.Error(),
			SuccessfullyBookedBookIDs: bookIDs,
		}, nil
	}

	// all books are successfully booked
	return form.SaveBookBookingResponse{
		Code:                      200,
		Message:                   "success",
		SuccessfullyBookedBookIDs: bookIDs,
	}, nil
}

// GetBooksListBySubject get multiple book details per page
func (b *BookSvc) GetBooksListBySubject(subject string, limit int, page int) (res []form.BooksList, err error) {
	// paging via offset
	offset := 0
	if page > 1 {
		offset = (page - 1) * limit
	}

	var bookData []form.BooksList

	// get books from openLibrary API
	data, err := Client.GetBooksListAPIBySubject(subject, limit, offset)
	for _, v := range data.Data {
		bookID := strings.Split(v.Key, "/")[2]
		bookData = append(bookData, form.BooksList{
			BookID:      bookID,
			BookTitle:   v.Title,
			BookAuthor:  v.Authors[0].Name,
			BookEdition: v.EditionCount,
		})

		// save book details to db (mock)
		global.BooksList[bookID] = form.BooksList{
			BookID:      bookID,
			BookTitle:   v.Title,
			BookAuthor:  v.Authors[0].Name,
			BookEdition: v.EditionCount,
		}
	}

	return bookData, err
}
