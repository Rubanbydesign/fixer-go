// +build !appengine

package fixer

import (
	"net/http"

	"golang.org/x/net/context"
)

func getClient(ctx context.Context) *http.Client {
	return &http.Client{}
}
