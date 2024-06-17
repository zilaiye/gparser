package base

type Statement struct {
	SourceCode
	ExplainStatement *ExplainStatement
	ExecStatement    *ExecStatement
}
