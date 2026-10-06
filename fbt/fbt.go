package fbt

import (
	"strings"

	"github.com/wjp-letgo/letgo/lib"
	tiktokConfig "github.com/wjp-letgo/tiktok-go/config"
	fbtentity "github.com/wjp-letgo/tiktok-go/fbt/entity"
)

// Fbt
type Fbt struct {
	Config *tiktokConfig.Config
}

func (a *Fbt) MerchantOnboardedRegions() *fbtentity.OnboardedRegionsResult {
	var result fbtentity.OnboardedRegionsResult
	params := lib.InRow{}
	err := a.Config.GetS3("/fbt/202409/merchants/onboarded_regions", params, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}

func (a *Fbt) FbtWarehouseList() *fbtentity.WarehouseListResult {
	var result fbtentity.WarehouseListResult
	params := lib.InRow{}
	err := a.Config.GetS3("/fbt/202408/warehouses", params, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}

func (a *Fbt) InboundOrders(ids []string) *fbtentity.InboundOrdersResult {
	var result fbtentity.InboundOrdersResult
	params := lib.InRow{
		"ids": strings.Join(ids, ","),
	}
	err := a.Config.GetS3("/fbt/202409/inbound_orders", params, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}

func (a *Fbt) SearchFbtInventory(pageSize int, pageToken string, body *fbtentity.SearchInventoryRequest) *fbtentity.SearchInventoryResult {
	var result fbtentity.SearchInventoryResult
	params := lib.InRow{
		"page_size": pageSize,
	}
	if pageToken != "" {
		params["page_token"] = pageToken
	}
	err := a.Config.HttpS3("/fbt/202408/inventory/search", params, body, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}

func (a *Fbt) SearchFbtInventoryRecord(pageSize int, pageToken string, body *fbtentity.SearchInventoryRecordRequest) *fbtentity.SearchInventoryRecordResult {
	var result fbtentity.SearchInventoryRecordResult
	params := lib.InRow{
		"page_size": pageSize,
	}
	if pageToken != "" {
		params["page_token"] = pageToken
	}
	err := a.Config.HttpS3("/fbt/202410/inventory_records/search", params, body, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}

func (a *Fbt) SearchFbtGoods(pageSize int, pageToken string, body *fbtentity.SearchGoodsRequest) *fbtentity.SearchGoodsResult {
	var result fbtentity.SearchGoodsResult
	params := lib.InRow{
		"page_size": pageSize,
	}
	if pageToken != "" {
		params["page_token"] = pageToken
	}
	err := a.Config.HttpS3("/fbt/202409/goods/search", params, body, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}
