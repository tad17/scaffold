
package ops

import (
    "fmt"
    "os"
    "path/filepath"
)

type Executor struct {
    DryRun bool
}

func (e *Executor) Apply(list []Operation) error {
    for _, op := range list {

        switch v := op.(type) {

        case MkdirOp:
            if err := e.applyMkdir(v); err != nil {
                return err
            }

        case WriteFileOp:
            if err := e.applyWriteFile(v); err != nil {
                return err
            }

        default:
            return fmt.Errorf("unknown operation type %T", op)
        }
    }
    return nil
}

func (e *Executor) applyMkdir(op MkdirOp) error {
    if e.DryRun {
        fmt.Println("[dry-run] mkdir", op.Path)
        return nil
    }
    return os.MkdirAll(op.Path, 0755)
}

func (e *Executor) applyWriteFile(op WriteFileOp) error {

    if e.DryRun {
        fmt.Println("[dry-run] write file", op.Path)
        return nil
    }

    dir := filepath.Dir(op.Path)

    if err := os.MkdirAll(dir, 0755); err != nil {
        return err
    }

    return os.WriteFile(op.Path, op.Content, 0644)
}




