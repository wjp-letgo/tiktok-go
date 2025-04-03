package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

// @json
type ShipPackageRequest struct {
	HandoverMethod string         `json:"handover_method,omitempty"`
	PickupSlot     ShipPickupSlot `json:"pickup_slot,omitempty"`
	SelfShipment   SelfShipment   `json:"self_shipment,omitempty"`
}

func (e *ShipPackageRequest) String() string {
	return lib.ObjectToString(e)
}

// @json
type ShipPickupSlot struct {
	StartTime int `json:"start_time,omitempty"`
	EndTime   int `json:"end_time,omitempty"`
}

func (e *ShipPickupSlot) String() string {
	return lib.ObjectToString(e)
}

// @json
type SelfShipment struct {
	TrackingNumber     string `json:"tracking_number,omitempty"`
	ShippingProviderId string `json:"shipping_provider_id,omitempty"`
}

func (e *SelfShipment) String() string {
	return lib.ObjectToString(e)
}
