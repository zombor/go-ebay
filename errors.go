package ebay

import "strings"

type ebayErrors []ebayResponseError

func (err ebayErrors) Error() string {
	var errors []string

	for _, e := range err {
		errors = append(errors, e.LongMessage)
	}

	return strings.Join(errors, ",")
}
