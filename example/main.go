package main

import (
	"fmt"

	"github.com/vanhtuan0409/rtokenizer"
)

func main() {
	t := rtokenizer.NewTokenizer(rtokenizer.Option{
		IgnoreSpaces:             true,
		IgnoreLineBreak:          true,
		UseBuiltInLineBreakToken: true,
		UseBuiltInSpaceToken:     true,
	})
	t.Add("number", `[0-9]*\.?[0-9]+`)
	t.Add("plus", `\+`)
	t.Add("minus", `-`)
	t.Add("multiply", `\*`)
	t.Add("divide", `/`)
	t.Add("openBracket", `\(`)
	t.Add("closeBracket", `\)`)

	tokens, err := t.Tokenize("1+2    * (3 + 4) - 5")
	if err != nil {
		fmt.Println(err)
	}

	for _, token := range tokens {
		fmt.Println(token.Type)
		fmt.Println(token.RawValue)
		fmt.Println("====")
	}
}
