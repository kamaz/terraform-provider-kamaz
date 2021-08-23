data "kamaz_jwks_pems" "foo" {
  jwks = jsonencode({
    keys = [
      {
        n   = "2BHFUUq8NqZ3pxxi_RJcSIMG5nJoZQ8Nbvf-lW5o7hJ9CmLA4SeUmDL2IVK6CSuskTPj_ohAp_gtOg3PCJvn33grPoJQu38MoMB8kDqA4U-u3A86GGEjWtk6LPo7dEkojZNQkzhZCnEMTuRMtBZXsLWNGJpY3UADA3rxnHnBP1wrSt27iXIE0C6-1N5z00R13r3L0aWC0MuAUgjI2H4dGMr8B3niJ-NjOVPCwG7xSWsCwsSitAuhPGHaDtenB23ZsFJjbuTuiguoSJ9A1qo9kzBOg32xda4derbWasu7Tk8p53PFxXDJGR_h7dM-nsJHl7lAUDqL8zOrf9XXlPTjwQ",
        e   = "AQAB",
        use = "sig",
        alg = "RS256",
        kty = "RSA",
        kid = "462949174f1eedf4f9f9434877be483b324140f5"
      },
      {
        alg = "RS256",
        use = "sig",
        n   = "xkgm0jU0J7SgrmmuLypjWO6J9MlF9vpRpsw84sme4EtWMUyAu4zT-X9Ten5wB9W2z0Gft5QOmFL99ueP3MeOqZsXGwW2UWVuQCpkD0bo4qDDqwbt8Cl31Qjb5RHeuvmwYpNQK_1ppb6dwlUCA2Y9AaE7UsZITlR7r5XiBNvOEZh0LTsjPcikCheAs6nPSMBbdIeM28vii1PgPYTU6x6dRBVBAExaRnRDPZZh4acgfKIpbOCMJm2tucqwYhx3Wr5Lhu56oZALK4lvP9SAgOZdG3BA48PKIdLOeiTP-DI_pHJhIn1N5lMCcmcpG3OKMvWo0tFMOGj8Or-mHqB_5I-L4w",
        e   = "AQAB",
        kid = "6ef4bd908591f697a8a9b893b03e6a77eb04e51f",
        kty = "RSA"
      }
    ]
  })
}
