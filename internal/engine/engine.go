package engine

import (
    "bytes"
    "fmt"
    "scaffold/internal/ast"
    "scaffold/internal/ops"
)

type Engine struct {
}

// Evaluate — главная точка входа.
// Возвращает операции, которые нужно выполнить.
func (e *Engine) Evaluate(root ast.Node, outputPath string) ([]ops.Operation, error) {

    data, err := e.render(root)
    if err != nil {
        return nil, err
    }

    return []ops.Operation{
        ops.WriteFileOp{
            Path:    outputPath,
            Content: data,
        },
    }, nil
}

func (e *Engine) render(node ast.Node) ([]byte, error) {

    var buf bytes.Buffer

    if err := e.walk(node, &buf); err != nil {
        return nil, err
    }

    return buf.Bytes(), nil
}

func (e *Engine) walk(node ast.Node, out *bytes.Buffer) error {

    switch v := node.(type) {

    case ast.TextNode:
        out.WriteString(v.Value)
        return nil

    case ast.SequenceNode:
        for _, item := range v.Items {
            if err := e.walk(item, out); err != nil {
                return err
            }
        }
        return nil

    case ast.IncludeNode:
        return fmt.Errorf("include not implemented yet")

    case ast.CallNode:
        return fmt.Errorf("call not implemented yet")

    default:
        return fmt.Errorf("unknown node type %T", node)
    }
}
