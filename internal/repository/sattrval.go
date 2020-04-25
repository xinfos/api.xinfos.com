package repository

import (
	"errors"
	"fmt"

	"api.xinfos.com/internal/model"
	"api.xinfos.com/internal/repository/cache"
)

//SAttrRepository - 系统属性仓库
type SAttrValRepository struct {
	c *cache.SAttrValCache
}

//NewSAttrValRepository - SAttrValRepository
func NewSAttrValRepository() *SAttrValRepository {
	return &SAttrValRepository{
		c: cache.NewSAttrValCache(),
	}
}

//Create - 根据model模型创建记录
func (repo *SAttrValRepository) Create(m *model.SAttrVal) (uint64, error) {
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
func (repo *SAttrValRepository) FindByID(id uint64) (*model.SAttrVal, error) {
	data := repo.c.Get(id)
	if data != nil && data.ID > 0 {
		return data, nil
	}
	data, _ = model.SAttrValModel().FindByID(id)
	if data != nil && data.ID == id {
		repo.c.Set(data)
	}
	return data, nil
}

func (repo *SAttrValRepository) FindAllByCatID(id uint64) ([]*model.SAttrVal, error) {
	k := fmt.Sprintf("catid:%d", id)
	data := repo.c.GetAll(k)
	if data != nil && len(data) > 0 {
		return data, nil
	}
	data, _ = model.SAttrValModel().FindAllByCatID(id)
	if data != nil && len(data) > 0 {
		repo.c.SetAll(k, data)
	}
	return data, nil
}

//FindBySAttrIDs - 根据ID获取信息
func (repo *SAttrValRepository) FindBySAttrIDs(ids []uint64) ([]*model.SAttrVal, error) {
	data, _ := model.SAttrValModel().FindBySAttrIDs(ids)
	if data != nil && len(data) > 0 {
		// repo.c.Set(data)
	}
	return data, nil
}

func (repo *SAttrValRepository) FindAll() (*model.SAttr, error) {
	return nil, nil
}
