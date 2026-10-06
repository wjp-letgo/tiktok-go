package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

// @json
type CategoryAssetsResult struct {
	Code      int                `json:"code"`
	Message   string             `json:"message"`
	Data      CategoryAssetsData `json:"data"`
	RequestId string             `json:"request_id"`
}

func (e *CategoryAssetsResult) String() string {
	return lib.ObjectToString(e)
}

// @json
type CategoryAssetsData struct {
	CategoryAssets CategoryAssets `json:"category_assets"`
}

func (e *CategoryAssetsData) String() string {
	return lib.ObjectToString(e)
}

type CategoryAssets []CategoryAsset

func (e *CategoryAssets) String() string {
	return lib.ObjectToString(e)
}

// @json
type CategoryAsset struct {
	TargetMarket string   `json:"target_market"`
	Cipher       string   `json:"cipher"`
	Category     Category `json:"category"`
}

func (e *CategoryAsset) String() string {
	return lib.ObjectToString(e)
}

// @json
type Category struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func (e *Category) String() string {
	return lib.ObjectToString(e)
}
