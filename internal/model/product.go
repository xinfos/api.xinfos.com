package model

import (
	"time"

	"api.xinfos.com/driver"
	"github.com/jinzhu/gorm"
)

type Product struct {
	ID        uint64    `json:"id"`
	ProductID uint64    `json:"product_id"`
	BrandID   uint64    `json:"brand_id"`
	CatID     uint64    `json:"cat_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	IsDelete  uint      `json:"is_delete"`
}

var product *Product

//ProductModel - Brand model
func ProductModel() *Product {
	return product
}

//TableName - Return table name
func (t *Product) TableName() string {
	return `t_product`
}

//IsCategoryHadSubChild - Check the Category Has Sub Child.
func (t *Product) IsCategoryHadSubChild(id uint64) bool {
	data, err := t.firstByMap("`parent_cat_id` = (?) AND `is_delete` = 2", []interface{}{id})
	if err == nil && data != nil {
		return true
	}
	return false
}

//IsBrandHasProduct - Check the brand products are included
func (t *Product) IsBrandHasProduct(id uint64) bool {
	data, err := t.firstByMap("`brand_id` = (?) AND `is_delete` = 2", []interface{}{id})
	if err == nil && data != nil {
		return true
	}
	return false
}

//FindByID - Find Data By ID
func (t *Product) FindByID(id uint64) (*Product, error) {
	return t.findByMap(map[string]interface{}{"cat_id": id, "is_delete": 2})
}

//FindByPID - Find Data By PID
func (t *Product) FindByPID(id uint64) (*Product, error) {
	return t.findByMap(map[string]interface{}{"parent_cat_id": id, "is_delete": 2})
}

//FindByName - Find Data By name
func (t *Product) FindByName(name string) (*Product, error) {
	return t.findByMap(map[string]interface{}{"name": name, "is_delete": 2})
}

//FindAllByIDs - Find List Data By Ids
func (t *Product) FindAllByIDs(ids []uint64) ([]*Product, error) {
	return t.findAllByQueryCondition("`cat_id` in (?) AND `is_delete` = 2", []interface{}{ids})
}

//FindAllByCatID - Find List Data By CatID
func (t *Product) FindAllByCatID(id uint64) ([]*Product, error) {
	return t.findAllByQueryCondition("`cat_id` in (?) AND `is_delete` = 2", []interface{}{id})
}

//FindAllByPID - Find List Data By pid
func (t *Product) FindAllByPID(id uint64) ([]*Product, error) {
	return t.findAllByQueryCondition("`parent_cat_id` in (?) AND `is_delete` = 2", []interface{}{id})
}

func (t *Product) firstByMap(query string, args []interface{}) (*Product, error) {
	var data Product
	if err := driver.DB.Table(t.TableName()).Where(query, args...).First(&data).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, err
		}
	}
	return &data, nil
}

func (t *Product) findByMap(wheremaps map[string]interface{}) (*Product, error) {
	var data Product
	if err := driver.DB.Table(t.TableName()).Where(wheremaps).Find(&data).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, err
		}
	}
	return &data, nil
}

func (t *Product) findAllByQueryCondition(query string, args []interface{}) ([]*Product, error) {
	var data []*Product
	if err := driver.DB.Table(t.TableName()).Where(query, args...).Find(&data).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, err
		}
	}
	return data, nil
}
