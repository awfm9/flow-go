// Code generated from parser/Strictus.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Strictus
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseStrictusListener is a complete listener for a parse tree produced by StrictusParser.
type BaseStrictusListener struct{}

var _ StrictusListener = &BaseStrictusListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseStrictusListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseStrictusListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseStrictusListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseStrictusListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseStrictusListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseStrictusListener) ExitProgram(ctx *ProgramContext) {}

// EnterDeclaration is called when production declaration is entered.
func (s *BaseStrictusListener) EnterDeclaration(ctx *DeclarationContext) {}

// ExitDeclaration is called when production declaration is exited.
func (s *BaseStrictusListener) ExitDeclaration(ctx *DeclarationContext) {}

// EnterFunctionDeclaration is called when production functionDeclaration is entered.
func (s *BaseStrictusListener) EnterFunctionDeclaration(ctx *FunctionDeclarationContext) {}

// ExitFunctionDeclaration is called when production functionDeclaration is exited.
func (s *BaseStrictusListener) ExitFunctionDeclaration(ctx *FunctionDeclarationContext) {}

// EnterParameterList is called when production parameterList is entered.
func (s *BaseStrictusListener) EnterParameterList(ctx *ParameterListContext) {}

// ExitParameterList is called when production parameterList is exited.
func (s *BaseStrictusListener) ExitParameterList(ctx *ParameterListContext) {}

// EnterParameter is called when production parameter is entered.
func (s *BaseStrictusListener) EnterParameter(ctx *ParameterContext) {}

// ExitParameter is called when production parameter is exited.
func (s *BaseStrictusListener) ExitParameter(ctx *ParameterContext) {}

// EnterFullType is called when production fullType is entered.
func (s *BaseStrictusListener) EnterFullType(ctx *FullTypeContext) {}

// ExitFullType is called when production fullType is exited.
func (s *BaseStrictusListener) ExitFullType(ctx *FullTypeContext) {}

// EnterTypeDimension is called when production typeDimension is entered.
func (s *BaseStrictusListener) EnterTypeDimension(ctx *TypeDimensionContext) {}

// ExitTypeDimension is called when production typeDimension is exited.
func (s *BaseStrictusListener) ExitTypeDimension(ctx *TypeDimensionContext) {}

// EnterBaseType is called when production baseType is entered.
func (s *BaseStrictusListener) EnterBaseType(ctx *BaseTypeContext) {}

// ExitBaseType is called when production baseType is exited.
func (s *BaseStrictusListener) ExitBaseType(ctx *BaseTypeContext) {}

// EnterFunctionType is called when production functionType is entered.
func (s *BaseStrictusListener) EnterFunctionType(ctx *FunctionTypeContext) {}

// ExitFunctionType is called when production functionType is exited.
func (s *BaseStrictusListener) ExitFunctionType(ctx *FunctionTypeContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseStrictusListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseStrictusListener) ExitBlock(ctx *BlockContext) {}

// EnterStatements is called when production statements is entered.
func (s *BaseStrictusListener) EnterStatements(ctx *StatementsContext) {}

// ExitStatements is called when production statements is exited.
func (s *BaseStrictusListener) ExitStatements(ctx *StatementsContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseStrictusListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseStrictusListener) ExitStatement(ctx *StatementContext) {}

// EnterReturnStatement is called when production returnStatement is entered.
func (s *BaseStrictusListener) EnterReturnStatement(ctx *ReturnStatementContext) {}

// ExitReturnStatement is called when production returnStatement is exited.
func (s *BaseStrictusListener) ExitReturnStatement(ctx *ReturnStatementContext) {}

// EnterIfStatement is called when production ifStatement is entered.
func (s *BaseStrictusListener) EnterIfStatement(ctx *IfStatementContext) {}

// ExitIfStatement is called when production ifStatement is exited.
func (s *BaseStrictusListener) ExitIfStatement(ctx *IfStatementContext) {}

// EnterWhileStatement is called when production whileStatement is entered.
func (s *BaseStrictusListener) EnterWhileStatement(ctx *WhileStatementContext) {}

