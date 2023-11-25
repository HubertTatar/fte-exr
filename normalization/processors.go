package normalization

import (
	snowball "github.com/kljensen/snowball/english"
	"strings"
)

// https://en.wikipedia.org/wiki/Most_common_words_in_English
var stopwords = map[string]struct{}{
	"a": {}, "and": {}, "be": {}, "have": {}, "i": {}, "am": {},
	"in": {}, "of": {}, "that": {}, "the": {}, "to": {}, "is": {},
	"it": {}, "for": {}, "not": {}, "on": {}, "with": {}, "are": {},
	"he": {}, "as": {}, "you": {}, "do": {}, "at": {}, "she": {},
	"this": {}, "but": {}, "his": {}, "by": {}, "from": {}, "they": {},
}

func toLowercase(tokens []string) []string {
	r := make([]string, len(tokens))
	for i, token := range tokens {
		r[i] = strings.ToLower(token)
	}
	return r
}

// https://en.wikipedia.org/wiki/Stop_word
func filterStopWords(tokens []string) []string {
	r := make([]string, 0, len(tokens))
	for _, token := range tokens {
		if _, ok := stopwords[token]; !ok {
			r = append(r, token)
		}
	}
	return r
}

// https://en.wikipedia.org/wiki/Stemming
func stemmer(tokens []string) []string {
	r := make([]string, len(tokens))
	for i, token := range tokens {
		r[i] = snowball.Stem(token, false)
	}
	return r
}
