// Package getids encodes the ids argument of a /get request.
package getids

import (
	"encoding/json"

	"git.sr.ht/~rockorager/go-jmap"
)

// Marshal encodes v, a /get request whose ids field is omitempty: with a
// reference only "#ids" is sent (RFC 8620 §3.7), and an empty non-nil ids is
// sent as [] instead of being dropped, which would mean every record (§5.1).
func Marshal(v any, ids []jmap.ID, ref *jmap.ResultReference) ([]byte, error) {
	raw, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	emptyList := ref == nil && ids != nil && len(ids) == 0
	if ref == nil && !emptyList {
		return raw, nil
	}
	var obj map[string]json.RawMessage
	if err := json.Unmarshal(raw, &obj); err != nil {
		return nil, err
	}
	if ref != nil {
		delete(obj, "ids")
	} else {
		obj["ids"] = json.RawMessage("[]")
	}
	return json.Marshal(obj)
}
