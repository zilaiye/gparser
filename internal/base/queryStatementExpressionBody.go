package base

type QueryStatementExpressionBody struct {
	SourceCode
	FromStatement *FromStatement
	RegularBody   *RegularBody
}
