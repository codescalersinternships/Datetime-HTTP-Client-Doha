package pkg

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type Client struct {
	Client *http.Client
	Url    string
}

var (
	Timeout         = 2 * time.Second
	PortDefault     = "8080"
	EndpointDefault = "datetime"
	UrlDefault      = "http://localhost:8080/datetime"
)

type options func(c *Client)

func WithURL(url string) options {
	return func(c *Client) {
		c.Url = url
	}
}

func (c *Client) LoadConfigFromENV() error{
	err := godotenv.Load("../.env")

	if err != nil {
		logrus.Errorf("error while loading .env file. Err: %s", err)
		return err
	}

	port := os.Getenv("PORT")
	endpoint := os.Getenv("ENDPOINT")

	c.Url = fmt.Sprintf("http://localhost:%s%s", port, endpoint)

	logrus.Println("create client with load env file")

	return nil
}

func NewClient(opt ...options) *Client {

	client := &Client{
		Client: &http.Client{
			Timeout: Timeout,
		},
		Url: UrlDefault,
	}

	for _, o := range opt {
		o(client)
	}

	logrus.Printf("new client created %v\n",client)

	return client
}
