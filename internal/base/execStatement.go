package base

type ExecStatement struct {
	SourceCode
	QueryStatementExpression *QueryStatementExpression
	LoadStatement            *LoadStatement
	ExportStatement          *ExportStatement
	ImportStatement          *ImportStatement
	ReplDumpStatement        *ReplDumpStatement
	ReplLoadStatement        *ReplLoadStatement
	ReplStatusStatement      *ReplStatusStatement
	DdlStatement             *DdlStatement
	DeleteStatement          *DeleteStatement
	UpdateStatement          *UpdateStatement
	SqlTransactionStatement  *SqlTransactionStatement
	MergeStatement           *MergeStatement
	PrepareStatement         *PrepareStatement
	ExecuteStatement         *ExecuteStatement
}
