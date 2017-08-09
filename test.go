package main

import (
    "fmt"
    "os"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage:\n", os.Args[0], "/path/to/file")
        return
    }

    file, err := os.Open(os.Args[1])
    if err != nil {
        fmt.Println(err)
        return
    }
    defer file.Close()

    stat, err := file.Stat()
    if err != nil {
        fmt.Println(err)
        return
    }

    tmp := make([]byte, 1)
    count := 0
    for i := int64(0); i < stat.Size(); i++ {
        _, err = file.Read(tmp)
        if err != nil {
            fmt.Println(err)
            return
        }
        if tmp[0] == '\n' {
            count++
        }
    }

    fmt.Println(count)
}