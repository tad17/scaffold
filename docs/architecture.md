# Scaffold Architecture

Goal: declarative project generator with DSL and execution engine.

Core layers:

- config  — YAML loading
- dsl     — custom syntax parsing
- ast     — execution tree
- engine  — evaluation runtime
- ops     — filesystem operations
- project — project context and metadata

Execution pipeline:

YAML → DSL parse → AST → evaluate → operations → apply
```

Это нужно, чтобы ты сам не потерял замысел через месяц 🙂


