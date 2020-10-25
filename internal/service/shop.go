package service

import (
	"fmt"

	"api.xinfos.com/internal/model"
	"api.xinfos.com/internal/repository"
	"api.xinfos.com/utils/errs"
)

type ShopDashboard struct {
	Info      *model.Shop           `json:"info"`
	Statistis *model.ShopStatistics `json:"statistis"`
}

//ShopService 店铺服务
type ShopService struct {
	Repo          *repository.ShopRepository
	StatistisRepo *repository.ShopStatisticsRepository
}

//NewShopService 实例化店铺服务
func NewShopService() *ShopService {
	return &ShopService{
		Repo:          repository.NewShopRepository(),
		StatistisRepo: repository.NewShopStatisticsRepository(),
	}
}

//Dashboard 店铺概况
func (s *ShopService) Dashboard(sellerID uint64) (*ShopDashboard, *errs.Errs) {
	errmsg := &errs.Errs{}
	dashboard := &ShopDashboard{}

	dashboard.Info, errmsg = s.Repo.FindBySellerID(sellerID)
	if errmsg != nil {
		return nil, errmsg
	}
	fmt.Println(dashboard.Info.ID)
	if dashboard.Info.ID <= 0 {
		return nil, errs.ErrShopNotFound
	}
	dashboard.Statistis, _ = s.StatistisRepo.FindByShopID(dashboard.Info.ID)

	return dashboard, nil
}

//Create 创建店铺
func (s *ShopService) Create(m *model.Shop) (uint64, *errs.Errs) {
	return s.Repo.Create(m)
}

//Delete 删除店铺
func (s *ShopService) Delete(shopID, sellerID uint64) *errs.Errs {
	return s.Repo.Delete(shopID, sellerID)
}

//Update 更新店铺
func (s *ShopService) Update(m *model.Shop) *errs.Errs {
	return s.Repo.Update(m)
}

//FindByID 查询单个店铺详情
func (s *ShopService) FindByID(shopID, sellerID uint64) (*model.Shop, *errs.Errs) {
	return s.Repo.FindByShopIDAndSellerID(shopID, sellerID)
}

//FindBySellerID 查询卖家，店铺详情
func (s *ShopService) FindBySellerID(sellerID uint64) (*model.Shop, *errs.Errs) {
	return s.Repo.FindBySellerID(sellerID)
}

//FindAll 获取店铺列表
func (s *ShopService) FindAll(query string, args []interface{}, orderby string, page, pageSize uint) (*repository.ShopList, *errs.Errs) {
	return s.Repo.FindAll(query, args, orderby, page, pageSize)
}
