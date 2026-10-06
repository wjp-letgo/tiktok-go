package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

// @json
type UpdatePackageDeliveryStatusRequest struct {
	Packages DeliveryPackages `json:"packages,omitempty"`
}

func (e *UpdatePackageDeliveryStatusRequest) String() string {
	return lib.ObjectToString(e)
}

type DeliveryPackages []DeliveryPackage

func (e *DeliveryPackages) String() string {
	return lib.ObjectToString(e)
}

// @json
type DeliveryPackage struct {
	Id                 string `json:"id,omitempty"`
	DeliveryType       string `json:"delivery_type,omitempty"`
	FailDeliveryReason string `json:"fail_delivery_reason,omitempty"`
	FileType           string `json:"file_type,omitempty"`
	FileUrl            string `json:"file_url,omitempty"`
}

func (e *DeliveryPackage) String() string {
	return lib.ObjectToString(e)
}
