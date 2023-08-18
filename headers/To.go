package headers

// To Header中To字段
type To struct {
	Address Address
}

// String 格式化
func (t *To) String() string {
	return t.Address.String()
}

// ParseTo 解析
func ParseTo(data string) *To {
	field := ScanHeaderField(data, "To:")
	if field == nil {
		return nil
	}
	address := ParseAddress(*field)
	return &To{
		Address: address,
	}
}
