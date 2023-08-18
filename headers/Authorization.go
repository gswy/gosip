package headers

type Authorization struct {
}

func (a *Authorization) String() string {
	return ""
}

func ParseAuthorization(data string) *Authorization {
	field := ScanHeaderField(data, "Authorization:")
	if field == nil {
		return nil
	}
	return &Authorization{}
}

type WWWAuthorization struct {
}

func (w *WWWAuthorization) String() string {
	return ""
}

func ParseWWWAuthorization(data string) *WWWAuthorization {
	field := ScanHeaderField(data, "WWW-Authorization:")
	if field == nil {
		return nil
	}
	return &WWWAuthorization{}
}
