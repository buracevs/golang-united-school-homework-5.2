package cache

import "time"

type Cache struct {
	values map[string]item
}

type item struct {
	data string
	ttl  time.Time
}

func NewCache() Cache {
	return Cache{
		values: make(map[string]item),
	}
}

func (c Cache) Get(key string) (string, bool) {
	value, ok := c.values[key]
	if ok {
		if value.ttl.IsZero() {
			return value.data, ok
		}
		if time.Now().Unix() < value.ttl.Unix() {
			return value.data, ok
		}
	}

	return "", false
}

func (c Cache) Put(key, value string) {
	c.values[key] = item{data: value}
}

func (c Cache) Keys() []string {
	var result []string
	for key, value := range c.values {
		if value.ttl.IsZero() {
			result = append(result, key)
		}
		if time.Now().Unix() < value.ttl.Unix() {
			result = append(result, key)
		}
	}
	return result
}

func (c Cache) PutTill(key, value string, deadline time.Time) {
	c.values[key] = item{data: value, ttl: deadline}
}
