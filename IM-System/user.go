package main

import (
	"net"
)

type User struct {
	Name   string
	Addr   string
	C      chan string
	conn   net.Conn
	server *Server
}

// NewUser 创建用户
func NewUser(conn net.Conn, server *Server) *User {
	userAddr := conn.RemoteAddr().String()

	user := &User{
		Name:   userAddr,
		Addr:   userAddr,
		C:      make(chan string),
		conn:   conn,
		server: server,
	}
	// 启动监听当前user
	go user.ListenMessage()
	return user
}

// ListenMessage 监听当前user
func (this *User) ListenMessage() {
	for {
		msg := <-this.C
		this.conn.Write([]byte(msg + "\n"))
	}
}

// Online 用户上线
func (this *User) Online() {
	this.server.MapLock.Lock()
	this.server.OnLineMap[this.Name] = this
	this.server.MapLock.Unlock()
	this.server.BroadCast(this, "已上线")
}

// 用户下线
func (this *User) Offline() {
	this.server.MapLock.Lock()
	this.server.OnLineMap[this.Name] = this
	this.server.MapLock.Unlock()
	this.server.BroadCast(this, "已下线")
}

func (this *User) DoMessage(msg string) {
	this.server.BroadCast(this, msg)
}
