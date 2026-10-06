package supplychain

import (
	tiktokConfig "github.com/wjp-letgo/tiktok-go/config"
	supplychainentity "github.com/wjp-letgo/tiktok-go/supplychain/entity"
)

// SupplyChain
type SupplyChain struct {
	Config *tiktokConfig.Config
}

func (a *SupplyChain) ConfirmPackageShipment(body *supplychainentity.ConfirmPackageShipmentRequest) *supplychainentity.ConfirmPackageShipmentResult {
	var result supplychainentity.ConfirmPackageShipmentResult
	err := a.Config.HttpS3("/supply_chain/202309/packages/sync", nil, body, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}
