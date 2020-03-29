package main

import (
        "io/ioutil"
        "net"
        "net/http"
        "strings"
)

func getDefaultHttpResponse() *http.Response {
        return &http.Response{
                StatusCode: 200,
                ProtoMajor: 1,
                ProtoMinor: 1,
        }
}

func writeInternalErrorResponse(response *http.Response) {
        response.StatusCode = http.StatusInternalServerError
        response.Header.Set("Accept-Type", "application/json")
        response.Body = ioutil.NopCloser(strings.NewReader("{'message': 'server internal error'}"))
}

func sendErrorMsg(resp *http.Response, conn *net.Conn) {
        writeInternalErrorResponse(resp)
        resp.Write(*conn)
        (*conn).Close()
}