package ebay

import "encoding/xml"

type ReviseFixedPriceItem struct {
	ItemID                string
	StartPrice            string `xml:",omitempty"`
	ConditionID           uint   `xml:",omitempty"`
	Quantity              uint
	Title                 string           `xml:",omitempty"`
	Description           string           `xml:",omitempty"`
	PayPalEmailAddress    string           `xml:",omitempty"`
	PictureDetails        *PictureDetails  `xml:",omitempty"`
	ShippingDetails       *ShippingDetails `xml:",omitempty"`
	PrimaryCategory       *PrimaryCategory
	ReturnPolicy          *ReturnPolicy          `xml:",omitempty"`
	ProductListingDetails *ProductListingDetails `xml:",omitempty"`
	ItemSpecifics         []ItemSpecifics        `xml:",omitempty"`
}

func (c ReviseFixedPriceItem) CallName() string {
	return "ReviseFixedPriceItem"
}

func (c ReviseFixedPriceItem) Body() interface{} {
	type Item struct {
		ReviseFixedPriceItem
	}

	return Item{c}
}

func (c ReviseFixedPriceItem) ParseResponse(r []byte) (EbayResponse, error) {
	var xmlResponse ReviseFixedPriceItemResponse
	err := xml.Unmarshal(r, &xmlResponse)

	return xmlResponse, err
}

type ReviseFixedPriceItemResponse struct {
	ebayResponse
}

func (r ReviseFixedPriceItemResponse) ResponseErrors() ebayErrors {
	return r.ebayResponse.Errors
}
