package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

// @json
type PackageHandoverTimeSlotsResult struct {
	Code      int                          `json:"code"`
	Message   string                       `json:"message"`
	Data      PackageHandoverTimeSlotsData `json:"data"`
	RequestId string                       `json:"request_id"`
}

func (e *PackageHandoverTimeSlotsResult) String() string {
	return lib.ObjectToString(e)
}

// @json
type PackageHandoverTimeSlotsData struct {
	CanPickup        bool        `json:"can_pickup"`
	CanDropOff       bool        `json:"can_drop_off"`
	CanVanCollection bool        `json:"can_van_collection"`
	DropOffPointUrl  string      `json:"drop_off_point_url"`
	PickupSlots      PickupSlots `json:"pickup_slots"`
}

func (e *PackageHandoverTimeSlotsData) String() string {
	return lib.ObjectToString(e)
}

type PickupSlots []PickupSlot

func (e *PickupSlots) String() string {
	return lib.ObjectToString(e)
}

// @json
type PickupSlot struct {
	StartTime int  `json:"start_time"`
	EndTime   int  `json:"end_time"`
	Avaliable bool `json:"avaliable"`
}

func (e *PickupSlot) String() string {
	return lib.ObjectToString(e)
}
