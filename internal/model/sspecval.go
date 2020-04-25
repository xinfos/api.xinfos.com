package model

import (
	"api.xinfos.com/driver"

	"github.com/jinzhu/gorm"
)

//SSpecVal -
type SSpecVal struct {
	ID         uint64 `json:"id"`
	SSpecValID uint64 `json:"s_spec_val_id"`
	SSpecID    uint64 `json:"s_spec_id"`
	Name       string `json:"name"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
	IsDelete   uint   `json:"-"`
}

var sSpecVal *SSpecVal

func SSpecValModel() *SSpecVal {
	return sSpecVal
}

func (t *SSpecVal) TableName() string {
	return `t_system_spu_attr_value`
}

//Create - create user
func (t *SSpecVal) Create() error {
	if err := driver.DB.Table(t.TableName()).Create(&t).Error; err != nil {
		return err
	}
	return nil
}

//FindByID -
func (t *SSpecVal) FindByID(SSpecValID uint64) (*SSpecVal, error) {
	return t.findByMap(map[string]interface{}{
		"s_spec_val_id": SSpecValID,
		"is_delete":     2,
	})
}

func (t *SSpecVal) findByMap(wheremaps map[string]interface{}) (*SSpecVal, error) {
	var data SSpecVal
	if err := driver.DB.Table(t.TableName()).Where(wheremaps).Find(&data).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, err
		}
	}
	return &data, nil
}
