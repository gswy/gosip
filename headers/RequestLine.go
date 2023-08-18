package headers

import "fmt"

// RequestLine 请求行
type RequestLine struct {
	Method     string
	RequestUri RequestURI
}

// String 转String方法
func (r *RequestLine) String() string {
	return fmt.Sprintf("%s %s", r.Method, r.RequestUri.String())
}

// ParseRequestLine 解析
func ParseRequestLine(data string) *RequestLine {
	return &RequestLine{}
}
