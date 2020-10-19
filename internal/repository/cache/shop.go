package cache

import (
	"time"

	"api.xinfos.com/driver"
	"api.xinfos.com/internal/model"
	"github.com/go-redis/redis"
	"github.com/vmihailenco/msgpack/v4"
)

const (
	//CacheShopKey 店铺详情
	CacheShopKey = "cache:Shop:%d:%d"
	//CacheShopListKey 店铺列表
	CacheShopListKey = "cache:ShopStaff:list:%s"
)

//ShopCache - 店铺缓存
type ShopCache struct {
	Rds *redis.Client
}

//NewShopCache - 实例化店铺缓存服务
func NewShopCache() *ShopCache {
	return &ShopCache{
		Rds: driver.Rds,
	}
}

//Get - 获取单个店铺员工信息
func (u *ShopCache) Get(key string) (m *model.Shop) {
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
func (u *ShopCache) Set(key string, m *model.Shop) bool {
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
