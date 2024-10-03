package models

type CurrencyInfo struct {
	Timestamp int64    `json:"timestamp"`
	Asks      []AskBid `json:"asks"`
	Bids      []AskBid `json:"bids"`
}

type AskBid struct {
	Price  string `json:"price"`
	Volume string `json:"volume"`
	Amount string `json:"amount"`
	Factor string `json:"factor"`
	Type   Type   `json:"type"`
}

type Type string
