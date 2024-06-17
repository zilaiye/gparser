package iparser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	parser "gparser/internal/iantlr/grammar"
)

type GHiveVisitor struct {
	// 需要一个结构存放 表与字段之间的关系
	// 需要一个结构存放字段与字段之间的关系
	*antlr.BaseParseTreeVisitor
	Columns     []string
	DB          string
	Table       string
	Alias       string
	targetDB    string
	targetTable string
}

func NewGHiveVisitor() *GHiveVisitor {
	return &GHiveVisitor{Columns: []string{}}
}

func (G *GHiveVisitor) Visit(tree antlr.ParseTree) interface{} {
	if tree == nil {
		return nil
	}
	//fmt.Println(tree.GetText())
	return tree.Accept(G)
}

func (G *GHiveVisitor) VisitChildren(node antlr.RuleNode) interface{} {
	fmt.Println(node.GetText())
	G.Visit(node.GetRuleContext())
	return nil
}
func (G *GHiveVisitor) VisitTerminal(node antlr.TerminalNode) interface{} {
	fmt.Println(node.GetText())

	return nil
}
func (G *GHiveVisitor) VisitErrorNode(node antlr.ErrorNode) interface{} {
	fmt.Println(node.GetText())
	return nil
}

func (G *GHiveVisitor) VisitStatement(ctx *parser.StatementContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	if ctx.ExecStatement() != nil {
		return G.Visit(ctx.ExecStatement())
	} else {
		return G.Visit(ctx.ExplainStatement())
	}
	return nil
}

