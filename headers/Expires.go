package headers

import "fmt"

// Expires 过期时间
type Expires struct {
	Value int
}

func (e *Expires) String() string {
	return fmt.Sprintf("%d", e.Value)
}

func ParseExpires(data string) *Expires {
	return &Expires{}
}
