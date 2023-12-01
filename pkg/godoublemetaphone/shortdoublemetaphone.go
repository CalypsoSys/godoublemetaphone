package godoublemetaphone

/**
 * shortdoublemetaphone.go
 *
 * An implemenatation of Lawrence Phillips' Double Metaphone phonetic matching
 * algorithm, published in C/C++ Users Journal, June, 2000.  This implementation
 * implements Lawrence's proposed optimization, whereby four-character metaphone keys
 * are represented as four nibbles in an unsigned short.  This dramatically improves
 * storage and search efficiency.
 *
 * This implementation is based on the C# implementation that was written
 * by Adam J. Nelson (anelson@nullpointer.net). That was based on the C++ template implementation, also by Adam Nelson.
 * For the latest version of this implementation, implementations
 * in other languages, and links to articles I've written on the use of my various
 * Double Metaphone implementations, see:
 * http;//www.nullpointer.net/anelson/
 */

const (
	METAPHONE_KEY_LENGTH = 6 //The length of the metaphone keys produced.  4 is sweet spot
)

type ShortDoubleMetaphone interface {
	PrimaryShortKey() uint16
	AlternateShortKey() uint16
}

const (
	//Constants representing the characters in a metaphone key
	METAPHONE_A     uint16 = 0x01
	METAPHONE_F     uint16 = 0x02
	METAPHONE_FX    uint16 = ((METAPHONE_F << 4) | METAPHONE_X)
	METAPHONE_H     uint16 = 0x03
	METAPHONE_J     uint16 = 0x04
	METAPHONE_K     uint16 = 0x05
	METAPHONE_KL    uint16 = ((METAPHONE_K << 4) | METAPHONE_L)
	METAPHONE_KN    uint16 = ((METAPHONE_K << 4) | METAPHONE_N)
	METAPHONE_KS    uint16 = ((METAPHONE_K << 4) | METAPHONE_S)
	METAPHONE_L     uint16 = 0x06
	METAPHONE_M     uint16 = 0x07
	METAPHONE_N     uint16 = 0x08
	METAPHONE_P     uint16 = 0x09
	METAPHONE_S     uint16 = 0x0A
	METAPHONE_SK    uint16 = ((METAPHONE_S << 4) | METAPHONE_K)
	METAPHONE_T     uint16 = 0x0B
	METAPHONE_TK    uint16 = ((METAPHONE_T << 4) | METAPHONE_K)
	METAPHONE_TS    uint16 = ((METAPHONE_T << 4) | METAPHONE_S)
	METAPHONE_R     uint16 = 0x0C
	METAPHONE_X     uint16 = 0x0D
	METAPHONE_0     uint16 = 0x0E
	METAPHONE_SPACE uint16 = 0x0F
	METAPHONE_NULL  uint16 = 0x00

	/// Sentinel value, used to denote an invalid key
	METAPHONE_INVALID_KEY uint16 = 0xffff
)

/// <summary>Subclass of DoubleMetaphone, Adam Nelson's (anelson@nullpointer.net)
///     C# implementation of Lawrence Phillips' Double Metaphone algorithm,
///     published in C/C++ Users Journal, June, 2000.
///
///     This subclass implements Lawrence's suggested optimization, whereby
///     four-letter metaphone keys are represented as four nibbles in an
///     unsigned short.  This greatly improves storage and search efficiency.</summary>
type shortDoubleMetaphone struct {
	dm *doubleMetaphone

	/// The ushort versions of the primary and alternate keys
	primaryShortKey   uint16
	alternateShortKey uint16
}

/// <summary>Initializes the base class with the given word, then computes
///     ushort representations of the metaphone keys computed by the
///     base class</summary>
///
/// <param name="word">Word for which to compute metaphone keys</param>
func NewShortDoubleMetaphone(word string) ShortDoubleMetaphone {
	sdm := &shortDoubleMetaphone{
		dm: newDoubleMetaphone(word, METAPHONE_KEY_LENGTH),
	}

	sdm.primaryShortKey = sdm.metaphoneKeyToShort(sdm.dm.PrimaryKey())
	if sdm.dm.AlternateKey() != nil {
		sdm.alternateShortKey = sdm.metaphoneKeyToShort(*sdm.dm.AlternateKey())
	} else {
		sdm.alternateShortKey = METAPHONE_INVALID_KEY
	}

	return sdm
}

/// <summary>The primary metaphone key, represented as a ushort</summary>
func (sdm *shortDoubleMetaphone) PrimaryShortKey() uint16 {
	return sdm.primaryShortKey
}

/// <summary>The alternative metaphone key, or METAPHONE_INVALID_KEY if the current
///     word has no alternate key by double metaphone</summary>
func (sdm *shortDoubleMetaphone) AlternateShortKey() uint16 {
	return sdm.alternateShortKey
}

/// <summary>Represents a string metaphone key as a ushort</summary>
///
/// <param name="metaphoneKey">String metaphone key.  Must be four chars long; if you change
///     METAPHONE_KEY_LENGTH in DoubleMetaphone, this will break.  Length
///     tests are not performed, for performance reasons.</param>
///
/// <returns>ushort representation of the given metahphone key</returns>
func (sdm *shortDoubleMetaphone) metaphoneKeyToShort(metaphoneKey string) uint16 {
	var result, charResult uint16
	var currentChar rune

	for currentCharIdx := 0; currentCharIdx < len(metaphoneKey); currentCharIdx++ {
		currentChar = rune(metaphoneKey[currentCharIdx])
		if currentChar == 'A' {
			charResult = METAPHONE_A
		} else if currentChar == 'P' {
			charResult = METAPHONE_P
		} else if currentChar == 'S' {
			charResult = METAPHONE_S
		} else if currentChar == 'K' {
			charResult = METAPHONE_K
		} else if currentChar == 'X' {
			charResult = METAPHONE_X
		} else if currentChar == 'J' {
			charResult = METAPHONE_J
		} else if currentChar == 'T' {
			charResult = METAPHONE_T
		} else if currentChar == 'F' {
			charResult = METAPHONE_F
		} else if currentChar == 'N' {
			charResult = METAPHONE_N
		} else if currentChar == 'H' {
			charResult = METAPHONE_H
		} else if currentChar == 'M' {
			charResult = METAPHONE_M
		} else if currentChar == 'L' {
			charResult = METAPHONE_L
		} else if currentChar == 'R' {
			charResult = METAPHONE_R
		} else if currentChar == ' ' {
			charResult = METAPHONE_SPACE
		} else if currentChar == '\x00' {
			charResult = METAPHONE_0
		} else {
			charResult = 0x00 //This should never happen
		}

		result <<= 4
		result |= charResult
	}
	return result
}
