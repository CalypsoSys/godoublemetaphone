package godoublemetaphone

import (
	"fmt"
	"strings"
)

/**
 * doublemetaphone.go
 *
 * An implemenatation of Lawrence Phillips' Double Metaphone phonetic matching
 * algorithm, published in C/C++ Users Journal, June, 2000.
 *
 * This implementation is based on the C# implementation that was written
 * by Adam J. Nelson (anelson@nullpointer.net). That was based on the C++ template implementation, also by Adam Nelson.
 * For the latest version of this implementation, implementations
 * in other languages, and links to articles I've written on the use of my various
 * Double Metaphone implementations, see:
 * http;//www.nullpointer.net/anelson/
 *
 */

const (
	METAPHONE_KEY_LENGTH = 4 //The length of the metaphone keys produced.  4 is sweet spot
)

type DoubleMetaphone interface {
	PrimaryKey() string
	AlternateKey() *string
	Word() string
}

type doubleMetaphone struct {
	///StringBuilders used to construct the keys
	primaryKey   []rune
	alternateKey []rune

	///Actual keys, populated after construction
	primaryKeyString   string
	alternateKeyString string

	///Variables to track the key length w/o having to grab the .Length attr
	primaryKeyLength   int
	alternateKeyLength int

	///Working copy of the word, and the original word
	word         string
	originalWord string

	///Length and last valid zero-based index into word
	length int
	last   int

	///Flag indicating if an alternate metaphone key was computed for the word
	hasAlternate bool
}

func NewDoubleMetaphone(word string) DoubleMetaphone {
	return newDoubleMetaphone(word)
}

func newDoubleMetaphone(word string) *doubleMetaphone {
	dm := &doubleMetaphone{
		//Leave room at the end for writing a bit beyond the length; keys are chopped at the end anyway
		//primaryKey:   [METAPHONE_KEY_LENGTH + 2]rune{},
		//alternateKey: [METAPHONE_KEY_LENGTH + 2]rune{},
		primaryKey:   []rune{},
		alternateKey: []rune{},
	}

	return dm
}

/// <summary>The primary metaphone key for the current word</summary>
func (dm *doubleMetaphone) PrimaryKey() string {
	return dm.primaryKeyString
}

/// <summary>The alternate metaphone key for the current word, or null if the current
///     word does not have an alternate key by Double Metaphone</summary>
func (dm *doubleMetaphone) AlternateKey() *string {
	if dm.hasAlternate {
		return &dm.alternateKeyString
	}

	return nil
}

/// <summary>Original word for which the keys were computed</summary>
func (dm *doubleMetaphone) Word() string {
	return dm.originalWord
}

/// <summary>Static wrapper around the class, enables computation of metaphone keys
///     without instantiating a class.</summary>
///
/// <param name="word">Word whose metaphone keys are to be computed</param>
/// <param name="primaryKey">Ref to var to receive primary metaphone key</param>
/// <param name="alternateKey">Ref to var to receive alternate metaphone key, or be set to null if
///     word has no alternate key by double metaphone</param>

/* TODO
static public void doubleMetaphone(String word, ref String primaryKey, ref String alternateKey)
{
	DoubleMetaphone mp = new DoubleMetaphone(word);

	primaryKey = mp.PrimaryKey;
	alternateKey = mp.AlternateKey;
}
*/

/// <summary>Sets a new current word for the instance, computing the new word's metaphone
///     keys</summary>
///
/// <param name="word">New word to set to current word.  Discards previous metaphone keys,
///     and computes new keys for this word</param>
func (dm *doubleMetaphone) computeKeys(word string) {
	dm.primaryKey = []rune{}
	dm.alternateKey = []rune{}

	dm.primaryKeyString = ""
	dm.alternateKeyString = ""

	dm.primaryKeyLength = 0
	dm.alternateKeyLength = 0

	dm.hasAlternate = false

	dm.originalWord = word

	//Copy word to an internal working buffer so it can be modified
	dm.word = word

	dm.length = len(dm.word)

	//Compute last valid index into word
	dm.last = dm.length - 1

	//Padd with four spaces, so word can be over-indexed without fear of exception
	dm.word = fmt.Sprintf("     %s", dm.word)

	//Convert to upper case, since metaphone is not case sensitive
	dm.word = strings.ToUpper(dm.word)

	//Now build the keys
	dm.buildMetaphoneKeys()
}

