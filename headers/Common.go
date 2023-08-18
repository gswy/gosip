package headers

import (
	"bufio"
	"fmt"
	"strings"
)

type Address struct {
	User *string
	Host *string
	Port *int
}

func (a *Address) String() string {
	return fmt.Sprintf("<sip:%s@%s:%d>", *a.User, *a.Host, *a.Port)
}

func ParseAddress(data string) *Address {
	return &Address{}
}

type RequestURI struct {
	User *string
	Host *string
	Port *int
}

func (r *RequestURI) String() string {
	return fmt.Sprintf("sip:%s@%s:%d", *r.User, *r.Host, *r.Port)
}

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
