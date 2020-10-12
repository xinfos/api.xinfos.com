package model

import (
	"time"

	"api.xinfos.com/driver"
	"github.com/jinzhu/gorm"
)

//SAttrTempalte - System Standard attribute template
type SAttrTempalte struct {
	ID           uint64    `json:"id"`
	CatID        uint64    `json:"cat_id"`
	Name         string    `json:"name"`
	SGroupID     uint64    `json:"s_group_id"`
	SAttrID      uint64    `json:"s_attr_id"`
	Displayorder uint64    `json:"displayorder"`
	State        uint8     `json:"state"`
	CreatedAt    time.Time `json:"-"`
	UpdatedAt    time.Time `json:"-"`
	IsDelete     uint      `json:"-"`
}

var sAttrTempalte *SAttrTempalte

//SAttrTempalteModel Return SAttrTempalteModel
func SAttrTempalteModel() *SAttrTempalte {
	return sAttrTempalte
}

//TableName - Return the corresponding database table
func (t *SAttrTempalte) TableName() string {
	return `t_system_spu_attr_group_map`
}

//Create - create user
func (t *SAttrTempalte) Create(AttrGroupIDs []uint64) error {

	if len(AttrGroupIDs) > 0 {
		tx := driver.DB.Begin()
		defer func() {
			if r := recover(); r != nil {
				tx.Rollback()
			}
		}()
		if err := tx.Error; err != nil {
			return err
		}
		if err := tx.Create(&t).Error; err != nil {
			tx.Rollback()
			return err
		}
		return nil
	}
	if err := driver.DB.Table(t.TableName()).Create(&t).Error; err != nil {
		return err
	}
	return nil
}

//FindByCatIDAndName - find SAttrTempalte by `cat_id` and `name`
func (t *SAttrTempalte) FindByCatIDAndName(catID uint64, name string) (*SAttrTempalte, error) {
	return t.findByMap(map[string]interface{}{"cat_id": catID, "name": name, "is_delete": 2})
}

func (t *SAttrTempalte) findByMap(wheremaps map[string]interface{}) (*SAttrTempalte, error) {
	var data SAttrTempalte
	if err := driver.DB.Table(t.TableName()).Where(wheremaps).Find(&data).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, err
		}
	}
	return &data, nil
}
