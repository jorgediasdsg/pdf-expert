package api

import (
	"context"
)

type contextKey string

const reqIDKey contextKey = "request_id"

func withRequestID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, reqIDKey, id)
}

func getRequestID(ctx context.Context) string {
	val, ok := ctx.Value(reqIDKey).(string)
	if !ok {
		return ""
	}
	return val
}