// ExitWhileStatement is called when production whileStatement is exited.
func (s *BaseStrictusListener) ExitWhileStatement(ctx *WhileStatementContext) {}

// EnterVariableDeclaration is called when production variableDeclaration is entered.
func (s *BaseStrictusListener) EnterVariableDeclaration(ctx *VariableDeclarationContext) {}

// ExitVariableDeclaration is called when production variableDeclaration is exited.
func (s *BaseStrictusListener) ExitVariableDeclaration(ctx *VariableDeclarationContext) {}

// EnterAssignment is called when production assignment is entered.
func (s *BaseStrictusListener) EnterAssignment(ctx *AssignmentContext) {}

// ExitAssignment is called when production assignment is exited.
func (s *BaseStrictusListener) ExitAssignment(ctx *AssignmentContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseStrictusListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseStrictusListener) ExitExpression(ctx *ExpressionContext) {}

// EnterConditionalExpression is called when production conditionalExpression is entered.
func (s *BaseStrictusListener) EnterConditionalExpression(ctx *ConditionalExpressionContext) {}

// ExitConditionalExpression is called when production conditionalExpression is exited.
func (s *BaseStrictusListener) ExitConditionalExpression(ctx *ConditionalExpressionContext) {}

// EnterOrExpression is called when production orExpression is entered.
func (s *BaseStrictusListener) EnterOrExpression(ctx *OrExpressionContext) {}

// ExitOrExpression is called when production orExpression is exited.
func (s *BaseStrictusListener) ExitOrExpression(ctx *OrExpressionContext) {}

// EnterAndExpression is called when production andExpression is entered.
func (s *BaseStrictusListener) EnterAndExpression(ctx *AndExpressionContext) {}

// ExitAndExpression is called when production andExpression is exited.
func (s *BaseStrictusListener) ExitAndExpression(ctx *AndExpressionContext) {}

// EnterEqualityExpression is called when production equalityExpression is entered.
func (s *BaseStrictusListener) EnterEqualityExpression(ctx *EqualityExpressionContext) {}

// ExitEqualityExpression is called when production equalityExpression is exited.
func (s *BaseStrictusListener) ExitEqualityExpression(ctx *EqualityExpressionContext) {}

// EnterRelationalExpression is called when production relationalExpression is entered.
func (s *BaseStrictusListener) EnterRelationalExpression(ctx *RelationalExpressionContext) {}

// ExitRelationalExpression is called when production relationalExpression is exited.
func (s *BaseStrictusListener) ExitRelationalExpression(ctx *RelationalExpressionContext) {}

// EnterAdditiveExpression is called when production additiveExpression is entered.
func (s *BaseStrictusListener) EnterAdditiveExpression(ctx *AdditiveExpressionContext) {}

// ExitAdditiveExpression is called when production additiveExpression is exited.
func (s *BaseStrictusListener) ExitAdditiveExpression(ctx *AdditiveExpressionContext) {}

// EnterMultiplicativeExpression is called when production multiplicativeExpression is entered.
func (s *BaseStrictusListener) EnterMultiplicativeExpression(ctx *MultiplicativeExpressionContext) {}

// ExitMultiplicativeExpression is called when production multiplicativeExpression is exited.
func (s *BaseStrictusListener) ExitMultiplicativeExpression(ctx *MultiplicativeExpressionContext) {}

// EnterUnaryExpression is called when production unaryExpression is entered.
func (s *BaseStrictusListener) EnterUnaryExpression(ctx *UnaryExpressionContext) {}

// ExitUnaryExpression is called when production unaryExpression is exited.
func (s *BaseStrictusListener) ExitUnaryExpression(ctx *UnaryExpressionContext) {}

// EnterPrimaryExpression is called when production primaryExpression is entered.
func (s *BaseStrictusListener) EnterPrimaryExpression(ctx *PrimaryExpressionContext) {}

// ExitPrimaryExpression is called when production primaryExpression is exited.
func (s *BaseStrictusListener) ExitPrimaryExpression(ctx *PrimaryExpressionContext) {}

