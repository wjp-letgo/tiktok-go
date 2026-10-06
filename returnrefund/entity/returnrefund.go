package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

// @json
type EmptyData struct {
}

func (e *EmptyData) String() string {
	return lib.ObjectToString(e)
}

// @json
type AftersaleEligibilityResult struct {
	Code      int                      `json:"code"`
	Message   string                   `json:"message"`
	Data      AftersaleEligibilityData `json:"data"`
	RequestId string                   `json:"request_id"`
}

func (e *AftersaleEligibilityResult) String() string {
	return lib.ObjectToString(e)
}

// @json
type AftersaleEligibilityData struct {
	SkuEligibility SkuEligibilities `json:"sku_eligibility"`
}

func (e *AftersaleEligibilityData) String() string {
	return lib.ObjectToString(e)
}

type SkuEligibilities []SkuEligibility

func (e *SkuEligibilities) String() string {
	return lib.ObjectToString(e)
}

// @json
type SkuEligibility struct {
	SkuId            string   `json:"sku_id"`
	LineItemId       string   `json:"line_item_id"`
	Returnable       bool     `json:"returnable"`
	Refundable       bool     `json:"refundable"`
	Cancelable       bool     `json:"cancelable"`
	IneligibleReasons []string `json:"ineligible_reasons"`
}

func (e *SkuEligibility) String() string {
	return lib.ObjectToString(e)
}

// @json
type RejectReasonsResult struct {
	Code      int               `json:"code"`
	Message   string            `json:"message"`
	Data      RejectReasonsData `json:"data"`
	RequestId string            `json:"request_id"`
}

func (e *RejectReasonsResult) String() string {
	return lib.ObjectToString(e)
}

// @json
type RejectReasonsData struct {
	Reasons Reasons `json:"reasons"`
}

func (e *RejectReasonsData) String() string {
	return lib.ObjectToString(e)
}

type Reasons []Reason

func (e *Reasons) String() string {
	return lib.ObjectToString(e)
}

// @json
type Reason struct {
	Name        string `json:"name"`
	Text        string `json:"text"`
}

func (e *Reason) String() string {
	return lib.ObjectToString(e)
}

// @json
type CreateReturnRequest struct {
	OrderId          string   `json:"order_id,omitempty"`
	OrderLineItemIds []string `json:"order_line_item_ids,omitempty"`
	ReturnReason     string   `json:"return_reason,omitempty"`
	ReturnType       string   `json:"return_type,omitempty"`
	RefundTotal      string   `json:"refund_total,omitempty"`
	Currency         string   `json:"currency,omitempty"`
}

func (e *CreateReturnRequest) String() string {
	return lib.ObjectToString(e)
}

// @json
type CreateReturnResult struct {
	Code      int              `json:"code"`
	Message   string           `json:"message"`
	Data      CreateReturnData `json:"data"`
	RequestId string           `json:"request_id"`
}

func (e *CreateReturnResult) String() string {
	return lib.ObjectToString(e)
}

// @json
type CreateReturnData struct {
	ReturnId string `json:"return_id"`
}

func (e *CreateReturnData) String() string {
	return lib.ObjectToString(e)
}

// @json
type ApproveReturnRequest struct {
	Decision string `json:"decision,omitempty"`
	BuyerKeepItem bool `json:"buyer_keep_item,omitempty"`
}

func (e *ApproveReturnRequest) String() string {
	return lib.ObjectToString(e)
}

// @json
type ApproveReturnResult struct {
	Code      int       `json:"code"`
	Message   string    `json:"message"`
	Data      EmptyData `json:"data"`
	RequestId string    `json:"request_id"`
}

func (e *ApproveReturnResult) String() string {
	return lib.ObjectToString(e)
}

// @json
type RejectReturnRequest struct {
	RejectReason string `json:"reject_reason,omitempty"`
	Comment      string `json:"comment,omitempty"`
}

