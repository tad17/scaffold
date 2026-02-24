
package config

import (
    "scaffold/internal/ast"
    "scaffold/internal/dsl"
    "scaffold/internal/engine"
)

func BuildTemplateStore(cfg *Config) *engine.TemplateStore {

    store := engine.NewTemplateStore()

    for name, lines := range cfg.Templates {

        var nodes []ast.Node

        for _, line := range lines {
            nodes = append(nodes, dsl.ParseLine(line))
        }

        store.Register(name, ast.SequenceNode{
            Items: nodes,
        })
    }

    return store
}

