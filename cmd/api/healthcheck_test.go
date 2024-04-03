package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"youshare-api.anvo.dev/internal/assert"
)

func TestHealthCheck(t *testing.T) {
	rr := httptest.NewRecorder()

	r, err := http.NewRequest("GET", "/v1/healthcheck", nil)
	if err != nil {
		t.Fatal(err)
	}

	app := application{
		config: config{
			env: "development",
		},
		logger: nil,
	}

	healthcheckHandler := http.HandlerFunc(app.healthcheckHandler)
	healthcheckHandler(rr, r)

	rs := rr.Result()
	assert.Equal(t, http.StatusOK, rs.StatusCode)
	defer rs.Body.Close()

	body, err := io.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}

	var data map[string]any
	err = json.Unmarshal(body, &data)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "available", data["status"])
	assert.Equal(t, "development", data["system_info"].(map[string]any)["environment"])
	assert.Equal(t, version, data["system_info"].(map[string]any)["version"])
}