/**
* Internal impl of double metaphone algorithm.  Populates dm.primaryKey and dm.alternateKey.  Modified copy-past of
* Phillips' original code
 */
func (dm *doubleMetaphone) buildMetaphoneKeys() {
	current := 0
	if dm.length < 1 {
		return
	}

	//skip these when at start of word
	if dm.areStringsAt(0, 2, "GN", "KN", "PN", "WR", "PS") {
		current += 1
	}

	//Initial 'X' is pronounced 'Z' e.g. 'Xavier'
	if dm.word[0] == 'X' {
		dm.addMetaphoneCharacter("S") //'Z' maps to 'S'
		current += 1
	}

	///////////main loop//////////////////////////
	for (dm.primaryKeyLength < METAPHONE_KEY_LENGTH) || (dm.alternateKeyLength < METAPHONE_KEY_LENGTH) {
		if current >= dm.length {
			break
		}

		switch dm.word[current] {
		case 'A':
			fallthrough
		case 'E':
			fallthrough
		case 'I':
			fallthrough
		case 'O':
			fallthrough
		case 'U':
			fallthrough
		case 'Y':
			if current == 0 {
				//all init vowels now map to 'A'
				dm.addMetaphoneCharacter("A")
			}
			current += 1

		case 'B':
			//"-mb", e.g", "dumb", already skipped over...
			dm.addMetaphoneCharacter("P")

			if dm.word[current+1] == 'B' {
				current += 2
			} else {
				current += 1
			}

		case 'Ç':
			dm.addMetaphoneCharacter("S")
			current += 1

		case 'C':
			//various germanic
			if (current > 1) &&
				!dm.isVowel(current-2) &&
				dm.areStringsAt((current-1), 3, "ACH") &&
				((dm.word[current+2] != 'I') && ((dm.word[current+2] != 'E') ||
					dm.areStringsAt((current-2), 6, "BACHER", "MACHER"))) {
				dm.addMetaphoneCharacter("K")
				current += 2
				break
			}

			//special case 'caesar'
			if (current == 0) && dm.areStringsAt(current, 6, "CAESAR") {
				dm.addMetaphoneCharacter("S")
				current += 2
				break
			}

			//italian 'chianti'
			if dm.areStringsAt(current, 4, "CHIA") {
				dm.addMetaphoneCharacter("K")
				current += 2
				break
			}

			if dm.areStringsAt(current, 2, "CH") {
				//find 'michael'
				if (current > 0) && dm.areStringsAt(current, 4, "CHAE") {
					dm.addMetaphoneCharacters("K", "X")
					current += 2
					break
				}

				//greek roots e.g. 'chemistry', 'chorus'
				if (current == 0) &&
					(dm.areStringsAt((current+1), 5, "HARAC", "HARIS") ||
						dm.areStringsAt((current+1), 3, "HOR", "HYM", "HIA", "HEM")) &&
					!dm.areStringsAt(0, 5, "CHORE") {
					dm.addMetaphoneCharacter("K")
					current += 2
					break
				}

				//germanic, greek, or otherwise 'ch' for 'kh' sound
				if (dm.areStringsAt(0, 4, "VAN ", "VON ") || dm.areStringsAt(0, 3, "SCH")) ||
					// 'architect but not 'arch', 'orchestra', 'orchid'
					dm.areStringsAt((current-2), 6, "ORCHES", "ARCHIT", "ORCHID") ||
					dm.areStringsAt((current+2), 1, "T", "S") ||
					((dm.areStringsAt((current-1), 1, "A", "O", "U", "E") || (current == 0)) &&
						//e.g., 'wachtler', 'wechsler', but not 'tichner'
						dm.areStringsAt((current+2), 1, "L", "R", "N", "M", "B", "H", "F", "V", "W", " ")) {
					dm.addMetaphoneCharacter("K")
				} else {
					if current > 0 {
						if dm.areStringsAt(0, 2, "MC") {
							//e.g., "McHugh"
							dm.addMetaphoneCharacter("K")
						} else {
							dm.addMetaphoneCharacters("X", "K")
						}
					} else {
						dm.addMetaphoneCharacter("X")
					}
				}
				current += 2
				break
			}
			//e.g, 'czerny'
			if dm.areStringsAt(current, 2, "CZ") && !dm.areStringsAt((current-2), 4, "WICZ") {
				dm.addMetaphoneCharacters("S", "X")
				current += 2
				break
			}

			//e.g., 'focaccia'
			if dm.areStringsAt((current + 1), 3, "CIA") {
				dm.addMetaphoneCharacter("X")
				current += 3
				break
			}

			//double 'C', but not if e.g. 'McClellan'
			if dm.areStringsAt(current, 2, "CC") && !((current == 1) && (dm.word[0] == 'M')) {
				//'bellocchio' but not 'bacchus'
				if dm.areStringsAt((current+2), 1, "I", "E", "H") && !dm.areStringsAt((current+2), 2, "HU") {
					//'accident', 'accede' 'succeed'
					if ((current == 1) && (dm.word[current-1] == 'A')) ||
						dm.areStringsAt((current-1), 5, "UCCEE", "UCCES") {
						dm.addMetaphoneCharacter("KS")
						//'bacci', 'bertucci', other italian
					} else {
						dm.addMetaphoneCharacter("X")
					}
					current += 3
					break
				} else { //Pierce's rule
					dm.addMetaphoneCharacter("K")
					current += 2
					break
				}
			}

			if dm.areStringsAt(current, 2, "CK", "CG", "CQ") {
				dm.addMetaphoneCharacter("K")
				current += 2
				break
			}

			if dm.areStringsAt(current, 2, "CI", "CE", "CY") {
				//italian vs. english
				if dm.areStringsAt(current, 3, "CIO", "CIE", "CIA") {
					dm.addMetaphoneCharacters("S", "X")
				} else {
					dm.addMetaphoneCharacter("S")
				}
				current += 2
				break
			}

			//else
			dm.addMetaphoneCharacter("K")

			//name sent in 'mac caffrey', 'mac gregor
			if dm.areStringsAt((current + 1), 2, " C", " Q", " G") {
				current += 3
			} else {
				if dm.areStringsAt((current+1), 1, "C", "K", "Q") &&
					!dm.areStringsAt((current+1), 2, "CE", "CI") {
					current += 2
				} else {
					current += 1
				}
			}

		case 'D':
			if dm.areStringsAt(current, 2, "DG") {
				if dm.areStringsAt((current + 2), 1, "I", "E", "Y") {
					//e.g. 'edge'
					dm.addMetaphoneCharacter("J")
					current += 3
					break
				} else {
					//e.g. 'edgar'
					dm.addMetaphoneCharacter("TK")
					current += 2
					break
				}
			}

			if dm.areStringsAt(current, 2, "DT", "DD") {
				dm.addMetaphoneCharacter("T")
				current += 2
				break
			}

			//else
			dm.addMetaphoneCharacter("T")
			current += 1

		case 'F':
			if dm.word[current+1] == 'F' {
				current += 2
			} else {
				current += 1
			}
			dm.addMetaphoneCharacter("F")

		case 'G':
			if dm.word[current+1] == 'H' {
				if (current > 0) && !dm.isVowel(current-1) {
					dm.addMetaphoneCharacter("K")
					current += 2
					break
				}

				if current < 3 {
					//'ghislane', ghiradelli
					if current == 0 {
						if dm.word[current+2] == 'I' {

							dm.addMetaphoneCharacter("J")
						} else {
							dm.addMetaphoneCharacter("K")
						}
						current += 2
						break
					}
				}
				//Parker's rule (with some further refinements) - e.g., 'hugh'
				if ((current > 1) && dm.areStringsAt((current-2), 1, "B", "H", "D")) ||
					//e.g., 'bough'
					((current > 2) && dm.areStringsAt((current-3), 1, "B", "H", "D")) ||
					//e.g., 'broughton'
					((current > 3) && dm.areStringsAt((current-4), 1, "B", "H")) {
					current += 2
					break
				} else {
					//e.g., 'laugh', 'McLaughlin', 'cough', 'gough', 'rough', 'tough'
					if (current > 2) &&
						(dm.word[current-1] == 'U') &&
						dm.areStringsAt((current-3), 1, "C", "G", "L", "R", "T") {
						dm.addMetaphoneCharacter("F")
					} else {
						if (current > 0) && dm.word[current-1] != 'I' {
							dm.addMetaphoneCharacter("K")
						}
					}

					current += 2
					break
				}
			}

			if dm.word[current+1] == 'N' {
				if (current == 1) && dm.isVowel(0) && !dm.isWordSlavoGermanic() {
					dm.addMetaphoneCharacters("KN", "N")
				} else {
					//not e.g. 'cagney'
					if !dm.areStringsAt((current+2), 2, "EY") &&
						(dm.word[current+1] != 'Y') && !dm.isWordSlavoGermanic() {
						dm.addMetaphoneCharacters("N", "KN")
					} else {
						dm.addMetaphoneCharacter("KN")
					}
				}
				current += 2
				break
			}

			//'tagliaro'
			if dm.areStringsAt((current+1), 2, "LI") && !dm.isWordSlavoGermanic() {
				dm.addMetaphoneCharacters("KL", "L")
				current += 2
				break
			}

			//-ges-,-gep-,-gel-, -gie- at beginning
			if (current == 0) &&
				((dm.word[current+1] == 'Y') ||
					dm.areStringsAt((current+1), 2, "ES", "EP", "EB", "EL", "EY", "IB", "IL", "IN", "IE", "EI", "ER")) {
				dm.addMetaphoneCharacters("K", "J")
				current += 2
				break
			}

			// -ger-,  -gy-
			if (dm.areStringsAt((current+1), 2, "ER") || (dm.word[current+1] == 'Y')) &&
				!dm.areStringsAt(0, 6, "DANGER", "RANGER", "MANGER") &&
				!dm.areStringsAt((current-1), 1, "E", "I") &&
				!dm.areStringsAt((current-1), 3, "RGY", "OGY") {
				dm.addMetaphoneCharacters("K", "J")
				current += 2
				break
			}

			// italian e.g, 'biaggi'
			if dm.areStringsAt((current+1), 1, "E", "I", "Y") || dm.areStringsAt((current-1), 4, "AGGI", "OGGI") {
				//obvious germanic
				if (dm.areStringsAt(0, 4, "VAN ", "VON ") || dm.areStringsAt(0, 3, "SCH")) ||
					dm.areStringsAt((current+1), 2, "ET") {
					dm.addMetaphoneCharacter("K")
				} else {
					//always soft if french ending
					if dm.areStringsAt((current + 1), 4, "IER ") {
						dm.addMetaphoneCharacter("J")
					} else {
						dm.addMetaphoneCharacters("J", "K")
					}
				}
				current += 2
				break
			}

			if dm.word[current+1] == 'G' {
				current += 2
			} else {
				current += 1
			}
			dm.addMetaphoneCharacter("K")

		case 'H':
			//only keep if first & before vowel or btw. 2 vowels
			if ((current == 0) || dm.isVowel(current-1)) &&
				dm.isVowel(current+1) {
				dm.addMetaphoneCharacter("H")
				current += 2
			} else { //also takes care of 'HH'
				current += 1
			}

		case 'J':
			//obvious spanish, 'jose', 'san jacinto'
			if dm.areStringsAt(current, 4, "JOSE") || dm.areStringsAt(0, 4, "SAN ") {
				if ((current == 0) && (dm.word[current+4] == ' ')) || dm.areStringsAt(0, 4, "SAN ") {
					dm.addMetaphoneCharacter("H")
				} else {
					dm.addMetaphoneCharacters("J", "H")
				}
				current += 1
				break
			}

			if (current == 0) && !dm.areStringsAt(current, 4, "JOSE") {
				dm.addMetaphoneCharacters("J", "A") //Yankelovich/Jankelowicz
			} else {
				//spanish pron. of e.g. 'bajador'
				if dm.isVowel(current-1) &&
					!dm.isWordSlavoGermanic() &&
					((dm.word[current+1] == 'A') || (dm.word[current+1] == 'O')) {
					dm.addMetaphoneCharacters("J", "H")
				} else {
					if current == dm.last {
						dm.addMetaphoneCharacters("J", " ")
					} else {
						if !dm.areStringsAt((current+1), 1, "L", "T", "K", "S", "N", "M", "B", "Z") &&
							!dm.areStringsAt((current-1), 1, "S", "K", "L") {
							dm.addMetaphoneCharacter("J")
						}
					}
				}
			}

			if dm.word[current+1] == 'J' { //it could happen!
				current += 2
			} else {
				current += 1
			}

		case 'K':
			if dm.word[current+1] == 'K' {
				current += 2
			} else {
				current += 1
			}
			dm.addMetaphoneCharacter("K")

		case 'L':
			if dm.word[current+1] == 'L' {
				//spanish e.g. 'cabrillo', 'gallegos'
				if ((current == (dm.length - 3)) &&
					dm.areStringsAt((current-1), 4, "ILLO", "ILLA", "ALLE")) ||
					((dm.areStringsAt((dm.last-1), 2, "AS", "OS") || dm.areStringsAt(dm.last, 1, "A", "O")) &&
						dm.areStringsAt((current-1), 4, "ALLE")) {
					dm.addMetaphoneCharacters("L", " ")
					current += 2
					break
				}
				current += 2
			} else {
				current += 1
			}
			dm.addMetaphoneCharacter("L")

		case 'M':
			if (dm.areStringsAt((current-1), 3, "UMB") &&
				(((current + 1) == dm.last) || dm.areStringsAt((current+2), 2, "ER"))) ||
				//'dumb','thumb'
				(dm.word[current+1] == 'M') {
				current += 2
			} else {
				current += 1
			}
			dm.addMetaphoneCharacter("M")

		case 'N':
			if dm.word[current+1] == 'N' {
				current += 2
			} else {
				current += 1
			}
			dm.addMetaphoneCharacter("N")

		case 'Ñ':
			current += 1
			dm.addMetaphoneCharacter("N")

		case 'P':
			if dm.word[current+1] == 'H' {
				dm.addMetaphoneCharacter("F")
				current += 2
				break
			}

			//also account for "campbell", "raspberry"
			if dm.areStringsAt((current + 1), 1, "P", "B") {
				current += 2
			} else {
				current += 1
			}
			dm.addMetaphoneCharacter("P")

		case 'Q':
			if dm.word[current+1] == 'Q' {
				current += 2
			} else {
				current += 1
			}
			dm.addMetaphoneCharacter("K")

		case 'R':
			//french e.g. 'rogier', but exclude 'hochmeier'
			if (current == dm.last) &&
				!dm.isWordSlavoGermanic() &&
				dm.areStringsAt((current-2), 2, "IE") &&
				!dm.areStringsAt((current-4), 2, "ME", "MA") {
				dm.addMetaphoneCharacters("", "R")
			} else {
				dm.addMetaphoneCharacter("R")
			}

			if dm.word[current+1] == 'R' {
				current += 2
			} else {
				current += 1
			}

		case 'S':
			//special cases 'island', 'isle', 'carlisle', 'carlysle'
			if dm.areStringsAt((current - 1), 3, "ISL", "YSL") {
				current += 1
				break
			}

			//special case 'sugar-'
			if (current == 0) && dm.areStringsAt(current, 5, "SUGAR") {
				dm.addMetaphoneCharacters("X", "S")
				current += 1
				break
			}

			if dm.areStringsAt(current, 2, "SH") {
				//germanic
				if dm.areStringsAt((current + 1), 4, "HEIM", "HOEK", "HOLM", "HOLZ") {
					dm.addMetaphoneCharacter("S")
				} else {
					dm.addMetaphoneCharacter("X")
				}
				current += 2
				break
			}

			//italian & armenian
			if dm.areStringsAt(current, 3, "SIO", "SIA") || dm.areStringsAt(current, 4, "SIAN") {
				if !dm.isWordSlavoGermanic() {
					dm.addMetaphoneCharacters("S", "X")
				} else {
					dm.addMetaphoneCharacter("S")
				}
				current += 3
				break
			}

			//german & anglicisations, e.g. 'smith' match 'schmidt', 'snider' match 'schneider'
			//also, -sz- in slavic language altho in hungarian it is pronounced 's'
			if ((current == 0) &&
				dm.areStringsAt((current+1), 1, "M", "N", "L", "W")) ||
				dm.areStringsAt((current+1), 1, "Z") {
				dm.addMetaphoneCharacters("S", "X")
				if dm.areStringsAt((current + 1), 1, "Z") {
					current += 2
				} else {
					current += 1
				}
				break
			}

			if dm.areStringsAt(current, 2, "SC") {
				//Schlesinger's rule
				if dm.word[current+2] == 'H' {
					//dutch origin, e.g. 'school', 'schooner'
					if dm.areStringsAt((current + 3), 2, "OO", "ER", "EN", "UY", "ED", "EM") {
						//'schermerhorn', 'schenker'
						if dm.areStringsAt((current + 3), 2, "ER", "EN") {
							dm.addMetaphoneCharacters("X", "SK")
						} else {
							dm.addMetaphoneCharacter("SK")
						}
						current += 3
						break
					} else {
						if (current == 0) && !dm.isVowel(3) && (dm.word[3] != 'W') {
							dm.addMetaphoneCharacters("X", "S")
						} else {
							dm.addMetaphoneCharacter("X")
						}
						current += 3
						break
					}
				}

				if dm.areStringsAt((current + 2), 1, "I", "E", "Y") {
					dm.addMetaphoneCharacter("S")
					current += 3
					break
				}
				//else
				dm.addMetaphoneCharacter("SK")
				current += 3
				break
			}

			//french e.g. 'resnais', 'artois'
			if (current == dm.last) && dm.areStringsAt((current-2), 2, "AI", "OI") {
				dm.addMetaphoneCharacters("", "S")
			} else {
				dm.addMetaphoneCharacter("S")
			}

			if dm.areStringsAt((current + 1), 1, "S", "Z") {
				current += 2
			} else {
				current += 1
			}

		case 'T':
			if dm.areStringsAt(current, 4, "TION") {
				dm.addMetaphoneCharacter("X")
				current += 3
				break
			}

			if dm.areStringsAt(current, 3, "TIA", "TCH") {
				dm.addMetaphoneCharacter("X")
				current += 3
				break
			}

			if dm.areStringsAt(current, 2, "TH") ||
				dm.areStringsAt(current, 3, "TTH") {
				//special case 'thomas', 'thames' or germanic
				if dm.areStringsAt((current+2), 2, "OM", "AM") ||
					dm.areStringsAt(0, 4, "VAN ", "VON ") ||
					dm.areStringsAt(0, 3, "SCH") {
					dm.addMetaphoneCharacter("T")
				} else {
					dm.addMetaphoneCharacters("0", "T")
				}
				current += 2
				break
			}

			if dm.areStringsAt((current + 1), 1, "T", "D") {
				current += 2
			} else {
				current += 1
			}
			dm.addMetaphoneCharacter("T")

		case 'V':
			if dm.word[current+1] == 'V' {
				current += 2
			} else {
				current += 1
			}
			dm.addMetaphoneCharacter("F")

		case 'W':
			//can also be in middle of word
			if dm.areStringsAt(current, 2, "WR") {
				dm.addMetaphoneCharacter("R")
				current += 2
				break
			}

			if (current == 0) &&
				(dm.isVowel(current+1) || dm.areStringsAt(current, 2, "WH")) {
				//Wasserman should match Vasserman
				if dm.isVowel(current + 1) {
					dm.addMetaphoneCharacters("A", "F")
				} else {
					//need Uomo to match Womo
					dm.addMetaphoneCharacter("A")
				}
			}

			//Arnow should match Arnoff
			if ((current == dm.last) && dm.isVowel(current-1)) ||
				dm.areStringsAt((current-1), 5, "EWSKI", "EWSKY", "OWSKI", "OWSKY") ||
				dm.areStringsAt(0, 3, "SCH") {
				dm.addMetaphoneCharacters("", "F")
				current += 1
				break
			}

			//polish e.g. 'filipowicz'
			if dm.areStringsAt(current, 4, "WICZ", "WITZ") {
				dm.addMetaphoneCharacters("TS", "FX")
				current += 4
				break
			}

			//else skip it
			current += 1

		case 'X':
			//french e.g. breaux
			if !((current == dm.last) &&
				(dm.areStringsAt((current-3), 3, "IAU", "EAU") ||
					dm.areStringsAt((current-2), 2, "AU", "OU"))) {
				dm.addMetaphoneCharacter("KS")
			}

			if dm.areStringsAt((current + 1), 1, "C", "X") {
				current += 2
			} else {
				current += 1
			}

		case 'Z':
			//chinese pinyin e.g. 'zhao'
			if dm.word[current+1] == 'H' {
				dm.addMetaphoneCharacter("J")
				current += 2
				break
			} else {

				if dm.areStringsAt((current+1), 2, "ZO", "ZI", "ZA") ||
					(dm.isWordSlavoGermanic() && ((current > 0) && dm.word[current-1] != 'T')) {
					dm.addMetaphoneCharacters("S", "TS")
				} else {
					dm.addMetaphoneCharacter("S")
				}
			}

			if dm.word[current+1] == 'Z' {
				current += 2
			} else {
				current += 1
			}

		default:
			current += 1
		}
	}

	//Finally, chop off the keys at the proscribed length
	if dm.primaryKeyLength > METAPHONE_KEY_LENGTH {
		dm.primaryKey = dm.primaryKey[:METAPHONE_KEY_LENGTH]
	}

	if dm.alternateKeyLength > METAPHONE_KEY_LENGTH {
		dm.alternateKey = dm.alternateKey[:METAPHONE_KEY_LENGTH]
	}

	dm.primaryKeyString = string(dm.primaryKey)
	dm.alternateKeyString = string(dm.alternateKey)
}

