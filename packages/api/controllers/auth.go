// Copyright 2016 Raka Teja. All rights reserved.
// Write with Love

package controllers

import (
	"github.com/maderaka/goapp/packages/account"
	"golang.org/x/net/context"
	"net/http"
)

type AuthController struct {
	*Controller
}

func NewAuthController() *AuthController {
	return &AuthController{&Controller{}}
}

func (auth *AuthController) Login(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	auth.response(w, "Login", nil)
}

func (auth *AuthController) Register(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var user account.User
		// Decode request body into entity
		// Do input validation before return the struct
		if errors := auth.decode(r, &user); len(errors) > 0 {
			auth.statusCode = http.StatusInternalServerError
			auth.response(w, nil, errors)
		} else {

			// Instance user service
			userService := account.NewUserService()

			// Create user
			// Check email address is already registered
			// Register user if he/she is not registered yet
			user, err := userService.Create(ctx, user)
			if err != nil {
				auth.response(w, nil, auth.errors(err))
			} else {
				auth.response(w, user, nil)
			}
		}
	}
}
