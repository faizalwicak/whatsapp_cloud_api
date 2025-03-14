package main

import (
	"fmt"
	"log"
	"os"

	"github.com/faizalwicak/whatsapp_cloud_api/whatsappapi"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	fmt.Println("Send message")
	whatsappClient, err := whatsappapi.NewWhatsappCloudApiClient(&whatsappapi.WhatsappConfig{
		ClientId: os.Getenv("WHATSAPP_CLIENT_ID"),
		Token:    os.Getenv("WHATSAPP_TOKEN"),
	})

	if err != nil {
		fmt.Println("Error:", err)
	}

	whatsappClient.SendTextMessage("Hello world!", "")
}
