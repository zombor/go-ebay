// +build integration

package ebay

import (
	"fmt"
	"time"
)

func (s *EbayIntegrationTestSuite) Test_VerifyAddFixedPriceItem() {
	now := time.Now().String()

	c := VerifyAddFixedPriceItem{
		Currency:       "USD",
		Country:        "US",
		ListingType:    "FixedPriceItem",
		Quantity:       10,
		PaymentMethods: "PayPal",
		ShippingDetails: &ShippingDetails{
			ShippingType: "Flat",
			ShippingServiceOptions: []ShippingServiceOption{
				ShippingServiceOption{
					ShippingService:               "USPSFirstClass",
					ShippingServiceCost:           1.00,
					ShippingServiceAdditionalCost: 0.40,
				},
			},
		},
		PictureDetails: &PictureDetails{
			PictureURL: "http://example.com/foo.png",
		},
		ListingDuration: "GTC",
		PrimaryCategory: &PrimaryCategory{
			CategoryID: "1234",
		},
		ReturnPolicy: &ReturnPolicy{
			ReturnsAccepted:          "ReturnsAccepted",
			ReturnsAcceptedOption:    "ReturnsAccepted",
			ReturnsWithinOption:      "Days_14",
			RefundOption:             "Days_14",
			ShippingCostPaidByOption: "Buyer",
		},
		PostalCode:         "60657",
		StartPrice:         "10.12",
		Title:              "The item title - " + now,
		Description:        "The item description - " + now,
		PayPalEmailAddress: "danabock@gmail.com",
		DispatchTimeMax:    1,
	}

	resp, err := s.ebayConf.RunCommand(c)

	s.NoError(err, fmt.Sprintf("%#v", resp))
}
