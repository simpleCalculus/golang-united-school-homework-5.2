package cache

import (
	"time"
)

type myTimeValue struct {
	value    string
	deadline time.Time
}

type Cache struct {
	simpleMap map[string]string
	timeMap   map[string]myTimeValue
}

func NewCache() Cache {
	sm := make(map[string]string)
	tm := make(map[string]myTimeValue)
	return Cache{sm, tm}
}

func (c Cache) Put(key, value string) {
	c.simpleMap[key] = value
}

func (c Cache) PutTill(key, value string, deadline time.Time) {
	c.timeMap[key] = myTimeValue{value, deadline}
}

func (c Cache) Get(key string) (string, bool) {
	value, ok := c.simpleMap[key]
	if ok {
		return value, ok
	}

	timeValue, ok := c.timeMap[key]
	if !ok {
		return "", ok
	}
	if timeValue.deadline.After(time.Now()) {
		return timeValue.value, ok
	}
	return "", false
}

func (c Cache) Keys() []string {
	keys := make([]string, 0)
	for k := range c.simpleMap {
		keys = append(keys, k)
	}
	for k, v := range c.timeMap {
		if v.deadline.After(time.Now()) {
			keys = append(keys, k)
		}
	}
	return keys
}
