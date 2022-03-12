package token

import (
	"strconv"
	"testing"
)

func TestTokenizerError(t *testing.T) {

	input := "@"

	tokenizer := NewTokenizer(input)
	_, err := tokenizer.TokenizeSingle()
	if err == nil {
		t.Error("no error returned")
	}
	/* same as writing
	_, ok := err.(*TokenizerError)
	if !ok {
		t.Error("unexpected error returned")
	}
	*/
	if _, ok := err.(*TokenizerError); !ok {
		t.Error("unexpected error returned")
	}

}

func TestTokenizeNumber(t *testing.T) {
	input := "100"
	expected, _ := strconv.Atoi(input)

	tokenizer := NewTokenizer(input)
	tokens := tokenizer.Tokenize()
	if !tokens[0].Equals(&NumberToken{expected}) {
		t.Error()
	}

}

func TestTokenizeVariable(t *testing.T) {
	input := "var123"

	tokenizer := NewTokenizer(input)
	tokens := tokenizer.Tokenize()
	if !tokens[0].Equals(&VariableToken{input}) {
		t.Error()
	}
}

func TestTokenizeString(t *testing.T) {
	input := "\"this is a String\""

	tokenizer := NewTokenizer(input)
	tokens := tokenizer.Tokenize()
	if !tokens[0].Equals(&QuoteStringToken{input}) {
		t.Error()
	}
}

func TestAllKeywords(t *testing.T) {
	type keywordTest struct {
		Input          string
		ExpectedResult Token
	}

	keywordTests := []keywordTest{
		{"if", &IfToken{}},
		{"else", &ElseToken{}},
		{"true", &TrueToken{}},
		{"false", &FalseToken{}},
		{"for", &ForToken{}},
		{"fn", &FnToken{}},
		{"bool", &BoolToken{}},
		{"const", &ConstToken{}},
		{"float", &FloatToken{}},
		{"int", &IntToken{}},
		{"print", &PrintToken{}},
		{"return", &ReturnToken{}},
		{"str", &StringToken{}},
		{"var", &VarToken{}}}

	for _, test := range keywordTests {
		t.Run(test.Input, func(t *testing.T) {
			tokenizer := NewTokenizer(test.Input)
			tokens := tokenizer.Tokenize()
			actual := tokens[0]
			if !actual.Equals(test.ExpectedResult) {
				t.Errorf("expected token: \"%s\", got: %s", test.ExpectedResult, actual)
			}
			actualString := actual.String()
			if actualString != test.Input {
				t.Errorf("String() returned wrong output, expected: \"%s, got: \"%s\"\"", test.Input, actualString)
			}
		})
	}

}

func TestAllSymbols(t *testing.T) {
	type symbolTest struct {
		Input          string
		ExpectedResult Token
	}

	symbolTests := []symbolTest{
		{"(", &LeftParenToken{}},
		{")", &RightParenToken{}},
		{"[", &LeftBracketToken{}},
		{"]", &RightBracketToken{}},
		{"{", &LeftCurlyToken{}},
		{"}", &RightCurlyToken{}},
		{"+", &PlusToken{}},
		{"-", &MinusToken{}},
		{"*", &MultToken{}},
		{"/", &DivToken{}},
		{"%", &ModToken{}},
		{"<", &LesserToken{}},
		{">", &GreaterToken{}},
		{".", &DotToken{}},
		{",", &CommaToken{}},
		{"!", &NegateToken{}},
		{"=", &AssignmentToken{}},
		{"==", &EqualsToken{}},
		{"->", &ArrowToken{}},
		{"||", &OrToken{}},
		{"&&", &AndToken{}},
		{"|>", &PipeOperatorToken{}},
		{"!=", &NotEqualsToken{}}}

	for _, test := range symbolTests {
		t.Run(test.Input, func(t *testing.T) {
			tokenizer := NewTokenizer(test.Input)
			tokens := tokenizer.Tokenize()
			actual := tokens[0]
			if !actual.Equals(test.ExpectedResult) {
				t.Errorf("expected token: \"%s\", got: %s", test.ExpectedResult, actual)
			}
			actualString := actual.String()
			if actualString != test.Input {
				t.Errorf("String() returned wrong output, expected: \"%s, got: \"%s\"\"", test.Input, actualString)
			}
		})
	}
}
func testTokens(input string, token Token) bool {
	tokenizer := NewTokenizer(input)
	tokens := tokenizer.Tokenize()
	return tokens[0].Equals(token)
}
func TestTokenizeIf(t *testing.T) {
	isValid := testTokens("if", &IfToken{})
	if !isValid {
		t.Error()
	}
}
func TestTokenizeElse(t *testing.T) {
	isValid := testTokens("else", &ElseToken{})
	if !isValid {
		t.Error()
	}
}
func TestTokenizeTrue(t *testing.T) {
	isValid := testTokens("true", &TrueToken{})
	if !isValid {
		t.Error()
	}
}
func TestTokenizeFalse(t *testing.T) {
	isValid := testTokens("false", &FalseToken{})
	if !isValid {
		t.Error()
	}
}
func TestTokenizeFor(t *testing.T) {
	isValid := testTokens("for", &ForToken{})
	if !isValid {
		t.Error()
	}
}
func TestTokenizefn(t *testing.T) {
	isValid := testTokens("fn", &FnToken{})
	if !isValid {
		t.Error()
	}
}
func TestTokenizeBool(t *testing.T) {
	isValid := testTokens("bool", &BoolToken{})
	if !isValid {
		t.Error()
	}
}
func TestTokenizeConst(t *testing.T) {
	isValid := testTokens("const", &ConstToken{})
	if !isValid {
		t.Error()
	}
}
func TestTokenizeFloat(t *testing.T) {
	isValid := testTokens("float", &FloatToken{})
	if !isValid {
		t.Error()
	}
}
func TestTokenizeInt(t *testing.T) {
	isValid := testTokens("int", &IntToken{})
	if !isValid {
		t.Error()
	}
}
func TestTokenizePrint(t *testing.T) {
	isValid := testTokens("print", &PrintToken{})
	if !isValid {
		t.Error()
	}
}
func TestTokenizeReturn(t *testing.T) {
	isValid := testTokens("return", &ReturnToken{})
	if !isValid {
		t.Error()
	}
}
func TestTokenizeStr(t *testing.T) {
	isValid := testTokens("str", &StringToken{})
	if !isValid {
		t.Error()
	}
}
func TestTokenizeVar(t *testing.T) {
	isValid := testTokens("var", &StringToken{})
	if !isValid {
		t.Error()
	}
}
func TestTokenizeLeftParen(t *testing.T) {
	isValid := testTokens("(", &LeftParenToken{})
	if !isValid {
		t.Error()
	}
}

