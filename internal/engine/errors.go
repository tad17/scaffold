package engine

import "fmt"

type ExecError struct {
    Op   string
    Name string
    Err  error
}

func (e ExecError) Error() string {
    return fmt.Sprintf("%s %s \n<- %v", e.Op, e.Name, e.Err)
}

func (e ExecError) Unwrap() error {
    return e.Err
}

