package main

// #cgo LDFLAGS: -L./output_library_dir -lfeles
// #include<stdio.h>
// #include<feles.h>
import "C"
import (
        "log"
)

func main() {

        C.echo()

        err := initDaemon()
        if err != nil {
                log.Fatal(err.Error())
        }

        startFelesDaemon()
}