func TestTokenizeRightParen(t *testing.T) {
	isValid := testTokens(")", &RightParenToken{})
	if !isValid {
		t.Error()
	}
}

func TestTokenizeLeftBracket(t *testing.T) {
	isValid := testTokens("[", &LeftBracketToken{})
	if !isValid {
		t.Error()
	}
}

func TestTokenizeRightBracket(t *testing.T) {
	isValid := testTokens("]", &RightBracketToken{})
	if !isValid {
		t.Error()
	}
}

func TestTokenizeLeftCurly(t *testing.T) {
	isValid := testTokens("{", &LeftCurlyToken{})
	if !isValid {
		t.Error()
	}
}

func TestTokenizeRightCurly(t *testing.T) {
	isValid := testTokens("}", &RightCurlyToken{})
	if !isValid {
		t.Error()
	}
}

func TestTokenizePlus(t *testing.T) {
	isValid := testTokens("+", &PlusToken{})
	if !isValid {
		t.Error()
	}
}
func TestTokenizeMinus(t *testing.T) {
	isValid := testTokens("-", &MinusToken{})
	if !isValid {
		t.Error()
	}
}

func TestTokenizeMult(t *testing.T) {
	isValid := testTokens("*", &MultToken{})
	if !isValid {
		t.Error()
	}
}

func TestTokenizeDiv(t *testing.T) {
	isValid := testTokens("/", &DivToken{})
	if !isValid {
		t.Error()
	}
}

func TestTokenizeMod(t *testing.T) {
	isValid := testTokens("%", &ModToken{})
	if !isValid {
		t.Error()
	}
}

func TestTokenizeLess(t *testing.T) {
	isValid := testTokens("<", &LesserToken{})
	if !isValid {
		t.Error()
	}
}

func TestTokenizeGreater(t *testing.T) {
	isValid := testTokens(">", &GreaterToken{})
	if !isValid {
		t.Error()
	}
}

func TestTokenizeDot(t *testing.T) {
	isValid := testTokens(".", &DotToken{})
	if !isValid {
		t.Error()
	}
}

func TestTokenizeComma(t *testing.T) {
	isValid := testTokens(",", &CommaToken{})
	if !isValid {
		t.Error()
	}
}

func TestTokenizeNegate(t *testing.T) {
	isValid := testTokens("!", &NegateToken{})
	if !isValid {
		t.Error()
	}
}
func TestTokenizeAssignment(t *testing.T) {
	isValid := testTokens("=", &AssignmentToken{})
	if !isValid {
		t.Error()
	}
}

func TestTokenizeEquals(t *testing.T) {
	isValid := testTokens("==", &EqualsToken{})
	if !isValid {
		t.Error()
	}
}

func TestTokenizeArrow(t *testing.T) {
	isValid := testTokens("->", &ArrowToken{})
	if !isValid {
		t.Error()
	}
}

func TestTokenizeOr(t *testing.T) {
	isValid := testTokens("||", &OrToken{})
	if !isValid {
		t.Error()
	}
}

func TestTokenizeAnd(t *testing.T) {
	isValid := testTokens("&&", &AndToken{})
	if !isValid {
		t.Error()
	}
}

func TestTokenizePipe(t *testing.T) {
	isValid := testTokens("|>", &PipeOperatorToken{})
	if !isValid {
		t.Error()
	}
}

func TestTokenizeNotEquals(t *testing.T) {
	isValid := testTokens("!=", &NotEqualsToken{})
	if !isValid {
		t.Error()
	}
}
