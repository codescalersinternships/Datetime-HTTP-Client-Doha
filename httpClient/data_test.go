package httpClient

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"
)

func TestGetResponse(t *testing.T) {

	t.Run("Data with plain text format", func(t *testing.T) {
		mockserver := httptest.NewServer(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

				w.Header().Set("Content-type", "text/plain")

				w.WriteHeader(http.StatusInternalServerError)

				_ , _ = w.Write([]byte(time.Now().Format(time.UnixDate)))
			}),
		)

		client := NewClient(WithURL(mockserver.URL))

		data, err := client.GetDateTime()

		expected := time.Now().Format(time.UnixDate)

		if !reflect.DeepEqual(expected, data.DateTime) {
			t.Errorf("expected %v, got %v", expected, data.DateTime)
		}

		if err != nil {
			t.Errorf("%v", err)
		}

	})

	t.Run("Data with plain text format", func(t *testing.T) {
		mockserver := httptest.NewServer(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

				w.Header().Set("Content-Type", "application/json")

				enc := json.NewEncoder(w)

				err := enc.Encode(&DataTimeResponse{
					DateTime: time.Now().UTC().Format(time.UnixDate),
				})

				if err != nil {
					t.Errorf("%v", err)
				}
			}),
		)

		client := NewClient(WithURL(mockserver.URL))

		data, err := client.GetDateTime()

		expected := &DataTimeResponse{
			DateTime: time.Now().UTC().Format(time.UnixDate),
		}

		if !reflect.DeepEqual(expected.DateTime, data.DateTime) {
			t.Errorf("expected %v, got %v", expected.DateTime, data.DateTime)
		}

		if err != nil {
			t.Errorf("%v", err)
		}
	})

	t.Run("invalid url", func(t *testing.T) {

		client := NewClient(WithURL("http://localhost:8080/"))

		_, err := client.GetDateTime()

		if err == nil {
			t.Errorf("error is %v", err)
		}

	})

}
