package main

import (
	"fmt"
	"net/http"

	"github.com/dohaelsawy/codescalers/datetime-client/pkg"
)




func main(){

	// req, _ := http.NewRequest("GET","http://localhost:8080/datetime",nil)


	config := pkg.SetConfigDefualt()
	c := pkg.NewClient(config)
	http.NewServeMux()

	data , _ := c.GetResponse()

	fmt.Println(data)
}