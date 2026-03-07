package main

import (
    "scaffold/internal/logx"
    "scaffold/internal/config"
    "scaffold/internal/engine"
    "scaffold/internal/ops"
)

func main() {

    logx.DebugEnabled = true // временно всегда включено

    cfg, err := config.Load("example.yaml")
    if err != nil {
        panic(err)
    }

    templates, err := config.BuildTemplateStore(cfg)
    if err != nil {
        panic(err)
    }

    // ===============================
    // 🔴 ВСТАВИТЬ СЮДА
    // ===============================

    graph := engine.BuildDependencyGraph(templates)

    if err := graph.DetectCycles(); err != nil {
        panic(err)
    }

    for k, v := range graph.Edges {
    logx.Debug("template %s depends on %v", k, v)
}
    // ===============================
    // дальше всё как было
    // ===============================

    registry := engine.NewFunctionRegistry()
    engine.RegisterBuiltins(registry)

    eng := engine.Engine{
        Functions: registry,
        Templates: templates,
    }

    root, ok := templates.Get("main")
    if !ok {
        panic("templates")
    }

    operations, err := eng.Evaluate(root, "test-output/yaml.txt")
    if err != nil {
        panic(err)
    }

    executor := ops.Executor{}
    if err := executor.Apply(operations); err != nil {
        panic(err)
    }
}