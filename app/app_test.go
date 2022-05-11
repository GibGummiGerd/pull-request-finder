package app

import (
	"fmt"
	"pull-request-finder/configuration"
	"pull-request-finder/models"
	"reflect"
	"testing"
)

func TestGetPullRequests(t *testing.T) {
	unmarshallError := fmt.Errorf("json: cannot unmarshal object into Go value of type []models.PullRequest")
	type args struct {
		config configuration.Configuration
	}
	tests := []struct {
		name    string
		args    args
		want    []models.PullRequest
		wantErr error
	}{
		{
			name: "Invalid request returns object which can't be unmarshalled",
			args: args{
				config: configuration.Configuration{},
			},
			want:    []models.PullRequest{},
			wantErr: unmarshallError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetPullRequests(tt.args.config)
			fmt.Println(err)
			if err != nil || tt.wantErr != nil {
				if tt.wantErr == nil {
					t.Fatalf("wanted no error, but got: \n%v", err)
				}
				if err == nil {
					t.Fatalf("got no error but wanted the following: \n%v", err)
				}
				if err.Error() != tt.wantErr.Error() {
					t.Fatalf("expected error: \n%v, \ngot: \n%v", tt.wantErr, err)
				}
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPullRequests() = %v, want %v", got, tt.want)
			}
		})
	}
}
