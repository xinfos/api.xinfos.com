package service

import (
	"api.xinfos.com/internal/model"
	"api.xinfos.com/internal/repository"
	"api.xinfos.com/utils/errs"
)

//ProductService - 用户服务
type ProductService struct {
	SAttrGroupMapRepo *repository.SAttrGroupMapRepository
}

//NewProductService - 商品服务
func NewProductService() *ProductService {
	return &ProductService{
		SAttrGroupMapRepo: repository.NewSAttrGroupMapRepository(),
	}
}

//BeforeCreateByCatID -
func (s *ProductService) BeforeCreateByCatID(catID uint64) ([]*model.SysSPUAttrGroup, *errs.Errs) {
	attrs, err := s.SAttrGroupMapRepo.FindAllByCatID(catID)
	if err != nil || len(attrs) <= 0 {
		return nil, errs.ErrProductNoAttr
	}
	return attrs, nil
}
