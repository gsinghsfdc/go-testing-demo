package mystrings

import (
	"fmt"
	"strings"
	"testing"

	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
)

func Example() {
	fmt.Printf("Created By: %v, and %v", Upper("Gurvinder"), Lower("Kenny"))
	// Output: Created By: GURVINDER, and kenny
}

func TestLower(t *testing.T) {
	ts := map[string]string{
		"HELLO":   "hello",
		"WORLD[]": "world[]",
		"":        "",
	}

	for s, v := range ts {
		t.Run(s, func(t *testing.T) {
			if r := Lower(s); r != v {
				t.Fatalf("%v != %v", r, v)
			}
		})
	}
}

func TestUpper(t *testing.T) {
	ts := map[string]string{
		"hello":      "HELLO",
		"first-name": "FIRST-NAME",
		"":           "",
	}

	for s, v := range ts {
		t.Run(s, func(t *testing.T) {
			if r := Upper(s); r != v {
				t.Fatalf("%v != %v", r, v)
			}
		})
	}
}

func TestLowerMatchesStrings(t *testing.T) {
	properties := gopter.NewProperties(nil)
	properties.Property("my implimentation matches strings.Lower", prop.ForAll(
		func(s string) bool {
			if r := Lower(s); r != strings.ToLower(s) {
				t.Logf("Lower(%+v) == %+v", s, r)
				return false
			}

			return true
		},
		genStringRange('A', 'z'),
	))
	properties.TestingRun(t)
}

func TestUppoerMatchesStrings(t *testing.T) {
	properties := gopter.NewProperties(nil)
	properties.Property("my implimentation matches strings.Upper", prop.ForAll(
		func(s string) bool {
			if r := Upper(s); r != strings.ToUpper(s) {
				t.Logf("Upper(%+v) == %+v", s, r)
				return false
			}

			return true
		},
		genStringRange('A', 'z'),
		//genStringRange(' ', '~'),
	))
	properties.TestingRun(t)
}

func genStringRange(min, max rune) gopter.Gen {
	return gopter.DeriveGen(
		func(rs []rune) string {
			return string(rs)
		},
		func(s string) []rune {
			return []rune(s)
		},
		gen.SliceOf(gen.RuneRange(min, max)),
	)
}
