package ebay

import (
	"flag"
	"testing"

	"github.com/stretchr/testify/suite"
)

var devID, appID, certID, ruName, authToken string

func init() {
	flag.StringVar(&devID, "devid", "", "EBay sandbox DevID")
	flag.StringVar(&devID, "appid", "", "EBay sandbox AppID")
	flag.StringVar(&devID, "certid", "", "EBay sandbox CertID")
	flag.StringVar(&devID, "runame", "", "EBay sandbox RuName")
	flag.StringVar(&devID, "authtoken", "", "EBay sandbox AuthToken")
}

type EbayIntegrationTestSuite struct {
	suite.Suite
	ebayConf EbayConf
}

func (s *EbayIntegrationTestSuite) SetupSuite() {
	s.ebayConf = EbayConf{
		DevId:     devID,
		AppId:     appID,
		CertId:    certID,
		RuName:    ruName,
		AuthToken: authToken,
		SiteId:    0,
	}.Sandbox()
}

func TestEbayIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(EbayIntegrationTestSuite))
}
