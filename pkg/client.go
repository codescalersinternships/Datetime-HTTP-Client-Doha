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
	Port string
	Endpoint string
}



type Client struct {
	Client *http.Client
	Url string
}



var (
	Timeout = 2 * time.Second

	UrlDefault  = "http://localhost:%s/%s"
	
	PortDefault = "8080"
	
	EndpointDefault = "datetime"
)



func SetConfigDefualt() Config{
	return Config{
		Port: PortDefault,
		Endpoint: EndpointDefault,
	}
}


func LoadConfigFromENV() Config{
	err := godotenv.Load("../.env")

	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
		return Config{}
	}

	port := os.Getenv("PORT")
	endpoint := os.Getenv("ENDPOINT")

	return Config{
		Port: port,
		Endpoint: endpoint,
	}
	
}


func NewClient(config Config) *Client{

	client :=  &Client{
		Client : &http.Client{
			Timeout: Timeout,
		},
		Url: fmt.Sprintf(UrlDefault, config.Port, config.Endpoint),
	}

	return client
}