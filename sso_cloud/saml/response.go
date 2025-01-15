package saml

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"encoding/xml"
	"fmt"
	"github.com/google/uuid"
	"os"
	"time"
)

type Response struct {
	XMLName      xml.Name `xml:"saml2p:Response"`
	Saml2p       string   `xml:"xmlns:saml2p,attr"`
	Destination  string   `xml:"Destination,attr"`
	InResponseTo string   `xml:"InResponseTo,attr"`
	ID           string   `xml:"ID,attr"`
	Version      string   `xml:"Version,attr"`
	IssueInstant string   `xml:"IssueInstant,attr"`
	//Ds           string    `xml:"xmlns:ds,attr"`
	Issuer    Issuer    `xml:"saml2:Issuer"`
	Status    Status    `xml:"saml2p:Status"`
	Assertion Assertion `xml:"saml2:Assertion"`
}

type Issuer struct {
	XMLName xml.Name `xml:"saml2:Issuer"`
	Value   string   `xml:",chardata"`
	Saml2   string   `xml:"xmlns:saml2,attr"`
}

type Status struct {
	XMLName    xml.Name   `xml:"saml2p:Status"`
	StatusCode StatusCode `xml:"saml2p:StatusCode"`
}
type StatusCode struct {
	XMLName xml.Name `xml:"saml2p:StatusCode"`
	Value   string   `xml:"Value,attr"`
}

type Assertion struct {
	XMLName        xml.Name       `xml:"saml2:Assertion"`
	Saml2          string         `xml:"xmlns:saml2,attr"`
	ID             string         `xml:"ID,attr"`
	Version        string         `xml:"Version,attr"`
	IssueInstant   string         `xml:"IssueInstant,attr"`
	Issuer         string         `xml:"saml2:Issuer"`
	Signature      Signature      `xml:"ds:Signature"`
	Subject        Subject        `xml:"saml2:Subject"`
	Conditions     Conditions     `xml:"saml2:Conditions"`
	Authentication Authentication `xml:"saml2:AuthnStatement"`
}

type Signature struct {
	XMLName        xml.Name   `xml:"ds:Signature"`
	Id             string     `xml:"Id,attr"`
	Ds             string     `xml:"xmlns:ds,attr"`
	Value          string     `xml:",chardata"`
	SignedInfo     SignedInfo `xml:"ds:SignedInfo"`
	SignatureValue string     `xml:"ds:SignatureValue"`
}

type SignedInfo struct {
	XMLName                xml.Name               `xml:"ds:SignedInfo"`
	CanonicalizationMethod CanonicalizationMethod `xml:"ds:CanonicalizationMethod"`
	SignatureMethod        SignatureMethod        `xml:"ds:SignatureMethod"`
	Reference              Reference              `xml:"ds:Reference"`
}

type CanonicalizationMethod struct {
	XMLName   xml.Name `xml:"ds:CanonicalizationMethod"`
	Algorithm string   `xml:"Algorithm,attr"`
}

type SignatureMethod struct {
	XMLName   xml.Name `xml:"ds:SignatureMethod"`
	Algorithm string   `xml:"Algorithm,attr"`
}

type Reference struct {
	XMLName      xml.Name     `xml:"ds:Reference"`
	URI          string       `xml:"URI,attr"`
	Transforms   Transforms   `xml:"ds:Transforms"`
	DigestMethod DigestMethod `xml:"ds:DigestMethod"`
	DigestValue  string       `xml:"ds:DigestValue"`
}

type Transforms struct {
	XMLName   xml.Name  `xml:"ds:Transforms"`
	Transform Transform `xml:"ds:Transform"`
}
type Transform struct {
	XMLName   xml.Name `xml:"ds:Transform"`
	Algorithm string   `xml:"Algorithm,attr"`
}

type DigestMethod struct {
	XMLName   xml.Name `xml:"ds:DigestMethod"`
	Algorithm string   `xml:"Algorithm,attr"`
}

