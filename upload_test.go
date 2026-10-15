package jmap

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// Any 2xx is a successful upload (RFC 8620 §6.1); servers answer 201 Created.
// This row guards the cherry-pick of 25f7b38d on yarilo-patches.
func TestUploadAcceptsAny2xx(t *testing.T) {
	const body = `{"accountId":"a1","blobId":"B1","type":"message/rfc822","size":3}`
	cases := []struct {
		name    string
		status  int
		ctype   string
		body    string
		wantErr bool
	}{
		{"201 created", http.StatusCreated, "application/json", body, false},
		{"200 ok", http.StatusOK, "application/json", body, false},
		{"304 refused", http.StatusNotModified, "", "", true},
		{"400 problem", http.StatusBadRequest, "application/json", `{"type":"urn:ietf:params:jmap:error:limit","status":400}`, true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if tc.ctype != "" {
					w.Header().Set("Content-Type", tc.ctype)
				}
				w.WriteHeader(tc.status)
				_, _ = w.Write([]byte(tc.body))
			}))
			defer srv.Close()
			c := &Client{
				HttpClient:      srv.Client(),
				SessionEndpoint: srv.URL,
				Session:         &Session{UploadURL: srv.URL + "/upload/{accountId}/"},
			}
			got, err := c.Upload("a1", strings.NewReader("abc"))
			if tc.wantErr {
				if err == nil {
					t.Fatalf("status %d was taken as an upload: %+v", tc.status, got)
				}
				var reqErr *RequestError
				if tc.status == http.StatusBadRequest && !errors.As(err, &reqErr) {
					t.Errorf("a problem body was not decoded: %v", err)
				}
				return
			}
			if err != nil {
				t.Fatalf("status %d refused: %v", tc.status, err)
			}
			if got.ID != "B1" || got.Account != "a1" {
				t.Errorf("upload response %+v", got)
			}
		})
	}
}
