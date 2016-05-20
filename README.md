# go-ebay

This package runs api requests against eBay's XML Trading API.

Please note that this package's api will probably change in the future. Please use vendoring before depending on it.

## Usage

You'll need the proper eBay api credentials.

```go
conf := EbayConf{
  DevID: "dev-id",
  AppId: "app-id",
  CertId: "cert-id",
  RuName: "ru-name",
  AuthToken: "auth-token",
}

// Use sandbox
ebay := conf.Sandbox()

// Use production
ebay := conf.Production()

// Run an API call
response, err := ebay.RunCommand(
  GetItem{ItemID: "the-item-id"},
)

fmt.Println(response.ItemID)
fmt.Println(response.BuyItNowPrice)
```

## Implemented API Calls

The following api calls are currently implemented, with limited field support:

 - AddFixedPriceItem
 - VerifyAddFixedPriceItem
 - ReviseFixedPriceItem
 - GetItem

## Contributing

 - Add tests
 - Add code
 - Make sure `make test` passes
 - If possible, also add integration tests (`make integration-test`). You'll need an eBay developer account to run these
 - Send pull request against master
