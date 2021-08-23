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
