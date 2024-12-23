package datacacher

import (
	"fmt"
	"testing"
)

type testContainer struct {
	*MapContainer[string, int]
}

func (c *testContainer) CacheGetValue(key string, id int) (any, error) {
	val, err := c.MapContainer.CacheGetValue(key, id)
	if err != nil {
		return nil, err
	}
	fmt.Printf("test container get value result, key: %v, id: %v, val: %v \n", key, id, val)
	return val, nil
}

func (c *testContainer) CacheSetValue(key string, id int, value any) error {
	fmt.Printf("test container set value, key: %v, id: %v, val: %v \n", key, id, value)
	err := c.MapContainer.CacheSetValue(key, id, value)
	if err != nil {
		return err
	}
	fmt.Printf("test container set value after, map: %v \n", c.m)
	return nil
}

func (c *testContainer) CacheClearValue(key string, id int) error {
	fmt.Printf("test container clear value, key: %v, id: %v \n", key, id)
	err := c.MapContainer.CacheClearValue(key, id)
	if err != nil {
		return err
	}
	fmt.Printf("test container clear value after, map: %v \n", c.m)
	return nil
}

const (
	CACHER_KEY_CALC_PI    = "calc_pi"
	CACHER_KEY_GET_NUMBER = "get_number"
)

type calcPi struct {
}

func (c *calcPi) Create(container *testContainer, param *Param) (any, error) {
	fmt.Println("calc pi, create")
	return 3.1415926, nil
}

type getNumber struct {
}

func (c *getNumber) Create(container *testContainer, param *Param) (any, error) {
	fmt.Println("get number, create")
	number := param.GetExtra("key").(int)
	return number + 100, nil
}

func prepare() (*testContainer, *Cacher[string, int, *testContainer]) {
	testC := &testContainer{
		MapContainer: NewMapContainer[string, int](),
	}
	cacher := NewCacher[*testContainer]()
	cacher.Register(CACHER_KEY_CALC_PI, NewBaseCachee(&calcPi{}))
	cacher.Register(CACHER_KEY_GET_NUMBER, NewBaseCachee(&getNumber{}))
	return testC, cacher
}

func TestBase(t *testing.T) {
	testC, cacher := prepare()
	val, err := cacher.Get(testC, CACHER_KEY_CALC_PI)
	fmt.Println("CALC PI 1", val, err)
	val, err = cacher.Get(testC, CACHER_KEY_CALC_PI)
	fmt.Println("CALC PI 2", val, err)

	fmt.Println("---------------")

	val, err = cacher.Get(testC, CACHER_KEY_GET_NUMBER, NewParam().SetExtra("key", 100))
	fmt.Println("GET NUMBER 1", val, err)
	val, err = cacher.Get(testC, CACHER_KEY_GET_NUMBER, NewParam().SetExtra("key", 100))
	fmt.Println("GET NUMBER 2", val, err)

	val, err = cacher.Get(testC, CACHER_KEY_GET_NUMBER,
		NewParam().SetId(10086).SetExtra("key", 1000))
	fmt.Println("GET NUMBER 3", val, err)

}

func TestInvalidIdInParam(t *testing.T) {
	testC, cacher := prepare()
	// should panic here
	cacher.Get(testC, CACHER_KEY_GET_NUMBER,
		NewParam().SetId(uint(10010)).SetExtra("key", 2000))
}

func TestClearValue(t *testing.T) {
	testC, cacher := prepare()

	val, err := cacher.Get(testC, CACHER_KEY_GET_NUMBER,
		NewParam().SetId(10086).SetExtra("key", 100))
	fmt.Println("1. GET NUMBER", val, err)
	val, err = cacher.Get(testC, CACHER_KEY_GET_NUMBER,
		NewParam().SetId(10086).SetExtra("key", 100))
	fmt.Println("2. GET NUMBER", val, err)

	cacher.Clear(testC, CACHER_KEY_GET_NUMBER)
	val, err = cacher.Try(testC, CACHER_KEY_GET_NUMBER,
		NewParam().SetId(10086).SetExtra("key", 100))
	fmt.Println("3. TRY NUMBER", val, err)

	cacher.Clear(testC, CACHER_KEY_GET_NUMBER,
		NewParam().SetId(10086))
	val, err = cacher.Try(testC, CACHER_KEY_GET_NUMBER,
		NewParam().SetId(10086).SetExtra("key", 100))
	fmt.Println("4. TRY NUMBER", val, err)
}
