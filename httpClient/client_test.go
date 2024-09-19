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

		if !reflect.DeepEqual(expect.url,client.Url){
			t.Errorf("errorrrrrrr, shoud your url %s equal to mine %s", client.Url, expect.url)
		}
	})

	t.Run("client with .env configs", func(t *testing.T) {

		client := NewClient()
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
