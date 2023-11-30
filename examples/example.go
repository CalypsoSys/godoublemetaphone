package main

import (
	"fmt"

	"github.com/CalypsoSys/godoublemetaphone"
)

func main() {
	primary, secondary := godoublemetaphone.NewDoubleMetaphone("SMITH")
	fmt.Printf("Metaphones for SMITH: first: %v, second: %v\n", primary, secondary)

	primary, secondary = godoublemetaphone.NewDoubleMetaphone("SMYTHE")
	fmt.Printf("Metaphones for SMYTHE: first: %v, second: %v\n", primary, secondary)

	primary, secondary = godoublemetaphone.NewDoubleMetaphone("SCHMIDT")
	fmt.Printf("Metaphones for SCHMIDT: first: %v, second: %v\n", primary, secondary)
}
