package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

// @json
type GlobalSellerWarehouseResult struct {
	Code      int                       `json:"code"`
	Message   string                    `json:"message"`
	Data      GlobalSellerWarehouseData `json:"data"`
	RequestId string                    `json:"request_id"`
}

func (e *GlobalSellerWarehouseResult) String() string {
	return lib.ObjectToString(e)
}

// @json
type GlobalSellerWarehouseData struct {
	GlobalWarehouses GlobalWarehouses `json:"global_warehouses"`
}

func (e *GlobalSellerWarehouseData) String() string {
	return lib.ObjectToString(e)
}

type GlobalWarehouses []GlobalWarehouse

func (e *GlobalWarehouses) String() string {
	return lib.ObjectToString(e)
}

// @json
type GlobalWarehouse struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Ownership string `json:"ownership"`
}

func (e *GlobalWarehouse) String() string {
	return lib.ObjectToString(e)
}
