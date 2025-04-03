package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

// @json
type SplitOrdersResult struct {
	Code      int             `json:"code"`
	Message   string          `json:"message"`
	Data      SplitOrdersData `json:"data"`
	RequestId string          `json:"request_id"`
}

func (e *SplitOrdersResult) String() string {
	return lib.ObjectToString(e)
}

// @json
type SplitOrdersData struct {
	Packages Packages `json:"packages"`
}

func (e *SplitOrdersData) String() string {
	return lib.ObjectToString(e)
}

type Packages []Package

func (e *Packages) String() string {
	return lib.ObjectToString(e)
}

// @json
type Package struct {
	SplittableGroupId string `json:"splittable_group_id"`
	Id                string `json:"id"`
}

func (e *Package) String() string {
	return lib.ObjectToString(e)
}
