package finance

import (
	"fmt"

	"github.com/wjp-letgo/letgo/lib"
	tiktokConfig "github.com/wjp-letgo/tiktok-go/config"
	financeentity "github.com/wjp-letgo/tiktok-go/finance/entity"
)

// Finance
type Finance struct {
	Config *tiktokConfig.Config
}

// 获取结算单列表
func (a *Finance) GetStatements(pageSize int, pageToken, sortField, sortOrder, statementTimeGe, statementTimeLt, paymentStatus string) *financeentity.StatementsResult {
	var result financeentity.StatementsResult
	params := lib.InRow{
		"page_size": pageSize,
	}
	if pageToken != "" {
		params["page_token"] = pageToken
	}
	if sortField != "" {
		params["sort_field"] = sortField
	} else {
		params["sort_field"] = "statement_time"
	}
	if sortOrder != "" {
		params["sort_order"] = sortOrder
	}
	if statementTimeGe != "" {
		params["statement_time_ge"] = statementTimeGe
	}
	if statementTimeLt != "" {
		params["statement_time_lt"] = statementTimeLt
	}
	if paymentStatus != "" {
		params["payment_status"] = paymentStatus
	}
	err := a.Config.GetS3("/finance/202309/statements", params, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}

// 获取结算单交易明细
func (a *Finance) GetStatementTransactions(statementId string, pageSize int, pageToken, sortField, sortOrder string) *financeentity.StatementTransactionsResult {
	var result financeentity.StatementTransactionsResult
	params := lib.InRow{
		"page_size": pageSize,
	}
	if pageToken != "" {
		params["page_token"] = pageToken
	}
	if sortField != "" {
		params["sort_field"] = sortField
	} else {
		params["sort_field"] = "order_create_time"
	}
	if sortOrder != "" {
		params["sort_order"] = sortOrder
	}
	err := a.Config.GetS3(fmt.Sprintf("/finance/202309/statements/%s/statement_transactions", statementId), params, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}

// 获取订单结算交易
func (a *Finance) GetOrderStatementTransactions(orderId string) *financeentity.OrderStatementTransactionsResult {
	var result financeentity.OrderStatementTransactionsResult
	params := lib.InRow{}
	err := a.Config.GetS3(fmt.Sprintf("/finance/202309/orders/%s/statement_transactions", orderId), params, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}

// 获取付款记录
func (a *Finance) GetPayments(pageSize int, pageToken, sortField, sortOrder, createTimeGe, createTimeLt string) *financeentity.PaymentsResult {
	var result financeentity.PaymentsResult
	params := lib.InRow{
		"page_size": pageSize,
	}
	if pageToken != "" {
		params["page_token"] = pageToken
	}
	if sortField != "" {
		params["sort_field"] = sortField
	} else {
		params["sort_field"] = "create_time"
	}
	if sortOrder != "" {
		params["sort_order"] = sortOrder
	}
	if createTimeGe != "" {
		params["create_time_ge"] = createTimeGe
	}
	if createTimeLt != "" {
		params["create_time_lt"] = createTimeLt
	}
	err := a.Config.GetS3("/finance/202309/payments", params, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}

// 获取提现记录
func (a *Finance) GetWithdrawals(types string, pageSize int, pageToken, createTimeGe, createTimeLt string) *financeentity.WithdrawalsResult {
	var result financeentity.WithdrawalsResult
	params := lib.InRow{
		"page_size": pageSize,
	}
	if types != "" {
		params["types"] = types
	} else {
		params["types"] = "WITHDRAW,SETTLE,TRANSFER,REVERSE"
	}
	if pageToken != "" {
		params["page_token"] = pageToken
	}
	if createTimeGe != "" {
		params["create_time_ge"] = createTimeGe
	}
	if createTimeLt != "" {
		params["create_time_lt"] = createTimeLt
	}
	err := a.Config.GetS3("/finance/202309/withdrawals", params, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}
