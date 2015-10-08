package api

import (
	c "github.com/maderaka/goapp/packages/api/controllers"
	"golang.org/x/net/context"
	h "net/http"
)

type MuxFunc map[string]func(ctx context.Context, w h.ResponseWriter, r *h.Request)

func PingRoute(mux MuxFunc) MuxFunc {
	// Ping endpoint
	ping := c.NewPingController()
	mux["/ping"] = ping.Index

	return mux
}

func AuthRoute(mux MuxFunc) MuxFunc {
	// User endpoints
	auth := c.NewAuthController()
	mux["/auth/register"] = auth.Register
	mux["/auth/login"] = auth.Login

	return mux
}
