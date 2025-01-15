//package main
//
//import (
//	"encoding/xml"
//	"fmt"
//)
//
//type Attribute struct {
//	XMLName xml.Name `xml:"saml:Attribute"`
//	Name    string   `xml:"Name,attr"`
//	Value   string   `xml:"saml:AttributeValue"`
//}
//
//type Assertion struct {
//	XMLName    xml.Name    `xml:"saml:Assertion"`
//	Attributes []Attribute `xml:"saml:AttributeStatement>saml:Attribute"`
//}
//
//func main() {
//	// 创建一个包含多个断言的文档
//	assertion := &Assertion{
//		Attributes: []Attribute{
//			{Name: "https://cloud.tencent.com/SAML/Attributes/Role", Value: "QCS::CAM::uin/12345678:roleName/role1"},
//			{Name: "https://cloud.tencent.com/SAML/Attributes/RoleSessionName", Value: "test"},
//		},
//	}
//
//	// 将文档转换为XML格式的字符串
//	output, err := xml.MarshalIndent(assertion, "", "  ")
//	if err != nil {
//		fmt.Printf("error: %v\n", err)
//	}
//
//	// 打印XML文档
//	fmt.Printf("%s\n", string(output))
//}

//package main
//
//import (
//	"encoding/xml"
//	"fmt"
//	"os"
//	"time"
//)
//
//type Response struct {
//	XMLName   xml.Name  `xml:"saml2p:Response"`
//	Saml2     string    `xml:"xmlns:saml2,attr"`
//	Saml2p    string    `xml:"xmlns:saml2p,attr"`
//	Ds        string    `xml:"xmlns:ds,attr"`
//	Issuer    string    `xml:"saml2:Issuer"`
//	Status    Status    `xml:"saml2p:Status"`
//	Assertion Assertion `xml:"saml2:Assertion"`
//}
//
//type Status struct {
//	XMLName xml.Name `xml:"saml2p:Status"`
//	// ...
//}
//
//type Assertion struct {
//	XMLName        xml.Name       `xml:"saml2:Assertion"`
//	Issuer         string         `xml:"saml2:Issuer"`
//	Signature      string         `xml:"ds:Signature"`
//	Subject        Subject        `xml:"saml2:Subject"`
//	Conditions     Conditions     `xml:"saml2:Conditions"`
//	Authentication Authentication `xml:"saml2:AuthnStatement"`
//}
//
////type Signature struct {
////	XMLName xml.Name `xml:"ds:Signature"`
////	Value   string   `xml:"ds:Signature"`
////}
//
//type Subject struct {
//	XMLName             xml.Name            `xml:"saml2:Subject"`
//	NameID              string              `xml:"saml2:NameID"`
//	SubjectConfirmation SubjectConfirmation `xml:"saml2:SubjectConfirmation"`
//}
//
//type SubjectConfirmation struct {
//	XMLName                 xml.Name                `xml:"saml2:SubjectConfirmation"`
//	SubjectConfirmationData SubjectConfirmationData `xml:"saml2:SubjectConfirmationData"`
//	// ...
//}
//type SubjectConfirmationData struct {
//	XMLName      xml.Name `xml:"saml2:SubjectConfirmationData"`
//	NotOnOrAfter string   `xml:"NotOnOrAfter,attr"`
//	Recipient    string   `xml:"Recipient,attr"`
//}
//
//type Conditions struct {
//	XMLName             xml.Name            `xml:"saml2:Conditions"`
//	AudienceRestriction AudienceRestriction `xml:"saml2:AudienceRestriction"`
//}
//
//type AudienceRestriction struct {
//	XMLName  xml.Name `xml:"saml2:AudienceRestriction"`
//	Audience string   `xml:"saml2:Audience"`
//}
//
//type Authentication struct {
//	XMLName xml.Name `xml:"saml2:AuthnStatement"`
//	// ...
//}
//
//func main() {
//	NameID := "12345678"
//	ExpirationTime := time.Now().Add(time.Hour * 2).Format("2006-01-02T15:04:05Z")
//	RecipientAddress := "https://signin.aliyun.com/saml/SSO"
//	accountID := "12345678"
//	Audience := fmt.Sprintf("https://signin.aliyun.com/%s/saml/SSO", accountID)
//	response := Response{
//		Saml2:  "urn:oasis:names:tc:SAML:2.0:assertion",
//		Saml2p: "urn:oasis:names:tc:SAML:2.0:protocol",
//		Ds:     "http://www.w3.org/2000/09/xmldsig#",
//		Issuer: "http://127.0.0.1:8000/sso/saml/login",
//		Status: Status{
//			// ...
//		},
//		Assertion: Assertion{
//			Issuer:    "http://127.0.0.1:8000/sso/saml/login",
//			Signature: "MIIDujCCAqKgAwIBAgIJAKtb15tbL/91MA0GCSqGSIb3DQEBCwUAMGsxCzAJBgNVBAYTAlhZMQwwCgYDVQQIDANYSVkxFTATBgNVBAcMDERlZmF1bHQgQ2l0eTEcMBoGA1UECgwTRGVmYXVsdCBDb21wYW55IEx0ZDEZMBcGCSqGSIb3DQEJARYKMTI0QHFxLmNvbTAeFw0yNDA1MjcwMzQ4MTBaFw0zNDA1MjUwMzQ4MTBaMGsxCzAJBgNVBAYTAlhZMQwwCgYDVQQIDANYSVkxFTATBgNVBAcMDERlZmF1bHQgQ2l0eTEcMBoGA1UECgwTRGVmYXVsdCBDb21wYW55IEx0ZDEZMBcGCSqGSIb3DQEJARYKMTI0QHFxLmNvbTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBANJsdxyh7TCX9lcTuV57B3D7wQ+nQ7cZSiW+pbwcti0gi/52KY+ebWP0DJGIWbz2z3OxhfkghhtK64Ex/TpdO3qzRAmgpsTBEJEegTDaU+4wJBqOPTsVf1/2gP+wjWZkYlC5Wmvx6juBBzcIlgAcDexzNWxNh+xBPy7agF7CNDYfBbvwLZv0cotNCmifL1vB4woxAUD62uKRBbCTlWWiZDccTk6c5ukgYdcUQSVshUpG4LNmR1/aFOGPC2fKmePilPVMYWDjfOdL9trmUuGB105jqI3IT9ulbu9P87K880YZZ+Wflc6B2bzfpjOpiu6DO6lPyGSpBGHPyezH9QQ2NkkCAwEAAaNhMF8wHQYDVR0OBBYEFG+7CWZYqOWEE4Lp6WPiDZ3/cWH5MB8GA1UdIwQYMBaAFG+7CWZYqOWEE4Lp6WPiDZ3/cWH5MA8GA1UdEQQIMAaHBAoBgx8wDAYDVR0TBAUwAwEB/zANBgkqhkiG9w0BAQsFAAOCAQEAGzygfrxxUM4MVL3yWOSn4qW+FepdVxZtdckXeKcb9BB33O9VH1hXP3o1J26Pqnu6aKIA8xsEInMon2Z23SCeyufreUzaK1AIcSlN9z0LsTHt72kAAajmLuWUUkV40am4/uB6H2kM750Q1YIE45j2acJ/Rp1V2S6KEnXv8it6wLItLDhs3LUVRvjcm6gGaAgRevt6L2TH9zy4KdsWTfsi2OCMPQ9KDAxcILyP9ZCZT61ek84uH0fAlkmqisLwyyMyRAJ1Xf3UHVyuAFkj0dT2p2zl9vFiOQrz6G5BOWWCGApZlzxvvJuivN+mXA34c9GZ60M/nIhocSXDSPEtDMGNgQ==",
//			Subject: Subject{
//				NameID: NameID,
//				SubjectConfirmation: SubjectConfirmation{
//					SubjectConfirmationData: SubjectConfirmationData{
//						NotOnOrAfter: ExpirationTime,
//						Recipient:    RecipientAddress,
//					},
//				},
//			},
//			Conditions: Conditions{
//				AudienceRestriction: AudienceRestriction{
//					Audience: Audience,
//				},
//			},
//			Authentication: Authentication{
//				// ...
//			},
//		},
//	}
//
//	xmlData, err := xml.MarshalIndent(response, "", "  ")
//	if err != nil {
//		fmt.Println("Error marshaling XML:", err)
//		return
//	}
//
//	os.Stdout.Write(xmlData)
//}

