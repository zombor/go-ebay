package ebay

import "encoding/xml"

type GetItem struct {
	ItemID string
}

func (c GetItem) CallName() string {
	return "GetItem"
}

func (c GetItem) Body() interface{} {
	type Item struct {
		GetItem
	}

	return Item{c}
}

func (c GetItem) ParseResponse(r []byte) (EbayResponse, error) {
	var xmlResponse GetItemResponse
	err := xml.Unmarshal(r, &xmlResponse)

	return xmlResponse, err
}

type GetItemResponse struct {
	ebayResponse

	ItemID        string
	BuyItNowPrice string
}

func (r GetItemResponse) ResponseErrors() ebayErrors {
	return r.ebayResponse.Errors
}
