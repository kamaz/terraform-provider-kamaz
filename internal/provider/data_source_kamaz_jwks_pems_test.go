package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var certsCheck = resource.ComposeTestCheckFunc(
	resource.TestCheckResourceAttr(
		"data.kamaz_jwks_pems.foo", "certs.0.cert_kid", "462949174f1eedf4f9f9434877be483b324140f5",
	),
	resource.TestCheckResourceAttr(
		"data.kamaz_jwks_pems.foo", "certs.0.cert_pem", "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA2BHFUUq8NqZ3pxxi/RJc\nSIMG5nJoZQ8Nbvf+lW5o7hJ9CmLA4SeUmDL2IVK6CSuskTPj/ohAp/gtOg3PCJvn\n33grPoJQu38MoMB8kDqA4U+u3A86GGEjWtk6LPo7dEkojZNQkzhZCnEMTuRMtBZX\nsLWNGJpY3UADA3rxnHnBP1wrSt27iXIE0C6+1N5z00R13r3L0aWC0MuAUgjI2H4d\nGMr8B3niJ+NjOVPCwG7xSWsCwsSitAuhPGHaDtenB23ZsFJjbuTuiguoSJ9A1qo9\nkzBOg32xda4derbWasu7Tk8p53PFxXDJGR/h7dM+nsJHl7lAUDqL8zOrf9XXlPTj\nwQIDAQAB\n-----END PUBLIC KEY-----\n",
	),
	resource.TestCheckResourceAttr(
		"data.kamaz_jwks_pems.foo", "certs.1.cert_kid", "6ef4bd908591f697a8a9b893b03e6a77eb04e51f",
	),
	resource.TestCheckResourceAttr(
		"data.kamaz_jwks_pems.foo", "certs.1.cert_pem", "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAxkgm0jU0J7SgrmmuLypj\nWO6J9MlF9vpRpsw84sme4EtWMUyAu4zT+X9Ten5wB9W2z0Gft5QOmFL99ueP3MeO\nqZsXGwW2UWVuQCpkD0bo4qDDqwbt8Cl31Qjb5RHeuvmwYpNQK/1ppb6dwlUCA2Y9\nAaE7UsZITlR7r5XiBNvOEZh0LTsjPcikCheAs6nPSMBbdIeM28vii1PgPYTU6x6d\nRBVBAExaRnRDPZZh4acgfKIpbOCMJm2tucqwYhx3Wr5Lhu56oZALK4lvP9SAgOZd\nG3BA48PKIdLOeiTP+DI/pHJhIn1N5lMCcmcpG3OKMvWo0tFMOGj8Or+mHqB/5I+L\n4wIDAQAB\n-----END PUBLIC KEY-----\n",
	),
	resource.TestCheckNoResourceAttr(
		"data.kamaz_jwks_pems.foo", "certs.2",
	),
	resource.TestCheckResourceAttr(
		"data.kamaz_jwks_pems.foo", "json", `{"certs":[{"pem":"-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA2BHFUUq8NqZ3pxxi/RJc\nSIMG5nJoZQ8Nbvf+lW5o7hJ9CmLA4SeUmDL2IVK6CSuskTPj/ohAp/gtOg3PCJvn\n33grPoJQu38MoMB8kDqA4U+u3A86GGEjWtk6LPo7dEkojZNQkzhZCnEMTuRMtBZX\nsLWNGJpY3UADA3rxnHnBP1wrSt27iXIE0C6+1N5z00R13r3L0aWC0MuAUgjI2H4d\nGMr8B3niJ+NjOVPCwG7xSWsCwsSitAuhPGHaDtenB23ZsFJjbuTuiguoSJ9A1qo9\nkzBOg32xda4derbWasu7Tk8p53PFxXDJGR/h7dM+nsJHl7lAUDqL8zOrf9XXlPTj\nwQIDAQAB\n-----END PUBLIC KEY-----\n","kid":"462949174f1eedf4f9f9434877be483b324140f5"},{"pem":"-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAxkgm0jU0J7SgrmmuLypj\nWO6J9MlF9vpRpsw84sme4EtWMUyAu4zT+X9Ten5wB9W2z0Gft5QOmFL99ueP3MeO\nqZsXGwW2UWVuQCpkD0bo4qDDqwbt8Cl31Qjb5RHeuvmwYpNQK/1ppb6dwlUCA2Y9\nAaE7UsZITlR7r5XiBNvOEZh0LTsjPcikCheAs6nPSMBbdIeM28vii1PgPYTU6x6d\nRBVBAExaRnRDPZZh4acgfKIpbOCMJm2tucqwYhx3Wr5Lhu56oZALK4lvP9SAgOZd\nG3BA48PKIdLOeiTP+DI/pHJhIn1N5lMCcmcpG3OKMvWo0tFMOGj8Or+mHqB/5I+L\n4wIDAQAB\n-----END PUBLIC KEY-----\n","kid":"6ef4bd908591f697a8a9b893b03e6a77eb04e51f"}]}`,
	),
)

