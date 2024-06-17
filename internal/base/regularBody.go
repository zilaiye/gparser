package base

type RegularBody struct {
	SourceCode
	InsertStatement *InsertStatement
	SelectStatement *SelectStatement
}

type InsertStatement struct {
	InsertClause    *InsertClause
	SelectStatement *SelectStatement
}
