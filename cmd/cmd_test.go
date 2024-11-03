package cmd

import (
	"net/http"
	"testing"
)

type MockHTTPClient struct {
	DoFunc func(*http.Request) (*http.Response, error)
}

func TestMap(t *testing.T) {

}
