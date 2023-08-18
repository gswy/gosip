package gosip

import (
	"github.com/gswy/gosip/headers"
	"net"
)

type Request struct {
	Header Header
	Body   *Body
}

type Response struct {
}

func (r *Response) decode() {

}

func (r *Response) encode() {

}

type Context struct {
	conn     *net.UDPConn
	addr     net.Addr
	Request  Request
	Response Response
}

func (c Context) RESP(status headers.StatusLine, response Response) {

}

type callback func(ctx *Context)
