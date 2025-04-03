package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

// @json
type ShippingProvidersResult struct {
	Code      int                   `json:"code"`
	Message   string                `json:"message"`
	Data      ShippingProvidersData `json:"data"`
	RequestId string                `json:"request_id"`
}

func (e *ShippingProvidersResult) String() string {
	return lib.ObjectToString(e)
}

// @json
type ShippingProvidersData struct {
	ShippingProviders ShippingProviders `json:"shipping_providers"`
}

func (e *ShippingProvidersData) String() string {
	return lib.ObjectToString(e)
}

type ShippingProviders []ShippingProvider

func (e *ShippingProviders) String() string {
	return lib.ObjectToString(e)
}

// @json
type ShippingProvider struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func (e *ShippingProvider) String() string {
	return lib.ObjectToString(e)
}
