package godoublemetaphone

import (
	"testing"
)

func compareStringPointers(str1 *string, str2 *string) bool {
	if str1 != nil && str2 != nil {
		return *str1 == *str2
	} else if str1 == nil && str2 == nil {
		return true
	}

	return false
}

func safeString(str *string) string {
	if str != nil {
		return *str
	}

	return "<nil>"
}

func stringPtr(str string) *string {
	return &str
}

func overLap(set1 []string, set2 []string) string {
	for _, str1 := range set1 {
		if str1 != "" {
			for _, str2 := range set2 {
				if str1 == str2 {
					return str1
				}
			}
		}
	}

	return ""
}

func TestSingleResult(t *testing.T) {
	tests := []struct {
		name          string
		arg           string
		wantPrimary   string
		wantAlternate *string
	}{
		{
			name:          "test aubrey",
			arg:           "aubrey",
			wantPrimary:   "APR",
			wantAlternate: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDoubleMetaphone(tt.arg); got.PrimaryKey() != tt.wantPrimary || !compareStringPointers(got.AlternateKey(), tt.wantAlternate) {
				t.Errorf("TestSingleResult = %s %s, want %s %s", got.PrimaryKey(), safeString(got.AlternateKey()), tt.wantPrimary, safeString(tt.wantAlternate))
			}
		})
	}
}

func TestDoubleResult(t *testing.T) {
	tests := []struct {
		name          string
		arg           string
		wantPrimary   string
		wantAlternate *string
	}{
		{
			name:          "test richard",
			arg:           "richard",
			wantPrimary:   "RXRT",
			wantAlternate: stringPtr("RKRT"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDoubleMetaphone(tt.arg); got.PrimaryKey() != tt.wantPrimary || !compareStringPointers(got.AlternateKey(), tt.wantAlternate) {
				t.Errorf("TestDoubleResult = %s %s, want %s %s", got.PrimaryKey(), safeString(got.AlternateKey()), tt.wantPrimary, safeString(tt.wantAlternate))
			}
		})
	}
}

