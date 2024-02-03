package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/dylanh/library-app/model/form"
	"github.com/gookit/rux"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_GetBooksListBySubject(t *testing.T) {
	type args struct {
		subject string
		limit   string
		page    string
	}

	emptyResSlice := make([]form.BooksList, 0)

	type response struct {
		Code    int              `json:"code"`
		Message string           `json:"message"`
		Data    []form.BooksList `json:"data"`
	}

	tests := []struct {
		name     string
		mock     func()
		args     args
		wantCode int
		want     []form.BooksList
	}{
		{
			name: "error request param",
			mock: func() {
				mockGetBooksListBySubject = func(subject string, limit int, page int) (res []form.BooksList, err error) {
					return emptyResSlice, fmt.Errorf("test error")
				}
			},
			args: args{
				subject: "love",
				limit:   "10",
				page:    "aaaaaaaa",
			},
			wantCode: 400,
			want:     emptyResSlice,
		},
		{
			name: "error GetBooksListBySubject",
			mock: func() {
				mockGetBooksListBySubject = func(subject string, limit int, page int) (res []form.BooksList, err error) {
					return emptyResSlice, fmt.Errorf("test error")
				}
			},
			args: args{
				subject: "love",
				limit:   "10",
				page:    "1",
			},
			wantCode: 500,
			want:     emptyResSlice,
		},
		{
			name: "success",
			mock: func() {
				mockGetBooksListBySubject = func(subject string, limit int, page int) (res []form.BooksList, err error) {
					return []form.BooksList{
						{
							BookID:      "1",
							BookTitle:   "Book 1",
							BookAuthor:  "Author 1",
							BookEdition: 1,
						},
					}, nil
				}
			},
			args: args{
				subject: "love",
				limit:   "10",
				page:    "1",
			},
			wantCode: 200,
			want: []form.BooksList{
				{
					BookID:      "1",
					BookTitle:   "Book 1",
					BookAuthor:  "Author 1",
					BookEdition: 1,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Client = &ClientMock{}
			u := &BooksApi{}
			tt.mock()

			router := rux.New()
			router.GET("/v1/books/{subject}", u.GetBooksListBySubject)
			recorder := httptest.NewRecorder()
			request, _ := http.NewRequest("GET", "/v1/books/"+tt.args.subject+"?limit="+tt.args.limit+"&page="+tt.args.page, nil)
			router.ServeHTTP(recorder, request)

			assert.Equal(t, tt.wantCode, recorder.Code, "error code")

			var r response
			want := make([]form.BooksList, 0)
			err := json.Unmarshal(recorder.Body.Bytes(), &r)
			assert.NoError(t, err, "unmarshall no error")
			if len(tt.want) > 0 {
				wantData, _ := json.Marshal(tt.want)
				err := json.Unmarshal(wantData, &want)
				assert.NoError(t, err, "unmarshall no error")
			}
			assert.Equal(t, want, r.Data, "handler response")
		})
	}
}

func Test_SaveBookBooking(t *testing.T) {
	type args struct {
		BookID         []string `json:"book-id"`
		PickUpSchedule string   `json:"pick-up-schedule"`
		UserID         int64    `json:"user-id"`
	}

	tests := []struct {
		name     string
		mock     func()
		args     args
		wantCode int
		want     form.SaveBookBookingResponse
	}{
		{
			name: "error request param",
			mock: func() {
				mockSaveBookBookingFunc = func(f *form.SaveBookBookingRequest) (res form.SaveBookBookingResponse, err error) {
					return form.SaveBookBookingResponse{}, fmt.Errorf("test error")
				}
			},
			args: args{
				BookID: []string{},
			},
			wantCode: 400,
			want: form.SaveBookBookingResponse{
				Code:                      400,
				Message:                   "error param",
				SuccessfullyBookedBookIDs: make([]string, 0),
			},
		},
		{
			name: "error SaveBookBooking",
			mock: func() {
				mockSaveBookBookingFunc = func(f *form.SaveBookBookingRequest) (res form.SaveBookBookingResponse, err error) {
					return form.SaveBookBookingResponse{
						Code:                      500,
						Message:                   "fail to book requested books",
						SuccessfullyBookedBookIDs: make([]string, 0),
					}, fmt.Errorf("test error")
				}
			},
			args: args{
				BookID:         []string{"1"},
				PickUpSchedule: "2021-01-01 12:12:12",
				UserID:         1,
			},
			wantCode: 500,
			want: form.SaveBookBookingResponse{
				Code:                      500,
				Message:                   "fail to book requested books",
				SuccessfullyBookedBookIDs: make([]string, 0),
			},
		},
		{
			name: "success",
			mock: func() {
				mockSaveBookBookingFunc = func(f *form.SaveBookBookingRequest) (res form.SaveBookBookingResponse, err error) {
					return form.SaveBookBookingResponse{
						Code:                      200,
						Message:                   "success",
						SuccessfullyBookedBookIDs: []string{"1"},
					}, nil
				}
			},
			args: args{
				BookID:         []string{"1"},
				PickUpSchedule: "2021-01-01 12:12:12",
				UserID:         1,
			},
			wantCode: 200,
			want: form.SaveBookBookingResponse{
				Code:                      200,
				Message:                   "success",
				SuccessfullyBookedBookIDs: []string{"1"},
			},
		},
		{
			name: "partial success",
			mock: func() {
				mockSaveBookBookingFunc = func(f *form.SaveBookBookingRequest) (res form.SaveBookBookingResponse, err error) {
					return form.SaveBookBookingResponse{
						Code:                      202,
						Message:                   "partial success with error: test error",
						SuccessfullyBookedBookIDs: []string{"1"},
					}, nil
				}
			},
			args: args{
				BookID:         []string{"1", "2"},
				PickUpSchedule: "2021-01-01 12:12:12",
				UserID:         1,
			},
			wantCode: 202,
			want: form.SaveBookBookingResponse{
				Code:                      202,
				Message:                   "partial success with error: test error",
				SuccessfullyBookedBookIDs: []string{"1"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Client = &ClientMock{}
			u := &BooksApi{}
			tt.mock()

			// request body
			b := args{
				BookID:         tt.args.BookID,
				PickUpSchedule: tt.args.PickUpSchedule,
				UserID:         tt.args.UserID,
			}
			var buf bytes.Buffer
			err := json.NewEncoder(&buf).Encode(b)
			if err != nil {
				log.Fatal(err)
			}

			router := rux.New()
			router.POST("/v1/books", u.SaveBookBooking)
			recorder := httptest.NewRecorder()
			request, _ := http.NewRequest("POST", "/v1/books", &buf)
			request.Header.Set("Content-Type", "application/json")
			router.ServeHTTP(recorder, request)

			assert.Equal(t, tt.wantCode, recorder.Code, "error code")

			var r form.SaveBookBookingResponse
			want := form.SaveBookBookingResponse{}
			fmt.Println(string(recorder.Body.Bytes()[:]))
			err = json.Unmarshal(recorder.Body.Bytes(), &r)
			assert.NoError(t, err, "unmarshall no error")

			if tt.want.SuccessfullyBookedBookIDs != nil {
				wantData, _ := json.Marshal(tt.want)
				err := json.Unmarshal(wantData, &want)
				assert.NoError(t, err, "unmarshall no error")
			}

			assert.Equal(t, want, r, "handler response")
		})
	}
}
