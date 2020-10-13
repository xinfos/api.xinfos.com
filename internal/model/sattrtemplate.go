package model

import (
	"time"

	"api.xinfos.com/driver"
	"github.com/jinzhu/gorm"
)

//SAttrTempalte - System Standard attribute template
type SAttrTempalte struct {
	ID        uint64    `json:"id"`
	CatID     uint64    `json:"cat_id"`
	Name      string    `json:"name"`
	State     uint8     `json:"state"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	IsDelete  uint      `json:"-"`
}

var sAttrTempalte *SAttrTempalte

//SAttrTempalteModel Return SAttrTempalteModel
func SAttrTempalteModel() *SAttrTempalte {
	return sAttrTempalte
}

//TableName - Return the corresponding database table
func (t *SAttrTempalte) TableName() string {
	return `t_system_spu_attr_template`
}

//Create - create user
func (t *SAttrTempalte) Create(attrGroupIDs, specAttrIDs []uint64) error {

	if len(attrGroupIDs) > 0 || len(specAttrIDs) > 0 {
		return t.txCreate(attrGroupIDs, specAttrIDs)
	}
	if err := driver.DB.Table(t.TableName()).Create(&t).Error; err != nil {
		return err
	}
	return nil
}

//txCreate - 开启事务创建属性模板
func (t *SAttrTempalte) txCreate(attrGroupIDs, specAttrIDs []uint64) error {
	tx := driver.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}

	//1.创建模板基础数据
	if err := tx.Create(&t).Error; err != nil {
		tx.Rollback()
		return err
	}

	//2.创建模板和属性组关系
	if len(attrGroupIDs) > 0 {
		templateGroupIdx := []Row{}
		for _, v := range attrGroupIDs {
			templateGroupIdx = append(templateGroupIdx, Row{
				t.ID,
				v,
				2,
			})
		}
		// (db *gorm.DB, table string, columns []string, rows []Row)
		if err := BatchInsertRawSQL(tx, sAttrTemplateGroupIdx.TableName(), []string{"s_template_id", "s_group_id", "is_delete"}, templateGroupIdx); err != nil {
			tx.Rollback()
			return err
		}
	}

	//3.创建模板和规格属性关系
	if len(specAttrIDs) > 0 {
		templateSpecIdx := []Row{}
		for _, v := range specAttrIDs {
			templateSpecIdx = append(templateSpecIdx, Row{
				t.ID,
				v,
				2,
			})
		}
		if err := BatchInsertRawSQL(tx, sAttrTemplateSpecIdx.TableName(), []string{"s_template_id", "s_attr_id", "is_delete"}, templateSpecIdx); err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()
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
