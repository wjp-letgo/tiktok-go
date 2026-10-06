package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

// @json
type UpdatePackageShippingInfoResult struct {
	Code      int                            `json:"code"`
	Message   string                         `json:"message"`
	Data      UpdatePackageShippingInfoData  `json:"data"`
	RequestId string                         `json:"request_id"`
}

func (e *UpdatePackageShippingInfoResult) String() string {
	return lib.ObjectToString(e)
}

// @json
type UpdatePackageShippingInfoData struct {
}

func (e *UpdatePackageShippingInfoData) String() string {
	return lib.ObjectToString(e)
}
