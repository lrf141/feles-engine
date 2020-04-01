package main

// #cgo LDFLAGS: -L./output_library_dir -lfeles -luuid /usr/lib/x86_64-linux-gnu/libuuid.so
// #include<stdio.h>
// #include<feles.h>
// #include<ns.h>
import "C"
import (
        "log"
)

func main() {

        C.echo()
        C.create_namespace()

        err := initDaemon()
        if err != nil {
                log.Fatal(err.Error())
        }

        startFelesDaemon()
}