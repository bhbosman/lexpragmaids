%{
package lexpragmaids

import "strings"
import yaccPragmaIdsToken "github.com/bhbosman/yaccpragmaids"

	const (
		INITIAL = iota
	)


func (self *Handler) readLexem() (*Lexem, error) {

%}

%yyc self.current
%yyn self.current = self.GetChar()
%yyt self.startCond

D  [0-9]+
E  [eE][-+]?{D}
F  {D}"."{D}?{E}?|{D}{E}?|"."{D}{E}?

%%
    self.buf = self.buf[:0]

IDL return NewLexemStringValue(yaccPragmaIdsToken.RWIdl, self.TokenName, string(self.buf))
LOCAL return NewLexemStringValue(yaccPragmaIdsToken.RWLocal, self.TokenName, string(self.buf))
RMI return NewLexemStringValue(yaccPragmaIdsToken.RWRmi, self.TokenName, string(self.buf))
DCE return NewLexemStringValue(yaccPragmaIdsToken.RWDce, self.TokenName, string(self.buf))

":"              return NewLexemNoValue(int(':'), self.TokenName)
"."              return NewLexemNoValue(int('.'), self.TokenName)
[a-zA-Z0-9-/;]+
		s := strings.TrimSpace(string(self.buf[:len(self.buf)]))
		return NewLexemStringValue(yaccPragmaIdsToken.Value, self.TokenName, s)
%%
    return NewLexemEofValue()
}

