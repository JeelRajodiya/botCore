package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	Token        string
	DatabaseId   string
	NotionSecret string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	Token = os.Getenv("DISCORD_TOKEN")
	NotionSecret = os.Getenv("NOTION_SECRET")
	DatabaseId = os.Getenv("NOTION_DATABASE_ID")

	if Token == "" {
		log.Fatal("Discord bot token not found in .env file")
	}

	if NotionSecret == "" {
		log.Fatal("NotionSecret not found in .env file")
	}

	if DatabaseId == "" {
		log.Fatal("DatabaseId not found in .env file")
	}
}
