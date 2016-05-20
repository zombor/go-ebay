package ebay

import (
	"encoding/xml"
	"fmt"
)

type ebayRequest struct {
	conf    EbayConf
	command Command
}

func (c ebayRequest) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	startElement := xml.StartElement{
		Name: xml.Name{
			Space: "urn:ebay:apis:eBLBaseComponents",
			Local: fmt.Sprintf("%sRequest", c.command.CallName()),
		},
	}

	err := e.EncodeToken(startElement)

	if err != nil {
		return err
	}

	type RequesterCredentials struct {
		EBayAuthToken string `xml:"eBayAuthToken"`
	}

	err = e.Encode(
		RequesterCredentials{
			EBayAuthToken: c.conf.AuthToken,
		},
	)

	if err != nil {
		return err
	}

	err = e.Encode(c.command.Body())

	if err != nil {
		return err
	}

	endElement := xml.EndElement{
		Name: xml.Name{
			Space: "urn:ebay:apis:eBLBaseComponents",
			Local: fmt.Sprintf("%sRequest", c.command.CallName()),
		},
	}

	err = e.EncodeToken(endElement)

	if err != nil {
		return err
	}

	return nil
}
