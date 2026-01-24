package ports

import "context"

type contextKey string

const authTokenKey contextKey = "auth_token"

func WithAuthToken(ctx context.Context, token string) context.Context {
	return context.WithValue(ctx, authTokenKey, token)
}

func AuthTokenFromContext(ctx context.Context) string {
	value, _ := ctx.Value(authTokenKey).(string)
	return value
}
