package pkg

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/sirupsen/logrus"
)

var (
	ConstBackoffTime = 4 * time.Second
)

type DataTimeResponse struct {
	DatewTime string `json:"datewtime"`
}

type ErrResponse struct {
	Err        error `json:"error"`
	StatusCode int   `json:"statuscode"`
}

func (e ErrResponse) Error() string {
	return fmt.Sprintf("error is %s , and http status code is %d", e.Err, e.StatusCode)
}

func AssignErrorResponse(err error, statuscode int) error {
	return ErrResponse{
		Err:        err,
		StatusCode: statuscode,
	}
}

func (c *Client) GetResponse() (DataTimeResponse, error) {

	req, err := http.NewRequest("GET", c.Url, nil)

	if err != nil {
		logrus.Errorf("from http.NewRequestWithContext function %s\n", err)
		return DataTimeResponse{}, AssignErrorResponse(err, req.Response.StatusCode)
	}

	req.Header.Add("Accept", "text/plain")
	req.Header.Add("Accept", "application/json")

	operation := func() (*http.Response, error) {
		res, err := c.Client.Do(req)
		return res, err
	}

	notify := func(err error, t time.Duration) {
		logrus.Printf("error: %v happened at time: %v", err, t)
	}

	b := backoff.NewConstantBackOff(ConstBackoffTime)

	res, err := backoff.RetryNotifyWithData(operation, b, notify)

	if err != nil {
		logrus.Errorf("from Client.Do function %s\n", err)
		return DataTimeResponse{}, AssignErrorResponse(err, req.Response.StatusCode)
	}
	defer res.Body.Close()

	// response part ------------

	if res.StatusCode != http.StatusOK {
		logrus.Errorf("Response status is not 200, it is %v", res.StatusCode)
		return DataTimeResponse{}, AssignErrorResponse(fmt.Errorf("unsupported header format"), res.StatusCode)
	}

	header := res.Header.Get("Content-Type")

	if !strings.Contains(header, "text/plain") && !strings.Contains(header, "application/json") {
		logrus.Errorf("unsupported header type, status code is  %d", res.StatusCode)
		return DataTimeResponse{}, AssignErrorResponse(err, req.Response.StatusCode)
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		logrus.Errorf("can't read response body. Err = %s", err)
		return DataTimeResponse{}, AssignErrorResponse(err, req.Response.StatusCode)
	}

	var datetime DataTimeResponse

	if strings.Contains(header, "application/json") {

		err = json.Unmarshal(body, &datetime)

		if err != nil {
			logrus.Errorf("unable to unmarchal body to json. Err = %s", err)
			return DataTimeResponse{}, AssignErrorResponse(err, req.Response.StatusCode)
		}

		return datetime, nil
	}

	datetime.DatewTime = string(body)

	return datetime, nil
}
