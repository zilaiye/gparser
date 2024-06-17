package graph

import "fmt"

type LineageNode struct {
	// 数据库的名称
	DB string
	// 表名
	TableName string
	// 列名
	ColumnName string
}

func (node LineageNode) Key() string {
	return fmt.Sprintf("%s.%s.%s", node.DB, node.TableName, node.ColumnName)
}
