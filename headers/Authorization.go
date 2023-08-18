package headers

type Authorization struct {
}

func (a *Authorization) String() string {
	return ""
}

func ParseAuthorization(data string) *Authorization {
	return &Authorization{}
}

type WWWAuthorization struct {
}

func (w *WWWAuthorization) String() string {
	return ""
}

func ParseWWWAuthorization(data string) *WWWAuthorization {
	return &WWWAuthorization{}
}
