package orders

import (
	"fmt"
	"strings"

	"github.com/wjp-letgo/letgo/lib"
	tiktokConfig "github.com/wjp-letgo/tiktok-go/config"
	ordersentity "github.com/wjp-letgo/tiktok-go/orders/entity"
)

// Orders
type Orders struct {
	Config *tiktokConfig.Config
}

// 订单列表
// pageSize：The number of results to be returned per page. Default: 20. Valid range: [1-100].
// Available values:
// - UNPAID: The order has been placed, but payment has not been completed.
// - ON_HOLD: The order has been accepted and is awaiting fulfillment. The buyer may still cancel without the seller’s approval. If order_type=PRE_ORDER, the product is still awaiting release so payment will only be authorized 1 day before the release, but the seller should start preparing for the release. (Applicable only for the US and UK market).
// - AWAITING_SHIPMENT: The order is ready to be shipped, but no items have been shipped yet.
// - PARTIALLY_SHIPPING: Some items in the order have been shipped, but not all.
// - AWAITING_COLLECTION: Shipping has been arranged, but the package is waiting to be collected by the carrier.
// - IN_TRANSIT: The package has been collected by the carrier and delivery is in progress.
// - DELIVERED: The package has been delivered to the buyer.
// - COMPLETED: The order has been completed, and no further returns or refunds are allowed.
// - CANCELLED: The order has been cancelled.
func (a *Orders) OrdersSearch(pageSize int, pageToken, sortOrder, sortField string, body *ordersentity.OrdersRequest) *ordersentity.OrdersResult {
	var result ordersentity.OrdersResult
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
	err := a.Config.HttpS3("/order/202309/orders/search", params, body, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}

// 获得订单详情
func (a *Orders) OrderDetail(ids []string) *ordersentity.OrderDetailResult {
	var result ordersentity.OrderDetailResult
	params := lib.InRow{
		"ids": strings.Join(ids, ","),
	}
	err := a.Config.GetS3("/order/202309/orders", params, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}

// 获取订单或行项目的详细定价计算信息，包括凭证、税费等。
func (a *Orders) PriceDetail(orderId string) *ordersentity.PriceDetailResult {
	var result ordersentity.PriceDetailResult
	params := lib.InRow{}
	err := a.Config.GetS3(fmt.Sprintf("/order/202407/orders/%s/price_detail", orderId), params, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}

func (a *Orders) AddExternalOrder(body *ordersentity.AddExternalOrderRequest) *ordersentity.AddExternalOrderResult {
	var result ordersentity.AddExternalOrderResult
	err := a.Config.HttpS3("/order/202406/orders/external_orders", nil, body, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}
