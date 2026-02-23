package ast

// Node — базовый интерфейс всех элементов AST.
type Node interface {
    isNode()
}

// --------------------
// Literal text
// --------------------

type TextNode struct {
    Value string
}

func (TextNode) isNode() {}

// --------------------
// Sequence of nodes
// --------------------

type SequenceNode struct {
    Items []Node
}

func (SequenceNode) isNode() {}

// --------------------
// Include reference
// <- name
// <- name(args)
// --------------------

type IncludeNode struct {
    Target string
    Args   []Value
}

func (IncludeNode) isNode() {}

// --------------------
// Function call
// header(...)
// --------------------

type CallNode struct {
    Func string
    Args []Value
}

func (CallNode) isNode() {}


