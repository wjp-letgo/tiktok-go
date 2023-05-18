package oauth

import (
	"fmt"
	"github.com/wjpxxx/letgo/lib"
	tiktokConfig "github.com/wjpxxx/tiktokgo/config"
	oauthentity "github.com/wjpxxx/tiktokgo/oauth/entity"
)

//OAuth
type OAuth struct {
	Config *tiktokConfig.Config
}

//AuthorizationURL
func (a *OAuth) AuthorizationURL(state string) string {
	return fmt.Sprintf("%s/oauth/authorize?app_key=%s&state=%s", a.Config.BaseURL, a.Config.AppKey, state)
}

//GetAccessToken 获得token
func (a *OAuth) GetAccessToken(code string) *oauthentity.GetAccessTokenResult {
	var result oauthentity.GetAccessTokenResult
	params := lib.InRow{
		"auth_code":  code,
		"grant_type": "authorized_code",
	}
	a.Config.HttpS("/api/v2/token/get", params, &result)
	return &result
}

//RefreshToken 刷新token
func (a *OAuth) RefreshToken(refreshToken string) *oauthentity.GetAccessTokenResult {
	var result oauthentity.GetAccessTokenResult
	params := lib.InRow{
		"refresh_token": refreshToken,
		"grant_type":    "refresh_token",
	}
	err := a.Config.HttpS("/api/v2/token/refresh", params, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}
