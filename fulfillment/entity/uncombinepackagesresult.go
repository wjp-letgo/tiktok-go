package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

// @json
type UncombinePackagesResult struct {
	Code      int                   `json:"code"`
	Message   string                `json:"message"`
	Data      UncombinePackagesData `json:"data"`
	RequestId string                `json:"request_id"`
}

func (e *UncombinePackagesResult) String() string {
	return lib.ObjectToString(e)
}

// @json
type UncombinePackagesData struct {
	Packages CombinePackages `json:"packages"`
}

func (e *UncombinePackagesData) String() string {
	return lib.ObjectToString(e)
}
