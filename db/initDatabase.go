package db

import (
	"os"

	supa "github.com/nedpals/supabase-go"
)

var DbClient *supa.Client

// I need to somehow based on if dev or prod instantiate different db

func createDB() error {

	supaClient := supa.CreateClient(os.Getenv("SUPABASE_URL"), os.Getenv("SUPABASE_KEY"))

	DbClient = supaClient

	return nil
}

func Init() error {

	createDB()
	return nil
}
