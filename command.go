package ebay

type Command interface {
	Body() interface{}
	CallName() string
	ParseResponse([]byte) (EbayResponse, error)
}
