package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

// @json
type UpdateShippingInfoResult struct {
	Code      int             `json:"code"`
	Message   string          `json:"message"`
	Data      UpdateShippingInfoData `json:"data"`
	RequestId string          `json:"request_id"`
}

func (e *UpdateShippingInfoResult) String() string {
	return lib.ObjectToString(e)
}

//@json
type UpdateShippingInfoData struct{

}

func (e *UpdateShippingInfoData) String() string {
	return lib.ObjectToString(e)
}