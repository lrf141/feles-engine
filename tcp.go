package main

import (
        "bufio"
        "fmt"
        "io/ioutil"
        "log"
        "net"
        "net/http"
        "strings"
)

func initDaemon() error {

        err := removeUnixDomainSock()
        if err != nil {
                return err
        }
        return nil
}

func startFelesDaemon() {

        srv, err := net.Listen(protoType, unixDomainSockDir)
        if err != nil {
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
                        log.Println(err.Error())
                        continue
                }
                go receiver(conn)
        }
}

func receiver(conn net.Conn) {

        log.Printf("Accept: %v\n", conn.RemoteAddr())

        resp := getDefaultHttpResponse()

        req, err := http.ReadRequest(bufio.NewReader(conn))
        if err != nil {
                log.Printf("Can't read HTTP request: %v\n", conn.RemoteAddr())
                sendErrorMsg(resp, &conn)
                return
        }

        fmt.Println(req)

        resp.Body = ioutil.NopCloser(strings.NewReader("Pong\n"))
        err = resp.Write(conn)
        if err != nil {
                log.Printf("Can't send HTTP response to %v\n", conn.RemoteAddr())
        }
        err = conn.Close()
        if err != nil {
                log.Println("Can't close conenction")
        }
}
