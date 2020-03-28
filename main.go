package main

// #cgo LDFLAGS: -L./output_library_dir -lfeles
// #include<stdio.h>
// #include<feles.h>
import "C"
import (
        "github.com/gorilla/mux"
        "io"
        "log"
        "net/http"
)

func main() {

        C.echo()

        srv := mux.NewRouter()

        srv.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
                _, err := io.WriteString(writer, "pong")
                if err != nil {
                        writer.WriteHeader(http.StatusInternalServerError)
                }
        })

        http.Handle("/", srv)

        if err := http.ListenAndServe(":8080", nil); err != nil {
                log.Fatal(err.Error())
        }
}