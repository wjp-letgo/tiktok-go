package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

// @json
type EligibleShippingServiceResult struct {
	Code      int                         `json:"code"`
	Message   string                      `json:"message"`
	Data      EligibleShippingServiceData `json:"data"`
	RequestId string                      `json:"request_id"`
}

func (e *EligibleShippingServiceResult) String() string {
	return lib.ObjectToString(e)
}

// @json
type EligibleShippingServiceData struct {
	OrderId          string           `json:"order_id"`
	OrderLineId      []string         `json:"order_line_id"`
	Weight           Weight           `json:"weight"`
	ShippingServices ShippingServices `json:"shipping_services"`
	Dimension        Dimension        `json:"dimension"`
}

func (e *EligibleShippingServiceData) String() string {
	return lib.ObjectToString(e)
}

type ShippingServices []ShippingService

func (e *ShippingServices) String() string {
	return lib.ObjectToString(e)
}

// @json
type ShippingService struct {
	Id                   string `json:"id"`
	Name                 string `json:"name"`
	Price                string `json:"price"`
	Currency             string `json:"currency"`
	EarliestDeliveryDays int    `json:"earliest_delivery_days"`
	LatestDeliveryDays   int    `json:"latest_delivery_days"`
	IsDefault            bool   `json:"is_default"`
	ShippingProviderName string `json:"shipping_provider_name"`
	ShippingProviderId   string `json:"shipping_provider_id"`
}

func (e *ShippingService) String() string {
	return lib.ObjectToString(e)
}
