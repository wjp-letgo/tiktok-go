package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

// @json
type OrderDetailResult struct {
	Code      int        `json:"code"`
	Message   string     `json:"message"`
	Data      OrderDetailData `json:"data"`
	RequestId string     `json:"request_id"`
}

func (e *OrderDetailResult) String() string {
	return lib.ObjectToString(e)
}


// @json
type OrderDetailData struct {
	Orders        Orders `json:"orders"`
}

func (e *OrderDetailData) String() string {
	return lib.ObjectToString(e)
}


