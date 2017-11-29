package rtokenizer

func removeSubstring(s string, from int, to int) string {
	return s[:from] + s[to:]
}

func filterTokens(tokens []*Token, tType TokenType) []*Token {
	trimmed := []*Token{}
	for _, token := range tokens {
		if token.Type != Space {
			trimmed = append(trimmed, token)
		}
	}
	return trimmed
}
