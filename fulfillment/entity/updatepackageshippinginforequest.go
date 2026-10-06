package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

// @json
type UpdatePackageShippingInfoRequest struct {
	TrackingNumber     string `json:"tracking_number,omitempty"`
	ShippingProviderId string `json:"shipping_provider_id,omitempty"`
}

func (e *UpdatePackageShippingInfoRequest) String() string {
	return lib.ObjectToString(e)
}
