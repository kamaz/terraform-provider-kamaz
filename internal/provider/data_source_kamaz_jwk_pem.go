package provider

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/lestrrat-go/jwx/jwk"
)

func dataSourceJwkPem() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "Datasource converts JWK to PEM.",

		ReadContext: dataSourceJwkPemRead,

		Schema: map[string]*schema.Schema{
			"key": {
				// This description is used by the documentation generator and the language server.
				Description:  "JWK key as json.",
				Type:         schema.TypeString,
				ValidateFunc: validation.StringIsJSON,
				Required:     true,
			},
			"cert_kid": {
				Description: "JWK key id (copies values from kid of original key value).",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"cert_pem": {
				Description: "PEM certificate from JWK",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"json": {
				Description: "JSON mapping between key id and pem",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

func dataSourceJwkPemRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	jwkKeyInput := d.Get("key").(string)

	jwkKey, err := jwk.ParseKey([]byte(jwkKeyInput))
	if err != nil {
		return diag.FromErr(err)
	}
	keyId := jwkKey.KeyID()
	d.Set("cert_kid", keyId)

	pemKey, err := jwk.Pem(jwkKey)
	if err != nil {
		return diag.FromErr(err)
	}
	d.Set("cert_pem", string(pemKey))

	pemKeyIdJson, err := json.Marshal(PemKeyId{
		Pem:   string(pemKey),
		KeyId: keyId,
	})
	if err != nil {
		return diag.FromErr(err)
	}
	d.Set("json", string(pemKeyIdJson))

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
