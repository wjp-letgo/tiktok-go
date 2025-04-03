package logistic

import (
	"fmt"

	"github.com/wjp-letgo/letgo/lib"
	tiktokConfig "github.com/wjp-letgo/tiktok-go/config"
	logisticentity "github.com/wjp-letgo/tiktok-go/logistic/entity"
)

//Logistic
type Logistic struct {
	Config *tiktokConfig.Config
}

//获得仓库列表
func (a *Logistic)WarehouseList()*logisticentity.WarehouseListResult{
	var result logisticentity.WarehouseListResult
	params := lib.InRow{
	}
	err := a.Config.GetS3("/logistics/202309/warehouses",params, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}
//此API检索与卖方相关联的所有全局仓库信息。仓库信息包括全局仓库ID、仓库名称和仓库所有权。
func (a *Logistic)GlobalSellerWarehouse()*logisticentity.GlobalSellerWarehouseResult{
	var result logisticentity.GlobalSellerWarehouseResult
	params := lib.InRow{
	}
	err := a.Config.GetS2("/logistics/202309/global_warehouses",params, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}

//此API用于获取通过卖方指定仓库提供的交付选项列表。
func (a *Logistic)WarehouseDeliveryOptions(warehouseId string)*logisticentity.WarehouseDeliveryOptionsResult{
	var result logisticentity.WarehouseDeliveryOptionsResult
	params := lib.InRow{
	}
	err := a.Config.GetS3(fmt.Sprintf("/logistics/202309/warehouses/%s/delivery_options",warehouseId),params, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}

//此API用于获取与指定交付选项相对应的发货提供商
func (a *Logistic)ShippingProviders(deliveryOptionId string)*logisticentity.ShippingProvidersResult{
	var result logisticentity.ShippingProvidersResult
	params := lib.InRow{
	}
	err := a.Config.GetS3(fmt.Sprintf("/logistics/202309/delivery_options/%s/shipping_providers",deliveryOptionId),params, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}