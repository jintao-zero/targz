# targz
Golang语言实现将文件或者文件夹打包并压缩成tar.gz文件的库
#Installation
run
    `go get github.com/jintao-zero/targz`
#Examples

    package main

    import (
        "fmt"
        "targz"
    )

    func main()  {
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
            fmt.Println("untargz fail", err)
            return
        }

        err = targz.Targz("test", "t.tar.gz")
        if err != nil {
            fmt.Println("targz fail", err)
            return
        }
    }
