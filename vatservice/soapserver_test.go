package vatservice

import (
	"encoding/json"
	"testing"
)

const request = `<Envelope xmlns="http://schemas.xmlsoap.org/soap/envelope/"><Body xmlns="http://schemas.xmlsoap.org/soap/envelope/"><checkVat xmlns="urn:ec.europa.eu:taxud:vies:services:checkVat:types"><countryCode>DE</countryCode><vatNumber>203071105</vatNumber></checkVat></Body></Envelope>`

func TestParseRequest(t *testing.T) {
	soapEnvelope, err := parseRequest([]byte(request))
	t.Log(soapEnvelope, err)
	jsonBytes, jsonErr := json.MarshalIndent(soapEnvelope, "", "   ")
	t.Log("json", jsonErr, string(jsonBytes))

}
