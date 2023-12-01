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
	dm := godoublemetaphone.NewDoubleMetaphone("Peace")
	fmt.Printf("Metaphones for Boxers: primary: %s, alternate: %v\n", dm.PrimaryKey(), safeString(dm.AlternateKey()))

	dm = godoublemetaphone.NewDoubleMetaphone("Piece")
	fmt.Printf("Metaphones for Dancers: primary: %s, alternate: %v\n", dm.PrimaryKey(), safeString(dm.AlternateKey()))

	dm = godoublemetaphone.NewDoubleMetaphone("Pace")
	fmt.Printf("Metaphones for Prancers: primary: %s, alternate: %v\n", dm.PrimaryKey(), safeString(dm.AlternateKey()))

	sdm := godoublemetaphone.NewShortDoubleMetaphone("Peace")
	fmt.Printf("ShortMetaphones for Boxers: primary: %d, alternate: %d\n", sdm.PrimaryShortKey(), sdm.AlternateShortKey())

	sdm = godoublemetaphone.NewShortDoubleMetaphone("Piece")
	fmt.Printf("ShortMetaphones for Dancers: primary: %d, alternate: %d\n", sdm.PrimaryShortKey(), sdm.AlternateShortKey())

	sdm = godoublemetaphone.NewShortDoubleMetaphone("Pace")
	fmt.Printf("ShortMetaphones for Prancers: primary: %d, alternate: %d\n", sdm.PrimaryShortKey(), sdm.AlternateShortKey())
}
