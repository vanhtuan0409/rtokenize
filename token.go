package rtokenizer

// TokenType alias of string as TokenType
type TokenType string

// General-Predefined TokenType
const (
	Space     TokenType = "Space"
	LineBreak TokenType = "LineBreak"
)

// Token Represent of Token
type Token struct {
	Type     TokenType
	RawValue string
	Start    int
	End      int
}
