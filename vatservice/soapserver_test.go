package vatservice

import (
	"encoding/json"
	"testing"
)

const (
	response = `<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/"><soap:Body><checkVatResponse xmlns="urn:ec.europa.eu:taxud:vies:services:checkVat:types"><countryCode>DE</countryCode><vatNumber>203071105</vatNumber><requestDate>2016-01-19+01:00</requestDate><valid>true</valid><name>---</name><address>---</address></checkVatResponse></soap:Body></soap:Envelope>`
	request  = `<Envelope xmlns="http://schemas.xmlsoap.org/soap/envelope/"><Body xmlns="http://schemas.xmlsoap.org/soap/envelope/"><checkVat xmlns="urn:ec.europa.eu:taxud:vies:services:checkVat:types"><countryCode>DE</countryCode><vatNumber>203071105</vatNumber></checkVat></Body></Envelope>`
)

func TestParseRequest(t *testing.T) {
	request, err := parseVatRequest([]byte(request))
	if err != nil {
		panic(err)
	}
	t.Log(request, err)
	jsonBytes, jsonErr := json.MarshalIndent(request, "", "   ")
	t.Log("json", jsonErr, string(jsonBytes))
	if request.VatNumber != "203071105" {
		t.Fatal("wrong vat number")
	}
}

/*
func TestParseRequest(t *testing.T) {
	soapEnvelope, err := parseRequest([]byte(response))
	t.Log(soapEnvelope, err)
	jsonBytes, jsonErr := json.MarshalIndent(soapEnvelope, "", "   ")
	t.Log("json", jsonErr, string(jsonBytes))

}
*/
