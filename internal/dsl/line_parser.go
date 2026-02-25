package dsl

import (
    "strings"
    "scaffold/internal/logx"
    "scaffold/internal/ast"
)

func ParseLine(line string) ast.Node {

    logx.Debug("parse line: %q", line)

    line = strings.TrimSpace(line)

    if !strings.HasPrefix(line, "<-") {
        logx.Debug("  text node -> %s", line)
        return ast.TextNode{
            Value: line + "\n",
        }
    }

    body := strings.TrimSpace(line[2:])

    // --- HEADER ---
    if strings.HasPrefix(body, "template ") {
        logx.Debug("  header -> %s", body)
        name := strings.TrimSpace(strings.TrimPrefix(body, "template "))
        return ast.TemplateHeaderNode{
            Name: name,
        }
    }

    // --- function call with args ---
    if strings.Contains(body, "(") && strings.HasSuffix(body, ")") {
        logx.Debug("  function -> %s", body)
        name, args := parseCall(body)
        return ast.CallNode{
            Func: name,
            Args: args,
        }
    }

    // --- simple include ---
    logx.Debug("  include -> body=%s", body)
    return ast.IncludeNode{
        Target: body,
    }
}

func parseCall(s string) (string, []ast.Value) {

    //logx.Debug("  parse call: %s", s)

    open := strings.Index(s, "(")
    name := strings.TrimSpace(s[:open])

    argsPart := strings.TrimSuffix(s[open+1:], ")")

    if strings.TrimSpace(argsPart) == "" {
        return name, nil
    }

    parts := splitArgs(argsPart)

    var values []ast.Value
    for _, p := range parts {
        values = append(values, parseValue(p))
    }

    return name, values
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
