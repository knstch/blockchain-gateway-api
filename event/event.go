package event

type TransactionFromSubscribedContract struct {
	Network string `json:"string"`
	Hash    string `json:"hash"`

	From string `json:"from"`
	To   string `json:"to,omitempty"`

	GasPrice             string `json:"gas_price,omitempty"`
	MaxFeePerGas         string `json:"max_fee_per_gas,omitempty"`
	MaxPriorityFeePerGas string `json:"max_priority_fee_per_gas,omitempty"`

	Value string `json:"value"`

	Data []byte `json:"data"`
}
