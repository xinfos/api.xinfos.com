package service

import (
	"api.xinfos.com/internal/model"
	"api.xinfos.com/internal/repository"
	"api.xinfos.com/utils/errs"
)

//CategoryService - Category Service
type CategoryService struct {
	Repo *repository.CategoryRepository
}

//NewCategoryService - New Category Service
func NewCategoryService() *CategoryService {
	return &CategoryService{
		Repo: repository.NewCategoryRepository(),
	}
}

//Create - Create a piece of category data
func (s *CategoryService) Create(m *model.Category) (uint64, *errs.Errs) {
	return s.Repo.Create(m)
}

//Delete - Delete a category by cat_id
func (s *CategoryService) Delete(id uint64) *errs.Errs {
	return s.Repo.Delete(id)
}

//Update - Update a category by cat_id
func (s *CategoryService) Update(m *model.Category) *errs.Errs {
	return s.Repo.Update(m)
}

//FindByID - Find category by id
func (s *CategoryService) FindByID(id uint64) (*model.Category, *errs.Errs) {
	return s.Repo.FindByID(id)
}

//FindAllByPID - Find sub category list by pid
func (s *CategoryService) FindAllByPID(id uint64) ([]*model.Category, *errs.Errs) {
	return s.Repo.FindAllByPID(id)
}
