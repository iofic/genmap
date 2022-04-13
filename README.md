[![Contributor Covenant](https://img.shields.io/badge/Contributor%20Covenant-2.0-4baaaa.svg)](CODE_OF_CONDUCT.md)

[![Build](https://github.com/iofic/genmap/actions/workflows/build.yml/badge.svg)](https://github.com/iofic/genmap/actions/workflows/build.yml)

# GenMap
GenMap provides a generic map, able to store any type, and retrieve values with type safety.

This is done by associating types with key types.

## Example
```go
package example

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
