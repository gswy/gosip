package headers

type To struct {
	Address *Address
}

func (t *To) String() string {
	return t.Address.String()
}

func ParseTo(data string) *To {
	return &To{}
}
