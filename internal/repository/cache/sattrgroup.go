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
	cacheSAttrGroupKey     = "cache:sysattrgroup:%d"
	cacheSAttrGroupListKey = "cache:sysattrgroup:list:%s"
)

//SAttrGroupCache - user cache
type SAttrGroupCache struct {
	Rds *redis.Client
}

func NewSAttrGroupCache() *SAttrGroupCache {
	return &SAttrGroupCache{
		Rds: driver.Rds,
	}
}

func (u *SAttrGroupCache) Get(groupID uint64) (m *model.SAttrGroup) {
	key := fmt.Sprintf(cacheSAttrKey, groupID)
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

func (u *SAttrGroupCache) Set(m *model.SAttrGroup) bool {
	k := fmt.Sprintf(cacheSAttrGroupKey, m.ID)
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

func (u *SAttrGroupCache) GetAll(key string) (m []*model.SAttrGroup) {
	k := fmt.Sprintf(cacheSAttrGroupListKey, key)
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

func (u *SAttrGroupCache) SetAll(key string, m []*model.SAttrGroup) bool {
	k := fmt.Sprintf(cacheSAttrGroupListKey, key)
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
