# Datetime HTTP Client Package 
Create an HTTP client in Go that consumes the datetime server APIs implemented in the previous project. https://github.com/codescalersinternships/home/issues/284

## Installation 
- 1. Download project
```golang
   git get https://github.com/codescalersinternships/Datetime-Server-Doha.git
```
- 2. import package :
```golang
   import httpClient "github.com/dohaelsawy/codescalers/datetime-client/pkg"
```
### Functions
- 1. first have instance from client with default values 
```golang
   client := httpClient.NewClient() // default url-> http://localhost:8080/datatime
```
- 2. to set your values you have 3 options:
     - to load from env file
       ```golang
       err := client.LoadConfigFromENV()
       ```
       use following env valraibles:
       ```
       PORT=8090
       ENDPOINT=/datetime
       ```
     - to pass your flag values in binary file
       ```golang
       make build // build a binary file
       ./client -endpoint=/datatime -port=8090
       client.SetClientUrl(endpoint, port) // to set port and end point into client url
       ```
      - to pass your url with client declaration
        ```golang
        client := httpClient.NewClient(pkg.WithURL(YourUrl))
        ``` 
- 3. gets your data from desire url
     ```golang
     data, err := client.GetResponse()
     ```

### Test
- to run all tests
```golang
  go test v ./...
```
### Format
- format all files inside project
```golang
gofmt -w .
```
