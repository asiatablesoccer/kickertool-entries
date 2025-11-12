// Package kickertool api wrapper
package kickertool

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// API .
type API struct {
	accessToken string
	client      *http.Client
}

// New API client
func New(accessToken string) *API {
	return &API{
		accessToken: accessToken,
		client:      http.DefaultClient,
	}
}

// DoAPIRequest builds http request
func (api API) DoAPIRequest(
	method string,
	urlPath string,
	header http.Header,
	body io.Reader,
	output any,
) error {
	var (
		err      error
		respBody io.ReadCloser
		url      = urlPath
	)
	if header == nil {
		header = make(http.Header)
	}
	header.Add("Authorization", fmt.Sprintf("Bearer %s", api.accessToken))
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return err
	}
	req.Header = header
	resp, err := api.client.Do(req)
	if err != nil {
		return err
	}
	respBody = resp.Body
	defer respBody.Close()
	err = json.NewDecoder(respBody).Decode(&output)
	if err != nil {
		return err
	}
	return err
}
