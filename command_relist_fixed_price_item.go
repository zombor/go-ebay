package ebay

import "encoding/xml"

type RelistFixedPriceItem struct {
	ItemID                string
	StartPrice            string                 `xml:",omitempty"`
	ConditionID           uint                   `xml:",omitempty"`
	Quantity              uint                   `xml:",omitempty"`
	Title                 string                 `xml:",omitempty"`
	Description           string                 `xml:",omitempty"`
	PayPalEmailAddress    string                 `xml:",omitempty"`
	PictureDetails        *PictureDetails        `xml:",omitempty"`
	ShippingDetails       *ShippingDetails       `xml:",omitempty"`
	ProductListingDetails *ProductListingDetails `xml:",omitempty"`
	ItemSpecifics         []ItemSpecifics        `xml:",omitempty"`
}

func (c RelistFixedPriceItem) CallName() string {
	return "RelistFixedPriceItem"
}

func (c RelistFixedPriceItem) Body() interface{} {
	type Item struct {
		RelistFixedPriceItem
	}

	return Item{c}
}

func (c RelistFixedPriceItem) ParseResponse(r []byte) (EbayResponse, error) {
	var xmlResponse RelistFixedPriceItemResponse
	err := xml.Unmarshal(r, &xmlResponse)

	return xmlResponse, err
}

type RelistFixedPriceItemResponse struct {
	ebayResponse
}

func (r RelistFixedPriceItemResponse) ResponseErrors() ebayErrors {
	return r.ebayResponse.Errors
}
