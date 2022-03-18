package token

import (
	"fmt"
	"strconv"
	"unicode"

	"github.com/vSterlin/goscript/utils"
)

type Tokenizer struct {
	keywordMap   map[string]Token
	oneSymbolMap map[string]Token
	twoSymbolMap map[string]Token
	input        []byte
	inputPos     int
}

func NewTokenizer(input string) *Tokenizer {

	keywordMap := map[string]Token{
		"true":   &TrueToken{},
		"false":  &FalseToken{},
		"if":     &IfToken{},
		"else":   &ElseToken{},
		"float":  &FloatToken{},
		"for":    &ForToken{},
		"fn":     &FnToken{},
		"int":    &IntToken{},
		"str":    &StringToken{},
		"bool":   &BoolToken{},
		"const":  &ConstToken{},
		"print":  &PrintToken{},
		"return": &ReturnToken{},
		"string": &StringToken{},
		"var":    &VarToken{},
	}

	oneSymbolMap := map[string]Token{
		"(": &LeftParenToken{},
		")": &RightParenToken{},
		"{": &LeftCurlyToken{},
		"}": &RightCurlyToken{},
		"[": &LeftBracketToken{},
		"]": &RightBracketToken{},
		"+": &PlusToken{},
		"-": &MinusToken{},
		"*": &MultToken{},
		"/": &DivToken{},
		"^": &PowerToken{},
		"%": &ModToken{},
		",": &CommaToken{},
		"<": &LesserToken{},
		">": &GreaterToken{},

		".": &DotToken{},
		"!": &NegateToken{},
		"=": &AssignmentToken{},
	}

	twoSymbolMap := map[string]Token{
		"||": &OrToken{},
		"&&": &AndToken{},
		"|>": &PipeOperatorToken{},
		"!=": &NotEqualsToken{},
		"==": &EqualsToken{},
		"->": &ArrowToken{},
	}

	return &Tokenizer{
		keywordMap:   keywordMap,
		oneSymbolMap: oneSymbolMap,
		twoSymbolMap: twoSymbolMap,
		input:        []byte(input),
		inputPos:     0,
	}
}

func (t *Tokenizer) IsTokenString(input string) bool {
	// if string is in map returns true
	_, ok := t.keywordMap[input]
	return ok
}

func (t *Tokenizer) tryTokenizeNumber() *NumberToken {
	var initialInputPos int = t.inputPos
	var digits string = ""

	for t.inputPos < len(t.input) && unicode.IsDigit(rune(t.input[t.inputPos])) {
		digits += string(t.input[t.inputPos])
		t.inputPos++
	}

	// if we find a dot then we need to stop and let the float tokenizer handle it
	if t.inputPos < len(t.input) && string(t.input[t.inputPos]) == "." {
		t.inputPos = initialInputPos
		return nil
	}

	if len(digits) > 0 {
		num, err := strconv.Atoi(digits)
		if err != nil {
			panic(err)
		}
		return &NumberToken{num}
	} else {
		t.inputPos = initialInputPos
		return nil
	}
}

func (t *Tokenizer) tryTokenizeFloat() *FloatNumberToken {
	var initialInputPos int = t.inputPos
	var firstHalfDigits string = ""
	var secondHalfDigits string = ""

	// loop and store first half of digits until we find a dot
	for t.inputPos < len(t.input) && unicode.IsDigit(rune(t.input[t.inputPos])) {
		firstHalfDigits += string(t.input[t.inputPos])
		t.inputPos++
	}

	// if we found a dot, we can start storing the second half of the number
	if t.inputPos < len(t.input) && string(t.input[t.inputPos]) == "." {
		t.inputPos++
		for t.inputPos < len(t.input) && unicode.IsDigit(rune(t.input[t.inputPos])) {
			secondHalfDigits += string(t.input[t.inputPos])
			t.inputPos++
		}
	}

	if len(firstHalfDigits) > 0 || len(secondHalfDigits) > 0 {
		num, err := strconv.ParseFloat(firstHalfDigits, 64)
		secondNum, err := strconv.ParseFloat(secondHalfDigits, 64)
		if err != nil {
			panic(err)
		}
		return &FloatNumberToken{num, secondNum}
	} else {
		t.inputPos = initialInputPos
		return nil
	}

}

