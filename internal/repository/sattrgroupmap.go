package repository

import (
	"fmt"

	"api.xinfos.com/internal/model"
	"api.xinfos.com/internal/repository/cache"
)

//SAttrGroupMapRepository - 系统属性仓库
type SAttrGroupMapRepository struct {
	c *cache.SAttrGroupMapCache
}

//NewSAttrGroupMapRepository - 初始化 NewSAttrGroupMapRepository
func NewSAttrGroupMapRepository() *SAttrGroupMapRepository {
	return &SAttrGroupMapRepository{
		c: cache.NewSAttrGroupMapCache(),
	}
}

//FindByID - 根据ID获取信息
func (repo *SAttrGroupMapRepository) FindByID(id uint64) (*model.SAttrGroupMap, error) {
	return nil, nil
}

//FindAllByCatID - 根据商品系统分类ID查询SPU系统关联属性
func (repo *SAttrGroupMapRepository) FindAllByCatID(id uint64) ([]*model.SysSPUAttrGroup, error) {
	k := fmt.Sprintf("catid:%d", id)
	data := repo.c.GetAll(k)
	if data != nil && len(data) > 0 {
		return data, nil
	}
	data, _ = model.SAttrGroupMapModel().FindAllByCatID(id)
	if data != nil && len(data) > 0 {
		repo.c.SetAll(k, data)
	}
	return data, nil
}
