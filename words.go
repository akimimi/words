package words

import (
	"bytes"
	"github.com/gertd/go-pluralize"
	"strings"
	"unicode"
)

// Camelize converts the slashed words into Camel-Case words
func Camelize(s string) string {
	convert := true
	s = strings.Map(
		func(r rune) rune {
			if convert {
				convert = false
				return unicode.ToUpper(r)
			}

			if r == '_' || r == ' ' {
				convert = true
			}
			return r
		},
		s)
	return strings.Replace(s, "_", "", -1)
}

// Capitalize converts the first letter in a word to uppercase
func Capitalize(s string) string {
	first := true
	return strings.Map(
		func(r rune) rune {
			if first {
				first = false
				return unicode.ToUpper(r)
			}
			return r
		},
		s)
}

// Lowercase converts the first letter in a word to lowercase
func Lowercase(s string) string {
	first := true
	return strings.Map(
		func(r rune) rune {
			if first {
				first = false
				return unicode.ToLower(r)
			}
			return r
		},
		s)
}

// Modelize converts string in to Camel-Case words with the last word singularized
func Modelize(s string) string {
	words := Words(s)
	var buffer bytes.Buffer
	for i := 0; i < len(words); i++ {
		if i >= len(words)-1 {
			client := pluralize.NewClient()
			b := client.Singular(words[i])
			buffer.WriteString(Camelize(b))
		} else {
			buffer.WriteString(Camelize(words[i]))
		}
	}
	return buffer.String()
}

// Tableize converts string in to slashed lower case words with the last word pluralized
func Tableize(s string) string {
	words := Words(s)
	var buffer bytes.Buffer
	for i := 0; i < len(words); i++ {
		prefix := ""
		if i > 0 {
			prefix = "_"
		}
		buffer.WriteString(prefix)
		if i >= len(words)-1 {
			client := pluralize.NewClient()
			b := client.Plural(words[i])
			buffer.WriteString(Lowercase(b))
		} else {
			buffer.WriteString(Lowercase(words[i]))
		}
	}
	return buffer.String()
}

// Words split string into words by white space, '_', and upper case letter
func Words(s string) []string {
	ws := strings.Fields(s)
	results := make([]string, 0)
	for _, s := range ws {
		results = AppendString(results, words(s))
	}
	return results
}

func words(s string) []string {
	points := make([]int, 0)
	i := 0
	strings.Map(
		func(r rune) rune {
			if i == 0 || unicode.IsUpper(r) || unicode.IsSpace(r) || r == '_' {
				points = append(points, i)
			}
			i++
			return r
		},
		s)
	results := make([]string, 0)
	for i := 0; i < len(points); i++ {
		f := points[i]
		if s[f] == '_' || unicode.IsSpace(rune(s[f])) {
			f++
		}
		str := ""
		if i >= len(points)-1 {
			str = s[f:]
		} else {
			str = s[f:points[i+1]]
		}
		if str != "" {
			results = append(results, str)
		}
	}
	return results
}

// AppendString appends a string slice to the end of data which is also a string slice.
func AppendString(slice, data []string) []string {
	m := len(slice)
	n := m + len(data)
	if n > cap(slice) { // if necessary, reallocate
		// allocate double what's needed, for future growth.
		newSlice := make([]string, (n+1)*2)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0:n]
	copy(slice[m:n], data)
	return slice
}
