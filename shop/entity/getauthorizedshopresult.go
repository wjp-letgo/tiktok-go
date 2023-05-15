package entity

import (
	"github.com/wjpxxx/letgo/lib"
)

//@json
type GetAuthorizedShopResult struct {
	Code      int                   `json:"code"`
	Message   string                `json:"message"`
	Data      GetAuthorizedShopData `json:"data"`
	RequestId string                `json:"request_id"`
}

//@json
type GetAuthorizedShopData struct {
	ShopList []ShopList `json:"shop_list"`
}

//@json
type ShopList struct {
	Region   string `json:"region"`
	ShopId   string `json:"shop_id"`
	ShopName string `json:"shop_name"`
	Type     int    `json:"type"`
}

func (getauthorizedshopresult *GetAuthorizedShopResult) String() string {
	return lib.ObjectToString(getauthorizedshopresult)
}
func (getauthorizedshopdata *GetAuthorizedShopData) String() string {
	return lib.ObjectToString(getauthorizedshopdata)
}
func (shoplist *ShopList) String() string {
	return lib.ObjectToString(shoplist)
}
