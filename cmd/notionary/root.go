package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	notionary "github.com/mathisbot/notionary-go"
	"github.com/spf13/cobra"
)

var client *notionary.Client

var rootCmd = &cobra.Command{
	Use:   "notionary",
	Short: "Notionary CLI",
}

func init() {
	godotenv.Load()
	token := os.Getenv("NOTION_API_KEY")
	if token == "" {
		log.Fatal("NOTION_API_KEY not set")
	}
	client = notionary.New(token)
}

func main() {
	rootCmd.AddCommand(pagesCmd, usersCmd, markdownCmd)
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}