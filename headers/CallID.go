package headers

import (
	"strings"
)

// CallID 对象
type CallID struct {
	Value string
}

// String 格式化
func (c *CallID) String() string {
	return c.Value
}

// ParseCallID 解析
func ParseCallID(callId string) *CallID {
	field := ScanHeaderField(callId, "Call-ID:")
	if field == nil {
		return nil
	}
	value := strings.Replace(*field, "Call-ID:", "", 1)
	value = strings.TrimSpace(value)
	return &CallID{Value: value}
}
