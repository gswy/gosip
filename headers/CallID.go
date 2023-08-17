package headers

type CallID struct {
	Value string
}

func (c *CallID) String() string {
	return c.Value
}
