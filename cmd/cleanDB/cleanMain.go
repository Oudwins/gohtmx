package main

import (
	"fmt"
	"log"

	"github.com/Oudwins/stackifyer/db"
	"github.com/joho/godotenv"
)

var tables = []string{
	"files",
}

func main() {

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Missing .env file")
	}

	if err := db.Init(); err != nil {
		log.Fatal(("Init DB"))
	}
	for _, key := range tables {
		var res map[string]interface{}
		err := db.DbClient.DB.From(key).Delete().Neq("id", "572af60c-9839-47df-9976-361ad08df070").Execute(&res)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(res)
		}
	}
}
