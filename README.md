# Datetime HTTP Client Package 
Create an HTTP client in Go that consumes the datetime server APIs implemented in the previous project. https://github.com/codescalersinternships/home/issues/284

## Installation 
- 1. Clone project
```golang
   git clone https://github.com/codescalersinternships/Datetime-Server-Doha.git
```
- 2. import package :
```golang
   import "github.com/dohaelsawy/codescalers/datetime-client/pkg"
```
### functions
- 1. first have instance from client with defaults values 
```golang
   client := pkg.NewClient() // default url-> http://localhost:8080/datatime
```
- 2. to set your values you have 2 options
     - to load from env file
       ```golang
       err := client.LoadConfigFromENV()
       ```
     - to pass your flag values in binary file
       ```golang
       make build // build a binary file
       ./client -endpoint=/datatime -port=8090
       ```
- 3. gets your data from desire url
     ```golang
     data, err := client.GetResponse()
     ```

### test
- to run all tests
```golang
  go test v ./...
```
### format
- format all files inside project
```golang
gofmt -w .
```
