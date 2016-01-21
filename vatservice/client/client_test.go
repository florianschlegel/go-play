package client

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

const request = `<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/"><soap:Body><checkVat xmlns="urn:ec.europa.eu:taxud:vies:services:checkVat:types"><countryCode>DE</countryCode><vatNumber>203071105</vatNumber></checkVat></soap:Body></soap:Envelope>`

//const request = ""

func TestCheckVat(t *testing.T) {
	//soapClient := NewCheckVatPortType("http://ec.europa.eu/taxation_customs/vies/services/checkVatService", false, nil)
	soapClient := NewCheckVatPortType("http://127.0.0.1:8080/", false, nil)
	response, err := soapClient.CheckVat(&CheckVat{
		CountryCode: "DE",
		VatNumber:   "789789789",
	})
	t.Log(response, err)
	// envelope := SOAPEnvelope{}
	//
	// envelope.Body.Content = request
	// buffer := new(bytes.Buffer)
	// encoder := xml.NewEncoder(buffer)
	//
	// if err := encoder.Encode(envelope.Body.Content); err != nil {
	// }
	//
	// if err := encoder.Flush(); err != nil {
	// }
	//
	// fmt.Println(string(buffer.Bytes()))

	//postRequest([]byte(request))

}

func postRequest(data []byte) error {
	url := "http://127.0.0.1:8080"

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/xml")
	req.Header.Set("SOAPAction", "operationCheckVat")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
	return err
}
