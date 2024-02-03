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
