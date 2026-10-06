package returnrefund

import (
	"fmt"

	"github.com/wjp-letgo/letgo/lib"
	tiktokConfig "github.com/wjp-letgo/tiktok-go/config"
	returnrefundentity "github.com/wjp-letgo/tiktok-go/returnrefund/entity"
)

// ReturnRefund
type ReturnRefund struct {
	Config *tiktokConfig.Config
}

func (a *ReturnRefund) AftersaleEligibility(orderId string) *returnrefundentity.AftersaleEligibilityResult {
	var result returnrefundentity.AftersaleEligibilityResult
	params := lib.InRow{}
	err := a.Config.GetS3(fmt.Sprintf("/return_refund/202309/orders/%s/aftersale_eligibility", orderId), params, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}

func (a *ReturnRefund) RejectReasons(returnOrCancelId, locale string) *returnrefundentity.RejectReasonsResult {
	var result returnrefundentity.RejectReasonsResult
	params := lib.InRow{}
	if returnOrCancelId != "" {
		params["return_or_cancel_id"] = returnOrCancelId
	}
	if locale != "" {
		params["locale"] = locale
	}
	err := a.Config.GetS3("/return_refund/202309/reject_reasons", params, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}

func (a *ReturnRefund) CreateReturn(idempotencyKey string, body *returnrefundentity.CreateReturnRequest) *returnrefundentity.CreateReturnResult {
	var result returnrefundentity.CreateReturnResult
	params := lib.InRow{}
	if idempotencyKey != "" {
		params["idempotency_key"] = idempotencyKey
	}
	err := a.Config.HttpS3("/return_refund/202309/returns", params, body, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}

func (a *ReturnRefund) ApproveReturn(returnId, idempotencyKey string, body *returnrefundentity.ApproveReturnRequest) *returnrefundentity.ApproveReturnResult {
	var result returnrefundentity.ApproveReturnResult
	params := lib.InRow{}
	if idempotencyKey != "" {
		params["idempotency_key"] = idempotencyKey
	}
	err := a.Config.HttpS3(fmt.Sprintf("/return_refund/202309/returns/%s/approve", returnId), params, body, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}

func (a *ReturnRefund) RejectReturn(returnId, idempotencyKey string, body *returnrefundentity.RejectReturnRequest) *returnrefundentity.RejectReturnResult {
	var result returnrefundentity.RejectReturnResult
	params := lib.InRow{}
	if idempotencyKey != "" {
		params["idempotency_key"] = idempotencyKey
	}
	err := a.Config.HttpS3(fmt.Sprintf("/return_refund/202309/returns/%s/reject", returnId), params, body, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}

func (a *ReturnRefund) SearchReturns(pageSize int, pageToken, sortField, sortOrder string, body *returnrefundentity.SearchReturnsRequest) *returnrefundentity.SearchReturnsResult {
	var result returnrefundentity.SearchReturnsResult
	params := lib.InRow{
		"page_size": pageSize,
	}
	if pageToken != "" {
		params["page_token"] = pageToken
	}
	if sortField != "" {
		params["sort_field"] = sortField
	}
	if sortOrder != "" {
		params["sort_order"] = sortOrder
	}
	err := a.Config.HttpS3("/return_refund/202309/returns/search", params, body, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}

func (a *ReturnRefund) ReturnRecords(returnId string) *returnrefundentity.ReturnRecordsResult {
	var result returnrefundentity.ReturnRecordsResult
	params := lib.InRow{}
	err := a.Config.GetS3(fmt.Sprintf("/return_refund/202309/returns/%s/records", returnId), params, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}

func (a *ReturnRefund) CancelOrder(body *returnrefundentity.CancelOrderRequest) *returnrefundentity.CancelOrderResult {
	var result returnrefundentity.CancelOrderResult
	err := a.Config.HttpS3("/return_refund/202309/cancellations", nil, body, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}

func (a *ReturnRefund) ApproveCancellation(cancelId, idempotencyKey string) *returnrefundentity.ApproveCancellationResult {
	var result returnrefundentity.ApproveCancellationResult
	params := lib.InRow{}
	if idempotencyKey != "" {
		params["idempotency_key"] = idempotencyKey
	}
	err := a.Config.HttpS3(fmt.Sprintf("/return_refund/202309/cancellations/%s/approve", cancelId), params, nil, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}

func (a *ReturnRefund) RejectCancellation(cancelId, idempotencyKey string, body *returnrefundentity.RejectCancellationRequest) *returnrefundentity.RejectCancellationResult {
	var result returnrefundentity.RejectCancellationResult
	params := lib.InRow{}
	if idempotencyKey != "" {
		params["idempotency_key"] = idempotencyKey
	}
	err := a.Config.HttpS3(fmt.Sprintf("/return_refund/202309/cancellations/%s/reject", cancelId), params, body, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}

func (a *ReturnRefund) SearchCancellations(pageSize int, pageToken, sortField, sortOrder string, body *returnrefundentity.SearchCancellationsRequest) *returnrefundentity.SearchCancellationsResult {
	var result returnrefundentity.SearchCancellationsResult
	params := lib.InRow{
		"page_size": pageSize,
	}
	if pageToken != "" {
		params["page_token"] = pageToken
	}
	if sortField != "" {
		params["sort_field"] = sortField
	}
	if sortOrder != "" {
		params["sort_order"] = sortOrder
	}
	err := a.Config.HttpS3("/return_refund/202309/cancellations/search", params, body, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}

func (a *ReturnRefund) CalculateRefund(body *returnrefundentity.CalculateRefundRequest) *returnrefundentity.CalculateRefundResult {
	var result returnrefundentity.CalculateRefundResult
	err := a.Config.HttpS3("/return_refund/202309/refunds/calculate", nil, body, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}
