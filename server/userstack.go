package main

import "net"

type UserStack struct {
	Users []*net.Conn
}
