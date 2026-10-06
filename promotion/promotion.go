package promotion

import (
	"fmt"

	"github.com/wjp-letgo/letgo/lib"
	tiktokConfig "github.com/wjp-letgo/tiktok-go/config"
	promotionentity "github.com/wjp-letgo/tiktok-go/promotion/entity"
)

// Promotion
type Promotion struct {
	Config *tiktokConfig.Config
}

func (a *Promotion) CreateActivity(body *promotionentity.CreateActivityRequest) *promotionentity.CreateActivityResult {
	var result promotionentity.CreateActivityResult
	err := a.Config.HttpS3("/promotion/202309/activities", nil, body, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}

func (a *Promotion) UpdateActivity(activityId string, body *promotionentity.UpdateActivityRequest) *promotionentity.UpdateActivityResult {
	var result promotionentity.UpdateActivityResult
	err := a.Config.PutS3(fmt.Sprintf("/promotion/202309/activities/%s", activityId), nil, body, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}

func (a *Promotion) DeactivateActivity(activityId string) *promotionentity.DeactivateActivityResult {
	var result promotionentity.DeactivateActivityResult
	err := a.Config.HttpS3(fmt.Sprintf("/promotion/202309/activities/%s/deactivate", activityId), nil, nil, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}

func (a *Promotion) GetActivity(activityId string) *promotionentity.ActivityDetailResult {
	var result promotionentity.ActivityDetailResult
	params := lib.InRow{}
	err := a.Config.GetS3(fmt.Sprintf("/promotion/202309/activities/%s", activityId), params, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}

func (a *Promotion) SearchActivities(pageSize int, pageToken string, body *promotionentity.SearchActivitiesRequest) *promotionentity.SearchActivitiesResult {
	var result promotionentity.SearchActivitiesResult
	params := lib.InRow{
		"page_size": pageSize,
	}
	if pageToken != "" {
		params["page_token"] = pageToken
	}
	err := a.Config.HttpS3("/promotion/202309/activities/search", params, body, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}

func (a *Promotion) UpdateActivityProduct(activityId string, body *promotionentity.UpdateActivityProductRequest) *promotionentity.UpdateActivityProductResult {
	var result promotionentity.UpdateActivityProductResult
	err := a.Config.PutS3(fmt.Sprintf("/promotion/202309/activities/%s/products", activityId), nil, body, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}

func (a *Promotion) RemoveActivityProduct(activityId string, body *promotionentity.RemoveActivityProductRequest) *promotionentity.RemoveActivityProductResult {
	var result promotionentity.RemoveActivityProductResult
	err := a.Config.DeleteS3(fmt.Sprintf("/promotion/202309/activities/%s/products", activityId), nil, body, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}

func (a *Promotion) SearchCoupons(pageSize int, pageToken string, body *promotionentity.SearchCouponsRequest) *promotionentity.SearchCouponsResult {
	var result promotionentity.SearchCouponsResult
	params := lib.InRow{
		"page_size": pageSize,
	}
	if pageToken != "" {
		params["page_token"] = pageToken
	}
	err := a.Config.HttpS3("/promotion/202309/coupons/search", params, body, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}

func (a *Promotion) GetCoupon(couponId string) *promotionentity.CouponDetailResult {
	var result promotionentity.CouponDetailResult
	params := lib.InRow{}
	err := a.Config.GetS3(fmt.Sprintf("/promotion/202309/coupons/%s", couponId), params, &result)
	if err != nil {
		result.Code = -1
		result.Message = err.Error()
	}
	return &result
}
