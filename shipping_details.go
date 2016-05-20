package ebay

type ShippingDetails struct {
	ShippingType                           string
	ShippingDiscountProfileID              string
	InternationalShippingDiscountProfileID string
	ShippingServiceOptions                 []ShippingServiceOption
	InternationalShippingServiceOption     []InternationalShippingServiceOption
}

type ShippingServiceOption struct {
	ShippingService               string
	ShippingServiceCost           float64
	ShippingServiceAdditionalCost float64
	FreeShipping                  bool
}

type InternationalShippingServiceOption struct {
	ShippingService               string
	ShippingServiceCost           float64
	ShippingServiceAdditionalCost float64
	ShipToLocation                []string
	ShippingServicePriority       int
}
