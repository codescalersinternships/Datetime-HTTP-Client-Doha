package main

import (
	"os"
	"github.com/sirupsen/logrus"
)

func main() {

	f, err := os.OpenFile("client.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		logrus.Errorf("while reading client.log file %s\n",err)
	}
	
	defer f.Close()

	logrus.Info("System is starting")
}
