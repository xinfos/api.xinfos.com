package repository

import (
	"api.xinfos.com/internal/model"
	"api.xinfos.com/internal/repository/cache"
	"api.xinfos.com/utils/errs"
)

//SAttrTemplateRepository - System defined standard classification attribute template
type SAttrTemplateRepository struct {
	c *cache.SAttrGroupCache
}

//NewSAttrTemplateRepository - 初始化 NewSAttrTemplateRepository
func NewSAttrTemplateRepository() *SAttrTemplateRepository {
	return &SAttrTemplateRepository{
		c: cache.NewSAttrGroupCache(),
	}
}

//Create - 根据model模型创建记录
func (repo *SAttrTemplateRepository) Create(m *model.SAttrTempalte, AttrGroupIDs []uint64) (uint64, *errs.Errs) {
	if m.CatID <= 0 {
		return 0, errs.ErrAttrGroupCreateFailCateNotFound
	}
	cate, err := model.CategoryModel().FindByID(m.CatID)
	if err != nil || cate.CatID != m.CatID {
		return 0, errs.ErrAttrGroupCreateFailCateNotFound
	}

	existsSAttrGroup, err := m.FindByCatIDAndName(m.CatID, m.Name)
	if err != nil {
		return 0, errs.ErrAttrGroupCreateFailCateNotFound
	}

	if existsSAttrGroup != nil && existsSAttrGroup.ID > 0 {
		return existsSAttrGroup.ID, nil
	}

	m.IsDelete = 2
	err = m.Create()
	if err != nil {
		return 0, errs.ErrAttrGroupCreateFail
	}
	if m == nil || m.ID <= 0 {
		return 0, errs.ErrAttrGroupCreateFail
	}
	return m.ID, nil

	return 0, nil
}
