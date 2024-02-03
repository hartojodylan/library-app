package books

import (
	"bytes"
	"github.com/dylanh/library-app/global"
	"github.com/dylanh/library-app/model/form"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"testing"
)

func Test_GetBookDetailsFromDB(t *testing.T) {
	type args struct {
		bookIDs []string
	}
	tests := []struct {
		name    string
		args    args
		want    []form.BooksList
		wantErr bool
		mock    func()
	}{
		{
			name: "no books exist",
			args: args{
				bookIDs: []string{"1", "2"},
			},
			want:    nil,
			wantErr: true,
			mock: func() {
				global.BooksList = map[string]form.BooksList{}
			},
		},
		{
			name: "success",
			args: args{
				bookIDs: []string{"1", "2"},
			},
			want: []form.BooksList{
				{
					BookID: "1",
				},
				{
					BookID: "2",
				},
			},
			wantErr: false,
			mock: func() {
				global.BooksList = map[string]form.BooksList{
					"1": {
						BookID: "1",
					},
					"2": {
						BookID: "2",
					},
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			r := &BookCore{}

			got, err := r.GetBookDetailsFromDB(tt.args.bookIDs)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBookDetailsFromDB() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equalf(t, tt.want, got, "GetBookDetailsFromDB(%v)", tt.args.bookIDs)
		})
	}
}

func Test_InsertBookBooking(t *testing.T) {
	type args struct {
		bookIDs        []string
		pickUpSchedule string
		userID         int64
	}
	type res struct {
		successBookIDs []string
		partialSuccess bool
	}
	tests := []struct {
		name    string
		args    args
		want    res
		wantErr bool
		mock    func()
	}{
		{
			name: "all books requested are booked",
			args: args{
				bookIDs:        []string{"1", "2"},
				pickUpSchedule: "2000-12-12 00:00:00",
				userID:         1,
			},
			want:    res{},
			wantErr: true,
			mock: func() {
				global.BooksBooking = map[string]form.BookBooking{
					"1": {
						BookID: "1",
					},
					"2": {
						BookID: "2",
					},
				}
			},
		},
		{
			name: "success partial, book 1 is booked",
			args: args{
				bookIDs:        []string{"1", "2"},
				pickUpSchedule: "2000-12-12 00:00:00",
				userID:         1,
			},
			want: res{
				successBookIDs: []string{"2"},
				partialSuccess: true,
			},
			wantErr: true,
			mock: func() {
				global.BooksBooking = map[string]form.BookBooking{
					"1": {
						BookID: "1",
					},
				}
			},
		},
		{
			name: "success",
			args: args{
				bookIDs:        []string{"1", "2"},
				pickUpSchedule: "2000-12-12 00:00:00",
				userID:         1,
			},
			want: res{
				successBookIDs: []string{"1", "2"},
				partialSuccess: false,
			},
			wantErr: false,
			mock: func() {
				global.BooksBooking = map[string]form.BookBooking{}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			r := &BookCore{}

			got, partialSuccess, err := r.InsertBookBooking(tt.args.bookIDs, tt.args.pickUpSchedule, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("InsertBookBooking() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equalf(t, tt.want, res{
				successBookIDs: got,
				partialSuccess: partialSuccess,
			}, "InsertBookBooking(%v)", tt.args.bookIDs)
		})
	}
}

func Test_GetBooksListAPIBySubject(t *testing.T) {
	// mock http
	Client = &ClientMock{}

	type args struct {
		subject string
		limit   int
		offset  int
	}
	tests := []struct {
		name    string
		args    args
		want    form.GetBooksListApiResponse
		wantErr bool
		mock    func()
	}{
		{
			name: "http error",
			args: args{
				subject: "love",
				limit:   10,
				offset:  0,
			},
			want:    form.GetBooksListApiResponse{},
			wantErr: true,
			mock: func() {
				mockDoFunc = func(req *http.Request) (*http.Response, error) {
					return nil, assert.AnError
				}
			},
		},
		{
			name: "json decode error",
			args: args{
				subject: "love",
				limit:   10,
				offset:  0,
			},
			want:    form.GetBooksListApiResponse{},
			wantErr: true,
			mock: func() {
				mockDoFunc = func(req *http.Request) (*http.Response, error) {
					invalidResponse := io.NopCloser(bytes.NewReader([]byte("<invalid json>")))
					return &http.Response{
						Body: invalidResponse,
					}, nil
				}
			},
		},
		{
			name: "success",
			args: args{
				subject: "love",
				limit:   10,
				offset:  0,
			},
			want: form.GetBooksListApiResponse{
				Data: []form.Works{
					{
						Key:          "/books/OL1M",
						Title:        "Book1",
						EditionCount: 1,
						Authors: []form.Authors{
							{Name: "Author1"},
						},
					},
				},
			},
			wantErr: false,
			mock: func() {
				mockDoFunc = func(req *http.Request) (*http.Response, error) {
					validResponse := io.NopCloser(bytes.NewReader([]byte(`{"works": [{"key": "/books/OL1M", "title": "Book1", "authors": [{"name": "Author1"}], "edition_count": 1}]}`)))
					return &http.Response{
						StatusCode: 200,
						Body:       validResponse,
					}, nil
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			r := &BookCore{}

			got, err := r.GetBooksListAPIBySubject(tt.args.subject, tt.args.limit, tt.args.offset)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBooksListAPIBySubject() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equalf(t, tt.want, got, "GetBooksListAPIBySubject(%v)", tt.args.subject)
		})
	}
}
