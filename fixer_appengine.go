// +build appengine

package fixer

import (
	"net/http"

	"golang.org/x/net/context"

	"google.golang.org/appengine/urlfetch"
)

func getClient(ctx context.Context) *http.Client {
	return urlfetch.Client(ctx)
}
