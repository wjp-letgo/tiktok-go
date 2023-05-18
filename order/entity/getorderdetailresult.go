package entity

import "github.com/wjpxxx/letgo/lib"

type GetOrderDetailResult struct {
	Code      int                        `json:"code"`
	Message   string                     `json:"message"`
	RequestId string                     `json:"request_id"`
	Data      []GetOrderDetailResultData `json:"data"`
}

type GetOrderDetailResultData struct {
	OrderList []OrderDetailResultDataOrderList `json:"order_list"`
}

type OrderDetailResultDataOrderList struct {
	OrderId                   string             `json:"order_id"`
	OrderStatus               int                `json:"order_status"`
	PaymentMethod             string             `json:"payment_method"`
	DeliveryOption            string             `json:"delivery_option"`
	ShippingProvider          string             `json:"shipping_provider"`
	ShippingProviderId        string             `json:"shipping_provider_id"`
	CreateTime                string             `json:"create_time"`
	PaidTime                  int                `json:"paid_time"`
	BuyerMessage              string             `json:"buyer_message"`
	PaymentInfo               PaymentInfo        `json:"payment_info"`
	RecipientAddress          RecipientAddress   `json:"recipient_address"`
	ItemList                  []ItemList         `json:"item_list"`
	CancelReason              string             `json:"cancel_reason"`
	CancelUser                string             `json:"cancel_user"`
	ExtStatus                 int                `json:"ext_status"`
	OrderStatusOld            string             `json:"order_status_old"`
	TrackingNumber            string             `json:"tracking_number"`
	RtsTime                   int                `json:"rts_time"`
	RtsSla                    int                `json:"rts_sla"`
	TtsSla                    int                `json:"tts_sla"`
	CancelOrderSla            int                `json:"cancel_order_sla"`
	UpdateTime                int                `json:"update_time"`
	PackageList               []PackageList      `json:"package_list"`
	ReceiverAddressUpdated    int                `json:"receiver_address_updated"`
	BuyerUid                  string             `json:"buyer_uid"`
	SplitOrCombineTag         string             `json:"split_or_combine_tag"`
	FulfillmentType           int                `json:"fulfillment_type"`
	SellerNote                string             `json:"seller_note"`
	WarehouseId               string             `json:"warehouse_id"`
	PaymentMethodType         int                `json:"payment_method_type"`
	PaymentMethodName         string             `json:"payment_method_name"`
	DeliveryOptionType        int                `json:"delivery_option_type"`
	DeliveryOptionDescription string             `json:"delivery_option_description"`
	DeliveryOptionId          string             `json:"delivery_option_id"`
	DeliverySla               int                `json:"delivery_sla"`
	IsCod                     bool               `json:"is_cod"`
	CodRiskType               int                `json:"cod_risk_type"`
	CodRiskReason             []string           `json:"cod_risk_reason"`
	NeedUploadInvoice         int                `json:"need_upload_invoice"`
	OrderLineList             []OrderLineList    `json:"order_line_list"`
	Cpf                       string             `json:"cpf"`
	DistrictInfoList          []DistrictInfoList `json:"district_info_list"`
}

type PaymentInfo struct {
	Currency                    string  `json:"currency"`
	SubTotal                    float32 `json:"sub_total"`
	ShippingFee                 float32 `json:"shipping_fee"`
	SellerDiscount              float32 `json:"seller_discount"`
	PlatformDiscount            float32 `json:"platform_discount"`
	TotalAmount                 float32 `json:"total_amount"`
	OriginalTotalProductPrice   float32 `json:"original_total_product_price"`
	OriginalShippingFee         float32 `json:"original_shipping_fee"`
	ShippingFeeSellerDiscount   float32 `json:"shipping_fee_seller_discount"`
	ShippingFeePlatformDiscount float32 `json:"shipping_fee_platform_discount"`
	Taxes                       float32 `json:"taxes"`
	SmallOrderFee               float32 `json:"small_order_fee"`
}

type RecipientAddress struct {
	FullAddress     string   `json:"full_address"`
	Region          string   `json:"region"`
	State           string   `json:"state"`
	City            string   `json:"city"`
	District        string   `json:"district"`
	Town            string   `json:"town"`
	Phone           string   `json:"phone"`
	Name            string   `json:"name"`
	Zipcode         string   `json:"zipcode"`
	AddressDetail   string   `json:"address_detail"`
	AddressLineList []string `json:"address_line_list"`
	RegionCode      string   `json:"region_code"`
}

type ItemList struct {
	SkuId                    string  `json:"sku_id"`
	ProductId                string  `json:"product_id"`
	SkuName                  string  `json:"sku_name"`
	Quantity                 int     `json:"quantity"`
	SellerSku                string  `json:"seller_sku"`
	ProductName              string  `json:"product_name"`
	SkuImage                 string  `json:"sku_image"`
	SkuOriginalPrice         float32 `json:"sku_original_price"`
	SkuSalePrice             float32 `json:"sku_sale_price"`
	SkuPlatformDiscount      float32 `json:"sku_platform_discount"`
	SkuSellerDiscount        float32 `json:"sku_seller_discount"`
	SkuExtStatus             int     `json:"sku_ext_status"`
	SkuDisplayStatus         int     `json:"sku_display_status"`
	SkuCancelReason          string  `json:"sku_cancel_reason"`
	SkuCancelUser            string  `json:"sku_cancel_user"`
	SkuRtsTime               int     `json:"sku_rts_time"`
	SkuType                  int     `json:"sku_type"`
	SkuPlatformDiscountTotal float32 `json:"sku_platform_discount_total"`
	SkuSmallOrderFee         float32 `json:"sku_small_order_fee"`
}

type PackageList struct {
	PackageId string `json:"package_id"`
}

type OrderLineList struct {
	OrderLineId          string  `json:"order_line_id"`
	SkuId                string  `json:"sku"`
	ExtStatus            int     `json:"ext_status"`
	DisplayStatus        int     `json:"display_status"`
	ProductId            string  `json:"product_id"`
	ProductName          string  `json:"product_name"`
	SkuName              string  `json:"sku_name"`
	SellerSku            string  `json:"seller_sku"`
	SkuImage             string  `json:"sku_image"`
	SkuType              int     `json:"sku_type"`
	OriginalPrice        float32 `json:"original_price"`
	SalePrice            float32 `json:"sale_price"`
	PlatformDiscount     float32 `json:"platform_discount"`
	SellerDiscount       float32 `json:"seller_discount"`
	RtsTime              int     `json:"rts_time"`
	CancelReason         string  `json:"cancel_reason"`
	CancelUser           string  `json:"cancel_user"`
	PackageId            int     `json:"package_id"`
	PackageStatus        int     `json:"package_status"`
	PackageFreezeStatus  int     `json:"package_freeze_status"`
	ShippingProviderId   string  `json:"shipping_provider_id"`
	ShippingProviderName string  `json:"shipping_provider_name"`
	TrackingNumber       string  `json:"tracking_number"`
}

type DistrictInfoList struct {
	AddressLevelName string `json:"address_level_name"`
	AddressName      string `json:"address_name"`
}

func (getorderdetailresult *GetOrderDetailResult) String() string {
	return lib.ObjectToString(getorderdetailresult)
}
