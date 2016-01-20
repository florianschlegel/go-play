package client

import (
	"log"
	"testing"
)

func TestCheckVat(t *testing.T) {
	soapClient := NewCheckVatPortType("http://ec.europa.eu/taxation_customs/vies/services/checkVatService", false, nil)
	request := &CheckVat{
		CountryCode: "DE",
		VatNumber:   "203071105",
	}
	checkVatResponse, err := soapClient.CheckVat(request)
	if err != nil {
		t.Fatal(err)
	}
	log.Println("response", checkVatResponse)
}
