package controllers

import (
	"golang.org/x/net/context"
	"net/http"
)

type PingController struct {
	*Controller
}

func NewPingController() *PingController {
	return &PingController{&Controller{}}
}

func (p *PingController) Index(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	p.response(w, "Pong!", nil)
}
