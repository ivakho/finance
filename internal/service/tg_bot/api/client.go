package api

import "net/http"

type Client struct {
	BaseURL string
	Client  *http.Client
}
