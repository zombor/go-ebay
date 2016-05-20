test:
	go test .

integration-test:
	go test --tags integration -devid=${EBAY_DEVID} -appid=${EBAY_APPID} -certid=${EBAY_CERTID} -runame=${EBAY_RUNAME} -authtoken=${EBAY_AUTHTOKEN}
