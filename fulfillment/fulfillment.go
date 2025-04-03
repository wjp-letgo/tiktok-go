package fulfillment

import (
	"fmt"
	"strings"

	"github.com/wjp-letgo/letgo/lib"
	tiktokConfig "github.com/wjp-letgo/tiktok-go/config"
	fulfillmententity "github.com/wjp-letgo/tiktok-go/fulfillment/entity"
)

// Fulfillment
type Fulfillment struct {
	Config *tiktokConfig.Config
}

// 查询下订单是否能拆单
func (a *Fulfillment) OrderSplitAttributes(orderIds []string) *fulfillmententity.OrdersSplitAttributesResult {
	var result fulfillmententity.OrdersSplitAttributesResult
	params := lib.InRow{
		"order_ids": strings.Join(orderIds, ","),
	}
	err := a.Config.GetS3("/fulfillment/202309/orders/split_attributes", params, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}

// 拆单
func (a *Fulfillment) SplitOrders(orderId string, body *fulfillmententity.SplitOrdersRequest) *fulfillmententity.SplitOrdersResult {
	var result fulfillmententity.SplitOrdersResult
	err := a.Config.HttpS3(fmt.Sprintf("/fulfillment/202309/orders/%s/split", orderId), nil, body, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}

// 当指定包裹的尺寸或重量时，使用此API（适用于美国）可查询可用运输服务的列表
func (a *Fulfillment) EligibleShippingService(orderId string, body *fulfillmententity.EligibleShippingServiceRequest) *fulfillmententity.EligibleShippingServiceResult {
	var result fulfillmententity.EligibleShippingServiceResult
	err := a.Config.HttpS3(fmt.Sprintf("/fulfillment/202309/orders/%s/shipping_services/query", orderId), nil, body, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}

// 如果您在一个第一英里捆绑包中将多个包裹发送到TikTok Shop仓库，您可以使用API在TikTokShop上创建第一英里捆绑，并获得捆绑包ID。
func (a *Fulfillment) FirstMileBundle(body *fulfillmententity.FirstmileBundleRequest) *fulfillmententity.FirstmileBundleResult {
	var result fulfillmententity.FirstmileBundleResult
	err := a.Config.HttpS4("/fulfillment/202407/bundles", nil, body, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}

// 使用此API运送订单（采购标签）。此API仅适用于美国地区。运费和交货时间仅为估计值，并基于您提供的包装尺寸和重量。根据包裹属性，下面列出的选项可能与您的配送订阅不同。
func (a *Fulfillment) CreatePackages(body *fulfillmententity.CreatePackagesRequest) *fulfillmententity.CreatePackagesResult {
	var result fulfillmententity.CreatePackagesResult
	err := a.Config.HttpS3("/fulfillment/202309/packages", nil, body, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}

// 根据指定条件检索包ID。包创建时间和信息更新时间是常见的查询条件。
func (a *Fulfillment) SearchPackage(pageSize int, pageToken, sortField, sortOrder string, body *fulfillmententity.SearchPackageRequest) *fulfillmententity.SearchPackageResult {
	var result fulfillmententity.SearchPackageResult
	params := lib.InRow{
		"page_size": pageSize,
	}
	if sortField != "" {
		params["sort_field"] = sortField
	}
	if sortOrder != "" {
		params["sort_order"] = sortOrder
	}
	if pageToken != "" {
		params["page_token"] = pageToken
	}
	err := a.Config.HttpS3("/fulfillment/202309/packages/search", params, body, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}

// 获得安排时间
func (a *Fulfillment) PackageHandoverTimeSlots(packageId string) *fulfillmententity.PackageHandoverTimeSlotsResult {
	var result fulfillmententity.PackageHandoverTimeSlotsResult
	params := lib.InRow{}
	err := a.Config.GetS3(fmt.Sprintf("/fulfillment/202309/packages/%s/handover_time_slots", packageId), params, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}

// 设置包裹发货
func (a *Fulfillment) ShipPackage(packageId string, body *fulfillmententity.ShipPackageRequest) *fulfillmententity.ShipPackageResult {
	var result fulfillmententity.ShipPackageResult
	err := a.Config.HttpS3(fmt.Sprintf("/fulfillment/202309/packages/%s/ship", packageId), nil, body, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}

//批量设置包裹发货
func (a *Fulfillment)BatchShipPackages(body *fulfillmententity.BatchShipPackagesRequest)*fulfillmententity.BatchShipPackagesResult{
	var result fulfillmententity.BatchShipPackagesResult
	err := a.Config.HttpS3("/fulfillment/202309/packages/ship", nil, body, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}

// 获得快递面单文档
func (a *Fulfillment) PackageShippingDocument(packageId, documentType, documentSize string) *fulfillmententity.PackageShippingDocumentResult {
	var result fulfillmententity.PackageShippingDocumentResult
	params := lib.InRow{
		"document_type": documentType,
	}
	if documentSize != "" {
		params["document_size"] = documentSize
	}
	err := a.Config.GetS3(fmt.Sprintf("/fulfillment/202309/packages/%s/shipping_documents", packageId), params, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}

// 获得包裹信息
func (a *Fulfillment)PackageDetail(packageId string)*fulfillmententity.PackageDetailResult{
	var result fulfillmententity.PackageDetailResult
	params := lib.InRow{
	}
	err := a.Config.GetS3(fmt.Sprintf("/fulfillment/202309/packages/%s", packageId), params, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}

//设置订单已发货
func (a *Fulfillment)UpdateShippingInfo(orderId string,body *fulfillmententity.UpdateShippingInfoRequest)*fulfillmententity.UpdateShippingInfoResult{
	var result fulfillmententity.UpdateShippingInfoResult
	err := a.Config.HttpS3(fmt.Sprintf("/fulfillment/202309/orders/%s/shipping_info/update", orderId), nil, body, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}