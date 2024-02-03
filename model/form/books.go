package form

// HTTP Form Section

// SaveBookBookingRequest create new user
type SaveBookBookingRequest struct {
	// BookID to be booked
	BookID []string `json:"book-id" form:"book-id" validate:"required" example:"harry potter"`
	// Pick up schedule
	PickUpSchedule string `json:"pick-up-schedule" form:"pick-up-schedule" validate:"required" example:"12-12-2012 12:12:12"`
	// UserID who's booking
	UserID int64 `json:"user-id" form:"user-id" validate:"required"`
}

// SaveBookBookingResponse general http response
type SaveBookBookingResponse struct {
	Code                      int      `json:"code"`
	Message                   string   `json:"message"`
	SuccessfullyBookedBookIDs []string `json:"data"`
}

type BooksList struct {
	BookID      string
	BookTitle   string
	BookAuthor  string
	BookEdition int64
}

type BookBooking struct {
	BookID         string
	PickUpSchedule string
	UserID         int64
}

type GetBooksListBySubjectResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    []BooksList `json:"data"`
}

type GetBooksListApiResponse struct {
	Data []Works `json:"works"`
}

type Works struct {
	Key          string    `json:"key"`
	Title        string    `json:"title"`
	EditionCount int64     `json:"edition_count"`
	Authors      []Authors `json:"authors"`
}

type Authors struct {
	Name string `json:"name"`
	Key  string `json:"key"`
}
