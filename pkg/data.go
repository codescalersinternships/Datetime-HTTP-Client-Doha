package pkg

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/sirupsen/logrus"
)

var (
	ErrInternalServer     = fmt.Errorf("http Status Internal Server Error %d \n", http.StatusInternalServerError)
	ErrNotFound     = fmt.Errorf("http Not Found Error %d\n", http.StatusNotFound)
	ErrNotSupportedHeader = fmt.Errorf("support only text and json format\n")
	
)

type DataTimeResponse struct {
	DatewTime string  `json:"datewtime"`
}

func (c *Client) GetResponse() (DataTimeResponse, error) {

	req, err := http.NewRequestWithContext(context.Background(), "GET", c.Url, nil)

	if err != nil {
		logrus.Errorf("from http.NewRequestWithContext function %s\n", err)
		return DataTimeResponse{}, err
	}

	req.Header.Add("Accept", "text/plain")
	req.Header.Add("Accept", "application/json")

	res, err := c.Client.Do(req)

	if err != nil {
		logrus.Errorf("from Client.Do function %s\n", err)
		return DataTimeResponse{}, err
	}
	defer res.Body.Close()

	// response part ------------

	if res.StatusCode != http.StatusOK {
		logrus.Errorf("Response status is not 200 %d",res.StatusCode)
		return DataTimeResponse{},fmt.Errorf(fmt.Sprintf("status %d Error",res.StatusCode))
	}

	header := res.Header.Get("Content-Type")

	if !strings.Contains(header, "text/plain") && !strings.Contains(header, "application/json") {
		logrus.Error(ErrNotSupportedHeader)
		return DataTimeResponse{}, ErrNotSupportedHeader
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		logrus.Errorf("can't read response body. Err = %s", err)
		return DataTimeResponse{}, err
	}

	var datetime DataTimeResponse

	if strings.Contains(header, "application/json") {

		err = json.Unmarshal(body, &datetime)

		if err != nil {
			logrus.Errorf("unable to unmarchal body to json. Err = %s", err)
			return DataTimeResponse{}, err
		}

		return datetime, nil
	}

	datetime.DatewTime = string(body)

	return datetime, nil
}
