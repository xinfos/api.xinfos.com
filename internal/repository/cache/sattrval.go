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
	cacheSAttrValKey     = "cache:sysattrval:%d"
	cacheSAttrValListKey = "cache:sysattrval:list:%s"
)

//SAttrValCache - SAttrValCache
type SAttrValCache struct {
	Rds *redis.Client
}

func NewSAttrValCache() *SAttrValCache {
	return &SAttrValCache{
		Rds: driver.Rds,
	}
}

func (u *SAttrValCache) Get(id uint64) (m *model.SAttrVal) {
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

func (u *SAttrValCache) Set(m *model.SAttrVal) bool {
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

func (u *SAttrValCache) GetAll(key string) (m []*model.SAttrVal) {
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

func (u *SAttrValCache) SetAll(key string, m []*model.SAttrVal) bool {
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
