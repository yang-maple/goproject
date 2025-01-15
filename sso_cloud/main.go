package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sso_cloud/ldap"
	"sso_cloud/saml"
)

func main() {
	//var ids string
	var ids, relay string
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "./static")
	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	router.GET("/sso/saml/login", func(c *gin.Context) {
		params := new(struct {
			SAMLRequest string `form:"SAMLRequest"`
			RelayState  string `form:"RelayState"`
		})
		err := c.ShouldBind(&params)
		if err != nil {
			c.JSON(400, gin.H{
				"code": 400,
				"msg":  err.Error(),
			})
		}
		ids = *(saml.XmlRequest(params.SAMLRequest))
		relay = params.RelayState
		c.HTML(http.StatusOK, "login.html", gin.H{
			"title": "Main website",
		})
	})
	router.POST("/signLogin", func(c *gin.Context) {
		parms := new(struct {
			Username string `form:"username"`
			Password string `form:"password"`
		})
		err := c.ShouldBind(&parms)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": http.StatusBadRequest,
				"msg":  err.Error(),
			})
		}
		if parms.Username == "" || parms.Password == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": http.StatusBadRequest,
				"msg":  "Username or password is empty!",
			})
			return
		}
		err = ldap.Authentication(parms.Username, parms.Password)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": http.StatusBadRequest,
				"msg":  "The Username or Password are incorrect!",
			})
			return
		}
		//c.JSON(200, gin.H{
		//	"code": 200,
		//	"msg":  *saml.XmlResponse(parms.Username, ids, "1656090720688399"),
		//})
		// 登录成功后，跳转到首页
		c.HTML(http.StatusOK, "successlogin.html", gin.H{
			"title": "Main website",
			"xml":   *saml.XmlResponse(parms.Username, ids, "1656090720688399"),
			"relay": relay,
		})
		//str := `<samlp:Response ID="_e1ba6c8a-fa0c-4185-b474-3965cfcdc92d" Version="2.0" IssueInstant="2024-07-25T08:39:57.984Z" Destination="https://signin.aliyun.com/saml/SSO" InResponseTo="a11afb7diejabe744c64587h3b94681" xmlns:samlp="urn:oasis:names:tc:SAML:2.0:protocol">`
		//str += `<saml2:Issuer>http://127.0.0.1:8000/sso/saml/login</saml2:Issuer>`
		//str += `<saml2p:Status></saml2p:Status>`
		//str += `<saml2p:AuthnStatement>`
		//str += `<saml2:Assertion>`
		//str += `<saml2:Issuer>http://127.0.0.1:8000/sso/saml/login</saml2:Issuer>`
		//str += `<ds:Signature>MIIDujCCAqKgAwIBAgIJAKtb15tbL/91MA0GCSqGSIb3DQEBCwUAMGsxCzAJBgNVBAYTAlhZMQwwCgYDVQQIDANYSVkxFTATBgNVBAcMDERlZmF1bHQgQ2l0eTEcMBoGA1UECgwTRGVmYXVsdCBDb21wYW55IEx0ZDEZMBcGCSqGSIb3DQEJARYKMTI0QHFxLmNvbTAeFw0yNDA1MjcwMzQ4MTBaFw0zNDA1MjUwMzQ4MTBaMGsxCzAJBgNVBAYTAlhZMQwwCgYDVQQIDANYSVkxFTATBgNVBAcMDERlZmF1bHQgQ2l0eTEcMBoGA1UECgwTRGVmYXVsdCBDb21wYW55IEx0ZDEZMBcGCSqGSIb3DQEJARYKMTI0QHFxLmNvbTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBANJsdxyh7TCX9lcTuV57B3D7wQ+nQ7cZSiW+pbwcti0gi/52KY+ebWP0DJGIWbz2z3OxhfkghhtK64Ex/TpdO3qzRAmgpsTBEJEegTDaU+4wJBqOPTsVf1/2gP+wjWZkYlC5Wmvx6juBBzcIlgAcDexzNWxNh+xBPy7agF7CNDYfBbvwLZv0cotNCmifL1vB4woxAUD62uKRBbCTlWWiZDccTk6c5ukgYdcUQSVshUpG4LNmR1/aFOGPC2fKmePilPVMYWDjfOdL9trmUuGB105jqI3IT9ulbu9P87K880YZZ+Wflc6B2bzfpjOpiu6DO6lPyGSpBGHPyezH9QQ2NkkCAwEAAaNhMF8wHQYDVR0OBBYEFG+7CWZYqOWEE4Lp6WPiDZ3/cWH5MB8GA1UdIwQYMBaAFG+7CWZYqOWEE4Lp6WPiDZ3/cWH5MA8GA1UdEQQIMAaHBAoBgx8wDAYDVR0TBAUwAwEB/zANBgkqhkiG9w0BAQsFAAOCAQEAGzygfrxxUM4MVL3yWOSn4qW+FepdVxZtdckXeKcb9BB33O9VH1hXP3o1J26Pqnu6aKIA8xsEInMon2Z23SCeyufreUzaK1AIcSlN9z0LsTHt72kAAajmLuWUUkV40am4/uB6H2kM750Q1YIE45j2acJ/Rp1V2S6KEnXv8it6wLItLDhs3LUVRvjcm6gGaAgRevt6L2TH9zy4KdsWTfsi2OCMPQ9KDAxcILyP9ZCZT61ek84uH0fAlkmqisLwyyMyRAJ1Xf3UHVyuAFkj0dT2p2zl9vFiOQrz6G5BOWWCGApZlzxvvJuivN+mXA34c9GZ60M/nIhocSXDSPEtDMGNgQ==</ds:Signature>`
		//str += `<saml2:Subject>`
		//str += `<saml2:NameID>xiy19425_admin</saml2:NameID>`
		//str += `<saml2:SubjectConfirmation>`
		//str += `<saml2:SubjectConfirmationData NotOnOrAfter="2024-07-25T17:36:07Z" Recipient="https://signin.aliyun.com/saml/SSO"></saml2:SubjectConfirmationData>`
		//str += `</saml2:SubjectConfirmation>`
		//str += `</saml2:Subject>`
		//str += `<saml2:Conditions>`
		//str += `<saml2:AudienceRestriction>`
		//str += `<saml2:Audience>https://signin.aliyun.com/saml/SSO</saml2:Audience>`
		//str += `</saml2:AudienceRestriction>`
		//str += `</saml2:Conditions>`
		//str += `<saml2:AuthnStatement></saml2:AuthnStatement>`
		//str += `</saml2:Assertion>`
		//str += `</saml2:Authn>`
		//str += `</saml2p:Response>`
		//// 登录成功后，跳转到首页
		//c.JSON(http.StatusOK, gin.H{
		//	"code":   200,
		//	"newxml": base64.StdEncoding.EncodeToString([]byte(str)),
		//})
	})
	router.Run(":8000")
}
