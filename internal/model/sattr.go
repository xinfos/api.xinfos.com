package model

import (
	"api.xinfos.com/driver"

	"github.com/jinzhu/gorm"
)

//SAttr -
type SAttr struct {
	ID          uint64 `json:"id"`
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

//SAttrModel 实例化 SAttr 模型
func SAttrModel() *SAttr {
	return sAttr
}

//TableName 返回对应的表名
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

//FindByID - 根据属性ID查询属性信息
func (t *SAttr) FindByID(id uint64) (*SAttr, error) {
	return t.findByMap(map[string]interface{}{
		"id":        id,
		"is_delete": 2,
	})
}

//FindAllByCatID 根据分类ID相关属性
func (t *SAttr) FindAllByCatID(id uint64) ([]*SAttr, error) {
	return t.findAllByQueryCondition("`cat_id` in (?) AND `is_delete` = 2", []interface{}{id})
}

//FindAllBySAttrIDs - 根据属性ID，批量查询属性信息
func (t *SAttr) FindAllBySAttrIDs(ids []uint64) ([]*SAttr, error) {
	return t.findAllByQueryCondition("`id` in (?) AND `is_delete` = 2", []interface{}{ids})
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
