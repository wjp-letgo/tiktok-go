package entity

import "github.com/wjpxxx/letgo/lib"

type GetOrderListResult struct {
	Code      int                      `json:"code"`
	Message   string                   `json:"message"`
	RequestId string                   `json:"request_id"`
	Data      []GetOrderListResultData `json:"data"`
}

type GetOrderListResultData struct {
	More       bool        `json:"more"`
	NextCursor string      `json:"next_cursor"`
	Total      int         `json:"total"`
	OrderList  []OrderList `json:"order_list"`
}

type OrderList struct {
	OrderId     string `json:"order_id"`
	OrderStatus int    `json:"order_status"`
	UpdateTime  int    `json:"update_time"`
}

func (getorderlistresult *GetOrderListResult) String() string {
	return lib.ObjectToString(getorderlistresult)
}

func (getorderlistresultdata *GetOrderListResultData) String() string {
	return lib.ObjectToString(getorderlistresultdata)
}

func (getorderlist *OrderList) String() string {
	return lib.ObjectToString(getorderlist)
}
