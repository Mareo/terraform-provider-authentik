package provider

import (
	"context"
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"goauthentik.io/api/v3"
)

func dataSourceOutpost() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceOutpostRead,
		Description: "Get outpost by pk, name or managed value",
		Schema: map[string]*schema.Schema{
			"pk": {
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{"name", "managed"},
				AtLeastOneOf:  []string{"pk", "name", "managed"},
			},
			"name": {
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{"pk"},
				AtLeastOneOf:  []string{"pk", "name", "managed"},
			},
			"type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"protocol_providers": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
			"token_identifier": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"service_connection": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"config": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"managed": {
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{"pk"},
				AtLeastOneOf:  []string{"pk", "name", "managed"},
			},
		},
	}
}

func mapFromOutpost(outpost api.Outpost) (map[string]interface{}, error) {
	m := map[string]interface{}{
		"pk":                 outpost.GetPk(),
		"name":               outpost.GetName(),
		"type":               string(outpost.GetType()),
		"protocol_providers": outpost.GetProviders(),
		"service_connection": outpost.GetServiceConnection(),
		"token_identifier":   outpost.GetTokenIdentifier(),
		"config":             "",
		"managed":            outpost.GetManaged(),
	}

	b, err := json.Marshal(outpost.GetConfig())
	if err != nil {
		return nil, err
	}
	m["config"] = string(b)

	return m, nil
}

func setOutpost(data *schema.ResourceData, outpost api.Outpost) diag.Diagnostics {
	m, err := mapFromOutpost(outpost)
	if err != nil {
		return diag.FromErr(err)
	}
	for key, value := range m {
		switch key {
		case "pk":
			data.SetId(value.(string))
			setWrapper(data, key, value.(string))
		case "protocol_providers":
			setWrapper(data, key, value.([]int32))
		default:
			setWrapper(data, key, value.(string))
		}
	}
	return diag.Diagnostics{}
}

func dataSourceOutpostReadByPk(ctx context.Context, d *schema.ResourceData, c *APIClient, pk string) diag.Diagnostics {
	req := c.client.OutpostsApi.OutpostsInstancesRetrieve(ctx, pk)

	res, hr, err := req.Execute()
	if err != nil {
		return httpToDiag(d, hr, err)
	}

	return setOutpost(d, *res)
}

func dataSourceOutpostReadByFilters(ctx context.Context, d *schema.ResourceData, c *APIClient) diag.Diagnostics {
	req := c.client.OutpostsApi.OutpostsInstancesList(ctx)

	if n, ok := d.GetOk("name"); ok {
		req = req.NameIexact(n.(string))
	}

	if m, ok := d.GetOk("managed"); ok {
		req = req.ManagedIexact(m.(string))
	}

	res, hr, err := req.Execute()
	if err != nil {
		return httpToDiag(d, hr, err)
	}

	if len(res.Results) < 1 {
		return diag.Errorf("No matching outposts found")
	}

	if len(res.Results) > 1 {
		return diag.Errorf("Multiple outposts found")
	}

	return setOutpost(d, res.Results[0])
}

func dataSourceOutpostRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*APIClient)

	if pk, ok := d.GetOk("pk"); ok {
		return dataSourceOutpostReadByPk(ctx, d, c, pk.(string))
	}

	return dataSourceOutpostReadByFilters(ctx, d, c)
}
