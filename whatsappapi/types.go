package whatsappapi

type MessageLanguageCode struct {
	Code string `json:"code"`
}

type MessageTemplate struct {
	Name     string              `json:"name"`
	Language MessageLanguageCode `json:"language"`
}

type MessageInteractiveHeader struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

type MessageInteractiveBody struct {
	Text string `json:"text"`
}

type MessageInteractiveFooter struct {
	Text string `json:"text"`
}

type MessageInteractiveAction struct {
	Name       string `json:"name"`
	Parameters struct {
		DisplayText string `json:"display_text"`
		URL         string `json:"url"`
	} `json:"parameters"`
}

type MessageInteractive struct {
	Type   string                    `json:"type"`
	Header *MessageInteractiveHeader `json:"header"`
	Body   *MessageInteractiveBody   `json:"body"`
	Footer *MessageInteractiveFooter `json:"footer"`
	Action *MessageInteractiveAction `json:"action"`
}

type MessageText struct {
	PreviewUrl *string `json:"preview_url"`
	Body       string  `json:"body"`
}

type MessagePayload struct {
	MessagingProduct string              `json:"messaging_product"`
	RecipientType    string              `json:"recipient_type"`
	To               string              `json:"to"`
	Type             string              `json:"type"`
	Template         *MessageTemplate    `json:"template"`
	Interactive      *MessageInteractive `json:"interactive"`
	Text             *MessageText        `json:"text"`
}
