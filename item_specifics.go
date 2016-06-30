package ebay

type ItemSpecifics struct {
	NameValueList []NameValueList
}

type NameValueList struct {
	Name string
	Value []string
}
