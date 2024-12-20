package datacacher

type ICachee interface {
	Key() any
	SetKey(any)

	Get(ICacheContainer, *Param) (any, error)
	Set(ICacheContainer, any, *Param) error

	Create(ICacheContainer, *Param) (any, error)

	Clear(ICacheContainer, *Param) error
	ClearAll()
}

type BaseCachee struct {
	ICalculator
}

func NewBaseCachee(calculator ICalculator) *BaseCachee {
	return &BaseCachee{
		ICalculator: calculator,
	}
}

func (c *BaseCachee) Key() any {
	return c
}

func (c *BaseCachee) SetKey(key any) {
}

func (c *BaseCachee) Get(container ICacheContainer, param *Param) (any, error) {
	key := c.Key()
	var id any = 0
	if param != nil {
		id = param.Id()
	}
	return container.CacheGetValue(key, id)
}

func (c *BaseCachee) Set(container ICacheContainer, val any, param *Param) error {
	key := c.Key()
	var id any = 0
	if param != nil {
		id = param.Id()
	}
	return container.CacheSetValue(key, id, val)
}

func (c *BaseCachee) Clear(container ICacheContainer, param *Param) error {
	key := c.Key()
	var id any = 0
	if param != nil {
		id = param.Id()
	}
	return container.CacheClearValue(key, id)
}

func (c *BaseCachee) ClearAll() {}

type ICalculator interface {
	Create(ICacheContainer, *Param) (any, error)
}

type ICacheContainer interface {
	CacheGetValue(key any, id any) (any, error)
	CacheSetValue(key any, id any, value any) error
	CacheClearValue(key any, id any) error
}
