package lexpragmaids

type Lexem struct {
	TokenName string
	TypeKind  int
	Value     string
	Eof       bool
}

func (l Lexem) CheckTarget() bool {
	return false
}

type LexemCollection struct {
	LexemValues []Lexem
}

func NewLexemEofValue() (*Lexem, error) {

	return &Lexem{
		Eof: true,
	}, nil
}

func NewLexemNoValue(typeKind int, tokenName func(int) string) (*Lexem, error) {
	return &Lexem{
		TypeKind:  typeKind,
		TokenName: tokenName(typeKind),
		Eof:       false,
	}, nil
}

func NewLexemStringValue(typeKind int, tokenName func(int) string, stringValue string) (*Lexem, error) {
	return &Lexem{
		TypeKind:  typeKind,
		TokenName: tokenName(typeKind),
		Value:     stringValue,
		Eof:       false,
	}, nil
}