func (G *GHiveVisitor) VisitExplainStatement(ctx *parser.ExplainStatementContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitExplainOption(ctx *parser.ExplainOptionContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitVectorizationOnly(ctx *parser.VectorizationOnlyContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitVectorizatonDetail(ctx *parser.VectorizatonDetailContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitExecStatement(ctx *parser.ExecStatementContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	if ctx.QueryStatementExpression() != nil {
		return G.Visit(ctx.QueryStatementExpression())
	}
	return nil
}

func (G *GHiveVisitor) VisitLoadStatement(ctx *parser.LoadStatementContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitReplicationClause(ctx *parser.ReplicationClauseContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitExportStatement(ctx *parser.ExportStatementContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitImportStatement(ctx *parser.ImportStatementContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitReplDumpStatement(ctx *parser.ReplDumpStatementContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitReplDbPolicy(ctx *parser.ReplDbPolicyContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitReplLoadStatement(ctx *parser.ReplLoadStatementContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitReplConfigs(ctx *parser.ReplConfigsContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitReplConfigsList(ctx *parser.ReplConfigsListContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitReplTableLevelPolicy(ctx *parser.ReplTableLevelPolicyContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitReplStatusStatement(ctx *parser.ReplStatusStatementContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitDdlStatement(ctx *parser.DdlStatementContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
	return nil
}

func (G *GHiveVisitor) VisitIfExists(ctx *parser.IfExistsContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitRestrictOrCascade(ctx *parser.RestrictOrCascadeContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitIfNotExists(ctx *parser.IfNotExistsContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitForce(ctx *parser.ForceContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitRewriteEnabled(ctx *parser.RewriteEnabledContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitRewriteDisabled(ctx *parser.RewriteDisabledContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitStoredAsDirs(ctx *parser.StoredAsDirsContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitOrReplace(ctx *parser.OrReplaceContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitCreateDatabaseStatement(ctx *parser.CreateDatabaseStatementContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitDbLocation(ctx *parser.DbLocationContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitDbManagedLocation(ctx *parser.DbManagedLocationContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitDbProperties(ctx *parser.DbPropertiesContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitDbPropertiesList(ctx *parser.DbPropertiesListContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitDbConnectorName(ctx *parser.DbConnectorNameContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitSwitchDatabaseStatement(ctx *parser.SwitchDatabaseStatementContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitDropDatabaseStatement(ctx *parser.DropDatabaseStatementContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitDatabaseComment(ctx *parser.DatabaseCommentContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitTruncateTableStatement(ctx *parser.TruncateTableStatementContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitDropTableStatement(ctx *parser.DropTableStatementContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitInputFileFormat(ctx *parser.InputFileFormatContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitTabTypeExpr(ctx *parser.TabTypeExprContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitPartTypeExpr(ctx *parser.PartTypeExprContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitTabPartColTypeExpr(ctx *parser.TabPartColTypeExprContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitDescStatement(ctx *parser.DescStatementContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitAnalyzeStatement(ctx *parser.AnalyzeStatementContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitFrom_in(ctx *parser.From_inContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitDb_schema(ctx *parser.Db_schemaContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitShowStatement(ctx *parser.ShowStatementContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitShowTablesFilterExpr(ctx *parser.ShowTablesFilterExprContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitLockStatement(ctx *parser.LockStatementContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitLockDatabase(ctx *parser.LockDatabaseContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitLockMode(ctx *parser.LockModeContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitUnlockStatement(ctx *parser.UnlockStatementContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitUnlockDatabase(ctx *parser.UnlockDatabaseContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitCreateRoleStatement(ctx *parser.CreateRoleStatementContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitDropRoleStatement(ctx *parser.DropRoleStatementContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitGrantPrivileges(ctx *parser.GrantPrivilegesContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitRevokePrivileges(ctx *parser.RevokePrivilegesContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitGrantRole(ctx *parser.GrantRoleContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitRevokeRole(ctx *parser.RevokeRoleContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitShowRoleGrants(ctx *parser.ShowRoleGrantsContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitShowRoles(ctx *parser.ShowRolesContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitShowCurrentRole(ctx *parser.ShowCurrentRoleContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitSetRole(ctx *parser.SetRoleContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitShowGrants(ctx *parser.ShowGrantsContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitShowRolePrincipals(ctx *parser.ShowRolePrincipalsContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitPrivilegeIncludeColObject(ctx *parser.PrivilegeIncludeColObjectContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitPrivilegeObject(ctx *parser.PrivilegeObjectContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitPrivObject(ctx *parser.PrivObjectContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitPrivObjectCols(ctx *parser.PrivObjectColsContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitPrivilegeList(ctx *parser.PrivilegeListContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitPrivlegeDef(ctx *parser.PrivlegeDefContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitPrivilegeType(ctx *parser.PrivilegeTypeContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitPrincipalSpecification(ctx *parser.PrincipalSpecificationContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitPrincipalName(ctx *parser.PrincipalNameContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitWithGrantOption(ctx *parser.WithGrantOptionContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitGrantOptionFor(ctx *parser.GrantOptionForContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitAdminOptionFor(ctx *parser.AdminOptionForContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitWithAdminOption(ctx *parser.WithAdminOptionContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitMetastoreCheck(ctx *parser.MetastoreCheckContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitResourceList(ctx *parser.ResourceListContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitResource(ctx *parser.ResourceContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitResourceType(ctx *parser.ResourceTypeContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitCreateFunctionStatement(ctx *parser.CreateFunctionStatementContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitDropFunctionStatement(ctx *parser.DropFunctionStatementContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitReloadFunctionsStatement(ctx *parser.ReloadFunctionsStatementContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitCreateMacroStatement(ctx *parser.CreateMacroStatementContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitDropMacroStatement(ctx *parser.DropMacroStatementContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitCreateViewStatement(ctx *parser.CreateViewStatementContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitViewPartition(ctx *parser.ViewPartitionContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitViewOrganization(ctx *parser.ViewOrganizationContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitViewClusterSpec(ctx *parser.ViewClusterSpecContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitViewComplexSpec(ctx *parser.ViewComplexSpecContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitViewDistSpec(ctx *parser.ViewDistSpecContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitViewSortSpec(ctx *parser.ViewSortSpecContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitDropViewStatement(ctx *parser.DropViewStatementContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitCreateMaterializedViewStatement(ctx *parser.CreateMaterializedViewStatementContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitDropMaterializedViewStatement(ctx *parser.DropMaterializedViewStatementContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitCreateScheduledQueryStatement(ctx *parser.CreateScheduledQueryStatementContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitDropScheduledQueryStatement(ctx *parser.DropScheduledQueryStatementContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitAlterScheduledQueryStatement(ctx *parser.AlterScheduledQueryStatementContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitAlterScheduledQueryChange(ctx *parser.AlterScheduledQueryChangeContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitScheduleSpec(ctx *parser.ScheduleSpecContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitExecutedAsSpec(ctx *parser.ExecutedAsSpecContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitDefinedAsSpec(ctx *parser.DefinedAsSpecContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitShowFunctionIdentifier(ctx *parser.ShowFunctionIdentifierContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitShowStmtIdentifier(ctx *parser.ShowStmtIdentifierContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitTableComment(ctx *parser.TableCommentContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitCreateTablePartitionSpec(ctx *parser.CreateTablePartitionSpecContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitCreateTablePartitionColumnTypeSpec(ctx *parser.CreateTablePartitionColumnTypeSpecContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitCreateTablePartitionColumnSpec(ctx *parser.CreateTablePartitionColumnSpecContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitPartitionTransformSpec(ctx *parser.PartitionTransformSpecContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitColumnNameTransformConstraint(ctx *parser.ColumnNameTransformConstraintContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitPartitionTransformType(ctx *parser.PartitionTransformTypeContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitTableBuckets(ctx *parser.TableBucketsContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitTableImplBuckets(ctx *parser.TableImplBucketsContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitTableSkewed(ctx *parser.TableSkewedContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitRowFormat(ctx *parser.RowFormatContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitRecordReader(ctx *parser.RecordReaderContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitRecordWriter(ctx *parser.RecordWriterContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitRowFormatSerde(ctx *parser.RowFormatSerdeContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitRowFormatDelimited(ctx *parser.RowFormatDelimitedContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitTableRowFormat(ctx *parser.TableRowFormatContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitTablePropertiesPrefixed(ctx *parser.TablePropertiesPrefixedContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitTableProperties(ctx *parser.TablePropertiesContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitTablePropertiesList(ctx *parser.TablePropertiesListContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitKeyValueProperty(ctx *parser.KeyValuePropertyContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitKeyProperty(ctx *parser.KeyPropertyContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitTableRowFormatFieldIdentifier(ctx *parser.TableRowFormatFieldIdentifierContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitTableRowFormatCollItemsIdentifier(ctx *parser.TableRowFormatCollItemsIdentifierContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitTableRowFormatMapKeysIdentifier(ctx *parser.TableRowFormatMapKeysIdentifierContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitTableRowFormatLinesIdentifier(ctx *parser.TableRowFormatLinesIdentifierContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitTableRowNullFormat(ctx *parser.TableRowNullFormatContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitTableFileFormat(ctx *parser.TableFileFormatContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitTableLocation(ctx *parser.TableLocationContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitColumnNameTypeList(ctx *parser.ColumnNameTypeListContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitColumnNameTypeOrConstraintList(ctx *parser.ColumnNameTypeOrConstraintListContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitColumnNameColonTypeList(ctx *parser.ColumnNameColonTypeListContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitColumnNameList(ctx *parser.ColumnNameListContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitColumnName(ctx *parser.ColumnNameContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitExtColumnName(ctx *parser.ExtColumnNameContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitColumnNameOrderList(ctx *parser.ColumnNameOrderListContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitColumnParenthesesList(ctx *parser.ColumnParenthesesListContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitEnableValidateSpecification(ctx *parser.EnableValidateSpecificationContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitEnableSpecification(ctx *parser.EnableSpecificationContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitValidateSpecification(ctx *parser.ValidateSpecificationContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitEnforcedSpecification(ctx *parser.EnforcedSpecificationContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitRelySpecification(ctx *parser.RelySpecificationContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitCreateConstraint(ctx *parser.CreateConstraintContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitAlterConstraintWithName(ctx *parser.AlterConstraintWithNameContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitTableLevelConstraint(ctx *parser.TableLevelConstraintContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitPkUkConstraint(ctx *parser.PkUkConstraintContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitCheckConstraint(ctx *parser.CheckConstraintContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitCreateForeignKey(ctx *parser.CreateForeignKeyContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitAlterForeignKeyWithName(ctx *parser.AlterForeignKeyWithNameContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitSkewedValueElement(ctx *parser.SkewedValueElementContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitSkewedColumnValuePairList(ctx *parser.SkewedColumnValuePairListContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitSkewedColumnValuePair(ctx *parser.SkewedColumnValuePairContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitSkewedColumnValues(ctx *parser.SkewedColumnValuesContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitSkewedColumnValue(ctx *parser.SkewedColumnValueContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitSkewedValueLocationElement(ctx *parser.SkewedValueLocationElementContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitOrderSpecification(ctx *parser.OrderSpecificationContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitNullOrdering(ctx *parser.NullOrderingContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitColumnNameOrder(ctx *parser.ColumnNameOrderContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitColumnNameCommentList(ctx *parser.ColumnNameCommentListContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitColumnNameComment(ctx *parser.ColumnNameCommentContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitOrderSpecificationRewrite(ctx *parser.OrderSpecificationRewriteContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitColumnRefOrder(ctx *parser.ColumnRefOrderContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitColumnNameType(ctx *parser.ColumnNameTypeContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitColumnNameTypeOrConstraint(ctx *parser.ColumnNameTypeOrConstraintContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitTableConstraint(ctx *parser.TableConstraintContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitColumnNameTypeConstraint(ctx *parser.ColumnNameTypeConstraintContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitColumnConstraint(ctx *parser.ColumnConstraintContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitForeignKeyConstraint(ctx *parser.ForeignKeyConstraintContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitColConstraint(ctx *parser.ColConstraintContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitAlterColumnConstraint(ctx *parser.AlterColumnConstraintContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitAlterForeignKeyConstraint(ctx *parser.AlterForeignKeyConstraintContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitAlterColConstraint(ctx *parser.AlterColConstraintContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitColumnConstraintType(ctx *parser.ColumnConstraintTypeContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitDefaultVal(ctx *parser.DefaultValContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitTableConstraintType(ctx *parser.TableConstraintTypeContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitConstraintOptsCreate(ctx *parser.ConstraintOptsCreateContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitConstraintOptsAlter(ctx *parser.ConstraintOptsAlterContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitColumnNameColonType(ctx *parser.ColumnNameColonTypeContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitColType(ctx *parser.ColTypeContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitColTypeList(ctx *parser.ColTypeListContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitType(ctx *parser.TypeContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitPrimitiveType(ctx *parser.PrimitiveTypeContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitListType(ctx *parser.ListTypeContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitStructType(ctx *parser.StructTypeContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitMapType(ctx *parser.MapTypeContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitUnionType(ctx *parser.UnionTypeContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitSetOperator(ctx *parser.SetOperatorContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitQueryStatementExpression(ctx *parser.QueryStatementExpressionContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	withClause := ctx.WithClause()
	query := ctx.QueryStatementExpressionBody()
	if withClause != nil {
		G.Visit(withClause)
	}
	if query != nil {
		G.Visit(query)
	}
	return nil
}

func (G *GHiveVisitor) VisitQueryStatementExpressionBody(ctx *parser.QueryStatementExpressionBodyContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	G.Visit(ctx.FromStatement())
	G.Visit(ctx.RegularBody())
	return nil
}

func (G *GHiveVisitor) VisitWithClause(ctx *parser.WithClauseContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitCteStatement(ctx *parser.CteStatementContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitFromStatement(ctx *parser.FromStatementContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitSingleFromStatement(ctx *parser.SingleFromStatementContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitRegularBody(ctx *parser.RegularBodyContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	G.Visit(ctx.InsertClause())
	G.Visit(ctx.SelectStatement())
	return nil
}

func (G *GHiveVisitor) VisitAtomSelectStatement(ctx *parser.AtomSelectStatementContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	G.Visit(ctx.GetS())
	G.Visit(ctx.GetF())
	G.Visit(ctx.GetW())
	G.Visit(ctx.GetH())
	G.Visit(ctx.GetWin())
	G.Visit(ctx.GetQ())
	G.Visit(ctx.SelectStatement())
	G.Visit(ctx.ValuesSource())
	return nil
}

func (G *GHiveVisitor) VisitSelectStatement(ctx *parser.SelectStatementContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	G.Visit(ctx.AtomSelectStatement())
	G.Visit(ctx.SetOpSelectStatement())
	G.Visit(ctx.OrderByClause())
	G.Visit(ctx.ClusterByClause())
	G.Visit(ctx.DistributeByClause())
	G.Visit(ctx.SortByClause())
	G.Visit(ctx.LimitClause())
	return nil
}

func (G *GHiveVisitor) VisitSetOpSelectStatement(ctx *parser.SetOpSelectStatementContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitSelectStatementWithCTE(ctx *parser.SelectStatementWithCTEContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitBody(ctx *parser.BodyContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitInsertClause(ctx *parser.InsertClauseContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	G.Visit(ctx.TableOrPartition())
	return nil
}

func (G *GHiveVisitor) VisitDestination(ctx *parser.DestinationContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitLimitClause(ctx *parser.LimitClauseContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitDeleteStatement(ctx *parser.DeleteStatementContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitColumnAssignmentClause(ctx *parser.ColumnAssignmentClauseContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitPrecedencePlusExpressionOrDefault(ctx *parser.PrecedencePlusExpressionOrDefaultContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitSetColumnsClause(ctx *parser.SetColumnsClauseContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitUpdateStatement(ctx *parser.UpdateStatementContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitSqlTransactionStatement(ctx *parser.SqlTransactionStatementContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitStartTransactionStatement(ctx *parser.StartTransactionStatementContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitTransactionMode(ctx *parser.TransactionModeContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitTransactionAccessMode(ctx *parser.TransactionAccessModeContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitIsolationLevel(ctx *parser.IsolationLevelContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitLevelOfIsolation(ctx *parser.LevelOfIsolationContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitCommitStatement(ctx *parser.CommitStatementContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitRollbackStatement(ctx *parser.RollbackStatementContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitSetAutoCommitStatement(ctx *parser.SetAutoCommitStatementContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitAbortTransactionStatement(ctx *parser.AbortTransactionStatementContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitAbortCompactionStatement(ctx *parser.AbortCompactionStatementContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitMergeStatement(ctx *parser.MergeStatementContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitWhenClauses(ctx *parser.WhenClausesContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitWhenNotMatchedClause(ctx *parser.WhenNotMatchedClauseContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitWhenMatchedAndClause(ctx *parser.WhenMatchedAndClauseContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitWhenMatchedThenClause(ctx *parser.WhenMatchedThenClauseContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitUpdateOrDelete(ctx *parser.UpdateOrDeleteContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitKillQueryStatement(ctx *parser.KillQueryStatementContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitCompactionId(ctx *parser.CompactionIdContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitCompactionPool(ctx *parser.CompactionPoolContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitCompactionType(ctx *parser.CompactionTypeContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitCompactionStatus(ctx *parser.CompactionStatusContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitAlterStatement(ctx *parser.AlterStatementContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitAlterTableStatementSuffix(ctx *parser.AlterTableStatementSuffixContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitAlterTblPartitionStatementSuffix(ctx *parser.AlterTblPartitionStatementSuffixContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitAlterStatementPartitionKeyType(ctx *parser.AlterStatementPartitionKeyTypeContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitAlterViewStatementSuffix(ctx *parser.AlterViewStatementSuffixContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitAlterMaterializedViewStatementSuffix(ctx *parser.AlterMaterializedViewStatementSuffixContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitAlterMaterializedViewSuffixRewrite(ctx *parser.AlterMaterializedViewSuffixRewriteContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitAlterMaterializedViewSuffixRebuild(ctx *parser.AlterMaterializedViewSuffixRebuildContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitAlterDatabaseStatementSuffix(ctx *parser.AlterDatabaseStatementSuffixContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitAlterDatabaseSuffixProperties(ctx *parser.AlterDatabaseSuffixPropertiesContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitAlterDatabaseSuffixSetOwner(ctx *parser.AlterDatabaseSuffixSetOwnerContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitAlterDatabaseSuffixSetLocation(ctx *parser.AlterDatabaseSuffixSetLocationContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitAlterDatabaseSuffixSetManagedLocation(ctx *parser.AlterDatabaseSuffixSetManagedLocationContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitAlterStatementSuffixRename(ctx *parser.AlterStatementSuffixRenameContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitAlterStatementSuffixAddCol(ctx *parser.AlterStatementSuffixAddColContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitAlterStatementSuffixAddConstraint(ctx *parser.AlterStatementSuffixAddConstraintContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitAlterStatementSuffixUpdateColumns(ctx *parser.AlterStatementSuffixUpdateColumnsContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitAlterStatementSuffixDropConstraint(ctx *parser.AlterStatementSuffixDropConstraintContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitAlterStatementSuffixRenameCol(ctx *parser.AlterStatementSuffixRenameColContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitAlterStatementSuffixUpdateStatsCol(ctx *parser.AlterStatementSuffixUpdateStatsColContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitAlterStatementSuffixUpdateStats(ctx *parser.AlterStatementSuffixUpdateStatsContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitAlterStatementChangeColPosition(ctx *parser.AlterStatementChangeColPositionContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitAlterStatementSuffixAddPartitions(ctx *parser.AlterStatementSuffixAddPartitionsContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitAlterStatementSuffixAddPartitionsElement(ctx *parser.AlterStatementSuffixAddPartitionsElementContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitAlterStatementSuffixTouch(ctx *parser.AlterStatementSuffixTouchContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitAlterStatementSuffixArchive(ctx *parser.AlterStatementSuffixArchiveContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitAlterStatementSuffixUnArchive(ctx *parser.AlterStatementSuffixUnArchiveContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitPartitionLocation(ctx *parser.PartitionLocationContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitAlterStatementSuffixDropPartitions(ctx *parser.AlterStatementSuffixDropPartitionsContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitAlterStatementSuffixProperties(ctx *parser.AlterStatementSuffixPropertiesContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitAlterViewSuffixProperties(ctx *parser.AlterViewSuffixPropertiesContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitAlterStatementSuffixSerdeProperties(ctx *parser.AlterStatementSuffixSerdePropertiesContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitTablePartitionPrefix(ctx *parser.TablePartitionPrefixContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitAlterStatementSuffixFileFormat(ctx *parser.AlterStatementSuffixFileFormatContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitAlterStatementSuffixClusterbySortby(ctx *parser.AlterStatementSuffixClusterbySortbyContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitAlterTblPartitionStatementSuffixSkewedLocation(ctx *parser.AlterTblPartitionStatementSuffixSkewedLocationContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitSkewedLocations(ctx *parser.SkewedLocationsContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitSkewedLocationsList(ctx *parser.SkewedLocationsListContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitSkewedLocationMap(ctx *parser.SkewedLocationMapContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitAlterStatementSuffixLocation(ctx *parser.AlterStatementSuffixLocationContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitAlterStatementSuffixSkewedby(ctx *parser.AlterStatementSuffixSkewedbyContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitAlterStatementSuffixExchangePartition(ctx *parser.AlterStatementSuffixExchangePartitionContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitAlterStatementSuffixRenamePart(ctx *parser.AlterStatementSuffixRenamePartContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitAlterStatementSuffixStatsPart(ctx *parser.AlterStatementSuffixStatsPartContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitAlterStatementSuffixMergeFiles(ctx *parser.AlterStatementSuffixMergeFilesContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitAlterStatementSuffixBucketNum(ctx *parser.AlterStatementSuffixBucketNumContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitBlocking(ctx *parser.BlockingContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitCompactPool(ctx *parser.CompactPoolContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitAlterStatementSuffixCompact(ctx *parser.AlterStatementSuffixCompactContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitAlterStatementSuffixSetOwner(ctx *parser.AlterStatementSuffixSetOwnerContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitAlterStatementSuffixSetPartSpec(ctx *parser.AlterStatementSuffixSetPartSpecContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitAlterStatementSuffixExecute(ctx *parser.AlterStatementSuffixExecuteContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitFileFormat(ctx *parser.FileFormatContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitAlterDataConnectorStatementSuffix(ctx *parser.AlterDataConnectorStatementSuffixContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitAlterDataConnectorSuffixProperties(ctx *parser.AlterDataConnectorSuffixPropertiesContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitAlterDataConnectorSuffixSetOwner(ctx *parser.AlterDataConnectorSuffixSetOwnerContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitAlterDataConnectorSuffixSetUrl(ctx *parser.AlterDataConnectorSuffixSetUrlContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitLikeTableOrFile(ctx *parser.LikeTableOrFileContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitCreateTableStatement(ctx *parser.CreateTableStatementContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitCreateDataConnectorStatement(ctx *parser.CreateDataConnectorStatementContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitDataConnectorComment(ctx *parser.DataConnectorCommentContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitDataConnectorUrl(ctx *parser.DataConnectorUrlContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitDataConnectorType(ctx *parser.DataConnectorTypeContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitDcProperties(ctx *parser.DcPropertiesContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitDropDataConnectorStatement(ctx *parser.DropDataConnectorStatementContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitTableAllColumns(ctx *parser.TableAllColumnsContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitTableOrColumn(ctx *parser.TableOrColumnContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	fmt.Println("visit TableOrColumn ")
	tableOrColumn := G.Visit(ctx.Id_()).(string)
	// G.Columns = append(G.Columns, tableOrColumn)
	return tableOrColumn
}

func (G *GHiveVisitor) VisitDefaultValue(ctx *parser.DefaultValueContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitExpressionList(ctx *parser.ExpressionListContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitAliasList(ctx *parser.AliasListContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitFromClause(ctx *parser.FromClauseContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	G.Visit(ctx.FromSource())
	return nil
}

func (G *GHiveVisitor) VisitFromSource(ctx *parser.FromSourceContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	G.Visit(ctx.JoinSource())
	for i := range ctx.AllUniqueJoinSource() {
		G.Visit(ctx.UniqueJoinSource(i))
	}
	return nil
}

func (G *GHiveVisitor) VisitAtomjoinSource(ctx *parser.AtomjoinSourceContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	G.Visit(ctx.TableSource())

	return nil
}

func (G *GHiveVisitor) VisitJoinSource(ctx *parser.JoinSourceContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	fmt.Println("visit JoinSource")
	G.Visit(ctx.AtomjoinSource())
	return nil
}

func (G *GHiveVisitor) VisitJoinSourcePart(ctx *parser.JoinSourcePartContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitUniqueJoinSource(ctx *parser.UniqueJoinSourceContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitUniqueJoinExpr(ctx *parser.UniqueJoinExprContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitUniqueJoinToken(ctx *parser.UniqueJoinTokenContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitJoinToken(ctx *parser.JoinTokenContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitLateralView(ctx *parser.LateralViewContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitTableAlias(ctx *parser.TableAliasContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitTableBucketSample(ctx *parser.TableBucketSampleContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitSplitSample(ctx *parser.SplitSampleContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitTableSample(ctx *parser.TableSampleContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitTableSource(ctx *parser.TableSourceContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	if ctx.GetAlias() != nil {
		alias := ctx.GetAlias().GetText()
		G.Alias = alias
	}
	G.Visit(ctx.TableName())
	return nil
}

func (G *GHiveVisitor) VisitAsOfClause(ctx *parser.AsOfClauseContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitUniqueJoinTableSource(ctx *parser.UniqueJoinTableSourceContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitTableName(ctx *parser.TableNameContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	fmt.Println("visit TableName")
	db := ctx.GetDb().GetText()
	tab := ctx.GetTab().GetText()
	G.DB = db
	G.Table = tab
	return nil
}

func (G *GHiveVisitor) VisitViewName(ctx *parser.ViewNameContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitSubQuerySource(ctx *parser.SubQuerySourceContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitPartitioningSpec(ctx *parser.PartitioningSpecContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitPartitionTableFunctionSource(ctx *parser.PartitionTableFunctionSourceContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitPartitionedTableFunction(ctx *parser.PartitionedTableFunctionContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitWhereClause(ctx *parser.WhereClauseContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitSearchCondition(ctx *parser.SearchConditionContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitValuesSource(ctx *parser.ValuesSourceContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitValuesClause(ctx *parser.ValuesClauseContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitValuesTableConstructor(ctx *parser.ValuesTableConstructorContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitValueRowConstructor(ctx *parser.ValueRowConstructorContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitFirstValueRowConstructor(ctx *parser.FirstValueRowConstructorContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitVirtualTableSource(ctx *parser.VirtualTableSourceContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitSelectClause(ctx *parser.SelectClauseContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	G.Visit(ctx.All_distinct())
	G.Visit(ctx.SelectList())
	G.Visit(ctx.SelectTrfmClause())
	G.Visit(ctx.TrfmClause())
	return nil
}

func (G *GHiveVisitor) VisitAll_distinct(ctx *parser.All_distinctContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitSelectList(ctx *parser.SelectListContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	for i := range ctx.AllSelectItem() {
		G.Visit(ctx.SelectItem(i))
	}
	return nil
}

func (G *GHiveVisitor) VisitSelectTrfmClause(ctx *parser.SelectTrfmClauseContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitSelectItem(ctx *parser.SelectItemContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	if ctx.KW_AS() != nil {
		// 增加一条边

	}
	G.Visit(ctx.TableAllColumns())
	G.Visit(ctx.Expression())
	G.Visit(ctx.KW_AS())
	for _, id := range ctx.AllId_() {
		G.Visit(id)
	}

	return nil
}

func (G *GHiveVisitor) VisitTrfmClause(ctx *parser.TrfmClauseContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitSelectExpression(ctx *parser.SelectExpressionContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitSelectExpressionList(ctx *parser.SelectExpressionListContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitWindow_clause(ctx *parser.Window_clauseContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitWindow_defn(ctx *parser.Window_defnContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitWindow_specification(ctx *parser.Window_specificationContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitWindow_frame(ctx *parser.Window_frameContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitWindow_range_expression(ctx *parser.Window_range_expressionContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitWindow_value_expression(ctx *parser.Window_value_expressionContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitWindow_frame_start_boundary(ctx *parser.Window_frame_start_boundaryContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitWindow_frame_boundary(ctx *parser.Window_frame_boundaryContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitGroupByClause(ctx *parser.GroupByClauseContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitGroupby_expression(ctx *parser.Groupby_expressionContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitGroupByEmpty(ctx *parser.GroupByEmptyContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitRollupStandard(ctx *parser.RollupStandardContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitRollupOldSyntax(ctx *parser.RollupOldSyntaxContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitGroupingSetExpression(ctx *parser.GroupingSetExpressionContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitGroupingSetExpressionMultiple(ctx *parser.GroupingSetExpressionMultipleContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitGroupingExpressionSingle(ctx *parser.GroupingExpressionSingleContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitHavingClause(ctx *parser.HavingClauseContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitQualifyClause(ctx *parser.QualifyClauseContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitHavingCondition(ctx *parser.HavingConditionContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitExpressionsInParenthesis(ctx *parser.ExpressionsInParenthesisContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitExpressionsNotInParenthesis(ctx *parser.ExpressionsNotInParenthesisContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitExpressionPart(ctx *parser.ExpressionPartContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitExpressionOrDefault(ctx *parser.ExpressionOrDefaultContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitFirstExpressionsWithAlias(ctx *parser.FirstExpressionsWithAliasContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitExpressionWithAlias(ctx *parser.ExpressionWithAliasContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitExpressions(ctx *parser.ExpressionsContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitColumnRefOrderInParenthesis(ctx *parser.ColumnRefOrderInParenthesisContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitColumnRefOrderNotInParenthesis(ctx *parser.ColumnRefOrderNotInParenthesisContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitOrderByClause(ctx *parser.OrderByClauseContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitClusterByClause(ctx *parser.ClusterByClauseContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitPartitionByClause(ctx *parser.PartitionByClauseContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitDistributeByClause(ctx *parser.DistributeByClauseContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitSortByClause(ctx *parser.SortByClauseContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitTrimFunction(ctx *parser.TrimFunctionContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitFunction_(ctx *parser.Function_Context) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitNull_treatment(ctx *parser.Null_treatmentContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitFunctionName(ctx *parser.FunctionNameContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitCastExpression(ctx *parser.CastExpressionContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitCaseExpression(ctx *parser.CaseExpressionContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitWhenExpression(ctx *parser.WhenExpressionContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitFloorExpression(ctx *parser.FloorExpressionContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitFloorDateQualifiers(ctx *parser.FloorDateQualifiersContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitExtractExpression(ctx *parser.ExtractExpressionContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitTimeQualifiers(ctx *parser.TimeQualifiersContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitConstant(ctx *parser.ConstantContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitPrepareStmtParam(ctx *parser.PrepareStmtParamContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitParameterIdx(ctx *parser.ParameterIdxContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitStringLiteralSequence(ctx *parser.StringLiteralSequenceContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitCharSetStringLiteral(ctx *parser.CharSetStringLiteralContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitDateLiteral(ctx *parser.DateLiteralContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitTimestampLiteral(ctx *parser.TimestampLiteralContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitTimestampLocalTZLiteral(ctx *parser.TimestampLocalTZLiteralContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitIntervalValue(ctx *parser.IntervalValueContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitIntervalLiteral(ctx *parser.IntervalLiteralContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitIntervalExpression(ctx *parser.IntervalExpressionContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitIntervalQualifiers(ctx *parser.IntervalQualifiersContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitExpression(ctx *parser.ExpressionContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	fmt.Println("visit Expression ")
	return G.Visit(ctx.PrecedenceOrExpression())
}

func (G *GHiveVisitor) VisitAtomExpression(ctx *parser.AtomExpressionContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	fmt.Println("visit AtomExpression")
	return G.Visit(ctx.TableOrColumn())
}

func (G *GHiveVisitor) VisitPrecedenceFieldExpression(ctx *parser.PrecedenceFieldExpressionContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	tableOrColumn := G.Visit(ctx.AtomExpression()).(string)
	if len(ctx.AllDOT()) > 0 {
		G.Columns = append(G.Columns, tableOrColumn+"."+ctx.Id_(0).GetText())
	} else {
		G.Columns = append(G.Columns, tableOrColumn)
	}

	return nil
}

func (G *GHiveVisitor) VisitPrecedenceUnaryOperator(ctx *parser.PrecedenceUnaryOperatorContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitPrecedenceUnaryPrefixExpression(ctx *parser.PrecedenceUnaryPrefixExpressionContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	G.Visit(ctx.PrecedenceFieldExpression())
	return nil
}

func (G *GHiveVisitor) VisitPrecedenceBitwiseXorOperator(ctx *parser.PrecedenceBitwiseXorOperatorContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitPrecedenceBitwiseXorExpression(ctx *parser.PrecedenceBitwiseXorExpressionContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	fmt.Println("visit PrecedenceBitwiseXorExpression")
	for i := range ctx.AllPrecedenceUnaryPrefixExpression() {
		G.Visit(ctx.PrecedenceUnaryPrefixExpression(i))
	}
	return nil
}

func (G *GHiveVisitor) VisitPrecedenceStarOperator(ctx *parser.PrecedenceStarOperatorContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitPrecedenceStarExpression(ctx *parser.PrecedenceStarExpressionContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	for i := range ctx.AllPrecedenceBitwiseXorExpression() {
		G.Visit(ctx.PrecedenceBitwiseXorExpression(i))
	}
	return nil
}

func (G *GHiveVisitor) VisitPrecedencePlusOperator(ctx *parser.PrecedencePlusOperatorContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitPrecedencePlusExpression(ctx *parser.PrecedencePlusExpressionContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	for i := range ctx.AllPrecedenceStarExpression() {
		G.Visit(ctx.PrecedenceStarExpression(i))
	}
	return nil
}

func (G *GHiveVisitor) VisitPrecedenceConcatenateOperator(ctx *parser.PrecedenceConcatenateOperatorContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitPrecedenceConcatenateExpression(ctx *parser.PrecedenceConcatenateExpressionContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	fmt.Println("visit PrecedenceConcatenateExpression")
	for i := range ctx.AllPrecedencePlusExpression() {
		G.Visit(ctx.PrecedencePlusExpression(i))
	}
	return nil
}

func (G *GHiveVisitor) VisitPrecedenceAmpersandOperator(ctx *parser.PrecedenceAmpersandOperatorContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitPrecedenceAmpersandExpression(ctx *parser.PrecedenceAmpersandExpressionContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	for i, _ := range ctx.AllPrecedenceConcatenateExpression() {
		G.Visit(ctx.PrecedenceConcatenateExpression(i))
	}
	return nil
}

func (G *GHiveVisitor) VisitPrecedenceBitwiseOrOperator(ctx *parser.PrecedenceBitwiseOrOperatorContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitPrecedenceBitwiseOrExpression(ctx *parser.PrecedenceBitwiseOrExpressionContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	for i, _ := range ctx.AllPrecedenceAmpersandExpression() {
		G.Visit(ctx.PrecedenceAmpersandExpression(i))
	}
	return nil
}

func (G *GHiveVisitor) VisitPrecedenceRegexpOperator(ctx *parser.PrecedenceRegexpOperatorContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitPrecedenceSimilarOperator(ctx *parser.PrecedenceSimilarOperatorContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitSubQueryExpression(ctx *parser.SubQueryExpressionContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitPrecedenceSimilarExpression(ctx *parser.PrecedenceSimilarExpressionContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	G.Visit(ctx.PrecedenceSimilarExpressionMain())
	if ctx.KW_EXISTS() != nil {
		// 这里后续是需要处理的
		G.Visit(ctx.SubQueryExpression())
	}
	return nil
}

func (G *GHiveVisitor) VisitPrecedenceSimilarExpressionMain(ctx *parser.PrecedenceSimilarExpressionMainContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	G.Visit(ctx.PrecedenceBitwiseOrExpression())
	return nil
}

func (G *GHiveVisitor) VisitPrecedenceSimilarExpressionPart(ctx *parser.PrecedenceSimilarExpressionPartContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitPrecedenceSimilarExpressionAtom(ctx *parser.PrecedenceSimilarExpressionAtomContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitPrecedenceSimilarExpressionQuantifierPredicate(ctx *parser.PrecedenceSimilarExpressionQuantifierPredicateContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitQuantifierType(ctx *parser.QuantifierTypeContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitPrecedenceSimilarExpressionIn(ctx *parser.PrecedenceSimilarExpressionInContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitPrecedenceSimilarExpressionPartNot(ctx *parser.PrecedenceSimilarExpressionPartNotContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitPrecedenceDistinctOperator(ctx *parser.PrecedenceDistinctOperatorContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitPrecedenceEqualOperator(ctx *parser.PrecedenceEqualOperatorContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitPrecedenceEqualExpression(ctx *parser.PrecedenceEqualExpressionContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	fmt.Println("visit PrecedenceEqualExpression")
	for i, _ := range ctx.AllPrecedenceSimilarExpression() {
		G.Visit(ctx.PrecedenceSimilarExpression(i))
	}
	return nil
}

func (G *GHiveVisitor) VisitIsCondition(ctx *parser.IsConditionContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitPrecedenceUnarySuffixExpression(ctx *parser.PrecedenceUnarySuffixExpressionContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	fmt.Println("visit PrecedenceUnarySuffixExpression")
	G.Visit(ctx.PrecedenceEqualExpression())
	return nil
}

func (G *GHiveVisitor) VisitPrecedenceNotOperator(ctx *parser.PrecedenceNotOperatorContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitPrecedenceNotExpression(ctx *parser.PrecedenceNotExpressionContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	fmt.Println("visit PrecedenceNotExpression ")
	G.Visit(ctx.PrecedenceUnarySuffixExpression())
	return nil
}

func (G *GHiveVisitor) VisitPrecedenceAndOperator(ctx *parser.PrecedenceAndOperatorContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitPrecedenceAndExpression(ctx *parser.PrecedenceAndExpressionContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	for i, _ := range ctx.AllPrecedenceNotExpression() {
		G.Visit(ctx.PrecedenceNotExpression(i))
	}
	return nil
}

func (G *GHiveVisitor) VisitPrecedenceOrOperator(ctx *parser.PrecedenceOrOperatorContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitPrecedenceOrExpression(ctx *parser.PrecedenceOrExpressionContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	for i, _ := range ctx.AllPrecedenceAndExpression() {
		G.Visit(ctx.PrecedenceAndExpression(i))
	}
	return nil
}

func (G *GHiveVisitor) VisitBooleanValue(ctx *parser.BooleanValueContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitBooleanValueTok(ctx *parser.BooleanValueTokContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitTableOrPartition(ctx *parser.TableOrPartitionContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	db := ctx.TableName().GetDb().GetText()
	tableName := ctx.TableName().GetTab().GetText()
	G.targetDB = db
	G.targetTable = tableName
	fmt.Printf("db = %s , table = %s \n", db, tableName)
	return nil
}

func (G *GHiveVisitor) VisitPartitionSpec(ctx *parser.PartitionSpecContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitPartitionVal(ctx *parser.PartitionValContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitPartitionSelectorSpec(ctx *parser.PartitionSelectorSpecContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitPartitionSelectorVal(ctx *parser.PartitionSelectorValContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitPartitionSelectorOperator(ctx *parser.PartitionSelectorOperatorContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitSubQuerySelectorOperator(ctx *parser.SubQuerySelectorOperatorContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitSysFuncNames(ctx *parser.SysFuncNamesContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitDescFuncNames(ctx *parser.DescFuncNamesContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitId_(ctx *parser.Id_Context) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	// 返回这个token
	return ctx.Identifier().GetSymbol().GetText()
}

func (G *GHiveVisitor) VisitFunctionIdentifier(ctx *parser.FunctionIdentifierContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitPrincipalIdentifier(ctx *parser.PrincipalIdentifierContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitNonReserved(ctx *parser.NonReservedContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitSql11ReservedKeywordsUsedAsFunctionName(ctx *parser.Sql11ReservedKeywordsUsedAsFunctionNameContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitHint(ctx *parser.HintContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitHintList(ctx *parser.HintListContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitHintItem(ctx *parser.HintItemContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitHintName(ctx *parser.HintNameContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitHintArgs(ctx *parser.HintArgsContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitHintArgName(ctx *parser.HintArgNameContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitPrepareStatement(ctx *parser.PrepareStatementContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitExecuteStatement(ctx *parser.ExecuteStatementContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitExecuteParamList(ctx *parser.ExecuteParamListContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitResourcePlanDdlStatements(ctx *parser.ResourcePlanDdlStatementsContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitRpAssign(ctx *parser.RpAssignContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitRpAssignList(ctx *parser.RpAssignListContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitRpUnassign(ctx *parser.RpUnassignContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitRpUnassignList(ctx *parser.RpUnassignListContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitCreateResourcePlanStatement(ctx *parser.CreateResourcePlanStatementContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitWithReplace(ctx *parser.WithReplaceContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitActivate(ctx *parser.ActivateContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitEnable(ctx *parser.EnableContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitDisable(ctx *parser.DisableContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitUnmanaged(ctx *parser.UnmanagedContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitAlterResourcePlanStatement(ctx *parser.AlterResourcePlanStatementContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitGlobalWmStatement(ctx *parser.GlobalWmStatementContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitReplaceResourcePlanStatement(ctx *parser.ReplaceResourcePlanStatementContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitDropResourcePlanStatement(ctx *parser.DropResourcePlanStatementContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitPoolPath(ctx *parser.PoolPathContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitTriggerExpression(ctx *parser.TriggerExpressionContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitTriggerExpressionStandalone(ctx *parser.TriggerExpressionStandaloneContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitTriggerOrExpression(ctx *parser.TriggerOrExpressionContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitTriggerAndExpression(ctx *parser.TriggerAndExpressionContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitTriggerAtomExpression(ctx *parser.TriggerAtomExpressionContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitTriggerLiteral(ctx *parser.TriggerLiteralContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitComparisionOperator(ctx *parser.ComparisionOperatorContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitTriggerActionExpression(ctx *parser.TriggerActionExpressionContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitTriggerActionExpressionStandalone(ctx *parser.TriggerActionExpressionStandaloneContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitCreateTriggerStatement(ctx *parser.CreateTriggerStatementContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitAlterTriggerStatement(ctx *parser.AlterTriggerStatementContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitDropTriggerStatement(ctx *parser.DropTriggerStatementContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitPoolAssign(ctx *parser.PoolAssignContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitPoolAssignList(ctx *parser.PoolAssignListContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitCreatePoolStatement(ctx *parser.CreatePoolStatementContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitAlterPoolStatement(ctx *parser.AlterPoolStatementContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitDropPoolStatement(ctx *parser.DropPoolStatementContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitCreateMappingStatement(ctx *parser.CreateMappingStatementContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitAlterMappingStatement(ctx *parser.AlterMappingStatementContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}

func (G *GHiveVisitor) VisitDropMappingStatement(ctx *parser.DropMappingStatementContext) interface{} {
	//TODO implement me
	fmt.Println(ctx.GetText())
	return nil
}
