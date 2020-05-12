package service

import (
	"api.xinfos.com/internal/repository"
)

//AttrService - Attr Service
type AttrService struct {
	Repo *repository.BrandRepository
}

//NewAttrService - New Attr Service
func NewAttrService() *AttrService {
	return &AttrService{
		Repo: repository.NewBrandRepository(),
	}
}

//Create - Create a piece of category attr data
// func (s *AttrService) Create(m *model.SAttr) (uint64, *errs.Errs) {
// 	return s.Repo.Create(m)
// }
