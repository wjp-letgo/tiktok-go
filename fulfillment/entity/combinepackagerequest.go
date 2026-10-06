package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

// @json
type CombinePackageRequest struct {
	CombinablePackages CombinablePackages `json:"combinable_packages,omitempty"`
}

func (e *CombinePackageRequest) String() string {
	return lib.ObjectToString(e)
}
