package riotapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httputil"
	"regexp"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

type EndpointMismatchError error
type AuthenticationError error

type Client struct {
	ApiKey string
	Debug  bool
}

var GlobalClient = Client{"", false}

var endpointRegex = regexp.MustCompile("^https:\\/\\/([a-zA-Z0-9]+).api.riotgames.com")

func (c *Client) GetAndUnmarshal(endpoint string, v interface{}) error {
	resp, err := c.AuthenticatedRequest(endpoint)
	if err != nil {
		return errors.Wrap(err, "GetAndUnmarshal")
	}

	defer resp.Body.Close()
	return errors.Wrap(json.NewDecoder(resp.Body).Decode(v), "GetAndUnmarshal")
}

func (c *Client) AuthenticatedRequest(endpoint string) (*http.Response, error) {
	if c.ApiKey == "" {
		return nil, AuthenticationError(fmt.Errorf("AuthenticatedRequest: You must set the GlobalClient's ApiKey before requesting any resources"))
	}

	region := endpointRegex.FindString(endpoint)
	if region == "" {
		return nil, EndpointMismatchError(fmt.Errorf("AuthenticatedRequest: The supplied endpoint: %s could not be matched", endpoint))
	}
	// todo: ratelimiting
	// https://developer.riotgames.com/rate-limiting.html
	log.Info(region)

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, errors.Wrap(err, "AuthenticatedRequest")
	}

	req.Header.Set("X-RIOT-TOKEN", c.ApiKey)

	client := &http.Client{}

	resp, err := client.Do(req)
	if c.Debug {
		dump, err := httputil.DumpResponse(resp, true)
		log.Info(string(dump), err)
	}

	return resp, errors.Wrap(err, "AuthenticatedRequest")
}
