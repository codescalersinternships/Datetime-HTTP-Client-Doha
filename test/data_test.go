package main

import (
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

		assert.Equal(t, expected, data.DatewaTime)
		assert.NoError(t, err)

	})

}
