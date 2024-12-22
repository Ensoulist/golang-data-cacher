package datacacher

import "fmt"

type ICachee[KeyType comparable, IdType comparable, ContainerType ICacheContainer[KeyType, IdType]] interface {
	Key() KeyType
	SetKey(KeyType)

	Get(ContainerType, *Param) (any, error)
	Set(ContainerType, any, *Param) error

	Create(ContainerType, *Param) (any, error)

	Clear(ContainerType, *Param) error
	ClearAll()
}

type BaseCachee[KeyType comparable, IdType comparable, ContainerType ICacheContainer[KeyType, IdType]] struct {
	ICalculator[ContainerType, KeyType, IdType]
	key KeyType
}

func NewBaseCachee[KeyType comparable, IdType comparable, ContainerType ICacheContainer[KeyType, IdType]](calculator ICalculator[ContainerType, KeyType, IdType]) *BaseCachee[KeyType, IdType, ContainerType] {
	return &BaseCachee[KeyType, IdType, ContainerType]{
		ICalculator: calculator,
	}
}

func (c *BaseCachee[KeyType, IdType, ContainerType]) Key() KeyType {
	return c.key
}

func (c *BaseCachee[KeyType, IdType, ContainerType]) SetKey(key KeyType) {
	c.key = key
}

func (c *BaseCachee[KeyType, IdType, ContainerType]) Get(container ContainerType, param *Param) (any, error) {
	key := c.Key()
	var id IdType
	if param != nil {
		id = getIdFromParam[IdType](param)
	}
	return container.CacheGetValue(key, id)
}

func (c *BaseCachee[KeyType, IdType, ContainerType]) Set(container ContainerType, val any, param *Param) error {
	key := c.Key()
	var id IdType
	if param != nil {
		id = getIdFromParam[IdType](param)
	}
	return container.CacheSetValue(key, id, val)
}

func (c *BaseCachee[KeyType, IdType, ContainerType]) Clear(container ContainerType, param *Param) error {
	key := c.Key()
	var id IdType
	if param != nil {
		id = getIdFromParam[IdType](param)
	}
	return container.CacheClearValue(key, id)
}

func (c *BaseCachee[KeyType, IdType, ContainerType]) ClearAll() {}

func getIdFromParam[IdType comparable](param *Param) IdType {
	var id IdType
	if param.Id() == nil {
		return id
	}
	var ok bool
	id, ok = param.Id().(IdType)
	if !ok {
		panic(fmt.Sprintf("invalid id type, should be castable to %T", id))
	}
	return id
}
