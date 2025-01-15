package saml

import (
	"bytes"
	"compress/flate"
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"io"
)

// AuthnRequest 定义saml request结构体
type AuthnRequest struct {
	XMLName                     xml.Name `xml:"urn:oasis:names:tc:SAML:2.0:protocol AuthnRequest"`
	AssertionConsumerServiceURL string   `xml:"AssertionConsumerServiceURL,attr"`
	Destination                 string   `xml:"Destination,attr"`
	ID                          string   `xml:"ID,attr"`
	IssueInstant                string   `xml:"IssueInstant,attr"`
	Version                     string   `xml:"Version,attr"`
	Issuer                      string   `xml:"urn:oasis:names:tc:SAML:2.0:assertion Issuer"`
}

func XmlRequest(request string) *string {
	// 解码base64
	decoded, err := base64.StdEncoding.DecodeString(request)
	if err != nil {
		fmt.Println("decode error:", err)
		return nil
	}
	// 解压
	reader := flate.NewReader(bytes.NewReader(decoded))
	uncompressedRequest, err := io.ReadAll(reader)
	if err != nil {
		fmt.Println("read error:", err)
		return nil
	}
	// 解析xml 获取requestId 并返回
	return getRequestId(uncompressedRequest)
}

func getRequestId(request []byte) (id *string) {
	var authnRequest AuthnRequest
	// 解析xml
	err := xml.Unmarshal(request, &authnRequest)
	if err != nil {
		fmt.Printf("error: %v", err)
		return nil
	}
	return &authnRequest.ID
}
