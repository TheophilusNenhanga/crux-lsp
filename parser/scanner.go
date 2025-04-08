package parser

type TokenType int

const (
	TOKEN_LEFT_PARN TokenType = iota
	TOKEN_RIGHT_PARN
	TOKEN_LEFT_BRACE
	TOKEN_RIGHT_BRACE
	TOKEN_LEFT_SQUARE
	TOKEN_RIGHT_SQUARE
	TOKEN_COMMA
	TOKEN_DOT
	TOKEN_MINUS
	TOKEN_PLUS
	TOKEN_SEMICOLON
	TOKEN_SLASH
	TOKEN_BACKSLASH
	TOKEN_STAR
	TOKEN_STAR_STAR
	TOKEN_PERCENT
	TOKEN_COLON
	TOKEN_SINGLE_QUOTE
	TOKEN_DOUBLE_QUOTE
	TOKEN_BANG_EQUAL
	TOKEN_EQUAL
	TOKEN_EQUAL_EQUAL
	TOKEN_GREATER
	TOKEN_GREATER_EQUAL
	TOKEN_LESS
	TOKEN_LESS_EQUAL
	TOKEN_LEFT_SHIFT
	TOKEN_RIGHT_SHIFT
	TOKEN_PLUS_EQUAL
	TOKEN_MINUS_EQUAL
	TOKEN_STAR_EQUAL
	TOKEN_SLASH_EQUAL
	TOKEN_IDENTIFIER
	TOKEN_STRING
	TOKEN_INT
	TOKEN_FLOAT
	TOKEN_AND
	TOKEN_NOT
	TOKEN_CLASS
	TOKEN_ELSE
	TOKEN_FALSE
	TOKEN_FOR
	TOKEN_FN
	TOKEN_IF
	TOKEN_NIL
	TOKEN_OR
	TOKEN_RETURN
	TOKEN_SUPER
	TOKEN_SELF
	TOKEN_TRUE
	TOKEN_LET
	TOKEN_WHILE
	TOKEN_ERROR
	TOKEN_BREAK
	TOKEN_CONTINUE
	TOKEN_USE
	TOKEN_FROM
	TOKEN_PUB
	TOKEN_AS
	TOKEN_EOF
	TOKEN_MATCH
	TOKEN_EQUAL_ARROW
	TOKEN_OK
	TOKEN_ERR
	TOKEN_DEFAULT
	TOKEN_GIVE
)

type Token struct {
	TokenType  TokenType
	Lexeme     string
	LineNumber int
}

type Scanner struct {
	source  string
	start   int
	current int
	line    int
}

func NewScanner(source string) *Scanner {
	return &Scanner{
		source:  source,
		start:   0,
		current: 0,
		line:    1,
	}
}

func (scanner *Scanner) Advance() byte {
	char := scanner.source[scanner.current]
	scanner.current++
	return char
}

func (scanner *Scanner) Peek() byte {
	return scanner.source[scanner.current]
}

func (scanner *Scanner) IsAtEnd() bool {
	return scanner.current >= len(scanner.source)
}

func (scanner *Scanner) PeekNext() byte {
	if scanner.IsAtEnd() {
		return 0
	}
	return scanner.source[scanner.current+1]
}

func IsIdentifierStarter(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || c == '_' || c == '$'
}

func IsAlpha(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || c == '_'
}

func (scanner *Scanner) SkipWhiteSpace() {
	for {
		char := scanner.Peek()
		switch char {
		case ' ', '\t', '\r':
			scanner.Advance()
			break
		case '\n':
			scanner.line++
			scanner.Advance()
			break
		case '/':
			if scanner.PeekNext() == '/' {
				for scanner.Peek() != '\n' && scanner.IsAtEnd() {
					scanner.Advance()
				}
			} else {
				return
			}
			break
		default:
			return
		}
	}
}
