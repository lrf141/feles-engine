package main

import "os"

func removeUnixDomainSock() error {
        return os.RemoveAll(unixDomainSockDir)
}