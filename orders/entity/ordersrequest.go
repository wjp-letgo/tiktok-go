package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

// @json
type OrdersRequest struct {
	OrderStatus          string   `json:"order_status,omitempty"`
	CreateTimeGe         int      `json:"create_time_ge,omitempty"`
	CreateTimeLt         int      `json:"create_time_lt,omitempty"`
	UpdateTimeGe         int      `json:"update_time_ge,omitempty"`
	UpdateTimeLt         int      `json:"update_time_lt,omitempty"`
	ShippingType         string   `json:"shipping_type,omitempty"`
	BuyerUserId          string   `json:"buyer_user_id,omitempty"`
	IsBuyerRequestCancel bool     `json:"is_buyer_request_cancel,omitempty"`
	WarehouseIds         []string `json:"warehouse_ids,omitempty"`
}

func (e *OrdersRequest) String() string {
	return lib.ObjectToString(e)
}
