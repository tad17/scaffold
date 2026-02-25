
package engine

import "scaffold/internal/ast"

func RegisterBuiltins(r *FunctionRegistry) {

    r.Register("header", func(args []ast.Value) (ast.Node, error) {

        title := "HEADER"

        if len(args) > 0 {
            if s, ok := args[0].(ast.StringValue); ok {
                title = s.Value
            }
        }

        return ast.TextNode{
            Value: "=== " + title + " ===\n",
        }, nil
    })

    r.Register("footer", func(args []ast.Value) (ast.Node, error) {
        return ast.TextNode{
            Value: "\n=== FOOTER ===\n",
        }, nil
    })
}
