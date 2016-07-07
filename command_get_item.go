package ebay

import "encoding/xml"

type GetItem struct {
	ItemID string
}

func (c GetItem) CallName() string {
	return "GetItem"
}

func (c GetItem) Body() interface{} {
	type ItemID struct {
		ItemID string `xml:",innerxml"`
	}

	return ItemID{c.ItemID}
}

func (c GetItem) ParseResponse(r []byte) (EbayResponse, error) {
	var xmlResponse GetItemResponse
	err := xml.Unmarshal(r, &xmlResponse)

	return xmlResponse, err
}

type GetItemResponse struct {
	ebayResponse

	Item struct {
		ItemID        string
		Quantity      int64
		SellingStatus struct {
			ListingStatus string
			QuantitySold  int64
		}
	}
}

func (r GetItemResponse) ResponseErrors() ebayErrors {
	return r.ebayResponse.Errors
}
