package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

// @json
type SearchExternalOrderResult struct {
	Code      int                     `json:"code"`
	Message   string                  `json:"message"`
	Data      SearchExternalOrderData `json:"data"`
	RequestId string                  `json:"request_id"`
}

func (e *SearchExternalOrderResult) String() string {
	return lib.ObjectToString(e)
}

// @json
type SearchExternalOrderData struct {
	Orders SearchExternalOrders `json:"orders"`
}

func (e *SearchExternalOrderData) String() string {
	return lib.ObjectToString(e)
}

type SearchExternalOrders []SearchExternalOrder

func (e *SearchExternalOrders) String() string {
	return lib.ObjectToString(e)
}

// @json
type SearchExternalOrder struct {
	Id            string        `json:"id"`
	ExternalOrder ExternalOrder `json:"external_order"`
}

func (e *SearchExternalOrder) String() string {
	return lib.ObjectToString(e)
}
