package headers

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

const (
	StatusTrying       = 100
	StatusOk           = 200
	StatusUnauthorized = 401
)

// StatusLine 响应状态
type StatusLine struct {
	StatusCode int
}

// String 转换String
func (s *StatusLine) String() string {
	msg := StatusText(s.StatusCode)
	return fmt.Sprintf("SIP/2.0 %d %s", s.StatusCode, msg)
}

// StatusText 返回状态信息
func StatusText(status int) string {
	switch status {
	case StatusTrying:
		return "Trying"
	case StatusOk:
		return "OK"
	case StatusUnauthorized:
		return "Unauthorized"
	default:
		return ""
	}
}

func ParseStatusLine(first string) (*StatusLine, error) {
	pattern := regexp.MustCompile(`SIP/2\.0\s+(?P<status>\d+)`)
	match := pattern.FindStringSubmatch(first)
	if len(match) != 2 {
		return nil, errors.New("非请求头解析")
	}
	status := match[pattern.SubexpIndex("status")]
	statusCode, _ := strconv.Atoi(status)
	return &StatusLine{
		StatusCode: statusCode,
	}, nil
}

// HandleOk 处理成功
func HandleOk() StatusLine {
	return StatusLine{StatusCode: 200}
}

// HandleUnauthorized 未授权
func HandleUnauthorized() StatusLine {
	return StatusLine{StatusCode: 401}
}

// HandleTrying 特殊状态
func HandleTrying() StatusLine {
	return StatusLine{StatusCode: 100}
}
