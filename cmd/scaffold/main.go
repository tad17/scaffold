package main

import (
    "scaffold/internal/ast"
    "scaffold/internal/engine"
    "scaffold/internal/ops"
)

func main() {

    // --- создаём AST вручную ---

    root := ast.SequenceNode{
        Items: []ast.Node{
            ast.TextNode{Value: "Hello from scaffold\n"},
            ast.TextNode{Value: "Engine works\n"},
        },
    }

    // --- запускаем engine ---

    eng := engine.Engine{}

    operations, err := eng.Evaluate(root, "test-output/engine.txt")
    if err != nil {
        panic(err)
    }

    // --- выполняем операции ---

    executor := ops.Executor{}
    if err := executor.Apply(operations); err != nil {
        panic(err)
    }
}
