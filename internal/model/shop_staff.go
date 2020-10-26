package model

import (
	"fmt"

	"api.xinfos.com/driver"

	"github.com/jinzhu/gorm"
)

//ShopStaff - 店铺员工模型
type ShopStaff struct {
	ID        uint64 `json:"id"`
	ShopID    uint64 `json:"shop_id"`
	StaffNo   string `json:"staff_no"`
	Name      string `json:"name"`
	Mobile    string `json:"mobile"`
	State     uint   `json:"state"`
	CreatedAt string `json:"-"`
	UpdatedAt string `json:"-"`
	IsDelete  uint   `json:"-"`
}

var shopStaff *ShopStaff

//ShopStaffModel - 声明一个空的店铺员工模型
func ShopStaffModel() *ShopStaff {
	return shopStaff
}

//TableName - 返回对应的表名
func (t *ShopStaff) TableName() string {
	return `t_shop_staff`
}

//Create - create user
func (t *ShopStaff) Create() error {
	if err := driver.DB.Table(t.TableName()).Create(&t).Error; err != nil {
		return err
	}
	return nil
}

//Delete - Delete a record
func (t *ShopStaff) Delete() error {
	if err := driver.DB.Table(t.TableName()).Where("id = (?)", t.ID).Update("is_delete", 1).Error; err != nil {
		return err
	}
	return nil
}

//Update - Update a record
func (t *ShopStaff) Update() error {
	if err := driver.DB.Table(t.TableName()).Where("id = (?)", t.ID).Update(t).Error; err != nil {
		return err
	}
	return nil
}

//IsStaffNoExists - Check the staff_no is exists
func (t *ShopStaff) IsStaffNoExists(staffNo string, shopID uint64) bool {
	data, err := t.firstByMap("shop_id = (?) AND `staffNo` = (?) AND `is_delete` = 2", []interface{}{shopID, staffNo})
	if err == nil && data != nil && data.ID > 0 {
		return true
	}
	return false
}

//FindByID - Find Data By ID
func (t *ShopStaff) FindByID(id uint64) (*ShopStaff, error) {
	return t.findByMap(map[string]interface{}{"id": id, "is_delete": 2})
}

//FindByName - Find Data By name
func (t *ShopStaff) FindByName(name string, shopID uint64) (*ShopStaff, error) {
	return t.findByMap(map[string]interface{}{"name": name, "shop_id": shopID, "is_delete": 2})
}

//FindAll - Find all by query maps
func (t *ShopStaff) FindAll(query map[string]interface{}, orderby string, page, pageSize uint) ([]*ShopStaff, int, error) {
	return t.findAllByQueryMap("", query, orderby, page, pageSize)
}

//FindAllByIDs - Find List Data By Ids
func (t *ShopStaff) FindAllByIDs(ids []uint64, shopID uint64) ([]*ShopStaff, error) {
	return t.findAllByQueryCondition("`id` in (?) AND `shop_id` = (?) AND `is_delete` = 2", []interface{}{ids, shopID})
}

//FindAllByName - Find List Data By pid
func (t *ShopStaff) FindAllByName(name string, shopID uint64) ([]*ShopStaff, error) {
	return t.findAllByQueryCondition("`name` = (?) AND `shop_id` = (?) AND `is_delete` = 2", []interface{}{name, shopID})
}

//FindAllByQuery - Find all by query string
func (t *ShopStaff) FindAllByQuery(query string, args []interface{}, orderby, groupBy string, page, pageSize uint) ([]*ShopStaff, int, error) {
	return t.findAllByQuery("", query, args, orderby, groupBy, page, pageSize)
}

func (t *ShopStaff) firstByMap(query string, args []interface{}) (*ShopStaff, error) {
	var data ShopStaff
	if err := driver.DB.Table(t.TableName()).Where(query, args...).First(&data).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, err
		}
	}
	return &data, nil
}

func (t *ShopStaff) findByMap(wheremaps map[string]interface{}) (*ShopStaff, error) {
	var data ShopStaff
	if err := driver.DB.Table(t.TableName()).Where(wheremaps).Find(&data).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, err
		}
	}
	return &data, nil
}

func (t *ShopStaff) findAllByQueryCondition(query string, args []interface{}) ([]*ShopStaff, error) {
	var data []*ShopStaff
	if err := driver.DB.Table(t.TableName()).Where(query, args...).Find(&data).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, err
		}
	}
	return data, nil
}

func (t *ShopStaff) findAllByQueryMap(fields string, query map[string]interface{}, orderBy string, page, pageSize uint) (data []*ShopStaff, count int, err error) {
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

func (t *ShopStaff) findAllByQuery(fields string, query string, args []interface{}, orderby, groupBy string, page, pageSize uint) (data []*ShopStaff, count int, err error) {

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
