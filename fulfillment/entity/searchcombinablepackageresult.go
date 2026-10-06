package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

// @json
type SearchCombinablePackageResult struct {
	Code      int                         `json:"code"`
	Message   string                      `json:"message"`
	Data      SearchCombinablePackageData `json:"data"`
	RequestId string                      `json:"request_id"`
}

func (e *SearchCombinablePackageResult) String() string {
	return lib.ObjectToString(e)
}

// @json
type SearchCombinablePackageData struct {
	CombinablePackages CombinablePackages `json:"combinable_packages"`
	TotalCount         int                `json:"total_count"`
	NextPageToken      string             `json:"next_page_token"`
}

func (e *SearchCombinablePackageData) String() string {
	return lib.ObjectToString(e)
}

type CombinablePackages []CombinablePackage

func (e *CombinablePackages) String() string {
	return lib.ObjectToString(e)
}

// @json
type CombinablePackage struct {
	Id       string   `json:"id"`
	OrderIds []string `json:"order_ids"`
}

func (e *CombinablePackage) String() string {
	return lib.ObjectToString(e)
}
