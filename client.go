package kel

import (
	"errors"
	"net/http"
	"net/url"
)

// Client represents the HTTP client used to perform requests against the Kel API.
type Client struct {
	*http.Client

	baseURL *url.URL

	ResourceGroups *ResourceGroupService
	Sites          *SiteService
}

// New returns a Kel HTTP client to perform requests.
func New(httpClient *http.Client, baseURL string) (*Client, error) {
	if httpClient == nil {
		return nil, errors.New("HTTP client is nil")
	}
	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}
	client := &Client{
		Client:  httpClient,
		baseURL: u,
	}
	client.ResourceGroups = &ResourceGroupService{client: client}
	client.Sites = &SiteService{client: client}
	return client, nil
}

func (client *Client) makeURL(path string) *url.URL {
	var u url.URL
	u = *client.baseURL
	u.Path += path
	return &u
}
