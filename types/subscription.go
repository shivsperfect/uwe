package types

import "time"

type Currency int

const (
	USD Currency = iota
	EUR
	GBP
)

func (c Currency) String() string {
	switch c {
	case USD:
		return "USD"
	case EUR:
		return "EUR"
	case GBP:
		return "GBP"
	default:
		return "INVALID"
	}
}

type Subscription struct {
	Amount     int       `json:"amount"`
	Currency   Currency  `json:"currency"`
	Period     int       `json:"period"`
	VAT        int       `json:"vat"`
	ExternalID string    `json:"external_id"`
	StartedAt  time.Time `json:"started_at"`
	CanceledAt time.Time `json:"canceled_at"`
}
