package vhttp

import (
	"net/http"
	"net/url"
	"strings"
)

var defaultClient = CreateClient()

// Get Send GET request
func Get(url string) (response *http.Response, err error) {
	resp, err := defaultClient.Get(url)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Post Send POST request
func Post(url string, contentType string, data url.Values) (response *http.Response, err error) {
	req, err := http.NewRequest("POST", url, strings.NewReader(data.Encode()))
	req.Header.Add("content-type", contentType)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()

	response, err = defaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func IsOK(resp *http.Response) bool {
	switch resp.StatusCode {
	case http.StatusOK:
	case http.StatusCreated:
	case http.StatusAccepted:
	case http.StatusAlreadyReported:
	case http.StatusNoContent:
		return true
	default:
		return false
	}
	return false
}
