package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

// @json
type MarkPackageAsShippedRequest struct {
	TrackingNumber     string   `json:"tracking_number,omitempty"`
	ShippingProviderId string   `json:"shipping_provider_id,omitempty"`
	OrderLineItemIds   []string `json:"order_line_item_ids,omitempty"`
}

func (e *MarkPackageAsShippedRequest) String() string {
	return lib.ObjectToString(e)
}
