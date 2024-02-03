package books

import (
	"fmt"
	"github.com/dylanh/library-app/model/form"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_SaveBookBooking(t *testing.T) {
	Client = &ClientMock{}

	type args struct {
		request form.SaveBookBookingRequest
	}
	tests := []struct {
		name    string
		args    args
		want    form.SaveBookBookingResponse
		wantErr bool
		mock    func()
	}{
		{
			name: "no books exist",
			args: args{
				request: form.SaveBookBookingRequest{
					BookID:         []string{"1"},
					PickUpSchedule: "2021-01-01",
					UserID:         1,
				},
			},
			want: form.SaveBookBookingResponse{
				Code:    400,
				Message: "fail to book requested books",
			},
			wantErr: true,
			mock: func() {
				mockGetBookDetailsFromDBFunc = func(bookIDs []string) (res []form.BooksList, err error) {
					return nil, fmt.Errorf("test error")
				}
			},
		},
		{
			name: "success",
			args: args{
				request: form.SaveBookBookingRequest{
					BookID:         []string{"1"},
					PickUpSchedule: "2021-01-01",
					UserID:         1,
				},
			},
			want: form.SaveBookBookingResponse{
				Code:                      200,
				Message:                   "success",
				SuccessfullyBookedBookIDs: []string{"1"},
			},
			wantErr: false,
			mock: func() {
				mockGetBookDetailsFromDBFunc = func(bookIDs []string) (res []form.BooksList, err error) {
					return []form.BooksList{
						{
							BookID: "1",
						},
					}, nil
				}
				mockInsertBookBookingFunc = func(bookIDs []string, pickUpSchedule string, userID int64) (successBookIDs []string, partialSuccess bool, err error) {
					return []string{"1"}, false, nil
				}
			},
		},
		{
			name: "partial success",
			args: args{
				request: form.SaveBookBookingRequest{
					BookID:         []string{"1", "2"},
					PickUpSchedule: "2021-01-01",
					UserID:         1,
				},
			},
			want: form.SaveBookBookingResponse{
				Code:                      202,
				Message:                   "partial success with error: test error",
				SuccessfullyBookedBookIDs: []string{"1"},
			},
			wantErr: false,
			mock: func() {
				mockGetBookDetailsFromDBFunc = func(bookIDs []string) (res []form.BooksList, err error) {
					return []form.BooksList{
						{
							BookID: "1",
						},
					}, nil
				}
				mockInsertBookBookingFunc = func(bookIDs []string, pickUpSchedule string, userID int64) (successBookIDs []string, partialSuccess bool, err error) {
					return []string{"1"}, true, fmt.Errorf("test error")
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := SaveBookBooking(&tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("SaveBookBooking() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equalf(t, tt.want, got, "SaveBookBooking(%v)", tt.args.request)
		})
	}
}

func Test_GetBooksListBySubject(t *testing.T) {
	Client = &ClientMock{}

	type args struct {
		subject string
		limit   int
		page    int
	}
	tests := []struct {
		name    string
		args    args
		want    []form.BooksList
		wantErr bool
		mock    func()
	}{
		{
			name: "success",
			args: args{
				subject: "love",
				limit:   10,
				page:    1,
			},
			want: []form.BooksList{
				{
					BookID:      "1",
					BookTitle:   "Book 1",
					BookAuthor:  "Author 1",
					BookEdition: 1,
				},
			},
			wantErr: false,
			mock: func() {
				mockGetBooksListAPIBySubject = func(subject string, limit int, offset int) (res form.GetBooksListApiResponse, err error) {
					return form.GetBooksListApiResponse{
						Data: []form.Works{
							{
								Key:   "a/b/1",
								Title: "Book 1",
								Authors: []form.Authors{
									{
										Name: "Author 1",
									},
								},
								EditionCount: 1,
							},
						},
					}, nil
				}
			},
		},
		{
			name: "http error",
			args: args{
				subject: "love",
				limit:   10,
				page:    1,
			},
			want:    nil,
			wantErr: true,
			mock: func() {
				mockGetBooksListAPIBySubject = func(subject string, limit int, offset int) (res form.GetBooksListApiResponse, err error) {
					return form.GetBooksListApiResponse{}, fmt.Errorf("test error")
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := GetBooksListBySubject(tt.args.subject, tt.args.limit, tt.args.page)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBooksListBySubject() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equalf(t, tt.want, got, "GetBooksListBySubject(%v)", tt.args.subject)
		})
	}
}
