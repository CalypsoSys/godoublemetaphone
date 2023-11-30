package main

import (
	"fmt"

	"github.com/CalypsoSys/godoublemetaphone/pkg/godoublemetaphone"
)

func safeString(str *string) string {
	if str != nil {
		return *str
	}

	return "<nil>"
}

func main() {
	dm := godoublemetaphone.NewDoubleMetaphone("SMITH")
	fmt.Printf("Metaphones for SMITH: first: %s, second: %v\n", dm.PrimaryKey(), safeString(dm.AlternateKey()))

	dm = godoublemetaphone.NewDoubleMetaphone("SMYTHE")
	fmt.Printf("Metaphones for SMYTHE: first: %s, second: %v\n", dm.PrimaryKey(), safeString(dm.AlternateKey()))

	dm = godoublemetaphone.NewDoubleMetaphone("SCHMIDT")
	fmt.Printf("Metaphones for SCHMIDT: first: %s, second: %v\n", dm.PrimaryKey(), safeString(dm.AlternateKey()))
}