func TestAccKamazJwksPems(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccKamazJwksPemsHeredoc,
				Check:  certsCheck,
			},
			{
				Config: testAccKamazJwksPemsJsonencode,
				Check:  certsCheck,
			},
			{
				Config: testAccKamazJwksPemsJwksUri,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(
						"data.kamaz_jwks_pems.foo", "certs.0.cert_kid",
					),
					resource.TestCheckResourceAttrSet(
						"data.kamaz_jwks_pems.foo", "certs.0.cert_pem",
					),
					resource.TestCheckResourceAttrSet(
						"data.kamaz_jwks_pems.foo", "json",
					),
				),
			},
		},
	})
}

const testAccKamazJwksPemsHeredoc = `
data "kamaz_jwks_pems" "foo" {
  jwks = <<EOT
{
	"keys": [
	  {
	    "n": "2BHFUUq8NqZ3pxxi_RJcSIMG5nJoZQ8Nbvf-lW5o7hJ9CmLA4SeUmDL2IVK6CSuskTPj_ohAp_gtOg3PCJvn33grPoJQu38MoMB8kDqA4U-u3A86GGEjWtk6LPo7dEkojZNQkzhZCnEMTuRMtBZXsLWNGJpY3UADA3rxnHnBP1wrSt27iXIE0C6-1N5z00R13r3L0aWC0MuAUgjI2H4dGMr8B3niJ-NjOVPCwG7xSWsCwsSitAuhPGHaDtenB23ZsFJjbuTuiguoSJ9A1qo9kzBOg32xda4derbWasu7Tk8p53PFxXDJGR_h7dM-nsJHl7lAUDqL8zOrf9XXlPTjwQ",
	    "e": "AQAB",
	    "use": "sig",
	    "alg": "RS256",
	    "kty": "RSA",
	    "kid": "462949174f1eedf4f9f9434877be483b324140f5"
	  },
	  {
	    "alg": "RS256",
	    "use": "sig",
	    "n": "xkgm0jU0J7SgrmmuLypjWO6J9MlF9vpRpsw84sme4EtWMUyAu4zT-X9Ten5wB9W2z0Gft5QOmFL99ueP3MeOqZsXGwW2UWVuQCpkD0bo4qDDqwbt8Cl31Qjb5RHeuvmwYpNQK_1ppb6dwlUCA2Y9AaE7UsZITlR7r5XiBNvOEZh0LTsjPcikCheAs6nPSMBbdIeM28vii1PgPYTU6x6dRBVBAExaRnRDPZZh4acgfKIpbOCMJm2tucqwYhx3Wr5Lhu56oZALK4lvP9SAgOZdG3BA48PKIdLOeiTP-DI_pHJhIn1N5lMCcmcpG3OKMvWo0tFMOGj8Or-mHqB_5I-L4w",
	    "e": "AQAB",
	    "kid": "6ef4bd908591f697a8a9b893b03e6a77eb04e51f",
	    "kty": "RSA"
	  }
	]
      }
EOT
}
`

const testAccKamazJwksPemsJsonencode = `
data "kamaz_jwks_pems" "foo" {
  jwks = jsonencode({
	  keys = [
	  {
	    n = "2BHFUUq8NqZ3pxxi_RJcSIMG5nJoZQ8Nbvf-lW5o7hJ9CmLA4SeUmDL2IVK6CSuskTPj_ohAp_gtOg3PCJvn33grPoJQu38MoMB8kDqA4U-u3A86GGEjWtk6LPo7dEkojZNQkzhZCnEMTuRMtBZXsLWNGJpY3UADA3rxnHnBP1wrSt27iXIE0C6-1N5z00R13r3L0aWC0MuAUgjI2H4dGMr8B3niJ-NjOVPCwG7xSWsCwsSitAuhPGHaDtenB23ZsFJjbuTuiguoSJ9A1qo9kzBOg32xda4derbWasu7Tk8p53PFxXDJGR_h7dM-nsJHl7lAUDqL8zOrf9XXlPTjwQ",
	    e = "AQAB",
	    use = "sig",
	    alg = "RS256",
	    kty = "RSA",
	    kid = "462949174f1eedf4f9f9434877be483b324140f5"
	  },
	  {
	    alg = "RS256",
	    use = "sig",
	    n = "xkgm0jU0J7SgrmmuLypjWO6J9MlF9vpRpsw84sme4EtWMUyAu4zT-X9Ten5wB9W2z0Gft5QOmFL99ueP3MeOqZsXGwW2UWVuQCpkD0bo4qDDqwbt8Cl31Qjb5RHeuvmwYpNQK_1ppb6dwlUCA2Y9AaE7UsZITlR7r5XiBNvOEZh0LTsjPcikCheAs6nPSMBbdIeM28vii1PgPYTU6x6dRBVBAExaRnRDPZZh4acgfKIpbOCMJm2tucqwYhx3Wr5Lhu56oZALK4lvP9SAgOZdG3BA48PKIdLOeiTP-DI_pHJhIn1N5lMCcmcpG3OKMvWo0tFMOGj8Or-mHqB_5I-L4w",
	    e = "AQAB",
	    kid = "6ef4bd908591f697a8a9b893b03e6a77eb04e51f",
	    kty = "RSA"
	  }
	  ]
  })
}
`

const testAccKamazJwksPemsJwksUri = `
data "kamaz_jwks_pems" "foo" {
  jwks_uri = "https://www.googleapis.com/oauth2/v3/certs"
}
`
