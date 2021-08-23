package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccKamazJwkPem(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccKamazJwkPemHeredoc,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.kamaz_jwk_pem.foo", "cert_kid", "462949174f1eedf4f9f9434877be483b324140f5"),
					resource.TestCheckResourceAttr(
						"data.kamaz_jwk_pem.foo", "cert_pem", "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA2BHFUUq8NqZ3pxxi/RJc\nSIMG5nJoZQ8Nbvf+lW5o7hJ9CmLA4SeUmDL2IVK6CSuskTPj/ohAp/gtOg3PCJvn\n33grPoJQu38MoMB8kDqA4U+u3A86GGEjWtk6LPo7dEkojZNQkzhZCnEMTuRMtBZX\nsLWNGJpY3UADA3rxnHnBP1wrSt27iXIE0C6+1N5z00R13r3L0aWC0MuAUgjI2H4d\nGMr8B3niJ+NjOVPCwG7xSWsCwsSitAuhPGHaDtenB23ZsFJjbuTuiguoSJ9A1qo9\nkzBOg32xda4derbWasu7Tk8p53PFxXDJGR/h7dM+nsJHl7lAUDqL8zOrf9XXlPTj\nwQIDAQAB\n-----END PUBLIC KEY-----\n"),
					resource.TestCheckResourceAttr(
						"data.kamaz_jwk_pem.foo", "json", `{"pem":"-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA2BHFUUq8NqZ3pxxi/RJc\nSIMG5nJoZQ8Nbvf+lW5o7hJ9CmLA4SeUmDL2IVK6CSuskTPj/ohAp/gtOg3PCJvn\n33grPoJQu38MoMB8kDqA4U+u3A86GGEjWtk6LPo7dEkojZNQkzhZCnEMTuRMtBZX\nsLWNGJpY3UADA3rxnHnBP1wrSt27iXIE0C6+1N5z00R13r3L0aWC0MuAUgjI2H4d\nGMr8B3niJ+NjOVPCwG7xSWsCwsSitAuhPGHaDtenB23ZsFJjbuTuiguoSJ9A1qo9\nkzBOg32xda4derbWasu7Tk8p53PFxXDJGR/h7dM+nsJHl7lAUDqL8zOrf9XXlPTj\nwQIDAQAB\n-----END PUBLIC KEY-----\n","kid":"462949174f1eedf4f9f9434877be483b324140f5"}`),
				),
			},
			{
				Config: testAccKamazJwkPemJsonencode,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.kamaz_jwk_pem.foo", "cert_kid", "462949174f1eedf4f9f9434877be483b324140f5"),
					resource.TestCheckResourceAttr(
						"data.kamaz_jwk_pem.foo", "cert_pem", "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA2BHFUUq8NqZ3pxxi/RJc\nSIMG5nJoZQ8Nbvf+lW5o7hJ9CmLA4SeUmDL2IVK6CSuskTPj/ohAp/gtOg3PCJvn\n33grPoJQu38MoMB8kDqA4U+u3A86GGEjWtk6LPo7dEkojZNQkzhZCnEMTuRMtBZX\nsLWNGJpY3UADA3rxnHnBP1wrSt27iXIE0C6+1N5z00R13r3L0aWC0MuAUgjI2H4d\nGMr8B3niJ+NjOVPCwG7xSWsCwsSitAuhPGHaDtenB23ZsFJjbuTuiguoSJ9A1qo9\nkzBOg32xda4derbWasu7Tk8p53PFxXDJGR/h7dM+nsJHl7lAUDqL8zOrf9XXlPTj\nwQIDAQAB\n-----END PUBLIC KEY-----\n"),
					resource.TestCheckResourceAttr(
						"data.kamaz_jwk_pem.foo", "json", `{"pem":"-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA2BHFUUq8NqZ3pxxi/RJc\nSIMG5nJoZQ8Nbvf+lW5o7hJ9CmLA4SeUmDL2IVK6CSuskTPj/ohAp/gtOg3PCJvn\n33grPoJQu38MoMB8kDqA4U+u3A86GGEjWtk6LPo7dEkojZNQkzhZCnEMTuRMtBZX\nsLWNGJpY3UADA3rxnHnBP1wrSt27iXIE0C6+1N5z00R13r3L0aWC0MuAUgjI2H4d\nGMr8B3niJ+NjOVPCwG7xSWsCwsSitAuhPGHaDtenB23ZsFJjbuTuiguoSJ9A1qo9\nkzBOg32xda4derbWasu7Tk8p53PFxXDJGR/h7dM+nsJHl7lAUDqL8zOrf9XXlPTj\nwQIDAQAB\n-----END PUBLIC KEY-----\n","kid":"462949174f1eedf4f9f9434877be483b324140f5"}`),
				),
			},
		},
	})
}

const testAccKamazJwkPemHeredoc = `
data "kamaz_jwk_pem" "foo" {
  key = <<EOT
  {
    "n": "2BHFUUq8NqZ3pxxi_RJcSIMG5nJoZQ8Nbvf-lW5o7hJ9CmLA4SeUmDL2IVK6CSuskTPj_ohAp_gtOg3PCJvn33grPoJQu38MoMB8kDqA4U-u3A86GGEjWtk6LPo7dEkojZNQkzhZCnEMTuRMtBZXsLWNGJpY3UADA3rxnHnBP1wrSt27iXIE0C6-1N5z00R13r3L0aWC0MuAUgjI2H4dGMr8B3niJ-NjOVPCwG7xSWsCwsSitAuhPGHaDtenB23ZsFJjbuTuiguoSJ9A1qo9kzBOg32xda4derbWasu7Tk8p53PFxXDJGR_h7dM-nsJHl7lAUDqL8zOrf9XXlPTjwQ",
    "e": "AQAB",
    "use": "sig",
    "alg": "RS256",
    "kty": "RSA",
    "kid": "462949174f1eedf4f9f9434877be483b324140f5"
  }
EOT
}
`

const testAccKamazJwkPemJsonencode = `
data "kamaz_jwk_pem" "foo" {
  key = jsonencode({
    n = "2BHFUUq8NqZ3pxxi_RJcSIMG5nJoZQ8Nbvf-lW5o7hJ9CmLA4SeUmDL2IVK6CSuskTPj_ohAp_gtOg3PCJvn33grPoJQu38MoMB8kDqA4U-u3A86GGEjWtk6LPo7dEkojZNQkzhZCnEMTuRMtBZXsLWNGJpY3UADA3rxnHnBP1wrSt27iXIE0C6-1N5z00R13r3L0aWC0MuAUgjI2H4dGMr8B3niJ-NjOVPCwG7xSWsCwsSitAuhPGHaDtenB23ZsFJjbuTuiguoSJ9A1qo9kzBOg32xda4derbWasu7Tk8p53PFxXDJGR_h7dM-nsJHl7lAUDqL8zOrf9XXlPTjwQ",
    e = "AQAB",
    use = "sig",
    alg = "RS256",
    kty = "RSA",
    kid = "462949174f1eedf4f9f9434877be483b324140f5"
  })
}
`
