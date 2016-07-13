package ebay

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

type EbayTestSuite struct {
	suite.Suite
	ebayConf EbayConf
}

func (s *EbayTestSuite) SetupSuite() {
	s.ebayConf = EbayConf{
		DevId:     "dev-id",
		AppId:     "app-id",
		CertId:    "cert-id",
		RuName:    "ru-name",
		AuthToken: "auth-token",
		SiteId:    1,
	}
}

func (s *EbayTestSuite) Test_do_SendsRequest() {
	var ebayCalled bool
	type testBody struct {
		Test string
	}

	c := funcEbayCommand{
		callName: "test-command",
		body:     testBody{"test"},
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.Equal("/ws/api.dll", r.URL.Path)
		s.Equal("dev-id", r.Header.Get("X-EBAY-API-DEV-NAME"))
		s.Equal("app-id", r.Header.Get("X-EBAY-API-APP-NAME"))
		s.Equal("cert-id", r.Header.Get("X-EBAY-API-CERT-NAME"))
		s.Equal("1", r.Header.Get("X-EBAY-API-SITEID"))
		s.Equal("text/xml", r.Header.Get("Content-Type"))

		body, err := ioutil.ReadAll(r.Body)
		s.NoError(err)

		s.Equal(fmt.Sprintf(`%s<%sRequest xmlns="urn:ebay:apis:eBLBaseComponents"><RequesterCredentials><eBayAuthToken>%s</eBayAuthToken></RequesterCredentials><testBody><Test>test</Test></testBody></%sRequest>`, "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n", c.callName, s.ebayConf.AuthToken, c.callName), string(body))

		ebayCalled = true

		fmt.Fprintf(w, `<VerifyAddFixedPriceItemResponse xmlns="urn:ebay:apis:eBLBaseComponents"><Timestamp>2016-01-24T20:35:54.398Z</Timestamp><Ack>Success</Ack></VerifyAddFixedPriceItemResponse>`)
	}))
	defer ts.Close()
	s.ebayConf.baseUrl = ts.URL

	_, err := s.ebayConf.RunCommand(c)

	s.NoError(err)
	s.True(ebayCalled)
}

func (s *EbayTestSuite) Test_NewSandbox() {
	c := s.ebayConf.Sandbox()

	s.Equal("https://api.sandbox.ebay.com", c.baseUrl)
}

func (s *EbayTestSuite) Test_NewProduction() {
	c := s.ebayConf.Production()

	s.Equal("https://api.ebay.com", c.baseUrl)
}

func (s *EbayTestSuite) Test_ReturnsErrorWhenXMLEncodingFails() {
	var ebayCalled bool
	type testBody struct {
		Test string `xml:",comment"`
	}

	c := funcEbayCommand{
		callName: "test-command",
		body:     testBody{"--foo"},
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ebayCalled = true
	}))
	defer ts.Close()
	s.ebayConf.baseUrl = ts.URL

	_, err := s.ebayConf.RunCommand(c)

	s.Require().Error(err)
	s.True(strings.HasPrefix(err.Error(), "xml: "))
	s.False(ebayCalled)
}

func (s *EbayTestSuite) Test_ReturnsErrorForEbayNon200() {
	var ebayCalled bool
	type testBody struct {
		Test string
	}

	c := funcEbayCommand{
		callName: "test-command",
		body:     testBody{"test"},
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)

		fmt.Fprintf(w, `<SomeBadError>Wat</SomeBadError>`)

		ebayCalled = true
	}))
	defer ts.Close()
	s.ebayConf.baseUrl = ts.URL

	_, err := s.ebayConf.RunCommand(c)

	s.True(ebayCalled)
	s.Error(err)
	if httpErr, ok := err.(httpError); s.True(ok) {
		s.Equal("400 - <SomeBadError>Wat</SomeBadError>", httpErr.Error())
	}
}

func (s *EbayTestSuite) Test_ReturnsErrorWhenParseResponseErrors() {
	parseErr := errors.New("boom")
	var ebayCalled bool

	type testBody struct {
		Test string
	}

	c := funcEbayCommand{
		callName: "test-command",
		body:     testBody{"test"},
		err:      parseErr,
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ebayCalled = true
	}))
	defer ts.Close()
	s.ebayConf.baseUrl = ts.URL

	_, err := s.ebayConf.RunCommand(c)
	s.Equal(parseErr, err)
}

func (s *EbayTestSuite) Test_ReturnsResponseErrorsWhenResponseParsesAsFailure() {
	parseErr := errors.New("boom")
	var ebayCalled bool

	type testBody struct {
		Test string
	}

	errs := []ebayResponseError{
		ebayResponseError{
			ShortMessage:        "",
			LongMessage:         "",
			ErrorCode:           1,
			SeverityCode:        "code",
			ErrorClassification: "classification",
		},
	}

	c := funcEbayCommand{
		callName: "test-command",
		body:     testBody{"test"},
		err:      parseErr,
		response: ebayResponse{
			Ack:    "Failure",
			Errors: errs,
		},
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ebayCalled = true
	}))
	defer ts.Close()
	s.ebayConf.baseUrl = ts.URL

	_, err := s.ebayConf.RunCommand(c)
	if ebayErr, ok := err.(ebayErrors); s.True(ok) {
		s.Equal(errs, ebayErr)
	}
}

func TestEbayTestSuite(t *testing.T) {
	suite.Run(t, new(EbayTestSuite))
}

type funcEbayCommand struct {
	callName string
	body     interface{}
	err      error
	response ebayResponse
}

func (f funcEbayCommand) CallName() string { return f.callName }

func (f funcEbayCommand) Body() interface{} {
	return f.body
}

func (f funcEbayCommand) ParseResponse([]byte) (EbayResponse, error) {
	return f.response, f.err
}
