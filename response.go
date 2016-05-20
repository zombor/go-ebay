package ebay

import "time"

type EbayResponse interface {
	Failure() bool
	ResponseErrors() ebayErrors
}

type ebayResponse struct {
	Timestamp time.Time
	Ack       string
	Errors    []ebayResponseError
}

func (r ebayResponse) Failure() bool {
	return r.Ack == "Failure"
}

func (r ebayResponse) ResponseErrors() ebayErrors {
	return r.Errors
}

type ebayResponseError struct {
	ShortMessage        string
	LongMessage         string
	ErrorCode           int
	SeverityCode        string
	ErrorClassification string
}
