package main

import (
    "fmt"
    "io"
    "net/http"
    "time"
)

var (
    
	Timeout = time.Duration(1) * time.Second
	Url = "http://localhost:8090/datetime"
)

func main() {
    c := http.Client{Timeout: Timeout}
    
	response, err := c.Get(Url)
    
	if err != nil {
        fmt.Printf("Error %s", err)
        return
    }

    defer response.Body.Close()

    body, err := io.ReadAll(response.Body)
	
	if err != nil {
		fmt.Printf("Error %s", err)
        return
	}
    
	fmt.Printf("Body : %s", body)
}