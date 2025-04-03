package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

// @json
type PriceDetailResult struct {
	Code      int             `json:"code"`
	Message   string          `json:"message"`
	Data      PriceDetailData `json:"data"`
	RequestId string          `json:"request_id"`
}

func (e *PriceDetailResult) String() string {
	return lib.ObjectToString(e)
}

// @json
type PriceDetailData struct {
	Currency                            string               `json:"currency"`
	Total                               string               `json:"total"`
	Payment                             string               `json:"payment"`
	SkuListPrice                        string               `json:"sku_list_price"`
	SkuSalePrice                        string               `json:"sku_sale_price"`
	Subtotal                            string               `json:"subtotal"`
	SubtotalDeductionSeller             string               `json:"subtotal_deduction_seller"`
	SubtotalDeductionPlatform           string               `json:"subtotal_deduction_platform"`
	SubtotalTaxAmount                   string               `json:"subtotal_tax_amount"`
	VoucherDeductionPlatform            string               `json:"voucher_deduction_platform"`
	VoucherDeductionSeller              string               `json:"voucher_deduction_seller"`
	ShippingListPrice                   string               `json:"shipping_list_price"`
	ShippingSalePrice                   string               `json:"shipping_sale_price"`
	ShippingFeeDeductionSeller          string               `json:"shipping_fee_deduction_seller"`
	ShippingFeeDeductionPlatform        string               `json:"shipping_fee_deduction_platform"`
	ShippingFeeDeductionPlatformVoucher string               `json:"shipping_fee_deduction_platform_voucher"`
	TaxAmount                           string               `json:"tax_amount"`
	TaxRate                             string               `json:"tax_rate"`
	NetPriceAmount                      string               `json:"net_price_amount"`
	CodFee                              string               `json:"cod_fee"`
	CodFeeNetAmount                     string               `json:"cod_fee_net_amount"`
	SkuGiftOriginalPrice                string               `json:"sku_gift_original_price"`
	SkuGiftNetPrice                     string               `json:"sku_gift_net_price"`
	LineItems                           PriceDetailLineItems `json:"line_items"`
}

func (e *PriceDetailData) String() string {
	return lib.ObjectToString(e)
}

type PriceDetailLineItems []PriceDetailLineItem

func (e *PriceDetailLineItems) String() string {
	return lib.ObjectToString(e)
}

// @json
type PriceDetailLineItem struct {
	Id                                  string `json:"id"`
	Currency                            string `json:"currency"`
	Total                               string `json:"total"`
	Payment                             string `json:"payment"`
	SkuListPrice                        string `json:"sku_list_price"`
	SkuSalePrice                        string `json:"sku_sale_price"`
	Subtotal                            string `json:"subtotal"`
	SubtotalDeductionSeller             string `json:"subtotal_deduction_seller"`
	SubtotalDeductionPlatform           string `json:"subtotal_deduction_platform"`
	SubtotalTaxAmount                   string `json:"subtotal_tax_amount"`
	VoucherDeductionPlatform            string `json:"voucher_deduction_platform"`
	VoucherDeductionSeller              string `json:"voucher_deduction_seller"`
	ShippingListPrice                   string `json:"shipping_list_price"`
	ShippingSalePrice                   string `json:"shipping_sale_price"`
	ShippingFeeDeductionSeller          string `json:"shipping_fee_deduction_seller"`
	ShippingFeeDeductionPlatform        string `json:"shipping_fee_deduction_platform"`
	ShippingFeeDeductionPlatformVoucher string `json:"shipping_fee_deduction_platform_voucher"`
	TaxAmount                           string `json:"tax_amount"`
	TaxRate                             string `json:"tax_rate"`
	NetPriceAmount                      string `json:"net_price_amount"`
	CodFee                              string `json:"cod_fee"`
	CodFeeAmount                        string `json:"cod_fee_amount"`
	SkuGiftOriginalPrice                string `json:"sku_gift_original_price"`
	SkuGiftNetPrice                     string `json:"sku_gift_net_price"`
}

func (e *PriceDetailLineItem) String() string {
	return lib.ObjectToString(e)
}
