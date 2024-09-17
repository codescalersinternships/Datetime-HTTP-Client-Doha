package main

import (
	"testing"

	pkg "github.com/dohaelsawy/codescalers/datetime-client/pkg"
	"github.com/go-playground/assert/v2"
)


type MockClient struct {
	url string
}



func TestClient(t *testing.T){

	t.Run("default client configs", func(t *testing.T) {
		mconfig := pkg.SetConfigDefualt()

		client := pkg.NewClient(mconfig)

		expect :=  MockClient{
			url: "http://localhost:8080/datetime",
		}

		assert.Equal(t, expect.url, client.Url)
	})

	t.Run("client with .env configs", func(t *testing.T) {
		mconfig := pkg.LoadConfigFromENV()

		client := pkg.NewClient(mconfig)

		expect:= MockClient{
			url: "http://localhost:8090/datetime",
		}

		assert.Equal(t, expect.url, client.Url)
	})
}