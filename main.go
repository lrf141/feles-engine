package main

// #cgo LDFLAGS: -L./output_library_dir -lfeles
// #include<stdio.h>
// #include<feles.h>
// #include<ns.h>
import "C"
import (
        "log"
)

func main() {

        C.echo()
        C.clone_ns()

        err := initDaemon()
        if err != nil {
                log.Fatal(err.Error())
        }

        startFelesDaemon()
}