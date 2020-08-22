package service

import (
	"api.xinfos.com/internal/model"
	"api.xinfos.com/internal/repository"
	"api.xinfos.com/utils/errs"
)

//ShopStaffService - Shop Staff Service
type ShopStaffService struct {
	Repo *repository.ShopStaffRepository
}

//NewShopStaffService - New ShopStaff Service
func NewShopStaffService() *ShopStaffService {
	return &ShopStaffService{
		Repo: repository.NewShopStaffRepository(),
	}
}

//Create - Create a new shop employee
func (s *ShopStaffService) Create(m *model.ShopStaff) (uint64, *errs.Errs) {
	return s.Repo.Create(m)
}

//Delete - Delete a shop employee
func (s *ShopStaffService) Delete(id uint64) *errs.Errs {
	return s.Repo.Delete(id)
}

//Update - Update a shop employee
func (s *ShopStaffService) Update(m *model.ShopStaff) *errs.Errs {
	return s.Repo.Update(m)
}

//FindByID - Find a shop employee by ID
func (s *ShopStaffService) FindByID(id uint64) (*model.ShopStaff, *errs.Errs) {
	return s.Repo.FindByID(id)
}

type ShopStaffList struct {
	List       []*model.ShopStaff `json:"list"`
	TotalCount int                `json:"total_count"`
}

//FindAll - Find list of shop employee data
func (s *ShopStaffService) FindAll(query map[string]interface{}, orderby string, page, pageSize uint) (*ShopStaffList, *errs.Errs) {
	data, count, err := s.Repo.FindAll(query, orderby, page, pageSize)
	if err != nil {
		return nil, err
	}
	return &ShopStaffList{
		List:       data,
		TotalCount: count,
	}, nil
}
