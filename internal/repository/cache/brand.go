package cache

import (
	"api.xinfos.com/driver"
	"api.xinfos.com/internal/model"
	"fmt"
	"time"

	"github.com/go-redis/redis"
	"github.com/vmihailenco/msgpack/v4"
)

const (
	cacheBrandKey     = "cache:brand:%d"
	cacheBrandListKey = "cache:brand:list:%s"
)

//BrandCache - Brand Cache
type BrandCache struct {
	Rds *redis.Client
}

func NewBrandCache() *BrandCache {
	return &BrandCache{
		Rds: driver.Rds,
	}
}

func (u *BrandCache) Get(id uint64) (m *model.Brand) {
	key := fmt.Sprintf(cacheBrandKey, id)
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

func (u *BrandCache) Set(m *model.Brand) bool {
	k := fmt.Sprintf(cacheBrandKey, m.BrandID)
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

func (u *BrandCache) GetAll(key string) (m []*model.Brand) {
	k := fmt.Sprintf(cacheBrandListKey, key)
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

func (u *BrandCache) SetAll(key string, m []*model.Brand) bool {
	k := fmt.Sprintf(cacheBrandListKey, key)
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
