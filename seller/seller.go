package seller

import (
	"github.com/wjp-letgo/letgo/lib"
	tiktokConfig "github.com/wjp-letgo/tiktok-go/config"
	sellerentity "github.com/wjp-letgo/tiktok-go/seller/entity"
)

//Seller
type Seller struct {
	Config *tiktokConfig.Config
}
//获得授权店铺列表
func (a *Seller) SellerShops()*sellerentity.ShopsResult{
	var result sellerentity.ShopsResult
	params := lib.InRow{
	}
	err := a.Config.GetS2("/seller/202309/shops", params, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}
//您可以在列出产品之前使用此API来检查卖家是否有能力列出全球产品。
func (a *Seller)Permissions()*sellerentity.PermissionsResult{
	var result sellerentity.PermissionsResult
	params := lib.InRow{
	}
	err := a.Config.GetS2("/seller/202309/permissions", params, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}