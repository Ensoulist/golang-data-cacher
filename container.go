package datacacher

type ICacheContainer interface {
	CacheGetValue(key any, id any) (any, error)
	CacheSetValue(key any, id any, value any) error
	CacheClearValue(key any, id any) error
}

type MapContainer struct {
	m map[any]map[any]any
}

func NewMapContainer() *MapContainer {
	return &MapContainer{
		m: make(map[any]map[any]any),
	}
}

func (c *MapContainer) CacheGetValue(key any, id any) (any, error) {
	if c.m[key] == nil {
		return nil, nil
	}
	return c.m[key][id], nil
}

func (c *MapContainer) CacheSetValue(key any, id any, value any) error {
	if c.m[key] == nil {
		c.m[key] = make(map[any]any)
	}
	c.m[key][id] = value
	return nil
}

func (c *MapContainer) CacheClearValue(key any, id any) error {
	if c.m[key] == nil {
		return nil
	}
	delete(c.m[key], id)
	if len(c.m[key]) == 0 {
		delete(c.m, key)
	}
	return nil
}
