package strlib

import (
	"bytes"
	"strings"
	"unicode"
	"unicode/utf8"
)

// SnakeCase converts camel case to snake case.
func SnakeCase(s string) string {
	return string(snakeCase(s, []rune{}, false))
}

func snakeCase(s string, rs []rune, wasLower bool) []rune {
	if len(s) == 0 {
		return rs
	}

	r, size := utf8.DecodeRuneInString(s)
	isUpper := unicode.IsUpper(r)
	if isUpper && wasLower {
		rs = append(rs, '_')
	}
	rs = append(rs, unicode.ToLower(r))

	return snakeCase(s[size:], rs, !isUpper)
}

// UpperCamelCase converts snake case to upper camel case
func UpperCamelCase(s string) string {
	ss := strings.Split(s, "_")
	if len(ss) == 1 {
		return strings.Title(s)
	}
	buf := bytes.Buffer{}
	for _, p := range ss {
		buf.WriteString(strings.Title(p))
	}
	return buf.String()
}

// LowerCamelCase converts snake case to lower camel case
func LowerCamelCase(s string) string {
	ss := strings.Split(s, "_")
	if len(ss) == 1 {
		return strings.ToLower(string(s[0])) + s[1:]
	}
	b := []byte{}
	for i, p := range ss {
		if i == 0 {
			b = append(b, strings.ToLower(string(p[0]))+p[1:]...)
			continue
		}
		b = append(b, strings.ToUpper(string(p[0]))+strings.ToLower(p[1:])...)
	}
	return string(b)
}

// Capitalize returns a copy of the string with its first character capitalized and the rest lowercased
func Capitalize(s string) string {
	return strings.ToUpper(string(s[0])) + strings.ToLower(s[1:])
}

// IsUpper returns true if all cased characters in the string are uppercase
func IsUpper(s string) bool {
	if len(s) < 1 {
		return false
	}
	return strings.ToUpper(s) == s
}

// IsLower returns true if all cased characters in the string are lowercase
func IsLower(s string) bool {
	if len(s) < 1 {
		return false
	}
	return strings.ToLower(s) == s
}

// Reverse returns reversed string
func Reverse(s string) string {
	length := len(s)
	r := make([]byte, length)
	for i := 0; i < length/2+1; i++ {
		r[i], r[length-i-1] = s[length-i-1], s[i]
	}
	return string(r)
}
