package ebay

import "encoding/xml"

type VerifyAddFixedPriceItem AddFixedPriceItem

func (c VerifyAddFixedPriceItem) CallName() string {
	return "VerifyAddFixedPriceItem"
}

func (c VerifyAddFixedPriceItem) Body() interface{} {
	type Item struct {
		VerifyAddFixedPriceItem
	}

	return Item{c}
}

func (c VerifyAddFixedPriceItem) ParseResponse(r []byte) (EbayResponse, error) {
	var xmlResponse AddFixedPriceItemResponse
	err := xml.Unmarshal(r, &xmlResponse)

	return xmlResponse, err
}
