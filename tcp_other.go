// +build !linux,!darwin

package main

import (
	"net"

	log "github.com/sirupsen/logrus"
)

func redirLocal(addr, server string, shadow func(net.Conn) net.Conn) {
	log.Printf("TCP redirect not supported")
}

func redir6Local(addr, server string, shadow func(net.Conn) net.Conn) {
	log.Printf("TCP6 redirect not supported")
}
