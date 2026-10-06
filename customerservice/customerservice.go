package customerservice

import (
	"fmt"

	"github.com/wjp-letgo/letgo/lib"
	tiktokConfig "github.com/wjp-letgo/tiktok-go/config"
	customerserviceentity "github.com/wjp-letgo/tiktok-go/customerservice/entity"
)

// CustomerService
type CustomerService struct {
	Config *tiktokConfig.Config
}

func (a *CustomerService) CreateConversation(body *customerserviceentity.CreateConversationRequest) *customerserviceentity.CreateConversationResult {
	var result customerserviceentity.CreateConversationResult
	err := a.Config.HttpS3("/customer_service/202309/conversations", nil, body, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}

func (a *CustomerService) GetConversations(pageSize int, pageToken, locale string) *customerserviceentity.ConversationsResult {
	var result customerserviceentity.ConversationsResult
	params := lib.InRow{
		"page_size": pageSize,
	}
	if pageToken != "" {
		params["page_token"] = pageToken
	}
	if locale != "" {
		params["locale"] = locale
	}
	err := a.Config.GetS3("/customer_service/202309/conversations", params, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}

func (a *CustomerService) GetConversationMessages(conversationId string, pageSize int, pageToken string) *customerserviceentity.MessagesResult {
	var result customerserviceentity.MessagesResult
	params := lib.InRow{
		"page_size": pageSize,
	}
	if pageToken != "" {
		params["page_token"] = pageToken
	}
	err := a.Config.GetS3(fmt.Sprintf("/customer_service/202309/conversations/%s/messages", conversationId), params, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}

func (a *CustomerService) SendMessage(conversationId string, body *customerserviceentity.SendMessageRequest) *customerserviceentity.SendMessageResult {
	var result customerserviceentity.SendMessageResult
	err := a.Config.HttpS3(fmt.Sprintf("/customer_service/202309/conversations/%s/messages", conversationId), nil, body, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}

func (a *CustomerService) ReadMessage(conversationId string) *customerserviceentity.ReadMessageResult {
	var result customerserviceentity.ReadMessageResult
	err := a.Config.HttpS3(fmt.Sprintf("/customer_service/202309/conversations/%s/messages/read", conversationId), nil, nil, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}

func (a *CustomerService) GetAgentSettings() *customerserviceentity.AgentSettingsResult {
	var result customerserviceentity.AgentSettingsResult
	params := lib.InRow{}
	err := a.Config.GetS3("/customer_service/202309/agents/settings", params, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}

func (a *CustomerService) UpdateAgentSettings(body *customerserviceentity.UpdateAgentSettingsRequest) *customerserviceentity.UpdateAgentSettingsResult {
	var result customerserviceentity.UpdateAgentSettingsResult
	err := a.Config.PutS3("/customer_service/202309/agents/settings", nil, body, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}

func (a *CustomerService) UploadBuyerMessagesImage(filePath string) *customerserviceentity.UploadImageResult {
	var result customerserviceentity.UploadImageResult
	values := lib.InRow{
		"@data": filePath,
	}
	err := a.Config.MultipartS3("/customer_service/202309/images/upload", nil, values, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}

func (a *CustomerService) GetPerformance(startDate, endDate string) *customerserviceentity.PerformanceResult {
	var result customerserviceentity.PerformanceResult
	params := lib.InRow{}
	if startDate != "" {
		params["start_date"] = startDate
	}
	if endDate != "" {
		params["end_date"] = endDate
	}
	err := a.Config.GetS3("/customer_service/202309/performance", params, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}
