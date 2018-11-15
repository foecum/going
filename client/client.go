package client //import "github.com/foecum/going/client"

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"time"
)

// Requester ...
type Requester struct {
	c        http.Client
	endpoint string
}

// Config contains configuration details for creating a client
type Config struct {
	Endpoint      string
	Timeout       time.Duration
	CheckRedirect func(req *http.Request, via []*http.Request) error
	Jar           http.CookieJar
}

// NewHTTPClient creates a new http client
func NewHTTPClient(cfg Config) (Requester, error) {
	c := http.Client{
		Transport: &http.Transport{
			Dial: func(network, address string) (net.Conn, error) {
				return net.DialTimeout(network, address, cfg.Timeout)
			},
		},
		CheckRedirect: cfg.CheckRedirect,
		Jar:           cfg.Jar,
	}

	return Requester{c: c, endpoint: cfg.Endpoint}, nil
}

// MakeRequest ...
func (r Requester) MakeRequest(method, path string, body io.Reader) (*json.Decoder, error) {
	url := fmt.Sprintf("%s%s", r.endpoint, path)

	req, err := http.NewRequest(method, url, body)

	if err != nil {
		return nil, err
	}

	resp, err := r.c.Do(req)
	if err != nil {
		return nil, err
	}
	return json.NewDecoder(resp.Body), nil
}
