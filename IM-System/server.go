package main

import (
	"fmt"
	"io"
	"net"
	"sync"
)

type Server struct {
	IP   string
	Port int

	OnLineMap map[string]*User
	MapLock   sync.RWMutex
	Message   chan string
}

func NewServer(ip string, port int) *Server {
	server := &Server{
		IP:        ip,
		Port:      port,
		OnLineMap: make(map[string]*User),
		Message:   make(chan string),
	}
	return server
}

func (this *Server) Start() {
	// socket listen
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.IP, this.Port))
	if err != nil {
		fmt.Println("net listen error", err)
		return
	}
	defer listener.Close()
	go this.ListenMessage()
	for {
		//accept
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener accept err", err)
			continue
		}
		// do handler
		go this.Handler(conn)
	}

}

// ListenMessage 监听广播消息
func (this *Server) ListenMessage() {
	for {
		msg := <-this.Message
		this.MapLock.Lock()
		for _, cli := range this.OnLineMap {
			cli.C <- msg
		}
		this.MapLock.Unlock()
	}
}

// 广播消息
func (this *Server) BroadCast(user *User, msg string) {
	senMgs := "[" + user.Addr + "]" + user.Name + ":" + msg
	this.Message <- senMgs
}

func (this *Server) Handler(conn net.Conn) {
	user := NewUser(conn, this)
	user.Online()
	// 接受客户端发送的消息
	go func() {
		buf := make([]byte, 4096)
		for true {
			n, err := conn.Read(buf)
			if n == 0 {
				user.Offline()
				return
			}
			if err != nil && err != io.EOF {
				return
			}
			msg := string(buf[:n-1])
			user.DoMessage(msg)
		}
	}()
	select {}
	//fmt.Println("server start")
}
