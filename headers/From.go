package headers

import (
	"fmt"
	"strings"
)

// From Header中From字段
type From struct {
	Address Address
	Tag     *string
}

// String 格式化
func (f *From) String() string {
	result := fmt.Sprintf("%s", f.Address.String())
	if f.Tag != nil {
		result += fmt.Sprintf(";tag=%s", *f.Tag)
	}
	return result
}

// ParseFrom 解析
func ParseFrom(data string) *From {
	field := ScanHeaderField(data, "From:")
	if field == nil {
		return nil
	}
	address := ParseAddress(*field)
	split := strings.Split(*field, ";")
	result := make(map[string]*string)
	for _, part := range split {
		if strings.HasPrefix(part, "From:") {
			continue
		}
		// 以等号分割
		pair := strings.SplitN(part, "=", 2)
		if len(pair) == 2 {
			result[pair[0]] = &pair[1]
		}
	}
	var tag *string
	if val, ok := result["tag"]; ok {
		tag = val
	}
	return &From{
		Address: address,
		Tag:     tag,
	}
}
