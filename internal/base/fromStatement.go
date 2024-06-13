package base

type FromStatement struct {
	SingleFromStatement *SingleFromStatement
	URPair              *[]URPair
}

type URPair struct {
	SetOperator         *SetOperator
	SingleFromStatement *SingleFromStatement
}
