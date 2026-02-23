
package ops

type WriteFileOp struct {
    Path    string
    Content []byte
}

func (WriteFileOp) isOperation() {}

