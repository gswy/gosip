package headers

import "fmt"

type Subject struct {
	Value string
}

func (s *Subject) String() string {
	return fmt.Sprintf("%s", s.Value)
}

func ParseSubject() *Subject {
	return &Subject{}
}
