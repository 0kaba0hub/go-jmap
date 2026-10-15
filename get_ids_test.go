package jmap_test

import (
	"encoding/json"
	"testing"

	"git.sr.ht/~rockorager/go-jmap"
	"git.sr.ht/~rockorager/go-jmap/core/push/subscription"
	"git.sr.ht/~rockorager/go-jmap/mail/email"
	"git.sr.ht/~rockorager/go-jmap/mail/emailsubmission"
	"git.sr.ht/~rockorager/go-jmap/mail/identity"
	"git.sr.ht/~rockorager/go-jmap/mail/mailbox"
	"git.sr.ht/~rockorager/go-jmap/mail/thread"
	"git.sr.ht/~rockorager/go-jmap/mail/vacationresponse"
)

// sentArgs is the arguments object m is sent with inside a request.
func sentArgs(t *testing.T, m jmap.Method) map[string]json.RawMessage {
	t.Helper()
	req := &jmap.Request{}
	req.Invoke(m)
	raw, err := json.Marshal(req)
	if err != nil {
		t.Fatal(err)
	}
	var env struct {
		Calls [][]json.RawMessage `json:"methodCalls"`
	}
	if err := json.Unmarshal(raw, &env); err != nil {
		t.Fatal(err)
	}
	var args map[string]json.RawMessage
	if err := json.Unmarshal(env.Calls[0][1], &args); err != nil {
		t.Fatal(err)
	}
	return args
}

// An empty ids asks for no records, a reference replaces ids, and a nil ids
// with no reference stays absent (RFC 8620 §3.7, §5.1).
func TestGetIDsEncoding(t *testing.T) {
	ref := &jmap.ResultReference{ResultOf: "0", Name: "Email/query", Path: "/ids"}
	some := []jmap.ID{"a"}
	cases := []struct {
		name    string
		m       jmap.Method
		ids     string // "" when the key must be absent
		withRef bool
	}{
		{"Email ids []", &email.Get{IDs: []jmap.ID{}}, "[]", false},
		{"Email #ids alone", &email.Get{ReferenceIDs: ref}, "", true},
		{"Email #ids with ids set", &email.Get{IDs: some, ReferenceIDs: ref}, "", true},
		{"Email nil ids", &email.Get{}, "", false},
		{"Email ids listed", &email.Get{IDs: some}, `["a"]`, false},
		{"Mailbox ids []", &mailbox.Get{IDs: []jmap.ID{}}, "[]", false},
		{"Mailbox #ids", &mailbox.Get{IDs: some, ReferenceIDs: ref}, "", true},
		{"Thread ids []", &thread.Get{IDs: []jmap.ID{}}, "[]", false},
		{"Thread #ids", &thread.Get{IDs: some, ReferenceIDs: ref}, "", true},
		{"Identity ids []", &identity.Get{IDs: []jmap.ID{}}, "[]", false},
		{"Identity #ids", &identity.Get{IDs: some, ReferenceIDs: ref}, "", true},
		{"EmailSubmission ids []", &emailsubmission.Get{IDs: []jmap.ID{}}, "[]", false},
		{"EmailSubmission #ids", &emailsubmission.Get{IDs: some, ReferenceIDs: ref}, "", true},
		{"PushSubscription ids []", &subscription.Get{IDs: []jmap.ID{}}, "[]", false},
		{"VacationResponse ids []", &vacationresponse.Get{IDs: []jmap.ID{}}, "[]", false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			args := sentArgs(t, tc.m)
			got, has := args["ids"]
			switch {
			case tc.ids == "" && has:
				t.Errorf("ids sent as %s, want no ids key", got)
			case tc.ids != "" && string(got) != tc.ids:
				t.Errorf("ids sent as %q, want %s", got, tc.ids)
			}
			if _, hasRef := args["#ids"]; hasRef != tc.withRef {
				t.Errorf("#ids present %v, want %v", hasRef, tc.withRef)
			}
		})
	}
}
