package event

import (
	"github.com/wjp-letgo/letgo/lib"
	tiktokConfig "github.com/wjp-letgo/tiktok-go/config"
	evententity "github.com/wjp-letgo/tiktok-go/event/entity"
)

// Event
type Event struct {
	Config *tiktokConfig.Config
}

// 获取店铺已订阅的 webhook
func (a *Event) GetShopWebhooks() *evententity.WebhooksResult {
	var result evententity.WebhooksResult
	params := lib.InRow{}
	err := a.Config.GetS3("/event/202309/webhooks", params, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}

// 更新店铺 webhook
func (a *Event) UpdateShopWebhook(body *evententity.UpdateWebhookRequest) *evententity.UpdateWebhookResult {
	var result evententity.UpdateWebhookResult
	err := a.Config.PutS3("/event/202309/webhooks", nil, body, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}

// 删除店铺 webhook
func (a *Event) DeleteShopWebhook(body *evententity.DeleteWebhookRequest) *evententity.DeleteWebhookResult {
	var result evententity.DeleteWebhookResult
	err := a.Config.DeleteS3("/event/202309/webhooks", nil, body, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}
