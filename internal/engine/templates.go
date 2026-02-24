
package engine

import "scaffold/internal/ast"

type TemplateStore struct {
    templates map[string]ast.Node
}

func NewTemplateStore() *TemplateStore {
    return &TemplateStore{
        templates: make(map[string]ast.Node),
    }
}

func (s *TemplateStore) Register(name string, node ast.Node) {
    s.templates[name] = node
}

func (s *TemplateStore) Get(name string) (ast.Node, bool) {
    n, ok := s.templates[name]
    return n, ok
}
