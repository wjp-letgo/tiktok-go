package authorization

import (
	"github.com/wjp-letgo/letgo/lib"
	tiktokConfig "github.com/wjp-letgo/tiktok-go/config"
	authorizationentity "github.com/wjp-letgo/tiktok-go/authorization/entity"
)

//Authorization
type Authorization struct {
	Config *tiktokConfig.Config
}
//应用程序访问商店数据之前需要获得卖家授权。使用此API可以检查哪些商店当前被授权使用某个应用程序，并获得相应的商店密码，用作商店相关API中的输入参数。
func (a *Authorization) Shops()*authorizationentity.ShopsResult{
	var result authorizationentity.ShopsResult
	params := lib.InRow{
	}
	err := a.Config.GetS2("/authorization/202309/shops", params, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}
//检索合作伙伴为应用程序授权的业务类别资产列表。
//应用程序访问合作伙伴的数据之前需要合作伙伴授权，并且此访问权限是根据业务类别授予的。使用此API可检查当前为应用程序授权的业务类别资产，并获取相应的类别资产密码，用作关联合作伙伴相关API中的输入参数。
//有关合作伙伴授权的更多信息，请参阅《合作伙伴授权指南》。
func (a *Authorization)CategoryAssets()*authorizationentity.CategoryAssetsResult{
	var result authorizationentity.CategoryAssetsResult
	params := lib.InRow{
	}
	err := a.Config.GetS2("/authorization/202405/category_assets", params, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}