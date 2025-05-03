package api

import (
	"net/http"
	"time"
)

const API_URL = "https://pokeapi.co/api/v2"

type Client struct {
	Client http.Client
}

func NewClient() *Client {
	return &Client{
		http.Client{
			Timeout: time.Minute,
		},
	}
}
