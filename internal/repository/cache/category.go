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
	cacheCategoryKey     = "cache:category:%d"
	cacheCategoryListKey = "cache:category:list:%s"
)

//CategoryCache - Category Cache
type CategoryCache struct {
	Rds *redis.Client
}

func NewCategoryCache() *CategoryCache {
	return &CategoryCache{
		Rds: driver.Rds,
	}
}

func (u *CategoryCache) Get(id uint64) (m *model.Category) {
	key := fmt.Sprintf(cacheCategoryKey, id)
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

func (u *CategoryCache) Set(m *model.Category) bool {
	k := fmt.Sprintf(cacheCategoryKey, m.CatID)
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

func (u *CategoryCache) GetAll(key string) (m []*model.Category) {
	k := fmt.Sprintf(cacheCategoryListKey, key)
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

func (u *CategoryCache) SetAll(key string, m []*model.Category) bool {
	k := fmt.Sprintf(cacheCategoryListKey, key)
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
