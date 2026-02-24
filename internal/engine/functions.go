
package engine

import "scaffold/internal/ast"

// TemplateFunc — функция шаблона.
// Получает аргументы, возвращает AST-фрагмент.
type TemplateFunc func(args []ast.Value) (ast.Node, error)

type FunctionRegistry struct {
    funcs map[string]TemplateFunc
}

func NewFunctionRegistry() *FunctionRegistry {
    return &FunctionRegistry{
        funcs: make(map[string]TemplateFunc),
    }
}

func (r *FunctionRegistry) Register(name string, fn TemplateFunc) {
    r.funcs[name] = fn
}

func (r *FunctionRegistry) Get(name string) (TemplateFunc, bool) {
    fn, ok := r.funcs[name]
    return fn, ok
}
