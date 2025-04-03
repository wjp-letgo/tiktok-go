package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

// @json
type EligibleShippingServiceRequest struct {
	OrderLineItemIds []string  `json:"order_line_item_ids,omitempty"`
	Weight           Weight    `json:"weight,omitempty"`
	Dimension        Dimension `json:"dimension,omitempty"`
}

func (e *EligibleShippingServiceRequest) String() string {
	return lib.ObjectToString(e)
}

// @json
type Weight struct {
	Value string `json:"value,omitempty"`
	Unit  string `json:"unit,omitempty"`
}

func (e *Weight) String() string {
	return lib.ObjectToString(e)
}

// @json
type Dimension struct {
	Length string `json:"length,omitempty"`
	Width  string `json:"width,omitempty"`
	Height string `json:"height,omitempty"`
	Unit   string `json:"unit,omitempty"`
}

func (e *Dimension) String() string {
	return lib.ObjectToString(e)
}
