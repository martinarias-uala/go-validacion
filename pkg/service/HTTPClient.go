package service

import "net/http"

type HTTPClient interface {
	Get(url string) (*http.Response, error)
}
