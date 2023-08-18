package headers

import "fmt"

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
		result += fmt.Sprintf(";report=%s", *v.RPort)
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
	return &Via{}
}
