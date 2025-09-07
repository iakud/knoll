// Code generated from krpc.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // krpc
import "github.com/antlr4-go/antlr/v4"

// krpcListener is a complete listener for a parse tree produced by krpcParser.
type krpcListener interface {
	antlr.ParseTreeListener

	// EnterKrpc is called when entering the krpc production.
	EnterKrpc(c *KrpcContext)

	// ExitKrpc is called when exiting the krpc production.
	ExitKrpc(c *KrpcContext)
}
