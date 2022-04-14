[![Contributor Covenant](https://img.shields.io/badge/Contributor%20Covenant-2.0-4baaaa.svg)](CODE_OF_CONDUCT.md)
[![Build](https://github.com/iofic/genmap/actions/workflows/build.yaml/badge.svg)](https://github.com/iofic/genmap/actions/workflows/build.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/iofic/genmap.svg)](https://pkg.go.dev/github.com/iofic/genmap)

# GenMap
GenMap provides a generic map, able to store any type, and retrieve values with type safety.

This is done by associating types with key types.

## Example

### Basic example

```go
package main

import (
	"fmt"
	"github.com/iofic/genmap"
)

const (
	id1 int = iota
	id2
)

type Foo struct {
	a int
}

type Bar struct {
	b int
}

func main() {
	// In this example we'll be storing values of types *Foo and *Bar, against int type keys.
	
	// Create keys for the objects we'll be storing in our map.
	// We need to declare both the value types, and the key type on keys. This allows for type-safety when fetching
	// values, as each key is 'tied' to a type.
	// key1 can only store and load values of type *Foo
	key1 := genmap.NewKey[*Foo, int](id1)
	// key2 can only store and load values of type *Bar
	key2 := genmap.NewKey[*Bar, int](id2)
	
	// Create a genmap.Map, we only need to declare the key type, as Maps can store any type.
	m := genmap.New[int]()
	
	// Create some values to store.
	foo := &Foo{a: 1}
	bar := &Bar{b: 2}
	
	genmap.Store[*Foo, int](m, key1, foo)
	genmap.Store[*Bar, int](m, key2, bar)
	
	// We can then use the keys to load those values.
	// No need to type cast as we specify the types in the generic function Load.
	fmt.Println(genmap.Load[*Foo, int](m, key1).a) // 1
	fmt.Println(genmap.Load[*Bar, int](m, key2).b) // 2
}
```

### Using 'string' functions
GenMap offers 'string' varieties of top level functions. This allows for calling code to use string keys, which is very
common, without having to specify the string type during initialisation.

```go
package main

import (
	"fmt"
	"github.com/iofic/genmap"
)

func main() {
	// This example uses the 'string' version of map and keys. This allows the calling code to worry less about key and
	// map types, and only specify value store and load types.

	// Create two keys, both of these store and load int type values, however use unique IDs to differentiate.
	primaryIDKey := genmap.NewStringKey[int]("primary")
	secondaryIDKey := genmap.NewStringKey[int]("secondary")

	stringMap := genmap.NewStringMap()

	// Set the initial values against the keys
	genmap.StoreS[int](stringMap, primaryIDKey, 1)
	genmap.StoreS[int](stringMap, secondaryIDKey, 2)

	// Fetch values via keys.
	fmt.Println(genmap.LoadS[int](stringMap, primaryIDKey))
	// Output: 1 true
	fmt.Println(genmap.LoadS[int](stringMap, secondaryIDKey))
	// Output: 2 true

	// Override the first key.
	genmap.StoreS[int](stringMap, primaryIDKey, 3)

	// Fetch new values, only the first key value should have been changed.
	fmt.Println(genmap.LoadS[int](stringMap, primaryIDKey))
	// Output: 3 true
	fmt.Println(genmap.LoadS[int](stringMap, secondaryIDKey))
	// Output: 2 true

	// Delete the second key.
	genmap.DeleteS[int](stringMap, secondaryIDKey)

	// Fetch a missing key.
	fmt.Println(genmap.LoadS[int](stringMap, primaryIDKey))
	// Output: 3 true
	fmt.Println(genmap.LoadS[int](stringMap, secondaryIDKey))
	// Output: 0 false
}
```
