package api

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/mgyosbel/pokedex/internal/cache"
)

const API_URL = "https://pokeapi.co/api/v2"

type Client struct {
	Client http.Client
	cc     *cache.Cache
}

func NewClient() *Client {
	return &Client{
		Client: http.Client{
			Timeout: time.Minute,
		},
		cc: cache.NewCache(60 * time.Second),
	}
}

func (c *Client) GetCached(endpoint string) ([]byte, error) {
	if cached, ok := c.cc.Get(endpoint); ok {
		return cached, nil
	}

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if res.StatusCode > 299 {
		return nil, fmt.Errorf("response failed with status code %d: %s", res.StatusCode, body)
	}
	if err != nil {
		return nil, err
	}

	c.cc.Add(endpoint, body)
	return body, nil
}
