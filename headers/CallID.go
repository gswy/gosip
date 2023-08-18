package headers

// CallID 对象
type CallID struct {
	Value string
}

// String 格式化
func (c *CallID) String() string {
	return c.Value
}

// ParseCallID 解析
func ParseCallID(data string) *CallID {
	field := ScanHeaderField(data, "Call-ID:")
	if field == nil {
		return nil
	}
	return &CallID{Value: ""}
}
