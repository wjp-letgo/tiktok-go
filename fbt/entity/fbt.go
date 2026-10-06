package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

// @json
type OnboardedRegionsResult struct {
	Code      int                  `json:"code"`
	Message   string               `json:"message"`
	Data      OnboardedRegionsData `json:"data"`
	RequestId string               `json:"request_id"`
}

func (e *OnboardedRegionsResult) String() string {
	return lib.ObjectToString(e)
}

// @json
type OnboardedRegionsData struct {
	Regions Regions `json:"regions"`
}

func (e *OnboardedRegionsData) String() string {
	return lib.ObjectToString(e)
}

type Regions []Region

func (e *Regions) String() string {
	return lib.ObjectToString(e)
}

// @json
type Region struct {
	Region string `json:"region"`
	Status string `json:"status"`
}

func (e *Region) String() string {
	return lib.ObjectToString(e)
}

// @json
type WarehouseListResult struct {
	Code      int               `json:"code"`
	Message   string            `json:"message"`
	Data      WarehouseListData `json:"data"`
	RequestId string            `json:"request_id"`
}

func (e *WarehouseListResult) String() string {
	return lib.ObjectToString(e)
}

// @json
type WarehouseListData struct {
	Warehouses Warehouses `json:"warehouses"`
}

func (e *WarehouseListData) String() string {
	return lib.ObjectToString(e)
}

type Warehouses []Warehouse

func (e *Warehouses) String() string {
	return lib.ObjectToString(e)
}

// @json
type Warehouse struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Region string `json:"region"`
	Type   string `json:"type"`
}

func (e *Warehouse) String() string {
	return lib.ObjectToString(e)
}

// @json
type InboundOrdersResult struct {
	Code      int               `json:"code"`
	Message   string            `json:"message"`
	Data      InboundOrdersData `json:"data"`
	RequestId string            `json:"request_id"`
}

func (e *InboundOrdersResult) String() string {
	return lib.ObjectToString(e)
}

// @json
type InboundOrdersData struct {
	InboundOrders InboundOrders `json:"inbound_orders"`
}

func (e *InboundOrdersData) String() string {
	return lib.ObjectToString(e)
}

type InboundOrders []InboundOrder

func (e *InboundOrders) String() string {
	return lib.ObjectToString(e)
}

// @json
type InboundOrder struct {
	Id         string `json:"id"`
	Status     string `json:"status"`
	CreateTime int    `json:"create_time"`
	UpdateTime int    `json:"update_time"`
	WarehouseId string `json:"warehouse_id"`
}

func (e *InboundOrder) String() string {
	return lib.ObjectToString(e)
}

// @json
type SearchInventoryRequest struct {
	GoodsIds     []string `json:"goods_ids,omitempty"`
	SkuIds       []string `json:"sku_ids,omitempty"`
	WarehouseIds []string `json:"warehouse_ids,omitempty"`
}

func (e *SearchInventoryRequest) String() string {
	return lib.ObjectToString(e)
}

// @json
type SearchInventoryResult struct {
	Code      int                 `json:"code"`
	Message   string              `json:"message"`
	Data      SearchInventoryData `json:"data"`
	RequestId string              `json:"request_id"`
}

func (e *SearchInventoryResult) String() string {
	return lib.ObjectToString(e)
}

// @json
type SearchInventoryData struct {
	NextPageToken string            `json:"next_page_token"`
	TotalCount    int               `json:"total_count"`
	Inventory     InventoryItems    `json:"inventory"`
}

func (e *SearchInventoryData) String() string {
	return lib.ObjectToString(e)
}

type InventoryItems []InventoryItem

func (e *InventoryItems) String() string {
	return lib.ObjectToString(e)
}

// @json
type InventoryItem struct {
	GoodsId     string `json:"goods_id"`
	SkuId       string `json:"sku_id"`
	WarehouseId string `json:"warehouse_id"`
	OnHand      int    `json:"on_hand"`
	Available   int    `json:"available"`
	Reserved    int    `json:"reserved"`
}

func (e *InventoryItem) String() string {
	return lib.ObjectToString(e)
}

// @json
type SearchInventoryRecordRequest struct {
	GoodsIds     []string `json:"goods_ids,omitempty"`
	SkuIds       []string `json:"sku_ids,omitempty"`
	WarehouseIds []string `json:"warehouse_ids,omitempty"`
	CreateTimeGe int      `json:"create_time_ge,omitempty"`
	CreateTimeLt int      `json:"create_time_lt,omitempty"`
}

func (e *SearchInventoryRecordRequest) String() string {
	return lib.ObjectToString(e)
}

// @json
type SearchInventoryRecordResult struct {
	Code      int                       `json:"code"`
	Message   string                    `json:"message"`
	Data      SearchInventoryRecordData `json:"data"`
	RequestId string                    `json:"request_id"`
}

func (e *SearchInventoryRecordResult) String() string {
	return lib.ObjectToString(e)
}

// @json
type SearchInventoryRecordData struct {
	NextPageToken     string            `json:"next_page_token"`
	TotalCount        int               `json:"total_count"`
	InventoryRecords  InventoryRecords  `json:"inventory_records"`
}

func (e *SearchInventoryRecordData) String() string {
	return lib.ObjectToString(e)
}

type InventoryRecords []InventoryRecord

func (e *InventoryRecords) String() string {
	return lib.ObjectToString(e)
}

// @json
type InventoryRecord struct {
	GoodsId     string `json:"goods_id"`
	SkuId       string `json:"sku_id"`
	WarehouseId string `json:"warehouse_id"`
	Type        string `json:"type"`
	Quantity    int    `json:"quantity"`
	CreateTime  int    `json:"create_time"`
}

func (e *InventoryRecord) String() string {
	return lib.ObjectToString(e)
}

// @json
type SearchGoodsRequest struct {
	GoodsIds []string `json:"goods_ids,omitempty"`
	SkuIds   []string `json:"sku_ids,omitempty"`
}

func (e *SearchGoodsRequest) String() string {
	return lib.ObjectToString(e)
}

// @json
type SearchGoodsResult struct {
	Code      int             `json:"code"`
	Message   string          `json:"message"`
	Data      SearchGoodsData `json:"data"`
	RequestId string          `json:"request_id"`
}

func (e *SearchGoodsResult) String() string {
	return lib.ObjectToString(e)
}

// @json
type SearchGoodsData struct {
	NextPageToken string `json:"next_page_token"`
	TotalCount    int    `json:"total_count"`
	Goods         Goods  `json:"goods"`
}

func (e *SearchGoodsData) String() string {
	return lib.ObjectToString(e)
}

type Goods []Good

func (e *Goods) String() string {
	return lib.ObjectToString(e)
}

// @json
type Good struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	SkuId      string `json:"sku_id"`
	Barcode    string `json:"barcode"`
	CreateTime int    `json:"create_time"`
}

func (e *Good) String() string {
	return lib.ObjectToString(e)
}
