package main

import (
	"flag"
	"fmt"
	"os"

	client "github.com/codescalersinternships/Datetime-HTTP-Client-Doha/httpClient"
	"github.com/sirupsen/logrus"
)

func main() {

	f, err := os.OpenFile("client.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		logrus.Errorf("while reading client.log file %s\n", err)
	}

	defer f.Close()
	logrus.SetOutput(f)

	logrus.Info("System is starting")

	var port string
	var endpoint string

	flag.StringVar(&port, "port", "8080", "the number of url port")
	flag.StringVar(&endpoint, "endpoint", "/datatime", "the end point in url")
	flag.Parse()

	client := client.NewClient()

	client.SetClientUrl(endpoint, port)

	data, err := client.GetResponse()

	if err != nil {
		logrus.Error(err)
	}

	fmt.Println(data)
}
