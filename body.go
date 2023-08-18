package gosip

import (
	"bufio"
	"fmt"
	"strings"
)

type Body struct {
	Value *string
}

// ParseBody 解析Body
func ParseBody(data string) *Body {
	scanner := bufio.NewScanner(strings.NewReader(data))
	result := ""
	isBody := false
	for scanner.Scan() {
		line := scanner.Text()
		if isBody {
			result += fmt.Sprintf("%s\n", line)
		}
		if strings.TrimSpace(line) == "" {
			isBody = true
		}
	}
	if isBody {
		return &Body{Value: &result}
	} else {
		return nil
	}
}
