package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

// @json
type FirstmileBundleResult struct {
	Code      int                 `json:"code"`
	Message   string              `json:"message"`
	Data      FirstmileBundleData `json:"data"`
	RequestId string              `json:"request_id"`
}

func (e *FirstmileBundleResult) String() string {
	return lib.ObjectToString(e)
}

// @json
type FirstmileBundleData struct {
	FirstMileBundleId string `json:"first_mile_bundle_id"`
	Url               string `json:"url"`
	Errors            Errors `json:"errors"`
}

func (e *FirstmileBundleData) String() string {
	return lib.ObjectToString(e)
}

type Errors []Error

func (e *Errors) String() string {
	return lib.ObjectToString(e)
}

// @json
type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Detail  Detail `json:"detail"`
}

func (e *Error) String() string {
	return lib.ObjectToString(e)
}

// @json
type Detail struct {
	OrderId string `json:"order_id"`
}

func (e *Detail) String() string {
	return lib.ObjectToString(e)
}
