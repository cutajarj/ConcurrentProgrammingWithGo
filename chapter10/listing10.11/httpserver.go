package main

import (
    "github.com/cutajarj/ConcurrentProgrammingWithGo/chapter10/listing10.9"
    "net"
)

/*
If you're having problems running this, please make sure that the port 8080 is available
Note that error handling is omitted so that listing is brief
*/
func main() {
    incomingConnections := make(chan net.Conn)
    listing10_9.StartHttpWorkers(3, incomingConnections)
    server, _ := net.Listen("tcp", "localhost:8080")
    defer server.Close()
    for {
        conn, _ := server.Accept()
        incomingConnections <- conn
    }
}
