package datacacher

import "testing"

func TestAnyKey(t *testing.T) {
	m := make(map[any]any)
	m["test"] = 1
	t.Log(m["test"])
	t.Log(m["test1"])
}
