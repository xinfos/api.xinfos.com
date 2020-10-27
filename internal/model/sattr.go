package model

import (
	"fmt"
	"time"

	"api.xinfos.com/driver"

	"github.com/jinzhu/gorm"
)

//SAttr -
type SAttr struct {
	ID          uint64    `json:"id"`
	Name        string    `json:"name"`
	FillType    uint      `json:"fill_type"`
	IsRequired  uint      `json:"is_required"`
	IsNumeric   uint      `json:"is_numeric"`
	Unit        string    `json:"unit"`
	IsGeneric   uint      `json:"is_generic"`
	IsSearching uint      `json:"is_searching"`
	Segments    string    `json:"segments"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
	IsDelete    uint      `json:"-"`
}

type SAttrBlock struct {
	ID   uint64 `json:"attr_id"`
	Name string `json:"attr_name"`
}

var sAttr *SAttr

//SAttrModel 实例化 SAttr 模型
func SAttrModel() *SAttr {
	return sAttr
}

//TableName 返回对应的表名
func (t *SAttr) TableName() string {
	return `t_sys_spu_attr`
}

//Create - create user
func (t *SAttr) Create() error {
	if err := driver.DB.Table(t.TableName()).Create(&t).Error; err != nil {
		return err
	}
	return nil
}

//Delete - Delete a record
func (t *SAttr) Delete() error {
	if err := driver.DB.Table(t.TableName()).Where("id = (?)", t.ID).Update("is_delete", 1).Error; err != nil {
		return err
	}
	return nil
}

//Update - Update a record
func (t *SAttr) Update() error {
	if err := driver.DB.Table(t.TableName()).Where("id = (?)", t.ID).Update(t).Error; err != nil {
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

//FindBySAttrName 根据属性名查询属性
func (t *SAttr) FindBySAttrName(name string) (*SAttr, error) {
	return t.findByMap(map[string]interface{}{
		"name":      name,
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

func (t *SAttr) FindAllBySAttrName(name string) ([]*SAttr, error) {
	return t.findAllByQueryCondition("`name` like (?) AND `is_delete` = 2", []interface{}{name})
}

//FindAllByQuery - Find all by query string
func (t *SAttr) FindAllByQuery(query string, args []interface{}, orderby, groupBy string, page, pageSize uint) ([]*SAttr, int, error) {
	return t.findAllByQuery("", query, args, orderby, groupBy, page, pageSize)
}

//IsSAttrNameExists - Check the attr name is exists
func (t *SAttr) IsSAttrNameExists(name string) bool {
	data, err := t.firstByMap("`name` = (?) AND `is_delete` = 2", []interface{}{name})
	if err == nil && data != nil && data.ID > 0 {
		return true
	}
	return false
}

func (t *SAttr) firstByMap(query string, args []interface{}) (*SAttr, error) {
	var data SAttr
	if err := driver.DB.Table(t.TableName()).Where(query, args...).First(&data).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, err
		}
	}
	return &data, nil
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

func (t *SAttr) findAllByQuery(fields string, query string, args []interface{}, orderby, groupBy string, page, pageSize uint) (data []*SAttr, count int, err error) {

	ts := driver.DB.Table(t.TableName()).Where(query, args...)
	if len(fields) > 0 {
		ts = ts.Select(fields)
	}

	if len(orderby) > 0 {
		ts = ts.Order(orderby)
	}
	if len(groupBy) > 0 {
		ts = ts.Group(groupBy)
	}

	if page <= 0 {
		page = 1
	}

	fmt.Println(pageSize)
	if err := ts.Offset((page - 1) * pageSize).Limit(pageSize).Find(&data).Error; err != nil {
		return nil, 0, err
	}

	if err := ts.Count(&count).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, 0, err
		}
	}
	return data, count, nil
}
