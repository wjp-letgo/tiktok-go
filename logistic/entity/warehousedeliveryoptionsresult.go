package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

// @json
type WarehouseDeliveryOptionsResult struct {
	Code      int                          `json:"code"`
	Message   string                       `json:"message"`
	Data      WarehouseDeliveryOptionsData `json:"data"`
	RequestId string                       `json:"request_id"`
}

func (e *WarehouseDeliveryOptionsResult) String() string {
	return lib.ObjectToString(e)
}

// @json
type WarehouseDeliveryOptionsData struct {
	DeliveryOptions DeliveryOptions `json:"delivery_options"`
}

func (e *WarehouseDeliveryOptionsData) String() string {
	return lib.ObjectToString(e)
}

type DeliveryOptions []DeliveryOption

func (e *DeliveryOptions) String() string {
	return lib.ObjectToString(e)
}

// @json
type DeliveryOption struct {
	Id             string         `json:"id"`
	Name           string         `json:"name"`
	Type           string         `json:"type"`
	Description    string         `json:"description"`
	DimensionLimit DimensionLimit `json:"dimension_limit"`
	WeightLimit    WeightLimit    `json:"weight_limit"`
	Platform       []string       `json:"platform"`
}

func (e *DeliveryOption) String() string {
	return lib.ObjectToString(e)
}

// @json
type DimensionLimit struct {
	MaxHeight int    `json:"max_height"`
	MaxLength int    `json:"max_length"`
	MaxWidth  int    `json:"max_width"`
	Unit      string `json:"unit"`
}

func (e *DimensionLimit) String() string {
	return lib.ObjectToString(e)
}

// @json
type WeightLimit struct {
	MaxWeight int    `json:"max_weight"`
	MinWeight int    `json:"min_weight"`
	Unit      string `json:"unit"`
}

func (e *WeightLimit) String() string {
	return lib.ObjectToString(e)
}