func TestGeneralWordList(t *testing.T) {
	tests := []struct {
		name          string
		arg           string
		wantPrimary   string
		wantAlternate *string
	}{
		{
			name:          "test richard",
			arg:           "richard",
			wantPrimary:   "RXRT",
			wantAlternate: stringPtr("RKRT"),
		},
		{
			name:          "test Jose",
			arg:           "Jose",
			wantPrimary:   "HS",
			wantAlternate: nil,
		},
		{
			name:          "test cambrillo",
			arg:           "cambrillo",
			wantPrimary:   "KMPRL",
			wantAlternate: stringPtr("KMPR"),
		},
		{
			name:          "test otto",
			arg:           "otto",
			wantPrimary:   "AT",
			wantAlternate: nil,
		},
		{
			name:          "test aubrey",
			arg:           "aubrey",
			wantPrimary:   "APR",
			wantAlternate: nil,
		},
		{
			name:          "test maurice",
			arg:           "maurice",
			wantPrimary:   "MRS",
			wantAlternate: nil,
		},
		{
			name:          "test auto",
			arg:           "auto",
			wantPrimary:   "AT",
			wantAlternate: nil,
		},
		{
			name:          "test maisey",
			arg:           "maisey",
			wantPrimary:   "MS",
			wantAlternate: nil,
		},
		{
			name:          "test catherine",
			arg:           "catherine",
			wantPrimary:   "K0RN",
			wantAlternate: stringPtr("KTRN"),
		},
		{
			name:          "test geoff",
			arg:           "geoff",
			wantPrimary:   "JF",
			wantAlternate: stringPtr("KF"),
		},
		{
			name:          "test Chile",
			arg:           "Chile",
			wantPrimary:   "XL",
			wantAlternate: nil,
		},
		{
			name:          "test katherine",
			arg:           "katherine",
			wantPrimary:   "K0RN",
			wantAlternate: stringPtr("KTRN"),
		},
		{
			name:          "test steven",
			arg:           "steven",
			wantPrimary:   "STFN",
			wantAlternate: nil,
		},
		{
			name:          "test zhang",
			arg:           "zhang",
			wantPrimary:   "JNK",
			wantAlternate: nil,
		},
		{
			name:          "test bob",
			arg:           "bob",
			wantPrimary:   "PP",
			wantAlternate: nil,
		},
		{
			name:          "test ray",
			arg:           "ray",
			wantPrimary:   "R",
			wantAlternate: nil,
		},
		{
			name:          "test Tux",
			arg:           "Tux",
			wantPrimary:   "TKS",
			wantAlternate: nil,
		},
		{
			name:          "test bryan",
			arg:           "bryan",
			wantPrimary:   "PRN",
			wantAlternate: nil,
		},
		{
			name:          "test bryce",
			arg:           "bryce",
			wantPrimary:   "PRS",
			wantAlternate: nil,
		},
		{
			name:          "test Rapelje",
			arg:           "Rapelje",
			wantPrimary:   "RPL",
			wantAlternate: nil,
		},
		{
			name:          "test richard",
			arg:           "richard",
			wantPrimary:   "RXRT",
			wantAlternate: stringPtr("RKRT"),
		},
		{
			name:          "test solilijs",
			arg:           "solilijs",
			wantPrimary:   "SLLS",
			wantAlternate: nil,
		},
		{
			name:          "test Dallas",
			arg:           "Dallas",
			wantPrimary:   "TLS",
			wantAlternate: nil,
		},
		{
			name:          "test Schwein",
			arg:           "Schwein",
			wantPrimary:   "XN",
			wantAlternate: stringPtr("XFN"),
		},
		{
			name:          "test dave",
			arg:           "dave",
			wantPrimary:   "TF",
			wantAlternate: nil,
		},
		{
			name:          "test eric",
			arg:           "eric",
			wantPrimary:   "ARK",
			wantAlternate: nil,
		},
		{
			name:          "test Parachute",
			arg:           "Parachute",
			wantPrimary:   "PRKT",
			wantAlternate: nil,
		},
		{
			name:          "test brian",
			arg:           "brian",
			wantPrimary:   "PRN",
			wantAlternate: nil,
		},
		{
			name:          "test randy",
			arg:           "randy",
			wantPrimary:   "RNT",
			wantAlternate: nil,
		},
		{
			name:          "test Through",
			arg:           "Through",
			wantPrimary:   "0R",
			wantAlternate: stringPtr("TR"),
		},
		{
			name:          "test Nowhere",
			arg:           "Nowhere",
			wantPrimary:   "NR",
			wantAlternate: nil,
		},
		{
			name:          "test heidi",
			arg:           "heidi",
			wantPrimary:   "HT",
			wantAlternate: nil,
		},
		{
			name:          "test Arnow",
			arg:           "Arnow",
			wantPrimary:   "ARN",
			wantAlternate: stringPtr("ARNF"),
		},
		{
			name:          "test Thumbail",
			arg:           "Thumbail",
			wantPrimary:   "0MPL",
			wantAlternate: stringPtr("TMPL"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDoubleMetaphone(tt.arg); got.PrimaryKey() != tt.wantPrimary || !compareStringPointers(got.AlternateKey(), tt.wantAlternate) {
				t.Errorf("TestGeneralWordList = %s %s, want %s %s", got.PrimaryKey(), safeString(got.AlternateKey()), tt.wantPrimary, safeString(tt.wantAlternate))
			}
		})
	}
}

func TestHomophones(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want string
	}{
		{
			name: "test tolled",
			arg:  "tolled",
			want: "told",
		},
		{
			name: "test katherine",
			arg:  "katherine",
			want: "catherine",
		},
		{
			name: "test brian",
			arg:  "brian",
			want: "bryan",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewDoubleMetaphone(tt.arg)
			want := NewDoubleMetaphone(tt.want)
			if got.PrimaryKey() != want.PrimaryKey() || !compareStringPointers(got.AlternateKey(), want.AlternateKey()) {
				t.Errorf("TestHomophones = %s %s, want %s %s", got.PrimaryKey(), safeString(got.AlternateKey()), want.PrimaryKey(), safeString(want.AlternateKey()))
			}
		})
	}
}

