package headers

import (
	"errors"
	"fmt"
	"regexp"
)

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
func ParseRequestLine(first string) (*RequestLine, error) {
	pattern := regexp.MustCompile(`(?P<method>[a-z|A-Z]+)\s+`)
	match := pattern.FindStringSubmatch(first)
	if len(match) != 2 {
		return nil, errors.New("非请求头解析")
	}
	method := match[pattern.SubexpIndex("method")]
	requestURI := ParseRequestURI(first)
	return &RequestLine{
		Method:     method,
		RequestUri: *requestURI,
	}, nil
}
