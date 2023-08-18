package gosip

import (
	"fmt"
	"github.com/gswy/gosip/headers"
	"log"
	"net"
)

type Request struct {
	Header Header
	Body   Body
}

type Response struct {
	Header Header
	Body   Body
}

type Context struct {
	conn     *net.UDPConn
	addr     net.Addr
	Request  Request
	Response Response
}

func (r *Response) String() string {
	return fmt.Sprintf("%s\n%s", r.Header.String(), r.Body.String())
}

func (r *Response) Bytes() []byte {
	return []byte(r.String())
}

// RESP 返回消息
func (c Context) RESP(status headers.StatusLine, response Response) {
	response.Header.StatusLine = &status
	_, err := c.conn.WriteTo(response.Bytes(), c.addr)
	if err != nil {
		log.Printf("[SIP] 回执消息错误: %v", err)
	}
}

type callback func(ctx *Context)
