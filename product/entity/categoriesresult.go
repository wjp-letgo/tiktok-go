package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

// @json
type CategoriesResult struct {
	Code      int            `json:"code"`
	Message   string         `json:"message"`
	Data      CategoriesData `json:"data"`
	RequestId string         `json:"request_id"`
}

func (e *CategoriesResult) String() string {
	return lib.ObjectToString(e)
}

// @json
type CategoriesData struct {
	Categories Categories `json:"categories"`
}

func (e *CategoriesData) String() string {
	return lib.ObjectToString(e)
}

// @json
type Categories []Categorie

func (e *Categories) String() string {
	return lib.ObjectToString(e)
}

// @json
type Categorie struct {
	Id                 string   `json:"id"`
	ParentId           string   `json:"parent_id"`
	LocalName          string   `json:"local_name"`
	IsLeaf             bool     `json:"is_leaf"`
	PermissionStatuses []string `json:"permission_statuses"`
}

func (e *Categorie) String() string {
	return lib.ObjectToString(e)
}
