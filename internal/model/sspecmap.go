package model

import (
	"api.xinfos.com/driver"

	"github.com/jinzhu/gorm"
)

//SSpecMap -
type SSpecMap struct {
	ID        uint64 `json:"id"`
	CatID     uint64 `json:"cat_id"`
	SSpecID   uint64 `json:"s_spec_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	IsDelete  uint   `json:"-"`
}

var sSpecMap *SSpecMap

func SSpecMapModel() *SSpecMap {
	return sSpecMap
}

func (t *SSpecMap) TableName() string {
	return `t_system_sku_spec_map`
}

//Create - create user
func (t *SSpecMap) Create() error {
	if err := driver.DB.Table(t.TableName()).Create(&t).Error; err != nil {
		return err
	}
	return nil
}

//FindByID - 根据`user_id`查询用户信息
func (t *SSpecMap) FindByID(SSpecID uint64) (*SSpecMap, error) {
	return t.findByMap(map[string]interface{}{
		"s_spec_id": SSpecID,
		"is_delete": 2,
	})
}

func (t *SSpecMap) findByMap(wheremaps map[string]interface{}) (*SSpecMap, error) {
	var data SSpecMap
	if err := driver.DB.Table(t.TableName()).Where(wheremaps).Find(&data).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, err
		}
	}
	return &data, nil
}