type Subject struct {
	XMLName             xml.Name            `xml:"saml2:Subject"`
	NameID              NameID              `xml:"saml2:NameID"`
	SubjectConfirmation SubjectConfirmation `xml:"saml2:SubjectConfirmation"`
}

type NameID struct {
	XMLName xml.Name `xml:"saml2:NameID"`
	Value   string   `xml:",chardata"`
	Format  string   `xml:"Format,attr"`
}

type SubjectConfirmation struct {
	XMLName                 xml.Name                `xml:"saml2:SubjectConfirmation"`
	Method                  string                  `xml:"Method,attr"`
	SubjectConfirmationData SubjectConfirmationData `xml:"saml2:SubjectConfirmationData"`
}
type SubjectConfirmationData struct {
	XMLName      xml.Name `xml:"saml2:SubjectConfirmationData"`
	InResponseTo string   `xml:"InResponseTo,attr"`
	NotOnOrAfter string   `xml:"NotOnOrAfter,attr"`
	Recipient    string   `xml:"Recipient,attr"`
}

type Conditions struct {
	XMLName             xml.Name            `xml:"saml2:Conditions"`
	NotBefore           string              `xml:"NotBefore,attr"`
	NotOnOrAfter        string              `xml:"NotOnOrAfter,attr"`
	AudienceRestriction AudienceRestriction `xml:"saml2:AudienceRestriction"`
}

type AudienceRestriction struct {
	XMLName  xml.Name `xml:"saml2:AudienceRestriction"`
	Audience string   `xml:"saml2:Audience"`
}

type Authentication struct {
	XMLName      xml.Name     `xml:"saml2:AuthnStatement"`
	AuthnInstant string       `xml:"AuthnInstant,attr"`
	SessionIndex string       `xml:"SessionIndex,attr"`
	AuthnContext AuthnContext `xml:"saml2:AuthnContext"`
	// ...
}

type AuthnContext struct {
	XMLName              xml.Name `xml:"saml2:AuthnContext"`
	AuthnContextClassRef string   `xml:"saml2:AuthnContextClassRef"`
}

const (
	Destination = "https://signin.aliyun.com/saml/SSO"
	IssUerURL   = "http://127.0.0.1:8000/sso/saml/login"
	Version     = "2.0"
)

