package core

import (
	"context"
	"net/http"
)

// Context creates a new context from request.
func Context(r *http.Request) (ctx context.Context) {
	ctx = r.Context()
	return
}
