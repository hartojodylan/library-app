package global

import "github.com/dylanh/library-app/model/form"

// Global variables are located here
var (
	// These global variables will act as a database

	// BookID is map key
	BooksBooking = make(map[string]form.BookBooking)
	// BookID is map key
	BooksList = make(map[string]form.BooksList)
)
