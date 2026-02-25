package logx

import "fmt"

var DebugEnabled = false

func Debug(format string, args ...any) {
    if !DebugEnabled {
        return
    }
    fmt.Printf("[DEBUG] "+format+"\n", args...)
}

func Info(format string, args ...any) {
    fmt.Printf("[INFO] "+format+"\n", args...)
}

func Error(format string, args ...any) {
    fmt.Printf("[ERROR] "+format+"\n", args...)
}

func Fatal(format string, args ...any) {
    Error(format, args...)
    panic(fmt.Sprintf(format, args...))
}