func TestSimilarNames1(t *testing.T) {
	tests := []struct {
		name          string
		arg           string
		wantPrimary   string
		wantAlternate *string
	}{
		{
			name:          "test Bartosz",
			arg:           "Bartosz",
			wantPrimary:   "PRTS",
			wantAlternate: stringPtr("PRTX"),
		},
		{
			name:          "test Bartosch",
			arg:           "Bartosch",
			wantPrimary:   "PRTX",
			wantAlternate: nil,
		},
		{
			name:          "test Bartos",
			arg:           "Bartos",
			wantPrimary:   "PRTS",
			wantAlternate: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDoubleMetaphone(tt.arg); got.PrimaryKey() != tt.wantPrimary || !compareStringPointers(got.AlternateKey(), tt.wantAlternate) {
				t.Errorf("TestSimilarNames1 = %s %s, want %s %s", got.PrimaryKey(), safeString(got.AlternateKey()), tt.wantPrimary, safeString(tt.wantAlternate))
			}
		})
	}
}

func TestSimilarNames2(t *testing.T) {
	tests := []struct {
		name string
		arg1 string
		arg2 string
		want string
	}{
		{
			name: "test Jablonski",
			arg1: "Jablonski",
			arg2: "Yablonsky",
			want: "APLNSK",
		},
		{
			name: "test Smith",
			arg1: "Smith",
			arg2: "Schmidt",
			want: "XMT",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewDoubleMetaphone(tt.arg1)
			want := NewDoubleMetaphone(tt.arg2)
			set1 := []string{got.PrimaryKey(), safeString(got.AlternateKey())}
			set2 := []string{want.PrimaryKey(), safeString(want.AlternateKey())}
			if overLap(set1, set2) != tt.want {
				t.Errorf("TestSimilarNames2 = arg1 %s %s arg2 %s %s, want %s", got.PrimaryKey(), safeString(got.AlternateKey()), want.PrimaryKey(), safeString(want.AlternateKey()), tt.want)
			}
		})
	}
}

func TestNonEnglishUnicode(t *testing.T) {
	tests := []struct {
		name          string
		arg           string
		wantPrimary   string
		wantAlternate *string
	}{
		{
			name:          "test andestādītu",
			arg:           "andestādītu",
			wantPrimary:   "ANTSTTT",
			wantAlternate: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDoubleMetaphone(tt.arg); got.PrimaryKey() != tt.wantPrimary || !compareStringPointers(got.AlternateKey(), tt.wantAlternate) {
				t.Errorf("TestNonEnglishUnicode = %s %s, want %s %s", got.PrimaryKey(), safeString(got.AlternateKey()), tt.wantPrimary, safeString(tt.wantAlternate))
			}
		})
	}
}

func TestVariousGerman(t *testing.T) {
	tests := []struct {
		name          string
		arg           string
		wantPrimary   string
		wantAlternate *string
	}{
		{
			name:          "test ach",
			arg:           "ach",
			wantPrimary:   "AK",
			wantAlternate: nil,
		},
		{
			name:          "test bacher",
			arg:           "bacher",
			wantPrimary:   "PKR",
			wantAlternate: nil,
		},
		{
			name:          "test macher",
			arg:           "macher",
			wantPrimary:   "MKR",
			wantAlternate: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDoubleMetaphone(tt.arg); got.PrimaryKey() != tt.wantPrimary || !compareStringPointers(got.AlternateKey(), tt.wantAlternate) {
				t.Errorf("TestVariousGerman = %s %s, want %s %s", got.PrimaryKey(), safeString(got.AlternateKey()), tt.wantPrimary, safeString(tt.wantAlternate))
			}
		})
	}
}

