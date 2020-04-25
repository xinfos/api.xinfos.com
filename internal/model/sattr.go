package model

import (
	"api.xinfos.com/driver"

	"github.com/jinzhu/gorm"
)

//SAttr -
type SAttr struct {
	ID          uint64 `json:"id"`
	SAttrID     uint64 `json:"s_attr_id"`
	SGroupID    uint64 `json:"s_group_id"`
	CatID       uint64 `json:"cat_id"`
	Name        string `json:"name"`
	FillType    uint   `json:"fill_type"`
	IsRequired  uint   `json:"is_required"`
	IsNumeric   uint   `json:"is_numeric"`
	Unit        string `json:"unit"`
	IsGeneric   uint   `json:"is_generic"`
	IsSearching uint   `json:"is_searching"`
	Segments    string `json:"segments"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	IsDelete    uint   `json:"-"`
}

var sAttr *SAttr

func SAttrModel() *SAttr {
	return sAttr
}

func (t *SAttr) TableName() string {
	return `t_system_spu_attr`
}

//Create - create user
func (t *SAttr) Create() error {
	if err := driver.DB.Table(t.TableName()).Create(&t).Error; err != nil {
		return err
	}
	return nil
}

//FindByID -
func (t *SAttr) FindByID(id uint64) (*SAttr, error) {
	return t.findByMap(map[string]interface{}{
		"id":        id,
		"is_delete": 2,
	})
}

func (t *SAttr) FindAllByCatID(id uint64) ([]*SAttr, error) {
	return t.findAllByQueryCondition("`cat_id` in (?) AND `is_delete` = 2", []interface{}{id})
}

func (t *SAttr) FindAllBySAttrIDs(ids []uint64) ([]*SAttr, error) {
	return t.findAllByQueryCondition("`s_attr_id` in (?) AND `is_delete` = 2", []interface{}{ids})
}

func (t *SAttr) findByMap(wheremaps map[string]interface{}) (*SAttr, error) {
	var data SAttr
	if err := driver.DB.Table(t.TableName()).Where(wheremaps).Find(&data).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, err
		}
	}
	return &data, nil
}

func (t *SAttr) findAllByQueryCondition(query string, args []interface{}) ([]*SAttr, error) {
	var data []*SAttr
	if err := driver.DB.Table(t.TableName()).Where(query, args...).Find(&data).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, err
		}
	}
	return data, nil
}
