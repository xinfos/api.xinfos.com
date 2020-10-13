package service

import (
	"api.xinfos.com/internal/model"
	"api.xinfos.com/internal/repository"
	"api.xinfos.com/utils/errs"
)

//AttrTemplateService - Attribute Template Service
type AttrTemplateService struct {
	Repo *repository.SAttrTemplateRepository
}

//NewAttrTemplateService - Init Attribute Template Service
func NewAttrTemplateService() *AttrTemplateService {
	return &AttrTemplateService{
		Repo: repository.NewSAttrTemplateRepository(),
	}
}

//Create - Create a single attribute group
func (s *AttrTemplateService) Create(m *model.SAttrTempalte, generalAttrsGroupIDs, specAttrsIDs []uint64) (uint64, *errs.Errs) {
	return s.Repo.Create(m, generalAttrsGroupIDs, specAttrsIDs)
}
