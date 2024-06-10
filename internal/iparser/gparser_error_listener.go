package iparser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
)

type GparserErrorListener struct {
	antlr.ErrorListener
	GrammarErrors []string
}

func NewGparserErrorListener() *GparserErrorListener {
	//listener := new(GparserErrorListener)
	//listener.GrammarErrors = make([]string, 0)
	//return listener
	return new(GparserErrorListener)
}

func (d *GparserErrorListener) SyntaxError(_ antlr.Recognizer, _ interface{}, line, col int, msg string, _ antlr.RecognitionException) {
	errMsg := fmt.Sprintf("line %d:%d %s", line, col, msg)
	d.GrammarErrors = append(d.GrammarErrors, errMsg)
}

func (d *GparserErrorListener) ReportAmbiguity(_ antlr.Parser, _ *antlr.DFA, _, _ int, _ bool, _ *antlr.BitSet, _ *antlr.ATNConfigSet) {
}

func (d *GparserErrorListener) ReportAttemptingFullContext(_ antlr.Parser, _ *antlr.DFA, _, _ int, _ *antlr.BitSet, _ *antlr.ATNConfigSet) {
}

func (d *GparserErrorListener) ReportContextSensitivity(_ antlr.Parser, _ *antlr.DFA, _, _, _ int, _ *antlr.ATNConfigSet) {
}
