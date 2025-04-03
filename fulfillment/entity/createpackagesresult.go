package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

// @json
type CreatePackagesResult struct {
	Code      int                `json:"code"`
	Message   string             `json:"message"`
	Data      CreatePackagesData `json:"data"`
	RequestId string             `json:"request_id"`
}

func (e *CreatePackagesResult) String() string {
	return lib.ObjectToString(e)
}

// @json
type CreatePackagesData struct {
	OrderId             string              `json:"order_id"`
	OrderLineItemIds    []string            `json:"order_line_item_ids"`
	Dimension           Dimension           `json:"dimension"`
	ShippingServiceInfo ShippingServiceInfo `json:"shipping_service_info"`
	PackageId           string              `json:"package_id"`
	Weight              Weight              `json:"weight"`
	CreateTime          int                 `json:"create_time"`
}

func (e *CreatePackagesData) String() string {
	return lib.ObjectToString(e)
}

// @json
type ShippingServiceInfo struct {
	Id                   string `json:"id"`
	Name                 string `json:"name"`
	Price                string `json:"price"`
	Currency             string `json:"currency"`
	EarliestDeliveryDays int    `json:"earliest_delivery_days"`
	LatestDeliveryDays   int    `json:"latest_delivery_days"`
	ShippingProviderId   string `json:"shipping_provider_id"`
	ShippingProviderName string `json:"shipping_provider_name"`
}

func (e *ShippingServiceInfo) String() string {
	return lib.ObjectToString(e)
}
