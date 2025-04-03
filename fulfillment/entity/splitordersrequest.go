package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

// @json
type SplitOrdersRequest struct {
	SplittableGroups SplittableGroups `json:"splittable_groups,omitempty"`
}

func (e *SplitOrdersRequest) String() string {
	return lib.ObjectToString(e)
}

type SplittableGroups []SplittableGroup

func (e *SplittableGroups) String() string {
	return lib.ObjectToString(e)
}

// @json
type SplittableGroup struct {
	Id               string   `json:"id,omitempty"`
	OrderLineItemIds []string `json:"order_line_item_ids,omitempty"`
}

func (e *SplittableGroup) String() string {
	return lib.ObjectToString(e)
}
