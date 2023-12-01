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
			wantPrimary:   "KMPR",
			wantAlternate: nil,
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
