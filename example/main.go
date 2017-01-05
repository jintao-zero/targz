package main

import (
    "fmt"
    "os"
    "targz"
)

func main()  {
    fmt.Print(os.Getwd())
    err := targz.Tar("test", "test.tar")
    if err != nil {
        fmt.Println("tar fail", err)
        return
    }
    err = targz.Gzip("test.tar", "test.tar.gz")
    if err != nil {
        fmt.Println("gzip fail", err)
        return
    }

    err = targz.UnTargz("test.tar.gz", "./")
    if err != nil {
        fmt.Println("untargz faile", err)
        return
    }

}
