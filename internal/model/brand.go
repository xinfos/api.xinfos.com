package model

import (
	"time"

	"api.xinfos.com/driver"

	"github.com/jinzhu/gorm"
)

type Brand struct {
	BrandID   uint64    `json:"brand_id" gorm:"PRIMARY_KEY"`
	BrandName string    `json:"brand_name"`
	BrandLogo string    `json:"brand_logo"`
	BrandDesc string    `json:"brand_desc"`
	CatID     uint64    `json:"cat_id"`
	SortOrder uint32    `json:"sort_order"`
	IsShow    uint      `json:"is_show"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	IsDelete  uint      `json:"is_delete"`
}

var brand *Brand

//BrandModel - Brand model
func BrandModel() *Brand {
	return brand
}

//TableName - Return table name
func (t *Brand) TableName() string {
	return `t_brand`
}

//Create - create user
func (t *Brand) Create() error {
	if err := driver.DB.Table(t.TableName()).Create(&t).Error; err != nil {
		return err
	}
	return nil
}

//Delete - Delete a record
func (t *Brand) Delete() error {
	if err := driver.DB.Table(t.TableName()).Where("brand_id = (?)", t.BrandID).Update("is_delete", 1).Error; err != nil {
		return err
	}
	return nil
}

//Update - Update a record
func (t *Brand) Update() error {
	if err := driver.DB.Table(t.TableName()).Where("brand_id = (?)", t.BrandID).Update(t).Error; err != nil {
		return err
	}
	return nil
}

//IsBrandNameExists - Check the brand name is exists
func (t *Brand) IsBrandNameExists(name string) bool {
	data, err := t.firstByMap("`brand_name` = (?) AND `is_delete` = 2", []interface{}{name})
	if err == nil && data != nil && data.BrandID > 0 {
		return true
	}
	return false
}

//FindByID - Find Data By ID
func (t *Brand) FindByID(id uint64) (*Brand, error) {
	return t.findByMap(map[string]interface{}{"brand_id": id, "is_delete": 2})
}

//FindByName - Find Data By name
func (t *Brand) FindByName(name string) (*Brand, error) {
	return t.findByMap(map[string]interface{}{"brand_name": name, "is_delete": 2})
}

//FindAllByIDs - Find List Data By Ids
func (t *Brand) FindAllByIDs(ids []uint64) ([]*Brand, error) {
	return t.findAllByQueryCondition("`brand_id` in (?) AND `is_delete` = 2", []interface{}{ids})
}

//FindAllByCatID - Find List Data By CatID
func (t *Brand) FindAllByCatID(id uint64) ([]*Brand, error) {
	return t.findAllByQueryCondition("`cat_id` in (?) AND `is_delete` = 2", []interface{}{id})
}

func (t *Brand) findByMap(wheremaps map[string]interface{}) (*Brand, error) {
	var data Brand
	if err := driver.DB.Table(t.TableName()).Where(wheremaps).Find(&data).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, err
		}
	}
	return &data, nil
}

func (t *Brand) firstByMap(query string, args []interface{}) (*Brand, error) {
	var data Brand
	if err := driver.DB.Table(t.TableName()).Where(query, args...).First(&data).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, err
		}
	}
	return &data, nil
}

func (t *Brand) findAllByQueryCondition(query string, args []interface{}) ([]*Brand, error) {
	var data []*Brand
	if err := driver.DB.Table(t.TableName()).Where(query, args...).Find(&data).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, err
		}
	}
	return data, nil
}
