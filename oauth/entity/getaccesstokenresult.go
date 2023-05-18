package entity

import (
	"github.com/wjpxxx/letgo/lib"
)

//@json
type GetAccessTokenResult struct {
	Code    int                `json:"code"`
	Message string             `json:"message"`
	Data    GetAccessTokenData `json:"data"`
	//RequestId string `json:"request_id"`
}

//@json
type GetAccessTokenData struct {
	AccessToken          string `json:"access_token"`
	AccessTokenExpireIn  int    `json:"access_token_expire_in"`
	RefreshToken         string `json:"refresh_token"`
	RefreshTokenExpireIn int    `json:"refresh_token_expire_in"`
	OpenId               string `json:"open_id"`
	SellerName           string `json:"seller_name"`
	//SellerBaseRegion     string `json:"seller_base_region"`
	RequestId string `json:"request_id"`
}

func (getaccesstokenresult *GetAccessTokenResult) String() string {
	return lib.ObjectToString(getaccesstokenresult)
}
func (getaccesstokendata *GetAccessTokenData) String() string {
	return lib.ObjectToString(getaccesstokendata)
}
