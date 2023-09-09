package lang

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type lexerTestCase struct {
	description string
	input       string
	tokens      []Token
}

func TestScanTokens(t *testing.T) {
	for _, test := range makeLexerTestCases() {
		t.Run(test.description, func(t *testing.T) {
			lexer := NewLexer(test.input)

			have := test.tokens
			want := lexer.ScanTokens()

			assert.Equal(t, want, have)
		})
	}
}

func makeLexerTestCases() []lexerTestCase {
	// ...it isn't nice that this is written by hand, however for these test cases
	// it is more than fine, and in a way proves that the results are predictable.
	return []lexerTestCase{
		{
			description: "a let statement assigning a number",
			input:       "let some_number = 10;",
			tokens: []Token{
				{
					TokenType: TokenTypeLet,
					Token:     "let",
					Where:     1,
				},
				{
					TokenType: TokenTypeIdentifier,
					Token:     "some_number",
					Where:     1,
				},
				{
					TokenType: TokenTypeEqual,
					Token:     "=",
					Where:     1,
				},
				{
					TokenType: TokenTypeNumber,
					Token:     "10",
					Value:     float64(10), // ...as our number values are 64-bit floating point types, we should be explicit here.
					Where:     1,
				},
				{
					TokenType: TokenTypeSemi,
					Token:     ";",
					Where:     1,
				},
				{
					TokenType: TokenTypeEOF,
					Where:     1,
				},
			},
		},
		{
			description: "a let statement assigning a string",
			input:       `let some_string = "Hello, World.";`,
			tokens: []Token{
				{
					TokenType: TokenTypeLet,
					Token:     "let",
					Where:     1,
				},
				{
					TokenType: TokenTypeIdentifier,
					Token:     "some_string",
					Where:     1,
				},
				{
					TokenType: TokenTypeEqual,
					Token:     "=",
					Where:     1,
				},
				{
					TokenType: TokenTypeString,
					Token:     `"Hello, World."`,
					Value:     "Hello, World.",
					Where:     1,
				},
				{
					TokenType: TokenTypeSemi,
					Token:     ";",
					Where:     1,
				},
				{
					TokenType: TokenTypeEOF,
					Where:     1,
				},
			},
		},
		{
			description: "a function declaration",
			input: `def some_func() {
				return 3.14;
			}`,
			tokens: []Token{
				{
					TokenType: TokenTypeDef,
					Token:     "def",
					Where:     1,
				},
				{
					TokenType: TokenTypeIdentifier,
					Token:     "some_func",
					Where:     1,
				},
				{
					TokenType: TokenTypeLeftParen,
					Token:     "(",
					Where:     1,
				},
				{
					TokenType: TokenTypeRightParen,
					Token:     ")",
					Where:     1,
				},
				{
					TokenType: TokenTypeLeftBrace,
					Token:     "{",
					Where:     1,
				},
				{
					TokenType: TokenTypeReturn,
					Token:     "return",
					Where:     2,
				},
				{
					TokenType: TokenTypeNumber,
					Token:     "3.14",
					Value:     float64(3.14),
					Where:     2,
				},
				{
					TokenType: TokenTypeSemi,
					Token:     ";",
					Where:     2,
				},
				{
					TokenType: TokenTypeRightBrace,
					Token:     "}",
					Where:     3,
				},
				{
					TokenType: TokenTypeEOF,
					Where:     3,
				},
			},
		},
	}
}
