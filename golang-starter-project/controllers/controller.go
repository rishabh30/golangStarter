package controllers

import (
	"golang-starter-project/dtos"
	"golang-starter-project/infrastructures"
	"golang-starter-project/services"
	"golang-starter-project/utils"
	"net/http"
)

type Controller struct {
	service services.Service
}

func (c *Controller) SetDependencies(service services.Service) {
	c.service = service
}

func (c *Controller) Hello(res http.ResponseWriter, req *http.Request) {
	infrastructures.Logger.Info().
		Str("correlationId", req.Context().Value("correlationId").(string)).
		Msg("Received new request to hello world")
	response := dtos.ResponseWithStatus{
		Status:  http.StatusOK,
		Message: "Hello World",
	}
	utils.Respond(req.Context(), res, response)
}