func TestVariousItalian(t *testing.T) {
	tests := []struct {
		name          string
		arg           string
		wantPrimary   string
		wantAlternate *string
	}{
		{
			name:          "test bacci",
			arg:           "bacci",
			wantPrimary:   "PX",
			wantAlternate: nil,
		},
		{
			name:          "test bertucci",
			arg:           "bertucci",
			wantPrimary:   "PRTX",
			wantAlternate: nil,
		},
		{
			name:          "test bellocchio",
			arg:           "bellocchio",
			wantPrimary:   "PLX",
			wantAlternate: nil,
		},
		{
			name:          "test bacchus",
			arg:           "bacchus",
			wantPrimary:   "PKS",
			wantAlternate: nil,
		},
		{
			name:          "test ",
			arg:           "focaccia",
			wantPrimary:   "FKX",
			wantAlternate: nil,
		},
		{
			name:          "test chianti",
			arg:           "chianti",
			wantPrimary:   "KNT",
			wantAlternate: nil,
		},
		{
			name:          "test ",
			arg:           "tagliaro",
			wantPrimary:   "TKLR",
			wantAlternate: stringPtr("TLR"),
		},
		{
			name:          "test biaggi",
			arg:           "biaggi",
			wantPrimary:   "PJ",
			wantAlternate: stringPtr("PK"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDoubleMetaphone(tt.arg); got.PrimaryKey() != tt.wantPrimary || !compareStringPointers(got.AlternateKey(), tt.wantAlternate) {
				t.Errorf("TestVariousItalian = %s %s, want %s %s", got.PrimaryKey(), safeString(got.AlternateKey()), tt.wantPrimary, safeString(tt.wantAlternate))
			}
		})
	}
}

func TestVariousSpanish(t *testing.T) {
	tests := []struct {
		name          string
		arg           string
		wantPrimary   string
		wantAlternate *string
	}{
		{
			name:          "test bajador",
			arg:           "bajador",
			wantPrimary:   "PJTR",
			wantAlternate: stringPtr("PHTR"),
		},
		{
			name:          "test cabrillo",
			arg:           "cabrillo",
			wantPrimary:   "KPRL",
			wantAlternate: stringPtr("KPR"),
		},
		{
			name:          "test gallegos",
			arg:           "gallegos",
			wantPrimary:   "KLKS",
			wantAlternate: stringPtr("KKS"),
		},
		{
			name:        "test San Jacinto",
			arg:         "San Jacinto",
			wantPrimary: "SNHSNT",
			//wantPrimary:   "SNHS",
			wantAlternate: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDoubleMetaphone(tt.arg); got.PrimaryKey() != tt.wantPrimary || !compareStringPointers(got.AlternateKey(), tt.wantAlternate) {
				t.Errorf("TestVariousSpanish = %s %s, want %s %s", got.PrimaryKey(), safeString(got.AlternateKey()), tt.wantPrimary, safeString(tt.wantAlternate))
			}
		})
	}
}

func TestVariousFrench(t *testing.T) {
	tests := []struct {
		name          string
		arg           string
		wantPrimary   string
		wantAlternate *string
	}{
		{
			name:          "test rogier",
			arg:           "rogier",
			wantPrimary:   "RJ",
			wantAlternate: stringPtr("RJR"),
		},
		{
			name:          "test breaux",
			arg:           "breaux",
			wantPrimary:   "PR",
			wantAlternate: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDoubleMetaphone(tt.arg); got.PrimaryKey() != tt.wantPrimary || !compareStringPointers(got.AlternateKey(), tt.wantAlternate) {
				t.Errorf("TestVariousFrench = %s %s, want %s %s", got.PrimaryKey(), safeString(got.AlternateKey()), tt.wantPrimary, safeString(tt.wantAlternate))
			}
		})
	}
}

