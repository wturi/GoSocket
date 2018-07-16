package main

import (
    "bufio"
    "fmt"
	"net"
	"time"
)


var ConnMap map[string]*net.TCPConn

func main() {
	var tcpAddr *net.TCPAddr
	ConnMap=make(map[string]*net.TCPConn)
	//tcpAddr, _ = net.ResolveTCPAddr("tcp", "127.0.0.1:5000")
	tcpAddr, _ = net.ResolveTCPAddr("tcp", "10.105.247.41:5000")
    tcpListener, error := net.ListenTCP("tcp", tcpAddr)

	if error!=nil{
		fmt.Println("server si error")
		tcpListener.Close()
	}

	fmt.Println("server is state")

    for {
        tcpConn, err := tcpListener.AcceptTCP()
        if err != nil {
            continue
        }

		

		fmt.Println("A client connected : " + tcpConn.RemoteAddr().String()+"---"+time.Now().String())
		ConnMap[tcpConn.RemoteAddr().String()]=tcpConn
        go tcpPipe(tcpConn)
    }

}

func tcpPipe(conn *net.TCPConn) {
    ipStr := conn.RemoteAddr().String()
    defer func() {
        fmt.Println("disconnected :" + ipStr+"---"+time.Now().String())
        conn.Close()
    }()
    reader := bufio.NewReader(conn)

    for {
        message, err := reader.ReadString('\n')
        if err != nil {
            return
        }

		
		fmt.Println(conn.RemoteAddr().String()+":"+string(message))
		//广播
		//boradcastMessage(conn.RemoteAddr().String() + ":" + string(message))

		//返回消息
        //fmt.Println(string(message))
        msg := time.Now().String() + "\n"
        b := []byte(msg)
        conn.Write(b)
    }
}

func boradcastMessage(message string) {
    b := []byte(message)
    // 遍历所有客户端并发送消息
    for _, conn := range ConnMap {
        conn.Write(b)
    }
}