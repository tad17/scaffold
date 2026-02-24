
package engine

import "scaffold/internal/ast"

func RegisterBuiltins(r *FunctionRegistry) {

    r.Register("header", func(args []ast.Value) (ast.Node, error) {

        return ast.TextNode{
            Value: "=== HEADER ===\n",
        }, nil
    })

    r.Register("footer", func(args []ast.Value) (ast.Node, error) {

        return ast.TextNode{
            Value: "\n=== FOOTER ===\n",
        }, nil
    })
}
