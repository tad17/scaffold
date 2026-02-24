
package dsl

import (
    "strings"
    "scaffold/internal/ast"
)

func ParseLine(line string) ast.Node {

    line = strings.TrimSpace(line)

    if strings.HasPrefix(line, "<-") {

        body := strings.TrimSpace(line[2:])

        // function call
        if strings.HasSuffix(body, "()") {
            name := strings.TrimSuffix(body, "()")
            return ast.CallNode{
                Func: name,
            }
        }

        // include
        return ast.IncludeNode{
            Target: body,
        }
    }

    return ast.TextNode{
        Value: line + "\n",
    }
}
