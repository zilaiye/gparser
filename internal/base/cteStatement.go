package base

type CteStatement struct {
	Id_                      *Id_
	ColumnNameList           []*ColumnNameList
	QueryStatementExpression *QueryStatementExpression
}
