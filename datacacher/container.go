package datacacher

type ICacheContainer[KeyType comparable, IdType comparable] interface {
	CacheGetValue(key KeyType, id IdType) (any, error)
	CacheSetValue(key KeyType, id IdType, value any) error
	CacheClearValue(key KeyType, id IdType) error
	CacheClearAll(key KeyType) error
}

type MapContainer[KeyType comparable, IdType comparable] struct {
	m map[KeyType]map[IdType]any
}

func NewMapContainer[KeyType comparable, IdType comparable]() *MapContainer[KeyType, IdType] {
	return &MapContainer[KeyType, IdType]{
		m: make(map[KeyType]map[IdType]any),
	}
}

func (c *MapContainer[KeyType, IdType]) CacheGetValue(key KeyType, id IdType) (any, error) {
	if c.m[key] == nil {
		return nil, nil
	}
	return c.m[key][id], nil
}

func (c *MapContainer[KeyType, IdType]) CacheSetValue(key KeyType, id IdType, value any) error {
	if c.m[key] == nil {
		c.m[key] = make(map[IdType]any)
	}
	c.m[key][id] = value
	return nil
}

func (c *MapContainer[KeyType, IdType]) CacheClearValue(key KeyType, id IdType) error {
	if c.m[key] == nil {
		return nil
	}
	delete(c.m[key], id)
	if len(c.m[key]) == 0 {
		delete(c.m, key)
	}
	return nil
}

func (c *MapContainer[KeyType, IdType]) CacheClearAll(key KeyType) error {
	delete(c.m, key)
	return nil
}
