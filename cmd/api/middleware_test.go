package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"youshare-api.anvo.dev/internal/assert"
)

func TestEnableCORS(t *testing.T) {
	rr := httptest.NewRecorder()

	r, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	app := application{
		config: config{
			env: "development",
		},
		logger: nil,
	}

	app.enableCORS(next).ServeHTTP(rr, r)

	expectedValue := "*"
	assert.Equal(t, expectedValue, rr.Header().Get("Access-Control-Allow-Origin"))
	assert.Equal(t, http.StatusOK, rr.Code)

	body, err := io.ReadAll(rr.Body)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "OK", string(body))
}

func TestAuthenticate(t *testing.T) {
	testCases := []struct {
		name                string
		authorizationHeader string
		expectedStatusCode  int
	}{
		{
			name:                "empty header",
			authorizationHeader: "",
			expectedStatusCode:  http.StatusOK,
		},
		{
			name:                "invalid token",
			authorizationHeader: "Bearer jwttoken",
			expectedStatusCode:  http.StatusUnauthorized,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			rr := httptest.NewRecorder()

			r, err := http.NewRequest("GET", "/", nil)
			if err != nil {
				t.Fatal(err)
			}

			next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte("OK"))
			})

			app := application{
				config: config{
					env: "development",
				},
				logger: nil,
			}

			if testCase.authorizationHeader != "" {
				r.Header.Add("Authorization", testCase.authorizationHeader)
			}

			app.authenticate(next).ServeHTTP(rr, r)

			assert.Equal(t, testCase.expectedStatusCode, rr.Code)
		})
	}
}
