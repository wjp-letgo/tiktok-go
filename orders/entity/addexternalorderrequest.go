package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

// @json
type AddExternalOrderRequest struct {
	Orders AddExternalOrders `json:"orders,omitempty"`
}

func (e *AddExternalOrderRequest) String() string {
	return lib.ObjectToString(e)
}

type AddExternalOrders []AddExternalOrder

func (e *AddExternalOrders) String() string {
	return lib.ObjectToString(e)
}

// @json
type AddExternalOrder struct {
	Id            string        `json:"id,omitempty"`
	ExternalOrder ExternalOrder `json:"external_order,omitempty"`
}

func (e *AddExternalOrder) String() string {
	return lib.ObjectToString(e)
}

// @json
type ExternalOrder struct {
	Id        string            `json:"id,omitempty"`
	Platform  string            `json:"platform,omitempty"`
	LineItems ExternalLineItems `json:"line_items,omitempty"`
}

func (e *ExternalOrder) String() string {
	return lib.ObjectToString(e)
}

type ExternalLineItems []ExternalLineItem

func (e *ExternalLineItems) String() string {
	return lib.ObjectToString(e)
}

// @json
type ExternalLineItem struct {
	Id       string `json:"id,omitempty"`
	OriginId string `json:"origin_id,omitempty"`
}

func (e *ExternalLineItem) String() string {
	return lib.ObjectToString(e)
}
