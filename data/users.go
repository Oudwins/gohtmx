package data

import "github.com/nedpals/supabase-go"

type AuthenticatedDetails struct {
	User *supabase.User
}
