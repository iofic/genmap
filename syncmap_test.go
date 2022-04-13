package genmap_test

import (
	"github.com/iofic/genmap"
	"testing"
)

func TestSyncMap(t *testing.T) {
	const (
		_ int = iota
		val1
		val2
	)

	m := genmap.NewSyncMap()

	ik := genmap.NewKey[int, string]("int key")

	genmap.StoreSync[int, string](m, ik, val1)

	v, ok := genmap.LoadSync[int, string](m, ik)
	if v != val1 || !ok {
		t.Errorf("Load() = %v %t, want %v %t", v, ok, val1, true)
	}

	genmap.StoreSync[int, string](m, ik, val2)
	v, ok = genmap.LoadSync[int, string](m, ik)
	if v != val2 || !ok {
		t.Errorf("Load() = %v %t, want %v %t", v, ok, val2, true)
	}

	genmap.DeleteSync[int, string](m, ik)
	v, ok = genmap.LoadSync[int, string](m, ik)
	if v != 0 || ok {
		t.Errorf("Load() = %v %t, want %v %t", v, ok, 0, false)
	}
}

func TestLoadAndDelete(t *testing.T) {
	const (
		_ int = iota
		val1
	)

	m := genmap.NewSyncMap()

	ik := genmap.NewKey[int, string]("int key")

	genmap.StoreSync[int, string](m, ik, val1)

	got, loaded := genmap.LoadAndDelete(m, ik)
	if got != val1 || !loaded {
		t.Errorf("LoadAndDelete() = %v %t, want %v %t", got, loaded, val1, true)
	}

	got1, ok := genmap.LoadSync(m, ik)
	if got1 != 0 || ok {
		t.Errorf("LoadSync = %v %t, want %v %t", got1, ok, 0, false)
	}
}

func TestLoadOrStore(t *testing.T) {
	const (
		_ int = iota
		val1
		val2
	)

	m := genmap.NewSyncMap()

	ik := genmap.NewKey[int, string]("int key")

	got, loaded := genmap.LoadOrStore(m, ik, val1)
	if got != val1 || loaded {
		t.Errorf("LoadOrStore = %v %t, want %v %t", got, loaded, val1, false)
	}

	got1, loaded1 := genmap.LoadOrStore(m, ik, val2)
	if got1 != val1 || !loaded1 {
		t.Errorf("LoadOrStore = %v %t, want %v %t", got1, loaded1, val1, true)
	}
}
