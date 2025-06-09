package whatsappapi

// base message
type Message interface {
	GetType() string
}

// Message payload
type MessagePayload struct {
	MessagingProduct string              `json:"messaging_product"`
	RecipientType    string              `json:"recipient_type"`
	To               string              `json:"to"`
	Type             string              `json:"type"`
	Template         *MessageTemplate    `json:"template"`
	Interactive      *MessageInteractive `json:"interactive"`
	Text             *MessageText        `json:"text"`
}

// Message type: template
type MessageTemplate struct {
	Name       string                       `json:"name"`
	Language   *MessageTemplateLanguageCode `json:"language"`
	Components []*Message                   `json:"components"`
}

func (m *MessageTemplate) GetType() string {
	return "template"
}

type MessageTemplateComponent struct {
	Type       string                      `json:"type"`
	Parameters []*MessageTemplateParameter `json:"parameters"`
}

type MessageTemplateParameter interface {
	GetParameterType() string
}

type MessageTemplateParamterText struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

func (m *MessageTemplateParamterText) GetParameterType() string {
	return "text"
}

type MessageTemplateParamterCurrency struct {
	Type     string                                   `json:"type"`
	Currency *MessageTemplateParamterCurrencyCurrency `json:"currency"`
}

type MessageTemplateParamterCurrencyCurrency struct {
	FallbackValue string `json:"fallback_value"`
	Code          string `json:"code"`
	Amount1000    string `json:"amount_1000"`
}

func (m *MessageTemplateParamterCurrency) GetParameterType() string {
	return "currency"
}

type MessageTemplateParamterDateTime struct {
	Type     string                                   `json:"type"`
	Currency *MessageTemplateParamterDateTimeDateTime `json:"date_time"`
}

type MessageTemplateParamterDateTimeDateTime struct {
	FallbackValue string `json:"fallback_value"`
}

func (m *MessageTemplateParamterDateTime) GetParameterType() string {
	return "text"
}

type MessageTemplateLanguageCode struct {
	Code string `json:"code"`
}

// message type: interactive
type MessageInteractive struct {
	Type   string                    `json:"type"`
	Header *MessageInteractiveHeader `json:"header"`
	Body   *MessageInteractiveBody   `json:"body"`
	Footer *MessageInteractiveFooter `json:"footer"`
	Action *MessageInteractiveAction `json:"action"`
}

func (m *MessageInteractive) GetType() string {
	return "interactive"
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

// message type text
type MessageText struct {
	PreviewUrl *string `json:"preview_url"`
	Body       string  `json:"body"`
}

func (m *MessageText) GetType() string {
	return "text"
}
