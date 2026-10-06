package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

// @json
type CreateActivityRequest struct {
	Title        string `json:"title,omitempty"`
	ActivityType string `json:"activity_type,omitempty"`
	BeginTime    int    `json:"begin_time,omitempty"`
	EndTime      int    `json:"end_time,omitempty"`
	ProductLevel string `json:"product_level,omitempty"`
}

func (e *CreateActivityRequest) String() string {
	return lib.ObjectToString(e)
}

// @json
type CreateActivityResult struct {
	Code      int                `json:"code"`
	Message   string             `json:"message"`
	Data      CreateActivityData `json:"data"`
	RequestId string             `json:"request_id"`
}

func (e *CreateActivityResult) String() string {
	return lib.ObjectToString(e)
}

// @json
type CreateActivityData struct {
	ActivityId string `json:"activity_id"`
}

func (e *CreateActivityData) String() string {
	return lib.ObjectToString(e)
}

// @json
type UpdateActivityRequest struct {
	Title     string `json:"title,omitempty"`
	BeginTime int    `json:"begin_time,omitempty"`
	EndTime   int    `json:"end_time,omitempty"`
}

func (e *UpdateActivityRequest) String() string {
	return lib.ObjectToString(e)
}

// @json
type UpdateActivityResult struct {
	Code      int       `json:"code"`
	Message   string    `json:"message"`
	Data      EmptyData `json:"data"`
	RequestId string    `json:"request_id"`
}

func (e *UpdateActivityResult) String() string {
	return lib.ObjectToString(e)
}

// @json
type DeactivateActivityResult struct {
	Code      int       `json:"code"`
	Message   string    `json:"message"`
	Data      EmptyData `json:"data"`
	RequestId string    `json:"request_id"`
}

func (e *DeactivateActivityResult) String() string {
	return lib.ObjectToString(e)
}

// @json
type EmptyData struct {
}

func (e *EmptyData) String() string {
	return lib.ObjectToString(e)
}

// @json
type ActivityDetailResult struct {
	Code      int          `json:"code"`
	Message   string       `json:"message"`
	Data      ActivityData `json:"data"`
	RequestId string       `json:"request_id"`
}

func (e *ActivityDetailResult) String() string {
	return lib.ObjectToString(e)
}

// @json
type ActivityData struct {
	Id           string           `json:"id"`
	Title        string           `json:"title"`
	ActivityType string           `json:"activity_type"`
	Status       string           `json:"status"`
	BeginTime    int              `json:"begin_time"`
	EndTime      int              `json:"end_time"`
	ProductLevel string           `json:"product_level"`
	CreateTime   int              `json:"create_time"`
	UpdateTime   int              `json:"update_time"`
	Products     ActivityProducts `json:"products"`
}

func (e *ActivityData) String() string {
	return lib.ObjectToString(e)
}

type ActivityProducts []ActivityProduct

func (e *ActivityProducts) String() string {
	return lib.ObjectToString(e)
}

// @json
type ActivityProduct struct {
	Id    string       `json:"id,omitempty"`
	Skus  ActivitySkus `json:"skus,omitempty"`
	Discount string    `json:"discount,omitempty"`
	QuantityLimit int  `json:"quantity_limit,omitempty"`
}

func (e *ActivityProduct) String() string {
	return lib.ObjectToString(e)
}

type ActivitySkus []ActivitySku

func (e *ActivitySkus) String() string {
	return lib.ObjectToString(e)
}

// @json
type ActivitySku struct {
	Id            string `json:"id,omitempty"`
	Discount      string `json:"discount,omitempty"`
	ActivityPrice string `json:"activity_price,omitempty"`
	QuantityLimit int    `json:"quantity_limit,omitempty"`
}

func (e *ActivitySku) String() string {
	return lib.ObjectToString(e)
}

// @json
type SearchActivitiesRequest struct {
	Status       string `json:"status,omitempty"`
	ActivityType string `json:"activity_type,omitempty"`
	Title        string `json:"title,omitempty"`
}

func (e *SearchActivitiesRequest) String() string {
	return lib.ObjectToString(e)
}

// @json
type SearchActivitiesResult struct {
	Code      int                  `json:"code"`
	Message   string               `json:"message"`
	Data      SearchActivitiesData `json:"data"`
	RequestId string               `json:"request_id"`
}

func (e *SearchActivitiesResult) String() string {
	return lib.ObjectToString(e)
}

// @json
type SearchActivitiesData struct {
	NextPageToken string           `json:"next_page_token"`
	TotalCount    int              `json:"total_count"`
	Activities    ActivityList     `json:"activities"`
}

func (e *SearchActivitiesData) String() string {
	return lib.ObjectToString(e)
}

type ActivityList []ActivityData

func (e *ActivityList) String() string {
	return lib.ObjectToString(e)
}

// @json
type UpdateActivityProductRequest struct {
	ActivityId string           `json:"activity_id,omitempty"`
	Products   ActivityProducts `json:"products,omitempty"`
}

func (e *UpdateActivityProductRequest) String() string {
	return lib.ObjectToString(e)
}

// @json
type UpdateActivityProductResult struct {
	Code      int       `json:"code"`
	Message   string    `json:"message"`
	Data      EmptyData `json:"data"`
	RequestId string    `json:"request_id"`
}

func (e *UpdateActivityProductResult) String() string {
	return lib.ObjectToString(e)
}

// @json
type RemoveActivityProductRequest struct {
	ProductIds []string `json:"product_ids,omitempty"`
	SkuIds     []string `json:"sku_ids,omitempty"`
}

func (e *RemoveActivityProductRequest) String() string {
	return lib.ObjectToString(e)
}

// @json
type RemoveActivityProductResult struct {
	Code      int       `json:"code"`
	Message   string    `json:"message"`
	Data      EmptyData `json:"data"`
	RequestId string    `json:"request_id"`
}

func (e *RemoveActivityProductResult) String() string {
	return lib.ObjectToString(e)
}

// @json
type SearchCouponsRequest struct {
	Status     string `json:"status,omitempty"`
	CouponType string `json:"coupon_type,omitempty"`
	Title      string `json:"title,omitempty"`
}

func (e *SearchCouponsRequest) String() string {
	return lib.ObjectToString(e)
}

// @json
type SearchCouponsResult struct {
	Code      int               `json:"code"`
	Message   string            `json:"message"`
	Data      SearchCouponsData `json:"data"`
	RequestId string            `json:"request_id"`
}

func (e *SearchCouponsResult) String() string {
	return lib.ObjectToString(e)
}

// @json
type SearchCouponsData struct {
	NextPageToken string  `json:"next_page_token"`
	TotalCount    int     `json:"total_count"`
	Coupons       Coupons `json:"coupons"`
}

func (e *SearchCouponsData) String() string {
	return lib.ObjectToString(e)
}

type Coupons []Coupon

func (e *Coupons) String() string {
	return lib.ObjectToString(e)
}

// @json
type Coupon struct {
	Id         string `json:"id"`
	Title      string `json:"title"`
	Status     string `json:"status"`
	CouponType string `json:"coupon_type"`
	BeginTime  int    `json:"begin_time"`
	EndTime    int    `json:"end_time"`
	CreateTime int    `json:"create_time"`
}

func (e *Coupon) String() string {
	return lib.ObjectToString(e)
}

// @json
type CouponDetailResult struct {
	Code      int        `json:"code"`
	Message   string     `json:"message"`
	Data      Coupon     `json:"data"`
	RequestId string     `json:"request_id"`
}

func (e *CouponDetailResult) String() string {
	return lib.ObjectToString(e)
}
