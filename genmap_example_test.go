package genmap_test

import (
	"fmt"

	"github.com/iofic/genmap"
)

func ExampleMap() {
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
	fmt.Println(genmap.LoadS[int](stringMap, secondaryIDKey))

	// Override the first key.
	genmap.StoreS[int](stringMap, primaryIDKey, 3)

	// Fetch new values, only the first key value should have been changed.
	fmt.Println(genmap.LoadS[int](stringMap, primaryIDKey))
	fmt.Println(genmap.LoadS[int](stringMap, secondaryIDKey))

	// Delete the second key.
	genmap.DeleteS[int](stringMap, secondaryIDKey)

	// Fetch a missing key.
	fmt.Println(genmap.LoadS[int](stringMap, primaryIDKey))
	fmt.Println(genmap.LoadS[int](stringMap, secondaryIDKey))

	// Output:
	// 1 true
	// 2 true
	// 3 true
	// 2 true
	// 3 true
	// 0 false
}
