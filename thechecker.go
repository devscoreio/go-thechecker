package thechecker

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"
)

var (
	protocol = "https://"
	domain   = "api.thechecker.co"
	endpoint = "verify"
)

type Client struct {
	apikey string
	http   *http.Client

	Response *Response
}

// NewClient returns a client for thechecker-go
func NewClient(apikey string) (client *Client, err error) {
	if apikey == "" {
		apikey = os.Getenv("THECHECKER_API")
	}
	client = &Client{apikey: apikey}
	client.configure()
	return client, nil
}

func (c *Client) configure() {
	c.Response = &AccountAPI{c}
}
func (c *Client) do(r *http.Request) (*http.Response, error) {
	resp, err := httpClient.Do(r)
	if err != nil {
		return nil, fmt.Errorf("%s. HTTP Response: %s", errUnknown, err.Error())
	}
	c.limit.Updated = time.Now().UTC()
	if resp.StatusCode == 400 {
		return nil, ErrStatus400
	}
	if resp.StatusCode == 403 {
		return nil, ErrStatus403
	}
	if resp.StatusCode == 404 {
		return nil, ErrStatus404
	}
	if resp.StatusCode == 405 {
		return nil, ErrStatus405
	}
	if resp.StatusCode == 422 {
		return nil, ErrStatus422
	}
	if resp.StatusCode == 500 {
		return nil, ErrStatus500
	}
	return resp, nil
}

func (c *Client) get(email) (*http.Request, error) {
	u := &url.URL{
		Scheme: protocol,
		Host:   domain,
		Path:   fmt.Sprintf("v1/%s", endPoint),
	}
	q := u.Query()
	q.Set("email", email)
	q.Set("api_key", c.apikey)
	u.RawQuery = q.Encode()
	r, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return r, nil
}
