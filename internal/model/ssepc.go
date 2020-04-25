package model

import (
	"api.xinfos.com/driver"

	"github.com/jinzhu/gorm"
)

//SSpec -
type SSpec struct {
	ID        uint64 `json:"id"`
	SSpecID   uint64 `json:"s_spec_id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	IsDelete  uint   `json:"-"`
}

var sSpec *SSpec

func SSpecModel() *SSpec {
	return sSpec
}

func (t *SSpec) TableName() string {
	return `t_system_spu_attr_value`
}

//Create - create user
func (t *SSpec) Create() error {
	if err := driver.DB.Table(t.TableName()).Create(&t).Error; err != nil {
		return err
	}
	return nil
}

//FindByID - 根据`user_id`查询用户信息
func (t *SSpec) FindByID(SSpecID uint64) (*SSpec, error) {
	return t.findByMap(map[string]interface{}{
		"s_spec_id": SSpecID,
		"is_delete": 2,
	})
}

func (t *SSpec) findByMap(wheremaps map[string]interface{}) (*SSpec, error) {
	var data SSpec
	if err := driver.DB.Table(t.TableName()).Where(wheremaps).Find(&data).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, err
		}
	}
	return &data, nil
}
