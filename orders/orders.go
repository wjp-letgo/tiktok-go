package orders

import (
	"fmt"
	"strings"

	"github.com/wjp-letgo/letgo/lib"
	tiktokConfig "github.com/wjp-letgo/tiktok-go/config"
	ordersentity "github.com/wjp-letgo/tiktok-go/orders/entity"
)

//Orders
type Orders struct {
	Config *tiktokConfig.Config
}

//订单列表
//pageSize：The number of results to be returned per page. Default: 20. Valid range: [1-100].
func (a *Orders) OrdersSearch(pageSize int,pageToken,sortOrder,sortField string,body *ordersentity.OrdersRequest)*ordersentity.OrdersResult{
	var result ordersentity.OrdersResult
	params := lib.InRow{
		"page_size":pageSize,
	}
	if sortField!=""{
		params["sort_field"]=sortField
	}
	if sortOrder!=""{
		params["sort_order"]=sortOrder
	}
	if pageToken!=""{
		params["page_token"]=pageToken
	}
	err := a.Config.HttpS3("/order/202309/orders/search",params, body, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}

//获得订单详情
func (a *Orders) OrderDetail(ids []string)*ordersentity.OrderDetailResult{
	var result ordersentity.OrderDetailResult
	params := lib.InRow{
		"ids":strings.Join(ids,","),
	}
	err := a.Config.GetS3("/order/202309/orders",params, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}

//获取订单或行项目的详细定价计算信息，包括凭证、税费等。
func (a *Orders) PriceDetail(orderId string)*ordersentity.PriceDetailResult{
	var result ordersentity.PriceDetailResult
	params := lib.InRow{
	}
	err := a.Config.GetS3(fmt.Sprintf("/order/202407/orders/%s/price_detail",orderId),params, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}

func (a *Orders) AddExternalOrder(body *ordersentity.AddExternalOrderRequest)*ordersentity.AddExternalOrderResult{
	var result ordersentity.AddExternalOrderResult
	err := a.Config.HttpS3("/order/202406/orders/external_orders",nil, body, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}
