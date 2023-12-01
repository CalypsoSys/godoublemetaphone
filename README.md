# godoublemmetaphone
Golang implementation of Lawrence Phillips' Double Metaphone phonetic matching  algorithm, published in C/C++ Users Journal, June, 2000.

Golang implementation of Lawrence's proposed optimization, whereby four-character metaphone keys
are represented as four nibbles in an unsigned short.

Metaphone is a phonetic algorithm, published by Lawrence Philips in 1990, for indexing words by their English pronunciation. It fundamentally improves on the Soundex algorithm by using information about variations and inconsistencies in English spelling and pronunciation to produce a more accurate encoding, which does a better job of matching words and names which sound similar. As with Soundex, similar-sounding words should share the same keys. Metaphone is available as a built-in operator in a number of systems.

It is called "Double" because it can return both a primary and a secondary code for a string; this accounts for some ambiguous cases as well as for multiple variants of surnames with common ancestry. For example, encoding the name "Smith" yields a primary code of SM0 and a secondary code of XMT, while the name "Schmidt" yields a primary code of XMT and a secondary code of SMT—both have XMT in common.

Double Metaphone tries to account for myriad irregularities in English of Slavic, Germanic, Celtic, Greek, French, Italian, Spanish, Chinese, and other origins. Thus it uses a much more complex ruleset for coding than its predecessor; for example, it tests for approximately 100 different contexts of the use of the letter C alone.

See
https://en.wikipedia.org/wiki/Metaphone#Double_Metaphone

# Usage
Call the `godoublemetaphone.NewDoubleMetaphone` or `godoublemetaphone.NewShortDoubleMetaphone` function(s) to retrieve a interface to get the primary and the alternate metaphones. Both function takes as a parameter a string and returns a interface to get the metaphones (in either string or uint16 format).

Example:
```
	dm := godoublemetaphone.NewDoubleMetaphone("Peace")
	fmt.Printf("Metaphones for Boxers: primary: %s, alternate: %v\n", dm.PrimaryKey(), safeString(dm.AlternateKey()))

	sdm = godoublemetaphone.NewShortDoubleMetaphone("Piece")
	fmt.Printf("ShortMetaphones for Dancers: primary: %d, alternate: %d\n", sdm.PrimaryShortKey(), sdm.AlternateShortKey())
```

