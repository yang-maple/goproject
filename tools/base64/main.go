package main

import (
	"bytes"
	"compress/flate"
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"io"
)

type AuthnRequest struct {
	XMLName                     xml.Name `xml:"urn:oasis:names:tc:SAML:2.0:protocol AuthnRequest"`
	AssertionConsumerServiceURL string   `xml:"AssertionConsumerServiceURL,attr"`
	Destination                 string   `xml:"Destination,attr"`
	ID                          string   `xml:"ID,attr"`
	IssueInstant                string   `xml:"IssueInstant,attr"`
	Version                     string   `xml:"Version,attr"`
	Issuer                      string   `xml:"urn:oasis:names:tc:SAML:2.0:assertion Issuer"`
}

// AuthnRequest 定义saml request结构体

func main() {
	str := "fZJRb4IwFIXf9ytI34FSFLARjJsxM3GRCO5hb12pWAOt4xaz/fuh6OYeZvrU5J5zbr5zx5PPurKOogGpVYw8ByNLKK4LqcoYbfK5HaFJ8jAGVlfkQKet2am1+GgFGGsKIBrT6Z60grYWTSaao+Ris17GaGfMAajrgiyVVA6r5FerHK5r92TlZtkKWbPORSpmztEnQTfvkdDB3fNohDF2AXQvqHQpFbLmuuHivEWMtqwCgazFLEaM8IDs+XZfBmHBimBQ7gY4kmERertR6XdDkDIAeRS/MoBWLBQYpkyMCCYDG4c2GeY4oEOfDgdO5PtvyEobbTTX1aNUPZS2UVQzkEAVqwVQw2k2fVlS4mD63g8Bfc7z1E5XWY6s1ytccoLb4VZAe5z3vQ6XYJT09Ol54+bW4b4Bu/aDkv/b8IJhgEc4JDiIIn80+qln7N6mJpfv3xNIvgE="
	//amlRequest, _ := url.QueryUnescape(str)
	decoded, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		fmt.Println("decode error:", err)
		return
	}
	reader := flate.NewReader(bytes.NewReader(decoded))
	uncompressedRequest, err := io.ReadAll(reader)
	fmt.Println(string(uncompressedRequest))

	//fmt.Println(string(decoded))
	var authnRequest AuthnRequest
	//
	err = xml.Unmarshal(uncompressedRequest, &authnRequest)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	fmt.Printf("AuthnRequest: %+v\n", authnRequest.Destination)
}
