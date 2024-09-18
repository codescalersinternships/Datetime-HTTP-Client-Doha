package main

import (
	"fmt"
	"os"

	"github.com/dohaelsawy/codescalers/datetime-client/pkg"
	"github.com/sirupsen/logrus"
)

func main() {

	f, err := os.OpenFile("client.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		logrus.Errorf("while reading client.log file %s\n",err)
	}
	
	defer f.Close()
	logrus.SetOutput(f)

	logrus.Info("System is starting")

	client := pkg.NewClient(pkg.WithURL("http://localhost:8080/datetime"))

	data , err := client.GetResponse()

	if err != nil {
		logrus.Error(err)
	}

	fmt.Println(data)
}
