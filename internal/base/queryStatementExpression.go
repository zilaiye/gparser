package base

type QueryStatementExpression struct {
	SourceCode
	WithClause                   *WithClause
	QueryStatementExpressionBody *QueryStatementExpressionBody
}
