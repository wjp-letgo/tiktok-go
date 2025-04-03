package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

// @json
type PackageDetailResult struct {
	Code      int               `json:"code"`
	Message   string            `json:"message"`
	Data      PackageDetailData `json:"data"`
	RequestId string            `json:"request_id"`
}

func (e *PackageDetailResult) String() string {
	return lib.ObjectToString(e)
}

// @json
type PackageDetailData struct {
	PackageId              string           `json:"package_id"`
	Orders                 Orders           `json:"orders"`
	PackageStatus          string           `json:"package_status"`
	SplitAndCombineTag     string           `json:"split_and_combine_tag"`
	HasMultiSkus           bool             `json:"has_multi_skus"`
	NoteTag                string           `json:"note_tag"`
	ShippingProviderName   string           `json:"shipping_provider_name"`
	ShippingProviderId     string           `json:"shipping_provider_id"`
	ShippingType           string           `json:"shipping_type"`
	DeliveryOptionName     string           `json:"delivery_option_name"`
	DeliveryOptionId       string           `json:"delivery_option_id"`
	TrackingNumber         string           `json:"tracking_number"`
	LastMileTrackingNumber string           `json:"last_mile_tracking_number"`
	PickupSlot             ShipPickupSlot   `json:"pickup_slot"`
	CreateTime             int              `json:"create_time"`
	HandoverMethod         string           `json:"handover_method"`
	OrderLineItemIds       []string         `json:"order_line_item_ids"`
	RecipientAddress       RecipientAddress `json:"recipient_address"`
	SenderAddress          SenderAddress    `json:"sender_address"`
	Weight                 Weight           `json:"weight"`
	Dimension              Dimension        `json:"dimension"`
	UpdateTime             int              `json:"update_time"`
	Insurance              Insurance        `json:"insurance"`
}

func (e *PackageDetailData) String() string {
	return lib.ObjectToString(e)
}

// @json
type RecipientAddress struct {
	FullAddress   string `json:"full_address"`
	PhoneNumber   string `json:"phone_number"`
	Name          string `json:"name"`
	PostalCode    string `json:"postal_code"`
	AddressDetail string `json:"address_detail"`
	RegionCode    string `json:"region_code"`
	AddressLine1  string `json:"address_line1"`
	AddressLine2  string `json:"address_line2"`
	AddressLine3  string `json:"address_line3"`
	AddressLine4  string `json:"address_line4"`
}

func (e *RecipientAddress) String() string {
	return lib.ObjectToString(e)
}

// @json
type SenderAddress struct {
	FullAddress   string `json:"full_address"`
	PhoneNumber   string `json:"phone_number"`
	Name          string `json:"name"`
	PostalCode    string `json:"postal_code"`
	AddressDetail string `json:"address_detail"`
	RegionCode    string `json:"region_code"`
	AddressLine1  string `json:"address_line1"`
	AddressLine2  string `json:"address_line2"`
	AddressLine3  string `json:"address_line3"`
	AddressLine4  string `json:"address_line4"`
}

func (e *SenderAddress) String() string {
	return lib.ObjectToString(e)
}

// @json
type Insurance struct {
	IsPurchased     bool   `json:"is_purchased"`
	CoverageAmount  string `json:"coverage_amount"`
	IsClaimEligible bool   `json:"is_claim_eligible"`
	ClaimStatus     string `json:"claim_status"`
}

func (e *Insurance) String() string {
	return lib.ObjectToString(e)
}