func TestVariousSlavic(t *testing.T) {
	tests := []struct {
		name          string
		arg           string
		wantPrimary   string
		wantAlternate *string
	}{
		{
			name:          "test Wewski",
			arg:           "Wewski",
			wantPrimary:   "ASK",
			wantAlternate: stringPtr("FFSK"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDoubleMetaphone(tt.arg); got.PrimaryKey() != tt.wantPrimary || !compareStringPointers(got.AlternateKey(), tt.wantAlternate) {
				t.Errorf("TestVariousSlavic = %s %s, want %s %s", got.PrimaryKey(), safeString(got.AlternateKey()), tt.wantPrimary, safeString(tt.wantAlternate))
			}
		})
	}
}

func TestVariousChinese(t *testing.T) {
	tests := []struct {
		name          string
		arg           string
		wantPrimary   string
		wantAlternate *string
	}{
		{
			name:          "test zhao",
			arg:           "zhao",
			wantPrimary:   "J",
			wantAlternate: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDoubleMetaphone(tt.arg); got.PrimaryKey() != tt.wantPrimary || !compareStringPointers(got.AlternateKey(), tt.wantAlternate) {
				t.Errorf("TestVariousChinese = %s %s, want %s %s", got.PrimaryKey(), safeString(got.AlternateKey()), tt.wantPrimary, safeString(tt.wantAlternate))
			}
		})
	}
}

func TestVariousDutch(t *testing.T) {
	tests := []struct {
		name          string
		arg           string
		wantPrimary   string
		wantAlternate *string
	}{
		{
			name:          "test school",
			arg:           "school",
			wantPrimary:   "SKL",
			wantAlternate: nil,
		},
		{
			name:          "test schooner",
			arg:           "schooner",
			wantPrimary:   "SKNR",
			wantAlternate: nil,
		},
		{
			name: "test schermerhorn",
			arg:  "schermerhorn",
			//wantPrimary:   "XRMR",
			//wantAlternate: stringPtr("SKRM"),
			wantPrimary:   "XRMRRN",
			wantAlternate: stringPtr("SKRMRRN"),
		},
		{
			name:        "test schenker",
			arg:         "schenker",
			wantPrimary: "XNKR",
			//wantAlternate: stringPtr("SKNK"),
			wantAlternate: stringPtr("SKNKR"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDoubleMetaphone(tt.arg); got.PrimaryKey() != tt.wantPrimary || !compareStringPointers(got.AlternateKey(), tt.wantAlternate) {
				t.Errorf("TestVariousDutch = %s %s, want %s %s", got.PrimaryKey(), safeString(got.AlternateKey()), tt.wantPrimary, safeString(tt.wantAlternate))
			}
		})
	}
}

func TestChWords(t *testing.T) {
	tests := []struct {
		name          string
		arg           string
		wantPrimary   string
		wantAlternate *string
	}{
		{
			name:          "test Charac",
			arg:           "Charac",
			wantPrimary:   "KRK",
			wantAlternate: nil,
		},
		{
			name:          "test Charis",
			arg:           "Charis",
			wantPrimary:   "KRS",
			wantAlternate: nil,
		},
		{
			name:          "test chord",
			arg:           "chord",
			wantAlternate: nil,
			wantPrimary:   "KRT",
		},
		{
			name:          "test Chym",
			arg:           "Chym",
			wantAlternate: nil,
			wantPrimary:   "KM",
		},
		{
			name:          "test Chia",
			arg:           "Chia",
			wantPrimary:   "K",
			wantAlternate: nil,
		},
		{
			name:          "test chem",
			arg:           "chem",
			wantPrimary:   "KM",
			wantAlternate: nil,
		},
		{
			name:          "test chore",
			arg:           "chore",
			wantPrimary:   "XR",
			wantAlternate: nil,
		},
		{
			name: "test orchestra",
			arg:  "orchestra",
			//wantPrimary: "ARKS",
			wantPrimary:   "ARKSTR",
			wantAlternate: nil,
		},
		{
			name: "test architect",
			arg:  "architect",
			//wantPrimary: "ARKT",
			wantPrimary:   "ARKTKT",
			wantAlternate: nil,
		},
		{
			name:          "test orchid",
			arg:           "orchid",
			wantPrimary:   "ARKT",
			wantAlternate: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDoubleMetaphone(tt.arg); got.PrimaryKey() != tt.wantPrimary || !compareStringPointers(got.AlternateKey(), tt.wantAlternate) {
				t.Errorf("TestChWords = %s %s, want %s %s", got.PrimaryKey(), safeString(got.AlternateKey()), tt.wantPrimary, safeString(tt.wantAlternate))
			}
		})
	}
}

