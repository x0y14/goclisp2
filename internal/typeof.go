package internal

type TypeKind int

const (
	Illegal TypeKind = iota
	TInt
	TStr
)

type TypeError struct {
	Msg string
}

func (e *TypeError) Error() string { return e.Msg }

func TypeOf(expr ExprNode) (TypeKind, error) {
	switch n := expr.(type) {
	case *IntNode:
		return TInt, nil
	case *StringNode:
		return TStr, nil
	case *BinaryNode:
		lt, err := TypeOf(n.Left)
		if err != nil {
			return 0, err
		}
		rt, err := TypeOf(n.Right)
		if err != nil {
			return 0, err
		}
		switch n.Op {
		case OpAdd, OpSub:
			if lt != TInt || rt != TInt {
				return 0, &TypeError{Msg: "arith expects int,int"}
			}
			return TInt, nil
		default:
			return 0, &TypeError{Msg: "unknown binop"}
		}
	default:
		return 0, &TypeError{Msg: "unknown expr"}
	}
}
