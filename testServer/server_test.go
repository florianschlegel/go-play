package testserver

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"testing"
)

type Person struct {
	Name    string   `json:"name" xml:"name"`
	Age     int      `json:"age" xml:"age"`
	Address *Address `json:"address" xml:"address"`
}

type Address struct {
	Street   string `json:"street"`
	City     string `json:"city" xml:",comment"`
	Addition string `json:"addition,omitempty" xml:"addition,attr"`
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

func TestServer(t *testing.T) {

	pers := &Person{
		Name: "Nicola",
		Age:  25,
		Address: &Address{
			Street: "Holzweg 12",
			City:   "München",
		},
	}

	envelope := SOAPEnvelope{}

	envelope.Body.Content = pers
	buffer := new(bytes.Buffer)
	encoder := xml.NewEncoder(buffer)

	if err := encoder.Encode(envelope); err != nil {
	}

	if err := encoder.Flush(); err != nil {
	}

	fmt.Println("BUFFER:  ", buffer)

	CallServer(buffer.Bytes())

	//xmlBytes, xmlErr := xml.Marshal(pers)

	/*
		foo := &Foo{}
		err := LoadRequest([]byte(requestFoo), foo)
		if err != nil {
			panic(foo)
		}
		jsonBytes, jsonErr := json.MarshalIndent(foo, "", "   ")
		t.Log("json", jsonErr, string(jsonBytes))
	*/
}

func CallServer(requestBytes []byte) {
	envelope := &SOAPEnvelope{
		Body: SOAPBody{
			Content: &Person{},
		},
	}

	xml.Unmarshal([]byte(requestBytes), envelope)

	fmt.Println("Envelope:  ", envelope)
}