/**
* Returns true if dm.word is classified as "slavo-germanic" by Phillips' algorithm
*
* @return true if word contains strings that Lawrence's algorithm considers indicative of
*         slavo-germanic origin; else false
 */
func (dm *doubleMetaphone) isWordSlavoGermanic() bool {
	if (strings.Contains(dm.word, "W")) ||
		(strings.Contains(dm.word, "K")) ||
		(strings.Contains(dm.word, "CZ")) ||
		(strings.Contains(dm.word, "WITZ")) {
		return true
	}

	return false
}

/**
* Returns true if letter at given position in word is a Roman vowel
*
* @param pos    Position at which to check for a vowel
*
* @return True if dm.word[pos] is a Roman vowel, else false
 */
func (dm *doubleMetaphone) isVowel(pos int) bool {
	if (pos < 0) || (pos >= dm.length) {
		return false
	}

	it := dm.word[pos]

	if (it == 'E') || (it == 'A') || (it == 'I') || (it == 'O') || (it == 'U') || (it == 'Y') {
		return true
	}

	return false
}

func (dm *doubleMetaphone) addMetaphoneCharacter(primaryCharacter string) {
	dm.addMetaphoneCharacterPtr(primaryCharacter, nil)
}

func (dm *doubleMetaphone) addMetaphoneCharacters(primaryCharacter string, alternateCharacter string) {
	dm.addMetaphoneCharacterPtr(primaryCharacter, &alternateCharacter)
}

