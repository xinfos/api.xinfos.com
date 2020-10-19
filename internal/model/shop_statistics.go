package model

import (
	"time"

	"api.xinfos.com/driver"
	"github.com/jinzhu/gorm"
)

//ShopStatistics 店铺统计
type ShopStatistics struct {
	ID                uint64    `json:"-"`                  //自增ID
	ShopID            uint64    `json:"-"`                  //店铺ID
	StayPaid          uint64    `json:"stay_paid"`          //待付款
	StayDelivered     uint64    `json:"stay_delivered"`     //待发货
	StaySign          uint64    `json:"stay_sign"`          //待签收
	StayRefund        uint64    `json:"stay_refund"`        //待退款
	StayComment       uint64    `json:"stay_comment"`       //待评价
	AbnormalLogistics uint64    `json:"abnormal_logistics"` //物流异常
	UpdatedAt         time.Time `json:"updated_at"`         //更新时间
	IsDelete          uint      `json:"-"`                  //是否删除 [1:是 | 2: 否]
}

var shopStatistics *ShopStatistics

//ShopStatisticsModel - Shop model
func ShopStatisticsModel() *ShopStatistics {
	return shopStatistics
}

//TableName - Return table name
func (t *ShopStatistics) TableName() string {
	return `t_shop_statistics`
}

//Create - create shop statistics
func (t *ShopStatistics) Create() error {
	if err := driver.DB.Table(t.TableName()).Create(&t).Error; err != nil {
		return err
	}
	return nil
}

//Delete - delete shop statistics
func (t *ShopStatistics) Delete() error {
	if err := driver.DB.Table(t.TableName()).Where("shop_id = (?)", t.ShopID).Update("is_delete", 1).Error; err != nil {
		return err
	}
	return nil
}

//Update - update shop statistics
func (t *ShopStatistics) Update() error {
	if err := driver.DB.Table(t.TableName()).Where("shop_id = (?)", t.ShopID).Update(t).Error; err != nil {
		return err
	}
	return nil
}

//FindByShopID - Find shop by shop_id
func (t *ShopStatistics) FindByShopID(shopID uint64) (*ShopStatistics, error) {
	return t.findByMap(map[string]interface{}{"shop_id": shopID, "is_delete": 2})
}

func (t *ShopStatistics) findByMap(wheremaps map[string]interface{}) (*ShopStatistics, error) {
	var data ShopStatistics
	if err := driver.DB.Table(t.TableName()).Where(wheremaps).Find(&data).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, err
		}
	}
	return &data, nil
}