func TestCcWords(t *testing.T) {
	tests := []struct {
		name          string
		arg           string
		wantPrimary   string
		wantAlternate *string
	}{
		{
			name: "test accident",
			arg:  "accident",
			//wantPrimary: "AKST",
			wantPrimary:   "AKSTNT",
			wantAlternate: nil,
		},
		{
			name:          "test accede",
			arg:           "accede",
			wantPrimary:   "AKST",
			wantAlternate: nil,
		},
		{
			name:          "test succeed",
			arg:           "succeed",
			wantPrimary:   "SKST",
			wantAlternate: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDoubleMetaphone(tt.arg); got.PrimaryKey() != tt.wantPrimary || !compareStringPointers(got.AlternateKey(), tt.wantAlternate) {
				t.Errorf("TestMcWords = %s %s, want %s %s", got.PrimaryKey(), safeString(got.AlternateKey()), tt.wantPrimary, safeString(tt.wantAlternate))
			}
		})
	}
}

func TestMcWords(t *testing.T) {
	tests := []struct {
		name          string
		arg           string
		wantPrimary   string
		wantAlternate *string
	}{
		{
			name:          "test mac caffrey",
			arg:           "mac caffrey",
			wantPrimary:   "MKFR",
			wantAlternate: nil,
		},
		{
			name: "test mac gregor",
			arg:  "mac gregor",
			//wantPrimary: "MKRK",
			wantPrimary:   "MKRKR",
			wantAlternate: nil,
		},
		{
			name:          "test mc crae",
			arg:           "mc crae",
			wantPrimary:   "MKR",
			wantAlternate: nil,
		},
		{
			name:          "test mcclain",
			arg:           "mcclain",
			wantPrimary:   "MKLN",
			wantAlternate: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDoubleMetaphone(tt.arg); got.PrimaryKey() != tt.wantPrimary || !compareStringPointers(got.AlternateKey(), tt.wantAlternate) {
				t.Errorf("TestMcWords = %s %s, want %s %s", got.PrimaryKey(), safeString(got.AlternateKey()), tt.wantPrimary, safeString(tt.wantAlternate))
			}
		})
	}
}

func TestGhWords(t *testing.T) {
	tests := []struct {
		name          string
		arg           string
		wantPrimary   string
		wantAlternate *string
	}{
		{
			name:          "test laugh",
			arg:           "laugh",
			wantPrimary:   "LF",
			wantAlternate: nil,
		},
		{
			name:          "test cough",
			arg:           "cough",
			wantPrimary:   "KF",
			wantAlternate: nil,
		},
		{
			name:          "test rough",
			arg:           "rough",
			wantPrimary:   "RF",
			wantAlternate: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDoubleMetaphone(tt.arg); got.PrimaryKey() != tt.wantPrimary || !compareStringPointers(got.AlternateKey(), tt.wantAlternate) {
				t.Errorf("TestGhWords = %s %s, want %s %s", got.PrimaryKey(), safeString(got.AlternateKey()), tt.wantPrimary, safeString(tt.wantAlternate))
			}
		})
	}
}

