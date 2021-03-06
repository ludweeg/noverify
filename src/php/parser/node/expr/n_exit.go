package expr

import (
	"github.com/VKCOM/noverify/src/php/parser/freefloating"
	"github.com/VKCOM/noverify/src/php/parser/node"
	"github.com/VKCOM/noverify/src/php/parser/position"
	"github.com/VKCOM/noverify/src/php/parser/walker"
)

// Exit node
type Exit struct {
	FreeFloating freefloating.Collection
	Die          bool
	Position     *position.Position
	Expr         node.Node
}

// NewExit node constructor
func NewExit(Expr node.Node) *Exit {
	return &Exit{
		FreeFloating: nil,
		Expr:         Expr,
	}
}

// SetPosition sets node position
func (n *Exit) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *Exit) GetPosition() *position.Position {
	return n.Position
}

func (n *Exit) GetFreeFloating() *freefloating.Collection {
	return &n.FreeFloating
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Exit) Walk(v walker.Visitor) {
	if !v.EnterNode(n) {
		return
	}

	if n.Expr != nil {
		n.Expr.Walk(v)
	}

	v.LeaveNode(n)
}