func XmlResponse(nameId string, ResponseTo string, accountID string) *string {
	//NameIDs := nameId
	ExpirationTime := time.Now().Add(time.Hour * -6).Format("2006-01-02T15:04:05Z")
	RecipientAddress := "https://signin.aliyun.com/saml/SSO"
	Audience := fmt.Sprintf("https://signin.aliyun.com/%s/saml/SSO", accountID)
	RandId := getRandID()
	//拼装xml
	response := Response{
		Saml2p:       "urn:oasis:names:tc:SAML:2.0:protocol",
		Destination:  Destination,
		InResponseTo: ResponseTo,
		ID:           RandId,
		Version:      Version,
		IssueInstant: time.Now().Add(time.Hour * -8).Format("2006-01-02T15:04:05Z"),
		//	Ds:     "http://www.w3.org/2000/09/xmldsig#",
		Issuer: Issuer{
			Value: IssUerURL,
			Saml2: "urn:oasis:names:tc:SAML:2.0:assertion",
		},
		Status: Status{
			StatusCode: StatusCode{
				Value: "urn:oasis:names:tc:SAML:2.0:status:Success",
			},
		},
		Assertion: Assertion{
			Issuer:       IssUerURL,
			Saml2:        "urn:oasis:names:tc:SAML:2.0:assertion",
			ID:           RandId,
			Version:      Version,
			IssueInstant: time.Now().Add(time.Hour * -8).Format("2006-01-02T15:04:05Z"),
			Signature: Signature{
				Id: "placeholder",
				Ds: "http://www.w3.org/2000/09/xmldsig#",
				SignedInfo: SignedInfo{
					CanonicalizationMethod: CanonicalizationMethod{
						Algorithm: "http://www.w3.org/TR/2001/REC-xml-c14n-20010315",
					},
					SignatureMethod: SignatureMethod{
						Algorithm: "http://www.w3.org/2001/04/xmldsig-more#rsa-sha256",
					},
					Reference: Reference{
						URI: "#" + RandId,
						Transforms: Transforms{
							Transform: Transform{
								Algorithm: "http://www.w3.org/2000/09/xmldsig#enveloped-signature",
							},
						},
						DigestMethod: DigestMethod{
							Algorithm: "http://www.w3.org/2001/04/xmldsig-more#rsa-sha256",
						},
						DigestValue: "",
					},
				},
			},
			Subject: Subject{
				NameID: NameID{
					Value:  nameId,
					Format: "urn:oasis:names:tc:SAML:1.1:nameid-format:unspecified",
				},
				SubjectConfirmation: SubjectConfirmation{
					Method: "urn:oasis:names:tc:SAML:2.0:cm:bearer",
					SubjectConfirmationData: SubjectConfirmationData{
						InResponseTo: ResponseTo,
						NotOnOrAfter: ExpirationTime,
						Recipient:    RecipientAddress,
					},
				},
			},
			Conditions: Conditions{
				NotBefore:    time.Now().Add(time.Hour * -8).Format("2006-01-02T15:04:05Z"),
				NotOnOrAfter: ExpirationTime,
				AudienceRestriction: AudienceRestriction{
					Audience: Audience,
				},
			},
			Authentication: Authentication{
				AuthnInstant: time.Now().Add(time.Hour * -8).Format("2006-01-02T15:04:05Z"),
				SessionIndex: RandId,
				AuthnContext: AuthnContext{
					AuthnContextClassRef: "urn:oasis:names:tc:SAML:2.0:ac:classes:PasswordProtectedTransport",
				},
			},
		},
	}

	// 计算摘要
	assertionXML, _ := xml.MarshalIndent(response.Assertion, "", "")
	digest := sha256.Sum256(assertionXML)
	digestStr := base64.StdEncoding.EncodeToString(digest[:])
	response.Assertion.Signature.SignedInfo.Reference.DigestValue = digestStr

	// 对SignedInfo进行签名
	signedInfoData, err := xml.Marshal(response.Assertion.Signature.SignedInfo)
	if err != nil {
		fmt.Println("XML marshaling error:", err)
		return nil
	}
	// hashed处理
	signedInfoDigest := sha256.Sum256(signedInfoData)
	fmt.Println(string(signedInfoDigest[:]))
	// 读取私钥
	privateKeyFile, err := os.ReadFile("static/file/private-key.pem")
	if err != nil {
		fmt.Println("Private key file read error:", err)
		return nil
	}
	// 解析私钥
	privateKeyBlock, _ := pem.Decode(privateKeyFile)
	if privateKeyBlock == nil || privateKeyBlock.Type != "RSA PRIVATE KEY" {
		fmt.Println("Private key decode error")
		return nil
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(privateKeyBlock.Bytes)
	if err != nil {
		fmt.Println("Private key parse error:", err)
		return nil
	}
	signatureValue, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, signedInfoDigest[:])
	if err != nil {
		fmt.Println("SignPKCS1v15 error:", err)
		return nil
	}
	//生成xml
	response.Assertion.Signature.SignatureValue = base64.StdEncoding.EncodeToString(signatureValue)
	//fmt.Println(response.Assertion.Signature.SignatureValue)

	xmlData, err := xml.MarshalIndent(response, "", "")
	if err != nil {
		fmt.Println("Error marshaling XML:", err)
		return nil
	}
	//os.Stdout.Write(xmlData)
	fmt.Println(response.IssueInstant)
	data := base64.StdEncoding.EncodeToString(xmlData)
	return &data
}

func getRandID() string {
	randId := uuid.New()
	return fmt.Sprintf("_%s", randId.String())
}
