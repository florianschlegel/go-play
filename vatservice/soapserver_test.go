package vatservice

import (
	"encoding/json"
	"encoding/xml"
	"log"
	"testing"
)

const (
	responseXML = `<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/"><soap:Body><checkVatResponse xmlns="urn:ec.europa.eu:taxud:vies:services:checkVat:types"><countryCode>DE</countryCode><vatNumber>203071105</vatNumber><requestDate>2016-01-19+01:00</requestDate><valid>true</valid><name>---</name><address>---</address></checkVatResponse></soap:Body></soap:Envelope>`
	//request  = `<Envelope xmlns="http://schemas.xmlsoap.org/soap/envelope/"><Body xmlns="http://schemas.xmlsoap.org/soap/envelope/"><checkVat xmlns="urn:ec.europa.eu:taxud:vies:services:checkVat:types"><countryCode>DE</countryCode><vatNumber>203071105</vatNumber></checkVat></Body></Envelope>`
	request    = `<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/"><soap:Body><checkVat xmlns="urn:ec.europa.eu:taxud:vies:services:checkVat:types"><countryCode>DE</countryCode><vatNumber>203071105</vatNumber></checkVat></soap:Body></soap:Envelope>`
	requestFoo = `<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/"><soap:Body><foo><bar>hello</bar></foo></soap:Body></soap:Envelope>`
)

func estStuff(t *testing.T) {
	response := &CheckVatResponse{}
	rawbody := []byte(responseXML)
	log.Println(string(rawbody))
	respEnvelope := new(SOAPEnvelope)
	respEnvelope.Body = SOAPBody{Content: response}
	err := xml.Unmarshal(rawbody, respEnvelope)
	t.Log(err, "vat number", response.VatNumber)
}

func testParseRequest(t *testing.T) {
	r, err := parseVatRequest([]byte(request))
	if err != nil {
		panic(err)
	}
	t.Log(request, err)
	jsonBytes, jsonErr := json.MarshalIndent(r, "", "   ")
	t.Log("json", jsonErr, string(jsonBytes))
	if r.VatNumber != "203071105" {
		t.Fatal("wrong vat number")
	}
}

type Foo struct {
	Bar string `xml:"bar"`
}

func TestParseRequestFoo(t *testing.T) {
	foo := &Foo{}
	err := LoadRequest([]byte(requestFoo), foo)
	if err != nil {
		panic(foo)
	}
	jsonBytes, jsonErr := json.MarshalIndent(foo, "", "   ")
	t.Log("json", jsonErr, string(jsonBytes))
}

/*
func TestParseRequest(t *testing.T) {
	soapEnvelope, err := parseRequest([]byte(response))
	t.Log(soapEnvelope, err)
	jsonBytes, jsonErr := json.MarshalIndent(soapEnvelope, "", "   ")
	t.Log("json", jsonErr, string(jsonBytes))

}
*/
