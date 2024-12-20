package datacacher

import (
	"errors"
	"fmt"
)

type Cacher struct {
	registed map[any]ICachee
}

func NewCacher() *Cacher {
	return &Cacher{
		registed: make(map[any]ICachee),
	}
}

func (c *Cacher) Register(key any, cachee ICachee) {
	if _, ok := c.registed[key]; ok {
		panic(fmt.Sprintf("cacher register, key conflicted: %v", key))
	}
	cachee.SetKey(key)
	c.registed[key] = cachee
}

func (c *Cacher) Get(container ICacheContainer, key any, param ...*Param) (any, error) {
	var p *Param
	if len(param) > 0 {
		p = param[0]
	}

	val, err := c.Try(container, key, p)
	if err != nil {
		return nil, err
	}
	if val != nil {
		return val, nil
	}
	return c.Update(container, key, p)
}

func (c *Cacher) Try(container ICacheContainer, key any, param ...*Param) (any, error) {
	var p *Param
	if len(param) > 0 {
		p = param[0]
	}
	cachee, ok := c.registed[key]
	if !ok {
		return nil, errors.New("try, cacher not found")
	}
	return cachee.Get(container, p)
}

func (c *Cacher) Update(container ICacheContainer, key any, param ...*Param) (any, error) {
	var p *Param
	if len(param) > 0 {
		p = param[0]
	}

	cachee, ok := c.registed[key]
	if !ok {
		return nil, errors.New("update, cacher not found")
	}
	val, err := cachee.Create(container, p)
	if err != nil {
		return nil, err
	}
	err = cachee.Set(container, val, p)
	if err != nil {
		return nil, err
	}
	return val, nil
}

func (c *Cacher) Clear(container ICacheContainer, key any, param ...*Param) (any, error) {
	var p *Param
	if len(param) > 0 {
		p = param[0]
	}
	oldVal, err := c.Try(container, key, p)
	if err != nil {
		return nil, err
	}
	if oldVal == nil {
		return nil, nil
	}
	cachee := c.registed[key]
	err = cachee.Clear(container, p)
	if err != nil {
		return nil, err
	}
	return oldVal, nil
}

func (c *Cacher) ClearAll() {
}
