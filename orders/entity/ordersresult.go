package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

// @json
type OrdersResult struct {
	Code      int        `json:"code"`
	Message   string     `json:"message"`
	Data      OrdersData `json:"data"`
	RequestId string     `json:"request_id"`
}

func (e *OrdersResult) String() string {
	return lib.ObjectToString(e)
}

// @json
type OrdersData struct {
	NextPageToken string `json:"next_page_token"`
	TotalCount    int    `json:"total_count"`
	Orders        Orders `json:"orders"`
}

func (e *OrdersData) String() string {
	return lib.ObjectToString(e)
}

type Orders []Order

func (e *Orders) String() string {
	return lib.ObjectToString(e)
}

// @json
type Order struct {
	Id                                 string           `json:"id"`
	BuyerMessage                       string           `json:"buyer_message"`
	CancellationInitiator              string           `json:"cancellation_initiator"`
	ShippingProviderId                 string           `json:"shipping_provider_id"`
	CreateTime                         int              `json:"create_time"`
	ShippingProvider                   string           `json:"shipping_provider"`
	Packages                           Packages         `json:"packages"`
	Payment                            Payment          `json:"payment"`
	RecipientAddress                   RecipientAddress `json:"recipient_address"`
	Status                             string           `json:"status"`
	FulfillmentType                    string           `json:"fulfillment_type"`
	DeliveryType                       string           `json:"delivery_type"`
	PaidTime                           int              `json:"paid_time"`
	RtsSlaTime                         int              `json:"rts_sla_time"`
	TtsSlaTime                         int              `json:"tts_sla_time"`
	CancelReason                       string           `json:"cancel_reason"`
	UpdateTime                         int              `json:"update_time"`
	PaymentMethodName                  string           `json:"payment_method_name"`
	RtsTime                            int              `json:"rts_time"`
	TrackingNumber                     string           `json:"tracking_number"`
	SplitOrCombineTag                  string           `json:"split_or_combine_tag"`
	HasUpdatedRecipientAddress         bool             `json:"has_updated_recipient_address"`
	CancelOrderSlaTime                 int              `json:"cancel_order_sla_time"`
	WarehouseId                        string           `json:"warehouse_id"`
	RequestCancelTime                  int              `json:"request_cancel_time"`
	ShippingType                       string           `json:"shipping_type"`
	UserId                             string           `json:"user_id"`
	SellerNote                         string           `json:"seller_note"`
	DeliverySlaTime                    int              `json:"delivery_sla_time"`
	IsCod                              bool             `json:"is_cod"`
	DeliveryOptionId                   string           `json:"delivery_option_id"`
	CancelTime                         int              `json:"cancel_time"`
	NeedUploadInvoice                  string           `json:"need_upload_invoice"`
	DeliveryOptionName                 string           `json:"delivery_option_name"`
	Cpf                                string           `json:"cpf"`
	LineItems                          LineItems        `json:"line_items"`
	BuyerEmail                         string           `json:"buyer_email"`
	DeliveryDueTime                    int              `json:"delivery_due_time"`
	IsSampleOrder                      bool             `json:"is_sample_order"`
	ShippingDueTime                    int              `json:"shipping_due_time"`
	CollectionDueTime                  int              `json:"collection_due_time"`
	DeliveryOptionRequiredDeliveryTime int              `json:"delivery_option_required_delivery_time"`
	IsOnHoldOrder                      bool             `json:"is_on_hold_order"`
	DeliveryTime                       int              `json:"delivery_time"`
	IsReplacementOrder                 bool             `json:"is_replacement_order"`
	CollectionTime                     int              `json:"collection_time"`
	ReplacedOrderId                    string           `json:"replaced_order_id"`
	IsBuyerRequestCancel               bool             `json:"is_buyer_request_cancel"`
	PickUpCutOffTime                   int              `json:"pick_up_cut_off_time"`
	FastDispatchSlaTime                int              `json:"fast_dispatch_sla_time"`
	CommercePlatform                   string           `json:"commerce_platform"`
	OrderType                          string           `json:"order_type"`
	ReleaseDate                        int              `json:"release_date"`
	HandlingDuration                   HandlingDuration `json:"handling_duration"`
	AutoCombineGroupId                 string           `json:"auto_combine_group_id"`
	IsExchangeOrder                    bool             `json:"is_exchange_order"`
	ExchangeSourceOrderId              string           `json:"exchange_source_order_id"`
	CpfName                            string           `json:"cpf_name"`
}

type LineItems []LineItem

func (e *LineItems) String() string {
	return lib.ObjectToString(e)
}

