package service

import (
	"api.xinfos.com/internal/model"
	"api.xinfos.com/internal/repository"
	"api.xinfos.com/utils/errs"
)

//BrandService - Brand Service
type BrandService struct {
	Repo *repository.BrandRepository
}

//NewBrandService - New Brand Service
func NewBrandService() *BrandService {
	return &BrandService{
		Repo: repository.NewBrandRepository(),
	}
}

//Create - Create a piece of brand data
func (s *BrandService) Create(m *model.Brand) (uint64, *errs.Errs) {
	return s.Repo.Create(m)
}

//Delete - Delete a category by id
func (s *BrandService) Delete(id uint64) *errs.Errs {
	return s.Repo.Delete(id)
}

//Update - Update a piece of brand data
func (s *BrandService) Update(m *model.Brand) *errs.Errs {
	return s.Repo.Update(m)
}

//FindByID - Find a piece of brand data by id
func (s *BrandService) FindByID(id uint64) (*model.Brand, *errs.Errs) {
	return s.Repo.FindByID(id)
}
