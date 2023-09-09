package lang

type TokenType byte

const (
	TokenTypeLeftParen TokenType = iota
	TokenTypeLeftBrace
	TokenTypeRightParen
	TokenTypeRightBrace
	TokenTypeComma
	TokenTypePeriod
	TokenTypePlus
	TokenTypeMinus
	TokenTypeSlash
	TokenTypeAsterisk
	TokenTypeSemi

	TokenTypeBang
	TokenTypeBangEqual
	TokenTypeEqual
	TokenTypeEqualEqual
	TokenTypeGreater
	TokenTypeGreaterEqual
	TokenTypeLess
	TokenTypeLessEqual

	TokenTypeIdentifier
	TokenTypeString
	TokenTypeNumber

	TokenTypeAnd
	TokenTypeOr
	TokenTypeLet
	TokenTypeIf
	TokenTypeElse
	TokenTypeFor
	TokenTypeWhile
	TokenTypeDef
	TokenTypeReturn
	TokenTypeClass
	TokenTypeOf
	TokenTypeSelf
	TokenTypeSuper

	TokenTypeNone
	TokenTypeTrue
	TokenTypeFalse
	TokenTypePrint

	TokenTypeEOF
)

var tokenTypeStringMap = map[TokenType]string{
	TokenTypeLeftParen:    "TOKEN_TYPE_LEFT_PAREN",
	TokenTypeLeftBrace:    "TOKEN_TYPE_LEFT_BRACE",
	TokenTypeRightParen:   "TOKEN_TYPE_RIGHT_PAREN",
	TokenTypeRightBrace:   "TOKEN_TYPE_RIGHT_BRACE",
	TokenTypeComma:        "TOKEN_TYPE_COMMA",
	TokenTypePeriod:       "TOKEN_TYPE_PERIOD",
	TokenTypePlus:         "TOKEN_TYPE_PLUS",
	TokenTypeMinus:        "TOKEN_TYPE_MINUS",
	TokenTypeSlash:        "TOKEN_TYPE_SLASH",
	TokenTypeAsterisk:     "TOKEN_TYPE_ASTERISK",
	TokenTypeSemi:         "TOKEN_TYPE_SEMI",
	TokenTypeBang:         "TOKEN_TYPE_BANG",
	TokenTypeBangEqual:    "TOKEN_TYPE_BANG_EQUAL",
	TokenTypeEqual:        "TOKEN_TYPE_EQUAL",
	TokenTypeEqualEqual:   "TOKEN_TYPE_EQUAL_EQUAL",
	TokenTypeGreater:      "TOKEN_TYPE_GREATER",
	TokenTypeGreaterEqual: "TOKEN_TYPE_GREATER_EQUAL",
	TokenTypeLess:         "TOKEN_TYPE_LESS",
	TokenTypeLessEqual:    "TOKEN_TYPE_LESS_EQUAL",
	TokenTypeIdentifier:   "TOKEN_TYPE_IDENTIFIER",
	TokenTypeString:       "TOKEN_TYPE_STRING",
	TokenTypeNumber:       "TOKEN_TYPE_NUMBER",
	TokenTypeAnd:          "TOKEN_TYPE_AND",
	TokenTypeOr:           "TOKEN_TYPE_OR",
	TokenTypeLet:          "TOKEN_TYPE_LET",
	TokenTypeIf:           "TOKEN_TYPE_IF",
	TokenTypeElse:         "TOKEN_TYPE_ELSE",
	TokenTypeFor:          "TOKEN_TYPE_FOR",
	TokenTypeWhile:        "TOKEN_TYPE_WHILE",
	TokenTypeDef:          "TOKEN_TYPE_DEF",
	TokenTypeReturn:       "TOKEN_TYPE_RETURN",
	TokenTypeClass:        "TOKEN_TYPE_CLASS",
	TokenTypeOf:           "TOKEN_TYPE_OF",
	TokenTypeSelf:         "TOKEN_TYPE_SELF",
	TokenTypeSuper:        "TOKEN_TYPE_SUPER",
	TokenTypeNone:         "TOKEN_TYPE_NONE",
	TokenTypeTrue:         "TOKEN_TYPE_TRUE",
	TokenTypeFalse:        "TOKEN_TYPE_FALSE",
	TokenTypePrint:        "TOKEN_TYPE_PRINT",
	TokenTypeEOF:          "TOKEN_TYPE_EOF",
}

var tokenTypeKeywordsMap = map[string]TokenType{
	"and":    TokenTypeAnd,
	"or":     TokenTypeOr,
	"let":    TokenTypeLet,
	"if":     TokenTypeIf,
	"else":   TokenTypeElse,
	"for":    TokenTypeFor,
	"while":  TokenTypeWhile,
	"def":    TokenTypeDef,
	"return": TokenTypeReturn,
	"class":  TokenTypeClass,
	"of":     TokenTypeOf,
	"self":   TokenTypeSelf,
	"super":  TokenTypeSuper,
	"none":   TokenTypeNone,
	"true":   TokenTypeTrue,
	"false":  TokenTypeFalse,
	"print":  TokenTypePrint,
}

func (t TokenType) String() string {
	str, ok := tokenTypeStringMap[t]

	if !ok {
		return "TOKEN_ILLEGAL"
	}

	return str
}

func (t TokenType) Index() int {
	return int(t)
}
