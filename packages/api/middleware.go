// Copyright 2016 Raka Teja. All rights reserved.
// Write with Love

package api

import (
	"fmt"
	"golang.org/x/net/context"
	"net/http"
)

// ------------------------------------------------------------------
// HTTP Authentication
// ------------------------------------------------------------------
type Authenticate struct{}

// Doing authenticate by reading values from headers
func (auth *Authenticate) do(w http.ResponseWriter, r *http.Request) bool {
	id := r.Header.Get(ApiIdHeaderName)
	user := r.Header.Get(ApiUserHeaderName)
	secret := r.Header.Get(ApiSecretHeaderName)

	passed := false
	if client, ok := clients[id]; ok {
		if client[ApiUserHeaderName] == user && client[ApiSecretHeaderName] == secret {
			passed = true
		}
	}

	if !passed {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Forbidden"))
	}
	return passed
}

// ------------------------------------------------------------------
// HTTP Request Context
// ------------------------------------------------------------------

// An interface for context handler
type ContextHandler interface {
	ServeHTTPContext(context.Context, http.ResponseWriter, *http.Request)
}

// An implementation of context handler interface
type ContextHandlerFunc struct {
}

// A method for serving http context
func (ch ContextHandlerFunc) ServeHTTPContext(ctx context.Context, rw http.ResponseWriter, req *http.Request) {
	// Authentication
	apiAuth := new(Authenticate)

	// Initialize mutex func handler
	mux := make(MuxFunc)

	// If passed the authentication above
	if ok := apiAuth.do(rw, req); ok {

		// Register routes
		mux = PingRoute(mux)
		mux = AuthRoute(mux)

		// Check registered route
		if h, ok := mux[req.URL.String()]; ok {
			h(ctx, rw, req)
		} else {
			fmt.Println("Not found")
		}
	}
}

// A ContextAdapter struct
type ContextAdapter struct {
	ctx     context.Context
	handler ContextHandler
}

// A function to instance ContextAdapter abstract data type
func NewContextAdapter(ctx context.Context) *ContextAdapter {
	return &ContextAdapter{ctx, ContextHandlerFunc{}}
}

// A context adapter method to serve every http request
func (ca *ContextAdapter) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	ca.handler.ServeHTTPContext(ca.ctx, rw, req)
}
