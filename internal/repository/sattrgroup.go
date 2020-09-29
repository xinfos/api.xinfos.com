package repository

import (
	"fmt"

	"api.xinfos.com/internal/model"
	"api.xinfos.com/internal/repository/cache"
	"api.xinfos.com/utils/errs"
)

//SAttrGroupRepository - System defind Attribute Group Repository
type SAttrGroupRepository struct {
	c *cache.SAttrGroupCache
}

//NewSAttrGroupRepository - 初始化 NewSAttrGroupRepository
func NewSAttrGroupRepository() *SAttrGroupRepository {
	return &SAttrGroupRepository{
		c: cache.NewSAttrGroupCache(),
	}
}

//Create - 根据model模型创建记录
func (repo *SAttrGroupRepository) Create(m *model.SAttrGroup) (uint64, *errs.Errs) {
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
	if existsSAttrGroup != nil || existsSAttrGroup.ID > 0 {
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
}

//FindByID - 根据ID获取信息
func (repo *SAttrGroupRepository) FindByID(id uint64) (*model.SAttrGroup, error) {
	data := repo.c.Get(id)
	if data != nil && data.ID > 0 {
		return data, nil
	}
	data, _ = model.SAttrGroupModel().FindBySGroupID(id)
	if data != nil && data.ID == id {
		repo.c.Set(data)
	}
	return data, nil
}

func (repo *SAttrGroupRepository) FindAllByCatID(id uint64) ([]*model.SAttrGroup, error) {
	k := fmt.Sprintf("catid:%d", id)
	data := repo.c.GetAll(k)
	if data != nil && len(data) > 0 {
		return data, nil
	}
	data, _ = model.SAttrGroupModel().FindAllByCatID(id)
	if data != nil && len(data) > 0 {
		repo.c.SetAll(k, data)
	}
	return data, nil
}

func (repo *SAttrGroupRepository) FindBySGroupIDs(ids []uint64) ([]*model.SAttrGroup, error) {
	data, _ := model.SAttrGroupModel().FindBySGroupIDs(ids)
	if data != nil && len(data) > 0 {

	}
	return data, nil
}

func (repo *SAttrGroupRepository) FindAll() (*model.SAttrGroup, error) {
	return nil, nil
}
