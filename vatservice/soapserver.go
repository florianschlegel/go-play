package vatservice

import "encoding/xml"

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

type CheckVatResponse struct {
	XMLName xml.Name `xml:"urn:ec.europa.eu:taxud:vies:services:checkVat:types checkVatResponse"`

	CountryCode string `xml:"countryCode,omitempty"`
	VatNumber   string `xml:"vatNumber,omitempty"`
	//RequestDate time.Time `xml:"requestDate,omitempty"`
	RequestDate string `xml:"requestDate,omitempty"`
	Valid       bool   `xml:"valid,omitempty"`
	Name        string `xml:"name,omitempty"`
	Address     string `xml:"address,omitempty"`
}

type VatRequest struct {
	XMLName     xml.Name `xml:"urn:ec.europa.eu:taxud:vies:services:checkVat:types checkVat"`
	CountryCode string   `xml:"countryCode"`
	VatNumber   string   `xml:"vatNumber"`
}

type EnvelopeWithVatResponse struct {
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

type EnvelopeWithVatRequest struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Body    *struct {
		XMLName         xml.Name    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`
		CheckVatRequest *VatRequest `xml:"checkVat,omitempty"`
	}
}

func (b *SOAPBody) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if b.Content == nil {
		return xml.UnmarshalError("Content must be a pointer to a struct")
	}

	var (
		token    xml.Token
		err      error
		consumed bool
	)
Loop:
	for {
		if token, err = d.Token(); err != nil {
			return err
		}

		if token == nil {
			break
		}

		switch se := token.(type) {
		case xml.StartElement:
			if consumed {
				return xml.UnmarshalError("Found multiple elements inside SOAP body; not wrapped-document/literal WS-I compliant")
			} else if se.Name.Space == "http://schemas.xmlsoap.org/soap/envelope/" && se.Name.Local == "Fault" {
				b.Fault = &SOAPFault{}
				b.Content = nil

				err = d.DecodeElement(b.Fault, &se)
				if err != nil {
					return err
				}

				consumed = true
			} else {
				if err = d.DecodeElement(b.Content, &se); err != nil {
					return err
				}

				consumed = true
			}
		case xml.EndElement:
			break Loop
		}
	}

	return nil
}

func (f *SOAPFault) Error() string {
	return f.String
}

func LoadRequest(requestBytes []byte, request interface{}) error {
	envelope := &SOAPEnvelope{
		Body: SOAPBody{
			Content: request,
		},
	}
	return xml.Unmarshal(requestBytes, envelope)
}

func parseVatRequest(requestBytes []byte) (vatRequest *VatRequest, err error) {

	vatRequest = &VatRequest{}
	/*
		envelope := SOAPEnvelope{
			Body: SOAPBody{
				Content: vatRequest,
			},
		}
	*/
	envelope := new(SOAPEnvelope)
	envelope.Body = SOAPBody{Content: vatRequest}
	err = xml.Unmarshal(requestBytes, envelope)
	return
	/*
		envelope := &EnvelopeWithVatRequest{}
		err = xml.Unmarshal(requestBytes, &envelope)
		return envelope.Body.CheckVatRequest, err
	*/
}
