package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlerStatusCode(t *testing.T) {
	testCases := []struct {
		desc     string
		endpoint string
		want     int
	}{
		{
			desc:     "Endpoint '/' returns 200",
			endpoint: "/",
			want:     http.StatusOK,
		},
		{
			desc:     "Endpoint '/health' returns 200",
			endpoint: "/health",
			want:     http.StatusOK,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, tC.endpoint, nil)
			write := httptest.NewRecorder()
			handler := ServerMux()
			handler.ServeHTTP(write, req)
			response := write.Result()
			got := response.StatusCode
			if got != tC.want {
				t.Errorf("GET %s got StatusCode %d; want StatusCode %d", tC.endpoint, got, tC.want)
			}
		})
	}
}

func TestHandlerResponseBody(t *testing.T) {
	testCases := []struct {
		desc     string
		endpoint string
		want     Response
	}{
		{
			desc:     "Endpoint '/' returns correct response body",
			endpoint: "/",
			want:     Response{http.StatusOK, "Hello, World"},
		},
		{
			desc:     "Endpoint '/health' returns correct response body",
			endpoint: "/health",
			want:     Response{http.StatusOK, "Service is healthy"},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, tC.endpoint, nil)
			write := httptest.NewRecorder()
			handler := ServerMux()
			handler.ServeHTTP(write, req)

			response := write.Result()

			var got Response
			err := json.NewDecoder(response.Body).Decode(&got)
			if err != nil {
				t.Errorf("Failed to decode response body %s, err %s", response.Body, err)
			}

			if got != tC.want {
				t.Errorf("GET %s got response body %+v; want response body %+v", tC.endpoint, got, tC.want)
			}
		})
	}
}
