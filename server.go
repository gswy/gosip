package gosip

import (
	"log"
	"net"
	"strings"
)

// Engine 服务类
type Engine struct {
	RouteGroup
	IP   string
	Port int
	node map[string]callback
	conn *net.UDPConn
}

// NewEngine 创建SIP服务器
func NewEngine(ip string, port int) *Engine {
	engine := &Engine{
		IP:   ip,
		Port: port,
		node: make(map[string]callback),
	}
	engine.RouteGroup = RouteGroup{engine: engine}
	return engine
}

// Run 运行SIP服务器
func (s *Engine) Run() error {
	var err error
	laddr := &net.UDPAddr{IP: net.ParseIP(s.IP), Port: s.Port}
	s.conn, err = net.ListenUDP("udp", laddr)
	if err != nil {
		return err
	}
	for {
		buffer := make([]byte, 1024)
		n, addr, err := s.conn.ReadFrom(buffer)
		if err != nil {
			log.Printf("[SIP] 请求数据获取失败: %v\n", err)
			continue
		}
		s.handle(addr, buffer[:n])
	}
}

// 根据不同的请求分发处理
func (s *Engine) handle(addr net.Addr, bytes []byte) {
	go func() {
		data := string(bytes)
		methodIndex := strings.Index(data, " ")
		method := strings.ToUpper(strings.TrimSpace(data[:methodIndex]))
		c, ok := s.node[method]
		if !ok {
			log.Printf("[SIP] 找不到对应的处理请求方法: %s\n", method)
			return
		}
		c(&Context{
			conn: s.conn,
			addr: addr,
			Request: Request{
				Header: ParseHeader(data),
				Body:   ParseBody(data),
			},
		})
	}()
}

// addMethod 注册路由方法
func (s *Engine) addMethod(method string, call callback) {
	s.node[method] = call
}
