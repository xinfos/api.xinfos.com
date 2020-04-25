package repository

import (
	"errors"
	"fmt"

	"api.xinfos.com/internal/model"
	"api.xinfos.com/internal/repository/cache"
)

//SAttrRepository - 系统属性仓库
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
func (repo *SAttrGroupRepository) Create(m *model.SAttrGroup) (uint64, error) {
	err := m.Create()
	if err != nil {
		return 0, err
	}
	if m == nil || m.ID <= 0 {
		return 0, errors.New("create fail")
	}
	return m.ID, nil
}

//FindByID - 根据ID获取信息
func (repo *SAttrGroupRepository) FindByID(id uint64) (*model.SAttrGroup, error) {
	data := repo.c.Get(id)
	if data != nil && data.ID > 0 {
		return data, nil
	}
	data, _ = model.SAttrGroupModel().FindByID(id)
	if data != nil && data.ID == id {
		repo.c.Set(data)
	}
	return data, nil
}

func (repo *SAttrGroupRepository) FindBySGroupID(id uint64) (*model.SAttrGroup, error) {
	data := repo.c.Get(id)
	if data != nil && data.ID > 0 {
		return data, nil
	}
	data, _ = model.SAttrGroupModel().FindByID(id)
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
