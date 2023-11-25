package normalization

func Normalize(tokens []string) []string {
	tokens = toLowercase(tokens)
	tokens = filterStopWords(tokens)
	tokens = stemmer(tokens)
	return tokens
}
