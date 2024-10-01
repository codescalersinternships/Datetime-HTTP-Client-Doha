// Package Create an HTTP client in Go that consumes the datetime server APIs implemented in [Datetime Saerver](codescalersinternships/home#284)

package httpClient

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"time"
)

// Client represent a http client
type Client struct {
	Client *http.Client
	Url    string
}

// Generic http client for default values
const (
	TimeoutDefault  = 2 * time.Second
	PortDefault     = "8080"
	EndpointDefault = "datetime"
	UrlDefault      = "http://localhost:8080/datetime"
)

// options represent an option for http client
type options func(c *Client)

// WithURL set url option while initalizing a new client
func WithURL(url string) options {
	return func(c *Client) {
		c.Url = url
	}
}

// WithEndpointAndPort set endpoint and port option while initalizing a new client
func WithEndpointAndPort(endpoint, port string) options {
	return func(c *Client) {
		c.Url = makeUrl(endpoint,port)
	}
}


// LoadConfigFromENV load port and endpoint from env file and return error if exist
func (c *Client) LoadConfigFromENV(path string) error {
	err := godotenv.Load(path)

	if err != nil {
		logrus.Errorf("error while loading .env file. Err: %s", err)
		return err
	}

	port := os.Getenv("PORT")
	endpoint := os.Getenv("ENDPOINT")

	c.Url = makeUrl(endpoint, port)

	logrus.Println("create client with load env file")

	return nil
}

// NewClient initalize new http client and take option url string
func NewClient(opt ...options) *Client {

	client := &Client{
		Client: &http.Client{
			Timeout: TimeoutDefault,
		},
		Url: UrlDefault,
	}

	for _, o := range opt {
		o(client)
	}

	logrus.Printf("new client created %v\n", client)

	return client
}

func makeUrl(endpoint, port string) string {
	return fmt.Sprintf("http://localhost:%s%s", port, endpoint)
}
