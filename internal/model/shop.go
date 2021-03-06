package model

import (
	"time"

	"api.xinfos.com/driver"

	"github.com/jinzhu/gorm"
)

//Shop - Shop Model struct
type Shop struct {
	ID        uint64    `json:"shop_id"`
	SellerID  uint64    `json:"seller_id"`
	Name      string    `json:"name"`
	Type      uint      `json:"type" binding:"required"`
	Location  string    `json:"location"`
	Address   string    `json:"address"`
	IsAgree   uint      `json:"is_agree" binding:"required"`
	Desc      string    `json:"desc"`
	Logo      string    `json:"logo"`
	URL       string    `json:"url"`
	CertType  uint      `json:"cert_type"`
	State     uint      `json:"state"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	IsDelete  uint      `json:"-"`
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
	if err := driver.DB.Table(t.TableName()).Where("id = (?)", t.ID).Update("is_delete", 1).Error; err != nil {
		return err
	}
	return nil
}

//Update - Update a record
func (t *Shop) Update() error {
	if err := driver.DB.Table(t.TableName()).Where("id = (?)", t.ID).Update(t).Error; err != nil {
		return err
	}
	return nil
}

//IsBrandNameExists - Check the brand name is exists
func (t *Shop) IsBrandNameExists(name string) bool {
	data, err := t.firstByMap("`name` = (?) AND `is_delete` = 2", []interface{}{name})
	if err == nil && data != nil && data.ID > 0 {
		return true
	}
	return false
}

//FindByID - Find Data By ID
func (t *Shop) FindByID(id uint64) (*Shop, error) {
	return t.findByMap(map[string]interface{}{"id": id, "is_delete": 2})
}

//FindByName - Find Data By name
func (t *Shop) FindByName(name string) (*Shop, error) {
	return t.findByMap(map[string]interface{}{"name": name, "is_delete": 2})
}

//FindByShopIDAndSellerID - Find shop by shop_id & seller_id
func (t *Shop) FindByShopIDAndSellerID(id, sellerID uint64) (*Shop, error) {
	return t.findByMap(map[string]interface{}{"id": id, "seller_id": sellerID, "is_delete": 2})
}

//FindBySellerID - Find shop by seller_id
func (t *Shop) FindBySellerID(sellerID uint64) (*Shop, error) {
	return t.findByMap(map[string]interface{}{"seller_id": sellerID, "is_delete": 2})
}

//FindAll - Find all by query maps
func (t *Shop) FindAll(query map[string]interface{}, orderby string, page, pageSize uint) ([]*Shop, int, error) {
	return t.findAllByMap("", query, orderby, page, pageSize)
}

//FindAllByIDs - Find List Data By Ids
func (t *Shop) FindAllByIDs(ids []uint64) ([]*Shop, error) {
	return t.findAllByQueryCondition("`id` in (?) AND `is_delete` = 2", []interface{}{ids})
}

//FindAllByQuery - Find all by query string
func (t *Shop) FindAllByQuery(query string, args []interface{}, orderby, groupBy string, page, pageSize uint) ([]*Shop, int, error) {
	return t.findAllByQuery("", query, args, orderby, groupBy, page, pageSize)
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

func (t *Shop) findAllByQuery(fields string, query string, args []interface{}, orderby, groupBy string, page, pageSize uint) (data []*Shop, count int, err error) {

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
