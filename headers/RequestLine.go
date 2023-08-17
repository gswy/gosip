package headers

import "fmt"

type RequestLine struct {
	Method     *string
	RequestUri *RequestURI
}

// String 转String方法
func (r RequestLine) String() string {
	return fmt.Sprintf("%s %s", *r.Method, r.RequestUri.String())
}
