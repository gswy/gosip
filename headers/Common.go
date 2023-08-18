package headers

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Address URL地址对象
type Address struct {
	User string
	Host string
	Port int
}

// String 格式化
func (a *Address) String() string {
	return fmt.Sprintf("<sip:%s@%s:%d>", a.User, a.Host, a.Port)
}

// ParseAddress 解析
func ParseAddress(address string) Address {
	pattern := regexp.MustCompile("<sip:(?P<user>\\d+)@(?P<host>[^:]+):(?P<port>\\d+)>")
	match := pattern.FindStringSubmatch(address)
	user := match[pattern.SubexpIndex("user")]
	host := match[pattern.SubexpIndex("host")]
	port := match[pattern.SubexpIndex("port")]
	portNumber, _ := strconv.Atoi(port)
	return Address{
		User: user,
		Host: host,
		Port: portNumber,
	}
}

type RequestURI struct {
	User string
	Host string
	Port int
}

// String 格式化
func (r *RequestURI) String() string {
	return fmt.Sprintf("sip:%s@%s:%d", r.User, r.Host, r.Port)
}

// ParseRequestURI 解析
func ParseRequestURI(data string) *RequestURI {
	pattern := regexp.MustCompile("sip:(?P<user>\\d+)@(?P<ip>[^:]+):(?P<port>\\d+)")
	match := pattern.FindStringSubmatch(data)
	user := match[pattern.SubexpIndex("user")]
	ip := match[pattern.SubexpIndex("ip")]
	port := match[pattern.SubexpIndex("port")]
	portNumber, _ := strconv.Atoi(port)
	return &RequestURI{
		User: user,
		Host: ip,
		Port: portNumber,
	}
}

// ScanHeaderField 逐行扫描头部字段行
func ScanHeaderField(data string, prefix string) *string {
	// 逐行读取文本
	scanner := bufio.NewScanner(strings.NewReader(data))
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(strings.TrimSpace(line), prefix) {
			return &line
		}
	}
	return nil
}

// ScanHeaderFirst 扫描头部信息
func ScanHeaderFirst(data string) *string {
	// 扫描第一行文本即返回
	// 逐行读取文本
	scanner := bufio.NewScanner(strings.NewReader(data))
	for scanner.Scan() {
		line := scanner.Text()
		return &line
	}
	return nil
}