func TestG3Words(t *testing.T) {
	tests := []struct {
		name          string
		arg           string
		wantPrimary   string
		wantAlternate *string
	}{
		{
			name:          "test gya",
			arg:           "gya",
			wantPrimary:   "K",
			wantAlternate: stringPtr("J"),
		},
		{
			name:          "test ges",
			arg:           "ges",
			wantPrimary:   "KS",
			wantAlternate: stringPtr("JS"),
		},
		{
			name:          "test gep",
			arg:           "gep",
			wantPrimary:   "KP",
			wantAlternate: stringPtr("JP"),
		},
		{
			name:          "test geb",
			arg:           "geb",
			wantPrimary:   "KP",
			wantAlternate: stringPtr("JP"),
		},
		{
			name:          "test gel",
			arg:           "gel",
			wantPrimary:   "KL",
			wantAlternate: stringPtr("JL"),
		},
		{
			name:          "test gey",
			arg:           "gey",
			wantPrimary:   "K",
			wantAlternate: stringPtr("J"),
		},
		{
			name:          "test gib",
			arg:           "gib",
			wantPrimary:   "KP",
			wantAlternate: stringPtr("JP"),
		},
		{
			name:          "test gil",
			arg:           "gil",
			wantPrimary:   "KL",
			wantAlternate: stringPtr("JL"),
		},
		{
			name:          "test gin",
			arg:           "gin",
			wantPrimary:   "KN",
			wantAlternate: stringPtr("JN"),
		},
		{
			name:          "test gie",
			arg:           "gie",
			wantPrimary:   "K",
			wantAlternate: stringPtr("J"),
		},
		{
			name:          "test gei",
			arg:           "gei",
			wantPrimary:   "K",
			wantAlternate: stringPtr("J"),
		},
		{
			name:          "test ger",
			arg:           "ger",
			wantPrimary:   "KR",
			wantAlternate: stringPtr("JR"),
		},
		{
			name:          "test danger",
			arg:           "danger",
			wantPrimary:   "TNJR",
			wantAlternate: stringPtr("TNKR"),
		},
		{
			name:          "test manager",
			arg:           "manager",
			wantPrimary:   "MNKR",
			wantAlternate: stringPtr("MNJR"),
		},
		{
			name:          "test dowager",
			arg:           "dowager",
			wantPrimary:   "TKR",
			wantAlternate: stringPtr("TJR"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDoubleMetaphone(tt.arg); got.PrimaryKey() != tt.wantPrimary || !compareStringPointers(got.AlternateKey(), tt.wantAlternate) {
				t.Errorf("TestG3Words = %s %s, want %s %s", got.PrimaryKey(), safeString(got.AlternateKey()), tt.wantPrimary, safeString(tt.wantAlternate))
			}
		})
	}
}

func TestPbWords(t *testing.T) {
	tests := []struct {
		name          string
		arg           string
		wantPrimary   string
		wantAlternate *string
	}{
		{
			name:          "test Campbell",
			arg:           "Campbell",
			wantPrimary:   "KMPL",
			wantAlternate: nil,
		},
		{
			name:          "test raspberry",
			arg:           "raspberry",
			wantPrimary:   "RSPR",
			wantAlternate: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDoubleMetaphone(tt.arg); got.PrimaryKey() != tt.wantPrimary || !compareStringPointers(got.AlternateKey(), tt.wantAlternate) {
				t.Errorf("TestPbWords = %s %s, want %s %s", got.PrimaryKey(), safeString(got.AlternateKey()), tt.wantPrimary, safeString(tt.wantAlternate))
			}
		})
	}
}

func TestThWords(t *testing.T) {
	tests := []struct {
		name          string
		arg           string
		wantPrimary   string
		wantAlternate *string
	}{
		{
			name:          "test Thomas",
			arg:           "Thomas",
			wantPrimary:   "TMS",
			wantAlternate: nil,
		},
		{
			name:          "test ",
			arg:           "Thames",
			wantPrimary:   "TMS",
			wantAlternate: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDoubleMetaphone(tt.arg); got.PrimaryKey() != tt.wantPrimary || !compareStringPointers(got.AlternateKey(), tt.wantAlternate) {
				t.Errorf("TestThWords = %s %s, want %s %s", got.PrimaryKey(), safeString(got.AlternateKey()), tt.wantPrimary, safeString(tt.wantAlternate))
			}
		})
	}
}
