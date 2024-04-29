package main

import (
	"fmt"
)

type Cache interface {
	Get(k string) (string, bool)
	Set(k, v string)
}

var _ Cache = (*cacheImpl)(nil)

func newCacheImpl() *cacheImpl {
	return &cacheImpl{
		cacheMap: map[string]string{},
	}
}

type cacheImpl struct {
	cacheMap map[string]string
}

func (c *cacheImpl) Get(k string) (string, bool) {
	s, ok := c.cacheMap[k]
	return s, ok
}

func (c *cacheImpl) Set(k, v string) {
	c.cacheMap[k] = v
}

// странно создавать объект для хранения данных и в конструкторе всегда хардкодить какие-то значения
// поэтому вынесла заполнение данными dbImpl через метод (d *dbImpl) Set(k, v string)
func newDbImpl(cache Cache) *dbImpl {
	return &dbImpl{cache: cache, dbs: map[string]string{}}
}

type dbImpl struct {
	cache Cache
	dbs   map[string]string
}

func (d *dbImpl) Get(k string) (string, bool) {
	v, ok := d.cache.Get(k)
	if ok {
		return fmt.Sprintf("answer from cache: key: %s, val: %s", k, v), ok
	}

	v, ok = d.dbs[k]
	return fmt.Sprintf("answer from dbs: key: %s, val: %s", k, v), ok
}

func (d *dbImpl) Set(k, v string) {
	d.dbs[k] = v
	d.cache.Set(k, v)
}

func main() {
	c := newCacheImpl()
	db := newDbImpl(c)
	db.Set("hello", "world")
	db.Set("test", "test")

	fmt.Println(db.Get("hello"))
	fmt.Println(db.Get("test"))

	fmt.Println(db.Get("unknown"))
}
