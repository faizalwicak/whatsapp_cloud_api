package tests

import (
	"log"
	"os"
	"testing"

	"github.com/faizalwicak/whatsapp_cloud_api/whatsappapi"
	"github.com/joho/godotenv"
)

func setupClient(t *testing.T) *whatsappapi.WhatsappClient {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	client, err := whatsappapi.NewWhatsappCloudApiClient(&whatsappapi.WhatsappConfig{
		Token:    os.Getenv("WHATSAPP_TOKEN"),
		ClientId: os.Getenv("WHATSAPP_CLIENT_ID"),
	})

	if err != nil {
		t.Fatalf("Error create client: %v", err)
	}

	return client
}

func tearDown() {

}

func TestSendMessage(t *testing.T) {
	client := setupClient(t)
	client.SendTemplateMessage("hello_world", "en_US", "msisdn")
}

func TestSendTextMessage(t *testing.T) {

	client := setupClient(t)

	client.SendTextMessage("hello", "msisdn")
}
