package base

type SelectStatement struct {
	SourceCode
	AtomSelectStatement  *AtomSelectStatement
	SetOpSelectStatement *SetOpSelectStatement
	OrderByClause        *OrderByClause
	ClusterByClause      *ClusterByClause
	DistributeByClause   *DistributeByClause
	SortByClause         *SortByClause
	LimitClause          *LimitClause
}
