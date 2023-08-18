package headers

import (
	"fmt"
	"regexp"
	"strconv"
)

// CSeq Header头中的CSeq字段
type CSeq struct {
	Sequence int
	Method   string
}

// String
func (c *CSeq) String() string {
	return fmt.Sprintf("%d %s", c.Sequence, c.Method)
}

// ParseCSeq 解析
func ParseCSeq(cseq string) *CSeq {
	field := ScanHeaderField(cseq, "CSeq:")
	if field == nil {
		return nil
	}
	pattern := regexp.MustCompile(`CSeq:\s+(?P<sequence>\d+)\s+(?P<method>[A-Z|a-z]+)`)
	match := pattern.FindStringSubmatch(*field)
	sequence := match[pattern.SubexpIndex("sequence")]
	method := match[pattern.SubexpIndex("method")]
	var seqNum int
	if r, ok := strconv.Atoi(sequence); ok == nil {
		seqNum = r
	}
	return &CSeq{
		Sequence: seqNum,
		Method:   method,
	}
}
