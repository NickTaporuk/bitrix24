package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type ExtendedBitrix24 struct {
	webhookURL string
}

// NewExtendedBitrix24 creates a new instance of ExtendedBitrix24
func NewExtendedBitrix24(webhookURL string) *ExtendedBitrix24 {
	return &ExtendedBitrix24{webhookURL: webhookURL}
}

// CallMethod calls a Bitrix24 API method
func (bx24 *ExtendedBitrix24) CallMethod(method string, params map[string]string) ([]byte, error) {
	// Build the API endpoint URL
	endpoint := fmt.Sprintf("%s/%s", bx24.webhookURL, method)
	reqURL, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	// Add parameters to the URL
	query := reqURL.Query()
	for k, v := range params {
		query.Set(k, v)
	}
	reqURL.RawQuery = query.Encode()

	// Make the GET request
	resp, err := http.Get(reqURL.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read and decode the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
