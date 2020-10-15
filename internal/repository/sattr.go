package repository

import (
	"errors"
	"fmt"

	"api.xinfos.com/internal/model"
	"api.xinfos.com/internal/repository/cache"
	"api.xinfos.com/utils/errs"
)

type AttrList struct {
	List            []*model.SAttr `json:"list"`
	CurrentPageNo   uint           `json:"current_page_no"`
	CurrentPageSize uint           `json:"current_page_size"`
	TotalCount      int            `json:"total_count"`
}

//SAttrRepository - 系统属性仓库
type SAttrRepository struct {
	c *cache.SAttrCache
}

//NewSAttrRepository - 初始化 SAttrRepository
func NewSAttrRepository() *SAttrRepository {
	return &SAttrRepository{
		c: cache.NewSAttrCache(),
	}
}

//Create - 根据model模型创建记录
func (repo *SAttrRepository) Create(m *model.SAttr) (uint64, error) {
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
func (repo *SAttrRepository) FindByID(id uint64) (*model.SAttr, *errs.Errs) {
	data := repo.c.Get(id)
	if data != nil && data.ID > 0 {
		return data, nil
	}
	data, _ = model.SAttrModel().FindByID(id)
	if data == nil || data.ID != id {
		return nil, nil
	}
	repo.c.Set(data)
	return data, nil
}

func (repo *SAttrRepository) FindAllByCatID(id uint64) ([]*model.SAttr, error) {
	k := fmt.Sprintf("catid:%d", id)
	data := repo.c.GetAll(k)
	if data != nil && len(data) > 0 {
		return data, nil
	}
	data, _ = model.SAttrModel().FindAllByCatID(id)
	if data != nil && len(data) > 0 {
		repo.c.SetAll(k, data)
	}
	return data, nil
}

func (repo *SAttrRepository) FindBySAttrIDs(ids []uint64) ([]*model.SAttr, error) {
	data, _ := model.SAttrModel().FindAllBySAttrIDs(ids)
	if data != nil && len(data) > 0 {
	}
	return data, nil
}

func (repo *SAttrRepository) FindAll(query string, args []interface{}, orderby string, page, pageSize uint) (*AttrList, *errs.Errs) {
	data, count, err := model.SAttrModel().FindAllByQuery(query, args, orderby, "", page, pageSize)
	if err != nil {
		return nil, errs.ErrBrandCreateFail
	}
	l := &AttrList{
		List:            data,
		CurrentPageNo:   page,
		CurrentPageSize: pageSize,
		TotalCount:      count,
	}
	return l, nil
}

//Query query attr
func (repo *SAttrRepository) Query(query string) ([]*model.SAttrBlock, *errs.Errs) {
	data, err := model.SAttrModel().FindAllBySAttrName(query)
	if err != nil {
		return nil, nil
	}
	sAttrsBlock := []*model.SAttrBlock{}
	if data != nil && len(data) > 0 {
		for _, v := range data {
			sAttrsBlock = append(sAttrsBlock, &model.SAttrBlock{
				ID:   v.ID,
				Name: v.Name,
			})
		}
		return sAttrsBlock, nil
	}
	return nil, nil
}
