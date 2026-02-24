package main

import (
    "scaffold/internal/ast"
    "scaffold/internal/engine"
    "scaffold/internal/ops"
)

func main() {

    registry := engine.NewFunctionRegistry()
    engine.RegisterBuiltins(registry)

    eng := engine.Engine{
        Functions: registry,
    }

    root := ast.SequenceNode{
        Items: []ast.Node{
            ast.CallNode{Func: "header"},
            ast.TextNode{Value: "Body content\n"},
            ast.CallNode{Func: "footer"},
        },
    }

    operations, err := eng.Evaluate(root, "test-output/functions.txt")
    if err != nil {
        panic(err)
    }

    executor := ops.Executor{}
    if err := executor.Apply(operations); err != nil {
        panic(err)
    }
}
