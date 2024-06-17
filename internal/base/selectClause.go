package base

type SelectClause struct {
	SourceCode
	All_distinct     *All_distinct
	SelectList       *SelectList
	SelectTrfmClause *SelectTrfmClause
	TrfmClause       *TrfmClause
}
