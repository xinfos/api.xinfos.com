package service

import (
	"api.xinfos.com/internal/model"
	"api.xinfos.com/internal/repository"
	"api.xinfos.com/utils/errs"
)

//AttrService - Attr Service
type AttrService struct {
	Repo *repository.SAttrRepository
}

//NewAttrService - New Attr Service
func NewAttrService() *AttrService {
	return &AttrService{
		Repo: repository.NewSAttrRepository(),
	}
}

//Create - Create a piece of category attr data
// func (s *AttrService) Create(m *model.SAttr) (uint64, *errs.Errs) {
// 	return s.Repo.Create(m)
// }

func (s *AttrService) Query(query string) ([]*model.SAttrBlock, *errs.Errs) {
	return s.Repo.Query(query)
}
