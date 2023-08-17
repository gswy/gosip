package headers

type To struct {
	Address *Address
}

func (t *To) String() string {
	return t.Address.String()
}