func (e *RejectReturnRequest) String() string {
	return lib.ObjectToString(e)
}

// @json
type RejectReturnResult struct {
	Code      int       `json:"code"`
	Message   string    `json:"message"`
	Data      EmptyData `json:"data"`
	RequestId string    `json:"request_id"`
}

func (e *RejectReturnResult) String() string {
	return lib.ObjectToString(e)
}

// @json
type SearchReturnsRequest struct {
	ReturnIds    []string `json:"return_ids,omitempty"`
	OrderIds     []string `json:"order_ids,omitempty"`
	BuyerUserIds []string `json:"buyer_user_ids,omitempty"`
	ReturnTypes  []string `json:"return_types,omitempty"`
	ReturnStatus []string `json:"return_status,omitempty"`
	CreateTimeGe int      `json:"create_time_ge,omitempty"`
	CreateTimeLt int      `json:"create_time_lt,omitempty"`
	UpdateTimeGe int      `json:"update_time_ge,omitempty"`
	UpdateTimeLt int      `json:"update_time_lt,omitempty"`
}

func (e *SearchReturnsRequest) String() string {
	return lib.ObjectToString(e)
}

// @json
type SearchReturnsResult struct {
	Code      int               `json:"code"`
	Message   string            `json:"message"`
	Data      SearchReturnsData `json:"data"`
	RequestId string            `json:"request_id"`
}

func (e *SearchReturnsResult) String() string {
	return lib.ObjectToString(e)
}

// @json
type SearchReturnsData struct {
	NextPageToken string  `json:"next_page_token"`
	TotalCount    int     `json:"total_count"`
	Returns       Returns `json:"return_orders"`
}

func (e *SearchReturnsData) String() string {
	return lib.ObjectToString(e)
}

type Returns []ReturnOrder

func (e *Returns) String() string {
	return lib.ObjectToString(e)
}

// @json
type ReturnOrder struct {
	Id             string   `json:"id"`
	OrderId        string   `json:"order_id"`
	ReturnType     string   `json:"return_type"`
	ReturnStatus   string   `json:"return_status"`
	ArbitrationStatus string `json:"arbitration_status"`
	Role           string   `json:"role"`
	ReturnReason   string   `json:"return_reason"`
	CreateTime     int      `json:"create_time"`
	UpdateTime     int      `json:"update_time"`
	RefundAmount   string   `json:"refund_amount"`
	Currency       string   `json:"currency"`
	SellerNextActionResponse []string `json:"seller_next_action_response"`
}

func (e *ReturnOrder) String() string {
	return lib.ObjectToString(e)
}

// @json
type ReturnRecordsResult struct {
	Code      int               `json:"code"`
	Message   string            `json:"message"`
	Data      ReturnRecordsData `json:"data"`
	RequestId string            `json:"request_id"`
}

func (e *ReturnRecordsResult) String() string {
	return lib.ObjectToString(e)
}

// @json
type ReturnRecordsData struct {
	Records Records `json:"records"`
}

func (e *ReturnRecordsData) String() string {
	return lib.ObjectToString(e)
}

type Records []Record

func (e *Records) String() string {
	return lib.ObjectToString(e)
}

// @json
type Record struct {
	Event     string `json:"event"`
	Note      string `json:"note"`
	CreateTime int   `json:"create_time"`
	Role      string `json:"role"`
}

func (e *Record) String() string {
	return lib.ObjectToString(e)
}

// @json
type CancelOrderRequest struct {
	OrderId          string   `json:"order_id,omitempty"`
	CancelReason     string   `json:"cancel_reason,omitempty"`
	OrderLineItemIds []string `json:"order_line_item_ids,omitempty"`
}

func (e *CancelOrderRequest) String() string {
	return lib.ObjectToString(e)
}

// @json
type CancelOrderResult struct {
	Code      int             `json:"code"`
	Message   string          `json:"message"`
	Data      CancelOrderData `json:"data"`
	RequestId string          `json:"request_id"`
}

