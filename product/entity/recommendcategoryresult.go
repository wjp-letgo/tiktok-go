package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

// @json
type RecommendCategoryResult struct {
	Code      int                   `json:"code"`
	Message   string                `json:"message"`
	Data      RecommendCategoryData `json:"data"`
	RequestId string                `json:"request_id"`
}

func (e *RecommendCategoryResult) String() string {
	return lib.ObjectToString(e)
}

// @json
type RecommendCategoryData struct {
	LeafCategoryId string              `json:"leaf_category_id"`
	Categories     RecommendCategories `json:"categories"`
}

func (e *RecommendCategoryData) String() string {
	return lib.ObjectToString(e)
}

type RecommendCategories []RecommendCategorie

func (e *RecommendCategories) String() string {
	return lib.ObjectToString(e)
}

// @json
type RecommendCategorie struct {
	Id                 string   `json:"id"`
	Name               string   `json:"name"`
	Level              int      `json:"level"`
	IsLeaf             bool     `json:"is_leaf"`
	PermissionStatuses []string `json:"permission_statuses"`
}

func (e *RecommendCategorie) String() string {
	return lib.ObjectToString(e)
}
