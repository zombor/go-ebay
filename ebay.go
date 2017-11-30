package ebay

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

type EbayConf struct {
	baseUrl string

	DevId, AppId, CertId string
	RuName, AuthToken    string
	SiteId               int
	Logger               func(...interface{})
}

func (e EbayConf) Sandbox() EbayConf {
	e.baseUrl = "https://api.sandbox.ebay.com"

	return e
}

func (e EbayConf) Production() EbayConf {
	e.baseUrl = "https://api.ebay.com"

	return e
}

func (e EbayConf) RunCommand(c Command) (EbayResponse, error) {
	ec := ebayRequest{
		conf:    e,
		command: c,
	}

	body := new(bytes.Buffer)
	body.Write([]byte(xml.Header))
	err := xml.NewEncoder(body).Encode(ec)

	if err != nil {
		return ebayResponse{}, err
	}

	if e.Logger != nil {
		e.Logger(body.String())
	}

	req, _ := http.NewRequest(
		"POST",
		fmt.Sprintf("%s/ws/api.dll", e.baseUrl),
		body,
	)

	req.Header.Add("X-EBAY-API-DEV-NAME", e.DevId)
	req.Header.Add("X-EBAY-API-APP-NAME", e.AppId)
	req.Header.Add("X-EBAY-API-CERT-NAME", e.CertId)
	req.Header.Add("X-EBAY-API-CALL-NAME", c.CallName())
	req.Header.Add("X-EBAY-API-SITEID", strconv.Itoa(e.SiteId))
	req.Header.Add("X-EBAY-API-COMPATIBILITY-LEVEL", strconv.Itoa(837))
	req.Header.Add("Content-Type", "text/xml")

	client := &http.Client{}
	resp, err := client.Do(req)

	if urlErr, ok := err.(*url.Error); ok { // TODO: how to unit test this?
		return ebayResponse{}, urlErr
	} else if resp.StatusCode != 200 {
		httpErr := httpError{
			statusCode: resp.StatusCode,
		}
		httpErr.body, _ = ioutil.ReadAll(resp.Body)

		return ebayResponse{}, httpErr
	}

	bodyContents, _ := ioutil.ReadAll(resp.Body)

	if e.Logger != nil {
		e.Logger(string(bodyContents))
	}

	response, err := c.ParseResponse(bodyContents)

	if response.Failure() {
		return response, ebayErrors(response.ResponseErrors())
	}

	return response, err
}
