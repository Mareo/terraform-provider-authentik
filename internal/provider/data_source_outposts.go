package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceOutposts() *schema.Resource {
	outpostSchema := map[string]*schema.Schema{}
	for k, v := range dataSourceOutpost().Schema {
		outpostSchema[k] = &schema.Schema{}
		*outpostSchema[k] = *v
		outpostSchema[k].Computed = true
		outpostSchema[k].Optional = false
		outpostSchema[k].Required = false
		outpostSchema[k].AtLeastOneOf = []string{}
		outpostSchema[k].ConflictsWith = []string{}
	}
	return &schema.Resource{
		ReadContext: dataSourceOutpostsRead,
		Description: "Get outposts list",
		Schema: map[string]*schema.Schema{
			"managed_icontains": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"managed_iexact": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name_icontains": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name_iexact": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ordering": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"providers_isnull": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"providers_by_pk": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
			"search": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"service_connection_name_icontains": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"service_connection_name_iexact": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"outposts": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: outpostSchema,
				},
			},
		},
	}
}

func dataSourceOutpostsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*APIClient)

	req := c.client.OutpostsApi.OutpostsInstancesList(ctx)

	for key := range dataSourceGroups().Schema {
		if v, ok := d.GetOk(key); ok {
			switch key {
			case "managed_icontains":
				req = req.ManagedIcontains(v.(string))
			case "managed_iexact":
				req = req.ManagedIexact(v.(string))
			case "name_icontains":
				req = req.NameIcontains(v.(string))
			case "name_iexact":
				req = req.NameIexact(v.(string))
			case "ordering":
				req = req.Ordering(v.(string))
			case "providers_isnull":
				req = req.ProvidersIsnull(v.(bool))
			case "providers_by_pk":
				providers := make([]int32, len(v.([]int)))
				for i, pk := range v.([]int) {
					providers[i] = int32(pk)
				}
				req = req.ProvidersByPk(providers)
			case "search":
				req = req.Search(v.(string))
			case "service_connection_name_icontains":
				req = req.ServiceConnectionNameIcontains(v.(string))
			case "service_connection_name_iexact":
				req = req.ServiceConnectionNameIexact(v.(string))
			}
		}
	}

	outposts := make([]map[string]interface{}, 0)
	for page := int32(1); true; page++ {
		req = req.Page(page)
		res, hr, err := req.Execute()
		if err != nil {
			return httpToDiag(d, hr, err)
		}

		for _, outpostRes := range res.Results {
			u, err := mapFromOutpost(outpostRes)
			if err != nil {
				return diag.FromErr(err)
			}
			outposts = append(outposts, u)
		}

		if res.Pagination.Next == 0 {
			break
		}
	}

	d.SetId("0")
	setWrapper(d, "outposts", outposts)
	return diag.Diagnostics{}
}
