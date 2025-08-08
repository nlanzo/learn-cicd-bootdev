package auth

import (
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestGetAPIKey(t *testing.T) {
	tests := map[string]struct {
		input   http.Header
		want    string
		wantErr error
	}{
		"valid API key": {
			input: http.Header{
				"Authorization": []string{"ApiKey 1234567890"},
			},
			want:    "1234567890",
			wantErr: nil,
		},
		"no API key": {
			input:   http.Header{},
			want:    "",
			wantErr: ErrNoAuthHeaderIncluded,
		},
		"malformed API key": {
			input: http.Header{
				"Authorization": []string{"1234567890"},
			},
			want:    "",
			wantErr: ErrMalformedAuthHeader,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got, gotErr := GetAPIKey(test.input)
			diff := cmp.Diff(test.want, got)
			if diff != "" {
				t.Errorf("GetAPIKey() value = %v, want %v", got, test.want)
			}
			diff = cmp.Diff(test.wantErr, gotErr, cmpopts.EquateErrors())
			if diff != "" {
				t.Errorf("GetAPIKey() error = %v, want %v", gotErr, test.wantErr)
			}
		})
	}
}
