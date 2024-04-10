package iteration

import (
	"strings"
	"unicode"
)


func Repeat(character string, repetitons int) string {
  var repetition string
  for i := 0; i < repetitons; i++ {
    repetition += character 
  }

	return repetition
}


func Lowercase(text string) string {
  strings.ToLowerSpecial(unicode.TurkishCase,)

  return ""
}
