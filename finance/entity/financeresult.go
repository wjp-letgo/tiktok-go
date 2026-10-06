package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

// @json
type StatementsResult struct {
	Code      int            `json:"code"`
	Message   string         `json:"message"`
	Data      StatementsData `json:"data"`
	RequestId string         `json:"request_id"`
}

func (e *StatementsResult) String() string {
	return lib.ObjectToString(e)
}

// @json
type StatementsData struct {
	NextPageToken string     `json:"next_page_token"`
	Statements    Statements `json:"statements"`
}

func (e *StatementsData) String() string {
	return lib.ObjectToString(e)
}

type Statements []Statement

func (e *Statements) String() string {
	return lib.ObjectToString(e)
}

// @json
type Statement struct {
	Id                string `json:"id"`
	StatementTime     int    `json:"statement_time"`
	SettlementAmount  string `json:"settlement_amount"`
	Currency          string `json:"currency"`
	RevenueAmount     string `json:"revenue_amount"`
	FeeAmount         string `json:"fee_amount"`
	AdjustmentAmount  string `json:"adjustment_amount"`
	PaymentStatus     string `json:"payment_status"`
	PaymentId         string `json:"payment_id"`
	NetSalesAmount    string `json:"net_sales_amount"`
	ShippingCostAmount string `json:"shipping_cost_amount"`
}

func (e *Statement) String() string {
	return lib.ObjectToString(e)
}

// @json
type StatementTransactionsResult struct {
	Code      int                       `json:"code"`
	Message   string                    `json:"message"`
	Data      StatementTransactionsData `json:"data"`
	RequestId string                    `json:"request_id"`
}

func (e *StatementTransactionsResult) String() string {
	return lib.ObjectToString(e)
}

// @json
type StatementTransactionsData struct {
	NextPageToken         string                 `json:"next_page_token"`
	Id                    string                 `json:"id"`
	StatementTransactions StatementTransactions  `json:"statement_transactions"`
}

func (e *StatementTransactionsData) String() string {
	return lib.ObjectToString(e)
}

type StatementTransactions []StatementTransaction

func (e *StatementTransactions) String() string {
	return lib.ObjectToString(e)
}

// @json
type StatementTransaction struct {
	Id               string `json:"id"`
	Type             string `json:"type"`
	OrderId          string `json:"order_id"`
	OrderCreateTime  int    `json:"order_create_time"`
	Currency         string `json:"currency"`
	SettlementAmount string `json:"settlement_amount"`
	RevenueAmount    string `json:"revenue_amount"`
	FeeAmount        string `json:"fee_amount"`
	AdjustmentAmount string `json:"adjustment_amount"`
}

func (e *StatementTransaction) String() string {
	return lib.ObjectToString(e)
}

// @json
type OrderStatementTransactionsResult struct {
	Code      int                            `json:"code"`
	Message   string                         `json:"message"`
	Data      OrderStatementTransactionsData `json:"data"`
	RequestId string                         `json:"request_id"`
}

func (e *OrderStatementTransactionsResult) String() string {
	return lib.ObjectToString(e)
}

// @json
type OrderStatementTransactionsData struct {
	OrderId               string                `json:"order_id"`
	SkuStatements         SkuStatements         `json:"sku_statements"`
	StatementTransactions StatementTransactions `json:"statement_transactions"`
}

func (e *OrderStatementTransactionsData) String() string {
	return lib.ObjectToString(e)
}

type SkuStatements []SkuStatement

func (e *SkuStatements) String() string {
	return lib.ObjectToString(e)
}

// @json
type SkuStatement struct {
	SkuId            string `json:"sku_id"`
	SkuName          string `json:"sku_name"`
	Quantity         int    `json:"quantity"`
	SettlementAmount string `json:"settlement_amount"`
	Currency         string `json:"currency"`
}

func (e *SkuStatement) String() string {
	return lib.ObjectToString(e)
}

// @json
type PaymentsResult struct {
	Code      int          `json:"code"`
	Message   string       `json:"message"`
	Data      PaymentsData `json:"data"`
	RequestId string       `json:"request_id"`
}

func (e *PaymentsResult) String() string {
	return lib.ObjectToString(e)
}

// @json
type PaymentsData struct {
	NextPageToken string   `json:"next_page_token"`
	Payments      Payments `json:"payments"`
}

func (e *PaymentsData) String() string {
	return lib.ObjectToString(e)
}

type Payments []Payment

func (e *Payments) String() string {
	return lib.ObjectToString(e)
}

// @json
type Payment struct {
	Id           string `json:"id"`
	CreateTime   int    `json:"create_time"`
	Amount       string `json:"amount"`
	Currency     string `json:"currency"`
	Status       string `json:"status"`
	BankAccount  string `json:"bank_account"`
	ExchangeRate string `json:"exchange_rate"`
}

func (e *Payment) String() string {
	return lib.ObjectToString(e)
}

// @json
type WithdrawalsResult struct {
	Code      int             `json:"code"`
	Message   string          `json:"message"`
	Data      WithdrawalsData `json:"data"`
	RequestId string          `json:"request_id"`
}

func (e *WithdrawalsResult) String() string {
	return lib.ObjectToString(e)
}

// @json
type WithdrawalsData struct {
	NextPageToken string      `json:"next_page_token"`
	Withdrawals   Withdrawals `json:"withdrawals"`
}

func (e *WithdrawalsData) String() string {
	return lib.ObjectToString(e)
}

type Withdrawals []Withdrawal

func (e *Withdrawals) String() string {
	return lib.ObjectToString(e)
}

// @json
type Withdrawal struct {
	Id         string `json:"id"`
	Type       string `json:"type"`
	Amount     string `json:"amount"`
	Currency   string `json:"currency"`
	Status     string `json:"status"`
	CreateTime int    `json:"create_time"`
}

func (e *Withdrawal) String() string {
	return lib.ObjectToString(e)
}
