// Code generated from krpc.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // krpc
import "github.com/antlr4-go/antlr/v4"

// BasekrpcListener is a complete listener for a parse tree produced by krpcParser.
type BasekrpcListener struct{}

var _ krpcListener = &BasekrpcListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasekrpcListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasekrpcListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasekrpcListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasekrpcListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterKrpc is called when production krpc is entered.
func (s *BasekrpcListener) EnterKrpc(ctx *KrpcContext) {}

// ExitKrpc is called when production krpc is exited.
func (s *BasekrpcListener) ExitKrpc(ctx *KrpcContext) {}
