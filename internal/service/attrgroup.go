package service

import (
	"api.xinfos.com/internal/model"
	"api.xinfos.com/internal/repository"
	"api.xinfos.com/utils/errs"
)

//AttrGroupService - Attribute Group Service
type AttrGroupService struct {
	Repo *repository.SAttrGroupRepository
}

//NewAttrGroupService - Attribute Group Service
func NewAttrGroupService() *AttrGroupService {
	return &AttrGroupService{
		Repo: repository.NewSAttrGroupRepository(),
	}
}

//Create - Create a single attribute group
func (s *AttrGroupService) Create(m *model.SAttrGroup) (uint64, *errs.Errs) {
	return s.Repo.Create(m)
}
