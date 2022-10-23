package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	api "goauthentik.io/api/v3"
)

func resourceOutpostProvider() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceOutpostProviderCreate,
		ReadContext:   resourceOutpostProviderRead,
		DeleteContext: resourceOutpostProviderDelete,
		Schema: map[string]*schema.Schema{
			"outpost": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"protocol_provider": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func setOutpostProviders(ctx context.Context, d *schema.ResourceData, c *APIClient, providers []int32) diag.Diagnostics {
	outpostReq := &api.PatchedOutpostRequest{}
	outpostReq.SetProviders(providers)
	req := c.client.OutpostsApi.OutpostsInstancesPartialUpdate(ctx, d.Id())
	req = req.PatchedOutpostRequest(*outpostReq)
	if _, hr, err := req.Execute(); err != nil {
		return httpToDiag(d, hr, err)
	}
	return diag.Diagnostics{}
}

func resourceOutpostProviderCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*APIClient)

	outpost, ok := d.GetOk("outpost")
	if !ok || outpost == "" {
		return diag.Errorf("No outpost was provided")
	}

	provider, ok := d.GetOk("protocol_provider")
	if !ok || provider == nil {
		return diag.Errorf("No protocol_provider was provided")
	}

	res, hr, err := c.client.OutpostsApi.OutpostsInstancesRetrieve(ctx, outpost.(string)).Execute()
	if err != nil {
		return httpToDiag(d, hr, err)
	}

	providers := make([]int32, 0, len(res.Providers)+1)
	found := false
	for _, value := range res.Providers {
		if provider.(int) == int(value) {
			found = true
		}
		providers = append(providers, value)
	}

	d.SetId(res.Pk)
	if !found {
		providers = append(providers, int32(provider.(int)))
		return setOutpostProviders(ctx, d, c, providers)
	}

	return diag.Diagnostics{}
}

func resourceOutpostProviderRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*APIClient)

	provider, ok := d.GetOk("protocol_provider")
	if !ok || provider == nil {
		return diag.Errorf("No protocol_provider was provided")
	}

	res, hr, err := c.client.OutpostsApi.OutpostsInstancesRetrieve(ctx, d.Id()).Execute()
	if err != nil {
		return httpToDiag(d, hr, err)
	}

	found := false
	for _, value := range res.Providers {
		if provider.(int) == int(value) {
			found = true
		}
	}

	if !found {
		d.SetId("")
	}

	return diag.Diagnostics{}
}

func resourceOutpostProviderDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*APIClient)

	provider, ok := d.GetOk("protocol_provider")
	if !ok || provider == nil {
		return diag.Errorf("No protocol_provider was provided")
	}

	res, hr, err := c.client.OutpostsApi.OutpostsInstancesRetrieve(ctx, d.Id()).Execute()
	if err != nil {
		return httpToDiag(d, hr, err)
	}

	providers := make([]int32, 0, len(res.Providers))
	found := false
	for _, value := range res.Providers {
		if provider != int(value) {
			found = true
			providers = append(providers, value)
		}
	}

	if found {
		return setOutpostProviders(ctx, d, c, providers)
	}

	return diag.Diagnostics{}
}