// EnterPrimaryExpressionSuffix is called when production primaryExpressionSuffix is entered.
func (s *BaseStrictusListener) EnterPrimaryExpressionSuffix(ctx *PrimaryExpressionSuffixContext) {}

// ExitPrimaryExpressionSuffix is called when production primaryExpressionSuffix is exited.
func (s *BaseStrictusListener) ExitPrimaryExpressionSuffix(ctx *PrimaryExpressionSuffixContext) {}

// EnterEqualityOp is called when production equalityOp is entered.
func (s *BaseStrictusListener) EnterEqualityOp(ctx *EqualityOpContext) {}

// ExitEqualityOp is called when production equalityOp is exited.
func (s *BaseStrictusListener) ExitEqualityOp(ctx *EqualityOpContext) {}

// EnterRelationalOp is called when production relationalOp is entered.
func (s *BaseStrictusListener) EnterRelationalOp(ctx *RelationalOpContext) {}

// ExitRelationalOp is called when production relationalOp is exited.
func (s *BaseStrictusListener) ExitRelationalOp(ctx *RelationalOpContext) {}

// EnterAdditiveOp is called when production additiveOp is entered.
func (s *BaseStrictusListener) EnterAdditiveOp(ctx *AdditiveOpContext) {}

// ExitAdditiveOp is called when production additiveOp is exited.
func (s *BaseStrictusListener) ExitAdditiveOp(ctx *AdditiveOpContext) {}

// EnterMultiplicativeOp is called when production multiplicativeOp is entered.
func (s *BaseStrictusListener) EnterMultiplicativeOp(ctx *MultiplicativeOpContext) {}

// ExitMultiplicativeOp is called when production multiplicativeOp is exited.
func (s *BaseStrictusListener) ExitMultiplicativeOp(ctx *MultiplicativeOpContext) {}

// EnterUnaryOp is called when production unaryOp is entered.
func (s *BaseStrictusListener) EnterUnaryOp(ctx *UnaryOpContext) {}

// ExitUnaryOp is called when production unaryOp is exited.
func (s *BaseStrictusListener) ExitUnaryOp(ctx *UnaryOpContext) {}

// EnterIdentifierExpression is called when production IdentifierExpression is entered.
func (s *BaseStrictusListener) EnterIdentifierExpression(ctx *IdentifierExpressionContext) {}

// ExitIdentifierExpression is called when production IdentifierExpression is exited.
func (s *BaseStrictusListener) ExitIdentifierExpression(ctx *IdentifierExpressionContext) {}

// EnterLiteralExpression is called when production LiteralExpression is entered.
func (s *BaseStrictusListener) EnterLiteralExpression(ctx *LiteralExpressionContext) {}

// ExitLiteralExpression is called when production LiteralExpression is exited.
func (s *BaseStrictusListener) ExitLiteralExpression(ctx *LiteralExpressionContext) {}

// EnterFunctionExpression is called when production FunctionExpression is entered.
func (s *BaseStrictusListener) EnterFunctionExpression(ctx *FunctionExpressionContext) {}

// ExitFunctionExpression is called when production FunctionExpression is exited.
func (s *BaseStrictusListener) ExitFunctionExpression(ctx *FunctionExpressionContext) {}

// EnterNestedExpression is called when production NestedExpression is entered.
func (s *BaseStrictusListener) EnterNestedExpression(ctx *NestedExpressionContext) {}

// ExitNestedExpression is called when production NestedExpression is exited.
func (s *BaseStrictusListener) ExitNestedExpression(ctx *NestedExpressionContext) {}

// EnterExpressionAccess is called when production expressionAccess is entered.
func (s *BaseStrictusListener) EnterExpressionAccess(ctx *ExpressionAccessContext) {}

// ExitExpressionAccess is called when production expressionAccess is exited.
func (s *BaseStrictusListener) ExitExpressionAccess(ctx *ExpressionAccessContext) {}

// EnterMemberAccess is called when production memberAccess is entered.
func (s *BaseStrictusListener) EnterMemberAccess(ctx *MemberAccessContext) {}

