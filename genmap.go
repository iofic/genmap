package genmap

// Map is a simple type-safe map, providing Store and Load functionality. Map is not thread safe.
// The Map generic type is the key type.
type Map[K comparable] interface {
	getMap() map[K]any
}

// genMap is a simple type-safe map, providing Store and Load functionality. genMap is not thread safe.
// The genMap generic type is the key type.
type genMap[K comparable] struct {
	m map[K]any
}

func (gm *genMap[K]) getMap() map[K]any {
	return gm.m
}

// New creates a new genMap. Maps returned from New is ready to use.
// The New generic type is the key type.
func New[K comparable]() Map[K] {
	m := genMap[K]{
		m: make(map[K]any),
	}
	return &m
}

// NewStringMap returns a map of string type. This is provided as a more concise Map implementation than using New.
func NewStringMap() Map[string] {
	return New[string]()
}

// Store stores a value against a key on the genMap provided. This will override any value stored against this key.
func Store[T any, K comparable](m Map[K], key Key[T, K], value T) {
	m.getMap()[key.Key()] = value
}

// StoreS is a shortcut for Store, for StringMaps.
func StoreS[T any](m Map[string], key Key[T, string], value T) {
	Store[T, string](m, key, value)
}

// Load will fetch a value from the genMap. If found, ok will be true, otherwise false.
func Load[T any, K comparable](m Map[K], key Key[T, K]) (value T, ok bool) {
	val, ok := m.getMap()[key.Key()].(T)
	return val, ok
}

// LoadS is a shortcut for Load, for StringMaps.
func LoadS[T any](m Map[string], key Key[T, string]) (value T, ok bool) {
	return Load[T, string](m, key)
}

// Delete will delete an entry from the Map which matches the Key provided. If there is no entry matching the Key,
// Delete is a no-op function.
func Delete[T any, K comparable](m Map[K], key Key[T, K]) {
	delete(m.getMap(), key.Key())
}

// DeleteS is a shortcut for Delete, for StringMaps.
func DeleteS[T any](m Map[string], key Key[T, string]) {
	Delete[T, string](m, key)
}
