package provider

import (
	"context"

	"github.com/goauthentik/terraform-provider-authentik/api"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCertificateKeyPair() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceCertificateKeyPairCreate,
		ReadContext:   resourceCertificateKeyPairRead,
		UpdateContext: resourceCertificateKeyPairUpdate,
		DeleteContext: resourceCertificateKeyPairDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"certificate_data": {
				Type:     schema.TypeString,
				Required: true,
			},
			"key_data": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceCertificateKeyPairSchemaToModel(d *schema.ResourceData) (*api.CertificateKeyPairRequest, diag.Diagnostics) {
	app := api.CertificateKeyPairRequest{
		Name:            d.Get("name").(string),
		CertificateData: d.Get("certificate_data").(string),
	}

	if l, ok := d.Get("key_data").(string); ok {
		app.KeyData = &l
	}
	return &app, nil
}

func resourceCertificateKeyPairCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*APIClient)

	app, diags := resourceCertificateKeyPairSchemaToModel(d)
	if diags != nil {
		return diags
	}

	res, hr, err := c.client.CryptoApi.CryptoCertificatekeypairsCreate(ctx).CertificateKeyPairRequest(*app).Execute()
	if err != nil {
		return httpToDiag(hr)
	}

	d.SetId(res.Pk)
	return resourceCertificateKeyPairRead(ctx, d, m)
}

func resourceCertificateKeyPairRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*APIClient)

	res, hr, err := c.client.CryptoApi.CryptoCertificatekeypairsRetrieve(ctx, d.Id()).Execute()
	if err != nil {
		return httpToDiag(hr)
	}

	d.Set("name", res.Name)

	rc, hr, err := c.client.CryptoApi.CryptoCertificatekeypairsViewCertificateRetrieve(ctx, d.Id()).Execute()
	if err != nil {
		return httpToDiag(hr)
	}
	d.Set("certificate_data", rc.Data)

	rk, hr, err := c.client.CryptoApi.CryptoCertificatekeypairsViewPrivateKeyRetrieve(ctx, d.Id()).Execute()
	if err != nil {
		return httpToDiag(hr)
	}
	d.Set("key_data", rk.Data)

	return diags
}

func resourceCertificateKeyPairUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*APIClient)

	app, di := resourceCertificateKeyPairSchemaToModel(d)
	if di != nil {
		return di
	}

	res, hr, err := c.client.CryptoApi.CryptoCertificatekeypairsUpdate(ctx, d.Id()).CertificateKeyPairRequest(*app).Execute()
	if err != nil {
		return httpToDiag(hr)
	}

	d.SetId(res.Pk)
	return resourceCertificateKeyPairRead(ctx, d, m)
}

func resourceCertificateKeyPairDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*APIClient)
	hr, err := c.client.CryptoApi.CryptoCertificatekeypairsDestroy(ctx, d.Id()).Execute()
	if err != nil {
		return httpToDiag(hr)
	}
	return diag.Diagnostics{}
}
