package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

// @json
type RecommendCategoryRequest struct {
	ProductTitle    string                  `json:"product_title,omitempty"`
	Description     string                  `json:"description,omitempty"`
	Images          RecommendCategoryImages `json:"images,omitempty"`
	CategoryVersion string                  `json:"category_version,omitempty"`
	ListingPlatform string                  `json:"listing_platform,omitempty"`
}

func (e *RecommendCategoryRequest) String() string {
	return lib.ObjectToString(e)
}

// @json
type RecommendCategoryImages []RecommendCategoryImage

func (e *RecommendCategoryImages) String() string {
	return lib.ObjectToString(e)
}

// @json
type RecommendCategoryImage struct {
	Uri string `json:"uri"`
}

func (e *RecommendCategoryImage) String() string {
	return lib.ObjectToString(e)
}