/**
* Appends a metaphone character to the primary, and a possibly different alternate,
* metaphone keys for the word.
*
* @param primaryCharacter
*               Primary character to append to primary key, and, if no alternate char is present,
*               the alternate key as well
* @param alternateCharacter
*               Alternate character to append to alternate key.  May be null or a zero-length string,
*               in which case the primary character will be appended to the alternate key instead
 */
func (dm *doubleMetaphone) addMetaphoneCharacterPtr(primaryCharacter string, alternateCharacterPtr *string) {
	//Is the primary character valid?
	if len(primaryCharacter) > 0 {
		idx := 0
		for idx < len(primaryCharacter) {
			dm.primaryKey = append(dm.primaryKey, rune(primaryCharacter[idx]))
			dm.primaryKeyLength++
			idx++
		}
	}

	//Is the alternate character valid?
	if alternateCharacterPtr != nil {
		alternateCharacter := *alternateCharacterPtr
		//Alternate character was provided.  If it is not zero-length, append it, else
		//append the primary string as long as it wasn't zero length and isn't a space character
		if len(alternateCharacter) != 0 {
			dm.hasAlternate = true
			if alternateCharacter[0] != ' ' {
				idx := 0
				for idx < len(alternateCharacter) {
					dm.alternateKey = append(dm.alternateKey, rune(alternateCharacter[idx]))
					dm.alternateKeyLength++
					idx++
				}
			}
		} else {
			//No, but if the primary character is valid, add that instead
			if len(primaryCharacter) > 0 && (primaryCharacter[0] != ' ') {
				idx := 0
				for idx < len(primaryCharacter) {
					dm.alternateKey = append(dm.alternateKey, rune(primaryCharacter[idx]))
					dm.alternateKeyLength++
					idx++
				}
			}
		}
	} else if len(primaryCharacter) > 0 {
		//Else, no alternate character was passed, but a primary was, so append the primary character to the alternate key
		idx := 0
		for idx < len(primaryCharacter) {
			dm.alternateKey = append(dm.alternateKey, rune(primaryCharacter[idx]))
			dm.alternateKeyLength++
			idx++
		}
	}
}

/**
* Tests if any of the strings passed as variable arguments are at the given start position and
* length within word
*
* @param start   Start position in dm.word
* @param length  Length of substring starting at start in dm.word to compare to the given strings
* @param strings params array of zero or more strings for which to search in dm.word
*
* @return true if any one string in the strings array was found in dm.word at the given position
*         and length
 */
func (dm *doubleMetaphone) areStringsAt(start int, length int, strs ...string) bool {
	if start < 0 {
		//Sometimes, as a result of expressions like "current - 2" for start,
		//start ends up negative.  Since no string can be present at a negative offset, this is always false
		return false
	}

	target := dm.word[start : start+length]

	for idx := 0; idx < len(strs); idx++ {
		if strs[idx] == target {
			return true
		}
	}

	return false
}
