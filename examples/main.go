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

	clientId, ok := os.LookupEnv("WHATSAPP_CLIENT_ID")
	if !ok || clientId == "" {
		fmt.Println("WHATSAPP_CLIENT_ID not define in env")
		return
	}

	token, ok := os.LookupEnv("WHATSAPP_TOKEN")
	if !ok || token == "" {
		fmt.Println("WHATSAPP_TOKEN not define in env")
		return
	}

	msisdn, ok := os.LookupEnv("TEST_MSISDN")
	if !ok || token == "" {
		fmt.Println("TEST_MSISDN not define in env")
		return
	}

	fmt.Println("Send message")
	whatsappClient, err := whatsappapi.NewWhatsappCloudApiClient(&whatsappapi.WhatsappConfig{
		ClientId: clientId,
		Token:    token,
	})

	if err != nil {
		fmt.Println("Error:", err)
	}

	whatsappClient.SendTextMessage("Hello world!", msisdn)
}
