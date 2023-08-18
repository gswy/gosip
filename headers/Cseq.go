package headers

import "fmt"

type CSeq struct {
	Sequence int
	Method   string
}

func (c *CSeq) String() string {
	return fmt.Sprintf("%d %s", c.Sequence, c.Method)
}

func ParseCSeq(data string) *CSeq {
	field := ScanHeaderField(data, "CSeq:")
	if field == nil {
		return nil
	}
	return &CSeq{}
}
