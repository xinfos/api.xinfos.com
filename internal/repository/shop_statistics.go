package repository

import (
	"fmt"

	"api.xinfos.com/internal/model"
	"api.xinfos.com/internal/repository/cache"
	"api.xinfos.com/utils/errs"
)

//ShopStatisticsRepository 店铺统计Repo
type ShopStatisticsRepository struct {
	c *cache.ShopStatisticsCache
}

//NewShopStatisticsRepository - 实例化店铺R统计Repo
func NewShopStatisticsRepository() *ShopStatisticsRepository {
	return &ShopStatisticsRepository{
		c: cache.NewShopStatisticsCache(),
	}
}

//Create 初始化店铺统计数据
func (repo *ShopStatisticsRepository) Create(m *model.ShopStatistics) (uint64, *errs.Errs) {

	//1.检查店铺名称是否已存在
	isExistsShopStatistics, err := m.FindByShopID(m.ShopID)
	if err == nil && isExistsShopStatistics != nil && isExistsShopStatistics.ID > 0 {
		return 0, errs.ErrShopStatisticsIsExists
	}

	//2.创建店铺
	m.IsDelete = 2
	err = m.Create()
	if err != nil {
		return 0, errs.ErrShopStatisticsCreateFail
	}
	return m.ID, nil
}

//Delete 删除店铺统计数据
func (repo *ShopStatisticsRepository) Delete(shopID uint64) *errs.Errs {

	//1.检查店铺是否已存在
	m, err := model.ShopStatisticsModel().FindByShopID(shopID)
	if err != nil || m == nil || m.ID <= 0 {
		return errs.ErrShopStatisticsNotFound
	}

	//2.删除店铺统计数据
	err = m.Delete()
	if err != nil {
		return errs.ErrShopStatisticsDeleteFail
		//TODO 删除店铺缓存
	}
	return nil
}

//Update 更新店铺统计
func (repo *ShopStatisticsRepository) Update(m *model.ShopStatistics) *errs.Errs {

	//1.检查当前店铺是否存在
	originalShopStatistics, err := m.FindByShopID(m.ShopID)
	if err != nil || originalShopStatistics == nil || originalShopStatistics.ID <= 0 {
		return errs.ErrShopNotFound
	}

	//2.检查哪些数据需要被更新
	//2.1 待发货
	if originalShopStatistics.StayDelivered != m.StayDelivered {
		originalShopStatistics.StayDelivered = m.StayDelivered
	}
	//2.1 待签收
	if originalShopStatistics.StaySign != m.StaySign {
		originalShopStatistics.StaySign = m.StaySign
	}
	//2.1 待退款
	if originalShopStatistics.StayRefund != m.StayRefund {
		originalShopStatistics.StayRefund = m.StayRefund
	}
	//2.1 待评价
	if originalShopStatistics.StayComment != m.StayComment {
		originalShopStatistics.StayComment = m.StayComment
	}
	//2.1 物流异常
	if originalShopStatistics.AbnormalLogistics != m.AbnormalLogistics {
		originalShopStatistics.AbnormalLogistics = m.AbnormalLogistics
	}

	//3.更新店铺
	err = originalShopStatistics.Update()
	if err != nil {
		return errs.ErrShopStatisticsUpdateFail
	}
	return nil
}

//FindByShopID 获取店铺统计数据
func (repo *ShopStatisticsRepository) FindByShopID(shopID uint64) (*model.ShopStatistics, *errs.Errs) {
	k := fmt.Sprintf(cache.CacheShopStatisticsKey, shopID)
	data := repo.c.Get(k)
	if data != nil && data.ID > 0 {
		return data, nil
	}
	data, err := model.ShopStatisticsModel().FindByShopID(shopID)
	if err != nil && data == nil && data.ShopID != shopID {
		return nil, errs.ErrShopStatisticsNotFound
	}
	repo.c.Set(k, data)
	return data, nil
}
