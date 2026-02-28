package dsl

import (
    "strings"
    "errors"
    "scaffold/internal/engine"
    "scaffold/internal/logx"
    "scaffold/internal/ast"
)

func ParseLine(line string) (ast.Node, error) {

    logx.Debug("parse line: %q", line)

    line = strings.TrimSpace(line)

    if !strings.HasPrefix(line, "<-") {
        logx.Debug("  text node -> %s", line)
        return ast.TextNode{
            Value: line + "\n",
        }, nil
    }

    body := strings.TrimSpace(line[2:])

    // --- HEADER ---
    if strings.HasPrefix(body, "template ") {
        logx.Debug("  header -> %s", body)
        name := strings.TrimSpace(strings.TrimPrefix(body, "template "))
        return ast.TemplateHeaderNode{
            Name: name,
        }, nil
    }

    // --- function call with args ---
    if strings.Contains(body, "(") && strings.HasSuffix(body, ")") {
        logx.Debug("  function -> %s", body)
        name, args, err := parseCall(body)
        if err != nil {
            return nil, engine.ExecError{
                Op: "func",
                Name: name,
                Err: err,
            }
        }
        return ast.CallNode{
            Func: name,
            Args: args,
        }, nil
    }

    // --- simple include ---
    logx.Debug("  include -> body=%s", body)
    return ast.IncludeNode{
        Target: body,
    }, nil
}

func parseCall(s string) (string, []ast.Value, error) {

    //logx.Debug("  parse call: %s", s)

    open := strings.Index(s, "(")
    name := strings.TrimSpace(s[:open])

    argsPart := strings.TrimSuffix(s[open+1:], ")")

    if strings.TrimSpace(argsPart) == "" {
        return name, nil, nil
    }

    parts := splitArgs(argsPart)

    var values []ast.Value
    for _, p := range parts {
        values = append(values, parseValue(p))
    }

    err := engine.ExecError{
                Op: "func",
                Name: "parseCall",
                Err: errors.New("ошибка для проверки"),
            }
    return name, values, err
}

func splitArgs(s string) []string {
    return strings.Split(s, ",")
}

func parseValue(s string) ast.Value {

    s = strings.TrimSpace(s)

    // string literal
    if strings.HasPrefix(s, `"`) && strings.HasSuffix(s, `"`) {
        return ast.StringValue{
            Value: s[1 : len(s)-1],
        }
    }

    // fallback — raw string
    return ast.StringValue{
        Value: s,
    }
}
