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
// The payload is a string representing the graph in deflate + base64 format
func (c *Client) GetRequestContext(ctx context.Context, payload string, graphFormat GraphFormat, imageFormat ImageFormat) (string, error) {

	// construct the url
	u, err := url.Parse(fmt.Sprintf("%s://%s", c.Config.Mode, c.Config.Host))
	if err != nil {
		return "", errors.Wrapf(err, "fail to create URL from %s", c.Config.Host)
	}
	u.Path = path.Join(u.Path, string(graphFormat), string(imageFormat), payload)

	// construct the request
	req, err := http.NewRequest("GET", u.String(), nil)
	timeoutCtx, cancel := context.WithTimeout(ctx, c.Config.Timeout)
	defer cancel()
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
		return "", errors.Errorf(
			"fail to generate the image, status %d",
			response.StatusCode)
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", errors.Wrap(err, "fail to read the response body")
	}
	return string(body), nil
}

// GetRequest executes a GET request on Kroki
// The payload is a string representing the graph in deflate + base64 format
func (c *Client) GetRequest(payload string, graphFormat GraphFormat, imageFormat ImageFormat) (string, error) {
	return c.GetRequestContext(context.Background(), payload, graphFormat, imageFormat)
}
