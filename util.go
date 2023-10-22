package main

import (
    "fmt"
    "os"
)

func invalid(message string) {
    fmt.Fprintln(os.Stderr, message)
    os.Exit(1)
}


func pathExists(path string) bool {
    _, err := os.Stat(path)
    return err == nil
}
