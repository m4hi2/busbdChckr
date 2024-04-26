package busInformation

import (
	"bytes"
	"encoding/json"
	"io"
)

func BodyBuffer(body interface{}) (io.Reader, error) {
	jsonStr, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	return bytes.NewBuffer(jsonStr), nil
}

func BuildUrl(path string) string {
	return BDTicketHost + BDTicketPort + path
}
