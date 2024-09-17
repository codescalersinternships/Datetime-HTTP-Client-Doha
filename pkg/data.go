package pkg

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

var (
	// ErrInternalServer     = fmt.Errorf("http Status Internal Server Error %d", http.StatusInternalServerError)
	ErrNotSupportedHeader = fmt.Errorf("support only text and json format")
)

type DataTimeResponse struct {
	DatewTime string  `json:"datewtime"`
}

func (c *Client) GetResponse() (DataTimeResponse, error) {

	req, err := http.NewRequestWithContext(context.Background(), "GET", c.Url, nil)

	if err != nil {
		log.Fatalf("error from get response function %s", err)
		return DataTimeResponse{}, err
	}

	req.Header.Add("Accept", "text/plain")
	req.Header.Add("Accept", "application/json")

	res, err := c.Client.Do(req)

	if err != nil {
		// log.Fatalf("error from get response function %s", err)
		return DataTimeResponse{}, err
	}
	defer res.Body.Close()

	// response part ------------

	if res.StatusCode != http.StatusOK {
		// log.Fatal(res.StatusCode)
		return DataTimeResponse{}, fmt.Errorf("Status %d Error", res.StatusCode)
	}

	header := res.Header.Get("Content-Type")

	if !strings.Contains(header, "text/plain") && !strings.Contains(header, "application/json") {
		// log.Fatal(ErrNotSupportedHeader)
		return DataTimeResponse{}, ErrNotSupportedHeader
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		// log.Fatalf("error from get response function %s", err)
		return DataTimeResponse{}, err
	}

	var datetime DataTimeResponse

	if strings.Contains(header, "application/json") {

		err = json.Unmarshal(body, &datetime)

		if err != nil {
			// log.Fatalf("error from get response function %s", err)
			return DataTimeResponse{}, err
		}

		return datetime, nil
	}

	datetime.DatewTime = string(body)

	return datetime, nil
}
