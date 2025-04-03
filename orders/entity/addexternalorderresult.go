package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

// @json
type AddExternalOrderResult struct {
	Code      int                  `json:"code"`
	Message   string               `json:"message"`
	Data      AddExternalOrderData `json:"data"`
	RequestId string               `json:"request_id"`
}

func (e *AddExternalOrderResult) String() string {
	return lib.ObjectToString(e)
}

// @json
type AddExternalOrderData struct {
	Errors Errors `json:"errors"`
}

func (e *AddExternalOrderData) String() string {
	return lib.ObjectToString(e)
}

type Errors []Error

func (e *Errors) String() string {
	return lib.ObjectToString(e)
}

// @json
type Error struct {
	Code    string                 `json:"code"`
	Message string                 `json:"message"`
	Detail  AddExternalOrderDetail `json:"detail"`
}

func (e *Error) String() string {
	return lib.ObjectToString(e)
}

// @json
type AddExternalOrderDetail struct {
	OrderId       string              `json:"order_id"`
	ExternalOrder ExternalOrderDetail `json:"external_order"`
}

func (e *AddExternalOrderDetail) String() string {
	return lib.ObjectToString(e)
}

// @json
type ExternalOrderDetail struct {
	Id       string `json:"id"`
	Platform string `json:"platform"`
}

func (e *ExternalOrderDetail) String() string {
	return lib.ObjectToString(e)
}
