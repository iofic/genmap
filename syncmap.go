package genmap

import "sync"

type SyncMap interface {
	getMap() *sync.Map
}

type genSyncMap struct {
	m *sync.Map
}

func (gsm *genSyncMap) getMap() *sync.Map {
	return gsm.m
}

// NewSyncMap returns a type-safe generic map built upon a sync.Map.
func NewSyncMap() SyncMap {
	var m sync.Map
	return &genSyncMap{
		m: &m,
	}
}

// DeleteSync deletes the value for a key.
func DeleteSync[T any, K comparable](m SyncMap, key Key[T, K]) {
	m.getMap().Delete(key.Key())
}

// StoreSync sets the value for a key.
func StoreSync[T any, K comparable](m SyncMap, key Key[T, K], value T) {
	m.getMap().Store(key.Key(), value)
}

// LoadSync returns the value stored in the map for a key, or nil if no value is present. The ok result indicates
// whether value was found in the map.
func LoadSync[T any, K comparable](m SyncMap, key Key[T, K]) (value T, ok bool) {
	val, ok := m.getMap().Load(key.Key())
	if val != nil {
		value = val.(T)
	}
	return value, ok
}

// LoadAndDelete deletes the value for a key, returning the previous value if any. The loaded result reports whether the
// key was present.
func LoadAndDelete[T any, K comparable](m SyncMap, key Key[T, K]) (value T, loaded bool) {
	val, loaded := m.getMap().LoadAndDelete(key.Key())
	if val != nil {
		value = val.(T)
	}
	return value, loaded
}

// LoadOrStore returns the existing value for the key if present. Otherwise, it stores and returns the given value. The
// loaded result is true if the value was loaded, false if stored.
func LoadOrStore[T any, K comparable](m SyncMap, key Key[T, K], value T) (actual T, loaded bool) {
	val, loaded := m.getMap().LoadOrStore(key.Key(), value)
	if val != nil {
		value = val.(T)
	}
	return value, loaded
}
