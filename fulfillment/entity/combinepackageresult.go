package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

// @json
type CombinePackageResult struct {
	Code      int                `json:"code"`
	Message   string             `json:"message"`
	Data      CombinePackageData `json:"data"`
	RequestId string             `json:"request_id"`
}

func (e *CombinePackageResult) String() string {
	return lib.ObjectToString(e)
}

// @json
type CombinePackageData struct {
	Packages CombinePackages     `json:"packages"`
	Errors   CombinePackageErrors `json:"errors"`
}

func (e *CombinePackageData) String() string {
	return lib.ObjectToString(e)
}

type CombinePackages []CombinePackageItem

func (e *CombinePackages) String() string {
	return lib.ObjectToString(e)
}

// @json
type CombinePackageItem struct {
	Id       string   `json:"id"`
	OrderIds []string `json:"order_ids"`
}

func (e *CombinePackageItem) String() string {
	return lib.ObjectToString(e)
}

type CombinePackageErrors []CombinePackageError

func (e *CombinePackageErrors) String() string {
	return lib.ObjectToString(e)
}

// @json
type CombinePackageError struct {
	Code    int                     `json:"code"`
	Message string                  `json:"message"`
	Detail  CombinePackageErrorDetail `json:"detail"`
}

func (e *CombinePackageError) String() string {
	return lib.ObjectToString(e)
}

// @json
type CombinePackageErrorDetail struct {
	PackageId string `json:"package_id"`
}

func (e *CombinePackageErrorDetail) String() string {
	return lib.ObjectToString(e)
}
