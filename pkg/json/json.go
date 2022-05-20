package json

import (
	"encoding/json"
	"io"
)

func ToString(reader io.Reader) string {
	var m map[string]interface{}
	_ = json.NewDecoder(reader).Decode(&m)
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
