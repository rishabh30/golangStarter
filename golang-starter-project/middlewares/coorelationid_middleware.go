package middlewares

import (
	"context"
	"net/http"

	"github.com/rs/xid"
)

func CoorelationIdMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		coorelationId := xid.New().String()
		ctx := context.WithValue(req.Context(), "correlationId", coorelationId)
		req = req.WithContext(ctx)
		next.ServeHTTP(res, req)
	})
}