package main

import "fmt"

func main() {
	str := `<saml2p:Response xmlns:saml2="urn:oasis:names:tc:SAML:2.0:assertion" xmlns:saml2p="urn:oasis:names:tc:SAML:2.0:protocol" xmlns:ds="http://www.w3.org/2000/09/xmldsig#">`
	str += `<saml2:Issuer>http://127.0.0.1:8000/sso/saml/login</saml2:Issuer>`
	str += `<saml2p:Status></saml2p:Status>`
	str += `<saml2p:AuthnStatement>\n`
	str += `<saml2:Assertion>`
	str += `<saml2:Issuer>http://127.0.0.1:8000/sso/saml/login</saml2:Issuer>`
	str += `<ds:Signature>MIIDujCCAqKgAwIBAgIJAKtb15tbL/91MA0GCSqGSIb3DQEBCwUAMGsxCzAJBgNVBAYTAlhZMQwwCgYDVQQIDANYSVkxFTATBgNVBAcMDERlZmF1bHQgQ2l0eTEcMBoGA1UECgwTRGVmYXVsdCBDb21wYW55IEx0ZDEZMBcGCSqGSIb3DQEJARYKMTI0QHFxLmNvbTAeFw0yNDA1MjcwMzQ4MTBaFw0zNDA1MjUwMzQ4MTBaMGsxCzAJBgNVBAYTAlhZMQwwCgYDVQQIDANYSVkxFTATBgNVBAcMDERlZmF1bHQgQ2l0eTEcMBoGA1UECgwTRGVmYXVsdCBDb21wYW55IEx0ZDEZMBcGCSqGSIb3DQEJARYKMTI0QHFxLmNvbTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBANJsdxyh7TCX9lcTuV57B3D7wQ+nQ7cZSiW+pbwcti0gi/52KY+ebWP0DJGIWbz2z3OxhfkghhtK64Ex/TpdO3qzRAmgpsTBEJEegTDaU+4wJBqOPTsVf1/2gP+wjWZkYlC5Wmvx6juBBzcIlgAcDexzNWxNh+xBPy7agF7CNDYfBbvwLZv0cotNCmifL1vB4woxAUD62uKRBbCTlWWiZDccTk6c5ukgYdcUQSVshUpG4LNmR1/aFOGPC2fKmePilPVMYWDjfOdL9trmUuGB105jqI3IT9ulbu9P87K880YZZ+Wflc6B2bzfpjOpiu6DO6lPyGSpBGHPyezH9QQ2NkkCAwEAAaNhMF8wHQYDVR0OBBYEFG+7CWZYqOWEE4Lp6WPiDZ3/cWH5MB8GA1UdIwQYMBaAFG+7CWZYqOWEE4Lp6WPiDZ3/cWH5MA8GA1UdEQQIMAaHBAoBgx8wDAYDVR0TBAUwAwEB/zANBgkqhkiG9w0BAQsFAAOCAQEAGzygfrxxUM4MVL3yWOSn4qW+FepdVxZtdckXeKcb9BB33O9VH1hXP3o1J26Pqnu6aKIA8xsEInMon2Z23SCeyufreUzaK1AIcSlN9z0LsTHt72kAAajmLuWUUkV40am4/uB6H2kM750Q1YIE45j2acJ/Rp1V2S6KEnXv8it6wLItLDhs3LUVRvjcm6gGaAgRevt6L2TH9zy4KdsWTfsi2OCMPQ9KDAxcILyP9ZCZT61ek84uH0fAlkmqisLwyyMyRAJ1Xf3UHVyuAFkj0dT2p2zl9vFiOQrz6G5BOWWCGApZlzxvvJuivN+mXA34c9GZ60M/nIhocSXDSPEtDMGNgQ==</ds:Signature>`
	str += `<saml2:Subject>`
	str += `<saml2:NameID>xiy19425_admin</saml2:NameID>`
	str += `<saml2:SubjectConfirmation>`
	str += `<saml2:SubjectConfirmationData NotOnOrAfter="2024-07-25T17:36:07Z" Recipient="https://signin.aliyun.com/saml/SSO"></saml2:SubjectConfirmationData>`
	str += `</saml2:SubjectConfirmation>`
	str += `</saml2:Subject>`
	str += `<saml2:Conditions>`
	str += `<saml2:AudienceRestriction>`
	str += `<saml2:Audience>https://signin.aliyun.com/saml/SSO</saml2:Audience>`
	str += `</saml2:AudienceRestriction>`
	str += `</saml2:Conditions>`
	str += `<saml2:AuthnStatement></saml2:AuthnStatement>`
	str += `</saml2:Assertion>`
	str += `</saml2:Authn>`
	str += `</saml2p:Response>`
	fmt.Println(str)
}
