package model

import (
	"time"

	"api.xinfos.com/driver"

	"github.com/jinzhu/gorm"
)

//Shop - Shop Model struct
type Shop struct {
	SID       uint64    `json:"sid" gorm:"PRIMARY_KEY"`
	SellerID  uint64    `json:"seller_id"`
	Name      string    `json:"name"`
	Desc      string    `json:"desc"`
	Logo      string    `json:"logo"`
	URL       string    `json:"url"`
	CertType  uint      `json:"cert_type"`
	State     uint      `json:"state"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	IsDelete  uint      `json:"is_delete"`
}

var shop *Shop

//ShopModel - Shop model
func ShopModel() *Shop {
	return shop
}

//TableName - Return table name
func (t *Shop) TableName() string {
	return `t_shop`
}

//Create - create user
func (t *Shop) Create() error {
	if err := driver.DB.Table(t.TableName()).Create(&t).Error; err != nil {
		return err
	}
	return nil
}

//Delete - Delete a record
func (t *Shop) Delete() error {
	if err := driver.DB.Table(t.TableName()).Where("s_id = (?)", t.SID).Update("is_delete", 1).Error; err != nil {
		return err
	}
	return nil
}

//Update - Update a record
func (t *Shop) Update() error {
	if err := driver.DB.Table(t.TableName()).Where("s_id = (?)", t.SID).Update(t).Error; err != nil {
		return err
	}
	return nil
}

//IsBrandNameExists - Check the brand name is exists
func (t *Shop) IsBrandNameExists(name string) bool {
	data, err := t.firstByMap("`name` = (?) AND `is_delete` = 2", []interface{}{name})
	if err == nil && data != nil && data.SID > 0 {
		return true
	}
	return false
}

//FindByID - Find Data By ID
func (t *Shop) FindByID(id uint64) (*Shop, error) {
	return t.findByMap(map[string]interface{}{"s_id": id, "is_delete": 2})
}

//FindByName - Find Data By name
func (t *Shop) FindByName(name string) (*Shop, error) {
	return t.findByMap(map[string]interface{}{"name": name, "is_delete": 2})
}

//FindAll - Find all by query maps
func (t *Shop) FindAll(query map[string]interface{}, orderby string, page, pageSize uint) ([]*Shop, int, error) {
	return t.findAllByMap("", query, orderby, page, pageSize)
}

//FindAllByIDs - Find List Data By Ids
func (t *Shop) FindAllByIDs(ids []uint64) ([]*Shop, error) {
	return t.findAllByQueryCondition("`s_id` in (?) AND `is_delete` = 2", []interface{}{ids})
}

func (t *Shop) findByMap(wheremaps map[string]interface{}) (*Shop, error) {
	var data Shop
	if err := driver.DB.Table(t.TableName()).Where(wheremaps).Find(&data).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, err
		}
	}
	return &data, nil
}

func (t *Shop) firstByMap(query string, args []interface{}) (*Shop, error) {
	var data Shop
	if err := driver.DB.Table(t.TableName()).Where(query, args...).First(&data).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, err
		}
	}
	return &data, nil
}

func (t *Shop) findAllByMap(fields string, query map[string]interface{}, orderBy string, page, pageSize uint) (data []*Shop, count int, err error) {
	offset := (page - 1) * pageSize
	if len(fields) <= 0 {
		fields = "*"
	}
	query["is_delete"] = 2
	dbSelector := driver.DB.Table(t.TableName()).Select(fields).Where(query)
	if err := dbSelector.Offset(offset).Limit(pageSize).Order(orderBy).Find(&data).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, count, err
		}
	}
	if err := dbSelector.Count(&count).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, 0, err
		}
	}
	return data, count, nil
}
func (t *Shop) findAllByQueryCondition(query string, args []interface{}) ([]*Shop, error) {
	var data []*Shop
	if err := driver.DB.Table(t.TableName()).Where(query, args...).Find(&data).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, err
		}
	}
	return data, nil
}
