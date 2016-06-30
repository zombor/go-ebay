package ebay

import (
	"bytes"
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_AddItem_CallName(t *testing.T) {
	s := AddItem{}

	assert.Equal(t, "AddItem", s.CallName())
}

func Test_AddItem_Body(t *testing.T) {
	s := AddItem{
		Currency:           "USD",
		Country:            "US",
		DispatchTimeMax:    1,
		ConditionID:        1,
		Title:              "title",
		Description:        "description",
		StartPrice:         "1.02",
		BuyItNowPrice:      "2.03",
		ListingType:        "type",
		Quantity:           1,
		PaymentMethods:     "paypal",
		PayPalEmailAddress: "email",
		ListingDuration:    "GTC",
		ShippingDetails: &ShippingDetails{
			ShippingType:                           "type",
			ShippingDiscountProfileID:              "profile id",
			InternationalShippingDiscountProfileID: "i profile id",
			ShippingServiceOptions: []ShippingServiceOption{
				ShippingServiceOption{
					ShippingService:               "service",
					ShippingServiceCost:           1.23,
					ShippingServiceAdditionalCost: 2.34,
					FreeShipping:                  false,
				},
			},
			InternationalShippingServiceOption: []InternationalShippingServiceOption{
				InternationalShippingServiceOption{
					ShippingService:               "service",
					ShippingServiceCost:           1.23,
					ShippingServiceAdditionalCost: 2.34,
					ShipToLocation:                []string{"ship", "to"},
					ShippingServicePriority:       1,
				},
			},
		},
		PrimaryCategory: &PrimaryCategory{"1"},
		Storefront:      &Storefront{"categoryid"},
		PostalCode:      "60657",
		ReturnPolicy: &ReturnPolicy{
			"accepted", "accepted option", "returns within option", "refund option",
		},
		PictureDetails: &PictureDetails{"url"},
		ProductListingDetails: &ProductListingDetails{
			UPC:      "upc",
			BrandMPN: BrandMPN{"foo", "bar"},
		},
		ItemSpecifics: []ItemSpecifics{
			ItemSpecifics{
				[]NameValueList{
					NameValueList{
						Name:  "name",
						Value: []string{"value"},
					},
				},
			},
		},
	}

	b := new(bytes.Buffer)
	xml.NewEncoder(b).Encode(s.Body())

	assert.Equal(t, "<Item><Currency>USD</Currency><Country>US</Country><DispatchTimeMax>1</DispatchTimeMax><ConditionID>1</ConditionID><Title>title</Title><Description>description</Description><StartPrice>1.02</StartPrice><BuyItNowPrice>2.03</BuyItNowPrice><ListingType>type</ListingType><Quantity>1</Quantity><PaymentMethods>paypal</PaymentMethods><PayPalEmailAddress>email</PayPalEmailAddress><ListingDuration>GTC</ListingDuration><ShippingDetails><ShippingType>type</ShippingType><ShippingDiscountProfileID>profile id</ShippingDiscountProfileID><InternationalShippingDiscountProfileID>i profile id</InternationalShippingDiscountProfileID><ShippingServiceOptions><ShippingService>service</ShippingService><ShippingServiceCost>1.23</ShippingServiceCost><ShippingServiceAdditionalCost>2.34</ShippingServiceAdditionalCost><FreeShipping>false</FreeShipping></ShippingServiceOptions><InternationalShippingServiceOption><ShippingService>service</ShippingService><ShippingServiceCost>1.23</ShippingServiceCost><ShippingServiceAdditionalCost>2.34</ShippingServiceAdditionalCost><ShipToLocation>ship</ShipToLocation><ShipToLocation>to</ShipToLocation><ShippingServicePriority>1</ShippingServicePriority></InternationalShippingServiceOption></ShippingDetails><PrimaryCategory><CategoryID>1</CategoryID></PrimaryCategory><Storefront><StoreCategoryID>categoryid</StoreCategoryID></Storefront><PostalCode>60657</PostalCode><ReturnPolicy><ReturnsAccepted>accepted</ReturnsAccepted><ReturnsAcceptedOption>accepted option</ReturnsAcceptedOption><ReturnsWithinOption>returns within option</ReturnsWithinOption><RefundOption>refund option</RefundOption></ReturnPolicy><PictureDetails><PictureURL>url</PictureURL></PictureDetails><ProductListingDetails><UPC>upc</UPC><BrandMPN><Brand>foo</Brand><MPN>bar</MPN></BrandMPN></ProductListingDetails><ItemSpecifics><NameValueList><Name>name</Name><Value>value</Value></NameValueList></ItemSpecifics></Item>", b.String())
}

func Test_AddItem_MissingFieldsBody(t *testing.T) {
	s := AddItem{
		Currency:        "USD",
		Country:         "US",
		StartPrice:      "1.23",
		ListingDuration: "GTC",
		PrimaryCategory: &PrimaryCategory{"1"},
	}

	b := new(bytes.Buffer)
	err := xml.NewEncoder(b).Encode(s.Body())

	assert.NoError(t, err)
	assert.Equal(t, "<Item><Currency>USD</Currency><Country>US</Country><StartPrice>1.23</StartPrice><ListingDuration>GTC</ListingDuration><PrimaryCategory><CategoryID>1</CategoryID></PrimaryCategory></Item>", b.String())
}
