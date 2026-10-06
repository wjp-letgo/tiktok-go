package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

// @json
type WebhooksResult struct {
	Code      int          `json:"code"`
	Message   string       `json:"message"`
	Data      WebhooksData `json:"data"`
	RequestId string       `json:"request_id"`
}

func (e *WebhooksResult) String() string {
	return lib.ObjectToString(e)
}

// @json
type WebhooksData struct {
	Webhooks Webhooks `json:"webhooks"`
}

func (e *WebhooksData) String() string {
	return lib.ObjectToString(e)
}

type Webhooks []Webhook

func (e *Webhooks) String() string {
	return lib.ObjectToString(e)
}

// @json
type Webhook struct {
	EventType string `json:"event_type"`
	Address   string `json:"address"`
	CreateTime int   `json:"create_time"`
	UpdateTime int   `json:"update_time"`
}

func (e *Webhook) String() string {
	return lib.ObjectToString(e)
}

// @json
type UpdateWebhookRequest struct {
	Address   string `json:"address,omitempty"`
	EventType string `json:"event_type,omitempty"`
}

func (e *UpdateWebhookRequest) String() string {
	return lib.ObjectToString(e)
}

// @json
type UpdateWebhookResult struct {
	Code      int         `json:"code"`
	Message   string      `json:"message"`
	Data      EmptyData   `json:"data"`
	RequestId string      `json:"request_id"`
}

func (e *UpdateWebhookResult) String() string {
	return lib.ObjectToString(e)
}

// @json
type DeleteWebhookRequest struct {
	EventType string `json:"event_type,omitempty"`
}

func (e *DeleteWebhookRequest) String() string {
	return lib.ObjectToString(e)
}

// @json
type DeleteWebhookResult struct {
	Code      int       `json:"code"`
	Message   string    `json:"message"`
	Data      EmptyData `json:"data"`
	RequestId string    `json:"request_id"`
}

func (e *DeleteWebhookResult) String() string {
	return lib.ObjectToString(e)
}

// @json
type EmptyData struct {
}

func (e *EmptyData) String() string {
	return lib.ObjectToString(e)
}
