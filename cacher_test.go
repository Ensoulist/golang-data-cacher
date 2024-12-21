package datacacher

import (
	"fmt"
	"testing"
)

type testContainer struct {
	*MapContainer
}

func (c *testContainer) CacheGetValue(key, id any) (any, error) {
	val, err := c.MapContainer.CacheGetValue(key, id)
	if err != nil {
		return nil, err
	}
	fmt.Println("test container get value result", key, id, val)
	return val, nil
}

func (c *testContainer) CacheSetValue(key, id, value any) error {
	fmt.Println("test container set value", key, id, value)
	err := c.MapContainer.CacheSetValue(key, id, value)
	if err != nil {
		return err
	}
	fmt.Println("test container set value after", c.m)
	return nil
}

func (c *testContainer) CacheClearValue(key, id any) error {
	fmt.Println("test container clear value", key, id)
	err := c.MapContainer.CacheClearValue(key, id)
	if err != nil {
		return err
	}
	fmt.Println("test container clear value after", c.m)
	return nil
}

const (
	CACHER_KEY_CALC_PI    = "calc_pi"
	CACHER_KEY_GET_NUMBER = "get_number"
)

type calcPi struct {
}

func (c *calcPi) Create(container ICacheContainer, param *Param) (any, error) {
	return 3.1415926, nil
}

type getNumber struct {
}

func (c *getNumber) Create(container ICacheContainer, param *Param) (any, error) {
	number := param.GetExtra("key").(int64)
	return number + 100, nil
}

func TestBase(t *testing.T) {
	testC := &testContainer{
		MapContainer: NewMapContainer(),
	}
	cacher := NewCacher()
	cacher.Register(CACHER_KEY_CALC_PI, NewBaseCachee(&calcPi{}))
	cacher.Register(CACHER_KEY_GET_NUMBER, NewBaseCachee(&getNumber{}))

	val, err := cacher.Get(testC, CACHER_KEY_CALC_PI)
	fmt.Println("CALC PI 1", val, err)
	val, err = cacher.Get(testC, CACHER_KEY_CALC_PI)
	fmt.Println("CALC PI 2", val, err)
}
