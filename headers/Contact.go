package headers

type Contact struct {
	Address *Address
}

func (c *Contact) String() string {
	return c.Address.String()
}
