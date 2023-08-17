package gosip

import (
	"fmt"
	"github.com/gswy/gosip/headers"
)

type Header struct {
	RequestLine      *headers.RequestLine
	Via              *headers.Via
	From             *headers.From
	To               *headers.To
	Contact          *headers.Contact
	CallID           *headers.CallID
	CSeq             *headers.CSeq
	Authorization    *headers.Authorization
	WWWAuthorization *headers.WWWAuthorization
}

func (h *Header) String() string {
	result := ""
	if h.Via != nil {
		result += fmt.Sprintf("Via: %s\n", h.Via.String())
	}
	if h.From != nil {
		result += fmt.Sprintf("From: %s\n", h.From.String())
	}
	if h.To != nil {
		result += fmt.Sprintf("To: %s\n", h.From.String())
	}
	if h.Contact != nil {
		result += fmt.Sprintf("Contact: %s\n", h.Contact.String())
	}
	if h.CallID != nil {
		result += fmt.Sprintf("Call-ID: %s\n", h.CallID.String())
	}
	if h.CSeq != nil {
		result += fmt.Sprintf("CSeq: %s\n", h.CSeq.String())
	}
	if h.Authorization != nil {
		result += fmt.Sprintf("Authorization: %s\n", h.Authorization.String())
	}
	if h.WWWAuthorization != nil {
		result += fmt.Sprintf("WWW-Authorization: %s\n", h.WWWAuthorization.String())
	}
	return result
}
