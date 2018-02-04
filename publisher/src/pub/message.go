package pub

import "encoding/json"

//Content
type Content struct {
	Header string `json:"header"`
	Body   string `json:"body"`
}

//Message
type Message struct {
	From    string  `json:"from"`
	Content Content `json:"content"`
}

//JSON Convert Message to JSON
func (m Message) JSON() ([]byte, error) {
	return json.Marshal(m)
}
