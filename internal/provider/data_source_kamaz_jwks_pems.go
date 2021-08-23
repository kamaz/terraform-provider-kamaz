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

func dataSourceJwksPems() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "Datasource converts JWKS to PEMs.",

		ReadContext: dataSourceJwksPemsRead,

		Schema: map[string]*schema.Schema{
			"jwks": {
				Description:  "A set of keys following JWK Set Format (https://datatracker.ietf.org/doc/html/rfc7517#page-10).",
				Type:         schema.TypeString,
				ValidateFunc: validation.StringIsJSON,
				Optional:     true,
				ExactlyOneOf: []string{"jwks", "jwks_uri"},
			},
			"jwks_uri": {
				Description:  "A url to JWK Set",
				Type:         schema.TypeString,
				ValidateFunc: validation.IsURLWithHTTPS,
				Optional:     true,
				ExactlyOneOf: []string{"jwks", "jwks_uri"},
			},
			"certs": {
				Description: "List of pem certificates generated from jwks.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cert_kid": {
							Description: "JWK key id (copies values from kid of original key value).",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"cert_pem": {
							Description: "PEM certificate from JWK.",
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
			"json": {
				Description: "JSON mapping between key id and pem for JWK Set",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

func createJwks(d *schema.ResourceData) (jwk.Set, error) {
	jwksInput := d.Get("jwks").(string)
	jwksUri := d.Get("jwks_uri").(string)

	if jwksUri != "" {
		return jwk.Fetch(context.Background(), jwksUri)
	}

	return jwk.Parse([]byte(jwksInput))
}

func dataSourceJwksPemsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	jwks, err := createJwks(d)
	if err != nil {
		return diag.FromErr(err)
	}

	var pems []PemKeyId
	var pemsMap []map[string]interface{}

	for it := jwks.Iterate(context.Background()); it.Next(context.Background()); {
		pair := it.Pair()
		key := pair.Value.(jwk.Key)
		pem, err := jwk.Pem(key)

		if err != nil {
			return diag.FromErr(err)
		}

		pems = append(pems, PemKeyId{
			Pem:   string(pem),
			KeyId: key.KeyID(),
		})
		pemsMap = append(pemsMap, map[string]interface{}{
			"cert_pem": string(pem),
			"cert_kid": key.KeyID(),
		})
	}

	d.Set("certs", pemsMap)

	pemKeyIdJson, err := json.Marshal(Certs{
		Certs: pems,
	})
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("json", string(pemKeyIdJson))

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
