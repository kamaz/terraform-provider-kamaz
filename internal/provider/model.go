package provider

type PemKeyId struct {
	Pem   string `json:"pem"`
	KeyId string `json:"kid"`
}

type Certs struct {
	Certs []PemKeyId `json:"certs"`
}
