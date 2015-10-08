// Copyright 2015 Raka Teja. All rights reserved.
// Write with Love

package controllers

import (
	"github.com/maderaka/goapp/packages/core"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

type Controller struct {
	statusCode int
}

func (ctrl *Controller) errors(err error) []core.ResponseError {

	return []core.ResponseError{core.ResponseError{"fatal-error", err.Error()}}
}

func (ctrl *Controller) decode(r *http.Request, v core.EntityValid) []core.ResponseError {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

	if err != nil {
		return ctrl.errors(err)
	}
	if err := r.Body.Close(); err != nil {
		return ctrl.errors(err)
	}

	// Decode json body into struct entity
	if err := json.Unmarshal(body, &v); err != nil {
		return ctrl.errors(err)
	} else {
		if errors := v.Valid(); len(errors) > 0 {
			return errors
		}
	}

	return nil
}

func (ctrl *Controller) response(w http.ResponseWriter, v interface{}, errs []core.ResponseError) {
	w.Header().Set("Content-Type", "application/json")
	if ctrl.statusCode == 0 {
		ctrl.statusCode = http.StatusOK
	}

	w.WriteHeader(ctrl.statusCode)
	res := core.Response{
		StatusCode: ctrl.statusCode,
		Message:    http.StatusText(ctrl.statusCode),
		Errors:     errs,
		Data:       v,
	}
	if err := json.NewEncoder(w).Encode(res); err != nil {
		panic(err)
	}
}
