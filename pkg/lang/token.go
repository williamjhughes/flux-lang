package lang

import "fmt"

type Token struct {
	TokenType TokenType
	Token     string
	Value     interface{}
	Where     int
}

func (t Token) String() string {
	return fmt.Sprintf("%s | %s = %v", t.TokenType, t.Token, t.Value)
}
