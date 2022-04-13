package genmap_test

import (
	"fmt"
	"github.com/iofic/genmap"
)

type MyStruct struct{}

func ExampleNewStringKey() {
	kFloat64 := genmap.NewStringKey[float64]("foobar")
	// The key serialisation of a StringKey is a combination of the string id provided, and the type name.
	fmt.Println(kFloat64.Key())

	kMyStruct := genmap.NewStringKey[MyStruct]("yazbaz")
	// The key serialisation of a StringKey with a custom type will contain the package name.
	fmt.Println(kMyStruct.Key())

	// Output:
	// foobar_float64
	// yazbaz_genmap_test.MyStruct
}
