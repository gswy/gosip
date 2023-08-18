package headers

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Via struct {
	Transport   string
	SentAddress string
	SentPort    int
	RPort       *string
	Branch      *string
	Received    *string
}

func (v *Via) String() string {
	result := fmt.Sprintf("SIP/2.0/%s %s:%d", v.Transport, v.SentAddress, v.SentPort)
	if v.RPort != nil {
		result += fmt.Sprintf(";rport")
	}
	if v.Branch != nil {
		result += fmt.Sprintf(";branch=%s", *v.Branch)
	}
	if v.Received != nil {
		result += fmt.Sprintf(";received=%s", *v.Received)
	}
	return result
}

func ParseVia(data string) *Via {
	// 扫描via行
	field := ScanHeaderField(data, "Via:")
	if field == nil {
		return nil
	}
	// 扫描到后解析via（正则匹配）
	pattern := regexp.MustCompile(`SIP/2\.0/(?P<transport>[A-Z|a-z]+)\s+(?P<sentAddr>[^:]+):(?P<sentPort>\d+)`)
	match := pattern.FindStringSubmatch(*field)
	transport := match[pattern.SubexpIndex("transport")]
	sentAddr := match[pattern.SubexpIndex("sentAddr")]
	sentPort := match[pattern.SubexpIndex("sentPort")]
	portNumber, _ := strconv.Atoi(sentPort)

	// 解析其他参数
	split := strings.Split(*field, ";")
	result := make(map[string]*string)
	for _, part := range split {
		if strings.HasPrefix(part, "Via:") {
			continue
		}
		// 以等号分割
		pair := strings.SplitN(part, "=", 2)
		if len(pair) == 2 {
			result[pair[0]] = &pair[1]
		} else {
			// 可能存在rport这类数据，做兼容处理，rport后的内容可省略
			// 因此可从sentPort取出默认数据进行赋值操作
			if pair[0] == "rport" {
				result["rport"] = &sentPort
			}
		}
	}
	var branch *string
	if val, ok := result["branch"]; ok {
		branch = val
	}
	var rport *string
	if val, ok := result["rport"]; ok {
		rport = val
	}
	var received *string
	if val, ok := result["received"]; ok {
		received = val
	}
	return &Via{
		Transport:   transport,
		SentAddress: sentAddr,
		SentPort:    portNumber,
		RPort:       rport,
		Branch:      branch,
		Received:    received,
	}
}
