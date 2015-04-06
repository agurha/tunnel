package proto

import (
	"github.com/agurha/tunnel/pkg/conn"
)

type Protocol interface {
	GetName() string
	WrapConn(conn.Conn, interface{}) conn.Conn
}
