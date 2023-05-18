package order

import (
	"github.com/wjpxxx/letgo/lib"
	tiktokConfig "github.com/wjpxxx/tiktokgo/config"
	orderentity "github.com/wjpxxx/tiktokgo/order/entity"
)

type order struct {
	Config *tiktokConfig.Config
}

//GetOrderList 获取订单列表
func (o *order) GetOrderList(deliveryOptionType, sortType, pageSize, startTime, endTime, orderStatus, createTimeFrom, createTimeTo, updateTimeFrom, updateTimeTo int, sortBy, cursor string) *orderentity.GetOrderListResult {
	var result orderentity.GetOrderListResult
	params := lib.InRow{
		"page_size": pageSize,
	}
	if startTime > 0 {
		params["start_time"] = startTime
	}
	if endTime > 0 {
		params["end_time"] = endTime
	}
	if orderStatus > 0 {
		params["order_status"] = orderStatus
	}
	if sortBy != "" {
		params["sort_by"] = sortBy
	}
	if cursor != "" {
		params["cursor"] = cursor
	}
	if createTimeFrom > 0 {
		params["create_time_from"] = createTimeFrom
	}
	if createTimeTo > 0 {
		params["create_time_to"] = createTimeTo
	}
	if updateTimeFrom > 0 {
		params["update_time_from"] = updateTimeFrom
	}
	if updateTimeTo > 0 {
		params["update_time_to"] = updateTimeTo
	}
	if sortType > 0 {
		params["sort_type"] = sortType
	}
	if deliveryOptionType > 0 {
		params["delivery_option_type"] = deliveryOptionType
	}

	o.Config.HttpPost("/api/orders/search", params, &result)
	return &result
}

//GetOrderDetail 获取订单详情
func (o *order) GetOrderDetail(OrderIdList []string) *orderentity.GetOrderDetailResult {
	var result orderentity.GetOrderDetailResult
	params := lib.InRow{
		"order_id_list": OrderIdList,
	}
	o.Config.HttpPost("/api/orders/search", params, &result)
	return &result
}
