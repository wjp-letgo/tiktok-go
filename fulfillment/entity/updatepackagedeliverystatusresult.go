package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

// @json
type UpdatePackageDeliveryStatusResult struct {
	Code      int                            `json:"code"`
	Message   string                         `json:"message"`
	Data      UpdatePackageDeliveryStatusData `json:"data"`
	RequestId string                         `json:"request_id"`
}

func (e *UpdatePackageDeliveryStatusResult) String() string {
	return lib.ObjectToString(e)
}

// @json
type UpdatePackageDeliveryStatusData struct {
	Errors CombinePackageErrors `json:"errors"`
}

func (e *UpdatePackageDeliveryStatusData) String() string {
	return lib.ObjectToString(e)
}
