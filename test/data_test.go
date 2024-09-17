package main

// import (
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"
// 	"time"

// 	"github.com/dohaelsawy/codescalers/datetime-client/pkg"
// 	"github.com/stretchr/testify/assert"
// )

// func TestGetResponse(t *testing.T){

// 	t.Run("Data with json format", func(t *testing.T) {
// 		mockserver := httptest.NewServer(
// 			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 				w.Header().Set("Content-type","application/json") 
// 				w.WriteHeader(http.StatusOK)
// 				w.Write([]byte(time.Now().Format(time.UnixDate)))
// 			}),
// 		)

// 		config := pkg.SetConfigDefualt(pkg.WithEndpoint("auhf"))
// 		client := pkg.NewClient(congif)

// 		data , err := client.GetResponse()




		
// 	})
	
// }