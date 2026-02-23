
package ops

type MkdirOp struct {
    Path string
}

func (MkdirOp) isOperation() {}

