package httpClient

import (
	"reflect"
	"testing"
)

type MockClient struct {
	url string
}

func TestClient(t *testing.T) {

	t.Run("default client configs", func(t *testing.T) {

		client := NewClient()

		expect := MockClient{
			url: "http://localhost:8080/datetime",
		}

		if !reflect.DeepEqual(expect.url, client.Url) {
			t.Errorf("expected %v, got %v", expect.url, client.Url)
		}
	})

	t.Run("client with .env configs", func(t *testing.T) {

		client := NewClient()
		err := client.LoadConfigFromENV("../.env")

		expect := MockClient{
			url: "http://localhost:8090/datetime",
		}

		if !reflect.DeepEqual(expect.url, client.Url) {
			t.Errorf("expected %v, got %v", expect.url, client.Url)
		}

		if err != nil {
			t.Errorf("error is %v", err)
		}
	})
}
