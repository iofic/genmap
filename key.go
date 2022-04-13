package genmap

import (
	"fmt"
)

// Key associates a type with a key, ensuring type safety when setting and fetching values of this type.
type Key[T any, K comparable] interface {
	// Key returns a serialisation of the key type. This should be unique for the type, and can be a combination of an
	// instance.
	// e.g. myStruct_objectName
	Key() K
}

type mapKey[T any, K comparable] struct {
	key K
}

// Key returns the serialisation of this key.
func (sk mapKey[T, K]) Key() K {
	return sk.key
}

// NewKey creates a new Key. The id provided should be unique to the object.
func NewKey[T any, K comparable](id K) Key[T, K] {
	return mapKey[T, K]{
		key: id,
	}
}

// NewStringKey creates a new Key with a string id. This can be used with the string variety of top level functions.
func NewStringKey[T any](id string) Key[T, string] {
	var t T
	keyName := fmt.Sprintf("%s_%T", id, t)
	return NewKey[T, string](keyName)
}
