package books

import (
	"encoding/json"
	"fmt"
	"github.com/dylanh/library-app/global"
	"github.com/dylanh/library-app/model/form"
)

func (b *BookCore) GetBookDetailsFromDB(bookIDs []string) (res []form.BooksList, err error) {
	for _, v := range bookIDs {
		// in sql this should be a query to get book details, missing books are ignored
		if global.BooksList[v] != (form.BooksList{}) {
			res = append(res, global.BooksList[v])
		}
	}

	if len(res) == 0 {
		return res, fmt.Errorf("no books exist with the given IDs")
	}

	return res, nil
}

func (b *BookCore) InsertBookBooking(bookIDs []string, pickUpSchedule string, userID int64) (successBookIDs []string, partialSuccess bool, err error) {
	var bookedBookIDs []string
	var successfullyBookedBookIDs []string

	for _, v := range bookIDs {
		//hit db to save book booking (mock)
		if global.BooksBooking[v] == (form.BookBooking{}) {
			global.BooksBooking[v] = form.BookBooking{
				BookID:         v,
				PickUpSchedule: pickUpSchedule,
				UserID:         userID,
			}
			successfullyBookedBookIDs = append(successfullyBookedBookIDs, v)
		} else {
			// if book is already booked, append to slice
			bookedBookIDs = append(bookedBookIDs, v)
		}
	}

	// if all requested books are booked
	if len(bookedBookIDs) == len(bookIDs) {
		return successfullyBookedBookIDs, false, fmt.Errorf("BookID: %v is already booked", bookedBookIDs)
	}

	// if there are some books that are already booked (partial success)
	if len(bookedBookIDs) > 0 {
		return successfullyBookedBookIDs, true, fmt.Errorf("BookID: %v is already booked", bookedBookIDs)
	}

	// success
	return successfullyBookedBookIDs, false, nil
}

func (b *BookCore) GetBooksListAPIBySubject(subject string, limit int, offset int) (res form.GetBooksListApiResponse, err error) {
	// hit GET http://openlibrary.org/subjects/love.json to get book details
	response, err := Get("http://openlibrary.org/subjects/" + subject + ".json" + "?limit=" + fmt.Sprintf("%v", limit) + "&offset=" + fmt.Sprintf("%v", offset))
	if err != nil {
		return res, err
	}

	// decode response
	err = json.NewDecoder(response.Body).Decode(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}
