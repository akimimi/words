package words

import (
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
      } else {
        if r == '_' || r == ' ' {
          convert = true
        }
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

// Modelize converts string in to Camel-Case words
func Modelize (s string) string {
  return Camelize(s)
}