func (e *CancelOrderResult) String() string {
	return lib.ObjectToString(e)
}

// @json
type CancelOrderData struct {
	CancelId string `json:"cancel_id"`
}

func (e *CancelOrderData) String() string {
	return lib.ObjectToString(e)
}

// @json
type ApproveCancellationResult struct {
	Code      int       `json:"code"`
	Message   string    `json:"message"`
	Data      EmptyData `json:"data"`
	RequestId string    `json:"request_id"`
}

func (e *ApproveCancellationResult) String() string {
	return lib.ObjectToString(e)
}

// @json
type RejectCancellationRequest struct {
	RejectReason string `json:"reject_reason,omitempty"`
	Comment      string `json:"comment,omitempty"`
}

func (e *RejectCancellationRequest) String() string {
	return lib.ObjectToString(e)
}

// @json
type RejectCancellationResult struct {
	Code      int       `json:"code"`
	Message   string    `json:"message"`
	Data      EmptyData `json:"data"`
	RequestId string    `json:"request_id"`
}

func (e *RejectCancellationResult) String() string {
	return lib.ObjectToString(e)
}

// @json
type SearchCancellationsRequest struct {
	CancelIds     []string `json:"cancel_ids,omitempty"`
	OrderIds      []string `json:"order_ids,omitempty"`
	BuyerUserIds  []string `json:"buyer_user_ids,omitempty"`
	CancelStatus  []string `json:"cancel_status,omitempty"`
	CancelTypes   []string `json:"cancel_types,omitempty"`
	CreateTimeGe  int      `json:"create_time_ge,omitempty"`
	CreateTimeLt  int      `json:"create_time_lt,omitempty"`
	UpdateTimeGe  int      `json:"update_time_ge,omitempty"`
	UpdateTimeLt  int      `json:"update_time_lt,omitempty"`
}

func (e *SearchCancellationsRequest) String() string {
	return lib.ObjectToString(e)
}

// @json
type SearchCancellationsResult struct {
	Code      int                     `json:"code"`
	Message   string                  `json:"message"`
	Data      SearchCancellationsData `json:"data"`
	RequestId string                  `json:"request_id"`
}

func (e *SearchCancellationsResult) String() string {
	return lib.ObjectToString(e)
}

// @json
type SearchCancellationsData struct {
	NextPageToken string        `json:"next_page_token"`
	TotalCount    int           `json:"total_count"`
	Cancellations Cancellations `json:"cancellations"`
}

func (e *SearchCancellationsData) String() string {
	return lib.ObjectToString(e)
}

type Cancellations []Cancellation

func (e *Cancellations) String() string {
	return lib.ObjectToString(e)
}

// @json
type Cancellation struct {
	CancelId     string `json:"cancel_id"`
	OrderId      string `json:"order_id"`
	CancelType   string `json:"cancel_type"`
	CancelStatus string `json:"cancel_status"`
	CancelReason string `json:"cancel_reason"`
	CreateTime   int    `json:"create_time"`
	UpdateTime   int    `json:"update_time"`
}

func (e *Cancellation) String() string {
	return lib.ObjectToString(e)
}

// @json
type CalculateRefundRequest struct {
	OrderId          string   `json:"order_id,omitempty"`
	OrderLineItemIds []string `json:"order_line_item_ids,omitempty"`
}

func (e *CalculateRefundRequest) String() string {
	return lib.ObjectToString(e)
}

// @json
type CalculateRefundResult struct {
	Code      int                 `json:"code"`
	Message   string              `json:"message"`
	Data      CalculateRefundData `json:"data"`
	RequestId string              `json:"request_id"`
}

func (e *CalculateRefundResult) String() string {
	return lib.ObjectToString(e)
}

// @json
type CalculateRefundData struct {
	RefundTotal string `json:"refund_total"`
	Currency    string `json:"currency"`
}

func (e *CalculateRefundData) String() string {
	return lib.ObjectToString(e)
}
