package headers

import "fmt"

type From struct {
	Address *Address
	Tag     *string
}

func (f *From) String() string {
	result := fmt.Sprintf("%s", f.Address.String())
	if f.Tag != nil {
		result += fmt.Sprintf(";tag=%s", *f.Tag)
	}
	return result
}