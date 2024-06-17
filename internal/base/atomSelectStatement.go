package base

type AtomSelectStatement struct {
	SourceCode
	Query           *Query
	SelectStatement *SelectStatement
	ValuesSource *ValuesSource
}

type Query struct {
	SelectClause  *SelectClause
	FromClause    *FromClause
	WhereClause   *WhereClause
	GroupByClause *GroupByClause
	HavingClause  *HavingClause
	Window_clause *Window_clause
	QualifyClause *QualifyClause
}
