package whatsappapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type WhatsappClient struct {
	whatsappConfig WhatsappConfig
}

type WhatsappConfig struct {
	BaseUrl  string
	ClientId string
	Token    string
}

func NewWhatsappCloudApiClient(wc *WhatsappConfig) (*WhatsappClient, error) {
	if wc == nil {
		wc = &WhatsappConfig{}
	}

	if wc.BaseUrl == "" {
		wc.BaseUrl = "https://graph.facebook.com/v22.0"
	}

	if wc.ClientId == "" {
		return nil, fmt.Errorf("client is required, whatsappconfig: %v", wc)
	}

	if wc.Token == "" {
		return nil, fmt.Errorf("token is required, whatsappconfig: %v", wc)
	}

	return &WhatsappClient{
		whatsappConfig: *wc,
	}, nil
}

func (c *WhatsappClient) SendTemplateMessage(templateName, languageCode, msisdn string) {
	payload := MessagePayload{
		MessagingProduct: "whatsapp",
		RecipientType:    "individual",
		To:               msisdn,
		Type:             "template",
		Template: &MessageTemplate{
			Name: templateName,
			Language: MessageLanguageCode{
				Code: languageCode,
			},
		},
	}
	c.SendMessage(msisdn, payload)
}

func (c *WhatsappClient) SendTextMessage(text string, msisdn string) {
	payload := MessagePayload{
		MessagingProduct: "whatsapp",
		RecipientType:    "individual",
		To:               msisdn,
		Type:             "text",
		Text: &MessageText{
			Body: text,
		},
	}
	c.SendMessage(msisdn, payload)
}

func (c *WhatsappClient) SendMessage(msisdn string, payload MessagePayload) {

	jsonData, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	fmt.Println(string(jsonData))

	url := fmt.Sprintf("%s/%s/messages", c.whatsappConfig.BaseUrl, c.whatsappConfig.ClientId)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	req.Header.Set("Authorization", "Bearer "+c.whatsappConfig.Token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println(resp.StatusCode)

	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
}
