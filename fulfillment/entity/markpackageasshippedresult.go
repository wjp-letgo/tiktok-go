package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

// @json
type MarkPackageAsShippedResult struct {
	Code      int                      `json:"code"`
	Message   string                   `json:"message"`
	Data      MarkPackageAsShippedData `json:"data"`
	RequestId string                   `json:"request_id"`
}

func (e *MarkPackageAsShippedResult) String() string {
	return lib.ObjectToString(e)
}

// @json
type MarkPackageAsShippedData struct {
	OrderId          string                      `json:"order_id"`
	OrderLineItemIds []string                    `json:"order_line_item_ids"`
	PackageId        string                      `json:"package_id"`
	Warning          MarkPackageAsShippedWarning `json:"warning"`
}

func (e *MarkPackageAsShippedData) String() string {
	return lib.ObjectToString(e)
}

// @json
type MarkPackageAsShippedWarning struct {
	Message string `json:"message"`
}

func (e *MarkPackageAsShippedWarning) String() string {
	return lib.ObjectToString(e)
}
