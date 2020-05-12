package repository

import (
	"api.xinfos.com/internal/model"
	"api.xinfos.com/internal/repository/cache"
	"api.xinfos.com/utils/errs"
)

//BrandRepository - 品牌仓库
type BrandRepository struct {
	c *cache.BrandCache
}

//NewBrandRepository - 初始化 BrandRepository
func NewBrandRepository() *BrandRepository {
	return &BrandRepository{
		c: cache.NewBrandCache(),
	}
}

//BrandList - Brand list strcut
type BrandList struct {
	List            []*model.Brand `json:"list"`
	CurrentPageNo   uint           `json:"current_page_no"`
	CurrentPageSize uint           `json:"current_page_size"`
	TotalCount      int            `json:"total_count"`
}

//Create - Create a brand
func (repo *BrandRepository) Create(m *model.Brand) (uint64, *errs.Errs) {
	//1.check if the category exists

	//2.Check the brand name is exists.
	if m.IsBrandNameExists(m.BrandName) {
		return 0, errs.ErrBrandCreateFailNameIsExists
	}
	//3.Create
	if err := m.Create(); err != nil {
		return 0, errs.ErrBrandCreateFail
	}
	return m.BrandID, nil
}

//Delete - Delete a brand
func (repo *BrandRepository) Delete(id uint64) *errs.Errs {
	//1.check if the brand exists.
	m, err := model.BrandModel().FindByID(id)
	if err != nil || m == nil || m.BrandID <= 0 {
		return errs.ErrBrandDeleteFailNotFound
	}
	//2.Check if the brand has any prodcut
	if model.ProductModel().IsBrandHasProduct(id) {
		return errs.ErrBrandDeleteFailHasProduct
	}
	//3.Delete
	if err := m.Delete(); err != nil {
		return errs.ErrBrandDeleteFail
	}
	return nil
}

//Update - Update a brand
func (repo *BrandRepository) Update(m *model.Brand) *errs.Errs {
	//1.Check if the brand exists.
	data, err := m.FindByID(m.BrandID)
	if err != nil || data == nil || m.BrandID <= 0 {
		return errs.ErrBrandNotFound
	}
	//2.If not equal to the original name, check the brand name is exists.
	if m.BrandName != data.BrandName {
		if m.IsBrandNameExists(m.BrandName) {
			return errs.ErrBrandUpdateFailNameIsExists
		}
	}
	//3.Update
	if err := m.Update(); err != nil {
		return errs.ErrBrandUpdateFail
	}
	return nil
}

//FindByID - Find category by cat_id
func (repo *BrandRepository) FindByID(id uint64) (*model.Brand, *errs.Errs) {
	data := repo.c.Get(id)
	if data != nil && data.BrandID > 0 {
		return data, nil
	}
	data, _ = model.BrandModel().FindByID(id)
	if data == nil || data.BrandID <= 0 {
		return nil, errs.ErrBrandNotFound
	}
	repo.c.Set(data)
	return data, nil
}

//FindAll - Find category by cat_id
func (repo *BrandRepository) FindAll(query map[string]interface{}, orderby string, page, pageSize uint) (*BrandList, *errs.Errs) {
	data, count, err := model.BrandModel().FindAll(query, orderby, page, pageSize)
	if err != nil {
		return nil, errs.ErrBrandCreateFail
	}
	l := &BrandList{
		List:            data,
		CurrentPageNo:   page,
		CurrentPageSize: pageSize,
		TotalCount:      count,
	}
	return l, nil
}
