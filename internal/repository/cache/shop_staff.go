package cache

import (
	"fmt"
	"time"

	"api.xinfos.com/driver"
	"api.xinfos.com/internal/model"
	"github.com/go-redis/redis"
	"github.com/vmihailenco/msgpack/v4"
)

const (
	cacheShopStaffKey     = "cache:ShopStaff:%d"
	cacheShopStaffListKey = "cache:ShopStaff:list:%s"
)

//ShopStaffCache - 店铺员工缓存
type ShopStaffCache struct {
	Rds *redis.Client
}

//NewShopStaffCache - 实例化店铺员工缓存服务
func NewShopStaffCache() *ShopStaffCache {
	return &ShopStaffCache{
		Rds: driver.Rds,
	}
}

//Get - 获取单个店铺员工信息
func (u *ShopStaffCache) Get(id uint64) (m *model.ShopStaff) {
	key := fmt.Sprintf(cacheSAttrKey, id)
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

//Set - 存储单个店铺员工信息
func (u *ShopStaffCache) Set(m *model.ShopStaff) bool {
	k := fmt.Sprintf(cacheSAttrValKey, m.ID)
	v, err := msgpack.Marshal(m)
	if err != nil {
		return false
	}
	isOk, err := u.Rds.SetNX(k, v, 100*time.Second).Result()
	if err != nil {
		return false
	}
	return isOk
}

//GetAll - 查询店铺员工列表
func (u *ShopStaffCache) GetAll(key string) (m []*model.ShopStaff) {
	k := fmt.Sprintf(cacheSAttrValListKey, key)
	v, err := u.Rds.Get(k).Bytes()
	if err != nil {
		return nil
	}
	err = msgpack.Unmarshal(v, &m)
	if err != nil {
		return nil
	}
	return m
}

//SetAll - 存储店铺员工列表
func (u *ShopStaffCache) SetAll(key string, m []*model.ShopStaff) bool {
	k := fmt.Sprintf(cacheSAttrValListKey, key)
	v, err := msgpack.Marshal(m)
	if err != nil {
		return false
	}
	isOk, err := u.Rds.SetNX(k, v, 100*time.Second).Result()
	if err != nil {
		return false
	}
	return isOk
}
