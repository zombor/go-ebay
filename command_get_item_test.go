package ebay

import (
	"bytes"
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetItem_CallName(t *testing.T) {
	s := GetItem{}

	assert.Equal(t, "GetItem", s.CallName())
}

func Test_GetItem_Body(t *testing.T) {
	s := GetItem{
		ItemID: "item-id",
	}

	b := new(bytes.Buffer)
	xml.NewEncoder(b).Encode(s.Body())

	assert.Equal(t, "<ItemID>item-id</ItemID>", b.String())
}
