package internal

type Node interface {
	Pos() Position
}

type StmtNode interface {
	Node
	stmtNode()
}

type ExprNode interface {
	Node
	exprNode()
}

type BlockNode struct {
	pos   Position
	Stmts []StmtNode
}

func (b *BlockNode) Pos() Position {
	return b.pos
}
func (b *BlockNode) stmtNode() {}

type DefineFunctionNode struct {
	pos    Position
	Name   string
	Params []IdentNode
	Block  *BlockNode
}

func (d *DefineFunctionNode) Pos() Position {
	return d.pos
}
func (d *DefineFunctionNode) stmtNode() {}

type IdentNode struct {
	pos  Position
	Name string
}

func (i *IdentNode) Pos() Position {
	return i.pos
}
func (i *IdentNode) exprNode() {}

type ExprStmt struct {
	pos Position
	X   ExprNode
}

func (s *ExprStmt) Pos() Position { return s.pos }
func (s *ExprStmt) stmtNode()     {}

type BinaryOp int

const (
	OpAdd BinaryOp = iota
	OpSub
)

type BinaryNode struct {
	pos         Position
	Left, Right ExprNode
	Op          BinaryOp
}

func (b *BinaryNode) Pos() Position {
	return b.pos
}
func (b *BinaryNode) exprNode() {}

type IntNode struct {
	pos Position
	V   int
}

func (i *IntNode) Pos() Position {
	return i.pos
}
func (i *IntNode) exprNode() {}

type StringNode struct {
	pos Position
	V   string
}

func (s *StringNode) Pos() Position {
	return s.pos
}
func (s *StringNode) exprNode() {}
