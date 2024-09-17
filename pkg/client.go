package pkg

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)


type Client struct {
	Client *http.Client
	Url string
}



var (
	Timeout = 2 * time.Second	
	PortDefault = "8080"
	EndpointDefault = "datetime"
	UrlDefault = "http://localhost:8080/datetime"
)

type options func(*Client)


func (c *Client) WithURL(url string) options{
	return func (*Client)  {
		c.Url = url
	}
}


func (c *Client) LoadConfigFromENV() {
	err := godotenv.Load("../.env")

	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
		return 
	}

	port := os.Getenv("PORT")
	endpoint := os.Getenv("ENDPOINT")

	c.Url = fmt.Sprintf("http://localhost:%s/%s",port,endpoint)
}


func NewClient(opt ...options) *Client{

	client :=  &Client{
		Client : &http.Client{
			Timeout: Timeout,
		},
		Url: UrlDefault,
	}

	for _, o := range opt {
		o(client)
	}

	return client
}