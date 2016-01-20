package vatservice

import "encoding/xml"

type VatRequest struct {
	XMLName     xml.Name `xml:"urn:ec.europa.eu:taxud:vies:services:checkVat:types checkVat"`
	CountryCode string   `xml:"countryCode"`
	VatNumber   string   `xml:"vatNumber"`
}

type SOAPEnvelope struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`

	Body SOAPBody
}

type SOAPHeader struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Header"`

	Header interface{}
}

type SOAPBody struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`

	Fault   *SOAPFault  `xml:",omitempty"`
	Content interface{} `xml:",omitempty"`
}

type SOAPFault struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault"`

	Code   string `xml:"faultcode,omitempty"`
	String string `xml:"faultstring,omitempty"`
	Actor  string `xml:"faultactor,omitempty"`
	Detail string `xml:"detail,omitempty"`
}

type EnvelopeWithVatRequest struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Body    *struct {
		XMLName          xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`
		CheckVatResponse *struct {
			XMLName xml.Name `xml:"urn:ec.europa.eu:taxud:vies:services:checkVat:types checkVatResponse"`

			CountryCode string `xml:"countryCode,omitempty"`
			VatNumber   string `xml:"vatNumber,omitempty"`
			//RequestDate time.Time `xml:"requestDate,omitempty"`
			RequestDate string `xml:"requestDate,omitempty"`
			Valid       bool   `xml:"valid,omitempty"`
			Name        string `xml:"name,omitempty"`
			Address     string `xml:"address,omitempty"`
		} `xml:"checkVatResponse,omitempty"`
	}
}

func parseVatRequest(requestBytes []byte) (vatRequest VatRequest, err error) {
	vatRequest = VatRequest{}
	envelope := SOAPEnvelope{
		Body: SOAPBody{
			Content: vatRequest,
		},
	}
	err = xml.Unmarshal(requestBytes, &envelope)
	return
}
