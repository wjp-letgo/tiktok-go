package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

// @json
type EmptyData struct {
}

func (e *EmptyData) String() string {
	return lib.ObjectToString(e)
}

// @json
type CreateConversationRequest struct {
	BuyerUserId string `json:"buyer_user_id,omitempty"`
}

func (e *CreateConversationRequest) String() string {
	return lib.ObjectToString(e)
}

// @json
type CreateConversationResult struct {
	Code      int                    `json:"code"`
	Message   string                 `json:"message"`
	Data      CreateConversationData `json:"data"`
	RequestId string                 `json:"request_id"`
}

func (e *CreateConversationResult) String() string {
	return lib.ObjectToString(e)
}

// @json
type CreateConversationData struct {
	ConversationId string `json:"conversation_id"`
}

func (e *CreateConversationData) String() string {
	return lib.ObjectToString(e)
}

// @json
type ConversationsResult struct {
	Code      int               `json:"code"`
	Message   string            `json:"message"`
	Data      ConversationsData `json:"data"`
	RequestId string            `json:"request_id"`
}

func (e *ConversationsResult) String() string {
	return lib.ObjectToString(e)
}

// @json
type ConversationsData struct {
	NextPageToken string        `json:"next_page_token"`
	Conversations Conversations `json:"conversations"`
}

func (e *ConversationsData) String() string {
	return lib.ObjectToString(e)
}

type Conversations []Conversation

func (e *Conversations) String() string {
	return lib.ObjectToString(e)
}

// @json
type Conversation struct {
	Id             string `json:"id"`
	BuyerUserId    string `json:"buyer_user_id"`
	CanSendMessage bool   `json:"can_send_message"`
	CreateTime     int    `json:"create_time"`
	UpdateTime     int    `json:"update_time"`
	UnreadCount    int    `json:"unread_count"`
}

func (e *Conversation) String() string {
	return lib.ObjectToString(e)
}

// @json
type MessagesResult struct {
	Code      int          `json:"code"`
	Message   string       `json:"message"`
	Data      MessagesData `json:"data"`
	RequestId string       `json:"request_id"`
}

func (e *MessagesResult) String() string {
	return lib.ObjectToString(e)
}

// @json
type MessagesData struct {
	NextPageToken string   `json:"next_page_token"`
	Messages      Messages `json:"messages"`
}

func (e *MessagesData) String() string {
	return lib.ObjectToString(e)
}

type Messages []Message

func (e *Messages) String() string {
	return lib.ObjectToString(e)
}

// @json
type Message struct {
	Id         string `json:"id"`
	Type       string `json:"type"`
	Content    string `json:"content"`
	CreateTime int    `json:"create_time"`
	IsVisible  bool   `json:"is_visible"`
	Sender     Sender `json:"sender"`
}

func (e *Message) String() string {
	return lib.ObjectToString(e)
}

// @json
type Sender struct {
	Id       string `json:"id"`
	Role     string `json:"role"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
}

func (e *Sender) String() string {
	return lib.ObjectToString(e)
}

// @json
type SendMessageRequest struct {
	Type    string `json:"type,omitempty"`
	Content string `json:"content,omitempty"`
}

func (e *SendMessageRequest) String() string {
	return lib.ObjectToString(e)
}

// @json
type SendMessageResult struct {
	Code      int             `json:"code"`
	Message   string          `json:"message"`
	Data      SendMessageData `json:"data"`
	RequestId string          `json:"request_id"`
}

func (e *SendMessageResult) String() string {
	return lib.ObjectToString(e)
}

// @json
type SendMessageData struct {
	MessageId string `json:"message_id"`
}

func (e *SendMessageData) String() string {
	return lib.ObjectToString(e)
}

// @json
type ReadMessageResult struct {
	Code      int       `json:"code"`
	Message   string    `json:"message"`
	Data      EmptyData `json:"data"`
	RequestId string    `json:"request_id"`
}

func (e *ReadMessageResult) String() string {
	return lib.ObjectToString(e)
}

// @json
type AgentSettingsResult struct {
	Code      int               `json:"code"`
	Message   string            `json:"message"`
	Data      AgentSettingsData `json:"data"`
	RequestId string            `json:"request_id"`
}

func (e *AgentSettingsResult) String() string {
	return lib.ObjectToString(e)
}

// @json
type AgentSettingsData struct {
	IsAutoReply bool   `json:"is_auto_reply"`
	Greeting    string `json:"greeting"`
}

func (e *AgentSettingsData) String() string {
	return lib.ObjectToString(e)
}

// @json
type UpdateAgentSettingsRequest struct {
	IsAutoReply bool   `json:"is_auto_reply,omitempty"`
	Greeting    string `json:"greeting,omitempty"`
}

func (e *UpdateAgentSettingsRequest) String() string {
	return lib.ObjectToString(e)
}

// @json
type UpdateAgentSettingsResult struct {
	Code      int       `json:"code"`
	Message   string    `json:"message"`
	Data      EmptyData `json:"data"`
	RequestId string    `json:"request_id"`
}

func (e *UpdateAgentSettingsResult) String() string {
	return lib.ObjectToString(e)
}

// @json
type PerformanceResult struct {
	Code      int             `json:"code"`
	Message   string          `json:"message"`
	Data      PerformanceData `json:"data"`
	RequestId string          `json:"request_id"`
}

func (e *PerformanceResult) String() string {
	return lib.ObjectToString(e)
}

// @json
type PerformanceData struct {
	ResponseRate          string `json:"response_rate"`
	ResponseTime          string `json:"response_time"`
	SatisfactionRate      string `json:"satisfaction_rate"`
	SupportSessionCount   int    `json:"support_session_count"`
}

func (e *PerformanceData) String() string {
	return lib.ObjectToString(e)
}

// @json
type UploadImageResult struct {
	Code      int             `json:"code"`
	Message   string          `json:"message"`
	Data      UploadImageData `json:"data"`
	RequestId string          `json:"request_id"`
}

func (e *UploadImageResult) String() string {
	return lib.ObjectToString(e)
}

// @json
type UploadImageData struct {
	Url string `json:"url"`
	Uri string `json:"uri"`
}

func (e *UploadImageData) String() string {
	return lib.ObjectToString(e)
}
