package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

// @json
type BatchShipPackagesRequest struct {
	Packages BatchShipPackages `json:"packages,omitempty"`
}

func (e *BatchShipPackagesRequest) String() string {
	return lib.ObjectToString(e)
}

type BatchShipPackages []BatchShipPackage

func (e *BatchShipPackages) String() string {
	return lib.ObjectToString(e)
}

type BatchShipPackage struct {
	Id             string         `json:"id,omitempty"`
	HandoverMethod string         `json:"handover_method,omitempty"`
	PickupSlot     ShipPickupSlot `json:"pickup_slot,omitempty"`
	SelfShipment   SelfShipment   `json:"self_shipment,omitempty"`
}

func (e *BatchShipPackage) String() string {
	return lib.ObjectToString(e)
}
