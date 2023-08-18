package headers

import (
	"fmt"
	"strconv"
	"strings"
)

// Expires 过期时间
type Expires struct {
	Value int
}

// String 格式化
func (e *Expires) String() string {
	return fmt.Sprintf("%d", e.Value)
}

// ParseExpires 解析
func ParseExpires(data string) *Expires {
	field := ScanHeaderField(data, "Expires:")
	if field == nil {
		return nil
	}
	replace := strings.Replace(*field, "Expires:", "", 1)
	space := strings.TrimSpace(replace)
	val, _ := strconv.Atoi(space)
	return &Expires{
		Value: val,
	}
}
