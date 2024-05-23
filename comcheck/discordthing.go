package main

import (
    "fmt"
    "os"
    "strings"
)

func main() {
    args := os.Args[1:]

    for _, arg := range args {
        if strings.Contains(arg, "01")  strings.Contains(arg, "galaxy")  arg == "galaxy 01" {
            fmt.Println("Alert!!!")
            return
        }
    }
}