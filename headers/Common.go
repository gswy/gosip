package headers

import (
	"bufio"
	"fmt"
	"strings"
)

// Address URL地址对象
type Address struct {
	User *string
	Host *string
	Port *int
}

// String 格式化
func (a *Address) String() string {
	return fmt.Sprintf("<sip:%s@%s:%d>", *a.User, *a.Host, *a.Port)
}

// ParseAddress 解析
func ParseAddress(data string) *Address {
	return &Address{}
}

type RequestURI struct {
	User *string
	Host *string
	Port *int
}

// String 格式化
func (r *RequestURI) String() string {
	return fmt.Sprintf("sip:%s@%s:%d", *r.User, *r.Host, *r.Port)
}

// ParseRequestURI 解析
func ParseRequestURI() *RequestURI {
	return &RequestURI{}
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
