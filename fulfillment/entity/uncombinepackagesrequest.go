package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

// @json
type UncombinePackagesRequest struct {
	OrderIds []string `json:"order_ids,omitempty"`
}

func (e *UncombinePackagesRequest) String() string {
	return lib.ObjectToString(e)
}
