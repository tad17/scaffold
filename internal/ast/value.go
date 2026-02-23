package ast

// Value — аргументы функций и include.
type Value interface {
    isValue()
}

type StringValue struct {
    Value string
}

func (StringValue) isValue() {}

type NumberValue struct {
    Value float64
}

func (NumberValue) isValue() {}

type BoolValue struct {
    Value bool
}

func (BoolValue) isValue() {}

