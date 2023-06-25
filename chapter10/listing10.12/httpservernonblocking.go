package main

import (
    "fmt"
    "github.com/cutajarj/ConcurrentProgrammingWithGo/chapter10/listing10.9"
    "net"
)

/*
To trigger the "Too Many Requests" message try running this:
seq 1 2000 | xargs -Iname  -P100  curl -s "http://localhost:8080/index.html" | grep Busy
*/
func main() {
    incomingConnections := make(chan net.Conn)
    listing10_9.StartHttpWorkers(3, incomingConnections)
    server, _ := net.Listen("tcp", "localhost:8080")
    defer server.Close()
    for {
        conn, _ := server.Accept()
        select {
        case incomingConnections <- conn:
        default:
            fmt.Println("Server is busy")
            conn.Write([]byte("HTTP/1.1 429 Too Many Requests\r\n\r\n" +
                "<html>Busy</html>\n"))
            conn.Close()
        }
    }
}
