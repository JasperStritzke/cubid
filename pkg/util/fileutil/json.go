package fileutil

import (
	"encoding/json"
	"io"
)

func NewPrettyEncoder(w io.Writer) *json.Encoder {
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "  ")
	return encoder
}

//NewDecoder is just implemented for uniformed handling
func NewDecoder(r io.Reader) *json.Decoder {
	return json.NewDecoder(r)
}
