package engine

import (
    "errors"
    "scaffold/internal/ast"
)    

type DependencyGraph struct {
    Edges map[string][]string
}

func NewDependencyGraph() *DependencyGraph {
    return &DependencyGraph{
        Edges: make(map[string][]string),
    }
}

func collectIncludes(node ast.Node, out *[]string) {

    switch v := node.(type) {

    case ast.IncludeNode:
        *out = append(*out, v.Target)

    case ast.SequenceNode:
        for _, n := range v.Items {
            collectIncludes(n, out)
        }
    }
}

func BuildDependencyGraph(store *TemplateStore) *DependencyGraph {

    g := NewDependencyGraph()

    for name, tmpl := range store.templates {

        var deps []string
        collectIncludes(tmpl, &deps)

        g.Edges[name] = deps
    }

    return g
}

func (g *DependencyGraph) DetectCycles() error {

    visited := map[string]bool{}
    stack := map[string]bool{}

    var visit func(string) error

    visit = func(n string) error {

        if stack[n] {
            return ExecError{
                Op: "dependency",
                Name: n,
                Err:  errors.New("cycle detected"),
            }
        }

        if visited[n] {
            return nil
        }

        visited[n] = true
        stack[n] = true

        for _, next := range g.Edges[n] {
            if err := visit(next); err != nil {
                return err
            }
        }

        stack[n] = false
        return nil
    }

    for n := range g.Edges {
        if err := visit(n); err != nil {
            return err
        }
    }

    return nil
}

