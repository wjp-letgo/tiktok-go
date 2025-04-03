package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

// @json
type CreatePackagesRequest struct {
	OrderId           string    `json:"order_id,omitempty"`
	OrderLineItemIds  []string  `json:"order_line_item_ids,omitempty"`
	Dimension         Dimension `json:"dimension,omitempty"`
	ShippingServiceId string    `json:"shipping_service_id,omitempty"`
	Weight            Weight    `json:"weight,omitempty"`
}

func (e *CreatePackagesRequest) String() string {
	return lib.ObjectToString(e)
}
