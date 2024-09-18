package main

import (
	"reflect"
	"testing"

	pkg "github.com/dohaelsawy/codescalers/datetime-client/pkg"
)

type MockClient struct {
	url string
}

func TestClient(t *testing.T) {

	t.Run("default client configs", func(t *testing.T) {

		client := pkg.NewClient()

		expect := MockClient{
			url: "http://localhost:8080/datetime",
		}

		if !reflect.DeepEqual(expect.url,client.Url){
			t.Errorf("errorrrrrrr, shoud your url %s equal to mine %s", client.Url, expect.url)
		}
	})

	t.Run("client with .env configs", func(t *testing.T) {

		client := pkg.NewClient()
		err := client.LoadConfigFromENV()

		expect := MockClient{
			url: "http://localhost:8090/datetime",
		}

		if !reflect.DeepEqual(expect.url,client.Url){
			t.Errorf("errorrrrrrr, shoud your url %s equal to mine %s", client.Url, expect.url)
		}
		
		if err != nil {
			t.Errorf("errorrrrrrr, there is an erroooorrrrrr%v",err)
		}
	})
}
