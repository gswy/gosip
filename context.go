package gosip

import (
	"fmt"
	"net"
	"net/http"
)

type Request struct {
	Header Header
	Body   Body
}

type Response struct {
}

func (r *Response) decode() {

}

func (r *Response) encode() {

}

type Context struct {
	conn *net.UDPConn
	addr net.Addr
	Request
	Response
}

func (c Context) RESP(status int, response Response) {
	text := http.StatusText(status)
	_, err := c.conn.WriteTo([]byte(fmt.Sprintf("SIP/2.0 %d %s\n"+
		"CSeq: 1 Register\n"+
		"Call-ID: f3g4h51686094899@192.168.0.100\n"+
		"From: <sip:34020000002000001026@192.168.0.100:5060>;tag=1121243483\n"+
		"To: <sip:34020000002000001026@192.168.0.100:5060>\n"+
		"Via: SIP/2.0/UDP 192.168.0.100:5060;rport;branch=z9hG4bK107008409"+
		"Content-Length: 0\n", status, text)), c.addr)
	if err != nil {
		fmt.Println("写入错误:", err)
	}
}

type callback func(ctx *Context)
