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
	cacheSAttrGroupMapKey     = "cache:sysattrgroupmap:%d"
	cacheSAttrGroupMapListKey = "cache:sysattrgroupmap:list:%s"
)

//SAttrGroupMapCache - SAttrGroupMapCache
type SAttrGroupMapCache struct {
	Rds *redis.Client
}

func NewSAttrGroupMapCache() *SAttrGroupMapCache {
	return &SAttrGroupMapCache{
		Rds: driver.Rds,
	}
}

func (u *SAttrGroupMapCache) Get(id uint64) (m *model.SAttrGroupMap) {
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

func (u *SAttrGroupMapCache) Set(m *model.SAttrGroupMap) bool {
	k := fmt.Sprintf(cacheSAttrGroupMapKey, m.CatID)
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

func (u *SAttrGroupMapCache) GetAll(key string) (m []*model.SysSPUAttrGroup) {
	k := fmt.Sprintf(cacheSAttrGroupMapListKey, key)
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

func (u *SAttrGroupMapCache) SetAll(key string, m []*model.SysSPUAttrGroup) bool {
	k := fmt.Sprintf(cacheSAttrGroupMapListKey, key)
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
