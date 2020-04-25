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
	cacheSAttrKey     = "cache:sysattr:%d"
	cacheSAttrListKey = "cache:sysattr:list:%s"
)

//SAttrCache - user cache
type SAttrCache struct {
	Rds *redis.Client
}

func NewSAttrCache() *SAttrCache {
	return &SAttrCache{
		Rds: driver.Rds,
	}
}

func (u *SAttrCache) Get(id uint64) (m *model.SAttr) {
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

func (u *SAttrCache) Set(m *model.SAttr) bool {
	k := fmt.Sprintf(cacheSAttrKey, m.ID)
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

func (u *SAttrCache) GetAll(key string) (m []*model.SAttr) {
	k := fmt.Sprintf(cacheSAttrListKey, key)
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

func (u *SAttrCache) SetAll(key string, m []*model.SAttr) bool {
	k := fmt.Sprintf(cacheSAttrListKey, key)
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
