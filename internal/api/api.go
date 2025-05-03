package api

import (
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
