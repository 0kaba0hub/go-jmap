package vacationresponse

import (
	"git.sr.ht/~rockorager/go-jmap"
	"git.sr.ht/~rockorager/go-jmap/internal/getids"
	"git.sr.ht/~rockorager/go-jmap/mail"
)

// Get vacation response details
// https://www.rfc-editor.org/rfc/rfc8621.html#section-8.1
type Get struct {
	Account jmap.ID `json:"accountId,omitempty"`

	IDs []jmap.ID `json:"ids,omitempty"`

	Properties []string `json:"properties,omitempty"`
}

// MarshalJSON keeps an empty ids as [] and sends "#ids" without "ids".
func (m Get) MarshalJSON() ([]byte, error) {
	type get Get
	return getids.Marshal(get(m), m.IDs, nil)
}

func (m *Get) Name() string { return "VacationResponse/get" }

func (m *Get) Requires() []jmap.URI { return []jmap.URI{mail.URI, URI} }

type GetResponse struct {
	Account jmap.ID `json:"accountId,omitempty"`

	State string `json:"state,omitempty"`

	List []*VacationResponse `json:"list,omitempty"`

	NotFound []jmap.ID `json:"notFound,omitempty"`
}

func newGetResponse() jmap.MethodResponse { return &GetResponse{} }
