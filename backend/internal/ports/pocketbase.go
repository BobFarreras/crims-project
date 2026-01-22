package ports

import "context"

// PocketBaseClient defineix el contracte minim per validar connexio amb la BaaS.
// Ha de ser substituible si canviem a Supabase.
type PocketBaseClient interface {
	Ping(ctx context.Context) error
	// NOU MÈTODE
	CreateUser(username, email, password, passwordConfirm, name string) error
}
