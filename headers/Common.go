package headers

import "fmt"

type Address struct {
	User *string
	Host *string
	Port *int
}

func (a *Address) String() string {
	return fmt.Sprintf("<sip:%s@%s:%d>", *a.User, *a.Host, *a.Port)
}

type RequestURI struct {
	User *string
	Host *string
	Port *int
}

func (r *RequestURI) String() string {
	return fmt.Sprintf("sip:%s@%s:%d", *r.User, *r.Host, *r.Port)
}
