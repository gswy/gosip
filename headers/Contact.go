package headers

// Contact 对象
type Contact struct {
	Address Address
}

// String 格式化
func (c *Contact) String() string {
	return c.Address.String()
}

// ParseContact 解析
func ParseContact(data string) *Contact {
	field := ScanHeaderField(data, "Contact:")
	if field == nil {
		return nil
	}
	return &Contact{
		Address: ParseAddress(data),
	}
}
