package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"golang-starter-project/dtos"
	"net/http"
)

func Respond(ctx context.Context, res http.ResponseWriter, response dtos.ResponseWithStatus) error {
	res.Header().Set("Content-Type", "application/json")
	res.Header().Set("Coorelation-Id", ctx.Value("correlationId").(string))
	res.WriteHeader(response.Status)
	jsonResponse, err := json.Marshal(response.Message)
	if err != nil {
		return fmt.Errorf("error while marshalling response: %v", err)
	}
	res.Write(jsonResponse)
	return nil
}
