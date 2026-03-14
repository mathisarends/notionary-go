package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	notionary "github.com/mathisbot/notionary-go"
)

func main() {
	godotenv.Load()

	token := os.Getenv("NOTION_API_KEY")
	if token == "" {
		log.Fatal("NOTION_API_KEY not set")
	}

	client := notionary.New(token)

	page, err := client.Pages.FindByTitle(context.Background(), "Jarvis")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(page)
}