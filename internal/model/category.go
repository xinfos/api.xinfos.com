package model

import (
	"time"

	"api.xinfos.com/driver"

	"github.com/jinzhu/gorm"
)

type Category struct {
	CatID     uint64    `json:"cat_id" gorm:"PRIMARY_KEY"`
	PID       uint64    `json:"pid" gorm:"Column:parent_cat_id"`
	Name      string    `json:"name"`
	Alias     string    `json:"alias"`
	Desc      string    `json:"desc"`
	Depth     uint      `json:"depth"`
	ShowInNav uint      `json:"show_in_nav"`
	IsShow    uint      `json:"is_show"`
	IsParent  uint      `json:"is_parent"`
	State     uint      `json:"state"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	IsDelete  uint      `json:"-"`
}

var category *Category

//CategoryModel - Category Model
func CategoryModel() *Category {
	return category
}

//TableName - Return table name
func (t *Category) TableName() string {
	return `t_category`
}

//GetRootCategory 获取根分类
func (t *Category) GetRootCategory() *Category {
	return &Category{
		CatID:     10000,
		PID:       0,
		Name:      "商品根分类",
		Alias:     "商品根分类",
		Desc:      "商品根分类",
		Depth:     0,
		ShowInNav: 0,
		IsShow:    0,
		IsParent:  1,
		State:     1,
	}
}

//Create - create user
func (t *Category) Create() error {
	if err := driver.DB.Table(t.TableName()).Create(&t).Error; err != nil {
		return err
	}
	return nil
}

//Delete - Delete a record
func (t *Category) Delete() error {
	if err := driver.DB.Table(t.TableName()).Where("cat_id = (?)", t.CatID).Update("is_delete", 1).Error; err != nil {
		return err
	}
	return nil
}

//Update - Update a record
func (t *Category) Update() error {
	if err := driver.DB.Table(t.TableName()).Where("cat_id = (?)", t.CatID).Update(t).Error; err != nil {
		return err
	}
	return nil
}

//IsCategoryHadSubChild - Check the Category Has Sub Child.
func (t *Category) IsCategoryHadSubChild(id uint64) bool {
	data, err := t.firstByMap("`parent_cat_id` = (?) AND `is_delete` = 2", []interface{}{id})
	if err == nil && data != nil && data.CatID > 0 {
		return true
	}
	return false
}

//IsCategoryNameExists - Check the category name is exists
func (t *Category) IsCategoryNameExists(name string) bool {
	data, err := t.firstByMap("`name` = (?) AND `is_delete` = 2", []interface{}{name})
	if err == nil && data != nil && data.CatID > 0 {
		return true
	}
	return false
}

//FindByID - Find Data By ID
func (t *Category) FindByID(id uint64) (*Category, error) {
	return t.findByMap(map[string]interface{}{"cat_id": id, "is_delete": 2})
}

//FindByPID - Find Data By PID
func (t *Category) FindByPID(id uint64) (*Category, error) {
	return t.findByMap(map[string]interface{}{"parent_cat_id": id, "is_delete": 2})
}

//FindByName - Find Data By name
func (t *Category) FindByName(name string) (*Category, error) {
	return t.findByMap(map[string]interface{}{"name": name, "is_delete": 2})
}

//FindAll - Find all by query maps
func (t *Category) FindAll(query map[string]interface{}, orderby string, page, pageSize uint) ([]*Category, int, error) {
	return t.findAllByMap("", query, orderby, page, pageSize)
}

//FindAllByQuery - Find all by query string
func (t *Category) FindAllByQuery(query string, args []interface{}, orderby, groupby string, page, pageSize uint) ([]*Category, int, error) {
	return t.findAllByQuery("", query, args, orderby, groupby, page, pageSize)
}

//FindAllByIDs - Find List Data By Ids
func (t *Category) FindAllByIDs(ids []uint64) ([]*Category, error) {
	return t.findAllByQueryCondition("`cat_id` in (?) AND `is_delete` = 2", []interface{}{ids})
}

//FindAllByCatID - Find List Data By CatID
func (t *Category) FindAllByCatID(id uint64) ([]*Category, error) {
	return t.findAllByQueryCondition("`cat_id` in (?) AND `is_delete` = 2", []interface{}{id})
}

//FindAllByPID - Find List Data By pid
func (t *Category) FindAllByPID(id uint64) ([]*Category, error) {
	return t.findAllByQueryCondition("`parent_cat_id` in (?) AND `is_delete` = 2", []interface{}{id})
}

func (t *Category) firstByMap(query string, args []interface{}) (*Category, error) {
	var data Category
	if err := driver.DB.Table(t.TableName()).Where(query, args...).First(&data).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, err
		}
	}
	return &data, nil
}

func (t *Category) findByMap(wheremaps map[string]interface{}) (*Category, error) {
	var data Category
	if err := driver.DB.Table(t.TableName()).Where(wheremaps).Find(&data).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, err
		}
	}
	return &data, nil
}

func (t *Category) findAllByMap(fields string, query map[string]interface{}, orderBy string, page, pageSize uint) (data []*Category, count int, err error) {
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

func (t *Category) findAllByQueryCondition(query string, args []interface{}) ([]*Category, error) {
	var data []*Category
	if err := driver.DB.Table(t.TableName()).Where(query, args...).Find(&data).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, err
		}
	}
	return data, nil
}

func (t *Category) findAllByQuery(fields string, query string, args []interface{}, orderby, groupby string, page, pageSize uint) (data []*Category, count int, err error) {

	ts := driver.DB.Table(t.TableName()).Where(query, args...)
	if len(fields) > 0 {
		ts = ts.Select(fields)
	}

	if len(orderby) > 0 {
		ts = ts.Order(orderby)
	}
	if len(groupby) > 0 {
		ts = ts.Group(groupby)
	}

	if page <= 0 {
		page = 1
	}

	if err := ts.Offset((page - 1) * pageSize).Limit(page).Find(&data).Error; err != nil {
		return nil, 0, err
	}

	if err := ts.Count(&count).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, 0, err
		}
	}
	return data, count, nil
}
