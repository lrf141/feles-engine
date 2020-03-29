package main

// #cgo LDFLAGS: -L./output_library_dir -lfeles
// #include<stdio.h>
// #include<feles.h>
import "C"
import (
        "bufio"
        "fmt"
        "io/ioutil"
        "log"
        "net"
        "net/http"
        "net/http/httputil"
        "os"
        "strings"
)

func main() {

        C.echo()

        err := os.RemoveAll(unixDomainSockDir)
        if err != nil {
                log.Fatal(err.Error())
        }

        // support unix domain socket only
        srv, err := net.Listen(protoType, unixDomainSockDir)
        if err != nil {
                log.Println("Server Listen Error")
                log.Fatal(err.Error())
        }

        defer func() {
                err := srv.Close()
                if err != nil {
                        log.Fatal(err.Error())
                }
        }()

        for {
                conn, err := srv.Accept()
                if err != nil {

                }
                go func() {
                        fmt.Printf("Accept:%v\n", conn.RemoteAddr())
                        request, _ := http.ReadRequest(bufio.NewReader(conn))
                        dump,_ := httputil.DumpRequest(request,true)
                        fmt.Println(string(dump))

                        response := http.Response{
                                StatusCode:       200,
                                ProtoMajor:       1,
                                ProtoMinor:       0,
                                Body:             ioutil.NopCloser(strings.NewReader("Hello I am Response.")),
                        }
                        err := response.Write(conn)
                        if err != nil {
                                log.Println(err)
                                response.StatusCode = http.StatusInternalServerError
                        }
                        err = conn.Close()
                        if err != nil {
                                log.Println("Connection could not close.")
                        }
                }()
        }
}