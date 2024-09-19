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

				w.WriteHeader(http.StatusOK)

				w.Write([]byte(time.Now().Format(time.UnixDate)))
			}),
		)

		client := NewClient(WithURL(mockserver.URL))

		data, err := client.GetResponse()

		expected := time.Now().Format(time.UnixDate)

		if !reflect.DeepEqual(expected,data.DatewTime){
			t.Errorf("errorrrrrrr, shoud your response %s equal to mine %s", data.DatewTime, expected)
		}
		
		if err != nil {
			t.Errorf("errorrrrrrr, there is an erroooorrrrrr%v",err)
		}

	})

	t.Run("Data with plain text format", func(t *testing.T) {
		mockserver := httptest.NewServer(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

				w.Header().Set("Content-Type", "application/json")

				enc := json.NewEncoder(w)

				err := enc.Encode(&DataTimeResponse{
					DatewTime: time.Now().UTC().Format(time.UnixDate),
				})

				if err != nil {
					t.Errorf("%v", err)
				}
			}),
		)

		client := NewClient(WithURL(mockserver.URL))

		data, err := client.GetResponse()

		expected := &DataTimeResponse{
			DatewTime: time.Now().UTC().Format(time.UnixDate),
		}

		if !reflect.DeepEqual(expected.DatewTime,data.DatewTime){
			t.Errorf("errorrrrrrr, shoud your response %s equal to mine %s", data.DatewTime , expected.DatewTime)
		}
		
		if err != nil {
			t.Errorf("errorrrrrrr, there is an erroooorrrrrr%v",err)
		}
	})


	t.Run("invalid url", func(t *testing.T) {

		client := NewClient(WithURL("http://localhost:8080/"))

		_, err := client.GetResponse()

		if err == nil {
			t.Errorf("expect error %v, found nil", err)
		}
		
	})

}
