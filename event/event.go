package event

import "encoding/json"

type TransactionFromSubscribedContract struct {
	Transaction json.RawMessage `json:"transaction"`
}
