package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

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
	Id           string  `json:"id"`
	Name         string  `json:"name"`
	EffectStatus string  `json:"effect_status"`
	Type         string  `json:"type"`
	SubType      string  `json:"sub_type"`
	IsDefault    bool    `json:"is_default"`
	Address      Address `json:"address"`
}

func (e *Warehouse) String() string {
	return lib.ObjectToString(e)
}

// @json
type Address struct {
	Region               string      `json:"region"`
	State                string      `json:"state"`
	City                 string      `json:"city"`
	Distict              string      `json:"distict"`
	Town                 string      `json:"town"`
	ContactPerson        string      `json:"contact_person"`
	FirstName            string      `json:"first_name"`
	LastName             string      `json:"last_name"`
	FirstNameLocalScript string      `json:"first_name_local_script"`
	LastNameLocalScript  string      `json:"last_name_local_script"`
	PostalCode           string      `json:"postal_code"`
	FullAddress          string      `json:"full_address"`
	RegionCode           string      `json:"region_code"`
	PhoneNumber          string      `json:"phone_number"`
	AddressLine1         string      `json:"address_line1"`
	AddressLine2         string      `json:"address_line2"`
	AddressLine3         string      `json:"address_line3"`
	AddressLine4         string      `json:"address_line4"`
	Geolocation          Geolocation `json:"geolocation"`
}

func (e *Address) String() string {
	return lib.ObjectToString(e)
}

// @json
type Geolocation struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

func (e *Geolocation) String() string {
	return lib.ObjectToString(e)
}
