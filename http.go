package kroki

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"

	"github.com/pkg/errors"
)

// GetRequestContext executes a GET request on Kroki using a context
// The payload is a string representing the diagram in deflate + base64 format
func (c *Client) GetRequestContext(ctx context.Context, payload string, diagramType DiagramType, imageFormat ImageFormat) (string, error) {

	// construct the url
	u, err := url.Parse(c.Config.URL)
	if err != nil {
		return "", errors.Wrapf(err, "fail to create the URL from %s", c.Config.URL)
	}
	u.Path = path.Join(u.Path, string(diagramType), string(imageFormat), payload)

	// construct the request
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return "", errors.Wrap(err, "fail to create the request")
	}
	timeoutCtx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()
	req.Header = http.Header{
		"Accept":     {"text/plain"},
		"User-Agent": {fmt.Sprintf("kroki-go %s", Version)},
	}
	req = req.WithContext(timeoutCtx)

	// execute the request
	client := http.DefaultClient
	response, err := client.Do(req)
	if err != nil {
		return "", errors.Wrap(err, "fail to generate the image")
	}

	// read the result
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		body, err := ioutil.ReadAll(response.Body)
		var message string
		if err != nil {
			message = ""
		} else {
			message = string(body)
		}
		return "", errors.Errorf(
			"fail to generate the image {status: %d, body: %s}",
			response.StatusCode, message)
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", errors.Wrap(err, "fail to read the response body")
	}
	return string(body), nil
}

// GetRequest executes a GET request on Kroki
// The payload is a string representing the diagram in deflate + base64 format
func (c *Client) GetRequest(payload string, diagramType DiagramType, imageFormat ImageFormat) (string, error) {
	return c.GetRequestContext(context.Background(), payload, diagramType, imageFormat)
}
