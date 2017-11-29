package rtokenizer

import (
	"errors"
	"regexp"
)

// Possible error
var (
	ErrUnknowToken = errors.New("Encounter an unknow token")
)

// Option Option config for Tokenizer struct
type Option struct {
	IgnoreSpaces    bool
	IgnoreLineBreak bool
}

// Tokenizer Tokenizer
type Tokenizer interface {
	Add(TokenType, string) error
	Tokenize(string) ([]*Token, error)
}

// NewTokenizer Return Tokenizer Struct
func NewTokenizer(option Option) Tokenizer {
	t := &tokenizer{
		Matchers: []tokenMatcher{},
		Option:   option,
	}

	// Add default pattern matching
	t.Add(Space, `\s+|\t+`)
	t.Add(LineBreak, `\n`)
	return t
}

type tokenMatcher struct {
	Pattern *regexp.Regexp
	Type    TokenType
}

type tokenizer struct {
	Matchers []tokenMatcher
	Option   Option
}

// Add Add matching pattern for a specific TokenType
func (t *tokenizer) Add(tType TokenType, pattern string) error {
	r, err := regexp.Compile("^(" + pattern + ")")
	if err != nil {
		return err
	}
	t.Matchers = append(t.Matchers, tokenMatcher{
		Type:    tType,
		Pattern: r,
	})
	return nil
}

// Tokenize Tokenize a string into a serial of Tokens
func (t *tokenizer) Tokenize(s string) ([]*Token, error) {
	result := []*Token{}
	remainStr := s

	// Loop while we still have str to tokenize
	for remainStr != "" {
		isMatched := false
		tType := TokenType("")
		start := -1
		end := -1

		// Loop through all matcher and find longest match
		for _, matcher := range t.Matchers {
			indice := matcher.Pattern.FindIndex([]byte(remainStr))
			if len(indice) == 2 {
				isMatched = true
				bestMatchLength := end - start
				currentMatchLength := indice[1] - indice[0]
				if currentMatchLength >= bestMatchLength {
					tType = matcher.Type
					start, end = indice[0], indice[1]
				}
			}
		}

		if !isMatched {
			return nil, ErrUnknowToken
		}

		// Append result with longest match
		result = append(result, &Token{
			Type:     tType,
			Start:    start,
			End:      end,
			RawValue: remainStr[start:end],
		})
		remainStr = removeSubstring(remainStr, start, end)
	}

	// Filter result
	if t.Option.IgnoreSpaces {
		return filterTokens(result, Space), nil
	}
	if t.Option.IgnoreLineBreak {
		return filterTokens(result, LineBreak), nil
	}

	return result, nil
}
