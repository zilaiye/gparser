package iparser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"github.com/golang-collections/collections/stack"
	"gparser/internal/base"
	parser "gparser/internal/iantlr/grammar"
)

type GparserParserListener struct {
	stack *stack.Stack
}

var _ parser.HiveParserListener = &GparserParserListener{}

func NewGparserParserListener() *GparserParserListener {
	return &GparserParserListener{stack: stack.New()}
}

// VisitTerminal is called when a terminal node is visited.
func (s *GparserParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *GparserParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *GparserParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println(ctx.GetText())
}

// ExitEveryRule is called when any rule is exited.
func (s *GparserParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStatement is called when production statement is entered.
func (s *GparserParserListener) EnterStatement(ctx *parser.StatementContext) {
	stmt := &base.Statement{}
	s.stack.Push(stmt)
}

// ExitStatement is called when production statement is exited.
func (s *GparserParserListener) ExitStatement(ctx *parser.StatementContext) {
	stmt := s.stack.Pop()
	fmt.Println(stmt)
}

// EnterExplainStatement is called when production explainStatement is entered.
func (s *GparserParserListener) EnterExplainStatement(ctx *parser.ExplainStatementContext) {}

// ExitExplainStatement is called when production explainStatement is exited.
func (s *GparserParserListener) ExitExplainStatement(ctx *parser.ExplainStatementContext) {}

// EnterExplainOption is called when production explainOption is entered.
func (s *GparserParserListener) EnterExplainOption(ctx *parser.ExplainOptionContext) {}

// ExitExplainOption is called when production explainOption is exited.
func (s *GparserParserListener) ExitExplainOption(ctx *parser.ExplainOptionContext) {}

// EnterVectorizationOnly is called when production vectorizationOnly is entered.
func (s *GparserParserListener) EnterVectorizationOnly(ctx *parser.VectorizationOnlyContext) {}

// ExitVectorizationOnly is called when production vectorizationOnly is exited.
func (s *GparserParserListener) ExitVectorizationOnly(ctx *parser.VectorizationOnlyContext) {}

// EnterVectorizatonDetail is called when production vectorizatonDetail is entered.
func (s *GparserParserListener) EnterVectorizatonDetail(ctx *parser.VectorizatonDetailContext) {}

// ExitVectorizatonDetail is called when production vectorizatonDetail is exited.
func (s *GparserParserListener) ExitVectorizatonDetail(ctx *parser.VectorizatonDetailContext) {}

// EnterExecStatement is called when production execStatement is entered.
func (s *GparserParserListener) EnterExecStatement(ctx *parser.ExecStatementContext) {
	execStmt := &base.ExecStatement{}
	s.stack.Push(execStmt)
}

// ExitExecStatement is called when production execStatement is exited.
func (s *GparserParserListener) ExitExecStatement(ctx *parser.ExecStatementContext) {
	if execStmt, ok := s.stack.Pop().(*base.ExecStatement); ok {
		peek := s.stack.Peek()
		peek.(*base.Statement).ExecStatement = execStmt
	} else {
		fmt.Println("ExitExecStatement encountered an error ! type not match ExecStatement ")
	}
}

// EnterLoadStatement is called when production loadStatement is entered.
func (s *GparserParserListener) EnterLoadStatement(ctx *parser.LoadStatementContext) {}

// ExitLoadStatement is called when production loadStatement is exited.
func (s *GparserParserListener) ExitLoadStatement(ctx *parser.LoadStatementContext) {}

// EnterReplicationClause is called when production replicationClause is entered.
func (s *GparserParserListener) EnterReplicationClause(ctx *parser.ReplicationClauseContext) {}

// ExitReplicationClause is called when production replicationClause is exited.
func (s *GparserParserListener) ExitReplicationClause(ctx *parser.ReplicationClauseContext) {}

// EnterExportStatement is called when production exportStatement is entered.
func (s *GparserParserListener) EnterExportStatement(ctx *parser.ExportStatementContext) {}

// ExitExportStatement is called when production exportStatement is exited.
func (s *GparserParserListener) ExitExportStatement(ctx *parser.ExportStatementContext) {}

// EnterImportStatement is called when production importStatement is entered.
func (s *GparserParserListener) EnterImportStatement(ctx *parser.ImportStatementContext) {}

// ExitImportStatement is called when production importStatement is exited.
func (s *GparserParserListener) ExitImportStatement(ctx *parser.ImportStatementContext) {}

// EnterReplDumpStatement is called when production replDumpStatement is entered.
func (s *GparserParserListener) EnterReplDumpStatement(ctx *parser.ReplDumpStatementContext) {}

// ExitReplDumpStatement is called when production replDumpStatement is exited.
func (s *GparserParserListener) ExitReplDumpStatement(ctx *parser.ReplDumpStatementContext) {}

// EnterReplDbPolicy is called when production replDbPolicy is entered.
func (s *GparserParserListener) EnterReplDbPolicy(ctx *parser.ReplDbPolicyContext) {}

// ExitReplDbPolicy is called when production replDbPolicy is exited.
func (s *GparserParserListener) ExitReplDbPolicy(ctx *parser.ReplDbPolicyContext) {}

// EnterReplLoadStatement is called when production replLoadStatement is entered.
func (s *GparserParserListener) EnterReplLoadStatement(ctx *parser.ReplLoadStatementContext) {}

// ExitReplLoadStatement is called when production replLoadStatement is exited.
func (s *GparserParserListener) ExitReplLoadStatement(ctx *parser.ReplLoadStatementContext) {}

// EnterReplConfigs is called when production replConfigs is entered.
func (s *GparserParserListener) EnterReplConfigs(ctx *parser.ReplConfigsContext) {}

// ExitReplConfigs is called when production replConfigs is exited.
func (s *GparserParserListener) ExitReplConfigs(ctx *parser.ReplConfigsContext) {}

// EnterReplConfigsList is called when production replConfigsList is entered.
func (s *GparserParserListener) EnterReplConfigsList(ctx *parser.ReplConfigsListContext) {}

// ExitReplConfigsList is called when production replConfigsList is exited.
func (s *GparserParserListener) ExitReplConfigsList(ctx *parser.ReplConfigsListContext) {}

// EnterReplTableLevelPolicy is called when production replTableLevelPolicy is entered.
func (s *GparserParserListener) EnterReplTableLevelPolicy(ctx *parser.ReplTableLevelPolicyContext) {}

// ExitReplTableLevelPolicy is called when production replTableLevelPolicy is exited.
func (s *GparserParserListener) ExitReplTableLevelPolicy(ctx *parser.ReplTableLevelPolicyContext) {}

// EnterReplStatusStatement is called when production replStatusStatement is entered.
func (s *GparserParserListener) EnterReplStatusStatement(ctx *parser.ReplStatusStatementContext) {}

// ExitReplStatusStatement is called when production replStatusStatement is exited.
func (s *GparserParserListener) ExitReplStatusStatement(ctx *parser.ReplStatusStatementContext) {}

// EnterDdlStatement is called when production ddlStatement is entered.
func (s *GparserParserListener) EnterDdlStatement(ctx *parser.DdlStatementContext) {}

// ExitDdlStatement is called when production ddlStatement is exited.
func (s *GparserParserListener) ExitDdlStatement(ctx *parser.DdlStatementContext) {}

// EnterIfExists is called when production ifExists is entered.
func (s *GparserParserListener) EnterIfExists(ctx *parser.IfExistsContext) {}

// ExitIfExists is called when production ifExists is exited.
func (s *GparserParserListener) ExitIfExists(ctx *parser.IfExistsContext) {}

// EnterRestrictOrCascade is called when production restrictOrCascade is entered.
func (s *GparserParserListener) EnterRestrictOrCascade(ctx *parser.RestrictOrCascadeContext) {}

// ExitRestrictOrCascade is called when production restrictOrCascade is exited.
func (s *GparserParserListener) ExitRestrictOrCascade(ctx *parser.RestrictOrCascadeContext) {}

// EnterIfNotExists is called when production ifNotExists is entered.
func (s *GparserParserListener) EnterIfNotExists(ctx *parser.IfNotExistsContext) {}

// ExitIfNotExists is called when production ifNotExists is exited.
func (s *GparserParserListener) ExitIfNotExists(ctx *parser.IfNotExistsContext) {}

// EnterForce is called when production force is entered.
func (s *GparserParserListener) EnterForce(ctx *parser.ForceContext) {}

// ExitForce is called when production force is exited.
func (s *GparserParserListener) ExitForce(ctx *parser.ForceContext) {}

// EnterRewriteEnabled is called when production rewriteEnabled is entered.
func (s *GparserParserListener) EnterRewriteEnabled(ctx *parser.RewriteEnabledContext) {}

// ExitRewriteEnabled is called when production rewriteEnabled is exited.
func (s *GparserParserListener) ExitRewriteEnabled(ctx *parser.RewriteEnabledContext) {}

// EnterRewriteDisabled is called when production rewriteDisabled is entered.
func (s *GparserParserListener) EnterRewriteDisabled(ctx *parser.RewriteDisabledContext) {}

// ExitRewriteDisabled is called when production rewriteDisabled is exited.
func (s *GparserParserListener) ExitRewriteDisabled(ctx *parser.RewriteDisabledContext) {}

// EnterStoredAsDirs is called when production storedAsDirs is entered.
func (s *GparserParserListener) EnterStoredAsDirs(ctx *parser.StoredAsDirsContext) {}

// ExitStoredAsDirs is called when production storedAsDirs is exited.
func (s *GparserParserListener) ExitStoredAsDirs(ctx *parser.StoredAsDirsContext) {}

// EnterOrReplace is called when production orReplace is entered.
func (s *GparserParserListener) EnterOrReplace(ctx *parser.OrReplaceContext) {}

// ExitOrReplace is called when production orReplace is exited.
func (s *GparserParserListener) ExitOrReplace(ctx *parser.OrReplaceContext) {}

// EnterCreateDatabaseStatement is called when production createDatabaseStatement is entered.
func (s *GparserParserListener) EnterCreateDatabaseStatement(ctx *parser.CreateDatabaseStatementContext) {
}

// ExitCreateDatabaseStatement is called when production createDatabaseStatement is exited.
func (s *GparserParserListener) ExitCreateDatabaseStatement(ctx *parser.CreateDatabaseStatementContext) {
}

// EnterDbLocation is called when production dbLocation is entered.
func (s *GparserParserListener) EnterDbLocation(ctx *parser.DbLocationContext) {}

// ExitDbLocation is called when production dbLocation is exited.
func (s *GparserParserListener) ExitDbLocation(ctx *parser.DbLocationContext) {}

// EnterDbManagedLocation is called when production dbManagedLocation is entered.
func (s *GparserParserListener) EnterDbManagedLocation(ctx *parser.DbManagedLocationContext) {}

// ExitDbManagedLocation is called when production dbManagedLocation is exited.
func (s *GparserParserListener) ExitDbManagedLocation(ctx *parser.DbManagedLocationContext) {}

// EnterDbProperties is called when production dbProperties is entered.
func (s *GparserParserListener) EnterDbProperties(ctx *parser.DbPropertiesContext) {}

// ExitDbProperties is called when production dbProperties is exited.
func (s *GparserParserListener) ExitDbProperties(ctx *parser.DbPropertiesContext) {}

// EnterDbPropertiesList is called when production dbPropertiesList is entered.
func (s *GparserParserListener) EnterDbPropertiesList(ctx *parser.DbPropertiesListContext) {}

// ExitDbPropertiesList is called when production dbPropertiesList is exited.
func (s *GparserParserListener) ExitDbPropertiesList(ctx *parser.DbPropertiesListContext) {}

// EnterDbConnectorName is called when production dbConnectorName is entered.
func (s *GparserParserListener) EnterDbConnectorName(ctx *parser.DbConnectorNameContext) {}

// ExitDbConnectorName is called when production dbConnectorName is exited.
func (s *GparserParserListener) ExitDbConnectorName(ctx *parser.DbConnectorNameContext) {}

// EnterSwitchDatabaseStatement is called when production switchDatabaseStatement is entered.
func (s *GparserParserListener) EnterSwitchDatabaseStatement(ctx *parser.SwitchDatabaseStatementContext) {
}

// ExitSwitchDatabaseStatement is called when production switchDatabaseStatement is exited.
func (s *GparserParserListener) ExitSwitchDatabaseStatement(ctx *parser.SwitchDatabaseStatementContext) {
}

// EnterDropDatabaseStatement is called when production dropDatabaseStatement is entered.
func (s *GparserParserListener) EnterDropDatabaseStatement(ctx *parser.DropDatabaseStatementContext) {
}

// ExitDropDatabaseStatement is called when production dropDatabaseStatement is exited.
func (s *GparserParserListener) ExitDropDatabaseStatement(ctx *parser.DropDatabaseStatementContext) {}

// EnterDatabaseComment is called when production databaseComment is entered.
func (s *GparserParserListener) EnterDatabaseComment(ctx *parser.DatabaseCommentContext) {}

// ExitDatabaseComment is called when production databaseComment is exited.
func (s *GparserParserListener) ExitDatabaseComment(ctx *parser.DatabaseCommentContext) {}

// EnterTruncateTableStatement is called when production truncateTableStatement is entered.
func (s *GparserParserListener) EnterTruncateTableStatement(ctx *parser.TruncateTableStatementContext) {
}

// ExitTruncateTableStatement is called when production truncateTableStatement is exited.
func (s *GparserParserListener) ExitTruncateTableStatement(ctx *parser.TruncateTableStatementContext) {
}

// EnterDropTableStatement is called when production dropTableStatement is entered.
func (s *GparserParserListener) EnterDropTableStatement(ctx *parser.DropTableStatementContext) {}

// ExitDropTableStatement is called when production dropTableStatement is exited.
func (s *GparserParserListener) ExitDropTableStatement(ctx *parser.DropTableStatementContext) {}

// EnterInputFileFormat is called when production inputFileFormat is entered.
func (s *GparserParserListener) EnterInputFileFormat(ctx *parser.InputFileFormatContext) {}

// ExitInputFileFormat is called when production inputFileFormat is exited.
func (s *GparserParserListener) ExitInputFileFormat(ctx *parser.InputFileFormatContext) {}

// EnterTabTypeExpr is called when production tabTypeExpr is entered.
func (s *GparserParserListener) EnterTabTypeExpr(ctx *parser.TabTypeExprContext) {}

// ExitTabTypeExpr is called when production tabTypeExpr is exited.
func (s *GparserParserListener) ExitTabTypeExpr(ctx *parser.TabTypeExprContext) {}

// EnterPartTypeExpr is called when production partTypeExpr is entered.
func (s *GparserParserListener) EnterPartTypeExpr(ctx *parser.PartTypeExprContext) {}

// ExitPartTypeExpr is called when production partTypeExpr is exited.
func (s *GparserParserListener) ExitPartTypeExpr(ctx *parser.PartTypeExprContext) {}

// EnterTabPartColTypeExpr is called when production tabPartColTypeExpr is entered.
func (s *GparserParserListener) EnterTabPartColTypeExpr(ctx *parser.TabPartColTypeExprContext) {}

// ExitTabPartColTypeExpr is called when production tabPartColTypeExpr is exited.
func (s *GparserParserListener) ExitTabPartColTypeExpr(ctx *parser.TabPartColTypeExprContext) {}

// EnterDescStatement is called when production descStatement is entered.
func (s *GparserParserListener) EnterDescStatement(ctx *parser.DescStatementContext) {}

// ExitDescStatement is called when production descStatement is exited.
func (s *GparserParserListener) ExitDescStatement(ctx *parser.DescStatementContext) {}

// EnterAnalyzeStatement is called when production analyzeStatement is entered.
func (s *GparserParserListener) EnterAnalyzeStatement(ctx *parser.AnalyzeStatementContext) {}

// ExitAnalyzeStatement is called when production analyzeStatement is exited.
func (s *GparserParserListener) ExitAnalyzeStatement(ctx *parser.AnalyzeStatementContext) {}

// EnterFrom_in is called when production from_in is entered.
func (s *GparserParserListener) EnterFrom_in(ctx *parser.From_inContext) {}

// ExitFrom_in is called when production from_in is exited.
func (s *GparserParserListener) ExitFrom_in(ctx *parser.From_inContext) {}

// EnterDb_schema is called when production db_schema is entered.
func (s *GparserParserListener) EnterDb_schema(ctx *parser.Db_schemaContext) {}

// ExitDb_schema is called when production db_schema is exited.
func (s *GparserParserListener) ExitDb_schema(ctx *parser.Db_schemaContext) {}

// EnterShowStatement is called when production showStatement is entered.
func (s *GparserParserListener) EnterShowStatement(ctx *parser.ShowStatementContext) {}

// ExitShowStatement is called when production showStatement is exited.
func (s *GparserParserListener) ExitShowStatement(ctx *parser.ShowStatementContext) {}

// EnterShowTablesFilterExpr is called when production showTablesFilterExpr is entered.
func (s *GparserParserListener) EnterShowTablesFilterExpr(ctx *parser.ShowTablesFilterExprContext) {}

// ExitShowTablesFilterExpr is called when production showTablesFilterExpr is exited.
func (s *GparserParserListener) ExitShowTablesFilterExpr(ctx *parser.ShowTablesFilterExprContext) {}

// EnterLockStatement is called when production lockStatement is entered.
func (s *GparserParserListener) EnterLockStatement(ctx *parser.LockStatementContext) {}

// ExitLockStatement is called when production lockStatement is exited.
func (s *GparserParserListener) ExitLockStatement(ctx *parser.LockStatementContext) {}

// EnterLockDatabase is called when production lockDatabase is entered.
func (s *GparserParserListener) EnterLockDatabase(ctx *parser.LockDatabaseContext) {}

// ExitLockDatabase is called when production lockDatabase is exited.
func (s *GparserParserListener) ExitLockDatabase(ctx *parser.LockDatabaseContext) {}

// EnterLockMode is called when production lockMode is entered.
func (s *GparserParserListener) EnterLockMode(ctx *parser.LockModeContext) {}

// ExitLockMode is called when production lockMode is exited.
func (s *GparserParserListener) ExitLockMode(ctx *parser.LockModeContext) {}

// EnterUnlockStatement is called when production unlockStatement is entered.
func (s *GparserParserListener) EnterUnlockStatement(ctx *parser.UnlockStatementContext) {}

// ExitUnlockStatement is called when production unlockStatement is exited.
func (s *GparserParserListener) ExitUnlockStatement(ctx *parser.UnlockStatementContext) {}

// EnterUnlockDatabase is called when production unlockDatabase is entered.
func (s *GparserParserListener) EnterUnlockDatabase(ctx *parser.UnlockDatabaseContext) {}

// ExitUnlockDatabase is called when production unlockDatabase is exited.
func (s *GparserParserListener) ExitUnlockDatabase(ctx *parser.UnlockDatabaseContext) {}

// EnterCreateRoleStatement is called when production createRoleStatement is entered.
func (s *GparserParserListener) EnterCreateRoleStatement(ctx *parser.CreateRoleStatementContext) {}

// ExitCreateRoleStatement is called when production createRoleStatement is exited.
func (s *GparserParserListener) ExitCreateRoleStatement(ctx *parser.CreateRoleStatementContext) {}

// EnterDropRoleStatement is called when production dropRoleStatement is entered.
func (s *GparserParserListener) EnterDropRoleStatement(ctx *parser.DropRoleStatementContext) {}

// ExitDropRoleStatement is called when production dropRoleStatement is exited.
func (s *GparserParserListener) ExitDropRoleStatement(ctx *parser.DropRoleStatementContext) {}

// EnterGrantPrivileges is called when production grantPrivileges is entered.
func (s *GparserParserListener) EnterGrantPrivileges(ctx *parser.GrantPrivilegesContext) {}

// ExitGrantPrivileges is called when production grantPrivileges is exited.
func (s *GparserParserListener) ExitGrantPrivileges(ctx *parser.GrantPrivilegesContext) {}

// EnterRevokePrivileges is called when production revokePrivileges is entered.
func (s *GparserParserListener) EnterRevokePrivileges(ctx *parser.RevokePrivilegesContext) {}

// ExitRevokePrivileges is called when production revokePrivileges is exited.
func (s *GparserParserListener) ExitRevokePrivileges(ctx *parser.RevokePrivilegesContext) {}

// EnterGrantRole is called when production grantRole is entered.
func (s *GparserParserListener) EnterGrantRole(ctx *parser.GrantRoleContext) {}

// ExitGrantRole is called when production grantRole is exited.
func (s *GparserParserListener) ExitGrantRole(ctx *parser.GrantRoleContext) {}

// EnterRevokeRole is called when production revokeRole is entered.
func (s *GparserParserListener) EnterRevokeRole(ctx *parser.RevokeRoleContext) {}

// ExitRevokeRole is called when production revokeRole is exited.
func (s *GparserParserListener) ExitRevokeRole(ctx *parser.RevokeRoleContext) {}

// EnterShowRoleGrants is called when production showRoleGrants is entered.
func (s *GparserParserListener) EnterShowRoleGrants(ctx *parser.ShowRoleGrantsContext) {}

// ExitShowRoleGrants is called when production showRoleGrants is exited.
func (s *GparserParserListener) ExitShowRoleGrants(ctx *parser.ShowRoleGrantsContext) {}

// EnterShowRoles is called when production showRoles is entered.
func (s *GparserParserListener) EnterShowRoles(ctx *parser.ShowRolesContext) {}

// ExitShowRoles is called when production showRoles is exited.
func (s *GparserParserListener) ExitShowRoles(ctx *parser.ShowRolesContext) {}

// EnterShowCurrentRole is called when production showCurrentRole is entered.
func (s *GparserParserListener) EnterShowCurrentRole(ctx *parser.ShowCurrentRoleContext) {}

// ExitShowCurrentRole is called when production showCurrentRole is exited.
func (s *GparserParserListener) ExitShowCurrentRole(ctx *parser.ShowCurrentRoleContext) {}

// EnterSetRole is called when production setRole is entered.
func (s *GparserParserListener) EnterSetRole(ctx *parser.SetRoleContext) {}

// ExitSetRole is called when production setRole is exited.
func (s *GparserParserListener) ExitSetRole(ctx *parser.SetRoleContext) {}

// EnterShowGrants is called when production showGrants is entered.
func (s *GparserParserListener) EnterShowGrants(ctx *parser.ShowGrantsContext) {}

// ExitShowGrants is called when production showGrants is exited.
func (s *GparserParserListener) ExitShowGrants(ctx *parser.ShowGrantsContext) {}

// EnterShowRolePrincipals is called when production showRolePrincipals is entered.
func (s *GparserParserListener) EnterShowRolePrincipals(ctx *parser.ShowRolePrincipalsContext) {}

// ExitShowRolePrincipals is called when production showRolePrincipals is exited.
func (s *GparserParserListener) ExitShowRolePrincipals(ctx *parser.ShowRolePrincipalsContext) {}

// EnterPrivilegeIncludeColObject is called when production privilegeIncludeColObject is entered.
func (s *GparserParserListener) EnterPrivilegeIncludeColObject(ctx *parser.PrivilegeIncludeColObjectContext) {
}

// ExitPrivilegeIncludeColObject is called when production privilegeIncludeColObject is exited.
func (s *GparserParserListener) ExitPrivilegeIncludeColObject(ctx *parser.PrivilegeIncludeColObjectContext) {
}

// EnterPrivilegeObject is called when production privilegeObject is entered.
func (s *GparserParserListener) EnterPrivilegeObject(ctx *parser.PrivilegeObjectContext) {}

// ExitPrivilegeObject is called when production privilegeObject is exited.
func (s *GparserParserListener) ExitPrivilegeObject(ctx *parser.PrivilegeObjectContext) {}

// EnterPrivObject is called when production privObject is entered.
func (s *GparserParserListener) EnterPrivObject(ctx *parser.PrivObjectContext) {}

// ExitPrivObject is called when production privObject is exited.
func (s *GparserParserListener) ExitPrivObject(ctx *parser.PrivObjectContext) {}

// EnterPrivObjectCols is called when production privObjectCols is entered.
func (s *GparserParserListener) EnterPrivObjectCols(ctx *parser.PrivObjectColsContext) {}

// ExitPrivObjectCols is called when production privObjectCols is exited.
func (s *GparserParserListener) ExitPrivObjectCols(ctx *parser.PrivObjectColsContext) {}

// EnterPrivilegeList is called when production privilegeList is entered.
func (s *GparserParserListener) EnterPrivilegeList(ctx *parser.PrivilegeListContext) {}

// ExitPrivilegeList is called when production privilegeList is exited.
func (s *GparserParserListener) ExitPrivilegeList(ctx *parser.PrivilegeListContext) {}

// EnterPrivlegeDef is called when production privlegeDef is entered.
func (s *GparserParserListener) EnterPrivlegeDef(ctx *parser.PrivlegeDefContext) {}

// ExitPrivlegeDef is called when production privlegeDef is exited.
func (s *GparserParserListener) ExitPrivlegeDef(ctx *parser.PrivlegeDefContext) {}

// EnterPrivilegeType is called when production privilegeType is entered.
func (s *GparserParserListener) EnterPrivilegeType(ctx *parser.PrivilegeTypeContext) {}

// ExitPrivilegeType is called when production privilegeType is exited.
func (s *GparserParserListener) ExitPrivilegeType(ctx *parser.PrivilegeTypeContext) {}

// EnterPrincipalSpecification is called when production principalSpecification is entered.
func (s *GparserParserListener) EnterPrincipalSpecification(ctx *parser.PrincipalSpecificationContext) {
}

// ExitPrincipalSpecification is called when production principalSpecification is exited.
func (s *GparserParserListener) ExitPrincipalSpecification(ctx *parser.PrincipalSpecificationContext) {
}

// EnterPrincipalName is called when production principalName is entered.
func (s *GparserParserListener) EnterPrincipalName(ctx *parser.PrincipalNameContext) {}

// ExitPrincipalName is called when production principalName is exited.
func (s *GparserParserListener) ExitPrincipalName(ctx *parser.PrincipalNameContext) {}

// EnterWithGrantOption is called when production withGrantOption is entered.
func (s *GparserParserListener) EnterWithGrantOption(ctx *parser.WithGrantOptionContext) {}

// ExitWithGrantOption is called when production withGrantOption is exited.
func (s *GparserParserListener) ExitWithGrantOption(ctx *parser.WithGrantOptionContext) {}

// EnterGrantOptionFor is called when production grantOptionFor is entered.
func (s *GparserParserListener) EnterGrantOptionFor(ctx *parser.GrantOptionForContext) {}

// ExitGrantOptionFor is called when production grantOptionFor is exited.
func (s *GparserParserListener) ExitGrantOptionFor(ctx *parser.GrantOptionForContext) {}

// EnterAdminOptionFor is called when production adminOptionFor is entered.
func (s *GparserParserListener) EnterAdminOptionFor(ctx *parser.AdminOptionForContext) {}

// ExitAdminOptionFor is called when production adminOptionFor is exited.
func (s *GparserParserListener) ExitAdminOptionFor(ctx *parser.AdminOptionForContext) {}

// EnterWithAdminOption is called when production withAdminOption is entered.
func (s *GparserParserListener) EnterWithAdminOption(ctx *parser.WithAdminOptionContext) {}

// ExitWithAdminOption is called when production withAdminOption is exited.
func (s *GparserParserListener) ExitWithAdminOption(ctx *parser.WithAdminOptionContext) {}

// EnterMetastoreCheck is called when production metastoreCheck is entered.
func (s *GparserParserListener) EnterMetastoreCheck(ctx *parser.MetastoreCheckContext) {}

// ExitMetastoreCheck is called when production metastoreCheck is exited.
func (s *GparserParserListener) ExitMetastoreCheck(ctx *parser.MetastoreCheckContext) {}

// EnterResourceList is called when production resourceList is entered.
func (s *GparserParserListener) EnterResourceList(ctx *parser.ResourceListContext) {}

// ExitResourceList is called when production resourceList is exited.
func (s *GparserParserListener) ExitResourceList(ctx *parser.ResourceListContext) {}

// EnterResource is called when production resource is entered.
func (s *GparserParserListener) EnterResource(ctx *parser.ResourceContext) {}

// ExitResource is called when production resource is exited.
func (s *GparserParserListener) ExitResource(ctx *parser.ResourceContext) {}

// EnterResourceType is called when production resourceType is entered.
func (s *GparserParserListener) EnterResourceType(ctx *parser.ResourceTypeContext) {}

// ExitResourceType is called when production resourceType is exited.
func (s *GparserParserListener) ExitResourceType(ctx *parser.ResourceTypeContext) {}

// EnterCreateFunctionStatement is called when production createFunctionStatement is entered.
func (s *GparserParserListener) EnterCreateFunctionStatement(ctx *parser.CreateFunctionStatementContext) {
}

// ExitCreateFunctionStatement is called when production createFunctionStatement is exited.
func (s *GparserParserListener) ExitCreateFunctionStatement(ctx *parser.CreateFunctionStatementContext) {
}

// EnterDropFunctionStatement is called when production dropFunctionStatement is entered.
func (s *GparserParserListener) EnterDropFunctionStatement(ctx *parser.DropFunctionStatementContext) {
}

// ExitDropFunctionStatement is called when production dropFunctionStatement is exited.
func (s *GparserParserListener) ExitDropFunctionStatement(ctx *parser.DropFunctionStatementContext) {}

// EnterReloadFunctionsStatement is called when production reloadFunctionsStatement is entered.
func (s *GparserParserListener) EnterReloadFunctionsStatement(ctx *parser.ReloadFunctionsStatementContext) {
}

// ExitReloadFunctionsStatement is called when production reloadFunctionsStatement is exited.
func (s *GparserParserListener) ExitReloadFunctionsStatement(ctx *parser.ReloadFunctionsStatementContext) {
}

// EnterCreateMacroStatement is called when production createMacroStatement is entered.
func (s *GparserParserListener) EnterCreateMacroStatement(ctx *parser.CreateMacroStatementContext) {}

// ExitCreateMacroStatement is called when production createMacroStatement is exited.
func (s *GparserParserListener) ExitCreateMacroStatement(ctx *parser.CreateMacroStatementContext) {}

// EnterDropMacroStatement is called when production dropMacroStatement is entered.
func (s *GparserParserListener) EnterDropMacroStatement(ctx *parser.DropMacroStatementContext) {}

// ExitDropMacroStatement is called when production dropMacroStatement is exited.
func (s *GparserParserListener) ExitDropMacroStatement(ctx *parser.DropMacroStatementContext) {}

// EnterCreateViewStatement is called when production createViewStatement is entered.
func (s *GparserParserListener) EnterCreateViewStatement(ctx *parser.CreateViewStatementContext) {}

// ExitCreateViewStatement is called when production createViewStatement is exited.
func (s *GparserParserListener) ExitCreateViewStatement(ctx *parser.CreateViewStatementContext) {}

// EnterViewPartition is called when production viewPartition is entered.
func (s *GparserParserListener) EnterViewPartition(ctx *parser.ViewPartitionContext) {}

// ExitViewPartition is called when production viewPartition is exited.
func (s *GparserParserListener) ExitViewPartition(ctx *parser.ViewPartitionContext) {}

// EnterViewOrganization is called when production viewOrganization is entered.
func (s *GparserParserListener) EnterViewOrganization(ctx *parser.ViewOrganizationContext) {}

// ExitViewOrganization is called when production viewOrganization is exited.
func (s *GparserParserListener) ExitViewOrganization(ctx *parser.ViewOrganizationContext) {}

// EnterViewClusterSpec is called when production viewClusterSpec is entered.
func (s *GparserParserListener) EnterViewClusterSpec(ctx *parser.ViewClusterSpecContext) {}

// ExitViewClusterSpec is called when production viewClusterSpec is exited.
func (s *GparserParserListener) ExitViewClusterSpec(ctx *parser.ViewClusterSpecContext) {}

// EnterViewComplexSpec is called when production viewComplexSpec is entered.
func (s *GparserParserListener) EnterViewComplexSpec(ctx *parser.ViewComplexSpecContext) {}

// ExitViewComplexSpec is called when production viewComplexSpec is exited.
func (s *GparserParserListener) ExitViewComplexSpec(ctx *parser.ViewComplexSpecContext) {}

// EnterViewDistSpec is called when production viewDistSpec is entered.
func (s *GparserParserListener) EnterViewDistSpec(ctx *parser.ViewDistSpecContext) {}

// ExitViewDistSpec is called when production viewDistSpec is exited.
func (s *GparserParserListener) ExitViewDistSpec(ctx *parser.ViewDistSpecContext) {}

// EnterViewSortSpec is called when production viewSortSpec is entered.
func (s *GparserParserListener) EnterViewSortSpec(ctx *parser.ViewSortSpecContext) {}

// ExitViewSortSpec is called when production viewSortSpec is exited.
func (s *GparserParserListener) ExitViewSortSpec(ctx *parser.ViewSortSpecContext) {}

// EnterDropViewStatement is called when production dropViewStatement is entered.
func (s *GparserParserListener) EnterDropViewStatement(ctx *parser.DropViewStatementContext) {}

// ExitDropViewStatement is called when production dropViewStatement is exited.
func (s *GparserParserListener) ExitDropViewStatement(ctx *parser.DropViewStatementContext) {}

// EnterCreateMaterializedViewStatement is called when production createMaterializedViewStatement is entered.
func (s *GparserParserListener) EnterCreateMaterializedViewStatement(ctx *parser.CreateMaterializedViewStatementContext) {
}

// ExitCreateMaterializedViewStatement is called when production createMaterializedViewStatement is exited.
func (s *GparserParserListener) ExitCreateMaterializedViewStatement(ctx *parser.CreateMaterializedViewStatementContext) {
}

// EnterDropMaterializedViewStatement is called when production dropMaterializedViewStatement is entered.
func (s *GparserParserListener) EnterDropMaterializedViewStatement(ctx *parser.DropMaterializedViewStatementContext) {
}

// ExitDropMaterializedViewStatement is called when production dropMaterializedViewStatement is exited.
func (s *GparserParserListener) ExitDropMaterializedViewStatement(ctx *parser.DropMaterializedViewStatementContext) {
}

// EnterCreateScheduledQueryStatement is called when production createScheduledQueryStatement is entered.
func (s *GparserParserListener) EnterCreateScheduledQueryStatement(ctx *parser.CreateScheduledQueryStatementContext) {
}

// ExitCreateScheduledQueryStatement is called when production createScheduledQueryStatement is exited.
func (s *GparserParserListener) ExitCreateScheduledQueryStatement(ctx *parser.CreateScheduledQueryStatementContext) {
}

// EnterDropScheduledQueryStatement is called when production dropScheduledQueryStatement is entered.
func (s *GparserParserListener) EnterDropScheduledQueryStatement(ctx *parser.DropScheduledQueryStatementContext) {
}

// ExitDropScheduledQueryStatement is called when production dropScheduledQueryStatement is exited.
func (s *GparserParserListener) ExitDropScheduledQueryStatement(ctx *parser.DropScheduledQueryStatementContext) {
}

// EnterAlterScheduledQueryStatement is called when production alterScheduledQueryStatement is entered.
func (s *GparserParserListener) EnterAlterScheduledQueryStatement(ctx *parser.AlterScheduledQueryStatementContext) {
}

// ExitAlterScheduledQueryStatement is called when production alterScheduledQueryStatement is exited.
func (s *GparserParserListener) ExitAlterScheduledQueryStatement(ctx *parser.AlterScheduledQueryStatementContext) {
}

// EnterAlterScheduledQueryChange is called when production alterScheduledQueryChange is entered.
func (s *GparserParserListener) EnterAlterScheduledQueryChange(ctx *parser.AlterScheduledQueryChangeContext) {
}

// ExitAlterScheduledQueryChange is called when production alterScheduledQueryChange is exited.
func (s *GparserParserListener) ExitAlterScheduledQueryChange(ctx *parser.AlterScheduledQueryChangeContext) {
}

// EnterScheduleSpec is called when production scheduleSpec is entered.
func (s *GparserParserListener) EnterScheduleSpec(ctx *parser.ScheduleSpecContext) {}

// ExitScheduleSpec is called when production scheduleSpec is exited.
func (s *GparserParserListener) ExitScheduleSpec(ctx *parser.ScheduleSpecContext) {}

// EnterExecutedAsSpec is called when production executedAsSpec is entered.
func (s *GparserParserListener) EnterExecutedAsSpec(ctx *parser.ExecutedAsSpecContext) {}

// ExitExecutedAsSpec is called when production executedAsSpec is exited.
func (s *GparserParserListener) ExitExecutedAsSpec(ctx *parser.ExecutedAsSpecContext) {}

// EnterDefinedAsSpec is called when production definedAsSpec is entered.
func (s *GparserParserListener) EnterDefinedAsSpec(ctx *parser.DefinedAsSpecContext) {}

// ExitDefinedAsSpec is called when production definedAsSpec is exited.
func (s *GparserParserListener) ExitDefinedAsSpec(ctx *parser.DefinedAsSpecContext) {}

// EnterShowFunctionIdentifier is called when production showFunctionIdentifier is entered.
func (s *GparserParserListener) EnterShowFunctionIdentifier(ctx *parser.ShowFunctionIdentifierContext) {
}

// ExitShowFunctionIdentifier is called when production showFunctionIdentifier is exited.
func (s *GparserParserListener) ExitShowFunctionIdentifier(ctx *parser.ShowFunctionIdentifierContext) {
}

// EnterShowStmtIdentifier is called when production showStmtIdentifier is entered.
func (s *GparserParserListener) EnterShowStmtIdentifier(ctx *parser.ShowStmtIdentifierContext) {}

// ExitShowStmtIdentifier is called when production showStmtIdentifier is exited.
func (s *GparserParserListener) ExitShowStmtIdentifier(ctx *parser.ShowStmtIdentifierContext) {}

// EnterTableComment is called when production tableComment is entered.
func (s *GparserParserListener) EnterTableComment(ctx *parser.TableCommentContext) {}

// ExitTableComment is called when production tableComment is exited.
func (s *GparserParserListener) ExitTableComment(ctx *parser.TableCommentContext) {}

// EnterCreateTablePartitionSpec is called when production createTablePartitionSpec is entered.
func (s *GparserParserListener) EnterCreateTablePartitionSpec(ctx *parser.CreateTablePartitionSpecContext) {
}

// ExitCreateTablePartitionSpec is called when production createTablePartitionSpec is exited.
func (s *GparserParserListener) ExitCreateTablePartitionSpec(ctx *parser.CreateTablePartitionSpecContext) {
}

// EnterCreateTablePartitionColumnTypeSpec is called when production createTablePartitionColumnTypeSpec is entered.
func (s *GparserParserListener) EnterCreateTablePartitionColumnTypeSpec(ctx *parser.CreateTablePartitionColumnTypeSpecContext) {
}

// ExitCreateTablePartitionColumnTypeSpec is called when production createTablePartitionColumnTypeSpec is exited.
func (s *GparserParserListener) ExitCreateTablePartitionColumnTypeSpec(ctx *parser.CreateTablePartitionColumnTypeSpecContext) {
}

// EnterCreateTablePartitionColumnSpec is called when production createTablePartitionColumnSpec is entered.
func (s *GparserParserListener) EnterCreateTablePartitionColumnSpec(ctx *parser.CreateTablePartitionColumnSpecContext) {
}

// ExitCreateTablePartitionColumnSpec is called when production createTablePartitionColumnSpec is exited.
func (s *GparserParserListener) ExitCreateTablePartitionColumnSpec(ctx *parser.CreateTablePartitionColumnSpecContext) {
}

// EnterPartitionTransformSpec is called when production partitionTransformSpec is entered.
func (s *GparserParserListener) EnterPartitionTransformSpec(ctx *parser.PartitionTransformSpecContext) {
}

// ExitPartitionTransformSpec is called when production partitionTransformSpec is exited.
func (s *GparserParserListener) ExitPartitionTransformSpec(ctx *parser.PartitionTransformSpecContext) {
}

// EnterColumnNameTransformConstraint is called when production columnNameTransformConstraint is entered.
func (s *GparserParserListener) EnterColumnNameTransformConstraint(ctx *parser.ColumnNameTransformConstraintContext) {
}

// ExitColumnNameTransformConstraint is called when production columnNameTransformConstraint is exited.
func (s *GparserParserListener) ExitColumnNameTransformConstraint(ctx *parser.ColumnNameTransformConstraintContext) {
}

// EnterPartitionTransformType is called when production partitionTransformType is entered.
func (s *GparserParserListener) EnterPartitionTransformType(ctx *parser.PartitionTransformTypeContext) {
}

// ExitPartitionTransformType is called when production partitionTransformType is exited.
func (s *GparserParserListener) ExitPartitionTransformType(ctx *parser.PartitionTransformTypeContext) {
}

// EnterTableBuckets is called when production tableBuckets is entered.
func (s *GparserParserListener) EnterTableBuckets(ctx *parser.TableBucketsContext) {}

// ExitTableBuckets is called when production tableBuckets is exited.
func (s *GparserParserListener) ExitTableBuckets(ctx *parser.TableBucketsContext) {}

// EnterTableImplBuckets is called when production tableImplBuckets is entered.
func (s *GparserParserListener) EnterTableImplBuckets(ctx *parser.TableImplBucketsContext) {}

// ExitTableImplBuckets is called when production tableImplBuckets is exited.
func (s *GparserParserListener) ExitTableImplBuckets(ctx *parser.TableImplBucketsContext) {}

// EnterTableSkewed is called when production tableSkewed is entered.
func (s *GparserParserListener) EnterTableSkewed(ctx *parser.TableSkewedContext) {}

// ExitTableSkewed is called when production tableSkewed is exited.
func (s *GparserParserListener) ExitTableSkewed(ctx *parser.TableSkewedContext) {}

// EnterRowFormat is called when production rowFormat is entered.
func (s *GparserParserListener) EnterRowFormat(ctx *parser.RowFormatContext) {}

// ExitRowFormat is called when production rowFormat is exited.
func (s *GparserParserListener) ExitRowFormat(ctx *parser.RowFormatContext) {}

// EnterRecordReader is called when production recordReader is entered.
func (s *GparserParserListener) EnterRecordReader(ctx *parser.RecordReaderContext) {}

// ExitRecordReader is called when production recordReader is exited.
func (s *GparserParserListener) ExitRecordReader(ctx *parser.RecordReaderContext) {}

// EnterRecordWriter is called when production recordWriter is entered.
func (s *GparserParserListener) EnterRecordWriter(ctx *parser.RecordWriterContext) {}

// ExitRecordWriter is called when production recordWriter is exited.
func (s *GparserParserListener) ExitRecordWriter(ctx *parser.RecordWriterContext) {}

// EnterRowFormatSerde is called when production rowFormatSerde is entered.
func (s *GparserParserListener) EnterRowFormatSerde(ctx *parser.RowFormatSerdeContext) {}

// ExitRowFormatSerde is called when production rowFormatSerde is exited.
func (s *GparserParserListener) ExitRowFormatSerde(ctx *parser.RowFormatSerdeContext) {}

// EnterRowFormatDelimited is called when production rowFormatDelimited is entered.
func (s *GparserParserListener) EnterRowFormatDelimited(ctx *parser.RowFormatDelimitedContext) {}

// ExitRowFormatDelimited is called when production rowFormatDelimited is exited.
func (s *GparserParserListener) ExitRowFormatDelimited(ctx *parser.RowFormatDelimitedContext) {}

// EnterTableRowFormat is called when production tableRowFormat is entered.
func (s *GparserParserListener) EnterTableRowFormat(ctx *parser.TableRowFormatContext) {}

// ExitTableRowFormat is called when production tableRowFormat is exited.
func (s *GparserParserListener) ExitTableRowFormat(ctx *parser.TableRowFormatContext) {}

// EnterTablePropertiesPrefixed is called when production tablePropertiesPrefixed is entered.
func (s *GparserParserListener) EnterTablePropertiesPrefixed(ctx *parser.TablePropertiesPrefixedContext) {
}

// ExitTablePropertiesPrefixed is called when production tablePropertiesPrefixed is exited.
func (s *GparserParserListener) ExitTablePropertiesPrefixed(ctx *parser.TablePropertiesPrefixedContext) {
}

// EnterTableProperties is called when production tableProperties is entered.
func (s *GparserParserListener) EnterTableProperties(ctx *parser.TablePropertiesContext) {}

// ExitTableProperties is called when production tableProperties is exited.
func (s *GparserParserListener) ExitTableProperties(ctx *parser.TablePropertiesContext) {}

// EnterTablePropertiesList is called when production tablePropertiesList is entered.
func (s *GparserParserListener) EnterTablePropertiesList(ctx *parser.TablePropertiesListContext) {}

// ExitTablePropertiesList is called when production tablePropertiesList is exited.
func (s *GparserParserListener) ExitTablePropertiesList(ctx *parser.TablePropertiesListContext) {}

// EnterKeyValueProperty is called when production keyValueProperty is entered.
func (s *GparserParserListener) EnterKeyValueProperty(ctx *parser.KeyValuePropertyContext) {}

// ExitKeyValueProperty is called when production keyValueProperty is exited.
func (s *GparserParserListener) ExitKeyValueProperty(ctx *parser.KeyValuePropertyContext) {}

// EnterKeyProperty is called when production keyProperty is entered.
func (s *GparserParserListener) EnterKeyProperty(ctx *parser.KeyPropertyContext) {}

// ExitKeyProperty is called when production keyProperty is exited.
func (s *GparserParserListener) ExitKeyProperty(ctx *parser.KeyPropertyContext) {}

// EnterTableRowFormatFieldIdentifier is called when production tableRowFormatFieldIdentifier is entered.
func (s *GparserParserListener) EnterTableRowFormatFieldIdentifier(ctx *parser.TableRowFormatFieldIdentifierContext) {
}

// ExitTableRowFormatFieldIdentifier is called when production tableRowFormatFieldIdentifier is exited.
func (s *GparserParserListener) ExitTableRowFormatFieldIdentifier(ctx *parser.TableRowFormatFieldIdentifierContext) {
}

// EnterTableRowFormatCollItemsIdentifier is called when production tableRowFormatCollItemsIdentifier is entered.
func (s *GparserParserListener) EnterTableRowFormatCollItemsIdentifier(ctx *parser.TableRowFormatCollItemsIdentifierContext) {
}

// ExitTableRowFormatCollItemsIdentifier is called when production tableRowFormatCollItemsIdentifier is exited.
func (s *GparserParserListener) ExitTableRowFormatCollItemsIdentifier(ctx *parser.TableRowFormatCollItemsIdentifierContext) {
}

// EnterTableRowFormatMapKeysIdentifier is called when production tableRowFormatMapKeysIdentifier is entered.
func (s *GparserParserListener) EnterTableRowFormatMapKeysIdentifier(ctx *parser.TableRowFormatMapKeysIdentifierContext) {
}

// ExitTableRowFormatMapKeysIdentifier is called when production tableRowFormatMapKeysIdentifier is exited.
func (s *GparserParserListener) ExitTableRowFormatMapKeysIdentifier(ctx *parser.TableRowFormatMapKeysIdentifierContext) {
}

// EnterTableRowFormatLinesIdentifier is called when production tableRowFormatLinesIdentifier is entered.
func (s *GparserParserListener) EnterTableRowFormatLinesIdentifier(ctx *parser.TableRowFormatLinesIdentifierContext) {
}

// ExitTableRowFormatLinesIdentifier is called when production tableRowFormatLinesIdentifier is exited.
func (s *GparserParserListener) ExitTableRowFormatLinesIdentifier(ctx *parser.TableRowFormatLinesIdentifierContext) {
}

// EnterTableRowNullFormat is called when production tableRowNullFormat is entered.
func (s *GparserParserListener) EnterTableRowNullFormat(ctx *parser.TableRowNullFormatContext) {}

// ExitTableRowNullFormat is called when production tableRowNullFormat is exited.
func (s *GparserParserListener) ExitTableRowNullFormat(ctx *parser.TableRowNullFormatContext) {}

// EnterTableFileFormat is called when production tableFileFormat is entered.
func (s *GparserParserListener) EnterTableFileFormat(ctx *parser.TableFileFormatContext) {}

// ExitTableFileFormat is called when production tableFileFormat is exited.
func (s *GparserParserListener) ExitTableFileFormat(ctx *parser.TableFileFormatContext) {}

// EnterTableLocation is called when production tableLocation is entered.
func (s *GparserParserListener) EnterTableLocation(ctx *parser.TableLocationContext) {}

// ExitTableLocation is called when production tableLocation is exited.
func (s *GparserParserListener) ExitTableLocation(ctx *parser.TableLocationContext) {}

// EnterColumnNameTypeList is called when production columnNameTypeList is entered.
func (s *GparserParserListener) EnterColumnNameTypeList(ctx *parser.ColumnNameTypeListContext) {}

// ExitColumnNameTypeList is called when production columnNameTypeList is exited.
func (s *GparserParserListener) ExitColumnNameTypeList(ctx *parser.ColumnNameTypeListContext) {}

// EnterColumnNameTypeOrConstraintList is called when production columnNameTypeOrConstraintList is entered.
func (s *GparserParserListener) EnterColumnNameTypeOrConstraintList(ctx *parser.ColumnNameTypeOrConstraintListContext) {
}

// ExitColumnNameTypeOrConstraintList is called when production columnNameTypeOrConstraintList is exited.
func (s *GparserParserListener) ExitColumnNameTypeOrConstraintList(ctx *parser.ColumnNameTypeOrConstraintListContext) {
}

// EnterColumnNameColonTypeList is called when production columnNameColonTypeList is entered.
func (s *GparserParserListener) EnterColumnNameColonTypeList(ctx *parser.ColumnNameColonTypeListContext) {
}

// ExitColumnNameColonTypeList is called when production columnNameColonTypeList is exited.
func (s *GparserParserListener) ExitColumnNameColonTypeList(ctx *parser.ColumnNameColonTypeListContext) {
}

// EnterColumnNameList is called when production columnNameList is entered.
func (s *GparserParserListener) EnterColumnNameList(ctx *parser.ColumnNameListContext) {}

// ExitColumnNameList is called when production columnNameList is exited.
func (s *GparserParserListener) ExitColumnNameList(ctx *parser.ColumnNameListContext) {}

// EnterColumnName is called when production columnName is entered.
func (s *GparserParserListener) EnterColumnName(ctx *parser.ColumnNameContext) {}

// ExitColumnName is called when production columnName is exited.
func (s *GparserParserListener) ExitColumnName(ctx *parser.ColumnNameContext) {}

// EnterExtColumnName is called when production extColumnName is entered.
func (s *GparserParserListener) EnterExtColumnName(ctx *parser.ExtColumnNameContext) {}

// ExitExtColumnName is called when production extColumnName is exited.
func (s *GparserParserListener) ExitExtColumnName(ctx *parser.ExtColumnNameContext) {}

// EnterColumnNameOrderList is called when production columnNameOrderList is entered.
func (s *GparserParserListener) EnterColumnNameOrderList(ctx *parser.ColumnNameOrderListContext) {}

// ExitColumnNameOrderList is called when production columnNameOrderList is exited.
func (s *GparserParserListener) ExitColumnNameOrderList(ctx *parser.ColumnNameOrderListContext) {}

// EnterColumnParenthesesList is called when production columnParenthesesList is entered.
func (s *GparserParserListener) EnterColumnParenthesesList(ctx *parser.ColumnParenthesesListContext) {
}

// ExitColumnParenthesesList is called when production columnParenthesesList is exited.
func (s *GparserParserListener) ExitColumnParenthesesList(ctx *parser.ColumnParenthesesListContext) {}

// EnterEnableValidateSpecification is called when production enableValidateSpecification is entered.
func (s *GparserParserListener) EnterEnableValidateSpecification(ctx *parser.EnableValidateSpecificationContext) {
}

// ExitEnableValidateSpecification is called when production enableValidateSpecification is exited.
func (s *GparserParserListener) ExitEnableValidateSpecification(ctx *parser.EnableValidateSpecificationContext) {
}

// EnterEnableSpecification is called when production enableSpecification is entered.
func (s *GparserParserListener) EnterEnableSpecification(ctx *parser.EnableSpecificationContext) {}

// ExitEnableSpecification is called when production enableSpecification is exited.
func (s *GparserParserListener) ExitEnableSpecification(ctx *parser.EnableSpecificationContext) {}

// EnterValidateSpecification is called when production validateSpecification is entered.
func (s *GparserParserListener) EnterValidateSpecification(ctx *parser.ValidateSpecificationContext) {
}

// ExitValidateSpecification is called when production validateSpecification is exited.
func (s *GparserParserListener) ExitValidateSpecification(ctx *parser.ValidateSpecificationContext) {}

// EnterEnforcedSpecification is called when production enforcedSpecification is entered.
func (s *GparserParserListener) EnterEnforcedSpecification(ctx *parser.EnforcedSpecificationContext) {
}

// ExitEnforcedSpecification is called when production enforcedSpecification is exited.
func (s *GparserParserListener) ExitEnforcedSpecification(ctx *parser.EnforcedSpecificationContext) {}

// EnterRelySpecification is called when production relySpecification is entered.
func (s *GparserParserListener) EnterRelySpecification(ctx *parser.RelySpecificationContext) {}

// ExitRelySpecification is called when production relySpecification is exited.
func (s *GparserParserListener) ExitRelySpecification(ctx *parser.RelySpecificationContext) {}

// EnterCreateConstraint is called when production createConstraint is entered.
func (s *GparserParserListener) EnterCreateConstraint(ctx *parser.CreateConstraintContext) {}

// ExitCreateConstraint is called when production createConstraint is exited.
func (s *GparserParserListener) ExitCreateConstraint(ctx *parser.CreateConstraintContext) {}

// EnterAlterConstraintWithName is called when production alterConstraintWithName is entered.
func (s *GparserParserListener) EnterAlterConstraintWithName(ctx *parser.AlterConstraintWithNameContext) {
}

// ExitAlterConstraintWithName is called when production alterConstraintWithName is exited.
func (s *GparserParserListener) ExitAlterConstraintWithName(ctx *parser.AlterConstraintWithNameContext) {
}

// EnterTableLevelConstraint is called when production tableLevelConstraint is entered.
func (s *GparserParserListener) EnterTableLevelConstraint(ctx *parser.TableLevelConstraintContext) {}

// ExitTableLevelConstraint is called when production tableLevelConstraint is exited.
func (s *GparserParserListener) ExitTableLevelConstraint(ctx *parser.TableLevelConstraintContext) {}

// EnterPkUkConstraint is called when production pkUkConstraint is entered.
func (s *GparserParserListener) EnterPkUkConstraint(ctx *parser.PkUkConstraintContext) {}

// ExitPkUkConstraint is called when production pkUkConstraint is exited.
func (s *GparserParserListener) ExitPkUkConstraint(ctx *parser.PkUkConstraintContext) {}

// EnterCheckConstraint is called when production checkConstraint is entered.
func (s *GparserParserListener) EnterCheckConstraint(ctx *parser.CheckConstraintContext) {}

// ExitCheckConstraint is called when production checkConstraint is exited.
func (s *GparserParserListener) ExitCheckConstraint(ctx *parser.CheckConstraintContext) {}

// EnterCreateForeignKey is called when production createForeignKey is entered.
func (s *GparserParserListener) EnterCreateForeignKey(ctx *parser.CreateForeignKeyContext) {}

// ExitCreateForeignKey is called when production createForeignKey is exited.
func (s *GparserParserListener) ExitCreateForeignKey(ctx *parser.CreateForeignKeyContext) {}

// EnterAlterForeignKeyWithName is called when production alterForeignKeyWithName is entered.
func (s *GparserParserListener) EnterAlterForeignKeyWithName(ctx *parser.AlterForeignKeyWithNameContext) {
}

// ExitAlterForeignKeyWithName is called when production alterForeignKeyWithName is exited.
func (s *GparserParserListener) ExitAlterForeignKeyWithName(ctx *parser.AlterForeignKeyWithNameContext) {
}

// EnterSkewedValueElement is called when production skewedValueElement is entered.
func (s *GparserParserListener) EnterSkewedValueElement(ctx *parser.SkewedValueElementContext) {}

// ExitSkewedValueElement is called when production skewedValueElement is exited.
func (s *GparserParserListener) ExitSkewedValueElement(ctx *parser.SkewedValueElementContext) {}

// EnterSkewedColumnValuePairList is called when production skewedColumnValuePairList is entered.
func (s *GparserParserListener) EnterSkewedColumnValuePairList(ctx *parser.SkewedColumnValuePairListContext) {
}

// ExitSkewedColumnValuePairList is called when production skewedColumnValuePairList is exited.
func (s *GparserParserListener) ExitSkewedColumnValuePairList(ctx *parser.SkewedColumnValuePairListContext) {
}

// EnterSkewedColumnValuePair is called when production skewedColumnValuePair is entered.
func (s *GparserParserListener) EnterSkewedColumnValuePair(ctx *parser.SkewedColumnValuePairContext) {
}

// ExitSkewedColumnValuePair is called when production skewedColumnValuePair is exited.
func (s *GparserParserListener) ExitSkewedColumnValuePair(ctx *parser.SkewedColumnValuePairContext) {}

// EnterSkewedColumnValues is called when production skewedColumnValues is entered.
func (s *GparserParserListener) EnterSkewedColumnValues(ctx *parser.SkewedColumnValuesContext) {}

// ExitSkewedColumnValues is called when production skewedColumnValues is exited.
func (s *GparserParserListener) ExitSkewedColumnValues(ctx *parser.SkewedColumnValuesContext) {}

// EnterSkewedColumnValue is called when production skewedColumnValue is entered.
func (s *GparserParserListener) EnterSkewedColumnValue(ctx *parser.SkewedColumnValueContext) {}

// ExitSkewedColumnValue is called when production skewedColumnValue is exited.
func (s *GparserParserListener) ExitSkewedColumnValue(ctx *parser.SkewedColumnValueContext) {}

// EnterSkewedValueLocationElement is called when production skewedValueLocationElement is entered.
func (s *GparserParserListener) EnterSkewedValueLocationElement(ctx *parser.SkewedValueLocationElementContext) {
}

// ExitSkewedValueLocationElement is called when production skewedValueLocationElement is exited.
func (s *GparserParserListener) ExitSkewedValueLocationElement(ctx *parser.SkewedValueLocationElementContext) {
}

// EnterOrderSpecification is called when production orderSpecification is entered.
func (s *GparserParserListener) EnterOrderSpecification(ctx *parser.OrderSpecificationContext) {}

// ExitOrderSpecification is called when production orderSpecification is exited.
func (s *GparserParserListener) ExitOrderSpecification(ctx *parser.OrderSpecificationContext) {}

// EnterNullOrdering is called when production nullOrdering is entered.
func (s *GparserParserListener) EnterNullOrdering(ctx *parser.NullOrderingContext) {}

// ExitNullOrdering is called when production nullOrdering is exited.
func (s *GparserParserListener) ExitNullOrdering(ctx *parser.NullOrderingContext) {}

// EnterColumnNameOrder is called when production columnNameOrder is entered.
func (s *GparserParserListener) EnterColumnNameOrder(ctx *parser.ColumnNameOrderContext) {}

// ExitColumnNameOrder is called when production columnNameOrder is exited.
func (s *GparserParserListener) ExitColumnNameOrder(ctx *parser.ColumnNameOrderContext) {}

// EnterColumnNameCommentList is called when production columnNameCommentList is entered.
func (s *GparserParserListener) EnterColumnNameCommentList(ctx *parser.ColumnNameCommentListContext) {
}

// ExitColumnNameCommentList is called when production columnNameCommentList is exited.
func (s *GparserParserListener) ExitColumnNameCommentList(ctx *parser.ColumnNameCommentListContext) {}

// EnterColumnNameComment is called when production columnNameComment is entered.
func (s *GparserParserListener) EnterColumnNameComment(ctx *parser.ColumnNameCommentContext) {}

// ExitColumnNameComment is called when production columnNameComment is exited.
func (s *GparserParserListener) ExitColumnNameComment(ctx *parser.ColumnNameCommentContext) {}

// EnterOrderSpecificationRewrite is called when production orderSpecificationRewrite is entered.
func (s *GparserParserListener) EnterOrderSpecificationRewrite(ctx *parser.OrderSpecificationRewriteContext) {
}

// ExitOrderSpecificationRewrite is called when production orderSpecificationRewrite is exited.
func (s *GparserParserListener) ExitOrderSpecificationRewrite(ctx *parser.OrderSpecificationRewriteContext) {
}

// EnterColumnRefOrder is called when production columnRefOrder is entered.
func (s *GparserParserListener) EnterColumnRefOrder(ctx *parser.ColumnRefOrderContext) {}

// ExitColumnRefOrder is called when production columnRefOrder is exited.
func (s *GparserParserListener) ExitColumnRefOrder(ctx *parser.ColumnRefOrderContext) {}

// EnterColumnNameType is called when production columnNameType is entered.
func (s *GparserParserListener) EnterColumnNameType(ctx *parser.ColumnNameTypeContext) {}

// ExitColumnNameType is called when production columnNameType is exited.
func (s *GparserParserListener) ExitColumnNameType(ctx *parser.ColumnNameTypeContext) {}

// EnterColumnNameTypeOrConstraint is called when production columnNameTypeOrConstraint is entered.
func (s *GparserParserListener) EnterColumnNameTypeOrConstraint(ctx *parser.ColumnNameTypeOrConstraintContext) {
}

// ExitColumnNameTypeOrConstraint is called when production columnNameTypeOrConstraint is exited.
func (s *GparserParserListener) ExitColumnNameTypeOrConstraint(ctx *parser.ColumnNameTypeOrConstraintContext) {
}

// EnterTableConstraint is called when production tableConstraint is entered.
func (s *GparserParserListener) EnterTableConstraint(ctx *parser.TableConstraintContext) {}

// ExitTableConstraint is called when production tableConstraint is exited.
func (s *GparserParserListener) ExitTableConstraint(ctx *parser.TableConstraintContext) {}

// EnterColumnNameTypeConstraint is called when production columnNameTypeConstraint is entered.
func (s *GparserParserListener) EnterColumnNameTypeConstraint(ctx *parser.ColumnNameTypeConstraintContext) {
}

// ExitColumnNameTypeConstraint is called when production columnNameTypeConstraint is exited.
func (s *GparserParserListener) ExitColumnNameTypeConstraint(ctx *parser.ColumnNameTypeConstraintContext) {
}

// EnterColumnConstraint is called when production columnConstraint is entered.
func (s *GparserParserListener) EnterColumnConstraint(ctx *parser.ColumnConstraintContext) {}

// ExitColumnConstraint is called when production columnConstraint is exited.
func (s *GparserParserListener) ExitColumnConstraint(ctx *parser.ColumnConstraintContext) {}

// EnterForeignKeyConstraint is called when production foreignKeyConstraint is entered.
func (s *GparserParserListener) EnterForeignKeyConstraint(ctx *parser.ForeignKeyConstraintContext) {}

// ExitForeignKeyConstraint is called when production foreignKeyConstraint is exited.
func (s *GparserParserListener) ExitForeignKeyConstraint(ctx *parser.ForeignKeyConstraintContext) {}

// EnterColConstraint is called when production colConstraint is entered.
func (s *GparserParserListener) EnterColConstraint(ctx *parser.ColConstraintContext) {}

// ExitColConstraint is called when production colConstraint is exited.
func (s *GparserParserListener) ExitColConstraint(ctx *parser.ColConstraintContext) {}

// EnterAlterColumnConstraint is called when production alterColumnConstraint is entered.
func (s *GparserParserListener) EnterAlterColumnConstraint(ctx *parser.AlterColumnConstraintContext) {
}

// ExitAlterColumnConstraint is called when production alterColumnConstraint is exited.
func (s *GparserParserListener) ExitAlterColumnConstraint(ctx *parser.AlterColumnConstraintContext) {}

// EnterAlterForeignKeyConstraint is called when production alterForeignKeyConstraint is entered.
func (s *GparserParserListener) EnterAlterForeignKeyConstraint(ctx *parser.AlterForeignKeyConstraintContext) {
}

// ExitAlterForeignKeyConstraint is called when production alterForeignKeyConstraint is exited.
func (s *GparserParserListener) ExitAlterForeignKeyConstraint(ctx *parser.AlterForeignKeyConstraintContext) {
}

// EnterAlterColConstraint is called when production alterColConstraint is entered.
func (s *GparserParserListener) EnterAlterColConstraint(ctx *parser.AlterColConstraintContext) {}

// ExitAlterColConstraint is called when production alterColConstraint is exited.
func (s *GparserParserListener) ExitAlterColConstraint(ctx *parser.AlterColConstraintContext) {}

// EnterColumnConstraintType is called when production columnConstraintType is entered.
func (s *GparserParserListener) EnterColumnConstraintType(ctx *parser.ColumnConstraintTypeContext) {}

// ExitColumnConstraintType is called when production columnConstraintType is exited.
func (s *GparserParserListener) ExitColumnConstraintType(ctx *parser.ColumnConstraintTypeContext) {}

// EnterDefaultVal is called when production defaultVal is entered.
func (s *GparserParserListener) EnterDefaultVal(ctx *parser.DefaultValContext) {}

// ExitDefaultVal is called when production defaultVal is exited.
func (s *GparserParserListener) ExitDefaultVal(ctx *parser.DefaultValContext) {}

// EnterTableConstraintType is called when production tableConstraintType is entered.
func (s *GparserParserListener) EnterTableConstraintType(ctx *parser.TableConstraintTypeContext) {}

// ExitTableConstraintType is called when production tableConstraintType is exited.
func (s *GparserParserListener) ExitTableConstraintType(ctx *parser.TableConstraintTypeContext) {}

// EnterConstraintOptsCreate is called when production constraintOptsCreate is entered.
func (s *GparserParserListener) EnterConstraintOptsCreate(ctx *parser.ConstraintOptsCreateContext) {}

// ExitConstraintOptsCreate is called when production constraintOptsCreate is exited.
func (s *GparserParserListener) ExitConstraintOptsCreate(ctx *parser.ConstraintOptsCreateContext) {}

// EnterConstraintOptsAlter is called when production constraintOptsAlter is entered.
func (s *GparserParserListener) EnterConstraintOptsAlter(ctx *parser.ConstraintOptsAlterContext) {}

// ExitConstraintOptsAlter is called when production constraintOptsAlter is exited.
func (s *GparserParserListener) ExitConstraintOptsAlter(ctx *parser.ConstraintOptsAlterContext) {}

// EnterColumnNameColonType is called when production columnNameColonType is entered.
func (s *GparserParserListener) EnterColumnNameColonType(ctx *parser.ColumnNameColonTypeContext) {}

// ExitColumnNameColonType is called when production columnNameColonType is exited.
func (s *GparserParserListener) ExitColumnNameColonType(ctx *parser.ColumnNameColonTypeContext) {}

// EnterColType is called when production colType is entered.
func (s *GparserParserListener) EnterColType(ctx *parser.ColTypeContext) {}

// ExitColType is called when production colType is exited.
func (s *GparserParserListener) ExitColType(ctx *parser.ColTypeContext) {}

// EnterColTypeList is called when production colTypeList is entered.
func (s *GparserParserListener) EnterColTypeList(ctx *parser.ColTypeListContext) {}

// ExitColTypeList is called when production colTypeList is exited.
func (s *GparserParserListener) ExitColTypeList(ctx *parser.ColTypeListContext) {}

// EnterType is called when production type is entered.
func (s *GparserParserListener) EnterType(ctx *parser.TypeContext) {}

// ExitType is called when production type is exited.
func (s *GparserParserListener) ExitType(ctx *parser.TypeContext) {}

// EnterPrimitiveType is called when production primitiveType is entered.
func (s *GparserParserListener) EnterPrimitiveType(ctx *parser.PrimitiveTypeContext) {}

// ExitPrimitiveType is called when production primitiveType is exited.
func (s *GparserParserListener) ExitPrimitiveType(ctx *parser.PrimitiveTypeContext) {}

// EnterListType is called when production listType is entered.
func (s *GparserParserListener) EnterListType(ctx *parser.ListTypeContext) {}

// ExitListType is called when production listType is exited.
func (s *GparserParserListener) ExitListType(ctx *parser.ListTypeContext) {}

// EnterStructType is called when production structType is entered.
func (s *GparserParserListener) EnterStructType(ctx *parser.StructTypeContext) {}

// ExitStructType is called when production structType is exited.
func (s *GparserParserListener) ExitStructType(ctx *parser.StructTypeContext) {}

// EnterMapType is called when production mapType is entered.
func (s *GparserParserListener) EnterMapType(ctx *parser.MapTypeContext) {}

// ExitMapType is called when production mapType is exited.
func (s *GparserParserListener) ExitMapType(ctx *parser.MapTypeContext) {}

// EnterUnionType is called when production unionType is entered.
func (s *GparserParserListener) EnterUnionType(ctx *parser.UnionTypeContext) {}

// ExitUnionType is called when production unionType is exited.
func (s *GparserParserListener) ExitUnionType(ctx *parser.UnionTypeContext) {}

// EnterSetOperator is called when production setOperator is entered.
func (s *GparserParserListener) EnterSetOperator(ctx *parser.SetOperatorContext) {}

// ExitSetOperator is called when production setOperator is exited.
func (s *GparserParserListener) ExitSetOperator(ctx *parser.SetOperatorContext) {}

// EnterQueryStatementExpression is called when production queryStatementExpression is entered.
func (s *GparserParserListener) EnterQueryStatementExpression(ctx *parser.QueryStatementExpressionContext) {
}

// ExitQueryStatementExpression is called when production queryStatementExpression is exited.
func (s *GparserParserListener) ExitQueryStatementExpression(ctx *parser.QueryStatementExpressionContext) {
}

// EnterQueryStatementExpressionBody is called when production queryStatementExpressionBody is entered.
func (s *GparserParserListener) EnterQueryStatementExpressionBody(ctx *parser.QueryStatementExpressionBodyContext) {
}

// ExitQueryStatementExpressionBody is called when production queryStatementExpressionBody is exited.
func (s *GparserParserListener) ExitQueryStatementExpressionBody(ctx *parser.QueryStatementExpressionBodyContext) {
}

// EnterWithClause is called when production withClause is entered.
func (s *GparserParserListener) EnterWithClause(ctx *parser.WithClauseContext) {}

// ExitWithClause is called when production withClause is exited.
func (s *GparserParserListener) ExitWithClause(ctx *parser.WithClauseContext) {}

// EnterCteStatement is called when production cteStatement is entered.
func (s *GparserParserListener) EnterCteStatement(ctx *parser.CteStatementContext) {}

// ExitCteStatement is called when production cteStatement is exited.
func (s *GparserParserListener) ExitCteStatement(ctx *parser.CteStatementContext) {}

// EnterFromStatement is called when production fromStatement is entered.
func (s *GparserParserListener) EnterFromStatement(ctx *parser.FromStatementContext) {}

// ExitFromStatement is called when production fromStatement is exited.
func (s *GparserParserListener) ExitFromStatement(ctx *parser.FromStatementContext) {}

// EnterSingleFromStatement is called when production singleFromStatement is entered.
func (s *GparserParserListener) EnterSingleFromStatement(ctx *parser.SingleFromStatementContext) {}

// ExitSingleFromStatement is called when production singleFromStatement is exited.
func (s *GparserParserListener) ExitSingleFromStatement(ctx *parser.SingleFromStatementContext) {}

// EnterRegularBody is called when production regularBody is entered.
func (s *GparserParserListener) EnterRegularBody(ctx *parser.RegularBodyContext) {}

// ExitRegularBody is called when production regularBody is exited.
func (s *GparserParserListener) ExitRegularBody(ctx *parser.RegularBodyContext) {}

// EnterAtomSelectStatement is called when production atomSelectStatement is entered.
func (s *GparserParserListener) EnterAtomSelectStatement(ctx *parser.AtomSelectStatementContext) {}

// ExitAtomSelectStatement is called when production atomSelectStatement is exited.
func (s *GparserParserListener) ExitAtomSelectStatement(ctx *parser.AtomSelectStatementContext) {}

// EnterSelectStatement is called when production selectStatement is entered.
func (s *GparserParserListener) EnterSelectStatement(ctx *parser.SelectStatementContext) {}

// ExitSelectStatement is called when production selectStatement is exited.
func (s *GparserParserListener) ExitSelectStatement(ctx *parser.SelectStatementContext) {}

// EnterSetOpSelectStatement is called when production setOpSelectStatement is entered.
func (s *GparserParserListener) EnterSetOpSelectStatement(ctx *parser.SetOpSelectStatementContext) {}

// ExitSetOpSelectStatement is called when production setOpSelectStatement is exited.
func (s *GparserParserListener) ExitSetOpSelectStatement(ctx *parser.SetOpSelectStatementContext) {}

// EnterSelectStatementWithCTE is called when production selectStatementWithCTE is entered.
func (s *GparserParserListener) EnterSelectStatementWithCTE(ctx *parser.SelectStatementWithCTEContext) {
}

// ExitSelectStatementWithCTE is called when production selectStatementWithCTE is exited.
func (s *GparserParserListener) ExitSelectStatementWithCTE(ctx *parser.SelectStatementWithCTEContext) {
}

// EnterBody is called when production body is entered.
func (s *GparserParserListener) EnterBody(ctx *parser.BodyContext) {}

// ExitBody is called when production body is exited.
func (s *GparserParserListener) ExitBody(ctx *parser.BodyContext) {}

// EnterInsertClause is called when production insertClause is entered.
func (s *GparserParserListener) EnterInsertClause(ctx *parser.InsertClauseContext) {}

// ExitInsertClause is called when production insertClause is exited.
func (s *GparserParserListener) ExitInsertClause(ctx *parser.InsertClauseContext) {}

// EnterDestination is called when production destination is entered.
func (s *GparserParserListener) EnterDestination(ctx *parser.DestinationContext) {}

// ExitDestination is called when production destination is exited.
func (s *GparserParserListener) ExitDestination(ctx *parser.DestinationContext) {}

// EnterLimitClause is called when production limitClause is entered.
func (s *GparserParserListener) EnterLimitClause(ctx *parser.LimitClauseContext) {}

// ExitLimitClause is called when production limitClause is exited.
func (s *GparserParserListener) ExitLimitClause(ctx *parser.LimitClauseContext) {}

// EnterDeleteStatement is called when production deleteStatement is entered.
func (s *GparserParserListener) EnterDeleteStatement(ctx *parser.DeleteStatementContext) {}

// ExitDeleteStatement is called when production deleteStatement is exited.
func (s *GparserParserListener) ExitDeleteStatement(ctx *parser.DeleteStatementContext) {}

// EnterColumnAssignmentClause is called when production columnAssignmentClause is entered.
func (s *GparserParserListener) EnterColumnAssignmentClause(ctx *parser.ColumnAssignmentClauseContext) {
}

// ExitColumnAssignmentClause is called when production columnAssignmentClause is exited.
func (s *GparserParserListener) ExitColumnAssignmentClause(ctx *parser.ColumnAssignmentClauseContext) {
}

// EnterPrecedencePlusExpressionOrDefault is called when production precedencePlusExpressionOrDefault is entered.
func (s *GparserParserListener) EnterPrecedencePlusExpressionOrDefault(ctx *parser.PrecedencePlusExpressionOrDefaultContext) {
}

// ExitPrecedencePlusExpressionOrDefault is called when production precedencePlusExpressionOrDefault is exited.
func (s *GparserParserListener) ExitPrecedencePlusExpressionOrDefault(ctx *parser.PrecedencePlusExpressionOrDefaultContext) {
}

// EnterSetColumnsClause is called when production setColumnsClause is entered.
func (s *GparserParserListener) EnterSetColumnsClause(ctx *parser.SetColumnsClauseContext) {}

// ExitSetColumnsClause is called when production setColumnsClause is exited.
func (s *GparserParserListener) ExitSetColumnsClause(ctx *parser.SetColumnsClauseContext) {}

// EnterUpdateStatement is called when production updateStatement is entered.
func (s *GparserParserListener) EnterUpdateStatement(ctx *parser.UpdateStatementContext) {}

// ExitUpdateStatement is called when production updateStatement is exited.
func (s *GparserParserListener) ExitUpdateStatement(ctx *parser.UpdateStatementContext) {}

// EnterSqlTransactionStatement is called when production sqlTransactionStatement is entered.
func (s *GparserParserListener) EnterSqlTransactionStatement(ctx *parser.SqlTransactionStatementContext) {
}

// ExitSqlTransactionStatement is called when production sqlTransactionStatement is exited.
func (s *GparserParserListener) ExitSqlTransactionStatement(ctx *parser.SqlTransactionStatementContext) {
}

// EnterStartTransactionStatement is called when production startTransactionStatement is entered.
func (s *GparserParserListener) EnterStartTransactionStatement(ctx *parser.StartTransactionStatementContext) {
}

// ExitStartTransactionStatement is called when production startTransactionStatement is exited.
func (s *GparserParserListener) ExitStartTransactionStatement(ctx *parser.StartTransactionStatementContext) {
}

// EnterTransactionMode is called when production transactionMode is entered.
func (s *GparserParserListener) EnterTransactionMode(ctx *parser.TransactionModeContext) {}

// ExitTransactionMode is called when production transactionMode is exited.
func (s *GparserParserListener) ExitTransactionMode(ctx *parser.TransactionModeContext) {}

// EnterTransactionAccessMode is called when production transactionAccessMode is entered.
func (s *GparserParserListener) EnterTransactionAccessMode(ctx *parser.TransactionAccessModeContext) {
}

// ExitTransactionAccessMode is called when production transactionAccessMode is exited.
func (s *GparserParserListener) ExitTransactionAccessMode(ctx *parser.TransactionAccessModeContext) {}

// EnterIsolationLevel is called when production isolationLevel is entered.
func (s *GparserParserListener) EnterIsolationLevel(ctx *parser.IsolationLevelContext) {}

// ExitIsolationLevel is called when production isolationLevel is exited.
func (s *GparserParserListener) ExitIsolationLevel(ctx *parser.IsolationLevelContext) {}

// EnterLevelOfIsolation is called when production levelOfIsolation is entered.
func (s *GparserParserListener) EnterLevelOfIsolation(ctx *parser.LevelOfIsolationContext) {}

// ExitLevelOfIsolation is called when production levelOfIsolation is exited.
func (s *GparserParserListener) ExitLevelOfIsolation(ctx *parser.LevelOfIsolationContext) {}

// EnterCommitStatement is called when production commitStatement is entered.
func (s *GparserParserListener) EnterCommitStatement(ctx *parser.CommitStatementContext) {}

// ExitCommitStatement is called when production commitStatement is exited.
func (s *GparserParserListener) ExitCommitStatement(ctx *parser.CommitStatementContext) {}

// EnterRollbackStatement is called when production rollbackStatement is entered.
func (s *GparserParserListener) EnterRollbackStatement(ctx *parser.RollbackStatementContext) {}

// ExitRollbackStatement is called when production rollbackStatement is exited.
func (s *GparserParserListener) ExitRollbackStatement(ctx *parser.RollbackStatementContext) {}

// EnterSetAutoCommitStatement is called when production setAutoCommitStatement is entered.
func (s *GparserParserListener) EnterSetAutoCommitStatement(ctx *parser.SetAutoCommitStatementContext) {
}

// ExitSetAutoCommitStatement is called when production setAutoCommitStatement is exited.
func (s *GparserParserListener) ExitSetAutoCommitStatement(ctx *parser.SetAutoCommitStatementContext) {
}

// EnterAbortTransactionStatement is called when production abortTransactionStatement is entered.
func (s *GparserParserListener) EnterAbortTransactionStatement(ctx *parser.AbortTransactionStatementContext) {
}

// ExitAbortTransactionStatement is called when production abortTransactionStatement is exited.
func (s *GparserParserListener) ExitAbortTransactionStatement(ctx *parser.AbortTransactionStatementContext) {
}

// EnterAbortCompactionStatement is called when production abortCompactionStatement is entered.
func (s *GparserParserListener) EnterAbortCompactionStatement(ctx *parser.AbortCompactionStatementContext) {
}

// ExitAbortCompactionStatement is called when production abortCompactionStatement is exited.
func (s *GparserParserListener) ExitAbortCompactionStatement(ctx *parser.AbortCompactionStatementContext) {
}

// EnterMergeStatement is called when production mergeStatement is entered.
func (s *GparserParserListener) EnterMergeStatement(ctx *parser.MergeStatementContext) {}

// ExitMergeStatement is called when production mergeStatement is exited.
func (s *GparserParserListener) ExitMergeStatement(ctx *parser.MergeStatementContext) {}

// EnterWhenClauses is called when production whenClauses is entered.
func (s *GparserParserListener) EnterWhenClauses(ctx *parser.WhenClausesContext) {}

// ExitWhenClauses is called when production whenClauses is exited.
func (s *GparserParserListener) ExitWhenClauses(ctx *parser.WhenClausesContext) {}

// EnterWhenNotMatchedClause is called when production whenNotMatchedClause is entered.
func (s *GparserParserListener) EnterWhenNotMatchedClause(ctx *parser.WhenNotMatchedClauseContext) {}

// ExitWhenNotMatchedClause is called when production whenNotMatchedClause is exited.
func (s *GparserParserListener) ExitWhenNotMatchedClause(ctx *parser.WhenNotMatchedClauseContext) {}

// EnterWhenMatchedAndClause is called when production whenMatchedAndClause is entered.
func (s *GparserParserListener) EnterWhenMatchedAndClause(ctx *parser.WhenMatchedAndClauseContext) {}

// ExitWhenMatchedAndClause is called when production whenMatchedAndClause is exited.
func (s *GparserParserListener) ExitWhenMatchedAndClause(ctx *parser.WhenMatchedAndClauseContext) {}

// EnterWhenMatchedThenClause is called when production whenMatchedThenClause is entered.
func (s *GparserParserListener) EnterWhenMatchedThenClause(ctx *parser.WhenMatchedThenClauseContext) {
}

// ExitWhenMatchedThenClause is called when production whenMatchedThenClause is exited.
func (s *GparserParserListener) ExitWhenMatchedThenClause(ctx *parser.WhenMatchedThenClauseContext) {}

// EnterUpdateOrDelete is called when production updateOrDelete is entered.
func (s *GparserParserListener) EnterUpdateOrDelete(ctx *parser.UpdateOrDeleteContext) {}

// ExitUpdateOrDelete is called when production updateOrDelete is exited.
func (s *GparserParserListener) ExitUpdateOrDelete(ctx *parser.UpdateOrDeleteContext) {}

// EnterKillQueryStatement is called when production killQueryStatement is entered.
func (s *GparserParserListener) EnterKillQueryStatement(ctx *parser.KillQueryStatementContext) {}

// ExitKillQueryStatement is called when production killQueryStatement is exited.
func (s *GparserParserListener) ExitKillQueryStatement(ctx *parser.KillQueryStatementContext) {}

// EnterCompactionId is called when production compactionId is entered.
func (s *GparserParserListener) EnterCompactionId(ctx *parser.CompactionIdContext) {}

// ExitCompactionId is called when production compactionId is exited.
func (s *GparserParserListener) ExitCompactionId(ctx *parser.CompactionIdContext) {}

// EnterCompactionPool is called when production compactionPool is entered.
func (s *GparserParserListener) EnterCompactionPool(ctx *parser.CompactionPoolContext) {}

// ExitCompactionPool is called when production compactionPool is exited.
func (s *GparserParserListener) ExitCompactionPool(ctx *parser.CompactionPoolContext) {}

// EnterCompactionType is called when production compactionType is entered.
func (s *GparserParserListener) EnterCompactionType(ctx *parser.CompactionTypeContext) {}

// ExitCompactionType is called when production compactionType is exited.
func (s *GparserParserListener) ExitCompactionType(ctx *parser.CompactionTypeContext) {}

// EnterCompactionStatus is called when production compactionStatus is entered.
func (s *GparserParserListener) EnterCompactionStatus(ctx *parser.CompactionStatusContext) {}

// ExitCompactionStatus is called when production compactionStatus is exited.
func (s *GparserParserListener) ExitCompactionStatus(ctx *parser.CompactionStatusContext) {}

// EnterAlterStatement is called when production alterStatement is entered.
func (s *GparserParserListener) EnterAlterStatement(ctx *parser.AlterStatementContext) {}

// ExitAlterStatement is called when production alterStatement is exited.
func (s *GparserParserListener) ExitAlterStatement(ctx *parser.AlterStatementContext) {}

// EnterAlterTableStatementSuffix is called when production alterTableStatementSuffix is entered.
func (s *GparserParserListener) EnterAlterTableStatementSuffix(ctx *parser.AlterTableStatementSuffixContext) {
}

// ExitAlterTableStatementSuffix is called when production alterTableStatementSuffix is exited.
func (s *GparserParserListener) ExitAlterTableStatementSuffix(ctx *parser.AlterTableStatementSuffixContext) {
}

// EnterAlterTblPartitionStatementSuffix is called when production alterTblPartitionStatementSuffix is entered.
func (s *GparserParserListener) EnterAlterTblPartitionStatementSuffix(ctx *parser.AlterTblPartitionStatementSuffixContext) {
}

// ExitAlterTblPartitionStatementSuffix is called when production alterTblPartitionStatementSuffix is exited.
func (s *GparserParserListener) ExitAlterTblPartitionStatementSuffix(ctx *parser.AlterTblPartitionStatementSuffixContext) {
}

// EnterAlterStatementPartitionKeyType is called when production alterStatementPartitionKeyType is entered.
func (s *GparserParserListener) EnterAlterStatementPartitionKeyType(ctx *parser.AlterStatementPartitionKeyTypeContext) {
}

// ExitAlterStatementPartitionKeyType is called when production alterStatementPartitionKeyType is exited.
func (s *GparserParserListener) ExitAlterStatementPartitionKeyType(ctx *parser.AlterStatementPartitionKeyTypeContext) {
}

// EnterAlterViewStatementSuffix is called when production alterViewStatementSuffix is entered.
func (s *GparserParserListener) EnterAlterViewStatementSuffix(ctx *parser.AlterViewStatementSuffixContext) {
}

// ExitAlterViewStatementSuffix is called when production alterViewStatementSuffix is exited.
func (s *GparserParserListener) ExitAlterViewStatementSuffix(ctx *parser.AlterViewStatementSuffixContext) {
}

// EnterAlterMaterializedViewStatementSuffix is called when production alterMaterializedViewStatementSuffix is entered.
func (s *GparserParserListener) EnterAlterMaterializedViewStatementSuffix(ctx *parser.AlterMaterializedViewStatementSuffixContext) {
}

// ExitAlterMaterializedViewStatementSuffix is called when production alterMaterializedViewStatementSuffix is exited.
func (s *GparserParserListener) ExitAlterMaterializedViewStatementSuffix(ctx *parser.AlterMaterializedViewStatementSuffixContext) {
}

// EnterAlterMaterializedViewSuffixRewrite is called when production alterMaterializedViewSuffixRewrite is entered.
func (s *GparserParserListener) EnterAlterMaterializedViewSuffixRewrite(ctx *parser.AlterMaterializedViewSuffixRewriteContext) {
}

// ExitAlterMaterializedViewSuffixRewrite is called when production alterMaterializedViewSuffixRewrite is exited.
func (s *GparserParserListener) ExitAlterMaterializedViewSuffixRewrite(ctx *parser.AlterMaterializedViewSuffixRewriteContext) {
}

// EnterAlterMaterializedViewSuffixRebuild is called when production alterMaterializedViewSuffixRebuild is entered.
func (s *GparserParserListener) EnterAlterMaterializedViewSuffixRebuild(ctx *parser.AlterMaterializedViewSuffixRebuildContext) {
}

// ExitAlterMaterializedViewSuffixRebuild is called when production alterMaterializedViewSuffixRebuild is exited.
func (s *GparserParserListener) ExitAlterMaterializedViewSuffixRebuild(ctx *parser.AlterMaterializedViewSuffixRebuildContext) {
}

// EnterAlterDatabaseStatementSuffix is called when production alterDatabaseStatementSuffix is entered.
func (s *GparserParserListener) EnterAlterDatabaseStatementSuffix(ctx *parser.AlterDatabaseStatementSuffixContext) {
}

// ExitAlterDatabaseStatementSuffix is called when production alterDatabaseStatementSuffix is exited.
func (s *GparserParserListener) ExitAlterDatabaseStatementSuffix(ctx *parser.AlterDatabaseStatementSuffixContext) {
}

// EnterAlterDatabaseSuffixProperties is called when production alterDatabaseSuffixProperties is entered.
func (s *GparserParserListener) EnterAlterDatabaseSuffixProperties(ctx *parser.AlterDatabaseSuffixPropertiesContext) {
}

// ExitAlterDatabaseSuffixProperties is called when production alterDatabaseSuffixProperties is exited.
func (s *GparserParserListener) ExitAlterDatabaseSuffixProperties(ctx *parser.AlterDatabaseSuffixPropertiesContext) {
}

// EnterAlterDatabaseSuffixSetOwner is called when production alterDatabaseSuffixSetOwner is entered.
func (s *GparserParserListener) EnterAlterDatabaseSuffixSetOwner(ctx *parser.AlterDatabaseSuffixSetOwnerContext) {
}

// ExitAlterDatabaseSuffixSetOwner is called when production alterDatabaseSuffixSetOwner is exited.
func (s *GparserParserListener) ExitAlterDatabaseSuffixSetOwner(ctx *parser.AlterDatabaseSuffixSetOwnerContext) {
}

// EnterAlterDatabaseSuffixSetLocation is called when production alterDatabaseSuffixSetLocation is entered.
func (s *GparserParserListener) EnterAlterDatabaseSuffixSetLocation(ctx *parser.AlterDatabaseSuffixSetLocationContext) {
}

// ExitAlterDatabaseSuffixSetLocation is called when production alterDatabaseSuffixSetLocation is exited.
func (s *GparserParserListener) ExitAlterDatabaseSuffixSetLocation(ctx *parser.AlterDatabaseSuffixSetLocationContext) {
}

// EnterAlterDatabaseSuffixSetManagedLocation is called when production alterDatabaseSuffixSetManagedLocation is entered.
func (s *GparserParserListener) EnterAlterDatabaseSuffixSetManagedLocation(ctx *parser.AlterDatabaseSuffixSetManagedLocationContext) {
}

// ExitAlterDatabaseSuffixSetManagedLocation is called when production alterDatabaseSuffixSetManagedLocation is exited.
func (s *GparserParserListener) ExitAlterDatabaseSuffixSetManagedLocation(ctx *parser.AlterDatabaseSuffixSetManagedLocationContext) {
}

// EnterAlterStatementSuffixRename is called when production alterStatementSuffixRename is entered.
func (s *GparserParserListener) EnterAlterStatementSuffixRename(ctx *parser.AlterStatementSuffixRenameContext) {
}

// ExitAlterStatementSuffixRename is called when production alterStatementSuffixRename is exited.
func (s *GparserParserListener) ExitAlterStatementSuffixRename(ctx *parser.AlterStatementSuffixRenameContext) {
}

// EnterAlterStatementSuffixAddCol is called when production alterStatementSuffixAddCol is entered.
func (s *GparserParserListener) EnterAlterStatementSuffixAddCol(ctx *parser.AlterStatementSuffixAddColContext) {
}

// ExitAlterStatementSuffixAddCol is called when production alterStatementSuffixAddCol is exited.
func (s *GparserParserListener) ExitAlterStatementSuffixAddCol(ctx *parser.AlterStatementSuffixAddColContext) {
}

// EnterAlterStatementSuffixAddConstraint is called when production alterStatementSuffixAddConstraint is entered.
func (s *GparserParserListener) EnterAlterStatementSuffixAddConstraint(ctx *parser.AlterStatementSuffixAddConstraintContext) {
}

// ExitAlterStatementSuffixAddConstraint is called when production alterStatementSuffixAddConstraint is exited.
func (s *GparserParserListener) ExitAlterStatementSuffixAddConstraint(ctx *parser.AlterStatementSuffixAddConstraintContext) {
}

// EnterAlterStatementSuffixUpdateColumns is called when production alterStatementSuffixUpdateColumns is entered.
func (s *GparserParserListener) EnterAlterStatementSuffixUpdateColumns(ctx *parser.AlterStatementSuffixUpdateColumnsContext) {
}

// ExitAlterStatementSuffixUpdateColumns is called when production alterStatementSuffixUpdateColumns is exited.
func (s *GparserParserListener) ExitAlterStatementSuffixUpdateColumns(ctx *parser.AlterStatementSuffixUpdateColumnsContext) {
}

// EnterAlterStatementSuffixDropConstraint is called when production alterStatementSuffixDropConstraint is entered.
func (s *GparserParserListener) EnterAlterStatementSuffixDropConstraint(ctx *parser.AlterStatementSuffixDropConstraintContext) {
}

// ExitAlterStatementSuffixDropConstraint is called when production alterStatementSuffixDropConstraint is exited.
func (s *GparserParserListener) ExitAlterStatementSuffixDropConstraint(ctx *parser.AlterStatementSuffixDropConstraintContext) {
}

// EnterAlterStatementSuffixRenameCol is called when production alterStatementSuffixRenameCol is entered.
func (s *GparserParserListener) EnterAlterStatementSuffixRenameCol(ctx *parser.AlterStatementSuffixRenameColContext) {
}

// ExitAlterStatementSuffixRenameCol is called when production alterStatementSuffixRenameCol is exited.
func (s *GparserParserListener) ExitAlterStatementSuffixRenameCol(ctx *parser.AlterStatementSuffixRenameColContext) {
}

// EnterAlterStatementSuffixUpdateStatsCol is called when production alterStatementSuffixUpdateStatsCol is entered.
func (s *GparserParserListener) EnterAlterStatementSuffixUpdateStatsCol(ctx *parser.AlterStatementSuffixUpdateStatsColContext) {
}

// ExitAlterStatementSuffixUpdateStatsCol is called when production alterStatementSuffixUpdateStatsCol is exited.
func (s *GparserParserListener) ExitAlterStatementSuffixUpdateStatsCol(ctx *parser.AlterStatementSuffixUpdateStatsColContext) {
}

// EnterAlterStatementSuffixUpdateStats is called when production alterStatementSuffixUpdateStats is entered.
func (s *GparserParserListener) EnterAlterStatementSuffixUpdateStats(ctx *parser.AlterStatementSuffixUpdateStatsContext) {
}

// ExitAlterStatementSuffixUpdateStats is called when production alterStatementSuffixUpdateStats is exited.
func (s *GparserParserListener) ExitAlterStatementSuffixUpdateStats(ctx *parser.AlterStatementSuffixUpdateStatsContext) {
}

// EnterAlterStatementChangeColPosition is called when production alterStatementChangeColPosition is entered.
func (s *GparserParserListener) EnterAlterStatementChangeColPosition(ctx *parser.AlterStatementChangeColPositionContext) {
}

// ExitAlterStatementChangeColPosition is called when production alterStatementChangeColPosition is exited.
func (s *GparserParserListener) ExitAlterStatementChangeColPosition(ctx *parser.AlterStatementChangeColPositionContext) {
}

// EnterAlterStatementSuffixAddPartitions is called when production alterStatementSuffixAddPartitions is entered.
func (s *GparserParserListener) EnterAlterStatementSuffixAddPartitions(ctx *parser.AlterStatementSuffixAddPartitionsContext) {
}

// ExitAlterStatementSuffixAddPartitions is called when production alterStatementSuffixAddPartitions is exited.
func (s *GparserParserListener) ExitAlterStatementSuffixAddPartitions(ctx *parser.AlterStatementSuffixAddPartitionsContext) {
}

// EnterAlterStatementSuffixAddPartitionsElement is called when production alterStatementSuffixAddPartitionsElement is entered.
func (s *GparserParserListener) EnterAlterStatementSuffixAddPartitionsElement(ctx *parser.AlterStatementSuffixAddPartitionsElementContext) {
}

// ExitAlterStatementSuffixAddPartitionsElement is called when production alterStatementSuffixAddPartitionsElement is exited.
func (s *GparserParserListener) ExitAlterStatementSuffixAddPartitionsElement(ctx *parser.AlterStatementSuffixAddPartitionsElementContext) {
}

// EnterAlterStatementSuffixTouch is called when production alterStatementSuffixTouch is entered.
func (s *GparserParserListener) EnterAlterStatementSuffixTouch(ctx *parser.AlterStatementSuffixTouchContext) {
}

// ExitAlterStatementSuffixTouch is called when production alterStatementSuffixTouch is exited.
func (s *GparserParserListener) ExitAlterStatementSuffixTouch(ctx *parser.AlterStatementSuffixTouchContext) {
}

// EnterAlterStatementSuffixArchive is called when production alterStatementSuffixArchive is entered.
func (s *GparserParserListener) EnterAlterStatementSuffixArchive(ctx *parser.AlterStatementSuffixArchiveContext) {
}

// ExitAlterStatementSuffixArchive is called when production alterStatementSuffixArchive is exited.
func (s *GparserParserListener) ExitAlterStatementSuffixArchive(ctx *parser.AlterStatementSuffixArchiveContext) {
}

// EnterAlterStatementSuffixUnArchive is called when production alterStatementSuffixUnArchive is entered.
func (s *GparserParserListener) EnterAlterStatementSuffixUnArchive(ctx *parser.AlterStatementSuffixUnArchiveContext) {
}

// ExitAlterStatementSuffixUnArchive is called when production alterStatementSuffixUnArchive is exited.
func (s *GparserParserListener) ExitAlterStatementSuffixUnArchive(ctx *parser.AlterStatementSuffixUnArchiveContext) {
}

// EnterPartitionLocation is called when production partitionLocation is entered.
func (s *GparserParserListener) EnterPartitionLocation(ctx *parser.PartitionLocationContext) {}

// ExitPartitionLocation is called when production partitionLocation is exited.
func (s *GparserParserListener) ExitPartitionLocation(ctx *parser.PartitionLocationContext) {}

// EnterAlterStatementSuffixDropPartitions is called when production alterStatementSuffixDropPartitions is entered.
func (s *GparserParserListener) EnterAlterStatementSuffixDropPartitions(ctx *parser.AlterStatementSuffixDropPartitionsContext) {
}

// ExitAlterStatementSuffixDropPartitions is called when production alterStatementSuffixDropPartitions is exited.
func (s *GparserParserListener) ExitAlterStatementSuffixDropPartitions(ctx *parser.AlterStatementSuffixDropPartitionsContext) {
}

// EnterAlterStatementSuffixProperties is called when production alterStatementSuffixProperties is entered.
func (s *GparserParserListener) EnterAlterStatementSuffixProperties(ctx *parser.AlterStatementSuffixPropertiesContext) {
}

// ExitAlterStatementSuffixProperties is called when production alterStatementSuffixProperties is exited.
func (s *GparserParserListener) ExitAlterStatementSuffixProperties(ctx *parser.AlterStatementSuffixPropertiesContext) {
}

// EnterAlterViewSuffixProperties is called when production alterViewSuffixProperties is entered.
func (s *GparserParserListener) EnterAlterViewSuffixProperties(ctx *parser.AlterViewSuffixPropertiesContext) {
}

// ExitAlterViewSuffixProperties is called when production alterViewSuffixProperties is exited.
func (s *GparserParserListener) ExitAlterViewSuffixProperties(ctx *parser.AlterViewSuffixPropertiesContext) {
}

// EnterAlterStatementSuffixSerdeProperties is called when production alterStatementSuffixSerdeProperties is entered.
func (s *GparserParserListener) EnterAlterStatementSuffixSerdeProperties(ctx *parser.AlterStatementSuffixSerdePropertiesContext) {
}

// ExitAlterStatementSuffixSerdeProperties is called when production alterStatementSuffixSerdeProperties is exited.
func (s *GparserParserListener) ExitAlterStatementSuffixSerdeProperties(ctx *parser.AlterStatementSuffixSerdePropertiesContext) {
}

// EnterTablePartitionPrefix is called when production tablePartitionPrefix is entered.
func (s *GparserParserListener) EnterTablePartitionPrefix(ctx *parser.TablePartitionPrefixContext) {}

// ExitTablePartitionPrefix is called when production tablePartitionPrefix is exited.
func (s *GparserParserListener) ExitTablePartitionPrefix(ctx *parser.TablePartitionPrefixContext) {}

// EnterAlterStatementSuffixFileFormat is called when production alterStatementSuffixFileFormat is entered.
func (s *GparserParserListener) EnterAlterStatementSuffixFileFormat(ctx *parser.AlterStatementSuffixFileFormatContext) {
}

// ExitAlterStatementSuffixFileFormat is called when production alterStatementSuffixFileFormat is exited.
func (s *GparserParserListener) ExitAlterStatementSuffixFileFormat(ctx *parser.AlterStatementSuffixFileFormatContext) {
}

// EnterAlterStatementSuffixClusterbySortby is called when production alterStatementSuffixClusterbySortby is entered.
func (s *GparserParserListener) EnterAlterStatementSuffixClusterbySortby(ctx *parser.AlterStatementSuffixClusterbySortbyContext) {
}

// ExitAlterStatementSuffixClusterbySortby is called when production alterStatementSuffixClusterbySortby is exited.
func (s *GparserParserListener) ExitAlterStatementSuffixClusterbySortby(ctx *parser.AlterStatementSuffixClusterbySortbyContext) {
}

// EnterAlterTblPartitionStatementSuffixSkewedLocation is called when production alterTblPartitionStatementSuffixSkewedLocation is entered.
func (s *GparserParserListener) EnterAlterTblPartitionStatementSuffixSkewedLocation(ctx *parser.AlterTblPartitionStatementSuffixSkewedLocationContext) {
}

// ExitAlterTblPartitionStatementSuffixSkewedLocation is called when production alterTblPartitionStatementSuffixSkewedLocation is exited.
func (s *GparserParserListener) ExitAlterTblPartitionStatementSuffixSkewedLocation(ctx *parser.AlterTblPartitionStatementSuffixSkewedLocationContext) {
}

// EnterSkewedLocations is called when production skewedLocations is entered.
func (s *GparserParserListener) EnterSkewedLocations(ctx *parser.SkewedLocationsContext) {}

// ExitSkewedLocations is called when production skewedLocations is exited.
func (s *GparserParserListener) ExitSkewedLocations(ctx *parser.SkewedLocationsContext) {}

// EnterSkewedLocationsList is called when production skewedLocationsList is entered.
func (s *GparserParserListener) EnterSkewedLocationsList(ctx *parser.SkewedLocationsListContext) {}

// ExitSkewedLocationsList is called when production skewedLocationsList is exited.
func (s *GparserParserListener) ExitSkewedLocationsList(ctx *parser.SkewedLocationsListContext) {}

// EnterSkewedLocationMap is called when production skewedLocationMap is entered.
func (s *GparserParserListener) EnterSkewedLocationMap(ctx *parser.SkewedLocationMapContext) {}

// ExitSkewedLocationMap is called when production skewedLocationMap is exited.
func (s *GparserParserListener) ExitSkewedLocationMap(ctx *parser.SkewedLocationMapContext) {}

// EnterAlterStatementSuffixLocation is called when production alterStatementSuffixLocation is entered.
func (s *GparserParserListener) EnterAlterStatementSuffixLocation(ctx *parser.AlterStatementSuffixLocationContext) {
}

// ExitAlterStatementSuffixLocation is called when production alterStatementSuffixLocation is exited.
func (s *GparserParserListener) ExitAlterStatementSuffixLocation(ctx *parser.AlterStatementSuffixLocationContext) {
}

// EnterAlterStatementSuffixSkewedby is called when production alterStatementSuffixSkewedby is entered.
func (s *GparserParserListener) EnterAlterStatementSuffixSkewedby(ctx *parser.AlterStatementSuffixSkewedbyContext) {
}

// ExitAlterStatementSuffixSkewedby is called when production alterStatementSuffixSkewedby is exited.
func (s *GparserParserListener) ExitAlterStatementSuffixSkewedby(ctx *parser.AlterStatementSuffixSkewedbyContext) {
}

// EnterAlterStatementSuffixExchangePartition is called when production alterStatementSuffixExchangePartition is entered.
func (s *GparserParserListener) EnterAlterStatementSuffixExchangePartition(ctx *parser.AlterStatementSuffixExchangePartitionContext) {
}

// ExitAlterStatementSuffixExchangePartition is called when production alterStatementSuffixExchangePartition is exited.
func (s *GparserParserListener) ExitAlterStatementSuffixExchangePartition(ctx *parser.AlterStatementSuffixExchangePartitionContext) {
}

// EnterAlterStatementSuffixRenamePart is called when production alterStatementSuffixRenamePart is entered.
func (s *GparserParserListener) EnterAlterStatementSuffixRenamePart(ctx *parser.AlterStatementSuffixRenamePartContext) {
}

// ExitAlterStatementSuffixRenamePart is called when production alterStatementSuffixRenamePart is exited.
func (s *GparserParserListener) ExitAlterStatementSuffixRenamePart(ctx *parser.AlterStatementSuffixRenamePartContext) {
}

// EnterAlterStatementSuffixStatsPart is called when production alterStatementSuffixStatsPart is entered.
func (s *GparserParserListener) EnterAlterStatementSuffixStatsPart(ctx *parser.AlterStatementSuffixStatsPartContext) {
}

// ExitAlterStatementSuffixStatsPart is called when production alterStatementSuffixStatsPart is exited.
func (s *GparserParserListener) ExitAlterStatementSuffixStatsPart(ctx *parser.AlterStatementSuffixStatsPartContext) {
}

// EnterAlterStatementSuffixMergeFiles is called when production alterStatementSuffixMergeFiles is entered.
func (s *GparserParserListener) EnterAlterStatementSuffixMergeFiles(ctx *parser.AlterStatementSuffixMergeFilesContext) {
}

// ExitAlterStatementSuffixMergeFiles is called when production alterStatementSuffixMergeFiles is exited.
func (s *GparserParserListener) ExitAlterStatementSuffixMergeFiles(ctx *parser.AlterStatementSuffixMergeFilesContext) {
}

// EnterAlterStatementSuffixBucketNum is called when production alterStatementSuffixBucketNum is entered.
func (s *GparserParserListener) EnterAlterStatementSuffixBucketNum(ctx *parser.AlterStatementSuffixBucketNumContext) {
}

// ExitAlterStatementSuffixBucketNum is called when production alterStatementSuffixBucketNum is exited.
func (s *GparserParserListener) ExitAlterStatementSuffixBucketNum(ctx *parser.AlterStatementSuffixBucketNumContext) {
}

// EnterBlocking is called when production blocking is entered.
func (s *GparserParserListener) EnterBlocking(ctx *parser.BlockingContext) {}

// ExitBlocking is called when production blocking is exited.
func (s *GparserParserListener) ExitBlocking(ctx *parser.BlockingContext) {}

// EnterCompactPool is called when production compactPool is entered.
func (s *GparserParserListener) EnterCompactPool(ctx *parser.CompactPoolContext) {}

// ExitCompactPool is called when production compactPool is exited.
func (s *GparserParserListener) ExitCompactPool(ctx *parser.CompactPoolContext) {}

// EnterAlterStatementSuffixCompact is called when production alterStatementSuffixCompact is entered.
func (s *GparserParserListener) EnterAlterStatementSuffixCompact(ctx *parser.AlterStatementSuffixCompactContext) {
}

// ExitAlterStatementSuffixCompact is called when production alterStatementSuffixCompact is exited.
func (s *GparserParserListener) ExitAlterStatementSuffixCompact(ctx *parser.AlterStatementSuffixCompactContext) {
}

// EnterAlterStatementSuffixSetOwner is called when production alterStatementSuffixSetOwner is entered.
func (s *GparserParserListener) EnterAlterStatementSuffixSetOwner(ctx *parser.AlterStatementSuffixSetOwnerContext) {
}

// ExitAlterStatementSuffixSetOwner is called when production alterStatementSuffixSetOwner is exited.
func (s *GparserParserListener) ExitAlterStatementSuffixSetOwner(ctx *parser.AlterStatementSuffixSetOwnerContext) {
}

// EnterAlterStatementSuffixSetPartSpec is called when production alterStatementSuffixSetPartSpec is entered.
func (s *GparserParserListener) EnterAlterStatementSuffixSetPartSpec(ctx *parser.AlterStatementSuffixSetPartSpecContext) {
}

// ExitAlterStatementSuffixSetPartSpec is called when production alterStatementSuffixSetPartSpec is exited.
func (s *GparserParserListener) ExitAlterStatementSuffixSetPartSpec(ctx *parser.AlterStatementSuffixSetPartSpecContext) {
}

// EnterAlterStatementSuffixExecute is called when production alterStatementSuffixExecute is entered.
func (s *GparserParserListener) EnterAlterStatementSuffixExecute(ctx *parser.AlterStatementSuffixExecuteContext) {
}

// ExitAlterStatementSuffixExecute is called when production alterStatementSuffixExecute is exited.
func (s *GparserParserListener) ExitAlterStatementSuffixExecute(ctx *parser.AlterStatementSuffixExecuteContext) {
}

// EnterFileFormat is called when production fileFormat is entered.
func (s *GparserParserListener) EnterFileFormat(ctx *parser.FileFormatContext) {}

// ExitFileFormat is called when production fileFormat is exited.
func (s *GparserParserListener) ExitFileFormat(ctx *parser.FileFormatContext) {}

// EnterAlterDataConnectorStatementSuffix is called when production alterDataConnectorStatementSuffix is entered.
func (s *GparserParserListener) EnterAlterDataConnectorStatementSuffix(ctx *parser.AlterDataConnectorStatementSuffixContext) {
}

// ExitAlterDataConnectorStatementSuffix is called when production alterDataConnectorStatementSuffix is exited.
func (s *GparserParserListener) ExitAlterDataConnectorStatementSuffix(ctx *parser.AlterDataConnectorStatementSuffixContext) {
}

// EnterAlterDataConnectorSuffixProperties is called when production alterDataConnectorSuffixProperties is entered.
func (s *GparserParserListener) EnterAlterDataConnectorSuffixProperties(ctx *parser.AlterDataConnectorSuffixPropertiesContext) {
}

// ExitAlterDataConnectorSuffixProperties is called when production alterDataConnectorSuffixProperties is exited.
func (s *GparserParserListener) ExitAlterDataConnectorSuffixProperties(ctx *parser.AlterDataConnectorSuffixPropertiesContext) {
}

// EnterAlterDataConnectorSuffixSetOwner is called when production alterDataConnectorSuffixSetOwner is entered.
func (s *GparserParserListener) EnterAlterDataConnectorSuffixSetOwner(ctx *parser.AlterDataConnectorSuffixSetOwnerContext) {
}

// ExitAlterDataConnectorSuffixSetOwner is called when production alterDataConnectorSuffixSetOwner is exited.
func (s *GparserParserListener) ExitAlterDataConnectorSuffixSetOwner(ctx *parser.AlterDataConnectorSuffixSetOwnerContext) {
}

// EnterAlterDataConnectorSuffixSetUrl is called when production alterDataConnectorSuffixSetUrl is entered.
func (s *GparserParserListener) EnterAlterDataConnectorSuffixSetUrl(ctx *parser.AlterDataConnectorSuffixSetUrlContext) {
}

// ExitAlterDataConnectorSuffixSetUrl is called when production alterDataConnectorSuffixSetUrl is exited.
func (s *GparserParserListener) ExitAlterDataConnectorSuffixSetUrl(ctx *parser.AlterDataConnectorSuffixSetUrlContext) {
}

// EnterLikeTableOrFile is called when production likeTableOrFile is entered.
func (s *GparserParserListener) EnterLikeTableOrFile(ctx *parser.LikeTableOrFileContext) {}

// ExitLikeTableOrFile is called when production likeTableOrFile is exited.
func (s *GparserParserListener) ExitLikeTableOrFile(ctx *parser.LikeTableOrFileContext) {}

// EnterCreateTableStatement is called when production createTableStatement is entered.
func (s *GparserParserListener) EnterCreateTableStatement(ctx *parser.CreateTableStatementContext) {}

// ExitCreateTableStatement is called when production createTableStatement is exited.
func (s *GparserParserListener) ExitCreateTableStatement(ctx *parser.CreateTableStatementContext) {}

// EnterCreateDataConnectorStatement is called when production createDataConnectorStatement is entered.
func (s *GparserParserListener) EnterCreateDataConnectorStatement(ctx *parser.CreateDataConnectorStatementContext) {
}

// ExitCreateDataConnectorStatement is called when production createDataConnectorStatement is exited.
func (s *GparserParserListener) ExitCreateDataConnectorStatement(ctx *parser.CreateDataConnectorStatementContext) {
}

// EnterDataConnectorComment is called when production dataConnectorComment is entered.
func (s *GparserParserListener) EnterDataConnectorComment(ctx *parser.DataConnectorCommentContext) {}

// ExitDataConnectorComment is called when production dataConnectorComment is exited.
func (s *GparserParserListener) ExitDataConnectorComment(ctx *parser.DataConnectorCommentContext) {}

// EnterDataConnectorUrl is called when production dataConnectorUrl is entered.
func (s *GparserParserListener) EnterDataConnectorUrl(ctx *parser.DataConnectorUrlContext) {}

// ExitDataConnectorUrl is called when production dataConnectorUrl is exited.
func (s *GparserParserListener) ExitDataConnectorUrl(ctx *parser.DataConnectorUrlContext) {}

// EnterDataConnectorType is called when production dataConnectorType is entered.
func (s *GparserParserListener) EnterDataConnectorType(ctx *parser.DataConnectorTypeContext) {}

// ExitDataConnectorType is called when production dataConnectorType is exited.
func (s *GparserParserListener) ExitDataConnectorType(ctx *parser.DataConnectorTypeContext) {}

// EnterDcProperties is called when production dcProperties is entered.
func (s *GparserParserListener) EnterDcProperties(ctx *parser.DcPropertiesContext) {}

// ExitDcProperties is called when production dcProperties is exited.
func (s *GparserParserListener) ExitDcProperties(ctx *parser.DcPropertiesContext) {}

// EnterDropDataConnectorStatement is called when production dropDataConnectorStatement is entered.
func (s *GparserParserListener) EnterDropDataConnectorStatement(ctx *parser.DropDataConnectorStatementContext) {
}

// ExitDropDataConnectorStatement is called when production dropDataConnectorStatement is exited.
func (s *GparserParserListener) ExitDropDataConnectorStatement(ctx *parser.DropDataConnectorStatementContext) {
}

// EnterTableAllColumns is called when production tableAllColumns is entered.
func (s *GparserParserListener) EnterTableAllColumns(ctx *parser.TableAllColumnsContext) {}

// ExitTableAllColumns is called when production tableAllColumns is exited.
func (s *GparserParserListener) ExitTableAllColumns(ctx *parser.TableAllColumnsContext) {}

// EnterTableOrColumn is called when production tableOrColumn is entered.
func (s *GparserParserListener) EnterTableOrColumn(ctx *parser.TableOrColumnContext) {}

// ExitTableOrColumn is called when production tableOrColumn is exited.
func (s *GparserParserListener) ExitTableOrColumn(ctx *parser.TableOrColumnContext) {}

// EnterDefaultValue is called when production defaultValue is entered.
func (s *GparserParserListener) EnterDefaultValue(ctx *parser.DefaultValueContext) {}

// ExitDefaultValue is called when production defaultValue is exited.
func (s *GparserParserListener) ExitDefaultValue(ctx *parser.DefaultValueContext) {}

// EnterExpressionList is called when production expressionList is entered.
func (s *GparserParserListener) EnterExpressionList(ctx *parser.ExpressionListContext) {}

// ExitExpressionList is called when production expressionList is exited.
func (s *GparserParserListener) ExitExpressionList(ctx *parser.ExpressionListContext) {}

// EnterAliasList is called when production aliasList is entered.
func (s *GparserParserListener) EnterAliasList(ctx *parser.AliasListContext) {}

// ExitAliasList is called when production aliasList is exited.
func (s *GparserParserListener) ExitAliasList(ctx *parser.AliasListContext) {}

// EnterFromClause is called when production fromClause is entered.
func (s *GparserParserListener) EnterFromClause(ctx *parser.FromClauseContext) {}

// ExitFromClause is called when production fromClause is exited.
func (s *GparserParserListener) ExitFromClause(ctx *parser.FromClauseContext) {}

// EnterFromSource is called when production fromSource is entered.
func (s *GparserParserListener) EnterFromSource(ctx *parser.FromSourceContext) {}

// ExitFromSource is called when production fromSource is exited.
func (s *GparserParserListener) ExitFromSource(ctx *parser.FromSourceContext) {}

// EnterAtomjoinSource is called when production atomjoinSource is entered.
func (s *GparserParserListener) EnterAtomjoinSource(ctx *parser.AtomjoinSourceContext) {}

// ExitAtomjoinSource is called when production atomjoinSource is exited.
func (s *GparserParserListener) ExitAtomjoinSource(ctx *parser.AtomjoinSourceContext) {}

// EnterJoinSource is called when production joinSource is entered.
func (s *GparserParserListener) EnterJoinSource(ctx *parser.JoinSourceContext) {}

// ExitJoinSource is called when production joinSource is exited.
func (s *GparserParserListener) ExitJoinSource(ctx *parser.JoinSourceContext) {}

// EnterJoinSourcePart is called when production joinSourcePart is entered.
func (s *GparserParserListener) EnterJoinSourcePart(ctx *parser.JoinSourcePartContext) {}

// ExitJoinSourcePart is called when production joinSourcePart is exited.
func (s *GparserParserListener) ExitJoinSourcePart(ctx *parser.JoinSourcePartContext) {}

// EnterUniqueJoinSource is called when production uniqueJoinSource is entered.
func (s *GparserParserListener) EnterUniqueJoinSource(ctx *parser.UniqueJoinSourceContext) {}

// ExitUniqueJoinSource is called when production uniqueJoinSource is exited.
func (s *GparserParserListener) ExitUniqueJoinSource(ctx *parser.UniqueJoinSourceContext) {}

// EnterUniqueJoinExpr is called when production uniqueJoinExpr is entered.
func (s *GparserParserListener) EnterUniqueJoinExpr(ctx *parser.UniqueJoinExprContext) {}

// ExitUniqueJoinExpr is called when production uniqueJoinExpr is exited.
func (s *GparserParserListener) ExitUniqueJoinExpr(ctx *parser.UniqueJoinExprContext) {}

// EnterUniqueJoinToken is called when production uniqueJoinToken is entered.
func (s *GparserParserListener) EnterUniqueJoinToken(ctx *parser.UniqueJoinTokenContext) {}

// ExitUniqueJoinToken is called when production uniqueJoinToken is exited.
func (s *GparserParserListener) ExitUniqueJoinToken(ctx *parser.UniqueJoinTokenContext) {}

// EnterJoinToken is called when production joinToken is entered.
func (s *GparserParserListener) EnterJoinToken(ctx *parser.JoinTokenContext) {}

// ExitJoinToken is called when production joinToken is exited.
func (s *GparserParserListener) ExitJoinToken(ctx *parser.JoinTokenContext) {}

// EnterLateralView is called when production lateralView is entered.
func (s *GparserParserListener) EnterLateralView(ctx *parser.LateralViewContext) {}

// ExitLateralView is called when production lateralView is exited.
func (s *GparserParserListener) ExitLateralView(ctx *parser.LateralViewContext) {}

// EnterTableAlias is called when production tableAlias is entered.
func (s *GparserParserListener) EnterTableAlias(ctx *parser.TableAliasContext) {}

// ExitTableAlias is called when production tableAlias is exited.
func (s *GparserParserListener) ExitTableAlias(ctx *parser.TableAliasContext) {}

// EnterTableBucketSample is called when production tableBucketSample is entered.
func (s *GparserParserListener) EnterTableBucketSample(ctx *parser.TableBucketSampleContext) {}

// ExitTableBucketSample is called when production tableBucketSample is exited.
func (s *GparserParserListener) ExitTableBucketSample(ctx *parser.TableBucketSampleContext) {}

// EnterSplitSample is called when production splitSample is entered.
func (s *GparserParserListener) EnterSplitSample(ctx *parser.SplitSampleContext) {}

// ExitSplitSample is called when production splitSample is exited.
func (s *GparserParserListener) ExitSplitSample(ctx *parser.SplitSampleContext) {}

// EnterTableSample is called when production tableSample is entered.
func (s *GparserParserListener) EnterTableSample(ctx *parser.TableSampleContext) {}

// ExitTableSample is called when production tableSample is exited.
func (s *GparserParserListener) ExitTableSample(ctx *parser.TableSampleContext) {}

// EnterTableSource is called when production tableSource is entered.
func (s *GparserParserListener) EnterTableSource(ctx *parser.TableSourceContext) {}

// ExitTableSource is called when production tableSource is exited.
func (s *GparserParserListener) ExitTableSource(ctx *parser.TableSourceContext) {}

// EnterAsOfClause is called when production asOfClause is entered.
func (s *GparserParserListener) EnterAsOfClause(ctx *parser.AsOfClauseContext) {}

// ExitAsOfClause is called when production asOfClause is exited.
func (s *GparserParserListener) ExitAsOfClause(ctx *parser.AsOfClauseContext) {}

// EnterUniqueJoinTableSource is called when production uniqueJoinTableSource is entered.
func (s *GparserParserListener) EnterUniqueJoinTableSource(ctx *parser.UniqueJoinTableSourceContext) {
}

// ExitUniqueJoinTableSource is called when production uniqueJoinTableSource is exited.
func (s *GparserParserListener) ExitUniqueJoinTableSource(ctx *parser.UniqueJoinTableSourceContext) {}

// EnterTableName is called when production tableName is entered.
func (s *GparserParserListener) EnterTableName(ctx *parser.TableNameContext) {}

// ExitTableName is called when production tableName is exited.
func (s *GparserParserListener) ExitTableName(ctx *parser.TableNameContext) {}

// EnterViewName is called when production viewName is entered.
func (s *GparserParserListener) EnterViewName(ctx *parser.ViewNameContext) {}

// ExitViewName is called when production viewName is exited.
func (s *GparserParserListener) ExitViewName(ctx *parser.ViewNameContext) {}

// EnterSubQuerySource is called when production subQuerySource is entered.
func (s *GparserParserListener) EnterSubQuerySource(ctx *parser.SubQuerySourceContext) {}

// ExitSubQuerySource is called when production subQuerySource is exited.
func (s *GparserParserListener) ExitSubQuerySource(ctx *parser.SubQuerySourceContext) {}

// EnterPartitioningSpec is called when production partitioningSpec is entered.
func (s *GparserParserListener) EnterPartitioningSpec(ctx *parser.PartitioningSpecContext) {}

// ExitPartitioningSpec is called when production partitioningSpec is exited.
func (s *GparserParserListener) ExitPartitioningSpec(ctx *parser.PartitioningSpecContext) {}

// EnterPartitionTableFunctionSource is called when production partitionTableFunctionSource is entered.
func (s *GparserParserListener) EnterPartitionTableFunctionSource(ctx *parser.PartitionTableFunctionSourceContext) {
}

// ExitPartitionTableFunctionSource is called when production partitionTableFunctionSource is exited.
func (s *GparserParserListener) ExitPartitionTableFunctionSource(ctx *parser.PartitionTableFunctionSourceContext) {
}

// EnterPartitionedTableFunction is called when production partitionedTableFunction is entered.
func (s *GparserParserListener) EnterPartitionedTableFunction(ctx *parser.PartitionedTableFunctionContext) {
}

// ExitPartitionedTableFunction is called when production partitionedTableFunction is exited.
func (s *GparserParserListener) ExitPartitionedTableFunction(ctx *parser.PartitionedTableFunctionContext) {
}

// EnterWhereClause is called when production whereClause is entered.
func (s *GparserParserListener) EnterWhereClause(ctx *parser.WhereClauseContext) {}

// ExitWhereClause is called when production whereClause is exited.
func (s *GparserParserListener) ExitWhereClause(ctx *parser.WhereClauseContext) {}

// EnterSearchCondition is called when production searchCondition is entered.
func (s *GparserParserListener) EnterSearchCondition(ctx *parser.SearchConditionContext) {}

// ExitSearchCondition is called when production searchCondition is exited.
func (s *GparserParserListener) ExitSearchCondition(ctx *parser.SearchConditionContext) {}

// EnterValuesSource is called when production valuesSource is entered.
func (s *GparserParserListener) EnterValuesSource(ctx *parser.ValuesSourceContext) {}

// ExitValuesSource is called when production valuesSource is exited.
func (s *GparserParserListener) ExitValuesSource(ctx *parser.ValuesSourceContext) {}

// EnterValuesClause is called when production valuesClause is entered.
func (s *GparserParserListener) EnterValuesClause(ctx *parser.ValuesClauseContext) {}

// ExitValuesClause is called when production valuesClause is exited.
func (s *GparserParserListener) ExitValuesClause(ctx *parser.ValuesClauseContext) {}

// EnterValuesTableConstructor is called when production valuesTableConstructor is entered.
func (s *GparserParserListener) EnterValuesTableConstructor(ctx *parser.ValuesTableConstructorContext) {
}

// ExitValuesTableConstructor is called when production valuesTableConstructor is exited.
func (s *GparserParserListener) ExitValuesTableConstructor(ctx *parser.ValuesTableConstructorContext) {
}

// EnterValueRowConstructor is called when production valueRowConstructor is entered.
func (s *GparserParserListener) EnterValueRowConstructor(ctx *parser.ValueRowConstructorContext) {}

// ExitValueRowConstructor is called when production valueRowConstructor is exited.
func (s *GparserParserListener) ExitValueRowConstructor(ctx *parser.ValueRowConstructorContext) {}

// EnterFirstValueRowConstructor is called when production firstValueRowConstructor is entered.
func (s *GparserParserListener) EnterFirstValueRowConstructor(ctx *parser.FirstValueRowConstructorContext) {
}

// ExitFirstValueRowConstructor is called when production firstValueRowConstructor is exited.
func (s *GparserParserListener) ExitFirstValueRowConstructor(ctx *parser.FirstValueRowConstructorContext) {
}

// EnterVirtualTableSource is called when production virtualTableSource is entered.
func (s *GparserParserListener) EnterVirtualTableSource(ctx *parser.VirtualTableSourceContext) {}

// ExitVirtualTableSource is called when production virtualTableSource is exited.
func (s *GparserParserListener) ExitVirtualTableSource(ctx *parser.VirtualTableSourceContext) {}

// EnterSelectClause is called when production selectClause is entered.
func (s *GparserParserListener) EnterSelectClause(ctx *parser.SelectClauseContext) {}

// ExitSelectClause is called when production selectClause is exited.
func (s *GparserParserListener) ExitSelectClause(ctx *parser.SelectClauseContext) {}

// EnterAll_distinct is called when production all_distinct is entered.
func (s *GparserParserListener) EnterAll_distinct(ctx *parser.All_distinctContext) {}

// ExitAll_distinct is called when production all_distinct is exited.
func (s *GparserParserListener) ExitAll_distinct(ctx *parser.All_distinctContext) {}

// EnterSelectList is called when production selectList is entered.
func (s *GparserParserListener) EnterSelectList(ctx *parser.SelectListContext) {}

// ExitSelectList is called when production selectList is exited.
func (s *GparserParserListener) ExitSelectList(ctx *parser.SelectListContext) {}

// EnterSelectTrfmClause is called when production selectTrfmClause is entered.
func (s *GparserParserListener) EnterSelectTrfmClause(ctx *parser.SelectTrfmClauseContext) {}

// ExitSelectTrfmClause is called when production selectTrfmClause is exited.
func (s *GparserParserListener) ExitSelectTrfmClause(ctx *parser.SelectTrfmClauseContext) {}

// EnterSelectItem is called when production selectItem is entered.
func (s *GparserParserListener) EnterSelectItem(ctx *parser.SelectItemContext) {}

// ExitSelectItem is called when production selectItem is exited.
func (s *GparserParserListener) ExitSelectItem(ctx *parser.SelectItemContext) {}

// EnterTrfmClause is called when production trfmClause is entered.
func (s *GparserParserListener) EnterTrfmClause(ctx *parser.TrfmClauseContext) {}

// ExitTrfmClause is called when production trfmClause is exited.
func (s *GparserParserListener) ExitTrfmClause(ctx *parser.TrfmClauseContext) {}

// EnterSelectExpression is called when production selectExpression is entered.
func (s *GparserParserListener) EnterSelectExpression(ctx *parser.SelectExpressionContext) {}

// ExitSelectExpression is called when production selectExpression is exited.
func (s *GparserParserListener) ExitSelectExpression(ctx *parser.SelectExpressionContext) {}

// EnterSelectExpressionList is called when production selectExpressionList is entered.
func (s *GparserParserListener) EnterSelectExpressionList(ctx *parser.SelectExpressionListContext) {}

// ExitSelectExpressionList is called when production selectExpressionList is exited.
func (s *GparserParserListener) ExitSelectExpressionList(ctx *parser.SelectExpressionListContext) {}

// EnterWindow_clause is called when production window_clause is entered.
func (s *GparserParserListener) EnterWindow_clause(ctx *parser.Window_clauseContext) {}

// ExitWindow_clause is called when production window_clause is exited.
func (s *GparserParserListener) ExitWindow_clause(ctx *parser.Window_clauseContext) {}

// EnterWindow_defn is called when production window_defn is entered.
func (s *GparserParserListener) EnterWindow_defn(ctx *parser.Window_defnContext) {}

// ExitWindow_defn is called when production window_defn is exited.
func (s *GparserParserListener) ExitWindow_defn(ctx *parser.Window_defnContext) {}

// EnterWindow_specification is called when production window_specification is entered.
func (s *GparserParserListener) EnterWindow_specification(ctx *parser.Window_specificationContext) {}

// ExitWindow_specification is called when production window_specification is exited.
func (s *GparserParserListener) ExitWindow_specification(ctx *parser.Window_specificationContext) {}

// EnterWindow_frame is called when production window_frame is entered.
func (s *GparserParserListener) EnterWindow_frame(ctx *parser.Window_frameContext) {}

// ExitWindow_frame is called when production window_frame is exited.
func (s *GparserParserListener) ExitWindow_frame(ctx *parser.Window_frameContext) {}

// EnterWindow_range_expression is called when production window_range_expression is entered.
func (s *GparserParserListener) EnterWindow_range_expression(ctx *parser.Window_range_expressionContext) {
}

// ExitWindow_range_expression is called when production window_range_expression is exited.
func (s *GparserParserListener) ExitWindow_range_expression(ctx *parser.Window_range_expressionContext) {
}

// EnterWindow_value_expression is called when production window_value_expression is entered.
func (s *GparserParserListener) EnterWindow_value_expression(ctx *parser.Window_value_expressionContext) {
}

// ExitWindow_value_expression is called when production window_value_expression is exited.
func (s *GparserParserListener) ExitWindow_value_expression(ctx *parser.Window_value_expressionContext) {
}

// EnterWindow_frame_start_boundary is called when production window_frame_start_boundary is entered.
func (s *GparserParserListener) EnterWindow_frame_start_boundary(ctx *parser.Window_frame_start_boundaryContext) {
}

// ExitWindow_frame_start_boundary is called when production window_frame_start_boundary is exited.
func (s *GparserParserListener) ExitWindow_frame_start_boundary(ctx *parser.Window_frame_start_boundaryContext) {
}

// EnterWindow_frame_boundary is called when production window_frame_boundary is entered.
func (s *GparserParserListener) EnterWindow_frame_boundary(ctx *parser.Window_frame_boundaryContext) {
}

// ExitWindow_frame_boundary is called when production window_frame_boundary is exited.
func (s *GparserParserListener) ExitWindow_frame_boundary(ctx *parser.Window_frame_boundaryContext) {}

// EnterGroupByClause is called when production groupByClause is entered.
func (s *GparserParserListener) EnterGroupByClause(ctx *parser.GroupByClauseContext) {}

// ExitGroupByClause is called when production groupByClause is exited.
func (s *GparserParserListener) ExitGroupByClause(ctx *parser.GroupByClauseContext) {}

// EnterGroupby_expression is called when production groupby_expression is entered.
func (s *GparserParserListener) EnterGroupby_expression(ctx *parser.Groupby_expressionContext) {}

// ExitGroupby_expression is called when production groupby_expression is exited.
func (s *GparserParserListener) ExitGroupby_expression(ctx *parser.Groupby_expressionContext) {}

// EnterGroupByEmpty is called when production groupByEmpty is entered.
func (s *GparserParserListener) EnterGroupByEmpty(ctx *parser.GroupByEmptyContext) {}

// ExitGroupByEmpty is called when production groupByEmpty is exited.
func (s *GparserParserListener) ExitGroupByEmpty(ctx *parser.GroupByEmptyContext) {}

// EnterRollupStandard is called when production rollupStandard is entered.
func (s *GparserParserListener) EnterRollupStandard(ctx *parser.RollupStandardContext) {}

// ExitRollupStandard is called when production rollupStandard is exited.
func (s *GparserParserListener) ExitRollupStandard(ctx *parser.RollupStandardContext) {}

// EnterRollupOldSyntax is called when production rollupOldSyntax is entered.
func (s *GparserParserListener) EnterRollupOldSyntax(ctx *parser.RollupOldSyntaxContext) {}

// ExitRollupOldSyntax is called when production rollupOldSyntax is exited.
func (s *GparserParserListener) ExitRollupOldSyntax(ctx *parser.RollupOldSyntaxContext) {}

// EnterGroupingSetExpression is called when production groupingSetExpression is entered.
func (s *GparserParserListener) EnterGroupingSetExpression(ctx *parser.GroupingSetExpressionContext) {
}

// ExitGroupingSetExpression is called when production groupingSetExpression is exited.
func (s *GparserParserListener) ExitGroupingSetExpression(ctx *parser.GroupingSetExpressionContext) {}

// EnterGroupingSetExpressionMultiple is called when production groupingSetExpressionMultiple is entered.
func (s *GparserParserListener) EnterGroupingSetExpressionMultiple(ctx *parser.GroupingSetExpressionMultipleContext) {
}

// ExitGroupingSetExpressionMultiple is called when production groupingSetExpressionMultiple is exited.
func (s *GparserParserListener) ExitGroupingSetExpressionMultiple(ctx *parser.GroupingSetExpressionMultipleContext) {
}

// EnterGroupingExpressionSingle is called when production groupingExpressionSingle is entered.
func (s *GparserParserListener) EnterGroupingExpressionSingle(ctx *parser.GroupingExpressionSingleContext) {
}

// ExitGroupingExpressionSingle is called when production groupingExpressionSingle is exited.
func (s *GparserParserListener) ExitGroupingExpressionSingle(ctx *parser.GroupingExpressionSingleContext) {
}

// EnterHavingClause is called when production havingClause is entered.
func (s *GparserParserListener) EnterHavingClause(ctx *parser.HavingClauseContext) {}

// ExitHavingClause is called when production havingClause is exited.
func (s *GparserParserListener) ExitHavingClause(ctx *parser.HavingClauseContext) {}

// EnterQualifyClause is called when production qualifyClause is entered.
func (s *GparserParserListener) EnterQualifyClause(ctx *parser.QualifyClauseContext) {}

// ExitQualifyClause is called when production qualifyClause is exited.
func (s *GparserParserListener) ExitQualifyClause(ctx *parser.QualifyClauseContext) {}

// EnterHavingCondition is called when production havingCondition is entered.
func (s *GparserParserListener) EnterHavingCondition(ctx *parser.HavingConditionContext) {}

// ExitHavingCondition is called when production havingCondition is exited.
func (s *GparserParserListener) ExitHavingCondition(ctx *parser.HavingConditionContext) {}

// EnterExpressionsInParenthesis is called when production expressionsInParenthesis is entered.
func (s *GparserParserListener) EnterExpressionsInParenthesis(ctx *parser.ExpressionsInParenthesisContext) {
}

// ExitExpressionsInParenthesis is called when production expressionsInParenthesis is exited.
func (s *GparserParserListener) ExitExpressionsInParenthesis(ctx *parser.ExpressionsInParenthesisContext) {
}

// EnterExpressionsNotInParenthesis is called when production expressionsNotInParenthesis is entered.
func (s *GparserParserListener) EnterExpressionsNotInParenthesis(ctx *parser.ExpressionsNotInParenthesisContext) {
}

// ExitExpressionsNotInParenthesis is called when production expressionsNotInParenthesis is exited.
func (s *GparserParserListener) ExitExpressionsNotInParenthesis(ctx *parser.ExpressionsNotInParenthesisContext) {
}

// EnterExpressionPart is called when production expressionPart is entered.
func (s *GparserParserListener) EnterExpressionPart(ctx *parser.ExpressionPartContext) {}

// ExitExpressionPart is called when production expressionPart is exited.
func (s *GparserParserListener) ExitExpressionPart(ctx *parser.ExpressionPartContext) {}

// EnterExpressionOrDefault is called when production expressionOrDefault is entered.
func (s *GparserParserListener) EnterExpressionOrDefault(ctx *parser.ExpressionOrDefaultContext) {}

// ExitExpressionOrDefault is called when production expressionOrDefault is exited.
func (s *GparserParserListener) ExitExpressionOrDefault(ctx *parser.ExpressionOrDefaultContext) {}

// EnterFirstExpressionsWithAlias is called when production firstExpressionsWithAlias is entered.
func (s *GparserParserListener) EnterFirstExpressionsWithAlias(ctx *parser.FirstExpressionsWithAliasContext) {
}

// ExitFirstExpressionsWithAlias is called when production firstExpressionsWithAlias is exited.
func (s *GparserParserListener) ExitFirstExpressionsWithAlias(ctx *parser.FirstExpressionsWithAliasContext) {
}

// EnterExpressionWithAlias is called when production expressionWithAlias is entered.
func (s *GparserParserListener) EnterExpressionWithAlias(ctx *parser.ExpressionWithAliasContext) {}

// ExitExpressionWithAlias is called when production expressionWithAlias is exited.
func (s *GparserParserListener) ExitExpressionWithAlias(ctx *parser.ExpressionWithAliasContext) {}

// EnterExpressions is called when production expressions is entered.
func (s *GparserParserListener) EnterExpressions(ctx *parser.ExpressionsContext) {}

// ExitExpressions is called when production expressions is exited.
func (s *GparserParserListener) ExitExpressions(ctx *parser.ExpressionsContext) {}

// EnterColumnRefOrderInParenthesis is called when production columnRefOrderInParenthesis is entered.
func (s *GparserParserListener) EnterColumnRefOrderInParenthesis(ctx *parser.ColumnRefOrderInParenthesisContext) {
}

// ExitColumnRefOrderInParenthesis is called when production columnRefOrderInParenthesis is exited.
func (s *GparserParserListener) ExitColumnRefOrderInParenthesis(ctx *parser.ColumnRefOrderInParenthesisContext) {
}

// EnterColumnRefOrderNotInParenthesis is called when production columnRefOrderNotInParenthesis is entered.
func (s *GparserParserListener) EnterColumnRefOrderNotInParenthesis(ctx *parser.ColumnRefOrderNotInParenthesisContext) {
}

// ExitColumnRefOrderNotInParenthesis is called when production columnRefOrderNotInParenthesis is exited.
func (s *GparserParserListener) ExitColumnRefOrderNotInParenthesis(ctx *parser.ColumnRefOrderNotInParenthesisContext) {
}

// EnterOrderByClause is called when production orderByClause is entered.
func (s *GparserParserListener) EnterOrderByClause(ctx *parser.OrderByClauseContext) {}

// ExitOrderByClause is called when production orderByClause is exited.
func (s *GparserParserListener) ExitOrderByClause(ctx *parser.OrderByClauseContext) {}

// EnterClusterByClause is called when production clusterByClause is entered.
func (s *GparserParserListener) EnterClusterByClause(ctx *parser.ClusterByClauseContext) {}

// ExitClusterByClause is called when production clusterByClause is exited.
func (s *GparserParserListener) ExitClusterByClause(ctx *parser.ClusterByClauseContext) {}

// EnterPartitionByClause is called when production partitionByClause is entered.
func (s *GparserParserListener) EnterPartitionByClause(ctx *parser.PartitionByClauseContext) {}

// ExitPartitionByClause is called when production partitionByClause is exited.
func (s *GparserParserListener) ExitPartitionByClause(ctx *parser.PartitionByClauseContext) {}

// EnterDistributeByClause is called when production distributeByClause is entered.
func (s *GparserParserListener) EnterDistributeByClause(ctx *parser.DistributeByClauseContext) {}

// ExitDistributeByClause is called when production distributeByClause is exited.
func (s *GparserParserListener) ExitDistributeByClause(ctx *parser.DistributeByClauseContext) {}

// EnterSortByClause is called when production sortByClause is entered.
func (s *GparserParserListener) EnterSortByClause(ctx *parser.SortByClauseContext) {}

// ExitSortByClause is called when production sortByClause is exited.
func (s *GparserParserListener) ExitSortByClause(ctx *parser.SortByClauseContext) {}

// EnterTrimFunction is called when production trimFunction is entered.
func (s *GparserParserListener) EnterTrimFunction(ctx *parser.TrimFunctionContext) {}

// ExitTrimFunction is called when production trimFunction is exited.
func (s *GparserParserListener) ExitTrimFunction(ctx *parser.TrimFunctionContext) {}

// EnterFunction_ is called when production function_ is entered.
func (s *GparserParserListener) EnterFunction_(ctx *parser.Function_Context) {}

// ExitFunction_ is called when production function_ is exited.
func (s *GparserParserListener) ExitFunction_(ctx *parser.Function_Context) {}

// EnterNull_treatment is called when production null_treatment is entered.
func (s *GparserParserListener) EnterNull_treatment(ctx *parser.Null_treatmentContext) {}

// ExitNull_treatment is called when production null_treatment is exited.
func (s *GparserParserListener) ExitNull_treatment(ctx *parser.Null_treatmentContext) {}

// EnterFunctionName is called when production functionName is entered.
func (s *GparserParserListener) EnterFunctionName(ctx *parser.FunctionNameContext) {}

// ExitFunctionName is called when production functionName is exited.
func (s *GparserParserListener) ExitFunctionName(ctx *parser.FunctionNameContext) {}

// EnterCastExpression is called when production castExpression is entered.
func (s *GparserParserListener) EnterCastExpression(ctx *parser.CastExpressionContext) {}

// ExitCastExpression is called when production castExpression is exited.
func (s *GparserParserListener) ExitCastExpression(ctx *parser.CastExpressionContext) {}

// EnterCaseExpression is called when production caseExpression is entered.
func (s *GparserParserListener) EnterCaseExpression(ctx *parser.CaseExpressionContext) {}

// ExitCaseExpression is called when production caseExpression is exited.
func (s *GparserParserListener) ExitCaseExpression(ctx *parser.CaseExpressionContext) {}

// EnterWhenExpression is called when production whenExpression is entered.
func (s *GparserParserListener) EnterWhenExpression(ctx *parser.WhenExpressionContext) {}

// ExitWhenExpression is called when production whenExpression is exited.
func (s *GparserParserListener) ExitWhenExpression(ctx *parser.WhenExpressionContext) {}

// EnterFloorExpression is called when production floorExpression is entered.
func (s *GparserParserListener) EnterFloorExpression(ctx *parser.FloorExpressionContext) {}

// ExitFloorExpression is called when production floorExpression is exited.
func (s *GparserParserListener) ExitFloorExpression(ctx *parser.FloorExpressionContext) {}

// EnterFloorDateQualifiers is called when production floorDateQualifiers is entered.
func (s *GparserParserListener) EnterFloorDateQualifiers(ctx *parser.FloorDateQualifiersContext) {}

// ExitFloorDateQualifiers is called when production floorDateQualifiers is exited.
func (s *GparserParserListener) ExitFloorDateQualifiers(ctx *parser.FloorDateQualifiersContext) {}

// EnterExtractExpression is called when production extractExpression is entered.
func (s *GparserParserListener) EnterExtractExpression(ctx *parser.ExtractExpressionContext) {}

// ExitExtractExpression is called when production extractExpression is exited.
func (s *GparserParserListener) ExitExtractExpression(ctx *parser.ExtractExpressionContext) {}

// EnterTimeQualifiers is called when production timeQualifiers is entered.
func (s *GparserParserListener) EnterTimeQualifiers(ctx *parser.TimeQualifiersContext) {}

// ExitTimeQualifiers is called when production timeQualifiers is exited.
func (s *GparserParserListener) ExitTimeQualifiers(ctx *parser.TimeQualifiersContext) {}

// EnterConstant is called when production constant is entered.
func (s *GparserParserListener) EnterConstant(ctx *parser.ConstantContext) {}

// ExitConstant is called when production constant is exited.
func (s *GparserParserListener) ExitConstant(ctx *parser.ConstantContext) {}

// EnterPrepareStmtParam is called when production prepareStmtParam is entered.
func (s *GparserParserListener) EnterPrepareStmtParam(ctx *parser.PrepareStmtParamContext) {}

// ExitPrepareStmtParam is called when production prepareStmtParam is exited.
func (s *GparserParserListener) ExitPrepareStmtParam(ctx *parser.PrepareStmtParamContext) {}

// EnterParameterIdx is called when production parameterIdx is entered.
func (s *GparserParserListener) EnterParameterIdx(ctx *parser.ParameterIdxContext) {}

// ExitParameterIdx is called when production parameterIdx is exited.
func (s *GparserParserListener) ExitParameterIdx(ctx *parser.ParameterIdxContext) {}

// EnterStringLiteralSequence is called when production stringLiteralSequence is entered.
func (s *GparserParserListener) EnterStringLiteralSequence(ctx *parser.StringLiteralSequenceContext) {
}

// ExitStringLiteralSequence is called when production stringLiteralSequence is exited.
func (s *GparserParserListener) ExitStringLiteralSequence(ctx *parser.StringLiteralSequenceContext) {}

// EnterCharSetStringLiteral is called when production charSetStringLiteral is entered.
func (s *GparserParserListener) EnterCharSetStringLiteral(ctx *parser.CharSetStringLiteralContext) {}

// ExitCharSetStringLiteral is called when production charSetStringLiteral is exited.
func (s *GparserParserListener) ExitCharSetStringLiteral(ctx *parser.CharSetStringLiteralContext) {}

// EnterDateLiteral is called when production dateLiteral is entered.
func (s *GparserParserListener) EnterDateLiteral(ctx *parser.DateLiteralContext) {}

// ExitDateLiteral is called when production dateLiteral is exited.
func (s *GparserParserListener) ExitDateLiteral(ctx *parser.DateLiteralContext) {}

// EnterTimestampLiteral is called when production timestampLiteral is entered.
func (s *GparserParserListener) EnterTimestampLiteral(ctx *parser.TimestampLiteralContext) {}

// ExitTimestampLiteral is called when production timestampLiteral is exited.
func (s *GparserParserListener) ExitTimestampLiteral(ctx *parser.TimestampLiteralContext) {}

// EnterTimestampLocalTZLiteral is called when production timestampLocalTZLiteral is entered.
func (s *GparserParserListener) EnterTimestampLocalTZLiteral(ctx *parser.TimestampLocalTZLiteralContext) {
}

// ExitTimestampLocalTZLiteral is called when production timestampLocalTZLiteral is exited.
func (s *GparserParserListener) ExitTimestampLocalTZLiteral(ctx *parser.TimestampLocalTZLiteralContext) {
}

// EnterIntervalValue is called when production intervalValue is entered.
func (s *GparserParserListener) EnterIntervalValue(ctx *parser.IntervalValueContext) {}

// ExitIntervalValue is called when production intervalValue is exited.
func (s *GparserParserListener) ExitIntervalValue(ctx *parser.IntervalValueContext) {}

// EnterIntervalLiteral is called when production intervalLiteral is entered.
func (s *GparserParserListener) EnterIntervalLiteral(ctx *parser.IntervalLiteralContext) {}

// ExitIntervalLiteral is called when production intervalLiteral is exited.
func (s *GparserParserListener) ExitIntervalLiteral(ctx *parser.IntervalLiteralContext) {}

// EnterIntervalExpression is called when production intervalExpression is entered.
func (s *GparserParserListener) EnterIntervalExpression(ctx *parser.IntervalExpressionContext) {}

// ExitIntervalExpression is called when production intervalExpression is exited.
func (s *GparserParserListener) ExitIntervalExpression(ctx *parser.IntervalExpressionContext) {}

// EnterIntervalQualifiers is called when production intervalQualifiers is entered.
func (s *GparserParserListener) EnterIntervalQualifiers(ctx *parser.IntervalQualifiersContext) {}

// ExitIntervalQualifiers is called when production intervalQualifiers is exited.
func (s *GparserParserListener) ExitIntervalQualifiers(ctx *parser.IntervalQualifiersContext) {}

// EnterExpression is called when production expression is entered.
func (s *GparserParserListener) EnterExpression(ctx *parser.ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *GparserParserListener) ExitExpression(ctx *parser.ExpressionContext) {}

// EnterAtomExpression is called when production atomExpression is entered.
func (s *GparserParserListener) EnterAtomExpression(ctx *parser.AtomExpressionContext) {}

// ExitAtomExpression is called when production atomExpression is exited.
func (s *GparserParserListener) ExitAtomExpression(ctx *parser.AtomExpressionContext) {}

// EnterPrecedenceFieldExpression is called when production precedenceFieldExpression is entered.
func (s *GparserParserListener) EnterPrecedenceFieldExpression(ctx *parser.PrecedenceFieldExpressionContext) {
}

// ExitPrecedenceFieldExpression is called when production precedenceFieldExpression is exited.
func (s *GparserParserListener) ExitPrecedenceFieldExpression(ctx *parser.PrecedenceFieldExpressionContext) {
}

// EnterPrecedenceUnaryOperator is called when production precedenceUnaryOperator is entered.
func (s *GparserParserListener) EnterPrecedenceUnaryOperator(ctx *parser.PrecedenceUnaryOperatorContext) {
}

// ExitPrecedenceUnaryOperator is called when production precedenceUnaryOperator is exited.
func (s *GparserParserListener) ExitPrecedenceUnaryOperator(ctx *parser.PrecedenceUnaryOperatorContext) {
}

// EnterPrecedenceUnaryPrefixExpression is called when production precedenceUnaryPrefixExpression is entered.
func (s *GparserParserListener) EnterPrecedenceUnaryPrefixExpression(ctx *parser.PrecedenceUnaryPrefixExpressionContext) {
}

// ExitPrecedenceUnaryPrefixExpression is called when production precedenceUnaryPrefixExpression is exited.
func (s *GparserParserListener) ExitPrecedenceUnaryPrefixExpression(ctx *parser.PrecedenceUnaryPrefixExpressionContext) {
}

// EnterPrecedenceBitwiseXorOperator is called when production precedenceBitwiseXorOperator is entered.
func (s *GparserParserListener) EnterPrecedenceBitwiseXorOperator(ctx *parser.PrecedenceBitwiseXorOperatorContext) {
}

// ExitPrecedenceBitwiseXorOperator is called when production precedenceBitwiseXorOperator is exited.
func (s *GparserParserListener) ExitPrecedenceBitwiseXorOperator(ctx *parser.PrecedenceBitwiseXorOperatorContext) {
}

// EnterPrecedenceBitwiseXorExpression is called when production precedenceBitwiseXorExpression is entered.
func (s *GparserParserListener) EnterPrecedenceBitwiseXorExpression(ctx *parser.PrecedenceBitwiseXorExpressionContext) {
}

// ExitPrecedenceBitwiseXorExpression is called when production precedenceBitwiseXorExpression is exited.
func (s *GparserParserListener) ExitPrecedenceBitwiseXorExpression(ctx *parser.PrecedenceBitwiseXorExpressionContext) {
}

// EnterPrecedenceStarOperator is called when production precedenceStarOperator is entered.
func (s *GparserParserListener) EnterPrecedenceStarOperator(ctx *parser.PrecedenceStarOperatorContext) {
}

// ExitPrecedenceStarOperator is called when production precedenceStarOperator is exited.
func (s *GparserParserListener) ExitPrecedenceStarOperator(ctx *parser.PrecedenceStarOperatorContext) {
}

// EnterPrecedenceStarExpression is called when production precedenceStarExpression is entered.
func (s *GparserParserListener) EnterPrecedenceStarExpression(ctx *parser.PrecedenceStarExpressionContext) {
}

// ExitPrecedenceStarExpression is called when production precedenceStarExpression is exited.
func (s *GparserParserListener) ExitPrecedenceStarExpression(ctx *parser.PrecedenceStarExpressionContext) {
}

// EnterPrecedencePlusOperator is called when production precedencePlusOperator is entered.
func (s *GparserParserListener) EnterPrecedencePlusOperator(ctx *parser.PrecedencePlusOperatorContext) {
}

// ExitPrecedencePlusOperator is called when production precedencePlusOperator is exited.
func (s *GparserParserListener) ExitPrecedencePlusOperator(ctx *parser.PrecedencePlusOperatorContext) {
}

// EnterPrecedencePlusExpression is called when production precedencePlusExpression is entered.
func (s *GparserParserListener) EnterPrecedencePlusExpression(ctx *parser.PrecedencePlusExpressionContext) {
}

// ExitPrecedencePlusExpression is called when production precedencePlusExpression is exited.
func (s *GparserParserListener) ExitPrecedencePlusExpression(ctx *parser.PrecedencePlusExpressionContext) {
}

// EnterPrecedenceConcatenateOperator is called when production precedenceConcatenateOperator is entered.
func (s *GparserParserListener) EnterPrecedenceConcatenateOperator(ctx *parser.PrecedenceConcatenateOperatorContext) {
}

// ExitPrecedenceConcatenateOperator is called when production precedenceConcatenateOperator is exited.
func (s *GparserParserListener) ExitPrecedenceConcatenateOperator(ctx *parser.PrecedenceConcatenateOperatorContext) {
}

// EnterPrecedenceConcatenateExpression is called when production precedenceConcatenateExpression is entered.
func (s *GparserParserListener) EnterPrecedenceConcatenateExpression(ctx *parser.PrecedenceConcatenateExpressionContext) {
}

// ExitPrecedenceConcatenateExpression is called when production precedenceConcatenateExpression is exited.
func (s *GparserParserListener) ExitPrecedenceConcatenateExpression(ctx *parser.PrecedenceConcatenateExpressionContext) {
}

// EnterPrecedenceAmpersandOperator is called when production precedenceAmpersandOperator is entered.
func (s *GparserParserListener) EnterPrecedenceAmpersandOperator(ctx *parser.PrecedenceAmpersandOperatorContext) {
}

// ExitPrecedenceAmpersandOperator is called when production precedenceAmpersandOperator is exited.
func (s *GparserParserListener) ExitPrecedenceAmpersandOperator(ctx *parser.PrecedenceAmpersandOperatorContext) {
}

// EnterPrecedenceAmpersandExpression is called when production precedenceAmpersandExpression is entered.
func (s *GparserParserListener) EnterPrecedenceAmpersandExpression(ctx *parser.PrecedenceAmpersandExpressionContext) {
}

// ExitPrecedenceAmpersandExpression is called when production precedenceAmpersandExpression is exited.
func (s *GparserParserListener) ExitPrecedenceAmpersandExpression(ctx *parser.PrecedenceAmpersandExpressionContext) {
}

// EnterPrecedenceBitwiseOrOperator is called when production precedenceBitwiseOrOperator is entered.
func (s *GparserParserListener) EnterPrecedenceBitwiseOrOperator(ctx *parser.PrecedenceBitwiseOrOperatorContext) {
}

// ExitPrecedenceBitwiseOrOperator is called when production precedenceBitwiseOrOperator is exited.
func (s *GparserParserListener) ExitPrecedenceBitwiseOrOperator(ctx *parser.PrecedenceBitwiseOrOperatorContext) {
}

// EnterPrecedenceBitwiseOrExpression is called when production precedenceBitwiseOrExpression is entered.
func (s *GparserParserListener) EnterPrecedenceBitwiseOrExpression(ctx *parser.PrecedenceBitwiseOrExpressionContext) {
}

// ExitPrecedenceBitwiseOrExpression is called when production precedenceBitwiseOrExpression is exited.
func (s *GparserParserListener) ExitPrecedenceBitwiseOrExpression(ctx *parser.PrecedenceBitwiseOrExpressionContext) {
}

// EnterPrecedenceRegexpOperator is called when production precedenceRegexpOperator is entered.
func (s *GparserParserListener) EnterPrecedenceRegexpOperator(ctx *parser.PrecedenceRegexpOperatorContext) {
}

// ExitPrecedenceRegexpOperator is called when production precedenceRegexpOperator is exited.
func (s *GparserParserListener) ExitPrecedenceRegexpOperator(ctx *parser.PrecedenceRegexpOperatorContext) {
}

// EnterPrecedenceSimilarOperator is called when production precedenceSimilarOperator is entered.
func (s *GparserParserListener) EnterPrecedenceSimilarOperator(ctx *parser.PrecedenceSimilarOperatorContext) {
}

// ExitPrecedenceSimilarOperator is called when production precedenceSimilarOperator is exited.
func (s *GparserParserListener) ExitPrecedenceSimilarOperator(ctx *parser.PrecedenceSimilarOperatorContext) {
}

// EnterSubQueryExpression is called when production subQueryExpression is entered.
func (s *GparserParserListener) EnterSubQueryExpression(ctx *parser.SubQueryExpressionContext) {}

// ExitSubQueryExpression is called when production subQueryExpression is exited.
func (s *GparserParserListener) ExitSubQueryExpression(ctx *parser.SubQueryExpressionContext) {}

// EnterPrecedenceSimilarExpression is called when production precedenceSimilarExpression is entered.
func (s *GparserParserListener) EnterPrecedenceSimilarExpression(ctx *parser.PrecedenceSimilarExpressionContext) {
}

// ExitPrecedenceSimilarExpression is called when production precedenceSimilarExpression is exited.
func (s *GparserParserListener) ExitPrecedenceSimilarExpression(ctx *parser.PrecedenceSimilarExpressionContext) {
}

// EnterPrecedenceSimilarExpressionMain is called when production precedenceSimilarExpressionMain is entered.
func (s *GparserParserListener) EnterPrecedenceSimilarExpressionMain(ctx *parser.PrecedenceSimilarExpressionMainContext) {
}

// ExitPrecedenceSimilarExpressionMain is called when production precedenceSimilarExpressionMain is exited.
func (s *GparserParserListener) ExitPrecedenceSimilarExpressionMain(ctx *parser.PrecedenceSimilarExpressionMainContext) {
}

// EnterPrecedenceSimilarExpressionPart is called when production precedenceSimilarExpressionPart is entered.
func (s *GparserParserListener) EnterPrecedenceSimilarExpressionPart(ctx *parser.PrecedenceSimilarExpressionPartContext) {
}

// ExitPrecedenceSimilarExpressionPart is called when production precedenceSimilarExpressionPart is exited.
func (s *GparserParserListener) ExitPrecedenceSimilarExpressionPart(ctx *parser.PrecedenceSimilarExpressionPartContext) {
}

// EnterPrecedenceSimilarExpressionAtom is called when production precedenceSimilarExpressionAtom is entered.
func (s *GparserParserListener) EnterPrecedenceSimilarExpressionAtom(ctx *parser.PrecedenceSimilarExpressionAtomContext) {
}

// ExitPrecedenceSimilarExpressionAtom is called when production precedenceSimilarExpressionAtom is exited.
func (s *GparserParserListener) ExitPrecedenceSimilarExpressionAtom(ctx *parser.PrecedenceSimilarExpressionAtomContext) {
}

// EnterPrecedenceSimilarExpressionQuantifierPredicate is called when production precedenceSimilarExpressionQuantifierPredicate is entered.
func (s *GparserParserListener) EnterPrecedenceSimilarExpressionQuantifierPredicate(ctx *parser.PrecedenceSimilarExpressionQuantifierPredicateContext) {
}

// ExitPrecedenceSimilarExpressionQuantifierPredicate is called when production precedenceSimilarExpressionQuantifierPredicate is exited.
func (s *GparserParserListener) ExitPrecedenceSimilarExpressionQuantifierPredicate(ctx *parser.PrecedenceSimilarExpressionQuantifierPredicateContext) {
}

// EnterQuantifierType is called when production quantifierType is entered.
func (s *GparserParserListener) EnterQuantifierType(ctx *parser.QuantifierTypeContext) {}

// ExitQuantifierType is called when production quantifierType is exited.
func (s *GparserParserListener) ExitQuantifierType(ctx *parser.QuantifierTypeContext) {}

// EnterPrecedenceSimilarExpressionIn is called when production precedenceSimilarExpressionIn is entered.
func (s *GparserParserListener) EnterPrecedenceSimilarExpressionIn(ctx *parser.PrecedenceSimilarExpressionInContext) {
}

// ExitPrecedenceSimilarExpressionIn is called when production precedenceSimilarExpressionIn is exited.
func (s *GparserParserListener) ExitPrecedenceSimilarExpressionIn(ctx *parser.PrecedenceSimilarExpressionInContext) {
}

// EnterPrecedenceSimilarExpressionPartNot is called when production precedenceSimilarExpressionPartNot is entered.
func (s *GparserParserListener) EnterPrecedenceSimilarExpressionPartNot(ctx *parser.PrecedenceSimilarExpressionPartNotContext) {
}

// ExitPrecedenceSimilarExpressionPartNot is called when production precedenceSimilarExpressionPartNot is exited.
func (s *GparserParserListener) ExitPrecedenceSimilarExpressionPartNot(ctx *parser.PrecedenceSimilarExpressionPartNotContext) {
}

// EnterPrecedenceDistinctOperator is called when production precedenceDistinctOperator is entered.
func (s *GparserParserListener) EnterPrecedenceDistinctOperator(ctx *parser.PrecedenceDistinctOperatorContext) {
}

// ExitPrecedenceDistinctOperator is called when production precedenceDistinctOperator is exited.
func (s *GparserParserListener) ExitPrecedenceDistinctOperator(ctx *parser.PrecedenceDistinctOperatorContext) {
}

// EnterPrecedenceEqualOperator is called when production precedenceEqualOperator is entered.
func (s *GparserParserListener) EnterPrecedenceEqualOperator(ctx *parser.PrecedenceEqualOperatorContext) {
}

// ExitPrecedenceEqualOperator is called when production precedenceEqualOperator is exited.
func (s *GparserParserListener) ExitPrecedenceEqualOperator(ctx *parser.PrecedenceEqualOperatorContext) {
}

// EnterPrecedenceEqualExpression is called when production precedenceEqualExpression is entered.
func (s *GparserParserListener) EnterPrecedenceEqualExpression(ctx *parser.PrecedenceEqualExpressionContext) {
}

// ExitPrecedenceEqualExpression is called when production precedenceEqualExpression is exited.
func (s *GparserParserListener) ExitPrecedenceEqualExpression(ctx *parser.PrecedenceEqualExpressionContext) {
}

// EnterIsCondition is called when production isCondition is entered.
func (s *GparserParserListener) EnterIsCondition(ctx *parser.IsConditionContext) {}

// ExitIsCondition is called when production isCondition is exited.
func (s *GparserParserListener) ExitIsCondition(ctx *parser.IsConditionContext) {}

// EnterPrecedenceUnarySuffixExpression is called when production precedenceUnarySuffixExpression is entered.
func (s *GparserParserListener) EnterPrecedenceUnarySuffixExpression(ctx *parser.PrecedenceUnarySuffixExpressionContext) {
}

// ExitPrecedenceUnarySuffixExpression is called when production precedenceUnarySuffixExpression is exited.
func (s *GparserParserListener) ExitPrecedenceUnarySuffixExpression(ctx *parser.PrecedenceUnarySuffixExpressionContext) {
}

// EnterPrecedenceNotOperator is called when production precedenceNotOperator is entered.
func (s *GparserParserListener) EnterPrecedenceNotOperator(ctx *parser.PrecedenceNotOperatorContext) {
}

// ExitPrecedenceNotOperator is called when production precedenceNotOperator is exited.
func (s *GparserParserListener) ExitPrecedenceNotOperator(ctx *parser.PrecedenceNotOperatorContext) {}

// EnterPrecedenceNotExpression is called when production precedenceNotExpression is entered.
func (s *GparserParserListener) EnterPrecedenceNotExpression(ctx *parser.PrecedenceNotExpressionContext) {
}

// ExitPrecedenceNotExpression is called when production precedenceNotExpression is exited.
func (s *GparserParserListener) ExitPrecedenceNotExpression(ctx *parser.PrecedenceNotExpressionContext) {
}

// EnterPrecedenceAndOperator is called when production precedenceAndOperator is entered.
func (s *GparserParserListener) EnterPrecedenceAndOperator(ctx *parser.PrecedenceAndOperatorContext) {
}

// ExitPrecedenceAndOperator is called when production precedenceAndOperator is exited.
func (s *GparserParserListener) ExitPrecedenceAndOperator(ctx *parser.PrecedenceAndOperatorContext) {}

// EnterPrecedenceAndExpression is called when production precedenceAndExpression is entered.
func (s *GparserParserListener) EnterPrecedenceAndExpression(ctx *parser.PrecedenceAndExpressionContext) {
}

// ExitPrecedenceAndExpression is called when production precedenceAndExpression is exited.
func (s *GparserParserListener) ExitPrecedenceAndExpression(ctx *parser.PrecedenceAndExpressionContext) {
}

// EnterPrecedenceOrOperator is called when production precedenceOrOperator is entered.
func (s *GparserParserListener) EnterPrecedenceOrOperator(ctx *parser.PrecedenceOrOperatorContext) {}

// ExitPrecedenceOrOperator is called when production precedenceOrOperator is exited.
func (s *GparserParserListener) ExitPrecedenceOrOperator(ctx *parser.PrecedenceOrOperatorContext) {}

// EnterPrecedenceOrExpression is called when production precedenceOrExpression is entered.
func (s *GparserParserListener) EnterPrecedenceOrExpression(ctx *parser.PrecedenceOrExpressionContext) {
}

// ExitPrecedenceOrExpression is called when production precedenceOrExpression is exited.
func (s *GparserParserListener) ExitPrecedenceOrExpression(ctx *parser.PrecedenceOrExpressionContext) {
}

// EnterBooleanValue is called when production booleanValue is entered.
func (s *GparserParserListener) EnterBooleanValue(ctx *parser.BooleanValueContext) {}

// ExitBooleanValue is called when production booleanValue is exited.
func (s *GparserParserListener) ExitBooleanValue(ctx *parser.BooleanValueContext) {}

// EnterBooleanValueTok is called when production booleanValueTok is entered.
func (s *GparserParserListener) EnterBooleanValueTok(ctx *parser.BooleanValueTokContext) {}

// ExitBooleanValueTok is called when production booleanValueTok is exited.
func (s *GparserParserListener) ExitBooleanValueTok(ctx *parser.BooleanValueTokContext) {}

// EnterTableOrPartition is called when production tableOrPartition is entered.
func (s *GparserParserListener) EnterTableOrPartition(ctx *parser.TableOrPartitionContext) {}

// ExitTableOrPartition is called when production tableOrPartition is exited.
func (s *GparserParserListener) ExitTableOrPartition(ctx *parser.TableOrPartitionContext) {}

// EnterPartitionSpec is called when production partitionSpec is entered.
func (s *GparserParserListener) EnterPartitionSpec(ctx *parser.PartitionSpecContext) {}

// ExitPartitionSpec is called when production partitionSpec is exited.
func (s *GparserParserListener) ExitPartitionSpec(ctx *parser.PartitionSpecContext) {}

// EnterPartitionVal is called when production partitionVal is entered.
func (s *GparserParserListener) EnterPartitionVal(ctx *parser.PartitionValContext) {}

// ExitPartitionVal is called when production partitionVal is exited.
func (s *GparserParserListener) ExitPartitionVal(ctx *parser.PartitionValContext) {}

// EnterPartitionSelectorSpec is called when production partitionSelectorSpec is entered.
func (s *GparserParserListener) EnterPartitionSelectorSpec(ctx *parser.PartitionSelectorSpecContext) {
}

// ExitPartitionSelectorSpec is called when production partitionSelectorSpec is exited.
func (s *GparserParserListener) ExitPartitionSelectorSpec(ctx *parser.PartitionSelectorSpecContext) {}

// EnterPartitionSelectorVal is called when production partitionSelectorVal is entered.
func (s *GparserParserListener) EnterPartitionSelectorVal(ctx *parser.PartitionSelectorValContext) {}

// ExitPartitionSelectorVal is called when production partitionSelectorVal is exited.
func (s *GparserParserListener) ExitPartitionSelectorVal(ctx *parser.PartitionSelectorValContext) {}

// EnterPartitionSelectorOperator is called when production partitionSelectorOperator is entered.
func (s *GparserParserListener) EnterPartitionSelectorOperator(ctx *parser.PartitionSelectorOperatorContext) {
}

// ExitPartitionSelectorOperator is called when production partitionSelectorOperator is exited.
func (s *GparserParserListener) ExitPartitionSelectorOperator(ctx *parser.PartitionSelectorOperatorContext) {
}

// EnterSubQuerySelectorOperator is called when production subQuerySelectorOperator is entered.
func (s *GparserParserListener) EnterSubQuerySelectorOperator(ctx *parser.SubQuerySelectorOperatorContext) {
}

// ExitSubQuerySelectorOperator is called when production subQuerySelectorOperator is exited.
func (s *GparserParserListener) ExitSubQuerySelectorOperator(ctx *parser.SubQuerySelectorOperatorContext) {
}

// EnterSysFuncNames is called when production sysFuncNames is entered.
func (s *GparserParserListener) EnterSysFuncNames(ctx *parser.SysFuncNamesContext) {}

// ExitSysFuncNames is called when production sysFuncNames is exited.
func (s *GparserParserListener) ExitSysFuncNames(ctx *parser.SysFuncNamesContext) {}

// EnterDescFuncNames is called when production descFuncNames is entered.
func (s *GparserParserListener) EnterDescFuncNames(ctx *parser.DescFuncNamesContext) {}

// ExitDescFuncNames is called when production descFuncNames is exited.
func (s *GparserParserListener) ExitDescFuncNames(ctx *parser.DescFuncNamesContext) {}

// EnterId_ is called when production id_ is entered.
func (s *GparserParserListener) EnterId_(ctx *parser.Id_Context) {}

// ExitId_ is called when production id_ is exited.
func (s *GparserParserListener) ExitId_(ctx *parser.Id_Context) {}

// EnterFunctionIdentifier is called when production functionIdentifier is entered.
func (s *GparserParserListener) EnterFunctionIdentifier(ctx *parser.FunctionIdentifierContext) {}

// ExitFunctionIdentifier is called when production functionIdentifier is exited.
func (s *GparserParserListener) ExitFunctionIdentifier(ctx *parser.FunctionIdentifierContext) {}

// EnterPrincipalIdentifier is called when production principalIdentifier is entered.
func (s *GparserParserListener) EnterPrincipalIdentifier(ctx *parser.PrincipalIdentifierContext) {}

// ExitPrincipalIdentifier is called when production principalIdentifier is exited.
func (s *GparserParserListener) ExitPrincipalIdentifier(ctx *parser.PrincipalIdentifierContext) {}

// EnterNonReserved is called when production nonReserved is entered.
func (s *GparserParserListener) EnterNonReserved(ctx *parser.NonReservedContext) {}

// ExitNonReserved is called when production nonReserved is exited.
func (s *GparserParserListener) ExitNonReserved(ctx *parser.NonReservedContext) {}

// EnterSql11ReservedKeywordsUsedAsFunctionName is called when production sql11ReservedKeywordsUsedAsFunctionName is entered.
func (s *GparserParserListener) EnterSql11ReservedKeywordsUsedAsFunctionName(ctx *parser.Sql11ReservedKeywordsUsedAsFunctionNameContext) {
}

// ExitSql11ReservedKeywordsUsedAsFunctionName is called when production sql11ReservedKeywordsUsedAsFunctionName is exited.
func (s *GparserParserListener) ExitSql11ReservedKeywordsUsedAsFunctionName(ctx *parser.Sql11ReservedKeywordsUsedAsFunctionNameContext) {
}

// EnterHint is called when production hint is entered.
func (s *GparserParserListener) EnterHint(ctx *parser.HintContext) {}

// ExitHint is called when production hint is exited.
func (s *GparserParserListener) ExitHint(ctx *parser.HintContext) {}

// EnterHintList is called when production hintList is entered.
func (s *GparserParserListener) EnterHintList(ctx *parser.HintListContext) {}

// ExitHintList is called when production hintList is exited.
func (s *GparserParserListener) ExitHintList(ctx *parser.HintListContext) {}

// EnterHintItem is called when production hintItem is entered.
func (s *GparserParserListener) EnterHintItem(ctx *parser.HintItemContext) {}

// ExitHintItem is called when production hintItem is exited.
func (s *GparserParserListener) ExitHintItem(ctx *parser.HintItemContext) {}

// EnterHintName is called when production hintName is entered.
func (s *GparserParserListener) EnterHintName(ctx *parser.HintNameContext) {}

// ExitHintName is called when production hintName is exited.
func (s *GparserParserListener) ExitHintName(ctx *parser.HintNameContext) {}

// EnterHintArgs is called when production hintArgs is entered.
func (s *GparserParserListener) EnterHintArgs(ctx *parser.HintArgsContext) {}

// ExitHintArgs is called when production hintArgs is exited.
func (s *GparserParserListener) ExitHintArgs(ctx *parser.HintArgsContext) {}

// EnterHintArgName is called when production hintArgName is entered.
func (s *GparserParserListener) EnterHintArgName(ctx *parser.HintArgNameContext) {}

// ExitHintArgName is called when production hintArgName is exited.
func (s *GparserParserListener) ExitHintArgName(ctx *parser.HintArgNameContext) {}

// EnterPrepareStatement is called when production prepareStatement is entered.
func (s *GparserParserListener) EnterPrepareStatement(ctx *parser.PrepareStatementContext) {}

// ExitPrepareStatement is called when production prepareStatement is exited.
func (s *GparserParserListener) ExitPrepareStatement(ctx *parser.PrepareStatementContext) {}

// EnterExecuteStatement is called when production executeStatement is entered.
func (s *GparserParserListener) EnterExecuteStatement(ctx *parser.ExecuteStatementContext) {}

// ExitExecuteStatement is called when production executeStatement is exited.
func (s *GparserParserListener) ExitExecuteStatement(ctx *parser.ExecuteStatementContext) {}

// EnterExecuteParamList is called when production executeParamList is entered.
func (s *GparserParserListener) EnterExecuteParamList(ctx *parser.ExecuteParamListContext) {}

// ExitExecuteParamList is called when production executeParamList is exited.
func (s *GparserParserListener) ExitExecuteParamList(ctx *parser.ExecuteParamListContext) {}

// EnterResourcePlanDdlStatements is called when production resourcePlanDdlStatements is entered.
func (s *GparserParserListener) EnterResourcePlanDdlStatements(ctx *parser.ResourcePlanDdlStatementsContext) {
}

// ExitResourcePlanDdlStatements is called when production resourcePlanDdlStatements is exited.
func (s *GparserParserListener) ExitResourcePlanDdlStatements(ctx *parser.ResourcePlanDdlStatementsContext) {
}

// EnterRpAssign is called when production rpAssign is entered.
func (s *GparserParserListener) EnterRpAssign(ctx *parser.RpAssignContext) {}

// ExitRpAssign is called when production rpAssign is exited.
func (s *GparserParserListener) ExitRpAssign(ctx *parser.RpAssignContext) {}

// EnterRpAssignList is called when production rpAssignList is entered.
func (s *GparserParserListener) EnterRpAssignList(ctx *parser.RpAssignListContext) {}

// ExitRpAssignList is called when production rpAssignList is exited.
func (s *GparserParserListener) ExitRpAssignList(ctx *parser.RpAssignListContext) {}

// EnterRpUnassign is called when production rpUnassign is entered.
func (s *GparserParserListener) EnterRpUnassign(ctx *parser.RpUnassignContext) {}

// ExitRpUnassign is called when production rpUnassign is exited.
func (s *GparserParserListener) ExitRpUnassign(ctx *parser.RpUnassignContext) {}

// EnterRpUnassignList is called when production rpUnassignList is entered.
func (s *GparserParserListener) EnterRpUnassignList(ctx *parser.RpUnassignListContext) {}

// ExitRpUnassignList is called when production rpUnassignList is exited.
func (s *GparserParserListener) ExitRpUnassignList(ctx *parser.RpUnassignListContext) {}

// EnterCreateResourcePlanStatement is called when production createResourcePlanStatement is entered.
func (s *GparserParserListener) EnterCreateResourcePlanStatement(ctx *parser.CreateResourcePlanStatementContext) {
}

// ExitCreateResourcePlanStatement is called when production createResourcePlanStatement is exited.
func (s *GparserParserListener) ExitCreateResourcePlanStatement(ctx *parser.CreateResourcePlanStatementContext) {
}

// EnterWithReplace is called when production withReplace is entered.
func (s *GparserParserListener) EnterWithReplace(ctx *parser.WithReplaceContext) {}

// ExitWithReplace is called when production withReplace is exited.
func (s *GparserParserListener) ExitWithReplace(ctx *parser.WithReplaceContext) {}

// EnterActivate is called when production activate is entered.
func (s *GparserParserListener) EnterActivate(ctx *parser.ActivateContext) {}

// ExitActivate is called when production activate is exited.
func (s *GparserParserListener) ExitActivate(ctx *parser.ActivateContext) {}

// EnterEnable is called when production enable is entered.
func (s *GparserParserListener) EnterEnable(ctx *parser.EnableContext) {}

// ExitEnable is called when production enable is exited.
func (s *GparserParserListener) ExitEnable(ctx *parser.EnableContext) {}

// EnterDisable is called when production disable is entered.
func (s *GparserParserListener) EnterDisable(ctx *parser.DisableContext) {}

// ExitDisable is called when production disable is exited.
func (s *GparserParserListener) ExitDisable(ctx *parser.DisableContext) {}

// EnterUnmanaged is called when production unmanaged is entered.
func (s *GparserParserListener) EnterUnmanaged(ctx *parser.UnmanagedContext) {}

// ExitUnmanaged is called when production unmanaged is exited.
func (s *GparserParserListener) ExitUnmanaged(ctx *parser.UnmanagedContext) {}

// EnterAlterResourcePlanStatement is called when production alterResourcePlanStatement is entered.
func (s *GparserParserListener) EnterAlterResourcePlanStatement(ctx *parser.AlterResourcePlanStatementContext) {
}

// ExitAlterResourcePlanStatement is called when production alterResourcePlanStatement is exited.
func (s *GparserParserListener) ExitAlterResourcePlanStatement(ctx *parser.AlterResourcePlanStatementContext) {
}

// EnterGlobalWmStatement is called when production globalWmStatement is entered.
func (s *GparserParserListener) EnterGlobalWmStatement(ctx *parser.GlobalWmStatementContext) {}

// ExitGlobalWmStatement is called when production globalWmStatement is exited.
func (s *GparserParserListener) ExitGlobalWmStatement(ctx *parser.GlobalWmStatementContext) {}

// EnterReplaceResourcePlanStatement is called when production replaceResourcePlanStatement is entered.
func (s *GparserParserListener) EnterReplaceResourcePlanStatement(ctx *parser.ReplaceResourcePlanStatementContext) {
}

// ExitReplaceResourcePlanStatement is called when production replaceResourcePlanStatement is exited.
func (s *GparserParserListener) ExitReplaceResourcePlanStatement(ctx *parser.ReplaceResourcePlanStatementContext) {
}

// EnterDropResourcePlanStatement is called when production dropResourcePlanStatement is entered.
func (s *GparserParserListener) EnterDropResourcePlanStatement(ctx *parser.DropResourcePlanStatementContext) {
}

// ExitDropResourcePlanStatement is called when production dropResourcePlanStatement is exited.
func (s *GparserParserListener) ExitDropResourcePlanStatement(ctx *parser.DropResourcePlanStatementContext) {
}

// EnterPoolPath is called when production poolPath is entered.
func (s *GparserParserListener) EnterPoolPath(ctx *parser.PoolPathContext) {}

// ExitPoolPath is called when production poolPath is exited.
func (s *GparserParserListener) ExitPoolPath(ctx *parser.PoolPathContext) {}

// EnterTriggerExpression is called when production triggerExpression is entered.
func (s *GparserParserListener) EnterTriggerExpression(ctx *parser.TriggerExpressionContext) {}

// ExitTriggerExpression is called when production triggerExpression is exited.
func (s *GparserParserListener) ExitTriggerExpression(ctx *parser.TriggerExpressionContext) {}

// EnterTriggerExpressionStandalone is called when production triggerExpressionStandalone is entered.
func (s *GparserParserListener) EnterTriggerExpressionStandalone(ctx *parser.TriggerExpressionStandaloneContext) {
}

// ExitTriggerExpressionStandalone is called when production triggerExpressionStandalone is exited.
func (s *GparserParserListener) ExitTriggerExpressionStandalone(ctx *parser.TriggerExpressionStandaloneContext) {
}

// EnterTriggerOrExpression is called when production triggerOrExpression is entered.
func (s *GparserParserListener) EnterTriggerOrExpression(ctx *parser.TriggerOrExpressionContext) {}

// ExitTriggerOrExpression is called when production triggerOrExpression is exited.
func (s *GparserParserListener) ExitTriggerOrExpression(ctx *parser.TriggerOrExpressionContext) {}

// EnterTriggerAndExpression is called when production triggerAndExpression is entered.
func (s *GparserParserListener) EnterTriggerAndExpression(ctx *parser.TriggerAndExpressionContext) {}

// ExitTriggerAndExpression is called when production triggerAndExpression is exited.
func (s *GparserParserListener) ExitTriggerAndExpression(ctx *parser.TriggerAndExpressionContext) {}

// EnterTriggerAtomExpression is called when production triggerAtomExpression is entered.
func (s *GparserParserListener) EnterTriggerAtomExpression(ctx *parser.TriggerAtomExpressionContext) {
}

// ExitTriggerAtomExpression is called when production triggerAtomExpression is exited.
func (s *GparserParserListener) ExitTriggerAtomExpression(ctx *parser.TriggerAtomExpressionContext) {}

// EnterTriggerLiteral is called when production triggerLiteral is entered.
func (s *GparserParserListener) EnterTriggerLiteral(ctx *parser.TriggerLiteralContext) {}

// ExitTriggerLiteral is called when production triggerLiteral is exited.
func (s *GparserParserListener) ExitTriggerLiteral(ctx *parser.TriggerLiteralContext) {}

// EnterComparisionOperator is called when production comparisionOperator is entered.
func (s *GparserParserListener) EnterComparisionOperator(ctx *parser.ComparisionOperatorContext) {}

// ExitComparisionOperator is called when production comparisionOperator is exited.
func (s *GparserParserListener) ExitComparisionOperator(ctx *parser.ComparisionOperatorContext) {}

// EnterTriggerActionExpression is called when production triggerActionExpression is entered.
func (s *GparserParserListener) EnterTriggerActionExpression(ctx *parser.TriggerActionExpressionContext) {
}

// ExitTriggerActionExpression is called when production triggerActionExpression is exited.
func (s *GparserParserListener) ExitTriggerActionExpression(ctx *parser.TriggerActionExpressionContext) {
}

// EnterTriggerActionExpressionStandalone is called when production triggerActionExpressionStandalone is entered.
func (s *GparserParserListener) EnterTriggerActionExpressionStandalone(ctx *parser.TriggerActionExpressionStandaloneContext) {
}

// ExitTriggerActionExpressionStandalone is called when production triggerActionExpressionStandalone is exited.
func (s *GparserParserListener) ExitTriggerActionExpressionStandalone(ctx *parser.TriggerActionExpressionStandaloneContext) {
}

// EnterCreateTriggerStatement is called when production createTriggerStatement is entered.
func (s *GparserParserListener) EnterCreateTriggerStatement(ctx *parser.CreateTriggerStatementContext) {
}

// ExitCreateTriggerStatement is called when production createTriggerStatement is exited.
func (s *GparserParserListener) ExitCreateTriggerStatement(ctx *parser.CreateTriggerStatementContext) {
}

// EnterAlterTriggerStatement is called when production alterTriggerStatement is entered.
func (s *GparserParserListener) EnterAlterTriggerStatement(ctx *parser.AlterTriggerStatementContext) {
}

// ExitAlterTriggerStatement is called when production alterTriggerStatement is exited.
func (s *GparserParserListener) ExitAlterTriggerStatement(ctx *parser.AlterTriggerStatementContext) {}

// EnterDropTriggerStatement is called when production dropTriggerStatement is entered.
func (s *GparserParserListener) EnterDropTriggerStatement(ctx *parser.DropTriggerStatementContext) {}

// ExitDropTriggerStatement is called when production dropTriggerStatement is exited.
func (s *GparserParserListener) ExitDropTriggerStatement(ctx *parser.DropTriggerStatementContext) {}

// EnterPoolAssign is called when production poolAssign is entered.
func (s *GparserParserListener) EnterPoolAssign(ctx *parser.PoolAssignContext) {}

// ExitPoolAssign is called when production poolAssign is exited.
func (s *GparserParserListener) ExitPoolAssign(ctx *parser.PoolAssignContext) {}

// EnterPoolAssignList is called when production poolAssignList is entered.
func (s *GparserParserListener) EnterPoolAssignList(ctx *parser.PoolAssignListContext) {}

// ExitPoolAssignList is called when production poolAssignList is exited.
func (s *GparserParserListener) ExitPoolAssignList(ctx *parser.PoolAssignListContext) {}

// EnterCreatePoolStatement is called when production createPoolStatement is entered.
func (s *GparserParserListener) EnterCreatePoolStatement(ctx *parser.CreatePoolStatementContext) {}

// ExitCreatePoolStatement is called when production createPoolStatement is exited.
func (s *GparserParserListener) ExitCreatePoolStatement(ctx *parser.CreatePoolStatementContext) {}

// EnterAlterPoolStatement is called when production alterPoolStatement is entered.
func (s *GparserParserListener) EnterAlterPoolStatement(ctx *parser.AlterPoolStatementContext) {}

// ExitAlterPoolStatement is called when production alterPoolStatement is exited.
func (s *GparserParserListener) ExitAlterPoolStatement(ctx *parser.AlterPoolStatementContext) {}

// EnterDropPoolStatement is called when production dropPoolStatement is entered.
func (s *GparserParserListener) EnterDropPoolStatement(ctx *parser.DropPoolStatementContext) {}

// ExitDropPoolStatement is called when production dropPoolStatement is exited.
func (s *GparserParserListener) ExitDropPoolStatement(ctx *parser.DropPoolStatementContext) {}

// EnterCreateMappingStatement is called when production createMappingStatement is entered.
func (s *GparserParserListener) EnterCreateMappingStatement(ctx *parser.CreateMappingStatementContext) {
}

// ExitCreateMappingStatement is called when production createMappingStatement is exited.
func (s *GparserParserListener) ExitCreateMappingStatement(ctx *parser.CreateMappingStatementContext) {
}

// EnterAlterMappingStatement is called when production alterMappingStatement is entered.
func (s *GparserParserListener) EnterAlterMappingStatement(ctx *parser.AlterMappingStatementContext) {
}

// ExitAlterMappingStatement is called when production alterMappingStatement is exited.
func (s *GparserParserListener) ExitAlterMappingStatement(ctx *parser.AlterMappingStatementContext) {}

// EnterDropMappingStatement is called when production dropMappingStatement is entered.
func (s *GparserParserListener) EnterDropMappingStatement(ctx *parser.DropMappingStatementContext) {}

// ExitDropMappingStatement is called when production dropMappingStatement is exited.
func (s *GparserParserListener) ExitDropMappingStatement(ctx *parser.DropMappingStatementContext) {}