func (t *Tokenizer) tryTokenizeSymbol() Token {
	for key, token := range t.twoSymbolMap {
		if t.prefixCharsEqual(key) {
			t.inputPos += len(key)
			return token
		}
	}

	for key, token := range t.oneSymbolMap {
		if t.prefixCharsEqual(key) {
			t.inputPos += len(key)
			return token
		}
	}
	return nil
}

func (t *Tokenizer) tryTokenizeVariable() *VariableToken {
	var initialInputPos int = t.inputPos
	var name string = ""

	if unicode.IsLetter(rune(t.input[t.inputPos])) {
		name += string(t.input[t.inputPos])
		t.inputPos++
		for t.inputPos < len(t.input) && (unicode.IsLetter(rune(t.input[t.inputPos])) || unicode.IsDigit(rune(t.input[t.inputPos]))) {
			name += string(t.input[t.inputPos])
			t.inputPos++
		}
	} else {
		// reset position
		t.inputPos = initialInputPos
		return nil
	}

	if t.IsTokenString(name) {
		// reset position
		t.inputPos = initialInputPos
		return nil
	} else {
		return &VariableToken{name}
	}
}

func (t *Tokenizer) tryTokenizeString() *QuoteStringToken {
	var initialInputPos int = t.inputPos
	var stringValue string = ""

	// checks for opening quote
	if t.input[t.inputPos] == '"' {
		t.inputPos++
		// loop until we find the closing quote
		for t.inputPos < len(t.input) && t.input[t.inputPos] != '"' {
			stringValue += string(t.input[t.inputPos])
			t.inputPos++

			if t.inputPos >= len(t.input) {
				panic(fmt.Sprintf("Unclosed string: '%s', add closing quote", stringValue))
			}
		}
		// checks for closing quote then returns string token
		if t.inputPos < len(t.input) && t.input[t.inputPos] == '"' {
			t.inputPos++
			return &QuoteStringToken{stringValue}
		}
	}

	// reset position since we didn't find an opening quote
	t.inputPos = initialInputPos
	return nil
}

func (t *Tokenizer) skipWhitespace() {
	for t.inputPos < len(t.input) && unicode.IsSpace(rune(t.input[t.inputPos])) {
		t.inputPos++
	}
}

func (t *Tokenizer) prefixCharsEqual(probe string) bool {
	targetPos := t.inputPos
	probePos := 0

	for targetPos < len(t.input) && probePos < len(probe) && probe[probePos] == t.input[targetPos] {
		probePos++
		targetPos++
	}
	return probePos == len(probe)
}

func (t *Tokenizer) tryTokenizeKeyword() Token {
	for key, token := range t.keywordMap {
		if t.prefixCharsEqual(key) {
			t.inputPos += len(key)
			return token
		}
	}
	return nil
}

func (t *Tokenizer) TokenizeSingle() (Token, error) {
	var keywordToken Token
	var symToken Token
	var num *NumberToken
	var floatNum *FloatNumberToken
	var varToken *VariableToken
	var strToken *QuoteStringToken
	t.skipWhitespace()

	if t.inputPos >= len(t.input) {
		return nil, nil
	} else if num = t.tryTokenizeNumber(); num != nil {
		return num, nil
	} else if floatNum = t.tryTokenizeFloat(); floatNum != nil {
		return floatNum, nil
	} else if varToken = t.tryTokenizeVariable(); varToken != nil {
		return varToken, nil
	} else if strToken = t.tryTokenizeString(); strToken != nil {
		return strToken, nil
	} else if symToken = t.tryTokenizeSymbol(); symToken != nil {
		return symToken, nil
	} else if keywordToken = t.tryTokenizeKeyword(); keywordToken != nil {
		return keywordToken, nil
	} else {
		return nil, (NewTokenizerError(fmt.Sprintf(("invalid character %c at position %d"), t.input[t.inputPos], t.inputPos)))
	}

}

func (t *Tokenizer) Tokenize() []Token {

	tokens := []Token{}

	// will print an error in nicer way if it panics
	defer utils.HandleError()

	for t.inputPos < len(t.input) {
		current, err := t.TokenizeSingle()
		if err != nil {
			panic(err)
		}
		tokens = append(tokens, current)
	}

	return tokens
}
