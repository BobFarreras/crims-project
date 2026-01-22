package ports

import "context"

// AuthResponse representa la resposta de PocketBase al fer login
type AuthResponse struct {
	Token  string `json:"token"`
	Record struct {
		ID       string `json:"id"`
		Username string `json:"username"`
		Email    string `json:"email"`
		Name     string `json:"name"`
	} `json:"record"`
}

type PocketBaseClient interface {
	Ping(ctx context.Context) error
	CreateUser(username, email, password, passwordConfirm, name string) error
	// 🔥 NOU MÈTODE
	AuthWithPassword(identity, password string) (*AuthResponse, error)
	RefreshAuth(token string) (*AuthResponse, error)
}
