package pkg

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)


type Config struct {
	port string
	endpoint string
}


type Client struct {
	client *http.Client
	url string
}

var (
	Timeout = 2 * time.Second

	UrlDefualt  = "http://localhost:%s/%s"
	
	PortDefault = "8080"
	
	EndpointDefault = "datetime"
)

func SetConfigDefualt() *Config{
	return &Config{
		port: PortDefault,
		endpoint: EndpointDefault,
	}
}


func LoadConfigFromENV() *Config{
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
		return &Config{}
	}

	port := os.Getenv("PORT")
	endpoint := os.Getenv("ENDPOINT")

	return &Config{
		port: port,
		endpoint: endpoint,
	}
	
}


func NewCLient(config Config) *Client{
	return &Client{
		client : &http.Client{
			Timeout: Timeout,
		},
		url: fmt.Sprintf(UrlDefualt, config.port, config.endpoint),
	}
}