package datacacher

import "time"

type TimeoutNode struct {
	val any
	tm  time.Time
}

type TimeoutCachee[KeyType comparable, IdType comparable, ContainerType ICacheContainer[KeyType, IdType]] struct {
	*BaseCachee[KeyType, IdType, ContainerType]
	timeout time.Duration
}

func WithTimeout[KeyType comparable, IdType comparable, ContainerType ICacheContainer[KeyType, IdType]](
	base *BaseCachee[KeyType, IdType, ContainerType], timeout time.Duration) *TimeoutCachee[KeyType, IdType, ContainerType] {
	return &TimeoutCachee[KeyType, IdType, ContainerType]{base, timeout}
}

func (c *TimeoutCachee[KeyType, IdType, ContainerType]) Get(container ContainerType, param *Param) (any, error) {
	nv, err := c.BaseCachee.Get(container, param)
	if err != nil {
		return nil, err
	}
	if nv == nil {
		return nil, nil
	}

	node := nv.(*TimeoutNode)
	if time.Now().After(node.tm.Add(c.timeout)) {
		return nil, nil
	}
	return node.val, nil
}

func (c *TimeoutCachee[KeyType, IdType, ContainerType]) Set(container ContainerType, val any, param *Param) error {
	return c.BaseCachee.Set(container, &TimeoutNode{val, time.Now()}, param)
}
