// Encode ASCII unicode code-points in a slice as runes.
// Copyright ©C. L. Banning, 2016.  See LICENSE file.

// Encode ASCII unicode code-points in a slice as runes.
//
// Examples:
//	data := []byte(`\u044d\u0442\u043e \u0442\u0435\u0441\u0442 \u0441\u043e\u043e\u0431\u0449\u0435\u043d\u0438\u0435`)
//	fmt.Println(unigo.EncodeToString(data) // prints: это тест сообщение
//
//	or:
//	fmt.Println(unigo.EncodeToString([]byte(`{"key":"wasn\u0027t"}`))) // prints: {"key":"wasn't"}
package unigo

import (
	"bytes"
	"fmt"
	"strconv"
	"unicode"
	"unicode/utf16"
	"unicode/utf8"
)

// Encode encodes ASCII unicode code-points in a slice to runes.
func Encode(b []byte) []byte {
	var buf []byte
	out := bytes.NewBuffer(buf)
	bb := make([]byte, 4)
	var n int
	// NOTE: the following is a hack of code in go/src/encoding/json/decode.go.
	for len(b) > 0 {
		if len(b) > 5 && b[0] == '\\' && b[1] == 'u' {
			rr := getu4(b)
			b = b[6:]
			if utf16.IsSurrogate(rr) {
				rr1 := getu4(b)
				if dec := utf16.DecodeRune(rr, rr1); dec != unicode.ReplacementChar {
					// A valid pair; consume.
					b = b[6:]
					n = utf8.EncodeRune(bb, dec)
					fmt.Fprintf(out, "%s", string(bb[:n]))
					continue
				} else {
					// Invalid surrogate; fall back to replacement rune.
					rr = unicode.ReplacementChar
				}
			}
			n = utf8.EncodeRune(bb, rr)
			fmt.Fprintf(out, "%s", string(bb[:n]))
			continue
		}
		fmt.Fprintf(out, "%s", string(b[:1]))
		if len(b) > 1 {
			b = b[1:]
		} else {
			b = nil
		}
	}
	return out.Bytes()
}

// EncodeToString casts the Encode return to a string.
func EncodeToString(b []byte) string {
	return string(Encode(b))
}

// [Literally, from: go/src/encoding/json/decode.go]
// getu4 decodes \uXXXX from the beginning of s, returning the hex value,
// or it returns -1.
func getu4(s []byte) rune {
	if len(s) < 6 || s[0] != '\\' || s[1] != 'u' {
		return -1
	}
	r, err := strconv.ParseUint(string(s[2:6]), 16, 64)
	if err != nil {
		return -1
	}
	return rune(r)
}
