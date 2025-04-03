package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

// @json
type UpdateShippingInfoRequest struct {
	TrackingNumber string `json:"tracking_number,omitempty"`
	ShippingProviderId string `json:"shipping_provider_id,omitempty"`
}

func (e *UpdateShippingInfoRequest) String() string {
	return lib.ObjectToString(e)
}