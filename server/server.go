package main

import (
    "fmt"
    "net"
    "bufio"
    "time"
)

func main() {
    fmt.Println("Starting the server ...")
    // 创建 listener
    listener, err := net.Listen("tcp", "localhost:50000")
    if err != nil {
        fmt.Println("Error listening", err.Error())
        return //终止程序
    }
    // 监听并接受来自客户端的连接
    for {
        conn, err := listener.Accept()
        if err != nil {
            fmt.Println("Error accepting", err.Error())
            return // 终止程序
        }
        go doServerStuff(conn)
    }
}

func doServerStuff(conn net.Conn) {
    for {
        buf := make([]byte, 512)
        len, err := conn.Read(buf)
        ipStr :=conn.RemoteAddr().String()
        if err != nil {
            fmt.Println("Error reading", err.Error())
            return //终止程序
        }
        fmt.Printf(ipStr+": %v", string(buf[:len]))

        reader:=bufio.NewReader(conn)

        for {
            message, err := reader.ReadString('\n')
            if err != nil {
                return
            }
    
            fmt.Println(string(message))
            msg := time.Now().String() + "\n"
            b := []byte(msg)
           conn.Write(b)
        }
    }
}