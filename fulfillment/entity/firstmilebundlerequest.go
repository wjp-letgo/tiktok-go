package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

// @json
type FirstmileBundleRequest struct {
	OrderIds           []string `json:"order_ids,omitempty"`
	HandoverMethod     string   `json:"handover_method,omitempty"`
	ShippingProviderId string   `json:"shipping_provider_id,omitempty"`
	TrackingNumber     string   `json:"tracking_number,omitempty"`
	PhoneTailNumber    string   `json:"phone_tail_number,omitempty"`
}

func (e *FirstmileBundleRequest) String() string {
	return lib.ObjectToString(e)
}
