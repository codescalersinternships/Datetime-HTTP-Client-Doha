package httpClient

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/sirupsen/logrus"
)

// Time takes for backoff to wait
const (
	ConstBackoffTime = 4 * time.Second
)

// DataTimeResponse represent date and time response
type DataTimeResponse struct {
	DateTime string `json:"date_time"`
}

// ErrResponse represent error response
type ErrResponse struct {
	Err        error `json:"error"`
	StatusCode int   `json:"statuscode"`
}

// Error implement custom error types
func (e ErrResponse) Error() string {
	return fmt.Sprintf("error is %v, and http status code is %d", e.Err, e.StatusCode)
}

// AssignErrorResponse takes error and status code and return an ErrResponse type
func assignErrorResponse(err error, statuscode int) error {
	return ErrResponse{
		Err:        err,
		StatusCode: statuscode,
	}
}

// GetResponse retrive DataTimeResponse and error if exist
func (c *Client) GetDateTime() (DataTimeResponse, error) {

	req, err := http.NewRequest("GET", c.Url, nil)

	if err != nil {
		logrus.Errorf("from http.NewRequestWithContext function %s\n", err)
		return DataTimeResponse{}, assignErrorResponse(err, req.Response.StatusCode)
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
		return DataTimeResponse{}, assignErrorResponse(err, req.Response.StatusCode)
	}
	defer res.Body.Close()

	// response part ------------

	return parseResponse(res)
}

func parseResponse(res *http.Response) (DataTimeResponse, error) {
	if res.StatusCode != http.StatusOK {
		logrus.Errorf("Response status is not 200, it is %v", res.StatusCode)
		return DataTimeResponse{}, assignErrorResponse(errors.New("response status is not ok"), res.StatusCode)
	}

	header := res.Header.Get("Content-Type")

	if !strings.Contains(header, "text/plain") && !strings.Contains(header, "application/json") {
		logrus.Errorf("unsupported header type, status code is  %d", res.StatusCode)
		return DataTimeResponse{}, assignErrorResponse(errors.New("unsupported header type"), res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		logrus.Errorf("can't read response body. Err = %s", err)
		return DataTimeResponse{}, assignErrorResponse(err, res.StatusCode)
	}

	var datetime DataTimeResponse

	if strings.Contains(header, "application/json") {

		err = json.Unmarshal(body, &datetime)

		if err != nil {
			logrus.Errorf("unable to unmarchal body to json. Err = %s", err)
			return DataTimeResponse{}, assignErrorResponse(err, res.StatusCode)
		}

		return datetime, nil
	}

	datetime.DateTime = string(body)

	return datetime, nil
}
