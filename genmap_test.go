package genmap_test

import (
	"github.com/iofic/genmap"
	"testing"
)

func TestMap(t *testing.T) {
	const (
		_ int = iota
		val1
		val2
	)

	m := genmap.New[string]()

	ik := genmap.NewKey[int, string]("int key")

	genmap.Store[int, string](m, ik, val1)

	v, ok := genmap.Load[int, string](m, ik)
	if v != val1 || !ok {
		t.Errorf("Load() = %v %t, want %v %t", v, ok, val1, true)
	}

	genmap.Store[int, string](m, ik, val2)
	v, ok = genmap.Load[int, string](m, ik)
	if v != val2 || !ok {
		t.Errorf("Load() = %v %t, want %v %t", v, ok, val2, true)
	}

	genmap.Delete[int, string](m, ik)
	v, ok = genmap.Load[int, string](m, ik)
	if v != 0 || ok {
		t.Errorf("Load() = %v %t, want %v %t", v, ok, 0, false)
	}
}

func FuzzMap(f *testing.F) {
	keys := []struct {
		key   string
		value int
	}{
		{
			key:   "a",
			value: 1,
		},
	}

	m := genmap.NewStringMap()

	for _, k := range keys {
		f.Add(k.key, k.value)
	}

	f.Fuzz(func(t *testing.T, key string, value int) {
		ik := genmap.NewStringKey[int](key)
		genmap.StoreS(m, ik, value)
		_, _ = genmap.LoadS(m, ik)
	})
}
