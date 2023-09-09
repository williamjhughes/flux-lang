package lang

import (
	"log"
	"strconv"
)

type Lexer struct {
	tokens []Token

	source  string
	start   int
	current int
	where   int
}

func NewLexer(source string) *Lexer {
	tokens := []Token{}

	return &Lexer{
		tokens:  tokens,
		source:  source,
		start:   0,
		current: 0,
		where:   1,
	}
}

func (l *Lexer) ScanTokens() []Token {
	for !l.isAtEnd() {
		l.start = l.current
		l.scan()
	}

	l.tokens = append(l.tokens, Token{
		TokenType: TokenTypeEOF,
		Where:     l.where,
	})

	return l.tokens
}

func (l *Lexer) scan() {
	char := l.advance()

	switch char {
	case '(':
		l.pushToken(TokenTypeLeftParen)
	case '{':
		l.pushToken(TokenTypeLeftBrace)
	case ')':
		l.pushToken(TokenTypeRightParen)
	case '}':
		l.pushToken(TokenTypeRightBrace)
	case ',':
		l.pushToken(TokenTypeComma)
	case '.':
		l.pushToken(TokenTypePeriod)
	case '+':
		l.pushToken(TokenTypePlus)
	case '-':
		l.pushToken(TokenTypeMinus)
	case '/':
		l.pushToken(TokenTypeSlash)
	case '*':
		l.pushToken(TokenTypeAsterisk)
	case ';':
		l.pushToken(TokenTypeSemi)
	case '#':
		for l.check() != '\n' && !l.isAtEnd() {
			l.advance()
		}
	case '!':
		if l.match('=') {
			l.pushToken(TokenTypeBangEqual)
			break
		}
		l.pushToken(TokenTypeBang)
	case '=':
		if l.match('=') {
			l.pushToken(TokenTypeEqualEqual)
			break
		}
		l.pushToken(TokenTypeEqual)
	case '>':
		if l.match('=') {
			l.pushToken(TokenTypeGreaterEqual)
			break
		}
		l.pushToken(TokenTypeGreater)
	case '<':
		if l.match('=') {
			l.pushToken(TokenTypeLessEqual)
			break
		}
		l.pushToken(TokenTypeLess)
	case ' ', '\r', '\t':
	case '\n':
		l.where++
	case '"':
		l.makeString()
	default:
		if l.isAlpha(char) {
			l.makeIdentifier()
			break
		}
		if l.isDigit(char) {
			l.makeNumber()
			break
		}
		log.Fatalf("error: unexpected character on line %d", l.where)
	}
}

func (l *Lexer) isAtEnd() bool {
	return l.current >= len(l.source)
}

func (l *Lexer) advance() rune {
	if l.current < len(l.source) {
		char := l.source[l.current]

		l.current++
		return rune(char)
	}

	return rune(0)
}

func (l *Lexer) pushTokenWithValue(tokenType TokenType, value interface{}) {
	token := string(l.source[l.start:l.current])

	l.tokens = append(l.tokens, Token{
		tokenType,
		token,
		value,
		l.where,
	})
}

func (l *Lexer) pushToken(tokenType TokenType) {
	l.pushTokenWithValue(tokenType, nil)
}

func (l *Lexer) match(expected rune) bool {
	if l.isAtEnd() || l.source[l.current] != byte(expected) {
		return false
	}

	l.current++
	return true
}

func (l *Lexer) checkNext() rune {
	if l.current+1 >= len(l.source) {
		return '\x00'
	}

	return rune(l.source[l.current+1])
}

func (l *Lexer) check() rune {
	if l.isAtEnd() {
		return '\x00'
	}

	return rune(l.source[l.current])
}

func (l *Lexer) isAlpha(char rune) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || char == '_'
}

func (l *Lexer) isDigit(char rune) bool {
	return char >= '0' && char <= '9'
}

func (l *Lexer) isAlphaNumeric(char rune) bool {
	return l.isAlpha(char) || l.isDigit(char)
}

func (l *Lexer) makeIdentifier() {
	for l.isAlphaNumeric(l.check()) {
		l.advance()
	}

	value := l.source[l.start:l.current]
	token, isFound := tokenTypeKeywordsMap[value]

	if !isFound {
		token = TokenTypeIdentifier
	}

	l.pushToken(token)
}

func (l *Lexer) makeString() {
	for l.check() != '"' && !l.isAtEnd() {
		if l.check() == '\n' {
			l.where++
		}

		l.advance()
	}

	if l.isAtEnd() {
		log.Fatalf("error: unterminated string on line %d", l.where)
		return
	}

	l.advance()

	value := l.source[l.start+1 : l.current-1] // ...remove the surrounding quotes.
	l.pushTokenWithValue(TokenTypeString, value)
}

func (l *Lexer) makeNumber() {
	if l.isDigit(l.check()) {
		l.advance()
	}

	if l.check() == '.' && l.isDigit(l.checkNext()) {
		l.advance()

		for l.isDigit(l.check()) {
			l.advance()
		}
	}

	stringValue := l.source[l.start:l.current]
	numberValue, err := strconv.ParseFloat(stringValue, 64) // ...our numbers will always be 64-bit floating point values.

	if err != nil {
		log.Fatalf("error: failed to parse number %s", err)
	}

	l.pushTokenWithValue(TokenTypeNumber, numberValue)
}
