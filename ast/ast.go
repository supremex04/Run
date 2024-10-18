package ast
import "run/token"

type Node interface {
	TokenLiteral() string
}
type Statement interface {
	Node 
	statementNode()
}
type Expression interface {
	Node
	expressionNode()
}