// @json
type LineItem struct {
	Id                   string              `json:"id"`
	SkuId                string              `json:"sku_id"`
	CombinedListingSkus  CombinedListingSkus `json:"combined_listing_skus"`
	DisplayStatus        string              `json:"display_status"`
	ProductName          string              `json:"product_name"`
	SellerSku            string              `json:"seller_sku"`
	SkuImage             string              `json:"sku_image"`
	SkuName              string              `json:"sku_name"`
	ProductId            string              `json:"product_id"`
	SalePrice            string              `json:"sale_price"`
	PlatformDiscount     string              `json:"platform_discount"`
	SellerDiscount       string              `json:"seller_discount"`
	SkuType              string              `json:"sku_type"`
	CancelReason         string              `json:"cancel_reason"`
	OriginalPrice        string              `json:"original_price"`
	RtsTime              int                 `json:"rts_time"`
	PackageStatus        string              `json:"package_status"`
	Currency             string              `json:"currency"`
	ShippingProviderName string              `json:"shipping_provider_name"`
	CancelUser           string              `json:"cancel_user"`
	ShippingProviderId   string              `json:"shipping_provider_id"`
	IsGift               bool                `json:"is_gift"`
	ItemTax              ItemTaxs            `json:"item_tax"`
	TrackingNumber       string              `json:"tracking_number"`
	PackageId            string              `json:"package_id"`
	RetailDeliveryFee    string              `json:"retail_delivery_fee"`
	BuyerServiceFee      string              `json:"buyer_service_fee"`
	SmallOrderFee        string              `json:"small_order_fee"`
	HandlingDurationDays string              `json:"handling_duration_days"`
	IsDangerousGood      bool                `json:"is_dangerous_good"`
}

func (e *LineItem) String() string {
	return lib.ObjectToString(e)
}

func (e *Order) String() string {
	return lib.ObjectToString(e)
}

type Packages []Package

func (e *Packages) String() string {
	return lib.ObjectToString(e)
}

// @json
type Package struct {
	Id string `json:"id"`
}

func (e *Package) String() string {
	return lib.ObjectToString(e)
}

// @json
type Payment struct {
	Currency                    string `json:"currency"`
	SubTotal                    string `json:"sub_total"`
	ShippingFee                 string `json:"shipping_fee"`
	SellerDiscount              string `json:"seller_discount"`
	PlatformDiscount            string `json:"platform_discount"`
	TotalAmount                 string `json:"total_amount"`
	OriginalTotalProductPrice   string `json:"original_total_product_price"`
	OriginalShippingFee         string `json:"original_shipping_fee"`
	ShippingFeeSellerDiscount   string `json:"shipping_fee_seller_discount"`
	ShippingFeePlatformDiscount string `json:"shipping_fee_platform_discount"`
	ShippingFeeCofundedDiscount string `json:"shipping_fee_cofunded_discount"`
	Tax                         string `json:"tax"`
	SmallOrderFee               string `json:"small_order_fee"`
	ShippingFeeTax              string `json:"shipping_fee_tax"`
	ProductTax                  string `json:"product_tax"`
	RetailDeliveryFee           string `json:"retail_delivery_fee"`
	BuyerServiceFee             string `json:"buyer_service_fee"`
	HandlingFee                 string `json:"handling_fee"`
	ShippingInsuranceFee        string `json:"shipping_insurance_fee"`
	ItemInsuranceFee            string `json:"item_insurance_fee"`
}

func (e *Payment) String() string {
	return lib.ObjectToString(e)
}

// @json
type RecipientAddress struct {
	FullAddress         string              `json:"full_address"`
	PhoneNumber         string              `json:"phone_number"`
	Name                string              `json:"name"`
	FirstName           string              `json:"first_name"`
	LastName            string              `json:"last_name"`
	AddressDetail       string              `json:"address_detail"`
	AddressLine1        string              `json:"address_line1"`
	AddressLine2        string              `json:"address_line2"`
	AddressLine3        string              `json:"address_line3"`
	AddressLine4        string              `json:"address_line4"`
	DistrictInfo        DistrictInfos       `json:"district_info"`
	DeliveryPreferences DeliveryPreferences `json:"delivery_preferences"`
	PostalCode          string              `json:"postal_code"`
	RegionCode          string              `json:"region_code"`
}

func (e *RecipientAddress) String() string {
	return lib.ObjectToString(e)
}

type DistrictInfos []DistrictInfo

func (e *DistrictInfos) String() string {
	return lib.ObjectToString(e)
}

// @json
type DistrictInfo struct {
	AddressLevelName string `json:"address_level_name"`
	AddressName      string `json:"address_name"`
	AddressLevel     string `json:"address_level"`
}

func (e *DistrictInfo) String() string {
	return lib.ObjectToString(e)
}

// @json
type DeliveryPreferences struct {
	DropOffLocation string `json:"drop_off_location"`
}

func (e *DeliveryPreferences) String() string {
	return lib.ObjectToString(e)
}

type CombinedListingSkus []CombinedListingSku

func (e *CombinedListingSkus) String() string {
	return lib.ObjectToString(e)
}

// @json
type CombinedListingSku struct {
	SkuId     string `json:"sku_id"`
	SkuCount  int    `json:"sku_count"`
	ProductId string `json:"product_id"`
	SellerSku string `json:"seller_sku"`
}

func (e *CombinedListingSku) String() string {
	return lib.ObjectToString(e)
}

type ItemTaxs []ItemTax

func (e *ItemTaxs) String() string {
	return lib.ObjectToString(e)
}

// @json
type ItemTax struct {
	TaxType   string `json:"tax_type"`
	TaxAmount string `json:"tax_amount"`
	TaxRate   string `json:"tax_rate"`
}

func (e *ItemTax) String() string {
	return lib.ObjectToString(e)
}

// @json
type HandlingDuration struct {
	Days string `json:"days"`
	Type string `json:"type"`
}

func (e *HandlingDuration) String() string {
	return lib.ObjectToString(e)
}
