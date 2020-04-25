package model

import (
	"api.xinfos.com/driver"

	"github.com/jinzhu/gorm"
)

//SAttrVal -
type SAttrVal struct {
	ID         uint64 `json:"id"`
	SAttrValID uint64 `json:"s_attr_val_id"`
	SAttrID    uint64 `json:"s_attr_id"`
	CatID      uint64 `json:"cat_id"`
	Value      string `json:"value"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
	IsDelete   uint   `json:"-"`
}

var sAttrVal *SAttrVal

func SAttrValModel() *SAttrVal {
	return sAttrVal
}

func (t *SAttrVal) TableName() string {
	return `t_system_spu_attr_value`
}

//Create - create user
func (t *SAttrVal) Create() error {
	if err := driver.DB.Table(t.TableName()).Create(&t).Error; err != nil {
		return err
	}
	return nil
}

func (t *SAttrVal) FindByID(id uint64) (*SAttrVal, error) {
	return t.findByMap(map[string]interface{}{
		"id":        id,
		"is_delete": 2,
	})
}

func (t *SAttrVal) FindBySAttrValID(id uint64) (*SAttrVal, error) {
	return t.findByMap(map[string]interface{}{
		"s_attr_val_id": id,
		"is_delete":     2,
	})
}

func (t *SAttrVal) FindBySAttrID(id uint64) (*SAttrVal, error) {
	return t.findByMap(map[string]interface{}{
		"s_attr_id": id,
		"is_delete": 2,
	})
}

func (t *SAttrVal) FindAllByCatID(id uint64) ([]*SAttrVal, error) {
	return t.findAllByQueryCondition("`cat_id` in (?) AND `is_delete` = 2", []interface{}{id})
}

func (t *SAttrVal) FindBySAttrIDs(ids []uint64) ([]*SAttrVal, error) {
	return t.findAllByQueryCondition("`s_attr_id` in (?) AND `is_delete` = 2", []interface{}{ids})
}

func (t *SAttrVal) findByMap(wheremaps map[string]interface{}) (*SAttrVal, error) {
	var data SAttrVal
	if err := driver.DB.Table(t.TableName()).Where(wheremaps).Find(&data).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, err
		}
	}
	return &data, nil
}

func (t *SAttrVal) findAllByQueryCondition(query string, args []interface{}) ([]*SAttrVal, error) {
	var data []*SAttrVal
	if err := driver.DB.Table(t.TableName()).Where(query, args...).Find(&data).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, err
		}
	}
	return data, nil
}
