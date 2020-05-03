package bml

import (
	"net/http"
	"net/http/cookiejar"
)

// NewClient ...
func NewClient() *http.Client {
	jar, _ := cookiejar.New(nil)

	return &http.Client{
		Jar: jar,
	}
}
