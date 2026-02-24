package main

import (
    "scaffold/internal/ast"
    "scaffold/internal/engine"
    "scaffold/internal/ops"
)

func main() {

    // --- registry функций ---
    registry := engine.NewFunctionRegistry()
    engine.RegisterBuiltins(registry)

    // --- store шаблонов ---
    templates := engine.NewTemplateStore()

    templates.Register("content",
        ast.TextNode{Value: "This is included content\n"},
    )

    // --- engine ---
    eng := engine.Engine{
        Functions: registry,
        Templates: templates,
    }

    // --- root AST ---
    root := ast.SequenceNode{
        Items: []ast.Node{
            ast.CallNode{Func: "header"},
            ast.IncludeNode{Target: "content"},
            ast.CallNode{Func: "footer"},
        },
    }

    operations, err := eng.Evaluate(root, "test-output/include.txt")
    if err != nil {
        panic(err)
    }

    executor := ops.Executor{}
    if err := executor.Apply(operations); err != nil {
        panic(err)
    }
}
