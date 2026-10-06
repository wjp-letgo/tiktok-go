package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

// @json
type ConfirmPackageShipmentRequest struct {
	WarehouseProviderId string          `json:"warehouse_provider_id,omitempty"`
	Package             SyncPackageInfo `json:"package,omitempty"`
}

func (e *ConfirmPackageShipmentRequest) String() string {
	return lib.ObjectToString(e)
}

// @json
type SyncPackageInfo struct {
	PackageId      string `json:"package_id,omitempty"`
	TrackingNumber string `json:"tracking_number,omitempty"`
}

func (e *SyncPackageInfo) String() string {
	return lib.ObjectToString(e)
}

// @json
type ConfirmPackageShipmentResult struct {
	Code      int       `json:"code"`
	Message   string    `json:"message"`
	Data      EmptyData `json:"data"`
	RequestId string    `json:"request_id"`
}

func (e *ConfirmPackageShipmentResult) String() string {
	return lib.ObjectToString(e)
}

// @json
type EmptyData struct {
}

func (e *EmptyData) String() string {
	return lib.ObjectToString(e)
}
