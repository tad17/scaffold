package engine

import (
    "bytes"
    "fmt"
    "scaffold/internal/ast"
    "scaffold/internal/ops"
)

type Engine struct {
    Functions *FunctionRegistry
    Templates *TemplateStore
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

        if e.Templates == nil {
            return fmt.Errorf("no template store configured")
        }

        tmpl, ok := e.Templates.Get(v.Target)
        if !ok {
            return fmt.Errorf("unknown template %s", v.Target)
        }

        return e.walk(tmpl, out)
    
    case ast.CallNode:

        if e.Functions == nil {
            return fmt.Errorf("no function registry configured")
        }

        fn, ok := e.Functions.Get(v.Func)
        if !ok {
            return fmt.Errorf("unknown function %s", v.Func)
        }

        result, err := fn(v.Args)
        if err != nil {
            return err
        }

        return e.walk(result, out)
    
    default:
        return fmt.Errorf("unknown node type %T", node)
    }
}