// ExitMemberAccess is called when production memberAccess is exited.
func (s *BaseStrictusListener) ExitMemberAccess(ctx *MemberAccessContext) {}

// EnterBracketExpression is called when production bracketExpression is entered.
func (s *BaseStrictusListener) EnterBracketExpression(ctx *BracketExpressionContext) {}

// ExitBracketExpression is called when production bracketExpression is exited.
func (s *BaseStrictusListener) ExitBracketExpression(ctx *BracketExpressionContext) {}

// EnterInvocation is called when production invocation is entered.
func (s *BaseStrictusListener) EnterInvocation(ctx *InvocationContext) {}

// ExitInvocation is called when production invocation is exited.
func (s *BaseStrictusListener) ExitInvocation(ctx *InvocationContext) {}

// EnterArgument is called when production argument is entered.
func (s *BaseStrictusListener) EnterArgument(ctx *ArgumentContext) {}

// ExitArgument is called when production argument is exited.
func (s *BaseStrictusListener) ExitArgument(ctx *ArgumentContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseStrictusListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseStrictusListener) ExitLiteral(ctx *LiteralContext) {}

// EnterBooleanLiteral is called when production booleanLiteral is entered.
func (s *BaseStrictusListener) EnterBooleanLiteral(ctx *BooleanLiteralContext) {}

// ExitBooleanLiteral is called when production booleanLiteral is exited.
func (s *BaseStrictusListener) ExitBooleanLiteral(ctx *BooleanLiteralContext) {}

// EnterDecimalLiteral is called when production DecimalLiteral is entered.
func (s *BaseStrictusListener) EnterDecimalLiteral(ctx *DecimalLiteralContext) {}

// ExitDecimalLiteral is called when production DecimalLiteral is exited.
func (s *BaseStrictusListener) ExitDecimalLiteral(ctx *DecimalLiteralContext) {}

// EnterBinaryLiteral is called when production BinaryLiteral is entered.
func (s *BaseStrictusListener) EnterBinaryLiteral(ctx *BinaryLiteralContext) {}

// ExitBinaryLiteral is called when production BinaryLiteral is exited.
func (s *BaseStrictusListener) ExitBinaryLiteral(ctx *BinaryLiteralContext) {}

// EnterOctalLiteral is called when production OctalLiteral is entered.
func (s *BaseStrictusListener) EnterOctalLiteral(ctx *OctalLiteralContext) {}

// ExitOctalLiteral is called when production OctalLiteral is exited.
func (s *BaseStrictusListener) ExitOctalLiteral(ctx *OctalLiteralContext) {}

// EnterHexadecimalLiteral is called when production HexadecimalLiteral is entered.
func (s *BaseStrictusListener) EnterHexadecimalLiteral(ctx *HexadecimalLiteralContext) {}

// ExitHexadecimalLiteral is called when production HexadecimalLiteral is exited.
func (s *BaseStrictusListener) ExitHexadecimalLiteral(ctx *HexadecimalLiteralContext) {}

// EnterInvalidNumberLiteral is called when production InvalidNumberLiteral is entered.
func (s *BaseStrictusListener) EnterInvalidNumberLiteral(ctx *InvalidNumberLiteralContext) {}

// ExitInvalidNumberLiteral is called when production InvalidNumberLiteral is exited.
func (s *BaseStrictusListener) ExitInvalidNumberLiteral(ctx *InvalidNumberLiteralContext) {}

// EnterArrayLiteral is called when production arrayLiteral is entered.
func (s *BaseStrictusListener) EnterArrayLiteral(ctx *ArrayLiteralContext) {}

// ExitArrayLiteral is called when production arrayLiteral is exited.
func (s *BaseStrictusListener) ExitArrayLiteral(ctx *ArrayLiteralContext) {}

// EnterEos is called when production eos is entered.
func (s *BaseStrictusListener) EnterEos(ctx *EosContext) {}

// ExitEos is called when production eos is exited.
func (s *BaseStrictusListener) ExitEos(ctx *EosContext) {}
