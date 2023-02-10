package main

import (
    "fmt"
    "os"
    "io"
)

func main() {
    args := os.Args[0:]
    opened_file, err := os.Open(args[0:])
    if err != nil {
        fmt.Println("Error: ", err)
        os.Exit(1)
    }
    io.Copy(opened_file, opened_file.Body)
}
