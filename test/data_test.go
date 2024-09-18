package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/dohaelsawy/codescalers/datetime-client/pkg"
	"github.com/stretchr/testify/assert"
)

func TestGetResponse(t *testing.T) {

	t.Run("Data with plain text format", func(t *testing.T) {
		mockserver := httptest.NewServer(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

				w.Header().Set("Content-type", "text/plain")

				w.WriteHeader(http.StatusOK)

				w.Write([]byte(time.Now().Format(time.UnixDate)))
			}),
		)

		client := pkg.NewClient(pkg.WithURL(mockserver.URL))

		data, err := client.GetResponse()

		expected := time.Now().Format(time.UnixDate)

		assert.Equal(t, expected, data.DatewTime)
		assert.NoError(t, err)

	})

	t.Run("Data with plain text format", func(t *testing.T) {
		mockserver := httptest.NewServer(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

				w.Header().Set("Content-Type", "application/json")

				enc := json.NewEncoder(w)

				err := enc.Encode(&pkg.DataTimeResponse{
					DatewTime: time.Now().UTC().Format(time.UnixDate),
				})

				if err != nil {
					t.Errorf("%v", err)
				}
			}),
		)

		client := pkg.NewClient(pkg.WithURL(mockserver.URL))

		data, err := client.GetResponse()

		expected := &pkg.DataTimeResponse{
			DatewTime: time.Now().UTC().Format(time.UnixDate),
		}

		assert.Equal(t, expected.DatewTime, data.DatewTime)
		assert.NoError(t, err)
	})


	t.Run("invalid url", func(t *testing.T) {

		client := pkg.NewClient(pkg.WithURL("http://localhost:8080/"))

		_, err := client.GetResponse()

		assert.NotEqual(t,nil,err)
		
	})

}
