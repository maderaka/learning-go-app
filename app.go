package main

import (
	"github.com/maderaka/goapp/packages/api"
	"github.com/maderaka/goapp/app"
	"net/http"
)

func main() {
	ctx := app.ContextValues()
	server := http.Server{
		Addr:    ":9090",
		Handler: api.NewContextAdapter(ctx),
	}

	server.ListenAndServe()
}
