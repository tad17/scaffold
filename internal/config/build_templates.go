
package config

import (
    "scaffold/internal/ast"
    "scaffold/internal/dsl"
    "scaffold/internal/engine"
)

func BuildTemplateStore(cfg *Config) (*engine.TemplateStore, error) {

    store := engine.NewTemplateStore()

    for name, lines := range cfg.Templates {

        var nodes []ast.Node

        for _, line := range lines {
            node, err := dsl.ParseLine(line)
            if err != nil {
                return nil, engine.ExecError{
                    Op: "function",
                    Name: "BuildTemplateStore",
                    Err: err,
                }
            }
            nodes = append(nodes, node)
        }

        store.Register(name, ast.SequenceNode{
            Items: nodes,
        })
    }

    return store, nil
}

