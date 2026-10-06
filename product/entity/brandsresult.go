package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

// @json
type BrandsResult struct {
	Code      int        `json:"code"`
	Message   string     `json:"message"`
	Data      BrandsData `json:"data"`
	RequestId string     `json:"request_id"`
}

func (e *BrandsResult) String() string {
	return lib.ObjectToString(e)
}

// @json
type BrandsData struct {
	Brands        Brands `json:"brands"`
	NextPageToken string `json:"next_page_token"`
	TotalCount    int    `json:"total_count"`
}

func (e *BrandsData) String() string {
	return lib.ObjectToString(e)
}

type Brands []Brand

func (e *Brands) String() string {
	return lib.ObjectToString(e)
}

// @json
type Brand struct {
	Id               string `json:"id"`
	Name             string `json:"name"`
	AuthorizedStatus string `json:"authorized_status"`
}

func (e *Brand) String() string {
	return lib.ObjectToString(e)
}
