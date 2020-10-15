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

//FindByID - Find category by id
func (s *AttrService) FindByID(id uint64) (*model.SAttr, *errs.Errs) {
	return s.Repo.FindByID(id)
}

//FindAll 查询属性相关的属性列表
func (s *AttrService) FindAll(query string, args []interface{}, orderby string, page, pageSize uint) (*repository.AttrList, *errs.Errs) {
	return s.Repo.FindAll(query, args, orderby, page, pageSize)
}

//Query 根据具体条件查询相关属性信息
func (s *AttrService) Query(query string) ([]*model.SAttrBlock, *errs.Errs) {
	return s.Repo.Query(query)
}
