package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

// @json
type SearchProductsRequest struct {
	Status             string   `json:"status,omitempty"`
	SellerSkus         []string `json:"seller_skus,omitempty"`
	CreateTimeGe       int      `json:"create_time_ge,omitempty"`
	CreateTimeLe       int      `json:"create_time_le,omitempty"`
	UpdateTimeGe       int      `json:"update_time_ge,omitempty"`
	UpdateTimeLe       int      `json:"update_time_le,omitempty"`
	CategoryVersion    string   `json:"category_version,omitempty"`
	ListingPlatforms   []string `json:"listing_platforms,omitempty"`
	ListingQualityTier string   `json:"listing_quality_tier,omitempty"`
}

func (e *SearchProductsRequest) String() string {
	return lib.ObjectToString(e)
}

// @json
type SearchProductsResult struct {
	Code      int                `json:"code"`
	Message   string             `json:"message"`
	Data      SearchProductsData `json:"data"`
	RequestId string             `json:"request_id"`
}

func (e *SearchProductsResult) String() string {
	return lib.ObjectToString(e)
}

// @json
type SearchProductsData struct {
	NextPageToken string           `json:"next_page_token"`
	TotalCount    int              `json:"total_count"`
	Products      SearchProductList `json:"products"`
}

func (e *SearchProductsData) String() string {
	return lib.ObjectToString(e)
}

type SearchProductList []SearchProduct

func (e *SearchProductList) String() string {
	return lib.ObjectToString(e)
}

// @json
type SearchProduct struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Status      string `json:"status"`
	CreateTime  int    `json:"create_time"`
	UpdateTime  int    `json:"update_time"`
	SalesRegions []string `json:"sales_regions"`
}

func (e *SearchProduct) String() string {
	return lib.ObjectToString(e)
}

// @json
type ProductIdsRequest struct {
	ProductIds []string `json:"product_ids,omitempty"`
}

func (e *ProductIdsRequest) String() string {
	return lib.ObjectToString(e)
}

// @json
type ProductIdsResult struct {
	Code      int            `json:"code"`
	Message   string         `json:"message"`
	Data      ProductIdsData `json:"data"`
	RequestId string         `json:"request_id"`
}

func (e *ProductIdsResult) String() string {
	return lib.ObjectToString(e)
}

// @json
type ProductIdsData struct {
	ProductIds []string `json:"product_ids"`
	Errors     ProductOpErrors `json:"errors"`
}

func (e *ProductIdsData) String() string {
	return lib.ObjectToString(e)
}

type ProductOpErrors []ProductOpError

func (e *ProductOpErrors) String() string {
	return lib.ObjectToString(e)
}

// @json
type ProductOpError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Detail  ProductOpErrorDetail `json:"detail"`
}

func (e *ProductOpError) String() string {
	return lib.ObjectToString(e)
}

// @json
type ProductOpErrorDetail struct {
	ProductId string `json:"product_id"`
}

func (e *ProductOpErrorDetail) String() string {
	return lib.ObjectToString(e)
}

// @json
type UpdatePriceRequest struct {
	Skus UpdatePriceSkus `json:"skus,omitempty"`
}

func (e *UpdatePriceRequest) String() string {
	return lib.ObjectToString(e)
}

type UpdatePriceSkus []UpdatePriceSku

func (e *UpdatePriceSkus) String() string {
	return lib.ObjectToString(e)
}

// @json
type UpdatePriceSku struct {
	Id    string       `json:"id,omitempty"`
	Price ProductPrice `json:"price,omitempty"`
}

func (e *UpdatePriceSku) String() string {
	return lib.ObjectToString(e)
}

// @json
type ProductPrice struct {
	Amount    string `json:"amount,omitempty"`
	Currency  string `json:"currency,omitempty"`
	SalePrice string `json:"sale_price,omitempty"`
}

func (e *ProductPrice) String() string {
	return lib.ObjectToString(e)
}

// @json
type UpdateInventoryRequest struct {
	Skus UpdateInventorySkus `json:"skus,omitempty"`
}

func (e *UpdateInventoryRequest) String() string {
	return lib.ObjectToString(e)
}

type UpdateInventorySkus []UpdateInventorySku

func (e *UpdateInventorySkus) String() string {
	return lib.ObjectToString(e)
}

// @json
type UpdateInventorySku struct {
	Id        string            `json:"id,omitempty"`
	Inventory SkuInventories    `json:"inventory,omitempty"`
}

func (e *UpdateInventorySku) String() string {
	return lib.ObjectToString(e)
}

type SkuInventories []SkuInventory

func (e *SkuInventories) String() string {
	return lib.ObjectToString(e)
}

// @json
type SkuInventory struct {
	WarehouseId string `json:"warehouse_id,omitempty"`
	Quantity    int    `json:"quantity,omitempty"`
}

func (e *SkuInventory) String() string {
	return lib.ObjectToString(e)
}

// @json
type InventorySearchRequest struct {
	ProductIds []string `json:"product_ids,omitempty"`
	SkuIds     []string `json:"sku_ids,omitempty"`
}

func (e *InventorySearchRequest) String() string {
	return lib.ObjectToString(e)
}

// @json
type InventorySearchResult struct {
	Code      int                 `json:"code"`
	Message   string              `json:"message"`
	Data      InventorySearchData `json:"data"`
	RequestId string              `json:"request_id"`
}

func (e *InventorySearchResult) String() string {
	return lib.ObjectToString(e)
}

// @json
type InventorySearchData struct {
	Inventory InventorySearchItems `json:"inventory"`
}

func (e *InventorySearchData) String() string {
	return lib.ObjectToString(e)
}

type InventorySearchItems []InventorySearchItem

func (e *InventorySearchItems) String() string {
	return lib.ObjectToString(e)
}

// @json
type InventorySearchItem struct {
	ProductId string         `json:"product_id"`
	SkuId     string         `json:"sku_id"`
	Warehouse SkuInventories `json:"warehouse_inventory"`
}

func (e *InventorySearchItem) String() string {
	return lib.ObjectToString(e)
}

// @json
type EmptyProductData struct {
}

func (e *EmptyProductData) String() string {
	return lib.ObjectToString(e)
}

// @json
type EmptyProductResult struct {
	Code      int              `json:"code"`
	Message   string           `json:"message"`
	Data      EmptyProductData `json:"data"`
	RequestId string           `json:"request_id"`
}

func (e *EmptyProductResult) String() string {
	return lib.ObjectToString(e)
}
