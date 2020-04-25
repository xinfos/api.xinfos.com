package model

import (
	"api.xinfos.com/driver"

	"github.com/jinzhu/gorm"
)

//SAttrGroup -
type SAttrGroup struct {
	ID        uint64 `json:"id"`
	SGroupID  uint64 `json:"s_group_id"`
	CatID     uint64 `json:"cat_id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	IsDelete  uint   `json:"-"`
}

var sAttrGroup *SAttrGroup

func SAttrGroupModel() *SAttrGroup {
	return sAttrGroup
}

func (t *SAttrGroup) TableName() string {
	return `t_system_spu_attr_group`
}

//Create - create user
func (t *SAttrGroup) Create() error {
	if err := driver.DB.Table(t.TableName()).Create(&t).Error; err != nil {
		return err
	}
	return nil
}

//FindByID -
func (t *SAttrGroup) FindByID(id uint64) (*SAttrGroup, error) {
	return t.findByMap(map[string]interface{}{
		"id":        id,
		"is_delete": 2,
	})
}

func (t *SAttrGroup) FindBySGroupID(id uint64) (*SAttrGroup, error) {
	return t.findByMap(map[string]interface{}{
		"s_group_id": id,
		"is_delete":  2,
	})
}

func (t *SAttrGroup) FindAllByCatID(id uint64) ([]*SAttrGroup, error) {
	return t.findAllByQueryCondition("`cat_id` = (?) AND `is_delete` = 2", []interface{}{id})
}

func (t *SAttrGroup) FindBySGroupIDs(ids []uint64) ([]*SAttrGroup, error) {
	return t.findAllByQueryCondition("`s_group_id` in (?) AND `is_delete` = 2", []interface{}{ids})
}

func (t *SAttrGroup) findByMap(wheremaps map[string]interface{}) (*SAttrGroup, error) {
	var data SAttrGroup
	if err := driver.DB.Table(t.TableName()).Where(wheremaps).Find(&data).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, err
		}
	}
	return &data, nil
}

func (t *SAttrGroup) findAllByMap(wheremaps map[string]interface{}) ([]*SAttrGroup, error) {
	var data []*SAttrGroup
	if err := driver.DB.Table(t.TableName()).Where(wheremaps).Find(&data).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, err
		}
	}
	return data, nil
}

func (t *SAttrGroup) findAllByQueryCondition(query string, args []interface{}) ([]*SAttrGroup, error) {
	var data []*SAttrGroup
	if err := driver.DB.Table(t.TableName()).Where(query, args...).Find(&data).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, err
		}
	}
	return data, nil
}
