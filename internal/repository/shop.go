package repository

import (
	"fmt"

	"api.xinfos.com/internal/model"
	"api.xinfos.com/internal/repository/cache"
	"api.xinfos.com/utils/errs"
)

//ShopList 店铺列表
type ShopList struct {
	List            []*model.Shop `json:"list"`
	CurrentPageNo   uint          `json:"current_page_no"`
	CurrentPageSize uint          `json:"current_page_size"`
	TotalCount      int           `json:"total_count"`
}

//ShopRepository 店铺Repo
type ShopRepository struct {
	c *cache.ShopCache
}

//NewShopRepository - 实例化店铺Repo
func NewShopRepository() *ShopRepository {
	return &ShopRepository{
		c: cache.NewShopCache(),
	}
}

//Create 创建店铺
func (repo *ShopRepository) Create(m *model.Shop) (uint64, *errs.Errs) {
	//1.检查店铺名称是否已存在
	isExistsShop, err := m.FindByName(m.Name)
	if err == nil && isExistsShop != nil && isExistsShop.ID > 0 {
		return 0, errs.ErrShopNameIsExists
	}
	//2.创建店铺
	m.IsDelete = 2
	err = m.Create()
	if err != nil {
		return 0, errs.ErrShopCreateFail
	}
	return m.ID, nil
}

//Delete 删除店铺
func (repo *ShopRepository) Delete(shopID, sellerID uint64) *errs.Errs {
	//1.检查店铺是否已存在
	m, err := model.ShopModel().FindByShopIDAndSellerID(shopID, sellerID)
	if err != nil || m == nil || m.ID <= 0 {
		return errs.ErrShopNotFound
	}

	//2.Delete shop
	err = m.Delete()
	if err != nil {
		return errs.ErrShopDeleteFail
		//TODO 删除店铺缓存
	}
	return nil
}

//Update 更新店铺
func (repo *ShopRepository) Update(m *model.Shop) *errs.Errs {
	//1.检查当前店铺是否存在
	originalShop, err := m.FindByShopIDAndSellerID(m.ID, m.SellerID)
	if err != nil || originalShop == nil || originalShop.ID <= 0 {
		return errs.ErrShopNotFound
	}

	//2.检查哪些数据需要被更新
	if originalShop.Name != m.Name {
		//2.1 检查新的店铺名称是否已存在
		isExistsShop, err := m.FindByName(m.Name)
		if err == nil && isExistsShop != nil && isExistsShop.ID > 0 {
			return errs.ErrShopNameIsExists
		}
		originalShop.Name = m.Name
	}

	if originalShop.Desc != m.Desc {
		originalShop.Desc = m.Desc
	}

	if originalShop.Logo != m.Logo {
		originalShop.Logo = m.Logo
	}

	if originalShop.URL != m.URL {
		originalShop.URL = m.URL
	}

	if originalShop.State != m.State {
		originalShop.State = m.State
	}

	//3.更新店铺
	err = originalShop.Update()
	if err != nil {
		return errs.ErrShopUpdateFail
	}
	return nil
}

//FindByShopIDAndSellerID 根据店铺ID和卖家ID，获取店铺详情
func (repo *ShopRepository) FindByShopIDAndSellerID(shopID, sellerID uint64) (*model.Shop, *errs.Errs) {
	k := fmt.Sprintf(cache.CacheShopKey, shopID, sellerID)
	data := repo.c.Get(k)
	if data != nil && data.ID > 0 {
		return data, nil
	}
	data, err := model.ShopModel().FindByShopIDAndSellerID(shopID, sellerID)
	if err != nil && data == nil && data.ID != shopID {
		return nil, nil
	}
	repo.c.Set(k, data)
	return data, nil
}

//FindBySellerID 根据卖家ID，获取店铺详情
func (repo *ShopRepository) FindBySellerID(sellerID uint64) (*model.Shop, *errs.Errs) {
	k := fmt.Sprintf(cache.CacheShopKey, 1, sellerID)
	data := repo.c.Get(k)
	if data != nil && data.ID > 0 {
		return data, nil
	}
	data, err := model.ShopModel().FindBySellerID(sellerID)
	if err != nil && data == nil && data.SellerID != sellerID {
		return nil, nil
	}
	repo.c.Set(k, data)
	return data, nil
}

//FindAll 店铺列表
func (repo *ShopRepository) FindAll(query string, args []interface{}, orderby string, page, pageSize uint) (*ShopList, *errs.Errs) {
	data, count, err := model.ShopModel().FindAllByQuery(query, args, orderby, "", page, pageSize)
	if err != nil {
		return nil, errs.ErrShopNotFound
	}
	l := &ShopList{
		List:            data,
		CurrentPageNo:   page,
		CurrentPageSize: pageSize,
		TotalCount:      count,
	}
	return l, nil
}
