package token

import (
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

	tokenizer := NewTokenizer(input)
	tokens := tokenizer.Tokenize()
	if !tokens[0].Equals(&NumberToken{100}) {
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
		{"bool", &BoolToken{}},
		{"const", &ConstToken{}},
		{"int", &IntToken{}},
		{"print", &PrintToken{}},
		{"return", &ReturnToken{}},
		{"str", &StringToken{}}}

	for _, test := range keywordTests {
		t.Run(test.Input, func(t *testing.T) {
			tokenizer := NewTokenizer(test.Input)
			tokens := tokenizer.Tokenize()
			actual := tokens[0]
			if !actual.Equals(test.ExpectedResult) {
				t.Errorf("expected token: \"%s\", got: %s", test.ExpectedResult, actual)
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
		{"!", &NegateToken{}},
		{"=", &AssignmentToken{}},
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
		})
	}
}
