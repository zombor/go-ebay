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
  Logger: logrus.Debug, // Optional
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

If you'd like to help out, either completing these api calls, or adding new ones would be great!

## Logging

You can log both the request and response XML if you pass a logger into the configuration struct.

Pass a function that satisfies the `func(...interface{})` signature. Both the stdlib logger and common loggers like logrus satisfy this function signature.

You can even use logrus structured logs to log with something like a request uuid:

```go
conf := EbayConf{
  DevID: "dev-id",
  AppId: "app-id",
  CertId: "cert-id",
  RuName: "ru-name",
  AuthToken: "auth-token",
  Logger: logrus.WithFields(logrus.Fields{
    "request_uuid": "something",
  }).Debug,
}

```

## Contributing

 - Add tests
 - Add code
 - Make sure `make test` passes
 - If possible, also add integration tests (`make integration-test`). You'll need an eBay developer account to run these
 - Send pull request against master
