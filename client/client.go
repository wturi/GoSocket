package main

import (
    "bufio"
    "fmt"
    "net"
)

func main() {
    var tcpAddr *net.TCPAddr
    //tcpAddr, _ = net.ResolveTCPAddr("tcp", "127.0.0.1:5000")
	tcpAddr, _ = net.ResolveTCPAddr("tcp", "115.159.98.198:5000")
    conn, _ := net.DialTCP("tcp", nil, tcpAddr)
    defer conn.Close()
    fmt.Println("connected!")

    go onMessageRecived(conn)

    // 控制台聊天功能加入
    for {
        var msg string
        fmt.Scanln(&msg)
        if msg == "quit" {
            break
        }
        b := []byte(msg + "\n")
        conn.Write(b)
    }
}

func onMessageRecived(conn *net.TCPConn) {
    reader := bufio.NewReader(conn)
    for {
        msg, err := reader.ReadString('\n')
        fmt.Println(msg)
        if err != nil {
            //quitSemaphore <- true
            break
        }
    }
}