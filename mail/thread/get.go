package thread

import (
	"git.sr.ht/~rockorager/go-jmap"
	"git.sr.ht/~rockorager/go-jmap/internal/getids"
	"git.sr.ht/~rockorager/go-jmap/mail"
)

// See RFC8621, Section 3.1.
type Get struct {
	Account jmap.ID `json:"accountId,omitempty"`

	IDs []jmap.ID `json:"ids,omitempty"`

	Properties []string `json:"properties,omitempty"`

	ReferenceIDs *jmap.ResultReference `json:"#ids,omitempty"`

	ReferenceProperties *jmap.ResultReference `json:"#properties,omitempty"`
}

// MarshalJSON keeps an empty ids as [] and sends "#ids" without "ids".
func (m Get) MarshalJSON() ([]byte, error) {
	type get Get
	return getids.Marshal(get(m), m.IDs, m.ReferenceIDs)
}

func (m *Get) Name() string { return "Thread/get" }

func (m *Get) Requires() []jmap.URI { return []jmap.URI{mail.URI} }

type GetResponse struct {
	Account jmap.ID `json:"accountId,omitempty"`

	State string `json:"state,omitempty"`

	List []*Thread `json:"list,omitempty"`

	NotFound []jmap.ID `json:"notFound,omitempty"`
}

func newGetResponse() jmap.MethodResponse { return &GetResponse{} }
