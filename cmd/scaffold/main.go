package main

import (
    "scaffold/internal/ops"
)

func main() {

    executor := ops.Executor{
        DryRun: false,
    }

    operations := []ops.Operation{
        ops.MkdirOp{
            Path: "test-output",
        },
        ops.WriteFileOp{
            Path: "test-output/hello.txt",
            Content: []byte("hello scaffold\n"),
        },
    }

    if err := executor.Apply(operations); err != nil {
        panic(err)
    }
}
