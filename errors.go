package ebay

import (
	"fmt"
	"strings"
)

type ebayErrors []ebayResponseError

func (err ebayErrors) Error() string {
	var errors []string

	for _, e := range err {
		errors = append(errors, e.LongMessage)
	}

	return strings.Join(errors, ",")
}

func (errs ebayErrors) RevisionError() bool {
	for _, err := range errs {
		if err.ErrorCode == 10039 || err.ErrorCode == 10029 || err.ErrorCode == 21916916 {
			return true
		}
	}

	return false
}

func (errs ebayErrors) ListingEnded() bool {
	for _, err := range errs {
		if err.ErrorCode == 291 || err.ErrorCode == 17 {
			return true
		}
	}

	return false
}

func (errs ebayErrors) ListingDeleted() bool {
	for _, err := range errs {
		if err.ErrorCode == 17 {
			return true
		}
	}

	return false
}

type httpError struct {
	statusCode int
	body       []byte
}

func (err httpError) Error() string {
	return fmt.Sprintf("%d - %s", err.statusCode, err.body)
}
