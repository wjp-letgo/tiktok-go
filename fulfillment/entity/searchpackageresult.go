package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

// @json
type SearchPackageResult struct {
	Code      int               `json:"code"`
	Message   string            `json:"message"`
	Data      SearchPackageData `json:"data"`
	RequestId string            `json:"request_id"`
}

func (e *SearchPackageResult) String() string {
	return lib.ObjectToString(e)
}

// @json
type SearchPackageData struct {
	NextPageToken string         `json:"next_page_token"`
	TotalCount    int            `json:"total_count"`
	Packages      SearchPackages `json:"packages"`
}

func (e *SearchPackageData) String() string {
	return lib.ObjectToString(e)
}

type SearchPackages []SearchPackage

func (e *SearchPackages) String() string {
	return lib.ObjectToString(e)
}

// @json
type SearchPackage struct {
	Id                   string   `json:"id"`
	Orders               Orders   `json:"orders"`
	CreateTime           int      `json:"create_time"`
	UpdateTime           int      `json:"update_time"`
	Status               string   `json:"status"`
	TrackingNumber       string   `json:"tracking_number"`
	ShippingProviderName string   `json:"shipping_provider_name"`
	ShippingProviderId   string   `json:"shipping_provider_id"`
	OrderLineItemIds     []string `json:"order_line_item_ids"`
}

func (e *SearchPackage) String() string {
	return lib.ObjectToString(e)
}

type Orders []Order

func (e *Orders) String() string {
	return lib.ObjectToString(e)
}

// @json
type Order struct {
	Id   string `json:"id"`
	Skus Skus   `json:"skus"`
}

func (e *Order) String() string {
	return lib.ObjectToString(e)
}

type Skus []Sku

func (e *Skus) String() string {
	return lib.ObjectToString(e)
}

// @json
type Sku struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
	Quantity int    `json:"quantity"`
}

func (e *Sku) String() string {
	return lib.ObjectToString(e)
}
