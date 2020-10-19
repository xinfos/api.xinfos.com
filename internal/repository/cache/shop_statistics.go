package cache

import (
	"time"

	"api.xinfos.com/driver"
	"api.xinfos.com/internal/model"
	"github.com/go-redis/redis"
	"github.com/vmihailenco/msgpack/v4"
)

const (
	//CacheShopStatisticsKey 店铺统计详情
	CacheShopStatisticsKey = "cache:Shop:Statistics:%d"
)

//ShopStatisticsCache - 店铺统计缓存
type ShopStatisticsCache struct {
	Rds *redis.Client
}

//NewShopStatisticsCache - 实例化店铺统计缓存服务
func NewShopStatisticsCache() *ShopStatisticsCache {
	return &ShopStatisticsCache{
		Rds: driver.Rds,
	}
}

//Get - 获取单个店铺员工信息
func (u *ShopStatisticsCache) Get(key string) (m *model.ShopStatistics) {
	v, err := u.Rds.Get(key).Bytes()
	if err != nil {
		return nil
	}
	err = msgpack.Unmarshal(v, &m)
	if err != nil {
		return nil
	}
	return m
}

//Set - 存储单个店铺信息
func (u *ShopStatisticsCache) Set(key string, m *model.ShopStatistics) bool {
	v, err := msgpack.Marshal(m)
	if err != nil {
		return false
	}
	isOk, err := u.Rds.SetNX(key, v, 100*time.Second).Result()
	if err != nil {
		return false
	}
	return isOk
}
