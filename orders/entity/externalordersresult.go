package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

// @json
type ExternalOrdersResult struct {
	Code      int                `json:"code"`
	Message   string             `json:"message"`
	Data      ExternalOrdersData `json:"data"`
	RequestId string             `json:"request_id"`
}

func (e *ExternalOrdersResult) String() string {
	return lib.ObjectToString(e)
}

// @json
type ExternalOrdersData struct {
	ExternalOrders ExternalOrderList `json:"external_orders"`
}

func (e *ExternalOrdersData) String() string {
	return lib.ObjectToString(e)
}

type ExternalOrderList []ExternalOrder

func (e *ExternalOrderList) String() string {
	return lib.ObjectToString(e)
}
